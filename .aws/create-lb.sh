#!/bin/bash

# TODO: replace SUBNET_ONE, SUBNET_TWO with SUBNETS

. .env

aws cloudformation update-stack --stack-name dev-duck-lb \
  --template-body file://lb.yml \
  --parameters ParameterKey=ProjectName,ParameterValue=$PROJECT \
  ParameterKey=SubnetOne,ParameterValue=$SUBNET_ONE \
  ParameterKey=SubnetTwo,ParameterValue=$SUBNET_TWO \
  ParameterKey=SG,ParameterValue=$LOAD_BALANCER_SG \
  --tags Key=project,Value=$PROJECT \
  --region ap-southeast-2 \
  --capabilities CAPABILITY_NAMED_IAM \
  --output yaml
