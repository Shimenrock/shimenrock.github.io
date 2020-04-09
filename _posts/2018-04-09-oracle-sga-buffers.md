---
title:  "view sga buffers"
published: true
summary: "Displays the status of buffers in the SGA"
read_time: false
comments: false
related: false
author_profile: false
categories: 
  - dbascript
tags: 
  - sga
  - oracle
  - buffer
---

# Displays the status of buffers in the SGA

```
lsnrctl status     //查看Service name

sqlplus sys/oracle@cdb1 as sysdba    //登录cdb

col name format a16
col pdb format a16
select name,pdb from v$services;

alter session set container=pdb_easyee;
```

```
SPOOL /tmp/oracle_sga_buffers.txt

SET PAGESIZE 999
SET LINESIZE 200
COLUMN TABLESPACE_NAME      FORMAT A18  
COLUMN TABLESPACE_NAME      FORMAT A18  
COLUMN object_name          FORMAT A30
COLUMN FREE                 FORMAT 9999
COLUMN XCUR                 FORMAT 9999
COLUMN SCUR                 FORMAT 9999
COLUMN CR                   FORMAT 9999
COLUMN READREC              FORMAT 9999
SELECT t.name AS tablespace_name,
       o.object_name,
       SUM(DECODE(bh.status, 'free', 1, 0)) AS free,
       SUM(DECODE(bh.status, 'xcur', 1, 0)) AS xcur,
       SUM(DECODE(bh.status, 'scur', 1, 0)) AS scur,
       SUM(DECODE(bh.status, 'cr', 1, 0)) AS cr,
       SUM(DECODE(bh.status, 'read', 1, 0)) AS read,
       SUM(DECODE(bh.status, 'mrec', 1, 0)) AS mrec,
       SUM(DECODE(bh.status, 'irec', 1, 0)) AS irec
FROM   v$bh bh
       JOIN dba_objects o ON o.object_id = bh.objd
       JOIN v$tablespace t ON t.ts# = bh.ts#
GROUP BY t.name, o.object_name;

SPOOL OFF;
```