---
title: "Mysql 5.6.46 二进制安装"
permalink: /dba/install-mysql57/
excerpt: "Mysql 5.6.46 二进制安装"
last_modified_at: 2020-02-08T21:36:11-04:00
categories: mysql
redirect_from:
  - /theme-setup/
toc: true
---

# 1.官网

https://downloads.mysql.com

mysql 5.6.46   394 MB   没有sys数据库
```
# vim /etc/selinux/config

# systemctl status firewalld.service
# systemctl stop firewalld.service
```
# 2.创建用户
```
# groupadd -r -g 1000 mysql
# useradd -d /home/mysql -r -m -s /sbin/nologin -g 1000 -u 1000 mysql
```

# 3.准备数据目录

**数据应该放在btrfs格式文件系统上，centos6放到lvm2文件系统上，支持快照**

```
# cd /mydata
# mkdir data
# mkdir /mydata/data/mysql_binlog
# chown mysql:mysql data
```

# 4.下载

```
# rpm -e mysql-server            #检查是否安装

# cd /tmp
# wget -c https://downloads.mysql.com/archives/get/p/23/file/mysql-5.7.28-linux-glibc2.12-x86_64.tar.gz
# md5sum mysql-5.7.28-linux-glibc2.12-x86_64.tar.gz
# tar xf mysql-5.7.28-linux-glibc2.12-x86_64.tar.gz -C /usr/local
# cd /usr/local/
# ln -sv 
# ln -snf mysql-5.7.28-linux-glibc2.12-x86_64 mysql
# cd mysql

```

# 5.安装

```
# yum -y install autoconf
[root@mysql220 mysql]# scripts/mysql_install_db --help
Usage: scripts/mysql_install_db [OPTIONS]
  --basedir=path       The path to the MySQL installation directory.
  --builddir=path      If using --srcdir with out-of-directory builds, you
                       will need to set this to the location of the build
                       directory where built files reside.
  --cross-bootstrap    For internal use.  Used when building the MySQL system
                       tables on a different host than the target.
  --datadir=path       The path to the MySQL data directory.
                       If missing, the directory will be created, but its
                       parent directory must already exist and be writable.
  --defaults-extra-file=name
                       Read this file after the global files are read.
  --defaults-file=name Only read default options from the given file name.
  --force              Causes mysql_install_db to run even if DNS does not
                       work.  In that case, grant table entries that
                       normally use hostnames will use IP addresses.
  --help               Display this help and exit.                     
  --ldata=path         The path to the MySQL data directory. Same as --datadir.
  --no-defaults        Don't read default options from any option file.
  --keep-my-cnf        Don't try to create my.cnf based on template. 
                       Useful for systems with working, updated my.cnf.
                       Deprecated, will be removed in future version.
  --random-passwords   Create and set a random password for all root accounts
                       and set the "password expired" flag,
                       also remove the anonymous accounts.
  --rpm                For internal use.  This option is used by RPM files
                       during the MySQL installation process.
  --skip-name-resolve  Use IP addresses rather than hostnames when creating
                       grant table entries.  This option can be useful if
                       your DNS does not work.
  --srcdir=path        The path to the MySQL source directory.  This option
                       uses the compiled binaries and support files within the
                       source tree, useful for if you don't want to install
                       MySQL yet and just want to create the system tables.
  --user=user_name     The login username to use for running mysqld.  Files
                       and directories created by mysqld will be owned by this
                       user.  You must be root to use this option.  By default
                       mysqld runs using your current login name and files and
                       directories that it creates will be owned by you.
Any other options are passed to the mysqld program.
```
```
生成mysql系统数据库
[root@mysql220 mysql]# scripts/mysql_install_db --datadir=/mydata/data --user=mysql
Installing MySQL system tables...2020-02-09 16:19:22 0 [Warning] 'THREAD_CONCURRENCY' is deprecated and will be removed in a future release.
2020-02-09 16:19:22 0 [Warning] TIMESTAMP with implicit DEFAULT value is deprecated. Please use --explicit_defaults_for_timestamp server option (see documentation for more details).
2020-02-09 16:19:22 0 [Note] Ignoring --secure-file-priv value as server is running with --bootstrap.
2020-02-09 16:19:22 0 [Note] ./bin/mysqld (mysqld 5.6.46-log) starting as process 6699 ...
2020-02-09 16:19:22 6699 [Note] InnoDB: Using atomics to ref count buffer pool pages
2020-02-09 16:19:22 6699 [Note] InnoDB: The InnoDB memory heap is disabled
2020-02-09 16:19:22 6699 [Note] InnoDB: Mutexes and rw_locks use GCC atomic builtins
2020-02-09 16:19:22 6699 [Note] InnoDB: Memory barrier is not used
2020-02-09 16:19:22 6699 [Note] InnoDB: Compressed tables use zlib 1.2.11
2020-02-09 16:19:22 6699 [Note] InnoDB: Using Linux native AIO
2020-02-09 16:19:22 6699 [Note] InnoDB: Using CPU crc32 instructions
2020-02-09 16:19:22 6699 [Note] InnoDB: Initializing buffer pool, size = 128.0M
2020-02-09 16:19:22 6699 [Note] InnoDB: Completed initialization of buffer pool
2020-02-09 16:19:22 6699 [Note] InnoDB: The first specified data file ./ibdata1 did not exist: a new database to be created!
2020-02-09 16:19:22 6699 [Note] InnoDB: Setting file ./ibdata1 size to 12 MB
2020-02-09 16:19:22 6699 [Note] InnoDB: Database physically writes the file full: wait...
2020-02-09 16:19:22 6699 [Note] InnoDB: Setting log file ./ib_logfile101 size to 48 MB
2020-02-09 16:19:23 6699 [Note] InnoDB: Setting log file ./ib_logfile1 size to 48 MB
2020-02-09 16:19:24 6699 [Note] InnoDB: Renaming log file ./ib_logfile101 to ./ib_logfile0
2020-02-09 16:19:24 6699 [Warning] InnoDB: New log files created, LSN=45781
2020-02-09 16:19:24 6699 [Note] InnoDB: Doublewrite buffer not found: creating new
2020-02-09 16:19:24 6699 [Note] InnoDB: Doublewrite buffer created
2020-02-09 16:19:24 6699 [Note] InnoDB: 128 rollback segment(s) are active.
2020-02-09 16:19:24 6699 [Warning] InnoDB: Creating foreign key constraint system tables.
2020-02-09 16:19:24 6699 [Note] InnoDB: Foreign key constraint system tables created
2020-02-09 16:19:24 6699 [Note] InnoDB: Creating tablespace and datafile system tables.
2020-02-09 16:19:24 6699 [Note] InnoDB: Tablespace and datafile system tables created.
2020-02-09 16:19:24 6699 [Note] InnoDB: Waiting for purge to start
2020-02-09 16:19:24 6699 [Note] InnoDB: 5.6.46 started; log sequence number 0
2020-02-09 16:19:24 6699 [Note] RSA private key file not found: /mydata/data//private_key.pem. Some authentication plugins will not work.
2020-02-09 16:19:24 6699 [Note] RSA public key file not found: /mydata/data//public_key.pem. Some authentication plugins will not work.
2020-02-09 16:19:24 6699 [Note] Binlog end
2020-02-09 16:19:24 6699 [Note] InnoDB: FTS optimize thread exiting.
2020-02-09 16:19:24 6699 [Note] InnoDB: Starting shutdown...
2020-02-09 16:19:26 6699 [Note] InnoDB: Shutdown completed; log sequence number 1625977
OK

Filling help tables...2020-02-09 16:19:26 0 [Warning] 'THREAD_CONCURRENCY' is deprecated and will be removed in a future release.
2020-02-09 16:19:26 0 [Warning] TIMESTAMP with implicit DEFAULT value is deprecated. Please use --explicit_defaults_for_timestamp server option (see documentation for more details).
2020-02-09 16:19:26 0 [Note] Ignoring --secure-file-priv value as server is running with --bootstrap.
2020-02-09 16:19:26 0 [Note] ./bin/mysqld (mysqld 5.6.46-log) starting as process 6721 ...
2020-02-09 16:19:26 6721 [Note] InnoDB: Using atomics to ref count buffer pool pages
2020-02-09 16:19:26 6721 [Note] InnoDB: The InnoDB memory heap is disabled
2020-02-09 16:19:26 6721 [Note] InnoDB: Mutexes and rw_locks use GCC atomic builtins
2020-02-09 16:19:26 6721 [Note] InnoDB: Memory barrier is not used
2020-02-09 16:19:26 6721 [Note] InnoDB: Compressed tables use zlib 1.2.11
2020-02-09 16:19:26 6721 [Note] InnoDB: Using Linux native AIO
2020-02-09 16:19:26 6721 [Note] InnoDB: Using CPU crc32 instructions
2020-02-09 16:19:26 6721 [Note] InnoDB: Initializing buffer pool, size = 128.0M
2020-02-09 16:19:26 6721 [Note] InnoDB: Completed initialization of buffer pool
2020-02-09 16:19:26 6721 [Note] InnoDB: Highest supported file format is Barracuda.
2020-02-09 16:19:26 6721 [Note] InnoDB: 128 rollback segment(s) are active.
2020-02-09 16:19:26 6721 [Note] InnoDB: Waiting for purge to start
2020-02-09 16:19:26 6721 [Note] InnoDB: 5.6.46 started; log sequence number 1625977
2020-02-09 16:19:26 6721 [Note] RSA private key file not found: /mydata/data//private_key.pem. Some authentication plugins will not work.
2020-02-09 16:19:26 6721 [Note] RSA public key file not found: /mydata/data//public_key.pem. Some authentication plugins will not work.
2020-02-09 16:19:26 6721 [Note] Binlog end
2020-02-09 16:19:26 6721 [Note] InnoDB: FTS optimize thread exiting.
2020-02-09 16:19:26 6721 [Note] InnoDB: Starting shutdown...
2020-02-09 16:19:28 6721 [Note] InnoDB: Shutdown completed; log sequence number 1625987
OK

To start mysqld at boot time you have to copy
support-files/mysql.server to the right place for your system

PLEASE REMEMBER TO SET A PASSWORD FOR THE MySQL root USER !
To do so, start the server, then issue the following commands:

  ./bin/mysqladmin -u root password 'new-password'
  ./bin/mysqladmin -u root -h mysql220 password 'new-password'

Alternatively you can run:

  ./bin/mysql_secure_installation

which will also give you the option of removing the test
databases and anonymous user created by default.  This is
strongly recommended for production servers.

See the manual for more instructions.

You can start the MySQL daemon with:

  cd . ; ./bin/mysqld_safe &

You can test the MySQL daemon with mysql-test-run.pl

  cd mysql-test ; perl mysql-test-run.pl

Please report any problems at http://bugs.mysql.com/

The latest information about MySQL is available on the web at

  http://www.mysql.com

Support MySQL by buying support/licenses at http://shop.mysql.com

New default config file was created as ./my.cnf and
will be used by default by the server when you start it.
You may edit this file to change server settings

WARNING: Default config file /etc/my.cnf exists on the system
This file will be read by default by the MySQL server
If you do not want to use this, either remove it, or use the
--defaults-file argument to mysqld_safe when starting the server

WARNING: Default config file /etc/mysql/my.cnf exists on the system
This file will be read by default by the MySQL server
If you do not want to use this, either remove it, or use the
--defaults-file argument to mysqld_safe when starting the server
```
```
[root@mysql220 mysql]# cp support-files/mysql.server /etc/rc.d/init.d/mysqld
[root@mysql220 mysql]# chkconfig --add mysqld
[root@mysql220 mysql]# chkconfig --list mysqld

注：该输出结果只显示 SysV 服务，并不包含
原生 systemd 服务。SysV 配置数据
可能被原生 systemd 配置覆盖。 

      要列出 systemd 服务，请执行 'systemctl list-unit-files'。
      查看在具体 target 启用的服务请执行
      'systemctl list-dependencies [target]'。

mysqld          0:关    1:关    2:开    3:开    4:开    5:开    6:关
ch# chkconfig --level 2345 mysqld on

# vim /root/.bash_profile 
 export PATH=$PATH:/usr/local/mysql/bin
或着
# vim /etc/profile.d/mysql.sh
PATH=/usr/local/mysql/bin:$PATH
```

# 6.配置

**准备配置文件**
- 配置格式: 类ini格式，为各程序均提供通过单个配置文件提供配置信息:[prog_name]
- 配置文件查找次序：
  
    /etc/my.cnf >  /etc/mysql/my.cnf > --default-extra-file=/PATH/TO/CONF_FILE > ~/.my.cnf

```
# mkdir /etc/mysql
# cp support-files/my-default.cnf /etc/mysql/my.cnf
```

1. MySQL的“utf8mb4”是真正的“UTF-8”。
2. MySQL的“utf8”是一种“专属的编码”，它能够编码的Unicode字符并不多。
3. utf8mb4的最低mysql版本支持版本为5.5.3+
4. SHOW VARIABLES WHERE Variable_name LIKE 'character_set_%' OR Variable_name LIKE 'collation%';
6. 必须保障以下变量都是utf8mb4

| 系统变量 | 描述 |
| --- | --- |
| character_set_client | (客户端来源数据使用的字符集) |
| character_set_connection | (连接层字符集) |
| character_set_database | (当前选中数据库的默认字符集) |
| character_set_results | (查询结果字符集) |
| character_set_server | (默认的内部操作字符集) |

```
[client]
port = 3306
socket = /tmp/mysql.sock

[mysqld]
port = 3306
user = mysql
basedir = /usr/local/mysql
datadir = /mydata/data/    //指定总目录 

socket = /tmp/mysql.sock
pid-file = /mydata/data/mysql.pid
log-error = /mydata/data/mysql_error.log

character-set-server = utf8mb4
init_connect='SET NAMES utf8mb4'

innodb_log_file_size = 256M
innodb_file_format = barracuda
innodb_strict_mode = 0
# 让每一个表数据库都是一个文件，方便管理
innodb_file_per_table = on
# 忽略名字的反向解析，加快速度
skip-name-resolve = on

#服务器ID，集群必填配置，区分机器编号，每台机器不同
server_id = 1

#开启二进制日志，行级记录，同步写入磁盘
log_bin = /mydata/data/mysql_binlog/mysql-bin
binlog_format = row
sync_binlog = 1

sql_mode='STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION'

symbolic-links=0
```
```

# service mysqld start
排错
# systemctl status mysqld.service

[root@mysql220 mariadb]# ps -ef|grep mysql
root      7140     1  0 16:54 pts/0    00:00:00 /bin/sh /usr/local/mysql/bin/mysqld_safe --datadir=/mydata/data/ --pid-file=/mydata/data/mysql.pid
mysql     7572  7140  0 16:54 pts/0    00:00:00 /usr/local/mysql/bin/mysqld --basedir=/usr/local/mysql --datadir=/mydata/data --plugin-dir=/usr/local/mysql/lib/plugin --user=mysql --log-error=/var/log/mariadb/mariadb.log --pid-file=/mydata/data/mysql.pid --socket=/mydata/data/mysql.sock --port=3306
```

# 7.安全初始化

```
/usr/local/mysql/bin/mysql_secure_installation
```
1. 当前密码空，回车
2. 设置root 密码
3. 删除匿名用户
4. 禁止root远程登陆
5. 移除测试库 no
6. 重载
   
```
[root@mysql220 data]# /usr/local/mysql/bin/mysql_secure_installation 



NOTE: RUNNING ALL PARTS OF THIS SCRIPT IS RECOMMENDED FOR ALL MySQL
      SERVERS IN PRODUCTION USE!  PLEASE READ EACH STEP CAREFULLY!

In order to log into MySQL to secure it, we'll need the current
password for the root user.  If you've just installed MySQL, and
you haven't set the root password yet, the password will be blank,
so you should just press enter here.

Enter current password for root (enter for none): 
OK, successfully used password, moving on...

Setting the root password ensures that nobody can log into the MySQL
root user without the proper authorisation.

Set root password? [Y/n] y
New password: 
Re-enter new password: 
Password updated successfully!
Reloading privilege tables..
 ... Success!


By default, a MySQL installation has an anonymous user, allowing anyone
to log into MySQL without having to have a user account created for
them.  This is intended only for testing, and to make the installation
go a bit smoother.  You should remove them before moving into a
production environment.

Remove anonymous users? [Y/n] y
 ... Success!

Normally, root should only be allowed to connect from 'localhost'.  This
ensures that someone cannot guess at the root password from the network.

Disallow root login remotely? [Y/n] y
 ... Success!

By default, MySQL comes with a database named 'test' that anyone can
access.  This is also intended only for testing, and should be removed
before moving into a production environment.

Remove test database and access to it? [Y/n] n
 ... skipping.

Reloading the privilege tables will ensure that all changes made so far
will take effect immediately.

Reload privilege tables now? [Y/n] y
 ... Success!




All done!  If you've completed all of the above steps, your MySQL
installation should now be secure.

Thanks for using MySQL!


Cleaning up...
```

# 8.部署sys数据库

https://github.com/mysql/mysql-sys
```
# yum -y install git
# git clone https://github.com/mysql/mysql-sys.git
# cd mysql-sys
# mysql -uroot -p < ./sys_56.sql

mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| mysql_binlog       |
| performance_schema |
| sys                |
| test               |
+--------------------+
6 rows in set (0.00 sec)
```

没有sys.schema_table_lock_waits 视图，mysql5.7中视图

# 9.utf8升级utf8mb4

## SQL修改
```
# 修改数据库
> ALTER DATABASE database_name CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

# 修改表
> ALTER TABLE table_name CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

# 修改表字段
> ALTER TABLE table_name CHANGE column_name column_name VARCHAR(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```
## mysql配置
```
default-character-set = utf8mb4
default-character-set = utf8mb4
character-set-client-handshake = FALSE
character-set-server = utf8mb4
collation-server = utf8mb4_unicode_ci
init_connect='SET NAMES utf8mb4'
```

utf8编码下，255长度的varchar长度约767，更改成utf8mb4后，最大只能支持191长度

# 10.升级至5.7.28

```
# service mysqld stop
# mkdir /mydata/data7

# wget -c https://downloads.mysql.com/archives/get/p/23/file/mysql-5.7.28-linux-glibc2.12-x86_64.tar.gz
# tar xf mysql-5.7.28-linux-glibc2.12-x86_64.tar.gz -C /usr/local/
# cd /usr/local/
# ln -snf mysql-5.7.28-linux-glibc2.12-x86_64 mysql
# cp /usr/local/mysql/support-files/mysql.server /etc/init.d/mysqld
# vim /etc/init.d/mysqld 
basedir=/usr/local/mysql
datadir=/mydata/data
# service mysqld start
Starting MySQL............ SUCCESS! 
# /usr/local/mysql/bin/mysql_upgrade -uroot -p
Enter password: 
Checking if update is needed.
Checking server version.
Running queries to upgrade MySQL server.
Checking system database.
mysql.columns_priv                                 OK
mysql.db                                           OK
mysql.engine_cost                                  OK
mysql.event                                        OK
mysql.func                                         OK
mysql.general_log                                  OK
mysql.gtid_executed                                OK
mysql.help_category                                OK
mysql.help_keyword                                 OK
mysql.help_relation                                OK
mysql.help_topic                                   OK
mysql.innodb_index_stats                           OK
mysql.innodb_table_stats                           OK
mysql.ndb_binlog_index                             OK
mysql.plugin                                       OK
mysql.proc                                         OK
mysql.procs_priv                                   OK
mysql.proxies_priv                                 OK
mysql.server_cost                                  OK
mysql.servers                                      OK
mysql.slave_master_info                            OK
mysql.slave_relay_log_info                         OK
mysql.slave_worker_info                            OK
mysql.slow_log                                     OK
mysql.tables_priv                                  OK
mysql.time_zone                                    OK
mysql.time_zone_leap_second                        OK
mysql.time_zone_name                               OK
mysql.time_zone_transition                         OK
mysql.time_zone_transition_type                    OK
mysql.user                                         OK
Found outdated sys schema version 1.5.1.
Upgrading the sys schema.
Checking databases.
sys.sys_config                                     OK
test.t                                             OK
Upgrade process completed successfully.
Checking if update is needed.

# service mysqld restart

[root@mysql220 local]# mysql -uroot -p
Enter password: 
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 2
Server version: 5.7.28-log MySQL Community Server (GPL)

Copyright (c) 2000, 2019, Oracle and/or its affiliates. All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> show databases
    -> ;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| mysql_binlog       |
| performance_schema |
| sys                |
| test               |
+--------------------+
6 rows in set (0.00 sec)

mysql>  drop database sys
    -> ;
Query OK, 101 rows affected (0.01 sec)

# cd /tmp/mysql-sys
# mysql -uroot -p
SOURCE ./sys_57.sql
```
