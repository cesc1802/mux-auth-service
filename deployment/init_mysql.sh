#!/usr/bin/env bash

docker run -d --name auth-service \
  -e MYSQL_ROOT_PASSWORD="123456" \
  -e MYSQL_DATABASE="auth_service" \
  -e MYSQL_USER="auth_service" \
  -e MYSQL_PASSWORD="123456" \
  -e MYSQL_AUTHENTICATION_PLUGIN="mysql_native_password" \
  -p 3306:3306 \
  bitnami/mysql:8.0