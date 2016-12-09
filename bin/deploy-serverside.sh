#!/usr/bin/env bash

SERVICE_DIRECTORY="/usr/local"

# Stop the running service if it exists
if [ -e "running_pid" ]; then
  sudo kill -9 "$(cat running_pid)"
  rm running_pid
fi

# Remove any old artifacts from the service directory
sudo rm -rf ${SERVICE_DIRECTORY}/terragen

# Create unique javascript bundle
PRODUCTION_BUNDLE=$(shasum build/static/main.js | cut -d ' ' -f 1)
mv build/static/main.js ./build/static/${PRODUCTION_BUNDLE}.js

# Deploy the new artifacts to the service directory
sudo mv build/ ${SERVICE_DIRECTORY}/terragen

# Run the service
echo $(${SERVICE_DIRECTORY}/terragen/terragen &> lastlog & echo $!) > running_pid

echo "Done Deploying and Restarting Service"
