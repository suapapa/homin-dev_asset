#!/bin/bash

set -e

CR=icn.vultrcr.com/homincr1
IMAGE_VERSION=latest # $1
IMAGE_TAG=$CR/asset:$IMAGE_VERSION 
docker buildx build --platform linux/amd64 -t $IMAGE_TAG .
docker push $IMAGE_TAG

if [ "$1" != "latest" ]; then
    IMAGE_TAG_LATEST=$CR/asset:latest 
    docker rmi $IMAGE_TAG_LATEST || true
    docker tag $IMAGE_TAG $IMAGE_TAG_LATEST
    docker push $IMAGE_TAG_LATEST
fi

# git tag -a $1 -m "add tag for $1"
# git push --tags

kubectl rollout restart deployment asset
