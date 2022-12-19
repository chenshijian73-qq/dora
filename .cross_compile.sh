#!/usr/bin/env bash

set -e

BUILD_VERSION="1.0.0-beta"
BUILD_DATE=$(date "+%F %T")
COMMIT_SHA1=$(git rev-parse HEAD)

TARGET_DIR="dist"
TARGET_NAME="dora"
PLATFORMS="darwin/amd64 darwin/arm64 linux/386 linux/amd64 linux/arm linux/arm64"
COMMANDS="json2csv version ctx sprint"

rm -rf ${TARGET_DIR}
mkdir ${TARGET_DIR}

if [ "$1" == "install" ]; then
  echo "install to ${GOPSTH}/bin/csj"
  go build -o ${GOPSTH}/bin/csj -ldflags \
    "-X 'github.com/chenshijian73-qq/doraemon/cmd.Version=$(BUILD_VERSION)' \
    -X 'github.com/chenshijian73-qq/doraemon/cmd.BUILD_DATE=$(BUILD_DATE)' \
    -X 'github.com/chenshijian73-qq/doraemon/cmd.CommitID=$(COMMIT_SHA1)' "
  for cmd in ${COMMANDS}; do
      echo "install => ${GOPATH}/bin/${cmd}"
      ln -sf ${GOPATH}/bin/dora ${GOPATH}/bin/${cmd}
  done
elif [ "$1" == "uninstall" ]; then
    echo "remove => ${GOPATH}/bin/csj"
    rm -f ${GOPATH}/bin/csj
    for cmd in ${COMMANDS}; do
        echo "remove => ${GOPATH}/bin/${cmd}"
        rm -f ${GOPATH}/bin/${cmd}
    done
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