---
title: "MYSQL 应用基础"
permalink: /dba/mysql-base/
excerpt: "MYSQL 应用基础"
last_modified_at: 2020-02-01T21:36:11-04:00
categories: mysql
redirect_from:
  - /theme-setup/
toc: true
---
<!--马哥22课笔记，日期：2月2日-->
<!--马哥23课笔记，日期：2月5日-->
# 1.关系型数据库

**关系模型：二维关系 row,column**

**关系数据库有六种范式，范式越高数据库冗余更小**
- 第一范式基本要求：每一列都是不可分割的原子数据项，即无重复的域，非关系型数据库则无此要求。**行的要求**
- 第二范式首先满足第一范式，每个实例或记录可以被唯一区分，即不能有重复的行。**列的要求**
- 第三范式必须满足第二范式，一个关系中不包含已在其他关系已包含的非主关键字信息。多张表不能存储相同的非主键字段。**表与表之间的关系**
- 一般关系型数据库满足到第三范式，第五范式过于完美，太理想化，完美本身就是不完美。

**因为要满足以上范式，关系型数据库经过拆表后，查询时又需要合表，导致影响性能，这是范式要求的必然结果**
{: .notice}

Relational RDBMS:
- MYSQL
  - MYSQL
  - MariaDB
  - Percona-Server
- PostgreSQL:pgsql --> EnterpriseDB
- Oracle
- MSSQL

# 2.Mariadb

[www.mariadb.com](https://mariadb.com/)  企业版
[www.mariadb.org](https://mariadb.org/)  社区版

- 分支1：10.5.12
- 分支2：5.5.67

MariaDB特性：
- 插件式存储引擎
- 存储引擎也称为表类型
- 支持更多的存储引擎
  - MyISAM 老，不支持事务 --> Aria
  - InnoDB 支持事务、锁等 --> XtraDB
- 诸多扩展和新特性
- 提供较多的测试组件
- true open source
- 支持多数据库存在，oracle只支持一个库，mysql每一个数据库只是一个目录而已

# 3.安装Mariadb

## 3.1.安装方法
- 源代码：编译安装
- 二进制格式程序包：展开至特定路径，并经过简单配置后即可使用
- 程序包管理管理程序包：
  - rpm
    - 项目官方提供
    - OS Vendor提供
  - deb
- CentOS 7直接安装，CentOS 6额外安装

### 3.1.1.创建用户
```
# groupadd -r -g 306 mysql
# useradd -r -g 306 -u 306 mysql
```

### 3.1.2.准备数据目录

**数据应该放在btrfs格式文件系统上，centos6放到lvm2文件系统上，支持快照**
```
# echo "- - -" > /sys/class/scsi_host/host0/scan
# fdisk -l
# fdisk /dev/sdb
欢迎使用 fdisk (util-linux 2.23.2)。

更改将停留在内存中，直到您决定将更改写入磁盘。
使用写入命令前请三思。

Device does not contain a recognized partition table
使用磁盘标识符 0x6d065a3c 创建新的 DOS 磁盘标签。

命令(输入 m 获取帮助)：n
Partition type:
   p   primary (0 primary, 0 extended, 4 free)
   e   extended
Select (default p): p
分区号 (1-4，默认 1)：1
起始 扇区 (2048-104857599，默认为 2048)：
将使用默认值 2048
Last 扇区, +扇区 or +size{K,M,G} (2048-104857599，默认为 104857599)：
将使用默认值 104857599
分区 1 已设置为 Linux 类型，大小设为 50 GiB

命令(输入 m 获取帮助)：w
The partition table has been altered!

Calling ioctl() to re-read partition table.
正在同步磁盘。
```
 如果是继续在sda磁盘上分区，分区号继续写的是3，并用partx -a /dev/sda重新识别分区。
```
[root@mysql220 mysql]# fdisk /dev/sdb
欢迎使用 fdisk (util-linux 2.23.2)。

更改将停留在内存中，直到您决定将更改写入磁盘。
使用写入命令前请三思。


命令(输入 m 获取帮助)：p

磁盘 /dev/sdb：53.7 GB, 53687091200 字节，104857600 个扇区
Units = 扇区 of 1 * 512 = 512 bytes
扇区大小(逻辑/物理)：512 字节 / 512 字节
I/O 大小(最小/最佳)：512 字节 / 512 字节
磁盘标签类型：dos
磁盘标识符：0x6d065a3c

   设备 Boot      Start         End      Blocks   Id  System
/dev/sdb1            2048   104857599    52427776   83  Linux

命令(输入 m 获取帮助)：t
已选择分区 1
Hex 代码(输入 L 列出所有代码)：8e  
已将分区“Linux”的类型更改为“Linux LVM”

命令(输入 m 获取帮助)：w
The partition table has been altered!

Calling ioctl() to re-read partition table.
正在同步磁盘。

[root@mysql220 mysql]# pvcreate /dev/sdb1
  Physical volume "/dev/sdb1" successfully created.
[root@mysql220 mysql]# vgs
  VG     #PV #LV #SN Attr   VSize   VFree
  centos   1   2   0 wz--n- <29.00g 4.00m
[root@mysql220 mysql]# vgcreate myvg /dev/sdb1
  Volume group "myvg" successfully created
[root@mysql220 mysql]# lvcreate -L 50G -n mydata myvg
  Volume group "myvg" has insufficient free space (12799 extents): 12800 required.
[root@mysql220 mysql]# lvcreate -L 49G -n mydata myvg
  Logical volume "mydata" created.

[root@mysql220 mysql]#  yum install xfsprogs
[root@mysql220 mysql]#  lsmod |grep xfs

[root@mysql220 mysql]# mkfs.xfs /dev/myvg/mydata 
meta-data=/dev/myvg/mydata       isize=512    agcount=4, agsize=3211264 blks
         =                       sectsz=512   attr=2, projid32bit=1
         =                       crc=1        finobt=0, sparse=0
data     =                       bsize=4096   blocks=12845056, imaxpct=25
         =                       sunit=0      swidth=0 blks
naming   =version 2              bsize=4096   ascii-ci=0 ftype=1
log      =internal log           bsize=4096   blocks=6272, version=2
         =                       sectsz=512   sunit=0 blks, lazy-count=1
realtime =none                   extsz=4096   blocks=0, rtextents=0
[root@mysql220 mysql]# mkdir /mydata
[root@mysql220 mysql]# vim /etc/fstab
[root@mysql220 mysql]#mount -a
```
```
/dev/myvg/mydata /mydata xfs defaults   0 0
```
```
# cd /mydata
# mkdir data
# chown mysql:mysql data
```

### 3.1.3.下载

```
# rpm -e mysql-server            #检查是否安装

# cd /tmp
# wget -c https://downloads.mariadb.org/interstitial/mariadb-5.5.67/bintar-linux-x86_64/mariadb-5.5.67-linux-x86_64.tar.gz/from/http%3A//mariadb.nethub.com.hk/?serve&change_mirror_from=1
# md5sum mariadb-5.5.67-linux-x86_64.tar.gz
# tar xf mariadb-5.5.67-linux-x86_64.tar.gz -C /usr/local
# cd /usr/local/
# ln -sv mariadb-5.5.67-linux-x86_64 mysql
# cd mysql
# chown -R root:mysql ./*
```

### 3.1.4.安装mariadb

```
[root@mysql220 mysql]# scripts/mysql_install_db --help
[root@mysql220 mysql]# scripts/mysql_install_db --datadir=/mydata/data --user=mysql
Installing MariaDB/MySQL system tables in '/mydata/data' ...
200206 16:40:32 [Note] ./bin/mysqld (mysqld 5.5.67-MariaDB) starting as process 9235 ...
OK
Filling help tables...
200206 16:40:32 [Note] ./bin/mysqld (mysqld 5.5.67-MariaDB) starting as process 9244 ...
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
cd '.' ; ./bin/mysqld_safe --datadir='/mydata/data'

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

# export PATH=$PATH:/usr/local/mysql/bin
```

### 3.1.5.配置mariadb

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
datadir = /mydata/data
innodb_file_per_table = on
skip_name_resolve = on
```
```
# mkdir /var/log/mariadb
# chown root:mysql /var/log/mariadb/
# cd /var/log/mariadb
# touch mariadb.log
# chown mysql:mysql mariadb.log

# service mysqld start
[root@mysql220 mariadb]# ps -ef|grep mysql
root     10906     1  0 17:12 pts/1    00:00:00 /bin/sh /usr/local/mysql/bin/mysqld_safe --datadir=/mydata/data --pid-file=/mydata/data/mysql220.pid
mysql    11299 10906  0 17:12 pts/1    00:00:00 /usr/local/mysql/bin/mysqld --basedir=/usr/local/mysql --datadir=/mydata/data --plugin-dir=/usr/local/mysql/lib/plugin --user=mysql --log-error=/var/log/mariadb/mariadb.log --pid-file=/mydata/data/mysql220.pid --socket=/tmp/mysql.sock --port=3306
```

### 3.1.5.mariadb程序组成

- 客户端
  - mysql CLI交互式客户端程序
  - mysqldump  
  - mysqladmin
- 服务端
  - mysqld_safe
  - mysqld
  - mysqld_multi

**服务器监听的两种socket 地址：**
- ip socket: 监听在tcp 3306 支持远程通信
- unix sock： 监听在sock文件上/tmp/mysql.sock  /var/lib/mysql/mysql.sock 仅支持本地通信
  - server：localhost ,127.0.0.1 自动选择unix sock

### 3.1.6.安全初始化

**命令行互交式客户端程序mysql**
  - -uUSERNAME  默认root
  - -hHOST      默认本地
  - -pPASSWORD  默认空
  - 支持命令历史

**账户分为两部分 'USERNAME'@'HOST'**  
- HOST限制远程连接
- 支持通配符
  - %匹配任意长度任意字符   172.16.%.%
  - _匹配任意单个字符

```
[root@mysql220 bin]# mysql
Welcome to the MariaDB monitor.  Commands end with ; or \g.
Your MariaDB connection id is 2
Server version: 5.5.67-MariaDB MariaDB Server

Copyright (c) 2000, 2018, Oracle, MariaDB Corporation Ab and others.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

MariaDB [(none)]> use mysql
Database changed
MariaDB [mysql]> select User,Host,Password From user;
+------+-----------+----------+
| User | Host      | Password |
+------+-----------+----------+
| root | localhost |          |
| root | mysql220  |          |
| root | 127.0.0.1 |          |
| root | ::1       |          |
|      | localhost |          |
|      | mysql220  |          |
+------+-----------+----------+
6 rows in set (0.00 sec)
```
**注意多个密码为空**

**安全初始化，全部设置密码，删除匿名用户**

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
```
[root@mysql220 bin]# mysql
ERROR 1045 (28000): Access denied for user 'root'@'localhost' (using password: NO)
[root@mysql220 bin]# mysql -uroot -p
Enter password: 
Welcome to the MariaDB monitor.  Commands end with ; or \g.
Your MariaDB connection id is 10
Server version: 5.5.67-MariaDB MariaDB Server

Copyright (c) 2000, 2018, Oracle, MariaDB Corporation Ab and others.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

MariaDB [(none)]> use mysql
Database changed
MariaDB [mysql]> select User,Host,Password From user;
+------+-----------+-------------------------------------------+
| User | Host      | Password                                  |
+------+-----------+-------------------------------------------+
| root | localhost | *85D84174C7E6E5EAB95615B2D11213BF94AC1F3E |
| root | 127.0.0.1 | *85D84174C7E6E5EAB95615B2D11213BF94AC1F3E |
| root | ::1       | *85D84174C7E6E5EAB95615B2D11213BF94AC1F3E |
+------+-----------+-------------------------------------------+
3 rows in set (0.00 sec)
```

### 3.1.7.命令分类

- 客户端命令：本地执行**
  - mysql>  help
  - 每个命令都有完整形式和简写形式 status , \s
- 服务端命令：通过mysql协议发往服务器执行并取回结构
  - 每个命令都必须有命令结束符合
  - select version();
  - select 1+1;
  - help create database;

 




