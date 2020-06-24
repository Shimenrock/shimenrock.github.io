---
title:  "view tablespaces"
published: true
summary: "查看表空间"
read_time: false
comments: false
related: false
author_profile: false
categories: 
  - dbascript
tags: 
  - tablespaces
  - oracle
---

# 查看ORACLE数据库里的表空间

```
lsnrctl status     //查看Service name

sqlplus sys/oracle@cdb1 as sysdba    //登录cdb

col name format a16
col pdb format a16
select name,pdb from v$services;

alter session set container=pdb_easyee;

SET LINESIZE 200

SELECT tablespace_name,
       block_size,
       extent_management,
       allocation_type,
       segment_space_management,
       status
FROM   dba_tablespaces
ORDER BY tablespace_name;

SET LINESIZE 200
SET PAGESIZE 100

SELECT UPPER(F.TABLESPACE_NAME) "tablespace_name",
       D.TOT_GROOTTE_MB "SIZE(M)",
       D.TOT_GROOTTE_MB - F.TOTAL_BYTES "USED(M)",
       TO_CHAR(ROUND((D.TOT_GROOTTE_MB - F.TOTAL_BYTES) / D.TOT_GROOTTE_MB * 100,
                     2),
               '990.99') "USED%",
       F.TOTAL_BYTES "FREE(M)",
       F.MAX_BYTES "MAX(M)"
  FROM (SELECT TABLESPACE_NAME,
               ROUND(SUM(BYTES) / (1024 * 1024), 2) TOTAL_BYTES,
               ROUND(MAX(BYTES) / (1024 * 1024), 2) MAX_BYTES
          FROM SYS.DBA_FREE_SPACE
         GROUP BY TABLESPACE_NAME) F,
       (SELECT DD.TABLESPACE_NAME,
               ROUND(SUM(DD.BYTES) / (1024 * 1024), 2) TOT_GROOTTE_MB
          FROM SYS.DBA_DATA_FILES DD
         GROUP BY DD.TABLESPACE_NAME) D
 WHERE D.TABLESPACE_NAME = F.TABLESPACE_NAME
 ORDER BY F.TABLESPACE_NAME;

select  FILE_ID,
        TABLESPACE_NAME,
        ROUND(BYTES/1024/1024,2) MB,
        STATUS,
        FILE_NAME 
        from dba_data_files;

//增加表空间例子
alter tablespace "tp_base" add datafile '+DATA/racdb/datafile/tp_base_07.dbf' size 10g autoextend on next 100m maxsize 30g
```