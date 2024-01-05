# tinyWebServer

## About

A simple and tiny web server for demonstrations purpouses.

## Tasks

```sh
docker build -t tarsoqueiroz/tinywebinfo:1.0 .

docker run --publish 8080:8080 tarsoqueiroz/tinywebinfo:1.0

docker image push tarsoqueiroz/tinywebinfo:1.0
docker image push tarsoqueiroz/tinywebinfo:2.0
docker image push tarsoqueiroz/tinywebinfo:3.0 
```
