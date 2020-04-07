---
title:  "view tablespaces"
published: true
summary: "查看表空间"
categories: 
  - dbascript
tags: 
  - tablespaces
---

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

```