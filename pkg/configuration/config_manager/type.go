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

package configmanager

import (
	"context"

	"github.com/fsnotify/fsnotify"

	appsv1 "github.com/apecloud/kubeblocks/apis/apps/v1"
	appsv1beta1 "github.com/apecloud/kubeblocks/apis/apps/v1beta1"
)

type ConfigHandler interface {
	OnlineUpdate(ctx context.Context, name string, updatedParams map[string]string) error
	VolumeHandle(ctx context.Context, event fsnotify.Event) error
	MountPoint() []string
}

type ConfigSpecInfo struct {
	*appsv1beta1.ReloadAction `json:",inline"`

	ReloadType      appsv1beta1.DynamicReloadType `json:"reloadType"`
	ConfigSpec      appsv1.ComponentConfigSpec    `json:"configSpec"`
	FormatterConfig appsv1beta1.FileFormatConfig  `json:"formatterConfig"`

	DownwardAPIOptions []appsv1beta1.DownwardAPIChangeTriggeredAction `json:"downwardAPIOptions"`

	// config volume mount path
	MountPoint string `json:"mountPoint"`
	TPLConfig  string `json:"tplConfig"`
}

type ConfigSpecMeta struct {
	ConfigSpecInfo `json:",inline"`

	ScriptConfig   []appsv1beta1.ScriptConfig
	ToolsImageSpec *appsv1beta1.ToolsSetup
}

type TPLScriptConfig struct {
	Scripts   string `json:"scripts"`
	FileRegex string `json:"fileRegex"`
	DataType  string `json:"dataType"`
	DSN       string `json:"dsn"`

	FormatterConfig appsv1beta1.FileFormatConfig `json:"formatterConfig"`
}

type ConfigLazyRenderedMeta struct {
	*appsv1.ComponentConfigSpec `json:",inline"`

	// secondary template path
	Templates       []string                     `json:"templates"`
	FormatterConfig appsv1beta1.FileFormatConfig `json:"formatterConfig"`
}
