#!/bin/bash

. .env

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
  ParameterKey=Subnets,ParameterValue=$SUBNETS \
  ParameterKey=LoadBalancerSG,ParameterValue=$LOAD_BALANCER_SG \
  ParameterKey=LoadBalancerListener,ParameterValue=$LOAD_BALANCER_LISTENER \
  --tags Key=project,Value=$PROJECT \
  --region ap-southeast-2 \
  --capabilities CAPABILITY_NAMED_IAM \
  --output yaml
