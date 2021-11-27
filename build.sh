#!/bin/sh

docker buildx build \
  --push -t ioannrove/elastic_beanstalk_example:nginx-latest \
  --platform=linux/amd64,linux/arm64 /Users/ivan/projects/elastic_beanstalk_example/nginx

docker buildx build \
  --push -t ioannrove/elastic_beanstalk_example:frontend-latest \
  --platform=linux/amd64,linux/arm64 /Users/ivan/projects/elastic_beanstalk_example/frontend

docker buildx build \
  --push -t ioannrove/elastic_beanstalk_example:backend-latest \
  --platform=linux/amd64,linux/arm64 /Users/ivan/projects/elastic_beanstalk_example/backend