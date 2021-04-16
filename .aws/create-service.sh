#!/bin/bash

CONTAINER_PORT=8080
PROJECT=dev-duck
SERVICE=DevDuck

aws cloudformation create-stack --stack-name dev-duck-service \
  --template-body file://service.yml \
  --parameters ParameterKey=ContainerImage,ParameterValue=$IMAGE \
  ParameterKey=ContainerPort,ParameterValue=$CONTAINER_PORT \
  ParameterKey=ServiceName,ParameterValue=$SERVICE \
  ParameterKey=ProjectName,ParameterValue=$PROJECT \
  --tags Key=project,Value=$PROJECT \
  --region ap-southeast-2 \
  --capabilities CAPABILITY_NAMED_IAM \
  --output yaml \
  --profile $PROFILE
