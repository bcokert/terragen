#!/usr/bin/env bash

DEPLOY_USER=admin

function print_usage {
  echo "Usage:"
  echo "  deploy.sh [-h|--help] [--dry] ssh_pem_file host"
  echo
  echo "Options:"
  echo "  -h|--help            Display this help"
  echo "  --dry                Print out intentions, but do not take any actions"
  echo
  echo "Arguments:"
  echo "  ssh_pem_file         The ssh key file to use to connect to the host"
  echo "  host                 The public address of the host to deploy to"
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

if [ ! -e "${PEM_FILE}" ]; then
  echo "Must provide a valid ssh key to deploy. Provided file ${PEM_FILE} does not exist"
  exit 1
fi

if [[ "${HOST}" == "" ]]; then
  echo "Must provide a host to deploy to"
  exit 1
fi


# Check that build has been done already
echo "Building service for target architecture"
env GOOS=linux GOARCH=amd64 go build -o build/terragen

if [ ! -e "build/static/main.js" ]; then
  echo "Must build website before deploying: yarn build"
  exit 1
fi

# Start deploy
echo "Deploying server artifacts to ${HOST}"
if [ ${DRY_RUN} = false ]; then
  scp -i ${PEM_FILE} -r build/ ${DEPLOY_USER}@${HOST}:/home/${DEPLOY_USER}
else
  echo "DRY RUN: scp -i ${PEM_FILE} -r build/ ${DEPLOY_USER}@${HOST}:/home/${DEPLOY_USER}"
fi

echo "Deploying local deploy script to ${HOST}"
if [ ${DRY_RUN} = false ]; then
  scp -i ${PEM_FILE} bin/deploy-serverside.sh ${DEPLOY_USER}@${HOST}:/home/${DEPLOY_USER}
else
  echo "DRY RUN: scp -i ${PEM_FILE} bin/deploy-serverside.sh ${DEPLOY_USER}@${HOST}:/home/${DEPLOY_USER}"
fi

echo "Running local deploy script on ${HOST}"
if [ ${DRY_RUN} = false ]; then
  ssh -i ${PEM_FILE} ${DEPLOY_USER}@${HOST} "/home/${DEPLOY_USER}/deploy-serverside.sh"
else
  echo "DRY RUN: scp  -i ${PEM_FILE} bin/deploy-serverside.sh ${DEPLOY_USER}@${HOST}:/home/${DEPLOY_USER}"
fi
