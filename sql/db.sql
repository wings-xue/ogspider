select count(*) from job;  --3866
select count(*) from zhaotoubiao; -- 3685
-- 181
delete from job;
delete from zhaotoubiao;

select count(*) from job where url ~ 'http.*?search.*?';  -- 181
select * from job where uuid = 'c438fda8011995484f9242fff095288c';


SELECT url FROM job
EXCEPT
SELECT req_id FROM zhaotoubiao;

select * from job where url = 'http://www.chinabidding.cn/zbgg/n5-Fca.html'; 
select status, count(status) from job group by status;
select * from job where status = 'retry';

select * from zhaotoubiao limit 1;



