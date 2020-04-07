---
title: "Atom使用方法"
published: true
categories: atom
permalink: how-to-use-atom.html
summary: "如何安装使用Atom"
toc: true
---

## 1.Download

https://atom.io/

## 2.Setting

- Core settings : utf8

## 3.Packages

**https://github.com/Glavin001/**

- atom-beautify 代码自动格式化
- minimap 预览所有代码 https://github.com/atom-minimpa/minimap.git
- linter 代码检查框架
- linter-eslint  java代码检查
- linter-csslint css代码检查
- pretty-json  json格式化
- emmet 快速书写代码
- autocomplete-paths 自动补全文件路径
- highlight-selected  高亮显示选择单词
- file-type-icons 文件图标

## 4.Packages Install

 setting > install > install Packages 输入插件名称 搜索

 -	1. 安装Atom
 -	2. 先安装上NET .Frameork框架，win7安装3.5的，win10安装4.5以上的
 -	3. 安装git for windows 64
 -	4. 安装node.js
 -	5. cd   C:\Users\自己电脑名字\.atom\packages
 -	6. git clone https://github.com/JoelBesada/activate-power-mode.git
 -	7. cd   C:\Users\自己电脑名字\.atom\packages\activate-power-mode
 -	8. npm install

## 5.markdown Packages

- markdown-preview-plus 增强预览 [安装前将markdown-preview插件Disable]
- markdown-scroll-sync  同步滚动
- language-markdown 代码自动补全，高亮显示
- markdown-img-paste 图片粘贴,图片保存到md文件相同目录下
- markdown-table-editor 表格编辑
- markdown-themeable-pdf pdf导出
- pdf-view

## 6.config proxy
  修改C:\Users\Administrator\.atom\.apm下的文件.apmrc
```
  http-proxy=http://127.0.0.1:1080
  https-proxy=http://127.0.0.1:1080
  strict-ssl=false
```

    apm config set strict-ssl false

    apm config set http-proxy null
    apm config set http-proxy socks5:127.0.0.1:1080
    apm config get http-proxy

    apm config set https-proxy null
    apm config set https-proxy socks5:127.0.0.1:1080
    apm config get https-proxy

    apm config list

## 7.npm 常用命令

- npm install xxx 安装模块
- npm install xxx -g 将模块安装到全局环境中。
- npm ls 查看安装的模块及依赖
- npm ls -g 查看全局安装的模块及依赖
- npm uninstall xxx  (-g) 卸载模块
- npm cache clean 清理缓存
- npm help xxx  查看帮助
- npm update moduleName   更新node模块

## 9.error

- TypeError: Right-hand side of 'instanceof' is not callable
  markdown-scroll-sync 版本2.1.2 markdown-preview-plus 3.*版本不兼容
  1. 卸载markdown-preview-plus3.x，安装markdown-preview-plus 2.4.16
     https://github.com/atom-community/markdown-preview-plus/tree/v2.4.16
     放到C:\Users\xxx\.atom\package
     执行 apm install markdown-preview-plus
  2. 不使用 markdown-scroll-sync
