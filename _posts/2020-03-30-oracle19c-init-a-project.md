---
title: "Oracle Database 19c Initialize a project"
published: false
related: true
header:
  teaser: /assets/images/unsplash-image-3.jpg
toc: true
toc_sticky: true
categories: 
  - oracle
---  

## 19c下表空间管理

### CDB表空间查看

以下视图适用于PDB

```
set pagesize 999
set linesize 120
col TABLESPACE_NAME format a16
col files format a6
col NAME format a8
col FAILGROUP  format a12
col PATH  format a15

SELECT UPPER(F.TABLESPACE_NAME) "TABLESPACE_NAME",
         D.DATAFILE_COUNT "FILES",
         ROUND((D.TOT_GROOTTE_MB), 0) "SIZE(M)",
         ROUND((D.TOT_GROOTTE_MB - F.TOTAL_BYTES), 0) "USED(M)",
         ROUND((D.TOT_GROOTTE_MB - F.TOTAL_BYTES) / D.TOT_GROOTTE_MB * 100, 0) "USED(%)",
         ROUND((F.TOTAL_BYTES), 0) "FREE(M)",
         ROUND((D.TOT_MAX_MB), 0) "EXTENSIBLE(M)"
      FROM (SELECT FF.TABLESPACE_NAME,
                  ROUND(SUM(FF.BYTES) / (1024 * 1024), 2) TOTAL_BYTES,
                  ROUND(MAX(FF.BYTES) / (1024 * 1024), 2) MAX_BYTES
             FROM SYS.DBA_FREE_SPACE FF
            GROUP BY TABLESPACE_NAME) F,
          (SELECT DD.TABLESPACE_NAME,
                  COUNT(*) DATAFILE_COUNT,
                  ROUND(SUM(DD.BYTES) / (1024 * 1024), 2) TOT_GROOTTE_MB,
                  ROUND(SUM(DD.MAXBYTES) / (1024 * 1024), 2) TOT_MAX_MB
             FROM SYS.DBA_DATA_FILES DD
            GROUP BY DD.TABLESPACE_NAME) D
    WHERE D.TABLESPACE_NAME = F.TABLESPACE_NAME
    ORDER BY F.TABLESPACE_NAME;
TABLESPACE_NAME   FILES    SIZE(M)    USED(M)    USED(%)    FREE(M) EXTENSIBLE(M)
---------------- ------ ---------- ---------- ---------- ---------- -------------
SYSAUX                1        510        476         93         34         32768
SYSTEM                1        890        882         99          8         32768
UNDOTBS1              1         65         65         99          0         32768
USERS                 1          5          3         54          2         32768

select TABLESPACE_NAME,file_name from dba_data_files;
TABLESPACE_NAME  FILE_NAME
---------------- ---------------------------------------------
SYSTEM           /u01/app/oracle/oradata/CDB1/system01.dbf
SYSAUX           /u01/app/oracle/oradata/CDB1/sysaux01.dbf
UNDOTBS1         /u01/app/oracle/oradata/CDB1/undotbs01.dbf
USERS            /u01/app/oracle/oradata/CDB1/users01.dbf

col FILE_NAME format a45
select tablespace_name,file_name,bytes/1024/1024 file_size,autoextensible from dba_temp_files;
TABLESPACE_NAME  FILE_NAME                                      FILE_SIZE AUTOEXTENSIBLE
---------------- --------------------------------------------- ---------- --------------
TEMP             /u01/app/oracle/oradata/CDB1/temp01.dbf               32 YES

col NAME format a60
col con_id format a7
select con_id,status,enabled,name,bytes/1024/1024 file_size from v_$tempfile;
    CON_ID STATUS  ENABLED    NAME                                                          FILE_SIZE
---------- ------- ---------- ------------------------------------------------------------ ----------
         1 ONLINE  READ WRITE /u01/app/oracle/oradata/CDB1/temp01.dbf                              32
         2 ONLINE  READ WRITE /u01/app/oracle/oradata/CDB1/pdbseed/temp012020-04-02_03-56-         36
                              28-909-AM.dbf                                                
         3 ONLINE  READ WRITE /u01/app/oracle/oradata/CDB1/pdb_easyee/temp01.dbf                   36

column name format a10
select name, con_id, dbid, con_uid, guid from v$containers order by con_id;
NAME        CON_ID       DBID    CON_UID GUID
---------- ------- ---------- ---------- --------------------------------
CDB$ROOT         1 1010358776          1 86B637B62FDF7A65E053F706E80A27CA
PDB$SEED         2 3821341309 3821341309 A24B236901A414BEE053D40BA8C053F9
PDB_EASYEE       3  714194978  714194978 A24B3A2ECA10193DE053D40BA8C04BF1
```

### PDB下创建表空间

```
create temporary tablespace pdbeasyee_temp
tempfile '/u01/app/oracle/oradata/CDB1/pdb_easyee/pdbeasyee_temp01.dbf'
size 32m
autoextend on
next 32m MAXSIZE unlimited  
extent management local;

create tablespace pdbeasyee_data
logging
datafile '/u01/app/oracle/oradata/CDB1/pdb_easyee/pdbeasyee_data01.dbf'
size 1024m
autoextend on
next 100m MAXSIZE unlimited  
extent management local;
```

### 创建业务用户  
```
-- Create the user 
create user easyee
  identified by easyee2020
  default tablespace PDBEASYEE_DATA
  temporary tablespace PDBEASYEE_TEMP
  profile DEFAULT;
-- Grant/Revoke role privileges 
grant connect to easyee;
grant resource to easyee;
-- Grant/Revoke system privileges 
grant unlimited tablespace to easyee;

```
### 登陆业务账户初始化数据结构






easyee
set linesize 300
set pagesize 300
col username format a30
col account_status format a30
col lock_date format a30
select username,account_status,lock_date from dba_users;

alter  user system identified by sjzrsj_2014;

https://blog.csdn.net/feiyanaffection/article/details/88394589
http://blog.itpub.net/26175573/viewspace-2122295/


alter session set container=pdb_easyee;

 column name format a8
SQL>  select name, con_id, dbid, con_uid, guid from v$containers order by con_id;