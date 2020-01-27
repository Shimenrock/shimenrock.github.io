---
title: "GIT日常命令整理"
published: true
categories: GIT
permalink: git-command.html
summary: "GIT日常命令整理"
tags: [GIT]
toc: true
---

- 进入工作区
```
cd e:/WORK_SYNC/102_GitHub/Blog
```
- 检查代码状态
```
git status
```
- 增加代码管理
```
git add README190929(0).md
git add articles/Git/
git add README190929\(0\).md
```
- 更新被修改(modified)和被删除(deleted)文件变化，不包括新文件(new)
```
git add -u   
```
- 添加所有变化
```
git add -A 
```
- 添加新文件(new)和被修改(modified)文件，不包括被删除(deleted)文件
```
git add .       
```                                             
- 提交
```
 git commit -m '20190929-2'' 
```
- 先检查
```
git checkout master 
```
- 和远端合并
```
git merge github/master
```
- 上传
```
git push github master 
```

## Duplicating a repository
- Clone to local
```
git clone git@github.com:geektime-geekbang/geektime-ELK.git {repository 上复制的地址}
```
- Mirroring a repository
```
$ git clone --bare https://github.com/exampleuser/old-repository.git
$ cd old-repository.git
$ git push --mirror https://github.com/exampleuser/new-repository.git
$ cd ..
$ rm -rf old-repository.git
```