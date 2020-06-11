---
title:  "UNDO Tablespace"
published: true
summary: "UNDO Tablespace"
read_time: false
comments: false
related: false
author_profile: false
categories: 
  - dbascript
tags: 
  - UNDO
  - tablespace
  - oracle
---

```
lsnrctl status     //查看Service name

sqlplus sys/oracle@cdb1 as sysdba    //登录cdb

col name format a16
col pdb format a16
select name,pdb from v$services;

alter session set container=pdb_easyee;
```

```
SET LINESIZE 200
SET PAGESIZE 999
// 查看当前数据所有表空间
SELECT   a.tablespace_name,
         ROUND (a.total_size) "total_size(MB)",
         ROUND (a.total_size) - ROUND (b.free_size, 3) "used_size(MB)",
         ROUND (b.free_size, 3) "free_size(MB)",
         ROUND (b.free_size / total_size * 100, 2) || '%' free_rate
  FROM   (  SELECT   tablespace_name, SUM (bytes) / 1024 / 1024 total_size
              FROM   dba_data_files
          GROUP BY   tablespace_name) a,
         (  SELECT   tablespace_name, SUM (bytes) / 1024 / 1024 free_size
              FROM   dba_free_space
          GROUP BY   tablespace_name) b
 WHERE   a.tablespace_name = b.tablespace_name(+);
// 查看表空间文件
COL file_name          FORMAT A60
COL TSIZE              FORMAT 99999.9999
select file_name,bytes/1024/1024 AS TSIZE from dba_data_files where tablespace_name like 'UNDOTBS1';
// 增加UNDO表空间文件
ALTER TABLESPACE UNDOTBS1 ADD DATAFILE '/u01/app/oracle/oradata/CDB1/pdb_easyee/undotbs01_2.dbf' SIZE 10M;
// 查看过期

SELECT   tablespace_name, 
         status, 
         SUM (bytes) / 1024 / 1024 "Bytes(M)"
  FROM   dba_undo_extents
GROUP BY   tablespace_name, status;
```