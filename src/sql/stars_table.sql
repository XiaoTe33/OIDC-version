-- `stars`
create database if not exists winter_examination_database;
use winter_examination_database;
drop table if exists stars;
create table if not exists stars
(
    star_id bigint auto_increment primary key not null comment '主键',
    star_user_id bigint comment '用户',
    star_goods_id bigint comment '收藏的商品'
)comment 'stars';