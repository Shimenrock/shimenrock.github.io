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

> * Oracle12cR2=12.2.0.1 
> * Oracle18C=12.2.0.2 
> * Oracle19c=12.2.0.3

<div class="notice">
  <p>参考文档：</p>
  <p><a href="https://docs.oracle.com/en/database/oracle/oracle-database/19/ladbi/index.html">https://docs.oracle.com/en/database/oracle/oracle-database/19/ladbi/index.html</a>
  </p>
</div>

评论：19.3.0.0.0 这个版本的安装方式，真是一股扑面而来的印度编程风，安装介质的目录就是ORACLE_BASE，如果你把介质放到了tmp目录，安装过程不注意的话，数据库软件就安装到了tmp目录，真是莫名其妙。

### 1 Oracle Database Installation Checklist

这一部分主要是操作系统版本要求：

> * The following Linux x86-64 kernels are supported: 
> * Oracle Linux 7.4 with the Unbreakable Enterprise Kernel 4: 4.1.12-124.19.2.el7uek.x86_64 or later
> * Oracle Linux 7.4 with the Unbreakable Enterprise Kernel 5: 4.14.35-1818.1.6.el7uek.x86_64 or later
> * Oracle Linux 7.5 with the Red Hat Compatible kernel: 3.10.0-862.11.6.el7.x86_64 or later
> * Red Hat Enterprise Linux 7.5: 3.10.0-862.11.6.el7.x86_64 or later
> * SUSE Linux Enterprise Server 12 SP3: 4.4.103-92.56-default or later

### 2 Checking and Configuring Server Hardware for Oracle Database

### 3 Automatically Configuring Oracle Linux with Oracle Preinstallation RPM

### 4 Configuring Operating Systems for Oracle Database on Linux

#### 4.1 解决软件依赖关系

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



#### 4.2 Confirming Host Name Resolution

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

#### 4.3 关闭防火墙和selinux

```
systemctl stop firewalld.service
systemctl disable firewalld.service

vi /etc/selinux/config
修改
SELINUX=disabled
```

#### 4.4 修改系统参数

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

### 5 Configuring Users, Groups and Environments for Oracle Grid Infrastructure and Oracle Database

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

### 6 Configuring Networks for Oracle Database

### 7 Supported Storage Options for Oracle Database and Oracle Grid Infrastructure

### 8 Configuring File System Storage for Oracle Database

### 9 Configuring Storage for Oracle Grid Infrastructure for a Standalone Server

### 10 Installing and Configuring Oracle Grid Infrastructure for a Standalone Server

### 11 Installing Oracle Database

安装介质下载：https://edelivery.oracle.com

> * V982063-01.zip  ORACLE Database 19.3.0.0.0 for Linux x86-64
> * V982064-01.zip  ORACLE Database Client 19.3.0.0.0 for Linux x86-64
> * V982068-01.zip  ORACLE Database Grid Infrastructure 19.3.0.0.0 for Linux x86-64

建立安装目录

```
mkdir -p /u01/app/oracle
mkdir –p /oradata
mkdir -p /u01/oraInventory
chown -R oracle:oinstall /u01/ /oradata/ /u01/oraInventory/ /archivelog/
chmod -R 775 /u01/app/oracle
```
		
设置环境变量

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

```
# su - oracle
$ cd /tmp
$ mkdir -p /u01/app/oracle/product/19.0.0/db_1 
$ unzip V982063-01.zip -d /u01/app/oracle/product/19.0.0/db_1 

$ export DISPLAY=192.168.11.211:0.0
$ cd /u01/app/oracle/product/19.0.0/db_1 
$ export CV_ASSUME_DISTID=RHEL7.6

图形安装
$ ./runInstaller
静默安装
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

#### 创建数据库

图形创建

```
dbca
```

静默创建

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

#### 创建监听

图形创建

```
netca
```

静默创建

```
netca -silent -responseFile /u01/app/oracle/product/19.0.0/db_1/assistants/netca/netca.rsp
```

#### 配置自启动

<div class="notice">
  <p>Centos7以后/etc/rc.local已经没有执行权限，因为这个文件是为了兼容性的问题而添加的，建议创建自己的systemd服务。</p>
</div>

```
vi /etc/oratab
orcl:/u01/app/oracle/product/19.0.0/db_1:Y
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
```

#### 生产系统配置

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


