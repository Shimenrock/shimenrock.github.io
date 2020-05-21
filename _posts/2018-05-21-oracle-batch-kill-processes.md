---
title:  "Batch kill processes"
published: true
summary: "Batch kill processes"
read_time: false
comments: false
related: false
author_profile: false
categories: 
  - dbascript
tags: 
  - kill
  - processes
  - oracle
---

# 操作系统批量杀死oracle会话

<div class="notice">
  <p><b>使用前一定确认数据库里不存在GB级别的数据增删改批量操作，避免大回退，否则本意用特殊手段恢复业务使用，反而导致数据库长时间回滚不能使用。</b></p>
</div>

```
ps -ef | grep oracle| grep LOCAL=NO|grep -v grep|awk '{print $2}'|xargs -i kill -9 {}
```