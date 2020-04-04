---
title: "Grafana 安装与配置"
permalink: /elk/grafana/
excerpt: "Grafana"
last_modified_at: 2020-04-04T21:36:11-04:00
categories: grafana
redirect_from:
  - /theme-setup/
toc: true
toc_sticky: true
---

## 1.相关资源

> * 官网：https://grafana.com
> * 文档：https://grafana.com/docs/
> * 数据源：https://grafana.com/docs/grafana/latest/features/datasources/
> * 下载：https://grafana.com/docs/grafana/latest/

## 2.安装
```
# cat /etc/redhat-release 
CentOS Linux release 7.6.1810 (Core) 
# yum install initscripts urw-fonts wget
# wget https://dl.grafana.com/oss/release/grafana-6.7.2-1.x86_64.rpm
# rpm -Uvh grafana-6.7.2-1.x86_64.rpm 

# sudo systemctl daemon-reload
# sudo systemctl start grafana-server
# sudo systemctl status grafana-server
# sudo systemctl enable grafana-server.service

访问：http://IP:3000，默认账号/密码：admin/admin
```

## 3.软件目录
```
# rpm -ql grafana-6.7.2 | more
```
- /usr/sbin/grafana-server          # 执行程序
- /etc/init.d/grafana-server        # init.d脚本
- /etc/sysconfig/grafana-server     # 环境变量
- /etc/grafana/grafana.ini          # 配置文件
- grafana-server.service            # systemd配置
- /var/log/grafana/grafana.log      # 日志
- /var/lib/grafana/grafana.db       # 默认数据库

## 4.环境变量配置

```
# cat /etc/passwd | grep grafana
grafana:x:997:995:grafana user:/usr/share/grafana:/sbin/nologin

# cat /etc/group | grep grafana
grafana:x:995:

# more /etc/sysconfig/grafana-server 

GRAFANA_USER=grafana

GRAFANA_GROUP=grafana

GRAFANA_HOME=/usr/share/grafana

LOG_DIR=/var/log/grafana

DATA_DIR=/var/lib/grafana

MAX_OPEN_FILES=10000

CONF_DIR=/etc/grafana

CONF_FILE=/etc/grafana/grafana.ini

RESTART_ON_UPGRADE=true

PLUGINS_DIR=/var/lib/grafana/plugins

PROVISIONING_CFG_DIR=/etc/grafana/provisioning

# Only used on systemd systems
PID_FILE_DIR=/var/run/grafana
```