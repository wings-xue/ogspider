# ogspider
爬虫平台

1. spider 从数据库获取job
2. spider 解析job 为request
3. spider将request通过chan传入engine


4. engine拿到request给scheduler   
5. scheduler调度request给engine
6. engine拿到request给download
7. download调用process下载获得response
8. download将response传入engine
9.  engine获取response给pipeline
10. pipeline 调用process处理response
11. pipeline 生成item，存入数据

pgsql install
 docker run --name ogpostgres -d -p 5432:5432 -e POSTGRES_PASSWORD= postgres


docker run -d -p 5433:80 --name pgadmin4 -e PGADMIN_DEFAULT_EMAIL=qq号码@qq.com -e PGADMIN_DEFAULT_PASSWORD=密码 dpage/pgadmin4