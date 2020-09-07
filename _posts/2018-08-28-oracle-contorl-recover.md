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
## 1.记录控制文件、数据文件头的scn
```
SYS@enmo>select checkpoint_change# from v$database;

CHECKPOINT_CHANGE#
------------------
          18502624

SYS@enmo>select name,checkpoint_change# from v$datafile;

NAME                                               CHECKPOINT_CHANGE#
-------------------------------------------------- ------------------
/u01/app/dbfile/enmo/system01.dbf                            18502624
/u01/app/dbfile/enmo/sysaux01.dbf                            18502624
/u01/app/dbfile/enmo/undotbs1.dbf                            18502624
/u01/app/dbfile/enmo/users01.dbf                             18502624
/u01/app/dbfile/enmo/chang01.dbf                             18502624
/u01/app/dbfile/enmo/lob_data01.dbf                          18502624
/u01/app/dbfile/enmo/oggdata.dbf                             18502624
/u01/app/dbfile/reccat.dbf                                   18502624

8 rows selected.

SYS@enmo>select name,checkpoint_change# from v$datafile_header;

NAME                                               CHECKPOINT_CHANGE#
-------------------------------------------------- ------------------
/u01/app/dbfile/enmo/system01.dbf                            18502624
/u01/app/dbfile/enmo/sysaux01.dbf                            18502624
/u01/app/dbfile/enmo/undotbs1.dbf                            18502624
/u01/app/dbfile/enmo/users01.dbf                             18502624
/u01/app/dbfile/enmo/chang01.dbf                             18502624
/u01/app/dbfile/enmo/lob_data01.dbf                          18502624
/u01/app/dbfile/enmo/oggdata.dbf                             18502624
/u01/app/dbfile/reccat.dbf                                   18502624

8 rows selected.
```

--正常运行时last_change#的值就是空
```
SYS@enmo>select name,last_change# from v$datafile;

NAME                                               LAST_CHANGE#
-------------------------------------------------- ------------
/u01/app/dbfile/enmo/system01.dbf
/u01/app/dbfile/enmo/sysaux01.dbf
/u01/app/dbfile/enmo/undotbs1.dbf
/u01/app/dbfile/enmo/users01.dbf
/u01/app/dbfile/enmo/chang01.dbf
/u01/app/dbfile/enmo/lob_data01.dbf
/u01/app/dbfile/enmo/oggdata.dbf
/u01/app/dbfile/reccat.dbf

8 rows selected.
```
## 2.关闭数据库并移动控制文件
```
SYS@enmo>shutdown immediate;
Database closed.
Database dismounted.
ORACLE instance shut down.

[oracle@ora:enmo enmo]$mv control01.ctl control01.ctl.bak
[oracle@ora:enmo enmo]$mv control02.ctl control02.ctl.bak
```
## 3.开启数据库到nomount;
```
SYS@enmo>startup nomount;
ORACLE instance started.

Total System Global Area 1048576000 bytes
Fixed Size                  8628640 bytes
Variable Size             796919392 bytes
Database Buffers          234881024 bytes
Redo Buffers                8146944 bytes
SYS@enmo>select status from v$instance;

STATUS
------------------------------------
STARTED
```
## 4.使用rman恢复历史备份的控制文件
```
RMAN> restore controlfile from autobackup;

Starting restore at 2020-03-13 12:51:11 

allocated channel: ORA_DISK_1
channel ORA_DISK_1: SID=10 device type=DISK

recovery area destination: /u01/app/fra
database name (or database unique name) used for search: ENMO
channel ORA_DISK_1: AUTOBACKUP /u01/app/fra/ENMO/autobackup/2020_03_10/o1_mf_s_1034681351_h6g2d7x3_.bkp found in the recovery area
AUTOBACKUP search with format "%F" not attempted because DBID was not set
channel ORA_DISK_1: restoring control file from AUTOBACKUP /u01/app/fra/ENMO/autobackup/2020_03_10/o1_mf_s_1034681351_h6g2d7x3_.bkp
channel ORA_DISK_1: control file restore from AUTOBACKUP complete
output file name=/u01/app/dbfile/enmo/control01.ctl
output file name=/u01/app/dbfile/enmo/control02.ctl
Finished restore at 2020-03-13 12:51:13 
```
## 5.更新数据库到mount
```
SYS@enmo>alter database mount;

Database altered.
```
## 6.查看控制文件和数据文件中的scn号
--可以看到数据文件头中记录的scn(18504054)大于控制文件中记录的scn(18261588)
```
SYS@enmo>set line 999
SYS@enmo>col name for a50
SYS@enmo>select status from v$instance;

STATUS
------------------------------------
MOUNTED

SYS@enmo>select checkpoint_change# from v$database;

CHECKPOINT_CHANGE#
------------------
          18261588

SYS@enmo>select name,checkpoint_change# from v$datafile;

NAME                                               CHECKPOINT_CHANGE#
-------------------------------------------------- ------------------
/u01/app/dbfile/enmo/system01.dbf                            18261588
/u01/app/dbfile/enmo/sysaux01.dbf                            18261588
/u01/app/dbfile/enmo/undotbs1.dbf                            18261588
/u01/app/dbfile/enmo/users01.dbf                             18261588
/u01/app/dbfile/enmo/chang01.dbf                             18261588
/u01/app/dbfile/enmo/lob_data01.dbf                          18261588
/u01/app/dbfile/enmo/oggdata.dbf                             18261588
/u01/app/dbfile/reccat.dbf                                   18261588

8 rows selected.

SYS@enmo>select name,checkpoint_change# from v$datafile_header;

NAME                                               CHECKPOINT_CHANGE#
-------------------------------------------------- ------------------
/u01/app/dbfile/enmo/system01.dbf                            18504054
/u01/app/dbfile/enmo/sysaux01.dbf                            18504054
/u01/app/dbfile/enmo/undotbs1.dbf                            18504054
/u01/app/dbfile/enmo/users01.dbf                             18504054
/u01/app/dbfile/enmo/chang01.dbf                             18504054
/u01/app/dbfile/enmo/lob_data01.dbf                          18504054
/u01/app/dbfile/enmo/oggdata.dbf                             18504054
/u01/app/dbfile/reccat.dbf                                   18504054

8 rows selected.

SYS@enmo>select name,last_change# from v$datafile;

NAME                                               LAST_CHANGE#
-------------------------------------------------- ------------
/u01/app/dbfile/enmo/system01.dbf
/u01/app/dbfile/enmo/sysaux01.dbf
/u01/app/dbfile/enmo/undotbs1.dbf
/u01/app/dbfile/enmo/users01.dbf
/u01/app/dbfile/enmo/chang01.dbf
/u01/app/dbfile/enmo/lob_data01.dbf
/u01/app/dbfile/enmo/oggdata.dbf
/u01/app/dbfile/reccat.dbf

8 rows selected.
```
## 7.更新数据库为open
```
SYS@enmo>alter database open;
alter database open
*
ERROR at line 1:
ORA-01589: must use RESETLOGS or NORESETLOGS option for database open

SYS@enmo>alter database open noresetlogs;
alter database open noresetlogs
*
ERROR at line 1:
ORA-01610: recovery using the BACKUP CONTROLFILE option must be done  --恢复备份的控制文件不可以noresetlogs方式开库

SYS@enmo>alter database open resetlogs;
alter database open resetlogs
*
ERROR at line 1:
ORA-01152: file 1 was not restored from a sufficiently old backup  --说明使用了一个旧的控制文件与现有的数据文件的scn对不上
ORA-01110: data file 1: '/u01/app/dbfile/enmo/system01.dbf'

SYS@enmo>recover database; 
ORA-00283: recovery session canceled due to errors  
ORA-01610: recovery using the BACKUP CONTROLFILE option must be done  --直接恢复报错,必须执行使用备份的控制文件进行恢复数据库


SYS@enmo>recover database using backup controlfile;
ORA-00279: change 18282978 generated at 03/10/2020 09:25:09 needed for thread 1
ORA-00289: suggestion : /u01/app/orarch/enmo/1_123_1033390448.arc
ORA-00280: change 18282978 for thread 1 is in sequence #123


Specify log: {<RET>=suggested | filename | AUTO | CANCEL}
auto
ORA-00279: change 18292695 generated at 03/10/2020 13:24:26 needed for thread 1
ORA-00289: suggestion : /u01/app/orarch/enmo/1_124_1033390448.arc
ORA-00280: change 18292695 for thread 1 is in sequence #124
ORA-00278: log file '/u01/app/orarch/enmo/1_123_1033390448.arc' no longer needed for this recovery


ORA-00279: change 18293343 generated at 03/10/2020 13:37:38 needed for thread 1
ORA-00289: suggestion : /u01/app/orarch/enmo/1_125_1033390448.arc
ORA-00280: change 18293343 for thread 1 is in sequence #125
ORA-00278: log file '/u01/app/orarch/enmo/1_124_1033390448.arc' no longer needed for this recovery


ORA-00279: change 18354143 generated at 03/11/2020 09:36:32 needed for thread 1
ORA-00289: suggestion : /u01/app/orarch/enmo/1_126_1033390448.arc
ORA-00280: change 18354143 for thread 1 is in sequence #126
ORA-00278: log file '/u01/app/orarch/enmo/1_125_1033390448.arc' no longer needed for this recovery


ORA-00279: change 18454671 generated at 03/12/2020 15:35:37 needed for thread 1
ORA-00289: suggestion : /u01/app/orarch/enmo/1_127_1033390448.arc
ORA-00280: change 18454671 for thread 1 is in sequence #127
ORA-00278: log file '/u01/app/orarch/enmo/1_126_1033390448.arc' no longer needed for this recovery


ORA-00279: change 18460259 generated at 03/12/2020 15:41:41 needed for thread 1
ORA-00289: suggestion : /u01/app/orarch/enmo/1_128_1033390448.arc
ORA-00280: change 18460259 for thread 1 is in sequence #128
ORA-00278: log file '/u01/app/orarch/enmo/1_127_1033390448.arc' no longer needed for this recovery


ORA-00279: change 18463068 generated at 03/12/2020 15:56:59 needed for thread 1
ORA-00289: suggestion : /u01/app/orarch/enmo/1_129_1033390448.arc
ORA-00280: change 18463068 for thread 1 is in sequence #129
ORA-00278: log file '/u01/app/orarch/enmo/1_128_1033390448.arc' no longer needed for this recovery


ORA-00279: change 18464597 generated at 03/12/2020 16:03:39 needed for thread 1
ORA-00289: suggestion : /u01/app/orarch/enmo/1_130_1033390448.arc
ORA-00280: change 18464597 for thread 1 is in sequence #130
ORA-00278: log file '/u01/app/orarch/enmo/1_129_1033390448.arc' no longer needed for this recovery


ORA-00279: change 18465698 generated at 03/12/2020 16:04:21 needed for thread 1
ORA-00289: suggestion : /u01/app/orarch/enmo/1_131_1033390448.arc
ORA-00280: change 18465698 for thread 1 is in sequence #131
ORA-00278: log file '/u01/app/orarch/enmo/1_130_1033390448.arc' no longer needed for this recovery


ORA-00279: change 18465714 generated at 03/12/2020 16:04:27 needed for thread 1
ORA-00289: suggestion : /u01/app/orarch/enmo/1_132_1033390448.arc
ORA-00280: change 18465714 for thread 1 is in sequence #132
ORA-00278: log file '/u01/app/orarch/enmo/1_131_1033390448.arc' no longer needed for this recovery


ORA-00279: change 18465757 generated at 03/12/2020 16:04:47 needed for thread 1
ORA-00289: suggestion : /u01/app/orarch/enmo/1_133_1033390448.arc
ORA-00280: change 18465757 for thread 1 is in sequence #133
ORA-00278: log file '/u01/app/orarch/enmo/1_132_1033390448.arc' no longer needed for this recovery


ORA-00279: change 18467282 generated at 03/12/2020 16:16:58 needed for thread 1
ORA-00289: suggestion : /u01/app/orarch/enmo/1_134_1033390448.arc
ORA-00280: change 18467282 for thread 1 is in sequence #134
ORA-00278: log file '/u01/app/orarch/enmo/1_133_1033390448.arc' no longer needed for this recovery


ORA-00279: change 18469175 generated at 03/12/2020 16:25:01 needed for thread 1
ORA-00289: suggestion : /u01/app/orarch/enmo/1_135_1033390448.arc
ORA-00280: change 18469175 for thread 1 is in sequence #135
ORA-00278: log file '/u01/app/orarch/enmo/1_134_1033390448.arc' no longer needed for this recovery


ORA-00279: change 18469629 generated at 03/12/2020 16:28:31 needed for thread 1
ORA-00289: suggestion : /u01/app/orarch/enmo/1_136_1033390448.arc
ORA-00280: change 18469629 for thread 1 is in sequence #136
ORA-00278: log file '/u01/app/orarch/enmo/1_135_1033390448.arc' no longer needed for this recovery


ORA-00279: change 18469828 generated at 03/12/2020 16:30:07 needed for thread 1
ORA-00289: suggestion : /u01/app/orarch/enmo/1_137_1033390448.arc
ORA-00280: change 18469828 for thread 1 is in sequence #137
ORA-00278: log file '/u01/app/orarch/enmo/1_136_1033390448.arc' no longer needed for this recovery


ORA-00279: change 18470068 generated at 03/12/2020 16:31:32 needed for thread 1
ORA-00289: suggestion : /u01/app/orarch/enmo/1_138_1033390448.arc
ORA-00280: change 18470068 for thread 1 is in sequence #138
ORA-00278: log file '/u01/app/orarch/enmo/1_137_1033390448.arc' no longer needed for this recovery


ORA-00279: change 18470071 generated at 03/12/2020 16:33:14 needed for thread 1
ORA-00289: suggestion : /u01/app/orarch/enmo/1_139_1033390448.arc
ORA-00280: change 18470071 for thread 1 is in sequence #139
ORA-00278: log file '/u01/app/orarch/enmo/1_138_1033390448.arc' no longer needed for this recovery


ORA-00279: change 18470909 generated at 03/12/2020 16:33:18 needed for thread 1
ORA-00289: suggestion : /u01/app/orarch/enmo/1_140_1033390448.arc
ORA-00280: change 18470909 for thread 1 is in sequence #140
ORA-00278: log file '/u01/app/orarch/enmo/1_139_1033390448.arc' no longer needed for this recovery


ORA-00279: change 18484607 generated at 03/13/2020 09:24:42 needed for thread 1
ORA-00289: suggestion : /u01/app/orarch/enmo/1_141_1033390448.arc
ORA-00280: change 18484607 for thread 1 is in sequence #141
ORA-00278: log file '/u01/app/orarch/enmo/1_140_1033390448.arc' no longer needed for this recovery


ORA-00279: change 18485561 generated at 03/13/2020 09:29:58 needed for thread 1
ORA-00289: suggestion : /u01/app/orarch/enmo/1_142_1033390448.arc
ORA-00280: change 18485561 for thread 1 is in sequence #142
ORA-00278: log file '/u01/app/orarch/enmo/1_141_1033390448.arc' no longer needed for this recovery


ORA-00279: change 18502620 generated at 03/13/2020 11:35:04 needed for thread 1
ORA-00289: suggestion : /u01/app/orarch/enmo/1_143_1033390448.arc
ORA-00280: change 18502620 for thread 1 is in sequence #143
ORA-00278: log file '/u01/app/orarch/enmo/1_142_1033390448.arc' no longer needed for this recovery


ORA-00279: change 18502623 generated at 03/13/2020 12:42:33 needed for thread 1
ORA-00289: suggestion : /u01/app/orarch/enmo/1_144_1033390448.arc
ORA-00280: change 18502623 for thread 1 is in sequence #144
ORA-00278: log file '/u01/app/orarch/enmo/1_143_1033390448.arc' no longer needed for this recovery


ORA-00279: change 18502793 generated at 03/13/2020 12:42:37 needed for thread 1
ORA-00289: suggestion : /u01/app/orarch/enmo/1_145_1033390448.arc
ORA-00280: change 18502793 for thread 1 is in sequence #145
ORA-00278: log file '/u01/app/orarch/enmo/1_144_1033390448.arc' no longer needed for this recovery


ORA-00279: change 18504054 generated at 03/13/2020 12:45:47 needed for thread 1
ORA-00289: suggestion : /u01/app/orarch/enmo/1_146_1033390448.arc
ORA-00280: change 18504054 for thread 1 is in sequence #146
ORA-00278: log file '/u01/app/orarch/enmo/1_145_1033390448.arc' no longer needed for this recovery


ORA-00308: cannot open archived log '/u01/app/orarch/enmo/1_146_1033390448.arc'    --提示需要146日志,但是归档目录没有,手动指定redo日志(具体是redo那个日志不太清楚,分别尝试每个redo)
ORA-27037: unable to obtain file status
Linux-x86_64 Error: 2: No such file or directory
Additional information: 7==

SYS@enmo>recover database using backup controlfile;
ORA-00279: change 18504054 generated at 03/13/2020 12:45:47 needed for thread 1
ORA-00289: suggestion : /u01/app/orarch/enmo/1_146_1033390448.arc
ORA-00280: change 18504054 for thread 1 is in sequence #146


Specify log: {<RET>=suggested | filename | AUTO | CANCEL}
/u01/app/oraredo/enmo/redo04a.rdo
Log applied.
Media recovery complete.
```
## 8.再次查看控制文件、数据文件头记录的scn号
```
SYS@enmo>set line 999
SYS@enmo>col name for a50
SYS@enmo>select status from v$instance;

STATUS
------------------------------------
MOUNTED

SYS@enmo>select checkpoint_change# from v$database;

CHECKPOINT_CHANGE#
------------------
          18261588

SYS@enmo>select name,checkpoint_change# from v$datafile;

NAME                                               CHECKPOINT_CHANGE#
-------------------------------------------------- ------------------
/u01/app/dbfile/enmo/system01.dbf                            18504055
/u01/app/dbfile/enmo/sysaux01.dbf                            18504055
/u01/app/dbfile/enmo/undotbs1.dbf                            18504055
/u01/app/dbfile/enmo/users01.dbf                             18504055
/u01/app/dbfile/enmo/chang01.dbf                             18504055
/u01/app/dbfile/enmo/lob_data01.dbf                          18504055
/u01/app/dbfile/enmo/oggdata.dbf                             18504055
/u01/app/dbfile/reccat.dbf                                   18504055

8 rows selected.

SYS@enmo>select name,checkpoint_change# from v$datafile_header;

NAME                                               CHECKPOINT_CHANGE#
-------------------------------------------------- ------------------
/u01/app/dbfile/enmo/system01.dbf                            18504055
/u01/app/dbfile/enmo/sysaux01.dbf                            18504055
/u01/app/dbfile/enmo/undotbs1.dbf                            18504055
/u01/app/dbfile/enmo/users01.dbf                             18504055
/u01/app/dbfile/enmo/chang01.dbf                             18504055
/u01/app/dbfile/enmo/lob_data01.dbf                          18504055
/u01/app/dbfile/enmo/oggdata.dbf                             18504055
/u01/app/dbfile/reccat.dbf                                   18504055

8 rows selected.

SYS@enmo>select name,last_change# from v$datafile;

NAME                                               LAST_CHANGE#
-------------------------------------------------- ------------
/u01/app/dbfile/enmo/system01.dbf                      18504055
/u01/app/dbfile/enmo/sysaux01.dbf                      18504055
/u01/app/dbfile/enmo/undotbs1.dbf                      18504055
/u01/app/dbfile/enmo/users01.dbf                       18504055
/u01/app/dbfile/enmo/chang01.dbf                       18504055
/u01/app/dbfile/enmo/lob_data01.dbf                    18504055
/u01/app/dbfile/enmo/oggdata.dbf                       18504055
/u01/app/dbfile/reccat.dbf                             18504055

8 rows selected.
```
## 9.使用resetlogs方式开启数据库
```
SYS@enmo>alter database open resetlogs;

Database altered.
SYS@enmo>set line 999
SYS@enmo>col name for a50
SYS@enmo>select status from v$instance;

STATUS
------------
OPEN

SYS@enmo>select checkpoint_change# from v$database;

CHECKPOINT_CHANGE#
------------------
          18504060

SYS@enmo>select name,checkpoint_change# from v$datafile;

NAME                                               CHECKPOINT_CHANGE#
-------------------------------------------------- ------------------
/u01/app/dbfile/enmo/system01.dbf                            18504060
/u01/app/dbfile/enmo/sysaux01.dbf                            18504060
/u01/app/dbfile/enmo/undotbs1.dbf                            18504060
/u01/app/dbfile/enmo/users01.dbf                             18504060
/u01/app/dbfile/enmo/chang01.dbf                             18504060
/u01/app/dbfile/enmo/lob_data01.dbf                          18504060
/u01/app/dbfile/enmo/oggdata.dbf                             18504060
/u01/app/dbfile/reccat.dbf                                   18504060

8 rows selected.

SYS@enmo>select name,checkpoint_change# from v$datafile_header;

NAME                                               CHECKPOINT_CHANGE#
-------------------------------------------------- ------------------
/u01/app/dbfile/enmo/system01.dbf                            18504060
/u01/app/dbfile/enmo/sysaux01.dbf                            18504060
/u01/app/dbfile/enmo/undotbs1.dbf                            18504060
/u01/app/dbfile/enmo/users01.dbf                             18504060
/u01/app/dbfile/enmo/chang01.dbf                             18504060
/u01/app/dbfile/enmo/lob_data01.dbf                          18504060
/u01/app/dbfile/enmo/oggdata.dbf                             18504060
/u01/app/dbfile/reccat.dbf                                   18504060

8 rows selected.

SYS@enmo>select name,last_change# from v$datafile;

NAME                                               LAST_CHANGE#
-------------------------------------------------- ------------
/u01/app/dbfile/enmo/system01.dbf
/u01/app/dbfile/enmo/sysaux01.dbf
/u01/app/dbfile/enmo/undotbs1.dbf
/u01/app/dbfile/enmo/users01.dbf
/u01/app/dbfile/enmo/chang01.dbf
/u01/app/dbfile/enmo/lob_data01.dbf
/u01/app/dbfile/enmo/oggdata.dbf
/u01/app/dbfile/reccat.dbf

8 rows selected.
```
## 10.查看化身的数据
```
RMAN> list incarnation;


List of Database Incarnations
DB Key  Inc Key DB Name  DB ID            STATUS  Reset SCN  Reset Time
------- ------- -------- ---------------- --- ---------- ----------
2       2       ENMO     850702656        PARENT  6835695    2019-09-20 18:11:28 
1       1       ENMO     850702656        PARENT  16566762   2020-02-26 12:54:08 
3       3       ENMO     850702656        CURRENT 18504056   2020-03-13 13:24:09  --从18504056开始新的数据库化身
```