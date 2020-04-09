---
title: "页面配置文档：布局-侧边栏自定义"
excerpt: "A post with custom sidebar content."
author_profile: false
categories: Layout
sidebar:
  - title: "Title"
    image: http://placehold.it/350x250
    image_alt: "image"
    text: "Some text here."
  - title: "Another Title"
    text: "More text here."
    nav: sidebar-sample
---

This post has a custom sidebar set in the post's YAML Front Matter.

An example of how that YAML could look is:

```yaml
sidebar:
  - title: "Title"
    image: http://placehold.it/350x250
    image_alt: "image"
    text: "Some text here."
  - title: "Another Title"
    text: "More text here."
    nav: sidebar-sample
```