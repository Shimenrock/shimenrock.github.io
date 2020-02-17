---
title: "页面配置文档：布局-配置标题图片"
header:
  overlay_image: /assets/images/unsplash-image-1.jpg
  caption: "Photo credit: [**Unsplash**](https://unsplash.com)"
  actions:
    - label: "Learn more"
      url: "https://unsplash.com"
categories:
  - Layout
tags:
  - edge case
  - image
  - layout
last_modified_at: 2018-03-20T16:00:52-04:00
---

在帖子上显示一个覆盖图像的标题
 
## 重叠滤镜

可以指定黑色叠加层的不透明度（值在0到1之间）

![transparent black overlay]({{ "/assets/images/mm-header-overlay-black-filter.jpg" | relative_url }})

```yaml
excerpt: "This post should [...]"
header:
  overlay_image: /assets/images/unsplash-image-1.jpg
  overlay_filter: 0.5 # same as adding an opacity of 0.5 to a black background
  caption: "Photo credit: [**Unsplash**](https://unsplash.com)"
  actions:
    - label: "More Info"
      url: "https://unsplash.com"
```

或者指定一个RGBA的滤镜层:

![transparent red overlay]({{ "/assets/images/mm-header-overlay-red-filter.jpg" | relative_url }})

```yaml
excerpt: "This post should [...]"
header:
  overlay_image: /assets/images/unsplash-image-1.jpg
  overlay_filter: rgba(255, 0, 0, 0.5)
  caption: "Photo credit: [**Unsplash**](https://unsplash.com)"
  actions:
    - label: "More Info"
      url: "https://unsplash.com"
```

或者指定纯色背景：
```yaml
title: "Layout: Header Overlay with Background Fill"
header:
  overlay_color: "#333"
```