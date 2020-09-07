---
title:  "Flashback table"
published: true
summary: "Flashback table"
read_time: false
comments: false
related: false
toc: true
toc_sticky: true
author_profile: false
categories: 
  - dbascript
tags: 
  - flashback
  - oracle
---

## 1.登陆数据库

``` 
lsnrctl status     //查看Service name

sqlplus sys/oracle@cdb1 as sysdba    //登录cdb

col name format a16
col pdb format a16
select name,pdb from v$services;

alter session set container=pdb_easyee;
```

## 2.开启闪回

``` 
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

## 3.创建测试表

``` 
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

## 4.闪回表

```
SET LINESIZE 200
SET PAGESIZE 999

COL BANNER FORMAT A100
COL BANNER_FULL FORMAT A100
COL BANNER_LEGACY FORMAT A100
select * from v$version;

select  re_id,
        re_date,
        re_scn,
        re_start,
        re_note  
        from RESTORE_TIME order by re_id desc;


delete from RESTORE_TIME where re_id=2;
select to_char(sysdate,'yyyy-mm-dd hh24:mi:ss') from dual;
SELECT current_scn, TO_CHAR(SYSTIMESTAMP, 'YYYY-MM-DD HH24:MI:SS') FROM v$database;
commit;

2020-07-14 05:08:34
CURRENT_SCN TO_CHAR(SYSTIMESTAM
----------- -------------------
    2503844 2020-07-14 05:09:26

-- sysdate-10/1440 十分钟之前
select * from RESTORE_TIME as of timestamp sysdate-10/1440;

select * from RESTORE_TIME as of timestamp to_timestamp('2020-07-14 05:08:34','yyyy-mm-dd hh24:mi:ss');

select * from RESTORE_TIME as of scn 2500000;


flashback table RESTORE_TIME to timestamp to_timestamp('2020-07-14 05:08:34','yyyy-mm-dd hh24:mi:ss');

select row_movement from dba_tables where table_name='RESTORE_TIME' and owner='TEST';

alter table RESTORE_TIME enable row movement;
 
select * from RESTORE_TIME;
alter table RESTORE_TIME disable row movement;

```

## 5. 闪回drop

```
 drop table RESTORE_TIME;

 show recyclebin;
 
SQL> show recyclebin;
ORIGINAL NAME    RECYCLEBIN NAME                OBJECT TYPE  DROP TIME
---------------- ------------------------------ ------------ -------------------
RESTORE_TIME     BIN$qmRXDSK2DFPgU9QLqMC3yg==$0 TABLE        2020-07-14:05:22:02

 flashback table RESTORE_TIME to before drop;

 select * from RESTORE_TIME;

flashback table t to before drop rename to tt;
```
## 6. 闪回注意事项

- 数据库闪回需要在mounted下进行，并且open时需要使用resetlogs
- 闪回DROP只能用于非系统表空间和本地管理的表空间，外键约束无法恢复，对方覆盖、重命名需注意
- 表DROP，对应的物化视图会被彻底删除，物化视图不会存放在recyclebin里
- 闪回表，如果在做过dml，然后进行了表结构修改、truncate等DDL操作，新增/删除结构无法做闪回
- 闪回归档，必须在assm管理tablespace和undo auto管理下进行
- 注意闪回区管理，防止磁盘爆满，闪回区空间不足等
- 主库做库的闪回，会影响备库，需要重新同步
- snapshot standby 不支持最高保护模式
 
## 6.相关查询

V$FLASHBACK_DATABASE_LOG  ##查看数据库可闪回的时间点/SCN等信息

V$flashback_database_stat ##查看闪回日志空间记录信息

1. 查看数据库状态

```
select NAME,OPEN_MODE ,DATABASE_ROLE,CURRENT_SCN,FLASHBACK_ON from v$database;
```

2. 获取当前数据库的系统时间和SCN

```
select to_char(systimestamp,'yyyy-mm-dd HH24:MI:SS') as sysdt , dbms_flashback.get_system_change_number scn from dual;
```

3. 查看数据库可恢复的时间点

```
select * from V$FLASHBACK_DATABASE_LOG;
```

4. 查看闪回日志空间情况

```
select * from V$flashback_database_stat;
```

5. SCN和timestamp装换关系查询

```
select scn,to_char(time_dp,'yyyy-mm-dd hh24:mi:ss')from sys.smon_scn_time;
```

6. 查看闪回restore_point

```
select scn, STORAGE_SIZE ,to_char(time,'yyyy-mm-dd hh24:mi:ss') time,NAME from v$restore_point;
```

## 7.通过时间戳查询历史数据

```
set time on

alter session set nls_date_format='yyyy-mm-dd hh24:mi:ss';

select sysdate from dual;

22:49:28 SQL> create table demo(id primary key, text) as select rownum,to_char(rownum,'099999999') from xmltable('1 to 5');

Table created.

22:49:36 SQL>  select * from demo;

        ID TEXT
---------- ----------
         1  000000001
         2  000000002
         3  000000003
         4  000000004
         5  000000005
```
### update the PK

```
22:49:53 SQL> update demo set id=id+1;

5 rows updated.

22:50:09 SQL> commit;

SQL> select /*+ gather_plan_statistics */ * from demo where id=3;

        ID TEXT
---------- ----------
         3  000000002
```

### QUERY as of before the update

```
22:58:36 SQL> select /*+ gather_plan_statistics */ * from demo as of timestamp (timestamp '2020-09-06 22:49:50') where id=3;

        ID TEXT
---------- ----------
         3  000000003        


select * from demo as of timestamp to_timestamp('2020-09-06 22:49:50','yyyy-mm-dd hh24:mi:ss');

set linesize 300
select plan_table_output from dbms_xplan.display_cursor(format=>'allstats last +cost ');
```

### can still use the index acces to the old value

```
23:04:10 SQL> select plan_table_output from dbms_xplan.display_cursor(format=>'allstats last +cost ');

PLAN_TABLE_OUTPUT
------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
SQL_ID  bxx2nh2frvvtw, child number 2
-------------------------------------
select /*+ gather_plan_statistics */ * from demo as of timestamp
(timestamp '2020-09-06 22:49:50') where id=3

Plan hash value: 776694569

------------------------------------------------------------------------------------------------------------------
| Id  | Operation                   | Name        | Starts | E-Rows | Cost (%CPU)| A-Rows |   A-Time   | Buffers |
------------------------------------------------------------------------------------------------------------------
|   0 | SELECT STATEMENT            |             |      1 |        |     1 (100)|      1 |00:00:00.01 |       6 |

PLAN_TABLE_OUTPUT
------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
|   1 |  TABLE ACCESS BY INDEX ROWID| DEMO        |      1 |      1 |     1   (0)|      1 |00:00:00.01 |       6 |
|*  2 |   INDEX UNIQUE SCAN         | SYS_C007534 |      1 |      1 |     0   (0)|      1 |00:00:00.01 |       3 |
------------------------------------------------------------------------------------------------------------------

Predicate Information (identified by operation id):
---------------------------------------------------

   2 - access("ID"=3)


20 rows selected.
```