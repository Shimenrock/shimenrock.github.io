---
title:  "oracle time"
published: true
summary: "查看数据库时间"
read_time: false
comments: false
related: false
author_profile: false
categories: 
  - dbascript
tags: 
  - time
  - oracle
---

# 查看数据库时间

```
lsnrctl status     //查看Service name

sqlplus sys/oracle@cdb1 as sysdba    //登录cdb

col name format a16
col pdb format a16
select name,pdb from v$services;

alter session set container=pdb_easyee;

select to_char(sysdate,'yyyy-mm-dd hh24:mi:ss') time from dual;

select dbtimezone from dual;

select SESSIONTIMEZONE from dual;
```