/*
Copyright (C) 2022-2025 ApeCloud Co., Ltd

This file is part of KubeBlocks project

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package controllerutil

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"slices"

	"github.com/StudioSol/set"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	appsv1 "github.com/apecloud/kubeblocks/apis/apps/v1"
	"github.com/apecloud/kubeblocks/apis/apps/v1alpha1"
	appsv1beta1 "github.com/apecloud/kubeblocks/apis/apps/v1beta1"
	parametersv1alpha1 "github.com/apecloud/kubeblocks/apis/parameters/v1alpha1"
	"github.com/apecloud/kubeblocks/pkg/configuration/core"
	"github.com/apecloud/kubeblocks/pkg/configuration/util"
	"github.com/apecloud/kubeblocks/pkg/configuration/validate"
	"github.com/apecloud/kubeblocks/pkg/constant"
)

type Result struct {
	Phase      v1alpha1.ConfigurationPhase `json:"phase"`
	Revision   string                      `json:"revision"`
	Policy     string                      `json:"policy"`
	ExecResult string                      `json:"execResult"`

	SucceedCount  int32 `json:"succeedCount"`
	ExpectedCount int32 `json:"expectedCount"`

	Retry   bool   `json:"retry"`
	Failed  bool   `json:"failed"`
	Message string `json:"message"`
}

// MergeAndValidateConfigs merges and validates configuration files
func MergeAndValidateConfigs(configConstraint appsv1beta1.ConfigConstraintSpec, baseConfigs map[string]string, cmKey []string, updatedParams []core.ParamPairs) (map[string]string, error) {
	var (
		err error
		fc  = configConstraint.FileFormatConfig

		newCfg         map[string]string
		configOperator core.ConfigOperator
		updatedKeys    = util.NewSet()
	)

	cmKeySet := core.FromCMKeysSelector(cmKey)
	configLoaderOption := core.CfgOption{
		Type:           core.CfgCmType,
		Log:            log.FromContext(context.TODO()),
		CfgType:        fc.Format,
		ConfigResource: core.FromConfigData(baseConfigs, cmKeySet),
	}
	if configOperator, err = core.NewConfigLoader(configLoaderOption); err != nil {
		return nil, err
	}

	// merge param to config file
	for _, params := range updatedParams {
		validUpdatedParameters := filterImmutableParameters(params.UpdatedParams, configConstraint.ImmutableParameters)
		if len(validUpdatedParameters) == 0 {
			continue
		}
		if err := configOperator.MergeFrom(validUpdatedParameters, core.NewCfgOptions(params.Key, core.WithFormatterConfig(fc))); err != nil {
			return nil, err
		}
		updatedKeys.Add(params.Key)
	}

	if newCfg, err = configOperator.ToCfgContent(); err != nil {
		return nil, core.WrapError(err, "failed to generate config file")
	}

	// The ToCfgContent interface returns the file contents of all keys, the configuration file is encoded and decoded into keys,
	// the content may be different with the original file, such as comments, blank lines, etc,
	// in order to minimize the impact on the original file, only update the changed part.
	updatedCfg := fromUpdatedConfig(newCfg, updatedKeys)
	if err = validate.NewConfigValidator(&configConstraint, validate.WithKeySelector(cmKey)).Validate(updatedCfg); err != nil {
		return nil, core.WrapError(err, "failed to validate updated config")
	}
	return core.MergeUpdatedConfig(baseConfigs, updatedCfg), nil
}

// fromUpdatedConfig filters out changed file contents.
func fromUpdatedConfig(m map[string]string, sets *set.LinkedHashSetString) map[string]string {
	if sets.Length() == 0 {
		return map[string]string{}
	}

	r := make(map[string]string, sets.Length())
	for key, v := range m {
		if sets.InArray(key) {
			r[key] = v
		}
	}
	return r
}

// IsApplyConfigChanged checks if the configuration is changed
func IsApplyConfigChanged(configMap *corev1.ConfigMap, item v1alpha1.ConfigurationItemDetail) bool {
	if configMap == nil {
		return false
	}

	lastAppliedVersion, ok := configMap.Annotations[constant.ConfigAppliedVersionAnnotationKey]
	if !ok {
		return false
	}
	var target v1alpha1.ConfigurationItemDetail
	if err := json.Unmarshal([]byte(lastAppliedVersion), &target); err != nil {
		return false
	}

	return reflect.DeepEqual(target, item)
}

// IsRerender checks if the configuration template is changed
func IsRerender(configMap *corev1.ConfigMap, item v1alpha1.ConfigurationItemDetail) bool {
	if configMap == nil {
		return true
	}
	if item.Version == "" && item.Payload.Data == nil && item.ImportTemplateRef == nil {
		return false
	}
	if version := configMap.Annotations[constant.CMConfigurationTemplateVersion]; version != item.Version {
		return true
	}

	var updatedVersion v1alpha1.ConfigurationItemDetail
	updatedVersionStr, ok := configMap.Annotations[constant.ConfigAppliedVersionAnnotationKey]
	if ok && updatedVersionStr != "" {
		if err := json.Unmarshal([]byte(updatedVersionStr), &updatedVersion); err != nil {
			return false
		}
	}
	return !reflect.DeepEqual(updatedVersion.Payload, item.Payload) ||
		!reflect.DeepEqual(updatedVersion.ImportTemplateRef, item.ImportTemplateRef)
}

// GetConfigSpecReconcilePhase gets the configuration phase
func GetConfigSpecReconcilePhase(configMap *corev1.ConfigMap,
	item v1alpha1.ConfigurationItemDetail,
	status *v1alpha1.ConfigurationItemDetailStatus) v1alpha1.ConfigurationPhase {
	if status == nil || status.Phase == "" {
		return v1alpha1.CCreatingPhase
	}
	if !IsApplyConfigChanged(configMap, item) {
		return v1alpha1.CPendingPhase
	}
	return status.Phase
}

func CheckAndPatchPayload(item *v1alpha1.ConfigurationItemDetail, payloadID string, payload interface{}) (bool, error) {
	if item == nil {
		return false, nil
	}
	if item.Payload.Data == nil {
		item.Payload.Data = make(map[string]interface{})
	}
	oldPayload, ok := item.Payload.Data[payloadID]
	if !ok && payload == nil {
		return false, nil
	}
	if payload == nil {
		delete(item.Payload.Data, payloadID)
		return true, nil
	}
	newPayload, err := buildPayloadAsUnstructuredObject(payload)
	if err != nil {
		return false, err
	}
	if oldPayload != nil && reflect.DeepEqual(oldPayload, newPayload) {
		return false, nil
	}
	item.Payload.Data[payloadID] = newPayload
	return true, nil
}

func buildPayloadAsUnstructuredObject(payload interface{}) (interface{}, error) {
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	var unstructuredObj any
	if err = json.Unmarshal(b, &unstructuredObj); err != nil {
		return nil, err
	}
	return unstructuredObj, nil
}

func ResourcesPayloadForComponent(resources corev1.ResourceRequirements) any {
	if len(resources.Requests) == 0 && len(resources.Limits) == 0 {
		return nil
	}

	return map[string]any{
		"limits":   resources.Limits,
		"requests": resources.Requests,
	}
}

func filterImmutableParameters(parameters map[string]any, immutableParams []string) map[string]any {
	if len(immutableParams) == 0 || len(parameters) == 0 {
		return parameters
	}

	validParameters := make(map[string]any, len(parameters))
	for key, val := range parameters {
		if !slices.Contains(immutableParams, key) {
			validParameters[key] = val
		}
	}
	return validParameters
}

func TransformConfigTemplate(configs []appsv1.ComponentConfigSpec) []appsv1.ComponentTemplateSpec {
	arr := make([]appsv1.ComponentTemplateSpec, 0, len(configs))
	for _, config := range configs {
		arr = append(arr, config.ComponentTemplateSpec)
	}
	return arr
}

func ResolveCmpdParametersDefs(ctx context.Context, reader client.Reader, cmpd *appsv1.ComponentDefinition) (*parametersv1alpha1.ParamConfigRenderer, []*parametersv1alpha1.ParametersDefinition, error) {
	var paramsDefs []*parametersv1alpha1.ParametersDefinition

	configRender, err := ResolveComponentConfigRender(ctx, reader, cmpd)
	if err != nil {
		return nil, nil, err
	}
	if configRender == nil || len(configRender.Spec.ParametersDefs) == 0 {
		return configRender, nil, nil
	}
	for _, defName := range configRender.Spec.ParametersDefs {
		paramsDef := &parametersv1alpha1.ParametersDefinition{}
		if err = reader.Get(ctx, client.ObjectKey{Name: defName}, paramsDef); err != nil {
			return nil, nil, err
		}
		if paramsDef.Status.Phase != parametersv1alpha1.PDAvailablePhase {
			return nil, nil, fmt.Errorf("the referenced ParametersDefinition is unavailable: %s", paramsDef.Name)
		}
		paramsDefs = append(paramsDefs, paramsDef)
	}
	return configRender, paramsDefs, nil
}

func ResolveComponentConfigRender(ctx context.Context, reader client.Reader, cmpd *appsv1.ComponentDefinition) (*parametersv1alpha1.ParamConfigRenderer, error) {
	configDefList := &parametersv1alpha1.ParamConfigRendererList{}
	if err := reader.List(ctx, configDefList); err != nil {
		return nil, err
	}

	checkAvailable := func(configDef parametersv1alpha1.ParamConfigRenderer) error {
		if configDef.Status.Phase != parametersv1alpha1.PDAvailablePhase {
			return fmt.Errorf("the referenced ParamConfigRenderer is unavailable: %s", configDef.Name)
		}
		return nil
	}

	for i, item := range configDefList.Items {
		if item.Spec.ComponentDef != cmpd.Name {
			continue
		}
		if item.Spec.ServiceVersion == "" || item.Spec.ServiceVersion == cmpd.Spec.ServiceVersion {
			return &configDefList.Items[i], checkAvailable(item)
		}
	}
	return nil, nil
}
