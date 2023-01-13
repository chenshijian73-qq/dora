#!/usr/bin/env bash

set -e

BUILD_VERSION="2.0.0"
BUILD_DATE=$(date "+%F %T")
COMMIT_SHA1=$(git rev-parse HEAD)

TARGET_DIR="dist"
TARGET_NAME="dora"
PLATFORMS="darwin/amd64 darwin/arm64 linux/386 linux/amd64 linux/arm linux/arm64"

rm -rf ${TARGET_DIR}
mkdir ${TARGET_DIR}

go_install_path=$(echo ${GOPATH}|awk -F ":" '{print$1}')
install_path="/usr/local/bin"

if [ "$1" == "install" ]; then
  echo "install to ${install_path}"
  go build -o ${go_install_path}/dora -ldflags \
    "-X 'github.com/chenshijian73-qq/doraemon/cmd.Version=${BUILD_VERSION}' \
    -X 'github.com/chenshijian73-qq/doraemon/cmd.BuildDate=${BUILD_DATE}' \
    -X 'github.com/chenshijian73-qq/doraemon/cmd.CommitID=${COMMIT_SHA1}' "
  echo "install => ${install_path}/dora"
  ln -sf ${go_install_path}/dora ${install_path}/dora
elif [ "$1" == "uninstall" ]; then
    echo "remove => ${GOPATH}/bin/csj"
    rm -f ${GOPATH}/bin/csj
    for cmd in ${COMMANDS}; do
        echo "remove => ${GOPATH}/bin/${cmd}"
        rm -f ${GOPATH}/bin/${cmd}
    done
    rm -rf ${install_path}/dora
else
  for pl in ${PLATFORMS}; do
      export GOOS=$(echo "${pl}" | cut -d'/' -f1)
      export GOARCH=$(echo "${pl}" | cut -d'/' -f2)
      export CGO_ENABLED=0

      export TARGET=${TARGET_DIR}/${TARGET_NAME}_${GOOS}_${GOARCH}
      if [ "${GOOS}" == "windows" ]; then
          export TARGET=${TARGET_DIR}/${cmd}_${GOOS}_${GOARCH}.exe
      fi

      echo "build => ${TARGET}"
      go build -trimpath -o "${TARGET}" \
              -ldflags    "-X 'github.com/chenshijian73-qq/doraemon/cmd.Version=${BUILD_VERSION}' \
                          -X 'github.com/chenshijian73-qq/doraemon/cmd.BuildDate=${BUILD_DATE}' \
                          -X 'github.com/chenshijian73-qq/doraemon/cmd.CommitID=${COMMIT_SHA1}' \
                          -w -s"
  done
fi