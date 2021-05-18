#!/bin/bash

PROJECT=dev-duck

aws cloudformation update-stack --stack-name dev-duck-lb \
  --template-body file://lb.yml \
  --parameters ParameterKey=ProjectName,ParameterValue=$PROJECT \
  ParameterKey=SubnetOne,ParameterValue=$SUBNET_ONE \
  ParameterKey=SubnetTwo,ParameterValue=$SUBNET_TWO \
  ParameterKey=SG,ParameterValue=$SG \
  --tags Key=project,Value=$PROJECT \
  --region ap-southeast-2 \
  --capabilities CAPABILITY_NAMED_IAM \
  --output yaml
