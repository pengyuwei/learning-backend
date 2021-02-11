select date_format(now(),'%Y-%m-%d %H:%I:%s %u %W'); 
select now(),current_timestamp(),localtime(),localtimestamp(),sysdate(),curdate(),curtime();
select utc_timestamp(),utc_date(),utc_time();

set @dt = '2021-2-11 18:16:00.123456';

select date(@dt); -- 2021-02-11
select time(@dt); -- 18:16:00.123456
select year(@dt); -- 2021
select quarter(@dt); -- 1
select month(@dt); -- 2
select week(@dt); -- 6
select day(@dt); -- 11
select hour(@dt); -- 18
select minute(@dt); -- 16
select second(@dt); -- 0
select microsecond(@dt); -- 123456

select extract(year from @dt); -- 2021
select extract(quarter from @dt); -- 1
select extract(month from @dt); -- 2
select extract(week from @dt); -- 6
select extract(day from @dt); -- 11
select extract(hour from @dt); -- 18
select extract(minute from @dt); -- 16
select extract(second from @dt); -- 0
select extract(microsecond from @dt); -- 123456
select extract(year_month from @dt); -- 202102
select extract(day_hour from @dt); -- 1118
select extract(day_minute from @dt); -- 111816
select extract(day_second from @dt); -- 1181600
select extract(day_microsecond from @dt); -- 11181600123456
select extract(hour_minute from @dt); -- 1816
select extract(hour_second from @dt); -- 181600
select extract(hour_microsecond from @dt); -- 181600123456
select extract(minute_second from @dt); -- 1600
select extract(minute_microsecond from @dt); -- 1600123456
select extract(second_microsecond from @dt); -- 123456

select dayofweek(@dt); -- 5
select dayofmonth(@dt);  -- 11, 几月
select dayofyear(@dt);  -- 42

select week(@dt); -- 6, 第几周
select week(@dt,3);  -- 6, 三天后是第几周
select weekofyear(@dt);  -- 6, 一年中的第几周
select dayofweek(@dt);  -- 5, “某天”在一周中的位置, （1 =Sunday, 2 = Monday,…, 7 = Saturday）
select weekday(@dt);  -- 3, (0 =Monday, 1 = Tuesday, …, 6 = Sunday)；
select yearweek(@dt);  -- 202106, 那年的第几周

select dayname(@dt);  -- Thursday
select monthname(@dt);  -- February

select last_day(@dt); -- 2021-02-28, 返回月份中的最后一天
select day (last_day(now())) as days; -- 28, 获取本月有多少天

set @dt=now();

select date_add(@dt, interval 1 day ); -- 加一天
select date_add(@dt, interval 1 hour ); -- 加一小时
select date_add(@dt, interval 1 minute ); -- 加一分
select date_add(@dt, interval 1 second ); -- 加一秒
select date_add(@dt, interval 1 microsecond); -- 加一毫秒
select date_add(@dt, interval 1 week); -- 加一周
select date_add(@dt, interval 1 month ); -- 加一月
select date_add(@dt, interval 1 quarter); -- 加一季度
select date_add(@dt, interval 1 year );  -- 加一年
select date_add(@dt, interval -1 day ); -- 减一天

select date_add(@dt, interval '01:01:01' hour_second); -- 加01:01:01
select date_add(@dt, interval '1 01:02:03' day_second); -- 加1天01:02:03

select date_sub(@dt, interval '1 1:1:1' day_second); -- 减去1天1:1:1

-- 函数参数“P” 的格式为“YYYYMM” 或者 “YYMM”，第二个参数“N” 表示增加或减去 N month（月）
select period_add(202102,2), period_add(20210202,-2); -- 202104, 20210112,日期加/减去2月

select datediff( '2021-02-18' , '2021-02-11' ); -- 7 返回天数
select timediff( '2021-02-11 08:08:08' , '2021-02-11 00:00:00' ); -- 08:08:08 返回 time 差值

select time_to_sec( '01:00:05' ); -- 3605
select sec_to_time(3605); -- 01:00:05

select to_days( '0000-01-01' ); -- 1
select to_days( '2021-02-11' ); -- '738197'
select from_days(0); -- '0000-00-00'
select from_days(733627); -- '2008-08-08'

select str_to_date( '02/11/2021' , '%m/%d/%Y' ); -- 2021-02-11

select date_format( '2008-08-08 22:23:00' , '%W %M %Y' ); -- Friday August 2008
select date_format('2008-08-08 22:23:01', '%Y%m%d%H%i%s'); -- 20080808222301
select time_format('22:23:01', '%H.%i.%s'); -- 22.23.01

-- 语法 ：get_format(date|time|datetime, 'eur'|'usa'|'jis'|'iso'|'internal'

select get_format( date , 'usa' ) ; -- '%m.%d.%Y'
select get_format( date , 'jis' ) ; -- '%Y-%m-%d'
select get_format( date , 'iso' ) ; -- '%Y-%m-%d'
select get_format( date , 'eur' ) ; -- '%d.%m.%Y'
select get_format( date , 'internal' ) ; -- '%Y%m%d'
select get_format(datetime, 'usa' ) ; -- '%Y-%m-%d %H.%i.%s'
select get_format(datetime, 'jis' ) ; -- '%Y-%m-%d %H:%i:%s'
select get_format(datetime, 'iso' ) ; -- '%Y-%m-%d %H:%i:%s'
select get_format(datetime, 'eur' ) ; -- '%Y-%m-%d %H.%i.%s'
select get_format(datetime, 'internal' ) ; -- '%Y%m%d%H%i%s'
select get_format( time , 'usa' ) ; -- '%h:%i:%s %p'
select get_format( time , 'jis' ) ; -- '%H:%i:%s'
select get_format( time , 'iso' ) ; -- '%H:%i:%s'
select get_format( time , 'eur' ) ; -- '%H.%i.%s'
select get_format( time , 'internal' ) ; -- '%H%i%s'

select makedate(2021,31); -- 2021-01-31
select makedate(2021,32); -- 2021-02-01
select maketime(23,10,30); -- 23:10:30
