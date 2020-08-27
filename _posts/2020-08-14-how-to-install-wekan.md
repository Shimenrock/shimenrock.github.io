---
title: "how to install wekan on Centos 7"
published: true
related: true
header:
  teaser: /assets/images/2020-08-14-wekan.jpg
toc: true
toc_sticky: true
categories: 
  - wekan
---  

官网：https://wekan.github.io/

GITHUB:https://github.com/wekan/wekan

## 1.root 登陆
 
基础配置 
```
# hostnamectl set-hostname wekan225
# yum install net-tools -y
# vi /etc/sysconfig/selinux 
SELINUX=disabled
# vi /etc/sysconfig/network-scripts/ifcfg-ens192 
DNS1="8.8.8.8"
DNS2="4.4.4.4"
DNS3="114.114.114.114"
# service network restart
# systemctl stop firewalld.service 
# systemctl disable firewalld.service 
```

## 2.Install snap
```
# yum makecache fast 

# yum install yum-plugin-copr epel-release -y

# yum copr enable ngompa/snapcore-el7

# yum install snapd -y

# systemctl status snapd.service
● snapd.service - Snap Daemon
   Loaded: loaded (/usr/lib/systemd/system/snapd.service; disabled; vendor preset: disabled)
   Active: inactive (dead)
# systemctl enable --now snapd.socket
Created symlink from /etc/systemd/system/sockets.target.wants/snapd.socket to /usr/lib/systemd/system/snapd.socket.

```
**创建系统快照**

## 3.Install Wekan. Set URL like (subdomain.)example.com(/suburl)

注意snap安装需要挂代理，或配置软路由，否则报错DNS解析失败
```
# snap install wekan
2020-08-14T19:08:24+08:00 INFO Waiting for automatic snapd restart...
wekan 4.23 from Lauri Ojansivu (xet7) installed

# snap list
Name   Version      Rev   Tracking       Publisher   Notes
core   16-2.45.3.1  9804  latest/stable  canonical✓  core
wekan  4.23         956   latest/stable  xet7        -

# snap interfaces|grep wekan
:network                   wekan
:network-bind              wekan
wekan:mongodb-slot         -

'snap interfaces' is deprecated; use 'snap connections'.
-                          wekan:mongodb-plug

# systemctl status snap.wekan.mongodb
● snap.wekan.mongodb.service - Service for snap application wekan.mongodb
   Loaded: loaded (/etc/systemd/system/snap.wekan.mongodb.service; enabled; vendor preset: disabled)
   Active: active (running) since 五 2020-08-14 19:08:45 CST; 1min 6s ago
 Main PID: 16340 (mongodb-control)
   CGroup: /system.slice/snap.wekan.mongodb.service
           ├─16340 /bin/bash /snap/wekan/956/bin/mongodb-control
           └─17612 mongod --dbpath /var/snap/wekan/common --syslog --journal --bind_ip 127.0.0.1 --port 27019 --quiet

8月 14 19:09:30 wekan mongod.27019[17612]: [conn3] received client metadata from 127.0.0.1:55984 conn3: { driver: { name: "nodejs", v....0.11" }
8月 14 19:09:30 wekan mongod.27019[17612]: [conn4] received client metadata from 127.0.0.1:55986 conn4: { driver: { name: "nodejs", v....0.11" }
8月 14 19:09:35 wekan mongod.27019[17612]: [conn3] received client metadata from 127.0.0.1:55984 conn3: { driver: { name: "nodejs", v....0.11" }
8月 14 19:09:35 wekan mongod.27019[17612]: [conn4] received client metadata from 127.0.0.1:55986 conn4: { driver: { name: "nodejs", v....0.11" }
8月 14 19:09:40 wekan mongod.27019[17612]: [conn3] received client metadata from 127.0.0.1:55984 conn3: { driver: { name: "nodejs", v....0.11" }
8月 14 19:09:40 wekan mongod.27019[17612]: [conn4] received client metadata from 127.0.0.1:55986 conn4: { driver: { name: "nodejs", v....0.11" }
8月 14 19:09:45 wekan mongod.27019[17612]: [conn3] received client metadata from 127.0.0.1:55984 conn3: { driver: { name: "nodejs", v....0.11" }
8月 14 19:09:45 wekan mongod.27019[17612]: [conn4] received client metadata from 127.0.0.1:55986 conn4: { driver: { name: "nodejs", v....0.11" }
8月 14 19:09:50 wekan mongod.27019[17612]: [conn3] received client metadata from 127.0.0.1:55984 conn3: { driver: { name: "nodejs", v....0.11" }
8月 14 19:09:50 wekan mongod.27019[17612]: [conn4] received client metadata from 127.0.0.1:55986 conn4: { driver: { name: "nodejs", v....0.11" }
Hint: Some lines were ellipsized, use -l to show in full.


# systemctl status snap.wekan.wekan
● snap.wekan.wekan.service - Service for snap application wekan.wekan
   Loaded: loaded (/etc/systemd/system/snap.wekan.wekan.service; enabled; vendor preset: disabled)
   Active: active (running) since 五 2020-08-14 19:08:45 CST; 1min 29s ago
 Main PID: 16357 (wekan-control)
   CGroup: /system.slice/snap.wekan.wekan.service
           ├─16357 /bin/bash /snap/wekan/956/bin/wekan-control
           └─17608 /snap/wekan/956/bin/node main.js

8月 14 19:08:50 wekan wekan.wekan[16357]: > Starting add-description-title-allowed migration.
8月 14 19:08:50 wekan wekan.wekan[16357]: > Finishing add-description-title-allowed migration.
8月 14 19:08:50 wekan wekan.wekan[16357]: > Starting add-description-text-allowed migration.
8月 14 19:08:50 wekan wekan.wekan[16357]: > Finishing add-description-text-allowed migration.
8月 14 19:08:50 wekan wekan.wekan[16357]: > Starting add-sort-field-to-boards migration.
8月 14 19:08:50 wekan wekan.wekan[16357]: > Finishing add-sort-field-to-boards migration.
8月 14 19:08:50 wekan wekan.wekan[16357]: > Starting add-default-profile-view migration.
8月 14 19:08:50 wekan wekan.wekan[16357]: > Finishing add-default-profile-view migration.
8月 14 19:08:50 wekan wekan.wekan[16357]: Meteor APM: completed instrumenting the app
8月 14 19:08:51 wekan wekan.wekan[16357]: {"line":"87","file":"percolate_synced-cron.js","message":"SyncedCron: Scheduled \"notificat...:"info"}
Hint: Some lines were ellipsized, use -l to show in full.
```
**创建系统快照**
```
# snap services wekan
Service        Startup   Current   Notes
wekan.caddy    disabled  inactive  -
wekan.mongodb  enabled   active    -
wekan.wekan    enabled   active    -
# snap get wekan
error: snap "wekan" has no configuration

# snap set wekan root-url='http://192.168.11.225:5000'
//局域网访问

# snap set wekan root-url='http://localhost:5000' 
//只能本机访问
# snap set wekan root-url='http://boards.example.com'
//互联网访问
# snap set wekan port='5000'   
//设置对外访问端口

//以下参考配置
# snap set wekan mongodb-bind-ip="0.0.0.0"   //允许ip访问,这里允许任何人访问
# snap set wekan mail-url='smtp://*********:q***ic@smtp.qq.com:465
# snap set wekan mail-from='来自<*********@qq.com>'

# snap stop wekan //关闭wekan服务
# snap start wekan //开启wekan服务,第一次安装完成后就已经启动服务了
# snap restart wekan //重启wekan服务
```

## 4.Set port where Wekan runs, for example 80 if http, or local port 3001, if running behing proxy like Caddy
```
# systemctl restart snap.wekan.wekan
# snap get wekan
Key       Value
port      5000
root-url  http://192.168.11.225:5000
# snap services wekan
Service        Startup   Current   Notes
wekan.caddy    disabled  inactive  -
wekan.mongodb  enabled   active    -
wekan.wekan    enabled   active    -
```
## 5.Install all Snap updates automatically between 02:00AM and 04:00AM
```
# snap set core refresh.schedule=02:00-04:00
```

## 6.访问系统

http://192.168.11.225:5000/

先注册一个新用户

用新用户登陆