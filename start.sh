#!/bin/bash
go run main.go> log.log &
echo $! >beego.pid

tail -f log.log
