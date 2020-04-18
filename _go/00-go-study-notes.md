---
title: "Go语言学习笔记"
permalink: /go/study_notes/
excerpt: "Go语言学习笔记"
last_modified_at: 2020-03-30T21:36:11-04:00
categories: go
redirect_from:
  - /theme-setup/
toc: true
toc_sticky: true
---

<!--
Go语言核心36讲
1. 学习路线
-->

<img src="{{ site.url }}{{ site.baseurl }}/assets/images/2020-03-30-gopath.png" alt="">

[https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/preface.md](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/preface.md)

<div class="notice">
  <b>Go语言基础知识</b><br />
  1. Go语言规范<br />
  2. Go语言命令<br />
  3. Go语言编程基础<br />
  4. Go语言并发编程<br />
  <b>学习资源</b><br />
  <ul>
    <li>Go语言官网 https://golang.google.cn/</li>
    <li>Go语言规范文档 https://golang.google.cn/ref/spec</li>
    <li>Go语言命令文档 https://golang.google.cn/cmd/go</li>
    <li>Go程序编辑器和IDE https://golang.google.cn/doc/editors.html</li>
    <li>Go语言wiki https://github.com/golang/go/wiki</li>
    <li>Go并发编程实战</li>
    <li>Go命令教程 https://github.com/hyper0x/go_command_tutorial</li>
    <li>Go语言第一课 http://www.imooc.com/learn/345</li>
  </ul>  
  <b>Go语言进阶</b><br />
  1. Go语言数据类型使用进阶<br />
  2. Go语言标准库使用进阶<br />
  3. Go语言并发编程进阶<br />
  <b>进阶学习资源</b><br />
  <ul>
    <li>Effective Go https://golang.google.cn/doc/effective_go.html</li>
    <li>Go语言内存模式 https://golang.google.cn/ref/mem</li>
    <li>Go程序诊断 https://golang.google.cn/doc/diagnostics.html</li>
    <li>Go by Example https://gobyexample.com/</li>
    <li>Awesome Go https://github.com/avelino/awesome-go</li>
    <li>Go 语言优秀资源整理 https://shockerli.net/post/go-awesome/</li>
  </ul> 
</div>

课件： https://github.com/hyper0x/Golang_Puzzlers


### 安装

https://golang.google.cn/dl/

```
# wget https://dl.google.com/go/go1.14.1.linux-amd64.tar.gz
# tar -xvf go1.14.1.linux-amd64.tar.gz 
# mv go /opt/
# vi /root/.bash_profile 
export GOROOT=/opt/go
export GOPATH=/home/go
export PATH=$PATH:$GOROOT/bin:$GOPATH

# go version
go version go1.14.1 linux/amd64
```

windows环境下创建环境变量

GOROOT：Go 语言安装根目录的路径，也就是 GO 语言的安装路径。
GOPATH：若干工作区目录的路径。是我们自己定义的工作空间。
GOBIN：GO 程序生成的可执行文件（executable file）的路径。

