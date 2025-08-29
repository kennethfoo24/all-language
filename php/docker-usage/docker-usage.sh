#!/bin/bash

# This script is used to track sandboxes usage

target_url='https://pluto.us1.prod.dog/v1/sandbox/usage'

# Prepare Json body
sandbox_path=$(echo $script_path | sed -n 's:.*sandbox/::p' )

json_post='{"sandbox_path":"'$sandbox_path'","user":"'$user'"}'

# Send request to Pluto
status_code=$(curl \
    -X POST $target_url \
    -H "Content-Type: application/json" -d "$json_post" \
    --connect-timeout 5 \
    --retry 2 --retry-max-time 10 \
    --write-out "%{http_code}"\
    -s)

if [ $status_code -ne 200 ] ; then
    # Restart the request, without retrying but with verbose option for debugging
    curl \
        -X POST $target_url \
        -H "Content-Type: application/json" -d "$json_post" \
        --connect-timeout 5 \
        -v
    echo -e "Can't push data to Pluto. Status code: $status_code\nPlease contact Supportability Engineering if the issue persists\n"
else
  echo -e "Usage posted successfully"
fi

exit 0