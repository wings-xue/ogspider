## Spider 流程


1. 第一次执行spider流程，包含以下步骤
   1. 创建BaseSpider对象
   2. 加载Name
   3. 加载Host
   4. 加载Fields
   5. 加载StartURL(StartURL列表或者func) 
   6. 加载Setting
   7. 加载默认Middlerware
   8. 加载表创建函数

   Scrper OpenSpider
   1. 注册爬虫， 收集爬虫配置
   1.  判断是否加载需要的字段
   2.  表创建
   3.  初始化startReq，并且传入engine

2. Scraper Process, engine传入response到spider，包含以下步骤
   1. find Response 对应的 spider
   2. for循环执行middleware， 产生request和response（和scrapy不同点在于，所有状态都写入response和request交给pipeline处理）
   3. 新产生的request交付engine
   4. response处理
      1. 成功 Statuecode=200
      2. oldReq 状态修改
    

   



## Spider Middleware
#### 参考
```
https://docs.scrapy.org/en/latest/topics/spider-middleware.html#writing-your-own-spider-middleware
```

#### 主要功能
1. 判断response是否出现错误 （https://docs.scrapy.org/en/latest/topics/spider-middleware.html#scrapy.spidermiddlewares.SpiderMiddleware.process_spider_input）
2. 拆分request （https://docs.scrapy.org/en/latest/topics/spider-middleware.html#scrapy.spidermiddlewares.SpiderMiddleware.process_spider_output）
