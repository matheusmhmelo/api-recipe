#!/bin/bash

echo "Compiling the API"
docker run -it --rm -v "$(pwd)":/go -e GOPATH= golang:1.14 sh -c "CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/recipe/"

rm ./docker/recipe
mv ./recipe ./docker/

docker build -t matheusmhmelo/recipe:"$1" docker/

docker push matheusmhmelo/recipe:"$1"

if [[ ! $(docker service ls | grep dm_recipe) = "" ]]; then
  docker service update dm_recipe --image matheusmhmelo/recipe:$1
else
  docker stack deploy -c docker-compose.yaml dm
fi