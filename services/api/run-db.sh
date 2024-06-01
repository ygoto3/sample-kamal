#!/bin/sh

docker container run -d --mount type=bind,src="$(pwd)"/kamal/config/init.sql,dst=/docker-entrypoint-initdb.d/init.sql --name dev-db -p 3306:3306 -e MYSQL_ROOT_PASSWORD=$MYSQL_ROOT_PASSWORD mysql:8.4