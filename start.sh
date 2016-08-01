#!/bin/bash
go run main.go &
echo $! >beego.pid

time=`date "+%Y%m%d%H%M%S"`
version=`git log | awk 'NR==1{ print $2 }'`
tar -cjf /home/zhanglida/data/projects_pakcage_backup/beegostudy/beego-study-"$time"-$version.tar.bz2  .

exit 0 ;
