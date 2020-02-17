---
title:  "AIX系统ssh服务不能启动,解决方法"
published: true
categories: AIX
permalink: aix-ssh-norun.html
summary: "AIX系统ssh服务不能启动,解决方法"
tags: [AIX]
---

### aix系统ssh服务不能启动

 install_assist安装助手，启动后根据/etc/inittab顺序加载服务，到install_assist会卡住的。

 说明：开机助手是安装机器后必须点掉的过程，只能说明系统安装的人员遗漏了这一步。

 启动过程删除了下边这行
```
install_assist:2:wait:/usr/sbin/install_assist </dev/console >/dev/console 2>&1
```
