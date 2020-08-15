---
title:  "Rman image copy"
published: true
summary: "Rman image copy"
read_time: false
comments: false
related: false
author_profile: false
categories: 
  - dbascript
tags: 
  - rman
  - oracle
---

# 相关概念
## Block Change Tracking
 - 块跟踪
 - 主要用于RMAN备份的增量备份，记录自从上一次备份以来数据块的变化，相关后台进程CTWR（Change Tracking Writer）改善增量备份性能，RMAN可以不再扫描整个文件以查找变更数据。
 - 从10g开始提供此功能。


## backup set
 - 默认的备份类型，把数据文件中已经使用过的数据块备份到一个或多个文件中，这样的文件叫做“备份片”，所有备份出来的文件组合成为“备份集”。
 - 一个通道会产生一个备份集
 - 启动了控制文件自动备份，控制文件所在的备份文件会单独生成一个备份集，不会与数据文件备份集合并在一起。
 - 指定了每个备份集中包含的数据文件个数(filesperset)，即便只有一个通道，也有可能生成多个备份集 。

## backup piece
 - 每个备份片是一个单独的输出文件。
 - 一个备份片的大小是有限制的；如果没有大小的限制，备份集就只由一个备份片构成。
 - 备份片的大小不能大于你的文件系统所支持的文件的最大值。
```
# 设置备份片大小
RMAN > configure channel device type disk maxpiecesize 1024M ;
```
 - 控制文件备份以后，会出现一个独立备份集。控制文件和数据文件不能放在同一个备份集里，因为数据文件所在的备份集以Oracle 数据块为最小单位，而控制文件所在备份集是以操作系统块作为最小单位。
 - 归档日志文件所在的备份集也是以操作系统块为最小单位，所以归档日志文件备份集和数据文件备份集不能在同一个备份集里面。

## Image copy
 - 一个数据文件生成一个镜像副本文件(数据库数据文件、归档重做日志或者控制文件的精确副本)
 - 过程由RMAN完成，RMAN复制的时候也是一个数据块一个数据块(Oacle block)的复制
 - 默认检测数据块是否出现物理损坏(默认不会进行逻辑损坏检查,需要手工启动)，
 - 不需要将表空间置为begin backup状态
 - 和备份集类型不同在于生成的镜像副本中包含使用过的数据块，也包含从来没有用过的数据块 。
 - 恢复时速度相对备份集来说要更快 ，恢复时可以不用拷贝，指定新位置即可。
 - 映像备份需要占用和生产数据文件相同的空间，所以在数据量较大的情况下是不建议采用的
 - 命令backup as copy
## Image copies
 - 映像级别备份
 - 相当于数据文件和归档日志的拷贝复制品，与原文件在存储空间上完全一致
 - 如果需要做一个部分恢复（比如某一数据文件）采用映像备份情况下只需要检索相应的映像即可，恢复速度非常快，
  

### Image copy相对于复制的优点：
 - RMAN 能够验证备份文件内数据块的有效性
### Image copy与backup set区别
 - image copy 是datafile、control file、archived log完全一致的副本，可以不用RMAN直接恢复数据库。
 - backup set 格式为 RMAN 自有格式,由backup piece构成，可包含多个数据文件，压缩、增量备份，必须由RMAN恢复。


```
lsnrctl status     //查看Service name

sqlplus sys/oracle@cdb1 as sysdba    //登录cdb

col name format a16
col pdb format a16
select name,pdb from v$services;

alter session set container=pdb_easyee;
```

### 1、开启rman块追踪技术
```
SQL> alter database enable block change tracking using file '/u01/app/oracle/oradata/CDB1/block_change_track_file.f';

Database altered.
```
### 2、进行映像级别的level 0备份
```
$oracle@ORACLE-212$export ORACLE_SID=cdb1
$oracle@ORACLE-212$echo $ORACLE_SID
cdb1

$rman target /

Recovery Manager: Release 19.0.0.0.0 - Production on Wed Aug 5 03:50:00 2020
Version 19.3.0.0.0

Copyright (c) 1982, 2019, Oracle and/or its affiliates.  All rights reserved.

connected to target database: CDB1 (DBID=1010358776)

RMAN> report schema;

using target database control file instead of recovery catalog
Report of database schema for database with db_unique_name CDB1

List of Permanent Datafiles
===========================
File Size(MB) Tablespace           RB segs Datafile Name
---- -------- -------------------- ------- ------------------------
1    900      SYSTEM               YES     /u01/app/oracle/oradata/CDB1/system01.dbf
3    580      SYSAUX               NO      /u01/app/oracle/oradata/CDB1/sysaux01.dbf
4    65       UNDOTBS1             YES     /u01/app/oracle/oradata/CDB1/undotbs01.dbf
5    260      PDB$SEED:SYSTEM      NO      /u01/app/oracle/oradata/CDB1/pdbseed/system01.dbf
6    280      PDB$SEED:SYSAUX      NO      /u01/app/oracle/oradata/CDB1/pdbseed/sysaux01.dbf
7    5        USERS                NO      /u01/app/oracle/oradata/CDB1/users01.dbf
8    100      PDB$SEED:UNDOTBS1    NO      /u01/app/oracle/oradata/CDB1/pdbseed/undotbs01.dbf
9    540      PDB_EASYEE:SYSTEM    YES     /u01/app/oracle/oradata/CDB1/pdb_easyee/system01.dbf
10   320      PDB_EASYEE:SYSAUX    NO      /u01/app/oracle/oradata/CDB1/pdb_easyee/sysaux01.dbf
11   680      PDB_EASYEE:UNDOTBS1  YES     /u01/app/oracle/oradata/CDB1/pdb_easyee/undotbs01.dbf
12   5        PDB_EASYEE:USERS     NO      /u01/app/oracle/oradata/CDB1/pdb_easyee/users01.dbf
13   1024     PDB_EASYEE:PDBEASYEE_DATA NO      /u01/app/oracle/oradata/CDB1/pdb_easyee/pdbeasyee_data01.dbf
14   1024     PDB_EASYEE:BTTEST    NO      /u01/app/oracle/oradata/CDB1/pdb_easyee/bttest.dbf
15   10       PDB_EASYEE:UNDOTBS1  YES     /u01/app/oracle/oradata/CDB1/pdb_easyee/undotbs01_2.dbf

List of Temporary Files
=======================
File Size(MB) Tablespace           Maxsize(MB) Tempfile Name
---- -------- -------------------- ----------- --------------------
1    32       TEMP                 32767       /u01/app/oracle/oradata/CDB1/temp01.dbf
2    36       PDB$SEED:TEMP        32767       /u01/app/oracle/oradata/CDB1/pdbseed/temp012020-04-02_03-56-28-909-AM.dbf
3    36       PDB_EASYEE:TEMP      32767       /u01/app/oracle/oradata/CDB1/pdb_easyee/temp01.dbf
4    32       PDB_EASYEE:PDBEASYEE_TEMP 32767       /u01/app/oracle/oradata/CDB1/pdb_easyee/pdbeasyee_temp01.dbf

run {
ALLOCATE CHANNEL "ch1" DEVICE TYPE DISK FORMAT "/u01/bk/%U";
BACKUP AS COPY TAG 'BASE01' INCREMENTAL LEVEL 0 DATABASE;
}

RMAN> run {
2> ALLOCATE CHANNEL "ch1" DEVICE TYPE DISK FORMAT "/u01/bk/%U";
3> BACKUP AS COPY TAG 'BASE01' INCREMENTAL LEVEL 0 DATABASE;
4> }

using target database control file instead of recovery catalog
allocated channel: ch1
channel ch1: SID=304 device type=DISK

Starting backup at 2020-08-05 03:50:56
channel ch1: starting datafile copy
input datafile file number=00013 name=/u01/app/oracle/oradata/CDB1/pdb_easyee/pdbeasyee_data01.dbf
output file name=/u01/bk/data_D-CDB1_I-1010358776_TS-PDBEASYEE_DATA_FNO-13_02v72kd0 tag=BASE01 RECID=4 STAMP=1047613881
channel ch1: datafile copy complete, elapsed time: 00:00:25
channel ch1: starting datafile copy
input datafile file number=00014 name=/u01/app/oracle/oradata/CDB1/pdb_easyee/bttest.dbf
output file name=/u01/bk/data_D-CDB1_I-1010358776_TS-BTTEST_FNO-14_03v72kdp tag=BASE01 RECID=5 STAMP=1047613908
channel ch1: datafile copy complete, elapsed time: 00:00:35
channel ch1: starting datafile copy
input datafile file number=00001 name=/u01/app/oracle/oradata/CDB1/system01.dbf
output file name=/u01/bk/data_D-CDB1_I-1010358776_TS-SYSTEM_FNO-1_04v72kes tag=BASE01 RECID=6 STAMP=1047613954
channel ch1: datafile copy complete, elapsed time: 00:00:45
channel ch1: starting datafile copy
input datafile file number=00011 name=/u01/app/oracle/oradata/CDB1/pdb_easyee/undotbs01.dbf
output file name=/u01/bk/data_D-CDB1_I-1010358776_TS-UNDOTBS1_FNO-11_05v72kg9 tag=BASE01 RECID=7 STAMP=1047613993
channel ch1: datafile copy complete, elapsed time: 00:00:35
channel ch1: starting datafile copy
input datafile file number=00003 name=/u01/app/oracle/oradata/CDB1/sysaux01.dbf
output file name=/u01/bk/data_D-CDB1_I-1010358776_TS-SYSAUX_FNO-3_06v72khd tag=BASE01 RECID=8 STAMP=1047614019
channel ch1: datafile copy complete, elapsed time: 00:00:25
channel ch1: starting datafile copy
input datafile file number=00009 name=/u01/app/oracle/oradata/CDB1/pdb_easyee/system01.dbf
output file name=/u01/bk/data_D-CDB1_I-1010358776_TS-SYSTEM_FNO-9_07v72ki6 tag=BASE01 RECID=9 STAMP=1047614044
channel ch1: datafile copy complete, elapsed time: 00:00:25
channel ch1: starting datafile copy
input datafile file number=00010 name=/u01/app/oracle/oradata/CDB1/pdb_easyee/sysaux01.dbf
output file name=/u01/bk/data_D-CDB1_I-1010358776_TS-SYSAUX_FNO-10_08v72kiv tag=BASE01 RECID=10 STAMP=1047614054
channel ch1: datafile copy complete, elapsed time: 00:00:15
channel ch1: starting datafile copy
input datafile file number=00006 name=/u01/app/oracle/oradata/CDB1/pdbseed/sysaux01.dbf
output file name=/u01/bk/data_D-CDB1_I-1010358776_TS-SYSAUX_FNO-6_09v72kje tag=BASE01 RECID=11 STAMP=1047614068
channel ch1: datafile copy complete, elapsed time: 00:00:07
channel ch1: starting datafile copy
input datafile file number=00005 name=/u01/app/oracle/oradata/CDB1/pdbseed/system01.dbf
output file name=/u01/bk/data_D-CDB1_I-1010358776_TS-SYSTEM_FNO-5_0av72kjl tag=BASE01 RECID=12 STAMP=1047614075
channel ch1: datafile copy complete, elapsed time: 00:00:07
channel ch1: starting datafile copy
input datafile file number=00008 name=/u01/app/oracle/oradata/CDB1/pdbseed/undotbs01.dbf
output file name=/u01/bk/data_D-CDB1_I-1010358776_TS-UNDOTBS1_FNO-8_0bv72kjs tag=BASE01 RECID=13 STAMP=1047614078
channel ch1: datafile copy complete, elapsed time: 00:00:03
channel ch1: starting datafile copy
input datafile file number=00004 name=/u01/app/oracle/oradata/CDB1/undotbs01.dbf
output file name=/u01/bk/data_D-CDB1_I-1010358776_TS-UNDOTBS1_FNO-4_0cv72kjv tag=BASE01 RECID=14 STAMP=1047614081
channel ch1: datafile copy complete, elapsed time: 00:00:03
channel ch1: starting datafile copy
input datafile file number=00015 name=/u01/app/oracle/oradata/CDB1/pdb_easyee/undotbs01_2.dbf
output file name=/u01/bk/data_D-CDB1_I-1010358776_TS-UNDOTBS1_FNO-15_0dv72kk2 tag=BASE01 RECID=15 STAMP=1047614082
channel ch1: datafile copy complete, elapsed time: 00:00:01
channel ch1: starting datafile copy
input datafile file number=00007 name=/u01/app/oracle/oradata/CDB1/users01.dbf
output file name=/u01/bk/data_D-CDB1_I-1010358776_TS-USERS_FNO-7_0ev72kk3 tag=BASE01 RECID=16 STAMP=1047614083
channel ch1: datafile copy complete, elapsed time: 00:00:01
channel ch1: starting datafile copy
input datafile file number=00012 name=/u01/app/oracle/oradata/CDB1/pdb_easyee/users01.dbf
output file name=/u01/bk/data_D-CDB1_I-1010358776_TS-USERS_FNO-12_0fv72kk4 tag=BASE01 RECID=17 STAMP=1047614084
channel ch1: datafile copy complete, elapsed time: 00:00:01
Finished backup at 2020-08-05 03:54:45

Starting Control File and SPFILE Autobackup at 2020-08-05 03:54:45
piece handle=/u01/app/oracle/flashback/CDB1/autobackup/2020_08_05/o1_mf_s_1047614085_hlnsg6g5_.bkp comment=NONE
Finished Control File and SPFILE Autobackup at 2020-08-05 03:54:46
released channel: ch1

oracle@ORACLE-212$ls -lh /u01/bk
total 5.7G
-rw-r----- 1 oracle oinstall 1.1G Aug  5 03:51 data_D-CDB1_I-1010358776_TS-BTTEST_FNO-14_03v72kdp
-rw-r----- 1 oracle oinstall 1.1G Aug  5 03:51 data_D-CDB1_I-1010358776_TS-PDBEASYEE_DATA_FNO-13_02v72kd0
-rw-r----- 1 oracle oinstall 321M Aug  5 03:54 data_D-CDB1_I-1010358776_TS-SYSAUX_FNO-10_08v72kiv
-rw-r----- 1 oracle oinstall 581M Aug  5 03:53 data_D-CDB1_I-1010358776_TS-SYSAUX_FNO-3_06v72khd
-rw-r----- 1 oracle oinstall 281M Aug  5 03:54 data_D-CDB1_I-1010358776_TS-SYSAUX_FNO-6_09v72kje
-rw-r----- 1 oracle oinstall 901M Aug  5 03:52 data_D-CDB1_I-1010358776_TS-SYSTEM_FNO-1_04v72kes
-rw-r----- 1 oracle oinstall 261M Aug  5 03:54 data_D-CDB1_I-1010358776_TS-SYSTEM_FNO-5_0av72kjl
-rw-r----- 1 oracle oinstall 541M Aug  5 03:53 data_D-CDB1_I-1010358776_TS-SYSTEM_FNO-9_07v72ki6
-rw-r----- 1 oracle oinstall 681M Aug  5 03:52 data_D-CDB1_I-1010358776_TS-UNDOTBS1_FNO-11_05v72kg9
-rw-r----- 1 oracle oinstall  11M Aug  5 03:54 data_D-CDB1_I-1010358776_TS-UNDOTBS1_FNO-15_0dv72kk2
-rw-r----- 1 oracle oinstall  66M Aug  5 03:54 data_D-CDB1_I-1010358776_TS-UNDOTBS1_FNO-4_0cv72kjv
-rw-r----- 1 oracle oinstall 101M Aug  5 03:54 data_D-CDB1_I-1010358776_TS-UNDOTBS1_FNO-8_0bv72kjs
-rw-r----- 1 oracle oinstall 5.1M Aug  5 03:54 data_D-CDB1_I-1010358776_TS-USERS_FNO-12_0fv72kk4
-rw-r----- 1 oracle oinstall 5.1M Aug  5 03:54 data_D-CDB1_I-1010358776_TS-USERS_FNO-7_0ev72kk3
```

### 3、进行映像级别的level 1备份
```
创建测试数据
SQL> create table rmantest1 as select * from rmantest;
SQL> select count(*) from rmantest1;


run {
ALLOCATE CHANNEL "ch1" DEVICE TYPE DISK FORMAT "/u01/bk/%U";
BACKUP TAG 'incr_update' INCREMENTAL LEVEL 1 DATABASE PLUS ARCHIVELOG;
}

RMAN> run {
2> ALLOCATE CHANNEL "ch1" DEVICE TYPE DISK FORMAT "/u01/bk/%U";
3> BACKUP TAG 'incr_update' INCREMENTAL LEVEL 1 DATABASE;
4> }

allocated channel: ch1
channel ch1: SID=304 device type=DISK

Starting backup at 2020-08-05 04:00:21
channel ch1: starting incremental level 1 datafile backup set
channel ch1: specifying datafile(s) in backup set
input datafile file number=00013 name=/u01/app/oracle/oradata/CDB1/pdb_easyee/pdbeasyee_data01.dbf
input datafile file number=00014 name=/u01/app/oracle/oradata/CDB1/pdb_easyee/bttest.dbf
input datafile file number=00011 name=/u01/app/oracle/oradata/CDB1/pdb_easyee/undotbs01.dbf
input datafile file number=00009 name=/u01/app/oracle/oradata/CDB1/pdb_easyee/system01.dbf
input datafile file number=00010 name=/u01/app/oracle/oradata/CDB1/pdb_easyee/sysaux01.dbf
input datafile file number=00015 name=/u01/app/oracle/oradata/CDB1/pdb_easyee/undotbs01_2.dbf
input datafile file number=00012 name=/u01/app/oracle/oradata/CDB1/pdb_easyee/users01.dbf
channel ch1: starting piece 1 at 2020-08-05 04:00:21
channel ch1: finished piece 1 at 2020-08-05 04:00:22
piece handle=/u01/bk/0hv72kul_1_1 tag=INCR_UPDATE comment=NONE
channel ch1: backup set complete, elapsed time: 00:00:01
channel ch1: starting incremental level 1 datafile backup set
channel ch1: specifying datafile(s) in backup set
input datafile file number=00001 name=/u01/app/oracle/oradata/CDB1/system01.dbf
input datafile file number=00003 name=/u01/app/oracle/oradata/CDB1/sysaux01.dbf
input datafile file number=00004 name=/u01/app/oracle/oradata/CDB1/undotbs01.dbf
input datafile file number=00007 name=/u01/app/oracle/oradata/CDB1/users01.dbf
channel ch1: starting piece 1 at 2020-08-05 04:00:23
channel ch1: finished piece 1 at 2020-08-05 04:00:26
piece handle=/u01/bk/0iv72kum_1_1 tag=INCR_UPDATE comment=NONE
channel ch1: backup set complete, elapsed time: 00:00:03
channel ch1: starting incremental level 1 datafile backup set
channel ch1: specifying datafile(s) in backup set
input datafile file number=00006 name=/u01/app/oracle/oradata/CDB1/pdbseed/sysaux01.dbf
skipping datafile 00006 because it has not changed
input datafile file number=00005 name=/u01/app/oracle/oradata/CDB1/pdbseed/system01.dbf
skipping datafile 00005 because it has not changed
input datafile file number=00008 name=/u01/app/oracle/oradata/CDB1/pdbseed/undotbs01.dbf
skipping datafile 00008 because it has not changed
channel ch1: backup cancelled because all files were skipped
Finished backup at 2020-08-05 04:00:26

Starting Control File and SPFILE Autobackup at 2020-08-05 04:00:26
piece handle=/u01/app/oracle/flashback/CDB1/autobackup/2020_08_05/o1_mf_s_1047614426_hlnsrtdc_.bkp comment=NONE
Finished Control File and SPFILE Autobackup at 2020-08-05 04:00:27
released channel: ch1

$ls -lh /u01/bk
total 5.7G
-rw-r----- 1 oracle oinstall 248K Aug  5 04:00 0hv72kul_1_1
-rw-r----- 1 oracle oinstall 9.5M Aug  5 04:00 0iv72kum_1_1
-rw-r----- 1 oracle oinstall 1.1G Aug  5 03:51 data_D-CDB1_I-1010358776_TS-BTTEST_FNO-14_03v72kdp
-rw-r----- 1 oracle oinstall 1.1G Aug  5 03:51 data_D-CDB1_I-1010358776_TS-PDBEASYEE_DATA_FNO-13_02v72kd0
-rw-r----- 1 oracle oinstall 321M Aug  5 03:54 data_D-CDB1_I-1010358776_TS-SYSAUX_FNO-10_08v72kiv
-rw-r----- 1 oracle oinstall 581M Aug  5 03:53 data_D-CDB1_I-1010358776_TS-SYSAUX_FNO-3_06v72khd
-rw-r----- 1 oracle oinstall 281M Aug  5 03:54 data_D-CDB1_I-1010358776_TS-SYSAUX_FNO-6_09v72kje
-rw-r----- 1 oracle oinstall 901M Aug  5 03:52 data_D-CDB1_I-1010358776_TS-SYSTEM_FNO-1_04v72kes
-rw-r----- 1 oracle oinstall 261M Aug  5 03:54 data_D-CDB1_I-1010358776_TS-SYSTEM_FNO-5_0av72kjl
-rw-r----- 1 oracle oinstall 541M Aug  5 03:53 data_D-CDB1_I-1010358776_TS-SYSTEM_FNO-9_07v72ki6
-rw-r----- 1 oracle oinstall 681M Aug  5 03:52 data_D-CDB1_I-1010358776_TS-UNDOTBS1_FNO-11_05v72kg9
-rw-r----- 1 oracle oinstall  11M Aug  5 03:54 data_D-CDB1_I-1010358776_TS-UNDOTBS1_FNO-15_0dv72kk2
-rw-r----- 1 oracle oinstall  66M Aug  5 03:54 data_D-CDB1_I-1010358776_TS-UNDOTBS1_FNO-4_0cv72kjv
-rw-r----- 1 oracle oinstall 101M Aug  5 03:54 data_D-CDB1_I-1010358776_TS-UNDOTBS1_FNO-8_0bv72kjs
-rw-r----- 1 oracle oinstall 5.1M Aug  5 03:54 data_D-CDB1_I-1010358776_TS-USERS_FNO-12_0fv72kk4
-rw-r----- 1 oracle oinstall 5.1M Aug  5 03:54 data_D-CDB1_I-1010358776_TS-USERS_FNO-7_0ev72kk3
```
### 4、用level 1的备份update level 0 中的映像
```
run {
ALLOCATE CHANNEL "ch1" DEVICE TYPE DISK FORMAT "/u01/bk/%U";
RECOVER COPY OF DATABASE WITH TAG 'BASE01'; 
 }

RMAN> run {
2> ALLOCATE CHANNEL "ch1" DEVICE TYPE DISK FORMAT "/u01/bk/%U";
3> RECOVER COPY OF DATABASE WITH TAG 'BASE01'; 
4>  }

allocated channel: ch1
channel ch1: SID=304 device type=DISK

Starting recover at 2020-08-05 04:21:52
no copy of datafile 5 found to recover
no copy of datafile 6 found to recover
no copy of datafile 8 found to recover
channel ch1: starting incremental datafile backup set restore
channel ch1: specifying datafile copies to recover
recovering datafile copy file number=00009 name=/u01/bk/data_D-CDB1_I-1010358776_TS-SYSTEM_FNO-9_07v72ki6
recovering datafile copy file number=00010 name=/u01/bk/data_D-CDB1_I-1010358776_TS-SYSAUX_FNO-10_08v72kiv
recovering datafile copy file number=00011 name=/u01/bk/data_D-CDB1_I-1010358776_TS-UNDOTBS1_FNO-11_05v72kg9
recovering datafile copy file number=00012 name=/u01/bk/data_D-CDB1_I-1010358776_TS-USERS_FNO-12_0fv72kk4
recovering datafile copy file number=00013 name=/u01/bk/data_D-CDB1_I-1010358776_TS-PDBEASYEE_DATA_FNO-13_02v72kd0
recovering datafile copy file number=00014 name=/u01/bk/data_D-CDB1_I-1010358776_TS-BTTEST_FNO-14_03v72kdp
recovering datafile copy file number=00015 name=/u01/bk/data_D-CDB1_I-1010358776_TS-UNDOTBS1_FNO-15_0dv72kk2
channel ch1: reading from backup piece /u01/bk/0hv72kul_1_1
channel ch1: piece handle=/u01/bk/0hv72kul_1_1 tag=INCR_UPDATE
channel ch1: restored backup piece 1
channel ch1: restore complete, elapsed time: 00:00:01
channel ch1: starting incremental datafile backup set restore
channel ch1: specifying datafile copies to recover
recovering datafile copy file number=00001 name=/u01/bk/data_D-CDB1_I-1010358776_TS-SYSTEM_FNO-1_04v72kes
recovering datafile copy file number=00003 name=/u01/bk/data_D-CDB1_I-1010358776_TS-SYSAUX_FNO-3_06v72khd
recovering datafile copy file number=00004 name=/u01/bk/data_D-CDB1_I-1010358776_TS-UNDOTBS1_FNO-4_0cv72kjv
recovering datafile copy file number=00007 name=/u01/bk/data_D-CDB1_I-1010358776_TS-USERS_FNO-7_0ev72kk3
channel ch1: reading from backup piece /u01/bk/0iv72kum_1_1
channel ch1: piece handle=/u01/bk/0iv72kum_1_1 tag=INCR_UPDATE
channel ch1: restored backup piece 1
channel ch1: restore complete, elapsed time: 00:00:03
Finished recover at 2020-08-05 04:21:57

Starting Control File and SPFILE Autobackup at 2020-08-05 04:21:57
piece handle=/u01/app/oracle/flashback/CDB1/autobackup/2020_08_05/o1_mf_s_1047615717_hlnv15jx_.bkp comment=NONE
Finished Control File and SPFILE Autobackup at 2020-08-05 04:21:58
released channel: ch1

oracle@ORACLE-212$ls -lh /u01/bk
total 5.7G
-rw-r----- 1 oracle oinstall 248K Aug  5 04:00 0hv72kul_1_1
-rw-r----- 1 oracle oinstall 9.5M Aug  5 04:00 0iv72kum_1_1
-rw-r----- 1 oracle oinstall 1.1G Aug  5 04:21 data_D-CDB1_I-1010358776_TS-BTTEST_FNO-14_03v72kdp
-rw-r----- 1 oracle oinstall 1.1G Aug  5 04:21 data_D-CDB1_I-1010358776_TS-PDBEASYEE_DATA_FNO-13_02v72kd0
-rw-r----- 1 oracle oinstall 321M Aug  5 04:21 data_D-CDB1_I-1010358776_TS-SYSAUX_FNO-10_08v72kiv
-rw-r----- 1 oracle oinstall 581M Aug  5 04:21 data_D-CDB1_I-1010358776_TS-SYSAUX_FNO-3_06v72khd
-rw-r----- 1 oracle oinstall 281M Aug  5 03:54 data_D-CDB1_I-1010358776_TS-SYSAUX_FNO-6_09v72kje
-rw-r----- 1 oracle oinstall 901M Aug  5 04:21 data_D-CDB1_I-1010358776_TS-SYSTEM_FNO-1_04v72kes
-rw-r----- 1 oracle oinstall 261M Aug  5 03:54 data_D-CDB1_I-1010358776_TS-SYSTEM_FNO-5_0av72kjl
-rw-r----- 1 oracle oinstall 541M Aug  5 04:21 data_D-CDB1_I-1010358776_TS-SYSTEM_FNO-9_07v72ki6
-rw-r----- 1 oracle oinstall 681M Aug  5 04:21 data_D-CDB1_I-1010358776_TS-UNDOTBS1_FNO-11_05v72kg9
-rw-r----- 1 oracle oinstall  11M Aug  5 04:21 data_D-CDB1_I-1010358776_TS-UNDOTBS1_FNO-15_0dv72kk2
-rw-r----- 1 oracle oinstall  66M Aug  5 04:21 data_D-CDB1_I-1010358776_TS-UNDOTBS1_FNO-4_0cv72kjv
-rw-r----- 1 oracle oinstall 101M Aug  5 03:54 data_D-CDB1_I-1010358776_TS-UNDOTBS1_FNO-8_0bv72kjs
-rw-r----- 1 oracle oinstall 5.1M Aug  5 04:21 data_D-CDB1_I-1010358776_TS-USERS_FNO-12_0fv72kk4
-rw-r----- 1 oracle oinstall 5.1M Aug  5 04:21 data_D-CDB1_I-1010358776_TS-USERS_FNO-7_0ev72kk3

$ls -lh /u01/app/oracle/flashback/CDB1/autobackup/
total 0
drwxr-x--- 2 oracle oinstall  46 Jun 23 01:47 2020_06_23
drwxr-x--- 2 oracle oinstall 126 Aug  5 04:21 2020_08_05
```
### 5、获取源库控制文件
```
SQL> alter database backup controlfile to trace;

2020-08-05T04:34:55.312846-04:00
alter database backup controlfile to trace
2020-08-05T04:34:55.314064-04:00
Backup controlfile written to trace file /u01/app/oracle/diag/rdbms/cdb1/cdb1/trace/cdb1_ora_3319.trc
Completed: alter database backup controlfile to trace
```

### 6. 创建新的对象，用于后续完全恢复时进行稽核
```
SQL> create table scott.test22 (id number);
SQL> alter system switch logfile;
SQL> shutdown immediate;
```

### 7. 创建 IMAGE COPY 数据库所需的参数文件(参数文件中的db_name不能修改，audit_file_dest与control_files路径要进行调整)
```
[oracle@sqlaudit ~]$ cd $ORACLE_HOME/dbs
[oracle@sqlaudit dbs]$ cp orapwsrcdb orapwsrcdbnew
[oracle@sqlaudit dbs]$ strings spfilesrcdb.ora > initsrcdbnew.ora
[oracle@sqlaudit dbs]$ cat initsrcdbnew.ora | grep '/oracle/'
*.audit_file_dest='/oracle/app/oracle/admin/srcdbnew/adump'
*.control_files='/oradata/srcdb_img/control01.ctl','/oradata/srcdb_img/control02.ctl'
[oracle@sqlaudit dbs]$ mkdir -p /oracle/app/oracle/admin/srcdbnew/adump
```
### 8. 启动数据库实例
```
[oracle@sqlaudit dbs]$ export ORACLE_SID=srcdbnew
[oracle@sqlaudit dbs]$ sqlplus / as sysdba
SQL> startup nomount;
```

### 9. 复制在线日志文件到新目录用于完全恢复
```
[oracle@sqlaudit archive]$ cp /oracle/app/oracle/oradata/srcdb/redo01.log /oradata/srcdb_img/redo01.log
[oracle@sqlaudit archive]$ cp /oracle/app/oracle/oradata/srcdb/redo02.log /oradata/srcdb_img/redo02.log
[oracle@sqlaudit archive]$ cp /oracle/app/oracle/oradata/srcdb/redo03.log /oradata/srcdb_img/redo03.log
```

### 10. 重建控制文件
```
CREATE CONTROLFILE REUSE DATABASE "SRCDB" NORESETLOGS
    MAXLOGFILES 16
    MAXLOGMEMBERS 3
    MAXDATAFILES 100
    MAXINSTANCES 8
    MAXLOGHISTORY 292
LOGFILE
  GROUP 1 '/oradata/srcdb_img/redo01.log'  SIZE 50M BLOCKSIZE 512,
  GROUP 2 '/oradata/srcdb_img/redo02.log'  SIZE 50M BLOCKSIZE 512,
  GROUP 3 '/oradata/srcdb_img/redo03.log'  SIZE 50M BLOCKSIZE 512
DATAFILE
  '/oradata/srcdb_img/SRCDB-LVL0-data_D-SRCDB_I-595837900_TS-SYSTEM_FNO-1_0nsq1niv',
  '/oradata/srcdb_img/SRCDB-LVL0-data_D-SRCDB_I-595837900_TS-SYSAUX_FNO-2_0osq1nj2',
  '/oradata/srcdb_img/SRCDB-LVL0-data_D-SRCDB_I-595837900_TS-UNDOTBS1_FNO-3_0qsq1njk',
  '/oradata/srcdb_img/SRCDB-LVL0-data_D-SRCDB_I-595837900_TS-USERS_FNO-4_0rsq1njl',
  '/oradata/srcdb_img/SRCDB-LVL0-data_D-SRCDB_I-595837900_TS-GGTBS_FNO-5_0psq1njh'
CHARACTER SET WE8MSWIN1252
;
SQL> SELECT NAME FROM V$DATAFILE;
```

### 11. 源库检查检查需要注册的日志文件
```
[oracle@sqlaudit dbs]$ export ORACLE_SID=srcdb
SQL> startup mount;
RMAN> list backup of archivelog all;
BS Key  Size      Device Type Elapsed Time Completion Time    
------- ---------- ----------- ------------ -------------------
19      3.50K      DISK        00:00:00    2018-01-31 07:48:46
        BP Key: 19  Status: AVAILABLE  Compressed: NO  Tag: SRCDB-IMAGE
        Piece Name: /oracle/app/oracle/product/11.2.0.4/db_1/dbs/15sq1nmu_1_1
  List of Archived Logs in backup set 19
  Thrd Seq    Low SCN    Low Time            Next SCN  Next Time
  ---- ------- ---------- ------------------- ---------- ---------
  1    11      969636    2018-01-31 07:48:41 969648    2018-01-31 07:48:46
```

### 12. 新库注册源库的日志文件
```
SQL> ALTER DATABASE REGISTER LOGFILE '/oracle/archive/1_11_961988430.dbf';
SQL> ALTER DATABASE REGISTER LOGFILE '/oracle/archive/1_12_961988430.dbf';
SQL> ALTER DATABASE REGISTER LOGFILE '/oracle/archive/1_13_961988430.dbf';
SQL> ALTER DATABASE REGISTER LOGFILE '/oracle/archive/1_14_961988430.dbf';
SQL> RECOVER DATABASE;
SQL> ALTER DATABASE OPEN;
```

### 13. 创建新的临时文件
```
SQL> ALTER TABLESPACE TEMP ADD TEMPFILE '/oradata/srcdb_img/temp01.dbf' SIZE 100M AUTOEXTEND OFF;
```