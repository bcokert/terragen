#!/usr/bin/env bash

PATTERN="${1}"

if [ "${PATTERN}" == "" ]; then
  PATTERN="*"
fi

DIR=`pwd`
LAST_HASH=""

function runIfChanged {
  NEW_HASH=`find . -type f -name "${PATTERN}" -ls | md5`
  SHOULD_RUN=false

  if [ "${NEW_HASH}" != "${LAST_HASH}" ]; then
    SHOULD_RUN=true
  fi
  LAST_HASH="${NEW_HASH}"

  if [ ${SHOULD_RUN} == true ]; then
    clear
    echo "Watching files under: ${DIR}"
    echo "Watching all files that match pattern: ${PATTERN}"
    make coverage
  fi
}

while true; do
  runIfChanged
  sleep 1
done
