---
title:  "Flashback table"
published: true
summary: "Flashback table"
read_time: false
comments: false
related: false
author_profile: false
categories: 
  - dbascript
tags: 
  - flashback
  - oracle
---

``` 登陆数据库
lsnrctl status     //查看Service name

sqlplus sys/oracle@cdb1 as sysdba    //登录cdb

col name format a16
col pdb format a16
select name,pdb from v$services;

alter session set container=pdb_easyee;
```
``` 开启闪回
//查看是否开启闪回

SQL> select flashback_on from v$database;

FLASHBACK_ON
------------------
NO

//查看是否配置了db_recover_file_dest

SQL> show parameter db_recovery

NAME                                 TYPE        VALUE
------------------------------------ ----------- ------------------------------
db_recovery_file_dest                string
db_recovery_file_dest_size           big integer 0

//创建闪回目录

$cd $ORACLE_BASE
$mkdir flashback

SQL> alter system set db_recovery_file_dest_size=8G scope=both;
alter system set db_recovery_file_dest_size=8G scope=both
*
ERROR at line 1:
ORA-65040: operation not allowed from within a pluggable database

SQL> alter session set container=cdb$root;

Session altered.

SQL> show con_name 

CON_NAME
------------------------------
CDB$ROOT

SQL> alter system set db_recovery_file_dest_size=8G scope=both;
SQL> alter system set db_recovery_file_dest='/u01/app/oracle/flashback'  scope=both;

SQL> shutdown immediate;

$export ORACLE_SID=cdb1
$echo $ORACLE_SID
cdb1
$sqlplus / as sysdba

SQL*Plus: Release 19.0.0.0.0 - Production on Tue Jun 23 01:31:51 2020
Version 19.3.0.0.0

Copyright (c) 1982, 2019, Oracle.  All rights reserved.

Connected to an idle instance.

SQL> startup mount;

SQL> show con_name

CON_NAME
------------------------------
CDB$ROOT

COL CON_ID FORMAT 99
COL DBID FORMAT 9999999999
COL GUID FORMAT A35
COL NAME FORMAT A10
COL OPEN_MODE FORMAT A10
select con_id, dbid, guid, name , open_mode from v$pdbs;
CON_ID        DBID GUID                                NAME       OPEN_MODE
------ ----------- ----------------------------------- ---------- ----------
     2  3821341309 A24B236901A414BEE053D40BA8C053F9    PDB$SEED   MOUNTED
     3   714194978 A24B3A2ECA10193DE053D40BA8C04BF1    PDB_EASYEE MOUNTED

SQL> alter database archivelog;

Database altered.

SQL> alter database flashback on;

Database altered.

SQL> alter database open;

Database altered.

SQL> select flashback_on from v$database;

FLASHBACK_ON
------------------
YES
```
``` 创建测试表
alter session set container=pdb_easyee;

-- 创建用户 
create user test
  identified by "test"
  default tablespace USERS
  temporary tablespace TEMP
  profile DEFAULT;

-- 赋权
grant connect to test;
grant resource to test;
grant execute on dbms_flashback to test;
grant select on sys.v_$instance to test;
ALTER USER test QUOTA UNLIMITED ON USERS;

COL USERNAME FORMAT A30
select USERNAME,ACCOUNT_STATUS,LOCK_DATE  from dba_users WHERE USERNAME='TEST';
alter user TEST identified by "test_123";

$ vim $ORACLE_HOME/network/admin/tnsname.ora
EASYEE =
  (DESCRIPTION =
    (ADDRESS_LIST =
      (ADDRESS = (PROTOCOL = TCP)(HOST = 192.168.11.212)(PORT = 1521))
    )
    (CONNECT_DATA =
      (SERVICE_NAME = pdb_easyee.oracle.com)
    )
  )
$sqlplus test/"test_123"@EASYEE
-- 创建表
 
create table RESTORE_TIME
(
  re_id   NUMBER not null,
  re_date DATE,
  re_scn VARCHAR2(20),
  re_start VARCHAR2(20),
  re_note VARCHAR2(20)
);
--创建序列
create sequence RESTORE_TIME1
start with 1
increment by 1
minvalue 1
maxvalue 9999999
nocycle
nocache
noorder;
--创建触发器
create or replace trigger RESTORE_TIME_trigger
before insert on RESTORE_TIME
for each row WHEN (new.re_id is null) 
begin
    select RESTORE_TIME1.nextval into:new.re_id from sys.dual;
end;
/
alter trigger  RESTORE_TIME_trigger enable;
--创建存储过程
create or replace procedure pro_RESTORE_TIME is
begin
 insert  into RESTORE_TIME (re_date,re_scn,re_start)
select sysdate,dbms_flashback.get_system_change_number,to_char(startup_time,'yyyy-mm-dd hh24:mi:ss') from v$instance;
commit;
end pro_RESTORE_TIME;
/
--测试存储过程
begin
  -- Call the procedure
  pro_RESTORE_TIME;
end;
/
--创建定时任务
set linesize 200
select  re_id,
        re_date,
        re_scn,
        re_start,
        re_note  
        from RESTORE_TIME order by re_id desc;

```
```
SET LINESIZE 200
SET PAGESIZE 999

COL BANNER FORMAT A100
COL BANNER_FULL FORMAT A100
COL BANNER_LEGACY FORMAT A100
select * from v$version;
```