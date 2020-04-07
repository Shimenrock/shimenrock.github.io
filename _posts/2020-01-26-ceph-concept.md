---
title: "CEPH 架构"
published: false
categories: ceph
permalink: ceph-concept.html
summary: "CEPH 架构"
---

**Software Defined Storage , SDS 软件定义存储**

## 架构概述

- Ceph monitor(MON)
  - 通过保存一份集群状态映射来维护整个集群的健康状态
  - 包括OSD map，MON map, PG map ,CRUSH map
  - 所有群集节点都向MON节点汇报状态信息，并分享它们状态中的任何变化
  - 不存储数据

- Ceph 对象存储设备(OSD)
  - 只要应用程序向Ceph集群发出写操作，数据就会被以对象形式存储在OSD中
  - 通常一个OSD守护进程会捆绑到集群中的一块物理磁盘上

- Ceph 元数据服务器(MDS)
  - 只为CephFS文件系统跟踪文件的层次结构和存储元数据。

- RADOS
  - Reliable Autonomic Distributed Object Store
  - 可靠 自主 分布式 对象存储，基础组件
  - RADOS层确保数据始终保持一致

- librados
  - 为编程语言提供访问RADOS接口方式
  - 为RBD,RGW,CephFS提供原生接口

- RADOS 块设备(RBD)
  - 提供持久块存储

- RADOS网关接口(RGW)
  - 提供对象存储服务
  - 使用librgw  librados
  - 提供了与Amazon S3和OpenStack Swift兼容的RESTful API

- CephFS
  - 提供一个使用Ceph存储集群存储用户数据的与POSIX兼容的文件系统
