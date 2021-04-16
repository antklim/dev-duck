#!/bin/bash

IMAGE=dev-duck
PROJECT=dev-duck

aws cloudformation create-stack --stack-name dev-duck-ecr \
  --template-body file://ecr.yml \
  --parameters ParameterKey=Name,ParameterValue=$IMAGE \
  ParameterKey=ProjectName,ParameterValue=$PROJECT \
  --tags Key=project,Value=$PROJECT \
  --region ap-southeast-2 \
  --capabilities CAPABILITY_NAMED_IAM \
  --output yaml \
  --profile $PROFILE