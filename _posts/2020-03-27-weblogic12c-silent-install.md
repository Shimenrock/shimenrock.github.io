---
title: "WebLogic 12c Silent Installation"
published: true
related: true
header:
  teaser: /assets/images/2020-03-27-weblogic12c.jpg
toc: true
toc_sticky: true
categories: 
  - oracle
tags: 
  - oracle
---  

### Install JDK

```
# tar -xvf jdk-8u241-linux-x64.tar.gz 
# mkdir -p /usr/java
# mv jdk1.8.0_241 /usr/java/
# alternatives --install /usr/bin/java java /usr/java/jdk1.8.0_241/bin/java 1
# alternatives --config java
# java -version
java version "1.8.0_241"
Java(TM) SE Runtime Environment (build 1.8.0_241-b07)
Java HotSpot(TM) 64-Bit Server VM (build 25.241-b07, mixed mode)
```

### Confirming Host Name Resolution

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

### Create a new group and user

```
groupadd -g 54321 oinstall
useradd -u 54321 -g oinstall oracle
passwd oracle
```

### Create the directories in which the Oracle software will be installed

```
mkdir -p /u01/app/oracle/middleware
mkdir -p /u01/app/oracle/config/domains
mkdir -p /u01/app/oracle/config/applications
mkdir -p /u01/software
chown -R oracle:oinstall /u01/app/oracle/middleware  /u01/app/oracle/config/
chmod -R 775 /u01/app/oracle/middleware /u01/app/oracle/config/
```

### set environment

```
vi /home/oracle/.bash_profile
export MW_HOME=/u01/app/oracle/middleware
export WLS_HOME=$MW_HOME/wlserver
export WL_HOME=$WLS_HOME

export PATH=$JAVA_HOME/bin:$PATH
```

```
# unzip V983364-01.zip 
Archive:  V983364-01.zip
  inflating: fmw_12.2.1.4.0_wls.jar  
  inflating: fmw_12214_readme.html 
```

### Create Response File

```
su - oracle

vi /u01/software/wls.rsp

[ENGINE]
Response File Version=1.0.0.0.0
[GENERIC]
ORACLE_HOME=/u01/app/oracle/middleware
INSTALL_TYPE=WebLogic Server
MYORACLESUPPORT_USERNAME=
MYORACLESUPPORT_PASSWORD=<SECURE VALUE>
DECLINE_SECURITY_UPDATES=true
SECURITY_UPDATES_VIA_MYORACLESUPPORT=false
PROXY_HOST=
PROXY_PORT=
PROXY_USER=
PROXY_PWD=<SECURE VALUE>
COLLECTOR_SUPPORTHUB_URL=

vi /u01/software/oraInst.loc

inventory_loc=/u01/app/oraInventory
inst_group=oinstall
```

### WebLogic Silent Installation

```
$JAVA_HOME/bin/java -Xmx1024m -jar /jdk8/fmw_12.2.1.4.0_wls.jar -silent -responseFile /u01/software/wls.rsp -invPtrLoc /u01/software/oraInst.loc

Launcher log file is /tmp/OraInstall2020-03-28_02-10-06AM/launcher2020-03-28_02-10-06AM.log.
Extracting the installer . . . . . Done
Checking if CPU speed is above 300 MHz.   Actual 3600.000 MHz    Passed
Checking swap space: must be greater than 512 MB.   Actual 8103 MB    Passed
Checking if this platform requires a 64-bit JVM.   Actual 64    Passed (64-bit not required)
Checking temp space: must be greater than 300 MB.   Actual 147579 MB    Passed
Preparing to launch the Oracle Universal Installer from /tmp/OraInstall2020-03-28_02-10-06AM
Log: /tmp/OraInstall2020-03-28_02-10-06AM/install2020-03-28_02-10-06AM.log
Copyright (c) 1996, 2019, Oracle and/or its affiliates. All rights reserved.
Reading response file..
Skipping Software Updates
Starting check : CertifiedVersions
Expected result: One of oracle-6, oracle-7, redhat-7, redhat-6, SuSE-11, SuSE-12, SuSE-15
Actual Result: redhat-null
Check complete. The overall result of this check is: Passed
CertifiedVersions Check: Success.


Starting check : CheckJDKVersion
Expected result: 1.8.0_191
Actual Result: 1.8.0_241
Check complete. The overall result of this check is: Passed
CheckJDKVersion Check: Success.


Validations are enabled for this session.
Verifying data
Copying Files
Percent Complete : 10
Percent Complete : 20
Percent Complete : 30
Percent Complete : 40
Percent Complete : 50
Percent Complete : 60
Percent Complete : 70
Percent Complete : 80
Percent Complete : 90
Percent Complete : 100

The installation of Oracle Fusion Middleware 12c WebLogic Server and Coherence 12.2.1.4.0 completed successfully.
Logs successfully copied to /u01/app/oraInventory/logs.
```

```
oracle@ORACLE-212$. $WLS_HOME/server/bin/setWLSEnv.sh
CLASSPATH=/usr/java/jdk1.8.0_241/lib/tools.jar:/u01/app/oracle/middleware/wlserver/modules/features/wlst.wls.classpath.jar:/u01/app/oracle/product/19.0.0/db_1/jdk/jre:/u01/app/oracle/product/19.0.0/db_1/jlib:/u01/app/oracle/product/19.0.0/db_1/rdbms/jlib

PATH=/u01/app/oracle/middleware/wlserver/server/bin:/u01/app/oracle/middleware/wlserver/../oracle_common/modules/thirdparty/org.apache.ant/1.10.5.0.0/apache-ant-1.10.5/bin:/usr/java/jdk1.8.0_241/jre/bin:/usr/java/jdk1.8.0_241/bin:/u01/app/oracle/product/19.0.0/db_1/bin:/bin:/home/oracle/.local/bin:/home/oracle/bin:/usr/local/bin:/usr/bin:/usr/local/sbin:/usr/sbin:/u01/app/oracle/middleware/wlserver/../oracle_common/modules/org.apache.maven_3.2.5/bin

Your environment has been set.
oracle@ORACLE-212$ java weblogic.version             

WebLogic Server 12.2.1.4.0 Thu Sep 12 04:04:29 GMT 2019 1974621

Use 'weblogic.version -verbose' to get subsystem information

Use 'weblogic.utils.Versions' to get version information for all modules
```

### Create Domain

```
$ $ORACLE_HOME/oracle_common/common/bin/config.sh
```
修改自启动密码
```
$ export DOMAIN_HOME=$ORACLE_BASE/config/domains/frsdomain
$ mkdir -p $DOMAIN_HOME/servers/AdminServer/security
$ echo "username=weblogic" > $DOMAIN_HOME/servers/AdminServer/security/boot.properties
$ echo "password=Password1" >> $DOMAIN_HOME/servers/AdminServer/security/boot.properties

$ $DOMAIN_HOME/startWebLogic.sh &
```