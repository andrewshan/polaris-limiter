#!/bin/bash

if [ $# != 1 ]; then
    echo "e.g.: bash $0 v1.0"
    exit 1
fi

docker_repository="polarismesh/polaris-limiter"
docker_tag=$1

echo "docker repository : ${docker_repository}, tag : ${docker_tag}"

arch_list=( "amd64" "arm64" )
platforms=""

for arch in ${arch_list[@]}; do
    export GOARCH=${arch}
    bash build.sh ${docker_tag}
    if [ $? != 0 ]; then
      echo "build polaris-limiter failed"
      exit 1
    fi

    mv polaris-limiter polaris-limiter-${arch}
    platforms+="linux/${arch},"
done

platforms=${platforms::-1}
extra_tags=""

pre_release=`echo ${docker_tag}|egrep "(alpha|beta|rc|[T|t]est)"|wc -l`
if [ ${pre_release} == 0 ]; then
  extra_tags="-t ${docker_repository}:latest"
fi

docker buildx build --network=host -t ${docker_repository}:${docker_tag} ${extra_tags} --platform ${platforms} --push ./
