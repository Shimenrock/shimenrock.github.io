---
title: "Kibana7 安装与配置"
permalink: /elk/kibana-7-setup/
excerpt: "How to quickly install and setup Kibana 7."
last_modified_at: 2020-01-26T21:36:11-04:00
categories: elk
redirect_from:
  - /theme-setup/
toc: true
---

## 06.Kibana7 安装与配置

### 1.下载地址 

<u>https://www.elastic.co/downloads/kibana</u>
<u>https://www.elastic.co/guide/en/kibana/current/setup.html</u>
<u>https://www.elastic.co/guide/en/kibana/current/known-plugins.html</u>

### 2.Centos7.6下安装 
```
wget https://artifacts.elastic.co/downloads/kibana/kibana-7.3.2-linux-x86_64.tar.gz
shasum -a 512 kibana-7.3.2-linux-x86_64.tar.gz 
tar -xzf kibana-7.3.2-linux-x86_64.tar.gz
cd kibana-7.3.2-linux-x86_64/
chown elk：elk  kibana-7.3.2-linux-x86_64/

firewall-cmd --permanent --add-port=5601/tcp
firewall-cmd --reload
```

### 3.配置

 Config/kibana.yml
 elasticsearch.url?? 指向elasticsearch实例
```
server.port: 5601
server.host: 192.168.11.203
elasticsearch.hosts: ["http://192.168.11.203:9200"]
```
### 4.运行

```
# su - elk
```
启动? bin/kibana

插件? bin/kibana-plugin list

### 5.访问

<u>http://localhost:5601</u>

sample data测试数据(电商网站定点、航空公司飞行记录、网站日志记录)

Dashboards 面板

### 6.Dev Tools
 查看有哪些节点

* Get /_cat/nodes?v    

 查看帮助文档

* cmd + /

### 7.安装插件

* bin/kibana-plugin install plugin_location
* bin/kibana-plugin list
* bin/kibana remove

