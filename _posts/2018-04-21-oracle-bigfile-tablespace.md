---
title:  "Bigfile Tablespace"
published: true
summary: "Bigfile Tablespace"
read_time: false
comments: false
related: false
author_profile: false
categories: 
  - dbascript
tags: 
  - bigfile
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

COL PROPERTY_NAME  FORMAT A20
COL PROPERTY_VALUE FORMAT A20
COL DESCRIPTION    FORMAT A30
SELECT * FROM database_properties WHERE property_name = 'DEFAULT_TBS_TYPE';

ALTER DATABASE SET DEFAULT bigfile TABLESPACE;
ALTER DATABASE SET DEFAULT smallfile TABLESPACE;

CREATE bigfile tablespace bttest 
      datafile '/u01/app/oracle/oradata/CDB1/pdb_easyee/bttest.dbf' size 1024m autoextend on next 1g
      extent management local uniform size 128m
      segment space management auto;

SELECT tablespace_name, bigfile FROM dba_tablespaces;

col file_name format a60
SELECT  file_name, 
        file_id, 
        relative_fno 
FROM dba_data_files;


SELECT tablespace_name,
       block_size,
       extent_management,
       allocation_type,
       segment_space_management,
       status
FROM   dba_tablespaces
ORDER BY tablespace_name;


```