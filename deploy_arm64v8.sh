#!/bin/bash
path=`pwd`

i=$(docker images | grep "arm64v8/gin-backend" | awk '{print $1}')
if test -z $i; then
echo "not found the docker image, start build image..."
docker build -f ./DockerFile.arm64v8 -t arm64v8/gin-backend:v0.9.0 ../gin-backend
fi

i=$(docker images | grep "arm64v8/gin-backend" | awk '{print $1}')
if test -z $i; then
echo "build image error, exit shell!"
exit
fi

c=$(docker ps -a | grep "arm64v8-gin_backend-mysql" | awk '{print $1}')
if test -z $c; then
echo "not found the mysql server, start mysql server..."
docker run -d \
    -p 3306:3306 \
    -e MYSQL_ROOT_PASSWORD=123456 \
    -e MYSQL_DATABASE=gin_backend \
    --name arm64v8-gin_backend-mysql \
    --restart always \
    arm64v8/mysql:8.0
echo "waiting for database initialization..."
sleep 20s
docker logs --tail=10 arm64v8-gin_backend-mysql
fi
#
#i=$(docker ps -a | grep "arm64v8-gin-backend" | awk '{print $1}')
#if test ! -z $i; then
#echo "the server container already exists, delete..."
#docker rm -f arm64v8-gin-backend
#fi
#
#echo "start starnet demo server..."
#docker run -d \
#-p 8086:8086 \
#-w /gin-backend \
#-v $path/conf:/gin-backend/conf \
#-v $path/log:/gin-backend/log \
#-v $path/crypto-config:/gin-backend/crypto-config \
#--name arm64v8-gin-backend \
#--restart always \
#arm64v8/gin-backend:v0.9.0 \
#bash -c "cd src&&./gin-backend -config ../conf/system_config.yaml"
#sleep 2s
#docker logs arm64v8-gin-backend
#echo "the starnet demo server has been started!"
#
#
#i=$(docker ps -a | grep "arm64v8-starnet-web" | awk '{print $1}')
#if test ! -z $i; then
#echo "the web container already exists, delete..."
#docker rm -f arm64v8-starnet-web
#fi
#
#echo "start starnet demo web server..."
#chmod -R 777 $path/web/
#docker run -d \
#-p 80:80 \
#-v $path/web/conf.d:/etc/nginx/conf.d \
#-v $path/web/resources:/usr/share/nginx/resources \
#--name arm64v8-starnet-web \
#--restart always \
#arm64v8/nginx:1.23.3
#echo "the starnet demo web server has been started!"