#!/bin/bash

REPO=dev-duck
PROJECT=dev-duck

aws cloudformation update-stack --stack-name dev-duck-ecr \
  --template-body file://ecr.yml \
  --parameters ParameterKey=Name,ParameterValue=$REPO \
  ParameterKey=ProjectName,ParameterValue=$PROJECT \
  --tags Key=project,Value=$PROJECT \
  --region ap-southeast-2 \
  --capabilities CAPABILITY_NAMED_IAM \
  --output yaml
