---
title: "Auto restart Weblogic Managed Server"
published: true
categories: weblogic
permalink: how-to-auto-restart-weblogic.html
summary: "如何通过定时脚本自动重启weblogic节点"
tags: [atom]
toc: true
---

- Admin Server http://192.168.0.10:9001/console
- Managed Server-1 http://192.168.0.10:8011
- Managed Server-2 http://192.168.0.10:8012

## 1.Create wlst Script
```
# cat start_server-1.py
connect('weblogic','weblogic','t3://192.168.0.10:9001')
start ('server-1')
disconnect();
```
```
# cat stop_server-1.py
connect('weblogic','weblogic','t3://192.168.0.10:9001')
shutdown('server-1')
disconnect();
```
## 2.Create linux shell script
```
# cat restart_server-1.sh
#!/bin/bash
export wlst=/weblogic/wlserver_10.3/common/bin/wlst.sh
export server1_start_wlst=/weblogic/user_projects/domains/ggfw_domain/wlst/start_server-1.py
export server1_stop_wlst=/weblogic/user_projects/domains/ggfw_domain/wlst/stop_server-1.py
export server1_restar_log=/weblogic/user_projects/domains/server_domain/wlst/server-1_restar.log
#
echo "***************************************************" >> $server-1_restar_log
date +%Y-%M-%d" "%H:%M:%S >> $server-1_restar_log
$wlst $server1_stop_wlst >> $server-1_restar_log
$wlst $server1_start_wlst >> $server-1_restar_log
```
## 3.Create crontab job
```
# crontab -e
* 1 * * * /weblogic/user_projects/domains/server_domain/wlst/restart_server-1.sh
# 每天凌晨1点对被管服务器重启
```
## 4.About log
```
***************************************************
.............................

Your environment has been set.

.............................

Initializing WebLogic Scripting Tool (WLST) ...

Welcome to WebLogic Server Administration Scripting Shell

Type help() for help on available commands

Connecting to t3://192.168.0.10:9001 with userid weblogic ...
Successfully connected to Admin Server 'SERVER_AdminServer' that belongs to domain 'server_domain'.

Warning: An insecure protocol was used to connect to the
server. To ensure on-the-wire security, the SSL port or
Admin port should be used instead.

..............................


Starting server server-1 ...........................
Server with name server-1 started successfully
Disconnected from weblogic server: Server_AdminServer
```
## WebLogic 脚本工具
http://edocs.weblogicfans.net/wls/docs92/config_scripting/index.html
## WLST 命令和变量参考
http://edocs.weblogicfans.net/wls/docs92/config_scripting/reference.html