-- `shops`
create database if not exists winter_examination_database;
use winter_examination_database;
drop table if exists shops;
create table if not exists shops
(
    shop_id bigint not null auto_increment comment '主键' primary key,
    shop_owner_id bigint not null comment '店长',
    shop_name varchar(30) not null comment '商店名',
    shop_is_delete varchar(3) default '0' not null comment '是否删除',
    shop_notice varchar(1024) default '' comment '店铺公告'
) comment 'shops';
