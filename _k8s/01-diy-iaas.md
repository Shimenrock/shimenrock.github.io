---
title: "DIY & Install VMware ESXi 6.7.0 update02"
permalink: /k8s/install-esxi-6/
excerpt: "DIY & Install VMware ESXi 6.7.0 update02"
last_modified_at: 2020-01-28T21:36:11-04:00
redirect_from:
  - /theme-setup/
toc: true
---

# Part1 DIY & Install VMware ESXi 6.7.0 update02

## 1.Hardware list

 The hardware environment is purchased as follows:

| hardware     | parameter         | price   | remarks                        |
| ------------ | ----------------- | ------- | ------------------------------ |
| CPU          | i3 8100           | CN¥ 999 | Intel                          |
| RAM          | DDR4 2666 16G*2   | CN¥ 800 | Apacer                         |
| hard disk 1  | M.2(NVMe) 256G    | CN¥ 229 | Colorful [ESXi 6.7 for Setup]  |
| hard disk 2  | SATA 2T           | CN¥ 500 | Western Digital [Storing data] |
| power        | 550W              | CN¥ 279 | aigo G-T550                    |
| fan          | Water cooling     | CN¥ 229 | aigo 冰塔v120 彩虹版           |
| Motherboard  | B365M AORUS ELITE | CN¥ 799 | GIGABYTE                       |
| Chassis      | 黑曼巴标准版      | CN¥ 249 | aigo                           |
| Network card | PCIE X1 82575EB   | CN¥ 95  | 蝶舞[Deploy soft routes]       |

总共预算4000元，解决软路由、NAS、MacOS学习、K8S集群学习等问题。
注：以上硬件顺利安装VMware ESXi 6.7，并识别，无需特别安装驱动。

## 2.Computer Assembly
![DIY01](/assets/images/DIY01.jpg)

![DIY02](/assets/images/DIY02.jpg)

![DIY03](/assets/images/DIY03.jpg)

![DIY04](/assets/images/DIY04.jpg)

## 3.Install VMware ESXi 6.7

**Build 13006603 2019年4月发布版本**

vmware-vmvisor-installer-6.7.0.update02 注意update02版本性能最好。
  - 使用rufus-2.5制作VMware ESXi 6.7安装U盘
  - 利用U盘引导安装,安装过程大概10分钟。
  - 如果之前有安装，安装过程会提示“VMFS Found”，选择“Install ESXI, preserve VMFS datastore”
  - 完成安装后，按F2将管理地址修改为静态IP地址。
https://www.youtube.com/redirect?v=Vaj5_NhEWVI&redir_token=RIvi6E3D-QWjYImNTrtSNm6rdq98MTU3NDYyNDQ1MUAxNTc0NTM4MDUx&event=video_description&q=http%3A%2F%2Fa.sssru.space%2Fauth%2Fregister%3Fcode%3DrvxADrosr4SuhV3W9SZggZbZto25guaV
4群41350354


## 4.ESXI配置自动关机
```
cp /sbin/powerOffVms /vmfs/volumes/datastore1/poweroffvms
cd /vmfs/volumes/datastore1/
vi poweroffvms
```
将“except Exception, e:” 修改为”except Exception as e:”，一共三处。
将“except vim.fault.ToolsUnavailable, e:”修改为“except vim.fault.ToolsUnavailable as e:”，一个一处
因为脚本事python2.7的语法，但是6.0已经是python3
```
vi auto-shutdown.sh

#! /bin/ash
echo "shutting down VMs..please wait.."
/vmfs/volumes/datastore1/poweroffvms
echo "done."
echo "shutting down the host now.."
/bin/poweroff

chmod +x /vmfs/volumes/datastore1/auto-shutdown.sh

编辑  /etc/rc.local.d/local.sh  在其文件文本 末尾的"exit 0" 之前

/bin/kill $(cat /var/run/crond.pid)

/bin/echo "30   0    *   *   *   /vmfs/volumes/datastore1/auto-shutdown.sh" >> /var/spool/cron/crontabs/root

/usr/lib/vmware/busybox/bin/busybox crond

```
