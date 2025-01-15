#!/bin/sh

mongod --auth --bind_ip_all
mongosh localhost:27017 -u admin -p admin < /docker-entrypoint-initdb.d/seed.js
