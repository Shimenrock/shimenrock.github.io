---
title: "mariadb 10.4.12 二进制安装"
permalink: /dba/install-mariadb10/
excerpt: "mariadb 10.4.12 二进制安装"
last_modified_at: 2020-02-08T21:36:11-04:00
categories: mariadb
redirect_from:
  - /theme-setup/
toc: true
---

**解决：mariadb上未移植sys库问题**
{: .notice}

# 1.官网

[www.mariadb.com](https://mariadb.com/)  企业版
[www.mariadb.org](https://mariadb.org/)  社区版

MariaDB 10.4.12 Stable 2020-01-28  898.6MB
MariaDB 10.2.31 Stable 2020-01-28  461.0MB


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
# chown mysql:mysql data
```

# 4.下载

```
# rpm -e mysql-server            #检查是否安装

# cd /tmp
# wget -c https://downloads.mariadb.org/interstitial/mariadb-10.2.31/bintar-linux-x86_64/mariadb-10.2.31-linux-x86_64.tar.gz/from/http%3A//mirrors.tuna.tsinghua.edu.cn/mariadb/
# md5sum mariadb-10.2.31-linux-x86_64.tar.gz 
# tar xf mariadb-10.2.31-linux-x86_64.tar.gz -C /usr/local
# cd /usr/local/
# ln -sv mariadb-10.2.31-linux-x86_64 mysql
# ln -snf mariadb-10.2.31-linux-x86_64 mysql
# cd mysql
```

# 5.安装mariadb

```
[root@mysql220 mysql]# scripts/mysql_install_db --help
Usage: scripts/mysql_install_db [OPTIONS]
  --auth-root-authentication-method=normal|socket
                       Chooses the authentication method for the created
                       initial root user. The historical behavior is 'normal'
                       to creates a root user that can login without password,
                       which can be insecure. The default behavior 'socket'
                       sets an invalid root password but allows the system root
                       user to login as MariaDB root without a password.
  --auth-root-socket-user=user
                       Used with --auth-root-authentication-method=socket. It
                       specifies the name of the second MariaDB root account,
                       as well as of the system account allowed to access it.
                       Defaults to the value of --user.
  --basedir=path       The path to the MariaDB installation directory.
  --builddir=path      If using --srcdir with out-of-directory builds, you
                       will need to set this to the location of the build
                       directory where built files reside.
  --cross-bootstrap    For internal use.  Used when building the MariaDB system
                       tables on a different host than the target.
  --datadir=path       The path to the MariaDB data directory.
  --defaults-extra-file=name
                       Read this file after the global files are read.
  --defaults-file=name Only read default options from the given file name.
  --defaults-group-suffix=name
                       In addition to the given groups, read also groups with
                       this suffix
  --force              Causes mysql_install_db to run even if DNS does not
                       work.  In that case, grant table entries that
                       normally use hostnames will use IP addresses.
  --help               Display this help and exit.                     
  --ldata=path         The path to the MariaDB data directory. Same as
                       --datadir.
  --no-defaults        Don't read default options from any option file.
  --defaults-file=path Read only this configuration file.
  --rpm                For internal use.  This option is used by RPM files
                       during the MariaDB installation process.
  --skip-name-resolve  Use IP addresses rather than hostnames when creating
                       grant table entries.  This option can be useful if
                       your DNS does not work.
  --skip-test-db       Don't install a test database.
  --srcdir=path        The path to the MariaDB source directory.  This option
                       uses the compiled binaries and support files within the
                       source tree, useful for if you don't want to install
                       MariaDB yet and just want to create the system tables.
  --user=user_name     The login username to use for running mysqld.  Files
                       and directories created by mysqld will be owned by this
                       user.  You must be root to use this option.  By default
                       mysqld runs using your current login name and files and
                       directories that it creates will be owned by you.

All other options are passed to the mysqld program
```
```
生成mysql系统数据库
[root@mysql220 mysql]# scripts/mysql_install_db --datadir=/mydata/data --user=mysql
Installing MariaDB/MySQL system tables in '/mydata/data' ...
OK

To start mysqld at boot time you have to copy
support-files/mysql.server to the right place for your system


PLEASE REMEMBER TO SET A PASSWORD FOR THE MariaDB root USER !
To do so, start the server, then issue the following commands:

'./bin/mysqladmin' -u root password 'new-password'
'./bin/mysqladmin' -u root -h mysql220 password 'new-password'

Alternatively you can run:
'./bin/mysql_secure_installation'

which will also give you the option of removing the test
databases and anonymous user created by default.  This is
strongly recommended for production servers.

See the MariaDB Knowledgebase at http://mariadb.com/kb or the
MySQL manual for more instructions.

You can start the MariaDB daemon with:
cd '.' ; ./bin/mysqld_safe --datadir='/mydata/data10'

You can test the MariaDB daemon with mysql-test-run.pl
cd './mysql-test' ; perl mysql-test-run.pl

Please report any problems at http://mariadb.org/jira

The latest information about MariaDB is available at http://mariadb.org/.
You can find additional information about the MySQL part at:
http://dev.mysql.com
Consider joining MariaDB's strong and vibrant community:
https://mariadb.org/get-involved/
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

# vim /root/.bash_profile 
 export PATH=$PATH:/usr/local/mysql/bin
或着
# vim /etc/profile.d/mysql.sh
PATH=/usr/local/mysql/bin:$PATH
```

# 6.配置mariadb

**准备配置文件**
- 配置格式: 类ini格式，为各程序均提供通过单个配置文件提供配置信息:[prog_name]
- 配置文件查找次序：
  
    /etc/my.cnf >  /etc/mysql/my.cnf > --default-extra-file=/PATH/TO/CONF_FILE > ~/.my.cnf

```
# mkdir /etc/mysql
# cp support-files/my-large.cnf /etc/mysql/my.cnf
```

**添加三个选项**
```
datadir = /mydata/data         //指定总目录，必须的
innodb_file_per_table = on     //让每一个表数据库都是一个文件，方便管理
skip_name_resolve = on         //忽略名字的反向解析，加快速度
```
```
# mkdir /var/log/mariadb
# chown root:mysql /var/log/mariadb/
# cd /var/log/mariadb
# touch mariadb.log
# chown mysql:mysql mariadb.log

# service mysqld start
排错
# systemctl status mysqld.service

[root@mysql220 mariadb]# ps -ef|grep mysql
root     10906     1  0 17:12 pts/1    00:00:00 /bin/sh /usr/local/mysql/bin/mysqld_safe --datadir=/mydata/data --pid-file=/mydata/data/mysql220.pid
mysql    11299 10906  0 17:12 pts/1    00:00:00 /usr/local/mysql/bin/mysqld --basedir=/usr/local/mysql --datadir=/mydata/data --plugin-dir=/usr/local/mysql/lib/plugin --user=mysql --log-error=/var/log/mariadb/mariadb.log --pid-file=/mydata/data/mysql220.pid --socket=/tmp/mysql.sock --port=3306
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
[root@mysql220 bin]# /usr/local/mysql/bin/mysql_secure_installation
/usr/local/mysql/bin/mysql_secure_installation:行393: find_mysql_client: 未找到命令

NOTE: RUNNING ALL PARTS OF THIS SCRIPT IS RECOMMENDED FOR ALL MariaDB
      SERVERS IN PRODUCTION USE!  PLEASE READ EACH STEP CAREFULLY!

In order to log into MariaDB to secure it, we'll need the current
password for the root user.  If you've just installed MariaDB, and
you haven't set the root password yet, the password will be blank,
so you should just press enter here.

Enter current password for root (enter for none): 
OK, successfully used password, moving on...

Setting the root password ensures that nobody can log into the MariaDB
root user without the proper authorisation.

Set root password? [Y/n] Y
New password: 
Re-enter new password: 
Sorry, passwords do not match.

New password: 
Re-enter new password: 
Sorry, passwords do not match.

New password: 
Re-enter new password: 
Sorry, passwords do not match.

New password: 
Re-enter new password: 
Password updated successfully!
Reloading privilege tables..
 ... Success!


By default, a MariaDB installation has an anonymous user, allowing anyone
to log into MariaDB without having to have a user account created for
them.  This is intended only for testing, and to make the installation
go a bit smoother.  You should remove them before moving into a
production environment.

Remove anonymous users? [Y/n] Y    
 ... Success!

Normally, root should only be allowed to connect from 'localhost'.  This
ensures that someone cannot guess at the root password from the network.

Disallow root login remotely? [Y/n] y
 ... Success!

By default, MariaDB comes with a database named 'test' that anyone can
access.  This is also intended only for testing, and should be removed
before moving into a production environment.

Remove test database and access to it? [Y/n] n
 ... skipping.

Reloading the privilege tables will ensure that all changes made so far
will take effect immediately.

Reload privilege tables now? [Y/n] y
 ... Success!

Cleaning up...

All done!  If you've completed all of the above steps, your MariaDB
installation should now be secure.

Thanks for using MariaDB!
```

# 8.部署sys数据库

https://github.com/FromDual/mariadb-sys
```
# yum -y install git
# git clone https://github.com/FromDual/mariadb-sys.git
# cd mariadb-sys
# mysql -uroot -p < ./sys_10.sql

MariaDB [(none)]> show databases 
    -> ;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| sys                |
| test               |
+--------------------+
5 rows in set (0.00 sec)
```

