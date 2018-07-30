#The Collector

[![Build Status](https://travis-ci.org/richardlt/the-collector.svg?branch=master)](https://travis-ci.org/richardlt/the-collector)
[![](https://badge.imagelayers.io/richardleterrier/the-collector:latest.svg)](https://imagelayers.io/?images=richardleterrier/the-collector:latest 'Get your own badge on imagelayers.io')

## How to run dev

sh'''
  make install
  npm start
  go run main.go start
'''

## How to run prod

sh'''
  docker network create the-collector
  docker run -d --name mongo --network the-collector mongo
  docker run -d -p 8080:8080 --name the-collector --network the-collector richardleterrier/the-collector --database-uri mongo:27017 start
'''

