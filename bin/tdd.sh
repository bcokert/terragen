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
    echo -e "\033[34mWatching files under: ${DIR}\033[0m"
    echo -e "\033[34mWatching all files that match pattern: ${PATTERN}\033[0m"
    make coverage
  fi
}

while true; do
  runIfChanged
  sleep 1
done
