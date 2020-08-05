---
title:  "Query time-consuming statement"
published: true
summary: "Query time-consuming statement"
read_time: false
comments: false
related: false
author_profile: false
categories: 
  - dbascript
tags: 
  - statement
  - time-consuming
  - session_longops
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
# 创建一亿条数据
SET LINESIZE 200
SET PAGESIZE 999
create table students
(
    ID int,
    userName varchar(100),
    userPass varchar(100),
    userAge int
)
/

create or replace Procedure Test2
as
num NUMBER;
v_begintime varchar2(20);
v_endtime varchar2(20);
v_str varchar2(10);
begin
v_begintime:=to_char(sysdate,'yyyy-mm-dd hh24:mi:ss');
 FOR i IN 1..10000000 LOOP
       insert into students values(i,'jack','jjjaa',23);
       num:=num+1;
        IF Mod(num,5000)=0 THEN 
               COMMIT; 
        END IF;
 END LOOP;
 v_endtime:=to_char(sysdate,'yyyy-mm-dd hh24:mi:ss');
 dbms_output.put_line('startTime:'||v_begintime);
 dbms_output.put_line('endTime:'||v_endtime);
end Test2;
/

exec Test2

SQL> select count(*) from students;

  COUNT(*)
----------
  10000000
```


```
select * from students where username='jack';

select b.*
 from v$session a, v$session_longops b
 where a.sid=b.sid
 and a.serial#=b.serial#
 /
 
 
OPNAME：		    指长时间执行的操作名.如：Table Scan
TARGET：		    被操作的object_name. 如：tableA 
TARGET_DESC：	    描述target的内容 
SOFAR：			    这个是需要着重去关注的，表示已要完成的工作数，如扫描了多少个块。
TOTALWORK：		    指目标对象一共有多少数量（预计）。如块的数量。
UNITS： 
START_TIME：	    进程的开始时间
LAST_UPDATE_TIM：   最后一次调用set_session_longops的时间
TIME_REMAINING：    估计还需要多少时间完成，单位为秒
ELAPSED_SECONDS：   指从开始操作时间到最后更新时间
CONTEXT：
MESSAGE：			对于操作的完整描述，包括进度和操作内容。 
USERNAME：		    与v$session中的一样。
SQL_ADDRESS：	    关联v$sql
SQL_HASH_VALUE：    关联v$sql
QCSID：				主要是并行查询一起使用。


SET LINE 9999  PAGESIZE 9999
COL USERNAME FORMAT A10 
COL SESSION_INFO FORMAT A30
COL TARGET FORMAT A20
COL OPNAME FORMAT A35 
COL MESSAGE FORMAT A80 
COL SOFAR_TOTALWORK FORMAT A20 
COL PROGRESS FORMAT A8

SELECT A.USERNAME, 
       (SELECT NB.SID || ',' || NB.SERIAL# || ',' || PR.SPID || ',' ||NB.OSUSER|| ',' ||NB.STATUS|| ',' ||NB.EVENT
          FROM GV$PROCESS PR, GV$SESSION NB
         WHERE NB.PADDR = PR.ADDR
           AND NB.SID = A.SID
           AND NB.SERIAL# = A.SERIAL#
           AND PR.INST_ID = NB.INST_ID) SESSION_INFO,
       A.TARGET,
       A.OPNAME,
       TO_CHAR(A.START_TIME, 'YYYY-MM-DD HH24:MI:SS') START_TIME,
       ROUND(A.SOFAR * 100 / A.TOTALWORK, 2) || '%' AS PROGRESS,
       (A.SOFAR || ':' || A.TOTALWORK) SOFAR_TOTALWORK,
       A.TIME_REMAINING TIME_REMAINING,
       A.ELAPSED_SECONDS ELAPSED_SECONDS,
       MESSAGE MESSAGE
  FROM GV$SESSION_LONGOPS A
 WHERE A.TIME_REMAINING <> 0
 ORDER BY  A.TIME_REMAINING DESC, A.SQL_ID, A.SID;
```