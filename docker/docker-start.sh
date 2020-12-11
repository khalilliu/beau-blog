#!/bin/sh

envsubset '$API_SERVER' < /etc/nginx/nginx.conf.tpl > /etc/nginx/nginx.conf
env nginx
./bbadmin