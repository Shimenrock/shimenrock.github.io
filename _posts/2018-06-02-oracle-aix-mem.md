---
title:  "Aix system memory usage"
published: true
summary: "Aix system memory usage"
read_time: false
comments: false
related: false
author_profile: false
categories: 
  - dbascript
tags: 
  - AIX
  - memory
---

# 内存使用量
```
svmon -G
inuse：是物理内存使用量，这里是以4K为单位，所以1037487*4096=4249546752(424M)
virtual:是虚拟内存使用量，这里是以4K为单位，所以378796*4096=1551548416（155M）
```
# 内存从大到小排序
```
ps aux | head -1 ; ps aux | sort -rn +3 | head -10
```
# 内存从大到小排序(详细执行命令)
```
ps -ealf | head -1 ; ps -ealf | sort -rn +9 | head  -10
```
# 根据某个命令或进程名，查看内存使用率
```
ps aux | head -1 ; ps aux | grep topas
```
# 显示使用物理内存最多的3个进程：
```
# svmon -uP -t 3|grep -p Pid|grep '^.*[0-9] '
 6553834 java             51279     8917        0    50938      N     Y     N
 4456680 java             34626     8874        0    34608      N     Y     N
 5701730 BESClient        29564     8882        0    25689      Y     Y     N
```
- 输出的格式顺序为 Pid Command Inuse Pin Pgsp Virtual 64-bit Mthrd 
- 可以计算出X程序所使用的实存为51279×4096＝210038784，约为210MB 
# 显示交换区使用物理内存最多的3个进程
```
# svmon -gP -t 3|grep -p Pid|grep &apos;^.*[0-9] &apos; 
 1966206 shlap64          26782     8880        0    26771      Y     N     N
       0 swapper           9872     8912        0     9872      Y     N     N
       1 init             22094     8836        0    22076      N     N     N
第一个程序X所使用的交换区大小约为 26782×4096 =10510336 字节，大约为10MB空间 
```
# 每隔三秒显示使用最多的段
```
# svmon -S -t 3 -i 3 
Vsid Esid Type Description Inuse Pin Pgsp Virtual 
4f08 -    clnt 37505 0 - - 
11e1 -    clnt 33623 0 - - 
8811 -    work kernel pinned heap 12637 6547 8091 19397 
可见，Vsid为4f08的段使用最多 
```
# 显示指定PID内存使用情况
```
svmon -pP 22674 
看PID为22674的进程所使用的为那些文件 
Pid Command nuse Pin Pgsp Virtual 64-bit Mthrd 
22674 java 29333 1611 2756 32404 N Y 

Vsid Esid Type Description Inuse Pin Pgsp Virtual 
0 0 work kernel seg 2979 1593 1659 4561 
a056 - work 43 16  3   46 
1e03 2 work process private 77 2   17  93 
1080 - pers /dev/hd2:69742 1 0   -   - 
f8bd f work shared library data 84 0   11  99 
60ee 8 work shmat/mmap 0 0   0   0 
70ec - pers /dev/hd2:69836 1 0   -   - 
```
# 通过ncheck命令，检查Vsid都使用了哪些文件。
```
ncheck a056
```

# 查看物理内存总量
```
# cat mem1.sh
#!/usr/bin/ksh
#mem totle
totalmem=$(vmstat -v|head -n 1|awk &apos;{print $1/256}&apos;)
echo "mem totle:"
echo $totalmem MB
echo
```
# 查看每个用户占用物理内存的数量
# cat mem2.sh
```
usermem=$(for username in `cat /etc/passwd|awk -F: &apos;{print $1}&apos;`
do
svmon -U $username|grep $username" "
done)
usermem=`echo "$usermem"|grep -v "0        0        0        0"|awk &apos;{print $1,$2/256,"MB"}&apos;`
echo "singe user pmem"
echo "$usermem"
usermem=$(echo "$usermem"|awk &apos;BEGIN{sum1=0;}{sum1=sum1+$2;}END{print sum1;}&apos;)
usermem=$(echo $usermem|awk -F\. &apos;{print $1}&apos;)
echo "singe user pmem :" $usermem MB
echo
```

