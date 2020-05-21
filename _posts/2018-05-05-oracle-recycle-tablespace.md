---
title:  "Recycle Tablespace"
published: true
summary: "Recycle Tablespace"
read_time: false
comments: false
related: false
author_profile: false
categories: 
  - dbascript
tags: 
  - recycle
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

// 检查表空间内容
SELECT tablespace_name,
       block_size,
       extent_management,
       allocation_type,
       segment_space_management,
       status
FROM   dba_tablespaces
ORDER BY tablespace_name;

COL TABLESPACE_NAME 	     FORMAT A20
COL SEGMENT_NAME 			 FORMAT A20
COL EGMENT_SIZE		         FORMAT 99999.9999
SELECT TABLESPACE_NAME       AS TABLESPACE_NAME
     , SEGMENT_NAME          AS SEGMENT_NAME
     , SUM(BYTES)/1024/1024  AS SEGMENT_SIZE 
FROM DBA_SEGMENTS
WHERE TABLESPACE_NAME='PDBEASYEE_DATA'
GROUP BY TABLESPACE_NAME,SEGMENT_NAME
ORDER BY 3;

COL OWNER                    FORMAT A20
COL SEGMENT_NAME     	     FORMAT A20
COL SEGMENT_TYPE 			 FORMAT A20
COL SEGMENT_SIZE		     FORMAT 99999.9999
SELECT OWNER                  AS OWNER
      ,SEGMENT_NAME           AS SEGMENT_NAME
      ,SEGMENT_TYPE           AS SEGMENT_TYPE
      ,SUM(BYTES)/1024/1024   AS SEGMENT_SIZE
FROM DBA_SEGMENTS
WHERE TABLESPACE_NAME='PDBEASYEE_DATA'
GROUP BY OWNER,SEGMENT_NAME,SEGMENT_TYPE
ORDER BY 4;

// 1.先清理回收站
purge DBA_RECYCLEBIN

// 2.-查询表空间文件可以回收的大小
SET LINESIZE 200
COL FILE                   FORMAT A4
COL NAME                   FORMAT A60
COL CurrentMB              FORMAT 999999
COL ResizeTo               FORMAT 999999.9999
COL ReleaseMB              FORMAT 999999.9999
COL ResizeCMD              FORMAT A120
select   a.file#                                    
        ,a.name                                     AS NAME
        ,a.bytes/1024/1024                          AS CurrentMB
        ,ceil(HWM * a.block_size)/1024/1024         AS ResizeTo
        ,(a.bytes - HWM * a.block_size)/1024/1024   AS ReleaseMB
        ,'alter database datafile '''||a.name||''' resize '||
         ceil(HWM * a.block_size/1024/1024) || 'M;' AS ResizeCMD
from v$datafile a,
       (select file_id,max(block_id+blocks-1)       HWM
        from dba_extents
        group by file_id) b
where a.file# = b.file_id(+)
and (a.bytes - HWM *block_size)>0
order by 5;
------------
 

//1.选择某个表空间中超过N个blocks的segments，通过此语句可以看出那个表占用的空间大。

select   segment_name
        ,segment_type
        ,blocks 
from dba_segments
where tablespace_name='PDBEASYEE_DATA'
and blocks > 1
order by blocks;

//1.分析表，得知表的一些信息

analyze table TABLENAME estimate statistics;
 
//2.执行完后再执行

select   initial_extent
        ,next_extent
        ,min_extents
        ,blocks
        ,empty_blocks 
from dba_tables
where table_name='HISHOLDSINFO' and owner='hs_his';

//3.使用alter table ... deallocate unused 命令回收表的空间

 alter table hs_his.HISHOLDSINFO' deallocate unused keep 1k;
```