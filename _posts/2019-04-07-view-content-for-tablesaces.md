---
title:  "view content for tablespaces"
published: true
summary: "查看表空间存储内容"
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
SET PAGESIZE 99
col TABLESPACE_NAME  format a16
col SEGMENT_NAME format a20
col SEGMENT_SIZE format 9999.9999

SELECT TABLESPACE_NAME       AS TABLESPACE_NAME
     , SEGMENT_NAME          AS SEGMENT_NAME
     , SUM(BYTES)/1024/1024  AS SEGMENT_SIZE 
FROM DBA_SEGMENTS
WHERE TABLESPACE_NAME='PDBEASYEE_DATA'
GROUP BY TABLESPACE_NAME,SEGMENT_NAME
ORDER BY 3;

col OWNER format a16
col SEGMENT_NAME format a26
col SEGMENT_TYPE format a12
col SEGMENT_SIZE format 9999.9999

SELECT OWNER                  AS OWNER
      ,SEGMENT_NAME           AS SEGMENT_NAME
      ,SEGMENT_TYPE           AS SEGMENT_TYPE
      ,SUM(BYTES)/1024/1024   AS SEGMENT_SIZE
FROM DBA_SEGMENTS
WHERE TABLESPACE_NAME='PDBEASYEE_DATA'
GROUP BY OWNER,SEGMENT_NAME,SEGMENT_TYPE
ORDER BY 4;
```