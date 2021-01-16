---
title:  "oracle control recover"
published: true
summary: "通过历史控制文件恢复Oracle数据库"
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

```
sqlplus sys/oracle@cdb1 as sysdba    //登录cdb

SQL> archive log list;
Database log mode              Archive Mode
Automatic archival             Enabled
Archive destination            USE_DB_RECOVERY_FILE_DEST
Oldest online log sequence     23
Next log sequence to archive   25
Current log sequence           25
```
rman target sys/oracle@cdb1 

RMAN> show all;

using target database control file instead of recovery catalog
RMAN configuration parameters for database with db_unique_name CDB1 are:
CONFIGURE RETENTION POLICY TO REDUNDANCY 1; # default
CONFIGURE BACKUP OPTIMIZATION OFF; # default
CONFIGURE DEFAULT DEVICE TYPE TO DISK; # default
CONFIGURE CONTROLFILE AUTOBACKUP ON; # default   控制文件自动备份
CONFIGURE CONTROLFILE AUTOBACKUP FORMAT FOR DEVICE TYPE DISK TO '%F'; # default
CONFIGURE DEVICE TYPE DISK PARALLELISM 1 BACKUP TYPE TO BACKUPSET; # default
CONFIGURE DATAFILE BACKUP COPIES FOR DEVICE TYPE DISK TO 1; # default
CONFIGURE ARCHIVELOG BACKUP COPIES FOR DEVICE TYPE DISK TO 1; # default
CONFIGURE MAXSETSIZE TO UNLIMITED; # default
CONFIGURE ENCRYPTION FOR DATABASE OFF; # default
CONFIGURE ENCRYPTION ALGORITHM 'AES128'; # default
CONFIGURE COMPRESSION ALGORITHM 'BASIC' AS OF RELEASE 'DEFAULT' OPTIMIZE FOR LOAD TRUE ; # default
CONFIGURE RMAN OUTPUT TO KEEP FOR 7 DAYS; # default
CONFIGURE ARCHIVELOG DELETION POLICY TO NONE; # default
CONFIGURE SNAPSHOT CONTROLFILE NAME TO '/u01/app/oracle/product/19.0.0/db_1/dbs/snapcf_cdb1.f'; # default

RMAN> run
2> {allocate channel c1 type disk;
3> sql 'alter system archive log current';
4> backup database format '/backup/full_%d_%T_%s_%p';
5> backup current controlfile format '/backup/ctl_%d_%T_%s_%p';
6> release channel c1;
7> }

allocated channel: c1
channel c1: SID=291 device type=DISK

sql statement: alter system archive log current

Starting backup at 2020-11-04 07:06:21
channel c1: starting full datafile backup set
channel c1: specifying datafile(s) in backup set
input datafile file number=00013 name=/u01/app/oracle/oradata/CDB1/pdb_easyee/pdbeasyee_data01.dbf
input datafile file number=00014 name=/u01/app/oracle/oradata/CDB1/pdb_easyee/bttest.dbf
input datafile file number=00011 name=/u01/app/oracle/oradata/CDB1/pdb_easyee/undotbs01.dbf
input datafile file number=00009 name=/u01/app/oracle/oradata/CDB1/pdb_easyee/system01.dbf
input datafile file number=00010 name=/u01/app/oracle/oradata/CDB1/pdb_easyee/sysaux01.dbf
input datafile file number=00015 name=/u01/app/oracle/oradata/CDB1/pdb_easyee/undotbs01_2.dbf
input datafile file number=00012 name=/u01/app/oracle/oradata/CDB1/pdb_easyee/users01.dbf
channel c1: starting piece 1 at 2020-11-04 07:06:21
channel c1: finished piece 1 at 2020-11-04 07:07:16
piece handle=/backup/full_CDB1_20201104_22_1 tag=TAG20201104T070621 comment=NONE
channel c1: backup set complete, elapsed time: 00:00:55
channel c1: starting full datafile backup set
channel c1: specifying datafile(s) in backup set
input datafile file number=00001 name=/u01/app/oracle/oradata/CDB1/system01.dbf
input datafile file number=00003 name=/u01/app/oracle/oradata/CDB1/sysaux01.dbf
input datafile file number=00004 name=/u01/app/oracle/oradata/CDB1/undotbs01.dbf
input datafile file number=00007 name=/u01/app/oracle/oradata/CDB1/users01.dbf
channel c1: starting piece 1 at 2020-11-04 07:07:16
channel c1: finished piece 1 at 2020-11-04 07:08:21
piece handle=/backup/full_CDB1_20201104_23_1 tag=TAG20201104T070621 comment=NONE
channel c1: backup set complete, elapsed time: 00:01:05
channel c1: starting full datafile backup set
channel c1: specifying datafile(s) in backup set
input datafile file number=00006 name=/u01/app/oracle/oradata/CDB1/pdbseed/sysaux01.dbf
input datafile file number=00005 name=/u01/app/oracle/oradata/CDB1/pdbseed/system01.dbf
input datafile file number=00008 name=/u01/app/oracle/oradata/CDB1/pdbseed/undotbs01.dbf
channel c1: starting piece 1 at 2020-11-04 07:08:21
channel c1: finished piece 1 at 2020-11-04 07:08:46
piece handle=/backup/full_CDB1_20201104_24_1 tag=TAG20201104T070621 comment=NONE
channel c1: backup set complete, elapsed time: 00:00:25
Finished backup at 2020-11-04 07:08:46

Starting backup at 2020-11-04 07:08:47
channel c1: starting full datafile backup set
channel c1: specifying datafile(s) in backup set
including current control file in backup set
channel c1: starting piece 1 at 2020-11-04 07:08:48
channel c1: finished piece 1 at 2020-11-04 07:08:49
piece handle=/backup/ctl_CDB1_20201104_25_1 tag=TAG20201104T070847 comment=NONE
channel c1: backup set complete, elapsed time: 00:00:01
Finished backup at 2020-11-04 07:08:49

Starting Control File and SPFILE Autobackup at 2020-11-04 07:08:49
piece handle=/u01/app/oracle/flashback/CDB1/autobackup/2020_11_04/o1_mf_s_1055574529_ht56gksc_.bkp comment=NONE
Finished Control File and SPFILE Autobackup at 2020-11-04 07:08:50

released channel: c1

run
{allocate channel c1 type disk;
sql 'alter system archive log current';
backup archivelog all format '/backup/arch_%d_%T_%s_%p' delete input;
backup current controlfile format '/backup/ctl_%d_%T_%s_%p';
release channel c1;
}

RMAN> list backup of controlfile; 

BS Key  Type LV Size       Device Type Elapsed Time Completion Time    
------- ---- -- ---------- ----------- ------------ -------------------
14      Full    17.95M     DISK        00:00:00     2020-11-04 07:12:25
        BP Key: 14   Status: AVAILABLE  Compressed: NO  Tag: TAG20201104T071225
        Piece Name: /u01/app/oracle/flashback/CDB1/autobackup/2020_11_04/o1_mf_s_1055574745_ht56o9kc_.bkp
  Control File Included: Ckp SCN: 2701595      Ckp time: 2020-11-04 07:12:25
```


## 1.记录控制文件、数据文件头的scn
```
系统检查点
SQL> select checkpoint_change# from v$database;

CHECKPOINT_CHANGE#
------------------
           2701487

set linesize 300
set pagesize 99
col NAME format a60
select name,checkpoint_change# from v$datafile;
NAME                                                         CHECKPOINT_CHANGE#
------------------------------------------------------------ ------------------
/u01/app/oracle/oradata/CDB1/system01.dbf                               2701487
/u01/app/oracle/oradata/CDB1/sysaux01.dbf                               2701487
/u01/app/oracle/oradata/CDB1/undotbs01.dbf                              2701487
/u01/app/oracle/oradata/CDB1/pdbseed/system01.dbf                       1957825
/u01/app/oracle/oradata/CDB1/pdbseed/sysaux01.dbf                       1957825
/u01/app/oracle/oradata/CDB1/users01.dbf                                2701487
/u01/app/oracle/oradata/CDB1/pdbseed/undotbs01.dbf                      1957825
/u01/app/oracle/oradata/CDB1/pdb_easyee/system01.dbf                    2701487
/u01/app/oracle/oradata/CDB1/pdb_easyee/sysaux01.dbf                    2701487
/u01/app/oracle/oradata/CDB1/pdb_easyee/undotbs01.dbf                   2701487
/u01/app/oracle/oradata/CDB1/pdb_easyee/users01.dbf                     2701487
/u01/app/oracle/oradata/CDB1/pdb_easyee/pdbeasyee_data01.dbf            2701487
/u01/app/oracle/oradata/CDB1/pdb_easyee/bttest.dbf                      2701487
/u01/app/oracle/oradata/CDB1/pdb_easyee/undotbs01_2.dbf                 2701487

14 rows selected.

select name,checkpoint_change# from v$datafile_header;

NAME                                                         CHECKPOINT_CHANGE#
------------------------------------------------------------ ------------------
/u01/app/oracle/oradata/CDB1/system01.dbf                               2701487
/u01/app/oracle/oradata/CDB1/sysaux01.dbf                               2701487
/u01/app/oracle/oradata/CDB1/undotbs01.dbf                              2701487
/u01/app/oracle/oradata/CDB1/pdbseed/system01.dbf                       1957825
/u01/app/oracle/oradata/CDB1/pdbseed/sysaux01.dbf                       1957825
/u01/app/oracle/oradata/CDB1/users01.dbf                                2701487
/u01/app/oracle/oradata/CDB1/pdbseed/undotbs01.dbf                      1957825
/u01/app/oracle/oradata/CDB1/pdb_easyee/system01.dbf                    2701487
/u01/app/oracle/oradata/CDB1/pdb_easyee/sysaux01.dbf                    2701487
/u01/app/oracle/oradata/CDB1/pdb_easyee/undotbs01.dbf                   2701487
/u01/app/oracle/oradata/CDB1/pdb_easyee/users01.dbf                     2701487
/u01/app/oracle/oradata/CDB1/pdb_easyee/pdbeasyee_data01.dbf            2701487
/u01/app/oracle/oradata/CDB1/pdb_easyee/bttest.dbf                      2701487
/u01/app/oracle/oradata/CDB1/pdb_easyee/undotbs01_2.dbf                 2701487

14 rows selected.
```

--正常运行时last_change#的值就是空
```
SQL> select name,last_change# from v$datafile;

NAME                                                         LAST_CHANGE#
------------------------------------------------------------ ------------
/u01/app/oracle/oradata/CDB1/system01.dbf
/u01/app/oracle/oradata/CDB1/sysaux01.dbf
/u01/app/oracle/oradata/CDB1/undotbs01.dbf
/u01/app/oracle/oradata/CDB1/pdbseed/system01.dbf                 1957825
/u01/app/oracle/oradata/CDB1/pdbseed/sysaux01.dbf                 1957825
/u01/app/oracle/oradata/CDB1/users01.dbf
/u01/app/oracle/oradata/CDB1/pdbseed/undotbs01.dbf                1957825
/u01/app/oracle/oradata/CDB1/pdb_easyee/system01.dbf
/u01/app/oracle/oradata/CDB1/pdb_easyee/sysaux01.dbf
/u01/app/oracle/oradata/CDB1/pdb_easyee/undotbs01.dbf
/u01/app/oracle/oradata/CDB1/pdb_easyee/users01.dbf
/u01/app/oracle/oradata/CDB1/pdb_easyee/pdbeasyee_data01.dbf
/u01/app/oracle/oradata/CDB1/pdb_easyee/bttest.dbf
/u01/app/oracle/oradata/CDB1/pdb_easyee/undotbs01_2.dbf

14 rows selected.
```
## 2.关闭数据库并移动控制文件
```
SQL>shutdown immediate;
Database closed.
Database dismounted.
ORACLE instance shut down.

oracle@ORACLE-212$pwd
/u01/app/oracle/oradata/CDB1
oracle@ORACLE-212$mv control01.ctl control01.ctl.bak
oracle@ORACLE-212$mv control02.ctl control02.ctl.bak
```
## 3.开启数据库到nomount;
```
oracle@ORACLE-212$export ORACLE_SID=cdb1
oracle@ORACLE-212$echo $ORACLE_SID      
cdb1
oracle@ORACLE-212$sqlplus / as sysdba

SQL*Plus: Release 19.0.0.0.0 - Production on Wed Nov 4 07:36:57 2020
Version 19.3.0.0.0

Copyright (c) 1982, 2019, Oracle.  All rights reserved.

Connected to an idle instance.

SQL> startup nomount;
ORACLE instance started.

Total System Global Area 2466250360 bytes
Fixed Size                  9137784 bytes
Variable Size             536870912 bytes
Database Buffers         1912602624 bytes
Redo Buffers                7639040 bytes
SQL>select status from v$instance;

STATUS
------------------------------------
STARTED
```
## 4.使用rman恢复历史备份的控制文件
```
oracle@ORACLE-212$export ORACLE_SID=cdb1
oracle@ORACLE-212$rman target /

Recovery Manager: Release 19.0.0.0.0 - Production on Wed Nov 4 07:38:51 2020
Version 19.3.0.0.0

Copyright (c) 1982, 2019, Oracle and/or its affiliates.  All rights reserved.

connected to target database: CDB1 (not mounted)
RMAN> restore controlfile from autobackup;

Starting restore at 2020-11-04 07:39:05
using target database control file instead of recovery catalog
allocated channel: ORA_DISK_1
channel ORA_DISK_1: SID=256 device type=DISK

recovery area destination: /u01/app/oracle/flashback
database name (or database unique name) used for search: CDB1
channel ORA_DISK_1: AUTOBACKUP /u01/app/oracle/flashback/CDB1/autobackup/2020_11_04/o1_mf_s_1055574745_ht56o9kc_.bkp found in the recovery area
AUTOBACKUP search with format "%F" not attempted because DBID was not set
channel ORA_DISK_1: restoring control file from AUTOBACKUP /u01/app/oracle/flashback/CDB1/autobackup/2020_11_04/o1_mf_s_1055574745_ht56o9kc_.bkp
channel ORA_DISK_1: control file restore from AUTOBACKUP complete
output file name=/u01/app/oracle/oradata/CDB1/control01.ctl
output file name=/u01/app/oracle/oradata/CDB1/control02.ctl
Finished restore at 2020-11-04 07:39:07
```
## 5.更新数据库到mount
```
SQL>alter database mount;

Database altered.
```
## 6.查看控制文件和数据文件中的scn号
--可以看到数据文件头中记录的scn(18504054)大于控制文件中记录的scn(18261588)
```
set line 999
col name for a50
select status from v$instance;

STATUS
------------------------------------
MOUNTED

select checkpoint_change# from v$database;

CHECKPOINT_CHANGE#
------------------
           2701487

set linesize 300
set pagesize 99
col NAME format a60
select name,checkpoint_change# from v$datafile;

NAME                                                         CHECKPOINT_CHANGE#
------------------------------------------------------------ ------------------
/u01/app/oracle/oradata/CDB1/system01.dbf                               2701487
/u01/app/oracle/oradata/CDB1/sysaux01.dbf                               2701487
/u01/app/oracle/oradata/CDB1/undotbs01.dbf                              2701487
/u01/app/oracle/oradata/CDB1/pdbseed/system01.dbf                       1957825
/u01/app/oracle/oradata/CDB1/pdbseed/sysaux01.dbf                       1957825
/u01/app/oracle/oradata/CDB1/users01.dbf                                2701487
/u01/app/oracle/oradata/CDB1/pdbseed/undotbs01.dbf                      1957825
/u01/app/oracle/oradata/CDB1/pdb_easyee/system01.dbf                    2701487
/u01/app/oracle/oradata/CDB1/pdb_easyee/sysaux01.dbf                    2701487
/u01/app/oracle/oradata/CDB1/pdb_easyee/undotbs01.dbf                   2701487
/u01/app/oracle/oradata/CDB1/pdb_easyee/users01.dbf                     2701487
/u01/app/oracle/oradata/CDB1/pdb_easyee/pdbeasyee_data01.dbf            2701487
/u01/app/oracle/oradata/CDB1/pdb_easyee/bttest.dbf                      2701487
/u01/app/oracle/oradata/CDB1/pdb_easyee/undotbs01_2.dbf                 2701487

14 rows selected.

select name,checkpoint_change# from v$datafile_header;

NAME                                                         CHECKPOINT_CHANGE#
------------------------------------------------------------ ------------------
/u01/app/oracle/oradata/CDB1/system01.dbf                               2702560
/u01/app/oracle/oradata/CDB1/sysaux01.dbf                               2702560
/u01/app/oracle/oradata/CDB1/undotbs01.dbf                              2702560
/u01/app/oracle/oradata/CDB1/pdbseed/system01.dbf                       1957825
/u01/app/oracle/oradata/CDB1/pdbseed/sysaux01.dbf                       1957825
/u01/app/oracle/oradata/CDB1/users01.dbf                                2702560
/u01/app/oracle/oradata/CDB1/pdbseed/undotbs01.dbf                      1957825
/u01/app/oracle/oradata/CDB1/pdb_easyee/system01.dbf                    2702455
/u01/app/oracle/oradata/CDB1/pdb_easyee/sysaux01.dbf                    2702455
/u01/app/oracle/oradata/CDB1/pdb_easyee/undotbs01.dbf                   2702455
/u01/app/oracle/oradata/CDB1/pdb_easyee/users01.dbf                     2702455
/u01/app/oracle/oradata/CDB1/pdb_easyee/pdbeasyee_data01.dbf            2702455
/u01/app/oracle/oradata/CDB1/pdb_easyee/bttest.dbf                      2702455
/u01/app/oracle/oradata/CDB1/pdb_easyee/undotbs01_2.dbf                 2702455

14 rows selected.

select name,last_change# from v$datafile;

NAME                                                         LAST_CHANGE#
------------------------------------------------------------ ------------
/u01/app/oracle/oradata/CDB1/system01.dbf
/u01/app/oracle/oradata/CDB1/sysaux01.dbf
/u01/app/oracle/oradata/CDB1/undotbs01.dbf
/u01/app/oracle/oradata/CDB1/pdbseed/system01.dbf
/u01/app/oracle/oradata/CDB1/pdbseed/sysaux01.dbf
/u01/app/oracle/oradata/CDB1/users01.dbf
/u01/app/oracle/oradata/CDB1/pdbseed/undotbs01.dbf
/u01/app/oracle/oradata/CDB1/pdb_easyee/system01.dbf
/u01/app/oracle/oradata/CDB1/pdb_easyee/sysaux01.dbf
/u01/app/oracle/oradata/CDB1/pdb_easyee/undotbs01.dbf
/u01/app/oracle/oradata/CDB1/pdb_easyee/users01.dbf
/u01/app/oracle/oradata/CDB1/pdb_easyee/pdbeasyee_data01.dbf
/u01/app/oracle/oradata/CDB1/pdb_easyee/bttest.dbf
/u01/app/oracle/oradata/CDB1/pdb_easyee/undotbs01_2.dbf

14 rows selected.
```
## 7.更新数据库为open
```
>alter database open;
alter database open
*
ERROR at line 1:
ORA-01589: must use RESETLOGS or NORESETLOGS option for database open

>alter database open noresetlogs;
alter database open noresetlogs
*
ERROR at line 1:
ORA-01610: recovery using the BACKUP CONTROLFILE option must be done  --恢复备份的控制文件不可以noresetlogs方式开库

>alter database open resetlogs;
alter database open resetlogs
*
ERROR at line 1:
ORA-01152: file 1 was not restored from a sufficiently old backup  --说明使用了一个旧的控制文件与现有的数据文件的scn对不上
ORA-01110: data file 1: '/u01/app/oracle/oradata/CDB1/system01.dbf'

>recover database; 
ORA-00283: recovery session canceled due to errors  
ORA-01610: recovery using the BACKUP CONTROLFILE option must be done  --直接恢复报错,必须执行使用备份的控制文件进行恢复数据库


>recover database using backup controlfile;
ORA-00279: change 2701595 generated at 11/04/2020 07:10:05 needed for thread 1
ORA-00289: suggestion : /u01/app/oracle/flashback/CDB1/archivelog/2020_11_04/o1_mf_1_28_%u_.arc
ORA-00280: change 2701595 for thread 1 is in sequence #28


Specify log: {<RET>=suggested | filename | AUTO | CANCEL}
auto
ORA-00279: change 18292695 generated at 03/10/2020 13:24:26 needed for thread 1
ORA-00289: suggestion : /u01/app/orarch/enmo/1_124_1033390448.arc
ORA-00280: change 18292695 for thread 1 is in sequence #124
ORA-00278: log file '/u01/app/orarch/enmo/1_123_1033390448.arc' no longer needed for this recovery

......
......
......


ORA-00308: cannot open archived log '/u01/app/oracle/flashback/CDB1/archivelog/2020_11_04/o1_mf_1_28_%u_.arc'    --提示需要28日志,但是归档目录没有,手动指定redo日志(具体是redo那个日志不太清楚,分别尝试每个redo)
ORA-27037: unable to obtain file status
Linux-x86_64 Error: 2: No such file or directory
Additional information: 7==

>recover database using backup controlfile;
ORA-00279: change 18504054 generated at 03/13/2020 12:45:47 needed for thread 1
ORA-00289: suggestion : /u01/app/orarch/enmo/1_146_1033390448.arc
ORA-00280: change 18504054 for thread 1 is in sequence #146

Specify log: {<RET>=suggested | filename | AUTO | CANCEL}
/u01/app/oracle/oradata/CDB1/redo01.log
Log applied.
Media recovery complete.
```
## 8.再次查看控制文件、数据文件头记录的scn号
```
set line 999
col name for a50
select status from v$instance;

STATUS
------------------------------------
MOUNTED

>select checkpoint_change# from v$database;

CHECKPOINT_CHANGE#
------------------
           2701487

set linesize 300
set pagesize 99
col NAME format a60
select name,checkpoint_change# from v$datafile;

NAME                                                         CHECKPOINT_CHANGE#
------------------------------------------------------------ ------------------
/u01/app/oracle/oradata/CDB1/system01.dbf                               2701487
/u01/app/oracle/oradata/CDB1/sysaux01.dbf                               2701487
/u01/app/oracle/oradata/CDB1/undotbs01.dbf                              2701487
/u01/app/oracle/oradata/CDB1/pdbseed/system01.dbf                       1957825
/u01/app/oracle/oradata/CDB1/pdbseed/sysaux01.dbf                       1957825
/u01/app/oracle/oradata/CDB1/users01.dbf                                2701487
/u01/app/oracle/oradata/CDB1/pdbseed/undotbs01.dbf                      1957825
/u01/app/oracle/oradata/CDB1/pdb_easyee/system01.dbf                    2702455
/u01/app/oracle/oradata/CDB1/pdb_easyee/sysaux01.dbf                    2702455
/u01/app/oracle/oradata/CDB1/pdb_easyee/undotbs01.dbf                   2702455
/u01/app/oracle/oradata/CDB1/pdb_easyee/users01.dbf                     2702455
/u01/app/oracle/oradata/CDB1/pdb_easyee/pdbeasyee_data01.dbf            2702455
/u01/app/oracle/oradata/CDB1/pdb_easyee/bttest.dbf                      2702455
/u01/app/oracle/oradata/CDB1/pdb_easyee/undotbs01_2.dbf                 2702455

14 rows selected.

>select name,checkpoint_change# from v$datafile_header;

NAME                                                         CHECKPOINT_CHANGE#
------------------------------------------------------------ ------------------
/u01/app/oracle/oradata/CDB1/system01.dbf                               2702560
/u01/app/oracle/oradata/CDB1/sysaux01.dbf                               2702560
/u01/app/oracle/oradata/CDB1/undotbs01.dbf                              2702560
/u01/app/oracle/oradata/CDB1/pdbseed/system01.dbf                       1957825
/u01/app/oracle/oradata/CDB1/pdbseed/sysaux01.dbf                       1957825
/u01/app/oracle/oradata/CDB1/users01.dbf                                2702560
/u01/app/oracle/oradata/CDB1/pdbseed/undotbs01.dbf                      1957825
/u01/app/oracle/oradata/CDB1/pdb_easyee/system01.dbf                    2702455
/u01/app/oracle/oradata/CDB1/pdb_easyee/sysaux01.dbf                    2702455
/u01/app/oracle/oradata/CDB1/pdb_easyee/undotbs01.dbf                   2702455
/u01/app/oracle/oradata/CDB1/pdb_easyee/users01.dbf                     2702455
/u01/app/oracle/oradata/CDB1/pdb_easyee/pdbeasyee_data01.dbf            2702455
/u01/app/oracle/oradata/CDB1/pdb_easyee/bttest.dbf                      2702455
/u01/app/oracle/oradata/CDB1/pdb_easyee/undotbs01_2.dbf                 2702455

14 rows selected.
SYS@enmo>select name,last_change# from v$datafile;

NAME                                                         LAST_CHANGE#
------------------------------------------------------------ ------------
/u01/app/oracle/oradata/CDB1/system01.dbf                         2702560
/u01/app/oracle/oradata/CDB1/sysaux01.dbf                         2702560
/u01/app/oracle/oradata/CDB1/undotbs01.dbf                        2702560
/u01/app/oracle/oradata/CDB1/pdbseed/system01.dbf
/u01/app/oracle/oradata/CDB1/pdbseed/sysaux01.dbf
/u01/app/oracle/oradata/CDB1/users01.dbf                          2702560
/u01/app/oracle/oradata/CDB1/pdbseed/undotbs01.dbf
/u01/app/oracle/oradata/CDB1/pdb_easyee/system01.dbf
/u01/app/oracle/oradata/CDB1/pdb_easyee/sysaux01.dbf
/u01/app/oracle/oradata/CDB1/pdb_easyee/undotbs01.dbf
/u01/app/oracle/oradata/CDB1/pdb_easyee/users01.dbf
/u01/app/oracle/oradata/CDB1/pdb_easyee/pdbeasyee_data01.dbf
/u01/app/oracle/oradata/CDB1/pdb_easyee/bttest.dbf
/u01/app/oracle/oradata/CDB1/pdb_easyee/undotbs01_2.dbf

14 rows selected.
```
## 9.使用resetlogs方式开启数据库
```
>alter database open resetlogs;

Database altered.

>select status from v$instance;

STATUS
------------
OPEN

>select checkpoint_change# from v$database;

CHECKPOINT_CHANGE#
------------------
           2703620

>select name,checkpoint_change# from v$datafile;

NAME                                                         CHECKPOINT_CHANGE#
------------------------------------------------------------ ------------------
/u01/app/oracle/oradata/CDB1/system01.dbf                               2703620
/u01/app/oracle/oradata/CDB1/sysaux01.dbf                               2703620
/u01/app/oracle/oradata/CDB1/undotbs01.dbf                              2703620
/u01/app/oracle/oradata/CDB1/pdbseed/system01.dbf                       1957825
/u01/app/oracle/oradata/CDB1/pdbseed/sysaux01.dbf                       1957825
/u01/app/oracle/oradata/CDB1/users01.dbf                                2703620
/u01/app/oracle/oradata/CDB1/pdbseed/undotbs01.dbf                      1957825
/u01/app/oracle/oradata/CDB1/pdb_easyee/system01.dbf                    2703835
/u01/app/oracle/oradata/CDB1/pdb_easyee/sysaux01.dbf                    2703835
/u01/app/oracle/oradata/CDB1/pdb_easyee/undotbs01.dbf                   2703835
/u01/app/oracle/oradata/CDB1/pdb_easyee/users01.dbf                     2703835
/u01/app/oracle/oradata/CDB1/pdb_easyee/pdbeasyee_data01.dbf            2703835
/u01/app/oracle/oradata/CDB1/pdb_easyee/bttest.dbf                      2703835
/u01/app/oracle/oradata/CDB1/pdb_easyee/undotbs01_2.dbf                 2703835

14 rows selected.

>select name,checkpoint_change# from v$datafile_header;

NAME                                                         CHECKPOINT_CHANGE#
------------------------------------------------------------ ------------------
/u01/app/oracle/oradata/CDB1/system01.dbf                               2703620
/u01/app/oracle/oradata/CDB1/sysaux01.dbf                               2703620
/u01/app/oracle/oradata/CDB1/undotbs01.dbf                              2703620
/u01/app/oracle/oradata/CDB1/pdbseed/system01.dbf                       1957825
/u01/app/oracle/oradata/CDB1/pdbseed/sysaux01.dbf                       1957825
/u01/app/oracle/oradata/CDB1/users01.dbf                                2703620
/u01/app/oracle/oradata/CDB1/pdbseed/undotbs01.dbf                      1957825
/u01/app/oracle/oradata/CDB1/pdb_easyee/system01.dbf                    2703835
/u01/app/oracle/oradata/CDB1/pdb_easyee/sysaux01.dbf                    2703835
/u01/app/oracle/oradata/CDB1/pdb_easyee/undotbs01.dbf                   2703835
/u01/app/oracle/oradata/CDB1/pdb_easyee/users01.dbf                     2703835
/u01/app/oracle/oradata/CDB1/pdb_easyee/pdbeasyee_data01.dbf            2703835
/u01/app/oracle/oradata/CDB1/pdb_easyee/bttest.dbf                      2703835
/u01/app/oracle/oradata/CDB1/pdb_easyee/undotbs01_2.dbf                 2703835

14 rows selected.

>select name,last_change# from v$datafile;

NAME                                                         LAST_CHANGE#
------------------------------------------------------------ ------------
/u01/app/oracle/oradata/CDB1/system01.dbf
/u01/app/oracle/oradata/CDB1/sysaux01.dbf
/u01/app/oracle/oradata/CDB1/undotbs01.dbf
/u01/app/oracle/oradata/CDB1/pdbseed/system01.dbf                 1957825
/u01/app/oracle/oradata/CDB1/pdbseed/sysaux01.dbf                 1957825
/u01/app/oracle/oradata/CDB1/users01.dbf
/u01/app/oracle/oradata/CDB1/pdbseed/undotbs01.dbf                1957825
/u01/app/oracle/oradata/CDB1/pdb_easyee/system01.dbf
/u01/app/oracle/oradata/CDB1/pdb_easyee/sysaux01.dbf
/u01/app/oracle/oradata/CDB1/pdb_easyee/undotbs01.dbf
/u01/app/oracle/oradata/CDB1/pdb_easyee/users01.dbf
/u01/app/oracle/oradata/CDB1/pdb_easyee/pdbeasyee_data01.dbf
/u01/app/oracle/oradata/CDB1/pdb_easyee/bttest.dbf
/u01/app/oracle/oradata/CDB1/pdb_easyee/undotbs01_2.dbf

14 rows selected.
```
## 10.查看化身的数据
```
RMAN> list incarnation;

List of Database Incarnations
DB Key  Inc Key DB Name  DB ID            STATUS  Reset SCN  Reset Time
------- ------- -------- ---------------- --- ---------- ----------
1       1       CDB1     1010358776       PARENT  1          2019-04-17 00:55:59
2       2       CDB1     1010358776       PARENT  1920977    2020-04-02 03:55:38
3       3       CDB1     1010358776       CURRENT 2702561    2020-11-04 07:56:30  --从2702561开始新的数据库化身

```