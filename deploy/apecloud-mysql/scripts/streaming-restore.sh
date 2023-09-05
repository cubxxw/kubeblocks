#! /bin/bash

if [ -z ${KB_LAST_COMPONENT_REPLICAS} ]; then
  exit 0
fi

component_replicas=${KB_COMPONENT_REPLICAS}
last_component_replicas=${KB_LAST_COMPONENT_REPLICAS}
if [ ${last_component_replicas} -ge ${component_replicas} ] || [ ${last_component_replicas} -eq 0 ]; then
  exit 0
fi

ordinal=${KB_POD_NAME##*-}
if [ ${ordinal} -lt ${last_component_replicas} ] || [ ${ordinal} -ge ${component_replicas} ]; then
  exit 0
fi

RESTORE_FILE=${DATA_VOLUME}/.xtrabackup_restore_done
if [ -f ${RESTORE_FILE} ]; then
  exit 0
fi

#while [ ! -f ${RESTORE_FILE} ];
#do
#    sleep 1
#done

checkrun() {
  "$@"
  local ret=$?
  if [ $ret -ne 0 ]; then
    exit $ret
  fi
}

RESTORE_TMP_DIR=${DATA_VOLUME}/restore-tmp
rm -rf ${RESTORE_TMP_DIR}
mkdir -p ${DATA_DIR} ${RESTORE_TMP_DIR}

cd ${RESTORE_TMP_DIR}
leader=${KB_LEADER}.${KB_CLUSTER_COMP_NAME}-headless
checkrun nc ${leader} 3502 | xbstream -x

checkrun xtrabackup --decompress  --target-dir=${RESTORE_TMP_DIR}
checkrun xtrabackup --prepare --target-dir=${RESTORE_TMP_DIR}
find . -name "*.qp" | xargs rm -f
checkrun xtrabackup --move-back --target-dir=${RESTORE_TMP_DIR} --datadir=${DATA_DIR}/ --log-bin=${LOG_BIN}

cd ${DATA_VOLUME}
rm -rf ${RESTORE_TMP_DIR}
touch ${RESTORE_FILE}
touch ${DATA_DIR}/.xtrabackup_restore
chmod -R 0777 ${DATA_DIR}

