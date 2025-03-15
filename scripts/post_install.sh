#!/bin/bash

echo "Post-install script started..."

if [ ! -d /var/log/ding ]; then
    mkdir -p /var/log/ding
    echo "Log directory created at /var/log/ding"
else
    echo "Log directory already exists."
fi

echo "Post-install script completed successfully."