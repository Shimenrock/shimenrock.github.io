---
title: "Cerebro 安装与配置"
permalink: /elk/cerebro-setup/
excerpt: "How to install and setup Cerebro."
last_modified_at: 2020-01-26T21:36:11-04:00
categories: elk
redirect_from:
  - /theme-setup/
toc: true
---

# 10.Cerebro 安装与配置

## 1.下载地址

[https://github.com/lmenezes/cerebro](https://github.com/lmenezes/cerebro)

[https://github.com/lmenezes/cerebro/releases](https://github.com/lmenezes/cerebro/releases)


## 2.Centos7.6 下安装
```
wget -c https://github.com/lmenezes/cerebro/releases/download/v0.8.4/cerebro-0.8.4.tgz
tar xf cerebro-0.8.4.tgz
chown -R elk:elk cerebro-0.8.4/
su - elk
cd /cerebro-0.8.4/conf
vi application.conf
修改
     username = admin
     password = Sjzrsj_2014

hosts = [
  {
     host="http://192.168.11.203:9200"
     name = "my-app"
  }
```


firewall-cmd --permanent --add-port=9000/tcp
firewall-cmd --reload

执行即可