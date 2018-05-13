#!/bin/bash


pid=`lsof -i :8080|grep LISTEN |awk '{print $2}'`
kill $pid
go run main.go&
#curl 'http://127.0.0.1:8080/test/testdb'
