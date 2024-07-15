#!/bin/bash
while getopts v: flag
do
    case "${flag}" in
        v) version=${OPTARG};;
    esac
done
docker buildx build -t hehelf/adm-req-vue:$version .
docker push hehelf/adm-req-vue:$version