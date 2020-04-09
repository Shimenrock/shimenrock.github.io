---
title:  "view user"
published: true
summary: "查看数据库中的用户"
read_time: false
comments: false
related: false
author_profile: false
categories: 
  - dbascript
tags: 
  - users
  - oracle
---

# 查看数据库中的用户

```
lsnrctl status     //查看Service name

sqlplus sys/oracle@cdb1 as sysdba    //登录cdb

col name format a16
col pdb format a16
select name,pdb from v$services;

alter session set container=pdb_easyee;
```

```
SQL> desc dba_users;
 Name                                      Null?    Type
 ----------------------------------------- -------- ----------------------------
 USERNAME                                  NOT NULL VARCHAR2(128)
 USER_ID                                   NOT NULL NUMBER
 PASSWORD                                           VARCHAR2(4000)
 ACCOUNT_STATUS                            NOT NULL VARCHAR2(32)
 LOCK_DATE                                          DATE
 EXPIRY_DATE                                        DATE
 DEFAULT_TABLESPACE                        NOT NULL VARCHAR2(30)
 TEMPORARY_TABLESPACE                      NOT NULL VARCHAR2(30)
 LOCAL_TEMP_TABLESPACE                              VARCHAR2(30)
 CREATED                                   NOT NULL DATE
 PROFILE                                   NOT NULL VARCHAR2(128)
 INITIAL_RSRC_CONSUMER_GROUP                        VARCHAR2(128)
 EXTERNAL_NAME                                      VARCHAR2(4000)
 PASSWORD_VERSIONS                                  VARCHAR2(17)
 EDITIONS_ENABLED                                   VARCHAR2(1)
 AUTHENTICATION_TYPE                                VARCHAR2(8)
 PROXY_ONLY_CONNECT                                 VARCHAR2(1)
 COMMON                                             VARCHAR2(3)
 LAST_LOGIN                                         TIMESTAMP(9) WITH TIME ZONE
 ORACLE_MAINTAINED                                  VARCHAR2(1)
 INHERITED                                          VARCHAR2(3)
 DEFAULT_COLLATION                                  VARCHAR2(100)
 IMPLICIT                                           VARCHAR2(3)
 ALL_SHARD                                          VARCHAR2(3)
 PASSWORD_CHANGE_DATE                               DATE

SPOOL /tmp/oracle_users.txt

SET LINESIZE 200
SET PAGESIZE 999
COLUMN  USERNAME             FORMAT a30                         HEADING 'NAME'
COLUMN  ACCOUNT_STATUS       FORMAT a20                         HEADING 'STATUS'
COLUMN  DEFAULT_TABLESPACE   FORMAT a20                         HEADING 'TABLESPACE'
 
ALTER SESSION SET NLS_DATE_FORMAT = 'YYYY-MM-DD';
select USERNAME,ACCOUNT_STATUS,DEFAULT_TABLESPACE,TO_CHAR(LAST_LOGIN, 'MM/DD/YYYY') from dba_users;

select * from dba_sys_privs where grantee='EASYEE' ;
SPOOL OFF;
```