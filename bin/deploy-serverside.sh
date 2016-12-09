#!/usr/bin/env bash

SERVICE_DIRECTORY="/usr/local/terragen"

# Stop the running service if it exists
if [ -e "running_pid" ]; then
  sudo kill -9 "$(cat running_pid)"
  rm running_pid
fi

# Remove any old artifacts from the service directory
sudo rm -rf ${SERVICE_DIRECTORY}

# Modify the js bundle to match the pass in hash
if [[ "${TERRAGEN_JAVASCRIPT_BUNDLE}" == "" ]]; then
  echo "No bundle hash was provided. Need to have a value for the variable TERRAGEN_JAVASCRIPT_BUNDLE before deploy"
  exit 1
fi
mv build/static/main.js ./build/static/${TERRAGEN_JAVASCRIPT_BUNDLE}.js

# Deploy the new artifacts to the service directory
sudo mv build/ ${SERVICE_DIRECTORY}

# Make sure port 80 is redirected to port 8080
if [[ "$(sudo iptables -t nat --line-numbers -n -L | grep 'tcp dpt:80 redir ports 8080')" == "" ]]; then
  sudo iptables -t nat -A PREROUTING -p tcp --dport 80 -j REDIRECT --to-port 8080
fi

# Run the service
echo $(${SERVICE_DIRECTORY}/terragen &> lastlog & echo $!) > running_pid

sleep 2

echo "Logs after two seconds of startup time:"
echo $(cat lastlog)
echo
echo "Done Deploying and Restarting Service"
