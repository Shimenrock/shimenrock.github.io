---
title: "esxi虚拟机增加裸硬盘"
published: true
related: true
header:
  teaser: /assets/images/2020-03-16-wmware.jpg
categories: 
  - esxi
---
<!-- https://kb.vmware.com/s/article/1017530?lang=zh_CN#q=%E8%A3%B8  -->

# 1.打开与 ESXi/ESX 主机的 SSH 服务。

# 2.运行以下命令列出已连接到 ESXi 主机的磁盘：

```
# ls -l /vmfs/devices/disks
total 1945740720
-rw-------    1 root     root     480103981056 Mar 16 21:14 t10.ATA_____Colorful_SL500_480GB____________________AC20190723A0100602__
-rw-------    1 root     root     500107862016 Mar 16 21:14 t10.ATA_____ST500DM0022D1BD142___________________________________W2AKNKPK
-rw-------    1 root     root     500105740800 Mar 16 21:14 t10.ATA_____ST500DM0022D1BD142___________________________________W2AKNKPK:1
-rw-------    1 root     root     256060514304 Mar 16 21:14 t10.NVMe____Colorful_CN600_256GB____________________0000000000000001
-rw-------    1 root     root       4161536 Mar 16 21:14 t10.NVMe____Colorful_CN600_256GB____________________0000000000000001:1
-rw-------    1 root     root     4293918720 Mar 16 21:14 t10.NVMe____Colorful_CN600_256GB____________________0000000000000001:2
-rw-------    1 root     root     248138505728 Mar 16 21:14 t10.NVMe____Colorful_CN600_256GB____________________0000000000000001:3
-rw-------    1 root     root     262127616 Mar 16 21:14 t10.NVMe____Colorful_CN600_256GB____________________0000000000000001:5
-rw-------    1 root     root     262127616 Mar 16 21:14 t10.NVMe____Colorful_CN600_256GB____________________0000000000000001:6
-rw-------    1 root     root     115326976 Mar 16 21:14 t10.NVMe____Colorful_CN600_256GB____________________0000000000000001:7
-rw-------    1 root     root     299876352 Mar 16 21:14 t10.NVMe____Colorful_CN600_256GB____________________0000000000000001:8
-rw-------    1 root     root     2684354560 Mar 16 21:14 t10.NVMe____Colorful_CN600_256GB____________________0000000000000001:9
```

# 3.在列表中，标识要配置为 RDM 的本地设备，并复制设备名称。

注意：设备名称将很可能使用 t10. 前缀，类似于以下名称：t10.F405E46494C4540046F455B64787D285941707D203F45765

# 4.要将设备配置为 RDM，并将 RDM 指针文件输出到您所选的目标，请运行以下命令：

```
# vmkfstools -z /vmfs/devices/disks/diskname /vmfs/volumes/datastorename/vmfolder/vmname.vmdk

例如：

# vmkfstools -z /vmfs/devices/disks/t10.ATA_____Colorful_SL500_480GB____________________AC20190723A0100602__ /vmfs/volumes/datastore1/WIN10-211/SL500.vmdk
```

注意：新创建的 RDM 指针文件的大小显示为与其映射到的裸设备相同的大小，它是一个虚拟文件，不占用任何存储空间。

```
# ls -Alh
total 0
-rw-------    1 root     root      447.1G Mar 16 21:21 SL500-rdmp.vmdk
-rw-------    1 root     root         473 Mar 16 21:21 SL500.vmdk
-rw-r--r--    1 root     root           0 Mar 16 21:20 WIN10-211.vmsd
-rwxr-xr-x    1 root     root        1.9K Mar 16 21:20 WIN10-211.vmx

# cat SL500.vmdk 
# Disk DescriptorFile
version=1
encoding="UTF-8"
CID=fffffffe
parentCID=ffffffff
createType="vmfsPassthroughRawDeviceMap"

# Extent description
RW 937703088 VMFSRDM "SL500-rdmp.vmdk"

# The Disk Data Base 
#DDB

ddb.adapterType = "lsilogic"
ddb.geometry.cylinders = "58369"
ddb.geometry.heads = "255"
ddb.geometry.sectors = "63"
ddb.longContentID = "fd23004d46cc50146cf6c777fffffffe"
ddb.uuid = "60 00 C2 9f 28 1e 65 bf-3d cc 5e f0 f5 42 1e f6"
ddb.virtualHWVersion = "14"
```

# 5.如果已创建 RDM 指针文件，请使用 vSphere Client 将 RDM 连接到虚拟机：

1. 右键单击要将 RDM 磁盘添加到的虚拟机。
2. 单击编辑设置。
3. 单击添加...
4. 选择硬盘。
5. 选择使用现有虚拟磁盘。
6. 浏览到步骤 5 中保存 RDM 指针的目录，选择 RDM 指针文件，然后单击下一步。
7. 选择要将磁盘连接到的虚拟 SCSI 控制器，然后单击下一步。
8. 单击完成。

# 6.现在，虚拟机清单中的新硬盘应显示为映射的裸 LUN。
注意：
由于此虚拟机现有已连接的本地磁盘迁移，因此无法使用 vMotion。
如果需要从虚拟机中移除本地 RDM 映射，只需对共享存储 RDM 应用相同步骤。在 vSphere Client 中，右键单击虚拟机，单击编辑设置，选择 RDM 然后单击从磁盘删除。此操作不会删除磁盘中的数据，而仅删除 RDM 映射文件。

