#!/usr/bin/env bash

DEPLOY_USER=admin
CONF_DIR=conf

function print_usage {
  echo "Usage:"
  echo "  deploy.sh [-h|--help] [--dry] ssh_pem_file host env"
  echo
  echo "Options:"
  echo "  -h|--help            Display this help"
  echo "  --dry                Print out intentions, but do not take any actions"
  echo
  echo "Arguments:"
  echo "  ssh_pem_file         The ssh key file to use to connect to the host"
  echo "  host                 The public address of the host to deploy to"
  echo "  env                  The environment name for this deploy. Should match a file in conf/<environment>.sh"
}

# Process options
DRY_RUN=false

if [[ $# == 0 ]]; then print_usage; exit 1; fi
while [[ $# > 0 ]] ; do key="$1"
case ${key} in
    -h|--help) print_usage; exit 0;;
    --dry) DRY_RUN=true;;
    -*) echo "Illegal Option: ${key}"; print_usage; exit 1;;
    *) break;
esac
shift
done

PEM_FILE=${1}
HOST=${2}
ENVIRONMENT=${3}

if [ ! -e "${PEM_FILE}" ]; then
  echo "Must provide a valid ssh key to deploy. Provided file ${PEM_FILE} does not exist"
  exit 1
fi

if [[ "${HOST}" == "" ]]; then
  echo "Must provide a host to deploy to"
  exit 1
fi

if [[ "${ENVIRONMENT}" == "" || "${ENVIRONMENT}" == "common" || "$(ls ${CONF_DIR} | grep ${ENVIRONMENT}.sh)" == "" ]]; then
  echo "Must provide an environment that matches a file in conf. Eg: 'dev'. Cannot use 'common'."
  exit 1
fi


# Check that build has been done already
echo "Building service for target architecture"
if [ ${DRY_RUN} = false ]; then
  env GOOS=linux GOARCH=amd64 go build -o build/terragen
else
  echo "DRY RUN: env GOOS=linux GOARCH=amd64 go build -o build/terragen"
fi

if [ ! -e "build/static/main.js" ]; then
  echo "Must build website before deploying: yarn build"
  exit 1
fi

echo "Creating hash for the javascript bundle"
JS_BUNDLE=$(shasum build/static/main.js | cut -d ' ' -f 1)
echo "Bundle will have the file name: ${JS_BUNDLE}.js"

# Start deploy
echo "Deploying server artifacts to ${HOST}"
if [ ${DRY_RUN} = false ]; then
  scp -i ${PEM_FILE} -r build/ ${DEPLOY_USER}@${HOST}:/home/${DEPLOY_USER}
else
  echo "DRY RUN: scp -i ${PEM_FILE} -r build/ ${DEPLOY_USER}@${HOST}:/home/${DEPLOY_USER}"
fi

echo "Running local deploy script on ${HOST} with the ${ENVIRONMENT} config"
if [ ${DRY_RUN} = false ]; then
  printf "$(cat conf/common.sh conf/${ENVIRONMENT}.sh) \nexport TERRAGEN_JAVASCRIPT_BUNDLE='${JS_BUNDLE}'\n $(cat bin/deploy-serverside.sh)" | ssh -T -i ${PEM_FILE} ${DEPLOY_USER}@${HOST}
else
  echo "DRY RUN: printf \"$(cat conf/common.sh conf/${ENVIRONMENT}.sh) \nexport TERRAGEN_JAVASCRIPT_BUNDLE='${JS_BUNDLE}'\n $(cat bin/deploy-serverside.sh)\" | ssh -T -i ${PEM_FILE} ${DEPLOY_USER}@${HOST}"
fi
