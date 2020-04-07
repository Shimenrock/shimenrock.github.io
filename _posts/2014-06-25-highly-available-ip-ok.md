---
title:  "11gR2 RAC 新特性之Highly Available IP（HAIP）"
published: true
categories: oracle
permalink: oracle-haip.html
summary: "11gR2 RAC 新特性之Highly Available IP（HAIP）"
toc: true
---

从11.2.0.2开始，Oracle 的集群软件Grid Infrastructure(GI)中新增了Redundant Interconnect with Highly Available IP(HAIP)，以实现集群私网的高可用性和负载均衡。

在11.2.0.2之前，私网的冗余一般是通过在OS上做网卡绑（如bonding, EtherChannel等）实现的，有了HAIP之后，无需使用网卡绑定就可以实现私网网卡的冗余。

在安装GI的过程中，可以定义多个私网网卡来实现私网的冗余。

安装后，HAIP地址自动设置为169.254.*.*,这个地址不可以手动设置。HAIP 最少为１个，最多为４个(1块网卡，1个HAIP;2块网卡，2个HAIP; 3块及以上，4个HAIP), 均匀的分布在私网的网卡上。

## 1. 查看HAIP资源状态
```
$ crsctl stat res -t -init
NAME
TARGET STATE SERVER STATE_DETAILS Cluster Resources
-------------------------------------------------------------------------------------------------
ora.cluster_interconnect.haip
ONLINE ONLINE node2 1
```
## 2.查看HAIP地址和分布情况。
```
#ifconfig -a
eth1 Link encap:Ethernet HWaddr 00:0C:29:4B:B7:66
inet addr:192.168.254.32 Bcast:192.168.254.255 Mask:255.255.255.0
inet6 addr: fe80::20c:29ff:fe4b:b766/64 Scope:Link
UP BROADCAST RUNNING MULTICAST MTU:1500 Metric:1
......

eth1:1 Link encap:Ethernet HWaddr 00:0C:29:4B:B7:66
inet addr:169.254.31.199 Bcast:169.254.127.255 Mask:255.255.128.0 <=====HAIP address one.
UP BROADCAST RUNNING MULTICAST MTU:1500 Metric:1
Interrupt:193 Base address:0x1800

eth2 Link encap:Ethernet HWaddr 00:0C:29:4B:B7:70
inet addr:192.168.254.33 Bcast:192.168.254.255 Mask:255.255.255.0
inet6 addr: fe80::20c:29ff:fe4b:b770/64 Scope:Link
UP BROADCAST RUNNING MULTICAST MTU:1500 Metric:1
......

eth2:1 Link encap:Ethernet HWaddr 00:0C:29:4B:B7:70
inet addr:169.254.185.222 Bcast:169.254.255.255 Mask:255.255.128.0 <=====HAIP address two.
UP BROADCAST RUNNING MULTICAST MTU:1500 Metric:1
Interrupt:169 Base address:0x1880
```
haip均匀的分布在两个私网网卡上。

## 3. 断掉网卡eth1之后。
```
eth2 Link encap:Ethernet HWaddr 00:0C:29:4B:B7:70
inet addr:192.168.254.33 Bcast:192.168.254.255 Mask:255.255.255.0
inet6 addr: fe80::20c:29ff:fe4b:b770/64 Scope:Link
UP BROADCAST RUNNING MULTICAST MTU:1500 Metric:1
RX packets:3206 errors:0 dropped:0 overruns:0 frame.:0
TX packets:3916 errors:0 dropped:0 overruns:0 carrier:0
collisions:0 txqueuelen:1000
RX bytes:1474658 (1.4 MiB) TX bytes:2838774 (2.7 MiB)
Interrupt:169 Base address:0x1880

eth2:1 Link encap:Ethernet HWaddr 00:0C:29:4B:B7:70
inet addr:169.254.185.222 Bcast:169.254.255.255 Mask:255.255.128.0 <=====HAIP address two.
UP BROADCAST RUNNING MULTICAST MTU:1500 Metric:1
Interrupt:169 Base address:0x1880

eth2:2 Link encap:Ethernet HWaddr 00:0C:29:4B:B7:70
inet addr:169.254.31.199 Bcast:169.254.127.255 Mask:255.255.128.0 <=====HAIP address one.
UP BROADCAST RUNNING MULTICAST MTU:1500 Metric:1
Interrupt:169 Base address:0x1880
```
HAIP one 漂移到了网卡eth2上。

## 4. 网卡eth1恢复之后。
```
eth1 Link encap:Ethernet HWaddr 00:0C:29:4B:B7:66
inet addr:192.168.254.32 Bcast:192.168.254.255 Mask:255.255.255.0
inet6 addr: fe80::20c:29ff:fe4b:b766/64 Scope:Link
UP BROADCAST RUNNING MULTICAST MTU:1500 Metric:1
......

eth1:1 Link encap:Ethernet HWaddr 00:0C:29:4B:B7:66
inet addr:169.254.31.199 Bcast:169.254.127.255 Mask:255.255.128.0 <=====HAIP address one.
UP BROADCAST RUNNING MULTICAST MTU:1500 Metric:1
Interrupt:193 Base address:0x1800

eth2 Link encap:Ethernet HWaddr 00:0C:29:4B:B7:70
inet addr:192.168.254.33 Bcast:192.168.254.255 Mask:255.255.255.0
inet6 addr: fe80::20c:29ff:fe4b:b770/64 Scope:Link
UP BROADCAST RUNNING MULTICAST MTU:1500 Metric:1
......

eth2:1 Link encap:Ethernet HWaddr 00:0C:29:4B:B7:70
inet addr:169.254.185.222 Bcast:169.254.255.255 Mask:255.255.128.0 <=====HAIP address two.
UP BROADCAST RUNNING MULTICAST MTU:1500 Metric:1
Interrupt:169 Base address:0x1880
```
HAIP one 回到了网卡eth1上。

**注意：HAIP地址失败不会对ocssd产生影响，也就是说HAIP失败，不会导致节点重启。**

## HAIP 对数据库和ASM的影响

数据库和ASM实例使用这个HAIP作为cluster interconnect,以下是alert.log的片段。
```
Cluster communication is configured to use the following interface(s) for this instance
169.254.31.199
169.254.185.222
cluster interconnect IPC version:Oracle UDP/IP (generic)
IPC Vendor 1 proto 2
```
Oracle数据库和ASM实例可以通过HAIP来实现私网通讯的高可用性和负载均衡。私网的流量会在这些私网网卡上实现负载均衡，
如果某个网卡出现了故障，它上面的HAIP会自动切换到别的可用的私网网卡上，从而不影响私网的通讯。

注意：HAIP 是不允许被手动停止或禁用的，除非是由于某些版本或者平台不支持。

关于HAIP的更多介绍，请参考My Oracle Support Note 文档1210883.1.