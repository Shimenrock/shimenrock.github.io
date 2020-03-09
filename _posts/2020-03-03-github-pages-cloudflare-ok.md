---
title: "通过CloudFlare为GitHub Page提供安全加速访问"
published: true
categories: 
  - github
tags: 
  - github
---
GitHub Page托管的都是静态HTML文件，不存在动态内容。

# 1.域名配置

官网：https://www.cloudflare.com/

配置自定义nameservers

- ERIC.NS.CLOUDFLARE.COM
- SUE.NS.CLOUDFLARE.COM

# 2.在git仓库中添加一个CNAME文件

```
echo "www.ju.je" > CNAME
git add -A
git commit -m “Added CNAME file.”
git push origin master
```

# 3.CLOUDFLARE配置DNS配置

添加两条record

| Type | Name | Content | TTL | Proxy status |
| --- | --- | --- | --- | --- |
| CNAME | www | xxx.github.io | Auto | Proxied |
| CNAME | ju.je | xxx.github.io | Auto | Proxied |

# 4.CLOUDFLARE配置SSL/TLS

选择 FULL 模式 SSL, 不要选择 FULL(Strict)

# 5.CLOUDFLARE添加Page Rules

primary Page Rule: 
- http://www.ju.je* 
- Always Use HTTPS

重定向Page Rule: 
- https://ju.je
- Forwording URL  --- 301 - Permanent Redirect
- https://www.ju.je 

# 6.CLOUDFLARE添加Page Rules的Cache

Cache Page Rule: 
- https://ju.je
- Cache Level
- Cache Everything



转载 ：https://blog.cloudflare.com/secure-and-fast-github-pages-with-cloudflare/