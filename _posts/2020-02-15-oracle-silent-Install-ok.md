---
title: "Oracle数据库静默安装"
published: true
categories: oracle
permalink: oracle-silent-install.html
summary: "Oracle数据库静默安装"
tags: [oracle]
---
<!-- 
**翻译转载** [https://oracle-base.com/articles/misc/oui-silent-installations](https://oracle-base.com/articles/misc/oui-silent-installations)
-->

ORACLE 安装介质在response目录包含一个名为db_install.rsp的响应文件，可以编辑用于静默安装。也可以通过OUI工具生成一个自定义响应文件。

使用“- record”参数告诉安装程序生成响应文件

```
./runInstaller -record -destinationFile /tmp/10gR2.rsp
```

而从11gR2开始，不再支持“- record”选项，会在OUI最后一个画面提示保存响应文件

![response](/assets/images/2020-02-15-save-response-file.jpg)

不同版本响应文件的详细注释
- [10gR2.rsp](https://oracle-base.com/articles/misc/10gR2.rsp)
- [11gR2.rsp](https://oracle-base.com/articles/misc/11gR2.rsp)
- [12cR1.rsp](https://oracle-base.com/articles/misc/12cR1.rsp)
- [12cR2.rsp](https://oracle-base.com/articles/misc/12cR2.rsp)

**安装命令：**

```
# Normal.
./runInstaller -silent -responseFile /tmp/12cR2.rsp

# Ignore Prerequisites.
./runInstaller -ignoreSysPrereqs -ignorePrereq -waitforcompletion -showProgress -silent -responseFile /tmp/12cR2.rsp
```

**参数：**

```
-silent            : Run in silent mode.
-responsefile      : Specified the location of the response file. 
-ignoreSysPrereqs  : Ignore the system prerequisite checks.
-ignorePrereq      : Ignore the general prerequisite checks.
-waitforcompletion : Stop the installer spawning as a separate process, so scripts happen in sequence.
-invPtrLoc         : Used to specify the location of the oraInst.loc file, which in turn specifies the inventory details.
-force             : Installation continues when a non-empty directory is used for the ORACLE_HOME.
-showProgress      : Displays line of "." to show something is happeing.
```

**通过默认响应文件指定参数安装**

```
# 12cR1 and below.

./runInstaller -ignoreSysPrereqs -ignorePrereq -waitforcompletion -showProgress -silent \
    -responseFile /tmp/database/response/db_install.rsp \
    oracle.install.option=INSTALL_DB_SWONLY \
    ORACLE_HOSTNAME=${ORACLE_HOSTNAME} \
    UNIX_GROUP_NAME=oinstall \
    INVENTORY_LOCATION=${ORA_INVENTORY} \
    SELECTED_LANGUAGES=en,en_GB \
    ORACLE_HOME=${ORACLE_HOME} \
    ORACLE_BASE=${ORACLE_BASE} \
    oracle.install.db.InstallEdition=EE \
    oracle.install.db.DBA_GROUP=dba \
    oracle.install.db.OPER_GROUP=dba \
    oracle.install.db.BACKUPDBA_GROUP=dba \
    oracle.install.db.DGDBA_GROUP=dba \
    oracle.install.db.KMDBA_GROUP=dba \
    SECURITY_UPDATES_VIA_MYORACLESUPPORT=false \
    DECLINE_SECURITY_UPDATES=true

# 12cR2.

./runInstaller -ignoreSysPrereqs -ignorePrereq -waitforcompletion -showProgress -silent \
    -responseFile /tmp/database/response/db_install.rsp \
    oracle.install.option=INSTALL_DB_SWONLY \
    ORACLE_HOSTNAME=${ORACLE_HOSTNAME} \
    UNIX_GROUP_NAME=oinstall \
    INVENTORY_LOCATION=${ORA_INVENTORY} \
    SELECTED_LANGUAGES=en,en_GB \
    ORACLE_HOME=${ORACLE_HOME} \
    ORACLE_BASE=${ORACLE_BASE} \
    oracle.install.db.InstallEdition=EE \
    oracle.install.db.OSDBA_GROUP=dba \
    oracle.install.db.OSBACKUPDBA_GROUP=dba \
    oracle.install.db.OSDGDBA_GROUP=dba \
    oracle.install.db.OSKMDBA_GROUP=dba \
    oracle.install.db.OSRACDBA_GROUP=dba \
    SECURITY_UPDATES_VIA_MYORACLESUPPORT=false \
    DECLINE_SECURITY_UPDATES=true    

# 18c.

cd $ORACLE_HOME
unzip -oq /path/to/software/LINUX.X64_180000_db_home.zip

./runInstaller -ignorePrereq -waitforcompletion -silent \
    -responseFile ${ORACLE_HOME}/install/response/db_install.rsp \
    oracle.install.option=INSTALL_DB_SWONLY \
    ORACLE_HOSTNAME=${ORACLE_HOSTNAME} \
    UNIX_GROUP_NAME=oinstall \
    INVENTORY_LOCATION=${ORA_INVENTORY} \
    SELECTED_LANGUAGES=en,en_GB \
    ORACLE_HOME=${ORACLE_HOME} \
    ORACLE_BASE=${ORACLE_BASE} \
    oracle.install.db.InstallEdition=EE \
    oracle.install.db.OSDBA_GROUP=dba \
    oracle.install.db.OSBACKUPDBA_GROUP=dba \
    oracle.install.db.OSDGDBA_GROUP=dba \
    oracle.install.db.OSKMDBA_GROUP=dba \
    oracle.install.db.OSRACDBA_GROUP=dba \
    SECURITY_UPDATES_VIA_MYORACLESUPPORT=false \
    DECLINE_SECURITY_UPDATES=true    
```