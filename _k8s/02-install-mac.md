---
title: "Install macOS with VMware ESXi 6.7"
permalink: /k8s/install-macos-esxi/
excerpt: "Install macOS with VMware ESXi 6.7"
last_modified_at: 2020-01-28T21:36:11-04:00
redirect_from:
  - /theme-setup/
toc: true
---

[GitHub](https://github.com/DrDonk)

[VMware Workstation macOS](https://github.com/DrDonk/unlocker)

[VMware ESXi macOS](https://github.com/DrDonk/esxi-unlocker)

```
[root@localhost:/tmp] tar xzvf esxi-unlocker-300.tgz
unlocker.tgz
esxi-install.sh
esxi-uninstall.sh
esxi-smctest.sh
readme.txt
[root@localhost:/tmp] ls
esxi-install.sh        esxi-uninstall.sh      probe.session          unlocker.tgz
esxi-smctest.sh        esxi-unlocker-300.tgz  readme.txt             vmware-root
[root@localhost:/tmp] ./esxi-install.sh
VMware Unlocker 3.0.0
===============================
Copyright: Dave Parsons 2011-18
Installing unlocker.tgz
Acquiring lock /tmp/bootbank.lck
Copying unlocker.tgz to /bootbank/unlocker.tgz
Editing /bootbank/boot.cfg to add module unlocker.tgz
Success - please now restart the server!
```

完成解锁后
上传OS 10.13 HSierra.cdr，并修改后缀为OS 10.13 HSierra.iso
https://www.youtube.com/watch?v=7n-_zaNBX5E

  - 1. 通过IP:80访问ESXI的web管理页面。
  - 2. 创建虚拟机》创建新虚拟机》选择客户机操作系统 “Mac OS”，选择客户机操作系统系列版本 “Apple Mac OS X 10.13(64位)”》选择存储》
  - 3. 给虚拟机适当的资源，这里给了2颗CPU,8G内存，SATA 128G高速硬盘（精简），USB3.0,千兆网卡，显存16M(一会修改为直通显卡)，选择安装镜像
  - 4. 安装，选择简体中文》 继续 》 继续 》 同意 》 实用工具 》 磁盘工具 》 选择VMware Virtual SATA Hard Drive Media 》 抹掉 》 选择初始化完成的硬盘，按左上角红×返回 》 继续 》。。。。。。。。
  - 5. 把darwin.iso加载到CD/DVD里面,双击安装VMware Tools
