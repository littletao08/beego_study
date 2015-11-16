#!/bin/bash
cat beego.pid | awk '{print $1}' | xargs pkill -9 -P