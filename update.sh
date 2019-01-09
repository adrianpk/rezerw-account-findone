#!/bin/bash
build.sh
echo "Deploying update..."
aws lambda update-function-code --function-name FindOneAccount \
--zip-file fileb://./deployment.zip \
--region eu-west-1 \
--profile admin
echo "Update completed"