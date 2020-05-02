---
title:  "whomi"
published: true
summary: "whomi"
read_time: false
comments: false
related: false
author_profile: false
categories: 
  - dbascript
tags: 
  - userenv
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
set serveroutput on 
begin 
    dbms_output.put_line('USER             :'||sys_context('userenv','session_user')); 
    dbms_output.put_line('SESSION ID       :'||sys_context('userenv','sid')); 
    dbms_output.put_line('CURRENT_SCHEMA   :'||sys_context('userenv','current_schema')); 
    dbms_output.put_line('INSTANCE NAME    :'||sys_context('userenv','instance_name')); 
--12c    dbms_output.put_line('CDB NAME         :'||sys_context('userenv','cbd_name')); 
--12c    dbms_output.put_line('CONTAINER NAME   :'||sys_context('userenv','con_name')); 
--12c    dbms_output.put_line('DATABASE ROLE    :'||sys_context('userenv','database_role')); 
--12c    dbms_output.put_line('OS USER          :'||sys_context('userenv','os_user')); 
    dbms_output.put_line('CLIENT IP ADDRESS:'||sys_context('userenv','ip_address')); 
    dbms_output.put_line('SERVER HOSTNAME  :'||sys_context('userenv','server_host')); 
    dbms_output.put_line('CLIENT HOSTNAME  :'||sys_context('userenv','host'));
end; 
/

SET LINESIZE 200
SET PAGESIZE 999
COL OWNER  FORMAT A5
COL NAME   FORMAT A8
COL TYPE   FORMAT A12
COL LINE   FORMAT 9999
COL TEXT   FORMAT A110
COL ORIGIN_CON_ID  FORMAT 99
select OWNER,TEXT from dba_source vs
    where 1=1
        and upper(vs.text) like '%SYS_CONTEXT%'
        and vs.name = 'STANDARD'
    order by vs.line;

SET LINESIZE 120
SET PAGESIZE 999
COL terminal  					FORMAT A12
COL language  					FORMAT A12
COL sessionid 					FORMAT A12
COL instance  					FORMAT A12
COL entryid   					FORMAT A12
COL isdba     					FORMAT A12
COL nls_territory       FORMAT A12
COL nls_currency        FORMAT A12
COL nls_calendar        FORMAT A12
COL nls_date_format     FORMAT A12
COL nls_date_language   FORMAT A12
COL NLS_SORT            FORMAT A12
COL current_user        FORMAT A12
COL current_userid      FORMAT A12
COL session_user        FORMAT A12
COL session_userid      FORMAT A12
COL proxy_user          FORMAT A12
COL proxy_userid        FORMAT A12
COL db_domain       		FORMAT A12
COL db_name         		FORMAT A12
COL host            		FORMAT A12
COL os_user       			FORMAT A12
COL external_name   		FORMAT A12
COL ip_address          FORMAT A12
COL network_protocol    FORMAT A12
COL bg_job_id           FORMAT A12
COL fg_job_id           FORMAT A12
select 
	SYS_CONTEXT('USERENV','TERMINAL') terminal, 
	SYS_CONTEXT('USERENV','LANGUAGE') language, 
	SYS_CONTEXT('USERENV','SESSIONID') sessionid, 
	SYS_CONTEXT('USERENV','INSTANCE') instance, 
	SYS_CONTEXT('USERENV','ENTRYID') entryid, 
	SYS_CONTEXT('USERENV','ISDBA') isdba, 
	SYS_CONTEXT('USERENV','NLS_TERRITORY') nls_territory, 
	SYS_CONTEXT('USERENV','NLS_CURRENCY') nls_currency,  
    SYS_CONTEXT('USERENV','NLS_CALENDAR') nls_calendar, 
    SYS_CONTEXT('USERENV','NLS_DATE_FORMAT') nls_date_format,  
    SYS_CONTEXT('USERENV','NLS_DATE_LANGUAGE') nls_date_language, 
    SYS_CONTEXT('USERENV','NLS_SORT') nls_sort,
	SYS_CONTEXT('USERENV','CURRENT_USER') current_user, 
	SYS_CONTEXT('USERENV','CURRENT_USERID') current_userid, 
	SYS_CONTEXT('USERENV','SESSION_USER') session_user, 
	SYS_CONTEXT('USERENV','SESSION_USERID') session_userid, 
	SYS_CONTEXT('USERENV','PROXY_USER') proxy_user, 
	SYS_CONTEXT('USERENV','PROXY_USERID') proxy_userid, 
	SYS_CONTEXT('USERENV','DB_DOMAIN') db_domain, 
	SYS_CONTEXT('USERENV','DB_NAME') db_name, 
	SYS_CONTEXT('USERENV','HOST') host, 
	SYS_CONTEXT('USERENV','OS_USER') os_user, 
	SYS_CONTEXT('USERENV','EXTERNAL_NAME') external_name, 
	SYS_CONTEXT('USERENV','IP_ADDRESS') ip_address,
	SYS_CONTEXT('USERENV','NETWORK_PROTOCOL') network_protocol,  
	SYS_CONTEXT('USERENV','BG_JOB_ID') bg_job_id, 
	SYS_CONTEXT('USERENV','FG_JOB_ID') fg_job_id   
FROM DUAL;


```