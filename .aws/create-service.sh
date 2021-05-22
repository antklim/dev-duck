#!/bin/bash

CONTAINER_PORT=8080
AUTH_CONTAINER_PORT=8090
PROJECT=dev-duck
SERVICE=DevDuck

aws cloudformation update-stack --stack-name dev-duck-service \
  --template-body file://service.yml \
  --parameters ParameterKey=ContainerImage,ParameterValue=$IMAGE \
  ParameterKey=Cluster,ParameterValue=$CLUSTER \
  ParameterKey=ContainerPort,ParameterValue=$CONTAINER_PORT \
  ParameterKey=AuthContainerImage,ParameterValue=$AUTH_IMAGE \
  ParameterKey=AuthContainerPort,ParameterValue=$AUTH_CONTAINER_PORT \
  ParameterKey=ServiceName,ParameterValue=$SERVICE \
  ParameterKey=ProjectName,ParameterValue=$PROJECT \
  ParameterKey=VPC,ParameterValue=$VPC \
  --tags Key=project,Value=$PROJECT \
  --region ap-southeast-2 \
  --capabilities CAPABILITY_NAMED_IAM \
  --output yaml
