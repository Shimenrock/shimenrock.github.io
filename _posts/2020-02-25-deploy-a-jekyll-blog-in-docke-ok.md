---
title: "如何用Docker部署Jekyll Blog"
published: true
categories: 
  - jekyll
tags: 
  - Docker
---

官方网站 https://github.com/envygeeks/jekyll-docker

- jekyll/jekyll: Default image.
- jekyll/minimal: Very minimal image.
- jekyll/builder: Includes tools.

```
 # docker search jekyll
NAME                     DESCRIPTION                                     STARS            OFFICIAL            AUTOMATED
jekyll/jekyll            Official Jekyll Docker Image                    255                                     [OK]

 # docker pull jekyll/jekyll

 # docker images
REPOSITORY                           TAG                 IMAGE ID            CREATED             SIZE
jekyll/jekyll                        latest              2754bfa01869        4 months ago        547MB

 # docker inspect jekyll/jekyll|grep JEKYLL
                "JEKYLL_VAR_DIR=/var/jekyll",
                "JEKYLL_DOCKER_TAG=4.0.0",
                "JEKYLL_VERSION=4.0.0",
                "JEKYLL_DOCKER_COMMIT=0d12d4bc90b266ae4a38dfce9a511c52dd6f0311",
                "JEKYLL_DOCKER_NAME=jekyll",
                "JEKYLL_DATA_DIR=/srv/jekyll",
                "JEKYLL_BIN=/usr/jekyll/bin",
                "JEKYLL_ENV=development",
                "JEKYLL_VAR_DIR=/var/jekyll",
                "JEKYLL_DOCKER_TAG=4.0.0",
                "JEKYLL_VERSION=4.0.0",
                "JEKYLL_DOCKER_COMMIT=0d12d4bc90b266ae4a38dfce9a511c52dd6f0311",
                "JEKYLL_DOCKER_NAME=jekyll",
                "JEKYLL_DATA_DIR=/srv/jekyll",
                "JEKYLL_BIN=/usr/jekyll/bin",
                "JEKYLL_ENV=development",
```
将自己的blog代码上传服务器，并赋权限
```
# chown -R 1000:1000 shimenrock-dba-script

# cd shimenrock-dba-script

# docker run --name blog --mount type=bind,source=$(pwd),target=/srv/jekyll -p 4000:4000  -it jekyll/jekyll jekyll serve
Warning: the running version of Bundler (2.0.2) is older than the version that created the lockfile (2.1.4). We suggest you upgrade to the latest version of Bundler by running `gem install bundler`.
Fetching gem metadata from https://rubygems.org/...........
Fetching gem metadata from https://rubygems.org/.
Resolving dependencies...
Fetching rake 12.3.0
Installing rake 12.3.0
Fetching RedCloth 4.3.2
Installing RedCloth 4.3.2 with native extensions
Fetching public_suffix 4.0.3
Installing public_suffix 4.0.3
Fetching addressable 2.7.0
Installing addressable 2.7.0
Using bundler 2.0.2
Fetching colorator 1.1.0
Installing colorator 1.1.0
Fetching concurrent-ruby 1.1.5
Installing concurrent-ruby 1.1.5
Fetching eventmachine 1.2.7
Installing eventmachine 1.2.7 with native extensions
Fetching http_parser.rb 0.6.0
Installing http_parser.rb 0.6.0 with native extensions
Fetching em-websocket 0.5.1
Installing em-websocket 0.5.1
Fetching ffi 1.12.1
Installing ffi 1.12.1 with native extensions
Fetching forwardable-extended 2.6.0
Installing forwardable-extended 2.6.0
Fetching i18n 1.8.2
Installing i18n 1.8.2
Fetching sassc 2.2.1
Installing sassc 2.2.1 with native extensions
Fetching jekyll-sass-converter 2.0.1
Installing jekyll-sass-converter 2.0.1
Fetching rb-fsevent 0.10.3
Installing rb-fsevent 0.10.3
Fetching rb-inotify 0.10.1
Installing rb-inotify 0.10.1
Fetching listen 3.2.1
Installing listen 3.2.1
Fetching jekyll-watch 2.2.1
Installing jekyll-watch 2.2.1
Fetching kramdown 2.1.0
Installing kramdown 2.1.0
Fetching kramdown-parser-gfm 1.1.0
Installing kramdown-parser-gfm 1.1.0
Fetching liquid 4.0.3
Installing liquid 4.0.3
Fetching mercenary 0.3.6
Installing mercenary 0.3.6
Fetching pathutil 0.16.2
Installing pathutil 0.16.2
Fetching rouge 3.15.0
Installing rouge 3.15.0
Fetching safe_yaml 1.0.5
Installing safe_yaml 1.0.5
Fetching unicode-display_width 1.6.1
Installing unicode-display_width 1.6.1
Fetching terminal-table 1.8.0
Installing terminal-table 1.8.0
Fetching jekyll 4.0.0
Installing jekyll 4.0.0
Fetching jekyll-seo-tag 2.6.1
Installing jekyll-seo-tag 2.6.1
Fetching jekyll-textile-converter 0.1.0
Installing jekyll-textile-converter 0.1.0
Using jekyll-theme-console 0.3.7 from source at `.`
Bundle complete! 5 Gemfile dependencies, 32 gems now installed.
Bundled gems are installed into `/usr/local/bundle`
Warning: the running version of Bundler (2.0.2) is older than the version that created the lockfile (2.1.4). We suggest you upgrade to the latest version of Bundler by running `gem install bundler`.
ruby 2.6.5p114 (2019-10-01 revision 67812) [x86_64-linux-musl]
Configuration file: /srv/jekyll/_config.yml
            Source: /srv/jekyll
       Destination: /srv/jekyll/_site
 Incremental build: disabled. Enable with --incremental
      Generating... 
                    done in 0.353 seconds.
 Auto-regeneration: enabled for '/srv/jekyll'
    Server address: http://0.0.0.0:4000
  Server running... press ctrl-c to stop.
```
完成软件更新后，退出

启动容器
```
# docker start blog

# curl http://0.0.0.0:4000
<!DOCTYPE html>
<html lang="en"><head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no" />
  <title>/</title><!-- Begin Jekyll SEO tag v2.6.1 -->
<meta name="generator" content="Jekyll v4.0.0" />
```
删除容器
```
# docker rm -f blog
```
进入容器
```
# docker exec -ti blog /bin/sh
```