---
title: "Oracle Database 19c Installation On CentOS 8"
published: true
related: true
header:
  teaser: /assets/images/2020-03-25-oracle19c.png
toc: true
toc_sticky: true
categories: 
  - oracle
tags: 
  - oracle
---  

继上一个11.2.0.4稳定版后，19c是12c产品的稳定版本了，可以用于生产系统了。

- Oracle12cR2=12.2.0.1 
- Oracle18C=12.2.0.2 
- Oracle19c=12.2.0.3 

<div class="notice">
  <p>参考文档：</p>
  <p><a href="https://docs.oracle.com/en/database/oracle/oracle-database/19/ladbi/index.html">https://docs.oracle.com/en/database/oracle/oracle-database/19/ladbi/index.html</a>
  </p>
</div>

注意 ：19.3.0.0.0 安装介质的目录就是ORACLE_BASE，如果你把介质放到了tmp目录，数据库软件就安装到了tmp目录。

## 1. Oracle Database Installation Checklist

<div class="notice">
  <p>这一部分主要是操作系统版本要求，The following Linux x86-64 kernels are supported:</p>
   <ul>
    <li>Oracle Linux 7.4 with the Unbreakable Enterprise Kernel 4: 4.1.12-124.19.2.el7uek.x86_64 or later</li>
    <li>Oracle Linux 7.4 with the Unbreakable Enterprise Kernel 5: 4.14.35-1818.1.6.el7uek.x86_64 or later</li>
    <li>Oracle Linux 7.5 with the Red Hat Compatible kernel: 3.10.0-862.11.6.el7.x86_64 or later</li>
    <li>Red Hat Enterprise Linux 7.5: 3.10.0-862.11.6.el7.x86_64 or later</li>
    <li>SUSE Linux Enterprise Server 12 SP3: 4.4.103-92.56-default or later</li>
  </ul>
</div>

## 2. Checking and Configuring Server Hardware for Oracle Database

## 3. Automatically Configuring Oracle Linux with Oracle Preinstallation RPM

## 4. Configuring Operating Systems for Oracle Database on Linux

### 4.1 解决软件依赖关系

```
# mkdir /media/cdrom
# mount /dev/cdrom /media/cdrom/
或着挂载ISO文件
# mount -o loop /mnt/iso/CentOS-8-x86_64-1905-dvd1.iso /mnt/cdrom
# cd /etc/yum.repos.d/
# vi CentOS-Media.repo
修改
[c8-media-BaseOS]
name=CentOS-BaseOS-$releasever - Media
baseurl=file:///media/cdrom/BaseOS
gpgcheck=0
enabled=1
gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-centosofficial

[c8-media-AppStream]
name=CentOS-AppStream-$releasever - Media
baseurl=file:///media/cdrom/AppStream
gpgcheck=0
enabled=1
gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-centosofficial
# vi CentOS-Base.repo
# vi CentOS-AppStream.repo 
# vi CentOS-Extras.repo
修改
enabled=0
# yum repolist all
# yum clean all
# yum makecache
```
```
# dnf install -y bc binutils  elfutils-libelf elfutils-libelf-devel fontconfig-devel glibc glibc-devel ksh 
# dnf install -y libaio libaio-devel libX11 libXau libXi libXtst libXrender libXrender-devel
# dnf install -y libgcc libstdc++ libstdc++-devel libxcb make 
# dnf install -y nfs-utils smartmontools sysstat tigervnc*
# dnf install -y librdmacm-devel unixODBC libnsl2 libnsl2.i686  libnsl* 
// for Oracle RAC and Oracle Clusterware
# dnf install -y net-tools 

// for Oracle ACFS 
# dnf install -y nfs-utils python  python-configshell  python-rtslib  python-six  targetcli 
// Centos8 必须选择 python2 还是 python36，系统默认安装的是python3，dnf search python
#  dnf install -y python36 python3-configshell python3-rtslib python3-six 

// [文档 ID 2254198.1]Missing or Ignored package compat-libstdc+±33-3.2.3 causes Text Issues
// CentOS 7、CentOS 8的安装包中不含compat-libstdc++* 和compat-libcap1，如果不使用 Oracle Text 那么可以忽略该包
// 这里在https://centos.pkgs.org 搜索下载后安装
#  rpm -ivh compat-lib*
```
下载：   
[compat-libstdc++-33-3.2.3-72.el7.x86_64.rpm](https://github.com/Shimenrock/shimenrock.github.io/blob/master/assets/down/compat-libstdc++-33-3.2.3-72.el7.x86_64.rpm?raw=true)

[compat-libcap1-1.10-7.el7.x86_64.rpm](https://github.com/Shimenrock/shimenrock.github.io/blob/master/assets/down/compat-libcap1-1.10-7.el7.x86_64.rpm?raw=true)



### 4.2 Confirming Host Name Resolution

```
# hostnamectl status
   Static hostname: ORACLE-212
         Icon name: computer-vm
           Chassis: vm
        Machine ID: e46e35085ec243108381fd151ac72b2a
           Boot ID: 797eacaa8e4d47238a0a8240597d6067
    Virtualization: vmware
  Operating System: CentOS Linux 8 (Core)
       CPE OS Name: cpe:/o:centos:centos:8
            Kernel: Linux 4.18.0-147.el8.x86_64
      Architecture: x86-64
# ping ORACLE-212
PING ORACLE-212(ORACLE-212 (fe80::cf8a:fc1d:564c:cc8d%ens192)) 56 data bytes
64 bytes from ORACLE-212 (fe80::cf8a:fc1d:564c:cc8d%ens192): icmp_seq=1 ttl=64 time=0.031 ms
#  vi /etc/hosts
增加 192.168.11.212 ORACLE-212
# ping ORACLE-212
PING ORACLE-212 (192.168.11.212) 56(84) bytes of data.
64 bytes from ORACLE-212 (192.168.11.212): icmp_seq=1 ttl=64 time=0.051 ms
```

### 4.3 关闭防火墙和selinux

```
systemctl stop firewalld.service
systemctl disable firewalld.service

vi /etc/selinux/config
修改
SELINUX=disabled
```

### 4.4 修改系统参数

```
# vi /etc/sysctl.conf     #8G内存

fs.file-max = 6815744
kernel.sem = 250 32000 100 128
kernel.shmmni = 4096
kernel.shmall = 1048576
kernel.shmmax = 4294967296
kernel.panic_on_oops = 1
net.core.rmem_default = 262144
net.core.rmem_max = 4194304
net.core.wmem_default = 262144
net.core.wmem_max = 1048576
net.ipv4.conf.all.rp_filter = 2
net.ipv4.conf.default.rp_filter = 2
fs.aio-max-nr = 1048576
net.ipv4.ip_local_port_range = 9000 65500

# sysctl -p
```

```
# vi /etc/pam.d/login
session    required    /lib64/security/pam_limits.so	
```

## 5 Configuring Users, Groups and Environments for Oracle Grid Infrastructure and Oracle Database

| | |
| --- | --- |
| 数据库软件安装 | oinstall |
| ASM存储管理组  | asmdba |
| ASM启动关闭管理组 | asmoper |
| 数据库管理组包括软件安装权限 | dba |
| 数据库有限特权管理组  | oper |
| 数据库备份管理组 | backupdba |
| Data Guard管理组  | dgdba |
| 数据加密管理组 | kmdba |
| Oracle RAC集群管理组 | racdba |

```
/usr/sbin/groupadd -g 54321 oinstall
/usr/sbin/groupadd -g 54327 asmdba
/usr/sbin/groupadd -g 54328 asmoper
/usr/sbin/groupadd -g 54322 dba
/usr/sbin/groupadd -g 54323 oper
/usr/sbin/groupadd -g 54324 backupdba
/usr/sbin/groupadd -g 54325 dgdba
/usr/sbin/groupadd -g 54326 kmdba
/usr/sbin/groupadd -g 54330 racdba
```
```
/usr/sbin/useradd -u 54321 -g oinstall -G dba,asmdba,backupdba,dgdba,kmdba,racdba oracle
/usr/sbin/useradd -u 54331 -g oinstall -G dba,asmdba,backupdba,dgdba,kmdba,racdba grid
/usr/sbin/usermod -g oinstall -G dba,asmdba,backupdba,dgdba,kmdba,racdba,oper oracle
# passwd oracle

$ id oracle
uid=54321(oracle) gid=54321(oinstall) groups=54321(oinstall),54322(dba), 
54323(oper),54324(backupdba),54325(dgdba),54326(kmdba),54327(asmdba),54330(racdba)
$ id grid
uid=54331(grid) gid=54321(oinstall) groups=54321(oinstall),54322(dba),
54327(asmdba),54328(asmoper),54329(asmadmin),54330(racdba)
```

```
# vi /etc/security/limits.conf

# Open file descriptors “ulimit -Sn” “ulimit -Hn”
oracle   soft   nofile    131072
oracle   hard   nofile    131072
# Number of processes available to a single user “ulimit -Su” “ulimit -Hu”
oracle   soft   nproc    131072
oracle   hard   nproc    131072
# Size of the stack segment of the process “ulimit -Ss” “ulimit -Hs”
oracle   soft   stack    10240
oracle   hard   stack    32768
# Maximum locked memory limit 最大锁定内存限制，HugePages至少要占当前内存的90％；禁用HugePages内存时，至少3 GB
oracle   hard   memlock    3500000
oracle   soft   memlock    3500000
```

## 6 Configuring Networks for Oracle Database

## 7 Supported Storage Options for Oracle Database and Oracle Grid Infrastructure

## 8 Configuring File System Storage for Oracle Database

## 9 Configuring Storage for Oracle Grid Infrastructure for a Standalone Server

## 10 Installing and Configuring Oracle Grid Infrastructure for a Standalone Server

## 11 Installing Oracle Database

### 11.1 安装介质下载

https://edelivery.oracle.com

> * V982063-01.zip  ORACLE Database 19.3.0.0.0 for Linux x86-64
> * V982064-01.zip  ORACLE Database Client 19.3.0.0.0 for Linux x86-64
> * V982068-01.zip  ORACLE Database Grid Infrastructure 19.3.0.0.0 for Linux x86-64

### 11.2 建立安装目录

```
mkdir -p /u01/app/oracle
mkdir –p /oradata
mkdir -p /u01/oraInventory
chown -R oracle:oinstall /u01/ /oradata/ /u01/oraInventory/ /archivelog/
chmod -R 775 /u01/app/oracle
```
		
### 11.3 设置环境变量

```
# su - oracle
$ vi ~/.bash_profile
# User oracle specific environment and startup programs
export TMP=/tmp
export TMPDIR=\$TMP

export ORACLE_SID=CDB19C 
export ORACLE_BASE=/u01/app/oracle 
export ORACLE_HOME=$ORACLE_BASE/product/19.0.0/db_1 
export ORA_INVENTORY=/u01/app/oraInventory 
export PDB_NAME=ORDERS  
export TNS_ADMIN=$ORACLE_HOME/network/admin
export ORACLE_TERM=xterm

export PATH=$ORACLE_HOME/bin:$PATH 
export LD_LIBRARY_PATH=$ORACLE_HOME/lib:/lib:/usr/lib 
export CLASSPATH=$ORACLE_HOME/jdk/jre:$ORACLE_HOME/jlib:$ORACLE_HOME/rdbms/jlib 

export NLS_DATE_FORMAT='yyyy-mm-dd hh24:mi:ss'
export LANG=en-US
export NLS_LANG=AMERICAN_AMERICA.ZHS16GBK
export PS1=`whoami`@`hostname`$

stty erase "^H"
stty erase ^?

umask 022
# 确保执行软件安装的用户创建具有644权限的文件

$ source ~/.bash_profile
echo $ORACLE_HOME				
echo $ORACLE_SID
```

### 11.4 图形安装

```
# su - oracle
$ cd /tmp
$ mkdir -p /u01/app/oracle/product/19.0.0/db_1 
$ unzip V982063-01.zip -d /u01/app/oracle/product/19.0.0/db_1 

$ export DISPLAY=192.168.11.211:0.0
$ cd /u01/app/oracle/product/19.0.0/db_1 
$ export CV_ASSUME_DISTID=RHEL7.6
```

$ ./runInstaller

### 11.5 静默安装

```
./runInstaller -ignorePrereq -waitforcompletion -silent \
-responseFile /u01/app/oracle/product/19.0.0/db_1/install/response/db_install.rsp \
oracle.install.option=INSTALL_DB_SWONLY \
ORACLE_HOSTNAME=henry \
UNIX_GROUP_NAME=oinstall \
INVENTORY_LOCATION=/u01/app/oraInventory \
SELECTED_LANGUAGES=en,en_GB \
ORACLE_HOME=/u01/app/oracle/product/19.0.0/db_1 \
ORACLE_BASE=/u01/app/oracle \
oracle.install.db.InstallEdition=EE \
oracle.install.db.OSDBA_GROUP=dba \
oracle.install.db.OSOPER_GROUP=oper \
oracle.install.db.OSBACKUPDBA_GROUP=dba \
oracle.install.db.OSDGDBA_GROUP=dba \
oracle.install.db.OSKMDBA_GROUP=dba \
oracle.install.db.OSRACDBA_GROUP=dba \
SECURITY_UPDATES_VIA_MYORACLESUPPORT=false \
DECLINE_SECURITY_UPDATES=true

As a root user, execute the following script(s):
    1. /u01/app/oraInventory/orainstRoot.sh
    2. /u01/app/oracle/product/19.0.0/db_1/root.sh
```

### 11.6 创建数据库

首先启动监听程序

```
lsnrctl start 
```

图形创建

```
export DISPLAY=X.X.X.X:0.0
dbca
```

<div class="notice">
  <p>DBCA</p>
   <ul>
    <li>1.Database Operation ：增加了管理可插拔数据库选项</li>
    <li>2.Creaton Mode : 选择Advanced configuration</li>
    <li>3.Deployment Type : 选择通用事务型数据库</li>
    <li>4.Database Identification : 不勾选Create as Container database 创建的即non-CDB 非容器数据库架构。</li>
    <li>5.Storage Option ：选择文件</li>
    <li>6.Fast Recovery Option ：默认</li>
    <li>7.Network Configuration ：监听状态为 UP，方便后面创建EM</li>
    <li>8.Data Vault Option ：默认</li>
    <li>9.Configuration Option ：根据业务对于进行配置，Sample schemas 测试对象</li>
    <li>10.Management Options ：EM(EM express、EM cloud control、老版 database control)</li>
    <li>11.User Credentials ：sys、system、pdbadmin 先设置测试密码oracle</li>
    <li>12.Creation Option ：All Initalization Parameters 修改初始化参数</li>
    <li>12.Creation Option ：Custormize Storage Locations 修改控制文件、表空间文件、redo日志文件路径</li>
    <li>13.summary ：汇总信息，创建静默安装响应文件。</li>
  </ul>
</div>
 
<div class="notice">
  <p>Storage Option这里自定义配置</p>
   <ul>
    <li>Global database name： cdb1.oracle.com</li>
    <li>SID: cdb1</li>
    <li>Create as Container database 勾选创建PDB</li>
    <li>Use Local Undo tablespace 勾选pdb拥有自己的本地undo表空间 </li>
    <li>PDB: pdb_easyee</li>
  </ul>
</div>

### 11.7 静默创建数据库

```
dbca -silent -createDatabase -templateName General_Purpose.dbc \
-gdbname CDB19C \
-sid CDB19C \
-databaseConfigType SI \
-createAsContainerDatabase TRUE \
-numberOfPDBs 1 \
-pdbName ORDERS \
-useLocalUndoForPDBs TRUE \
-pdbAdminPassword oracle \
-sysPassword oracle \
-systemPassword oracle \
-characterSet AL32UTF8 \
-memoryPercentage 40
```

### 11.8 配置CDB自启动

<div class="notice">
  <p>Centos7以后/etc/rc.local已经没有执行权限，因为这个文件是为了兼容性的问题而添加的，建议创建自己的systemd服务。</p>
</div>

```
vi /etc/oratab
cdb1:/u01/app/oracle/product/19.0.0/db_1:Y
```
```
vi /usr/bin/ora_auto_start.sh

#! /bin/bash 
#  script  For oracle19c.service
/u01/app/oracle/product/19.0.0/db_1/bin/lsnrctl start
/u01/app/oracle/product/19.0.0/db_1/bin/dbstart /u01/app/oracle/product/19.0.0/db_1 

chmod +x /usr/bin/ora_auto_start.sh
```
```
vi /etc/systemd/system/oracle19c.service

[Unit]
Description=Oracle19c
After=syslog.target network.target
[Service]
LimitMEMLOCK=infinity
LimitNOFILE=65535
Type=oneshot
RemainAfterExit=yes
User=oracle
Environment="ORACLE_HOME=/u01/app/oracle/product/19.0.0/db_1"
ExecStart=/usr/bin/ora_auto_start.sh
[Install]
WantedBy=multi-user.target
```
```
systemctl enable oracle19c

reboot

systemctl status oracle19c

SQL> show con_name
```

### 11.8 配置PDB自启动

```
$export ORACLE_SID=cdb1
$sqlplus / as sysdba
SQL> show pdbs

    CON_ID CON_NAME                       OPEN MODE  RESTRICTED
---------- ------------------------------ ---------- ----------
         2 PDB$SEED                       READ ONLY  NO
         3 PDB_EASYEE                     MOUNTED

set linesize 300
set pagesize 120
column NAME format a30
col pdb format a30
select PDB, INST_ID, NAME from gv$services order by 1;
select CON_ID, NAME, OPEN_MODE from V$PDBS;
select con_id, dbid, guid, name , open_mode from v$pdbs;
```
**方法1**
```
alter pluggable database PDB_EASYEE open;
或着
alter session set container=pdb_easyee;
startup

alter session set container=CDB$ROOT;
alter pluggable database PDB_EASYEE save state;
alter pluggable database PDB_EASYEE close immediate;
```
**方法2**
```
 CREATE OR REPLACE TRIGGER open_pdbs
AFTER STARTUP ON DATABASE
BEGIN
EXECUTE IMMEDIATE 'ALTER PLUGGABLE DATABASE ALL OPEN';
END open_pdbs;
/
```

1.启动或者关闭一个或多个 PDB，指定的名称为一个以逗号分隔的列表。
```
ALTER PLUGGABLE DATABASE pdb1,pdb2 OPEN READ ONLY FORCE;
ALTER PLUGGABLE DATABASE pdb1,pdb2 CLOSE IMMEDIATE;
```
2.启动或者关闭 all pdbs
```
ALTER PLUGGABLE DATABASE ALL OPEN;
ALTER PLUGGABLE DATABASE ALL CLOSE IMMEDIATE;
```
3.排除某些pdbs不启动，启动所有为排除的。
```
ALTER PLUGGABLE DATABASE ALL EXCEPT pdb1 OPEN;
ALTER PLUGGABLE DATABASE ALL EXCEPT pdb1 CLOSE IMMEDIATE;
```

### 11.9 生产系统配置

> * 密码有效期限制
```   
ALTER PROFILE DEFAULT LIMIT PASSWORD_LIFE_TIME UNLIMITED;
```
> * 密码区分大小写
```
alter system set sec_case_sensitive_logon = false;
```
> * 密码错误10次自动锁定
```
ALTER PROFILE DEFAULT LIMIT FAILED_LOGIN_ATTEMPTS UNLIMITED;
```

### 11.10 PLSQL Developer配置

```
$lsnrctl status

LSNRCTL for Linux: Version 19.0.0.0.0 - Production on 02-APR-2020 05:18:04

Copyright (c) 1991, 2019, Oracle.  All rights reserved.

Connecting to (DESCRIPTION=(ADDRESS=(PROTOCOL=TCP)(HOST=ORACLE-212)(PORT=1521)))
STATUS of the LISTENER
------------------------
Alias                     LISTENER
Version                   TNSLSNR for Linux: Version 19.0.0.0.0 - Production
Start Date                02-APR-2020 05:08:25
Uptime                    0 days 0 hr. 9 min. 39 sec
Trace Level               off
Security                  ON: Local OS Authentication
SNMP                      OFF
Listener Parameter File   /u01/app/oracle/product/19.0.0/db_1/network/admin/listener.ora
Listener Log File         /u01/app/oracle/diag/tnslsnr/ORACLE-212/listener/alert/log.xml
Listening Endpoints Summary...
  (DESCRIPTION=(ADDRESS=(PROTOCOL=tcp)(HOST=ORACLE-212)(PORT=1521)))
  (DESCRIPTION=(ADDRESS=(PROTOCOL=ipc)(KEY=EXTPROC1521)))
Services Summary...
Service "86b637b62fdf7a65e053f706e80a27ca.oracle.com" has 1 instance(s).
  Instance "cdb1", status READY, has 1 handler(s) for this service...
Service "a24b3a2eca10193de053d40ba8c04bf1.oracle.com" has 1 instance(s).
  Instance "cdb1", status READY, has 1 handler(s) for this service...
Service "cdb1.oracle.com" has 1 instance(s).
  Instance "cdb1", status READY, has 1 handler(s) for this service...
Service "cdb1XDB.oracle.com" has 1 instance(s).
  Instance "cdb1", status READY, has 1 handler(s) for this service...
Service "pdb_easyee.oracle.com" has 1 instance(s).
  Instance "cdb1", status READY, has 1 handler(s) for this service...
The command completed successfully
oracle@ORACLE-212$
```

创建 tnsname ： CDB1 服务名 cdb1.oracle.com
创建 tnsname :  PDB_EASYEE 服务名 pdb_easyee.oracle.com

测试登陆 sys/oracle system/oracle  pdbadmin/oracle

## 多组合容器数据库优点

- 强大硬件资源有富余（一体机），有利于资源充分利用，比较创建多实例方案，可插拔数据库便于管理；同理硬件资源有限，考虑创建容器数据库是否有意义？
- 多database方案，内存不能共享使用
- 数据库上云后，多database方案和多schema方案不利于云管理。
- 多租户优点1：以更低成本集中管理多个数据库，实例开销比较低，存储成本较低
- 多租户优点2：降低DBA资源成本并维护安全性
- 多租户优点3：提供隔离

## 容器 V$CONTAINERS 

- 根容器 （在创建CDB时创建的第一个容器）
  - 包含ORACLE系统提供的公用对象和元数据
  - 包含ORACLE系统提供的公用用户和角色(公共账号必须c##开头，可以访问CDB和PDB)
- 可插拔数据库容器PDB
  - 应用程序的容器
    - 包含表空间 （永久和临时）
    - 方案/对象/权限
    - 已创建、克隆、移走、插入
  - 特点种子PDB
    - PDB$SEED 提供新PDB的快速预配

## CDB与非CDB区别

非CDB中，用户和ORACLE的数据、元数据、数据字典属于混合存储
容器数据库，分离SYSTEM和用户数据
  - 系统的容器： ORACLE元数据
  - 应用程序容器：用户元数据，用户数据

## 数据字典视图

- CDB_xxx  所有PDB中多租户容器数据库中的所有对象
  - DBA_xxx 容器或可插插数据库中的所有对象
    - ALL_xx 可由当前用户访问的对象
      - USER_xxx 当前用户拥有的对象
  select view_name from dba_views where view_name like 'CDB%';

- CDB_pdbs CDB中所有的PDB
- CDB_tablespaces  CDB中的所有表空间
- CDB_users CDB内的所有用户

select table_name FROM dict where table_name like 'DBA%'

## PDB命令行创建

方式1：使用PDB$SEED种子容器生成副本
```
数据文件创建到DB_CREATE_FILE_DEST
create pluggable database pdb1 admin user pdb1admin identified by pdb1admin;
指定文件创建路径
create pluggable database pdb1 admin user pdb1admin identified by pdb1admin create_file_dest='+PDBDATA';
```
方式2： 将PDB$SEED的数据文件目录进行转换
```
CREATE PLUGGABLE DATABASE pdb_easyee 
  ADMIN USER adm_easyee IDENTIFIED BY easyee2020
  ROLES = (dba)
  DEFAULT TABLESPACE PDB_EASYEE
    DATAFILE '/oradata/orcl/pdbeasyee/easyee01.dbf' SIZE 250M AUTOEXTEND ON
  FILE_NAME_CONVERT = ('/u01/app/oracle/oradata/ORCL/pdbseed/',
                       '/oradata/orcl/pdbeasyee/')
  STORAGE (MAXSIZE 2G)
  PATH_PREFIX = '/oradata/orcl/pdbeasyee/';

```
PDB默认创建SYSTEM,SYSAUX,TEMP,UNDO表空间