
本地运行
    bundle exec jekyll serve
谷歌分析配置
https://analytics.google.com/analytics
管理>>媒体管理>>创建媒体资源
这是针对此媒体资源的全局网站代码 (gtag.js) 跟踪代码。请复制此代码，并将其作为第一个项目粘贴到您要跟踪的每个网页的 <HEAD> 标记中。如果您的网页上已经有全局网站代码，则只需将以下代码段中的 config 行添加到现有的全局网站代码。
<!-- Global site tag (gtag.js) - Google Analytics -->
<script async src="https://www.googletagmanager.com/gtag/js?id=UA-135881568-2"></script>
<script>
  window.dataLayer = window.dataLayer || [];
  function gtag(){dataLayer.push(arguments);}
  gtag('js', new Date());

  gtag('config', 'UA-135881568-2');
</script>