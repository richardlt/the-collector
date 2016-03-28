#!/bin/sh

rm -rf ./node_modules
rm -rf ./bower_components
rm the-collector

npm install
bower install

gulp bundle
go build

rm -rf ./the-collector-package
mkdir the-collector-package

cp ./the-collector ./the-collector-package
mkdir -p ./the-collector-package/client/dist
cp -r ./client/dist/ ./the-collector-package/client/dist

rm ./the-collector.zip

cd the-collector-package
zip -ro ../the-collector.zip *
