-- `goods`
create database if not exists winter_examination_database;
use winter_examination_database;
drop table if exists goods;
create table if not exists goods
(
    goods_id bigint default 0 not null comment '主键' ,
    goods_is_deleted tinyint default 0 not null comment '是否删除(0-未删, 1-已删)',
    goods_name varchar(256) not null comment '商品名',
    goods_amount int not null default 0 comment '库存',
    good_is_star varchar(1) default '0' comment '是否被收藏',
    goods_kind varchar(256) default '其他' not null comment '商品种类',
    goods_price float not null comment '商品价格',
    goods_sold_amount bigint default 0 not null comment '销量',
    goods_score varchar(20) default '5.0' not null comment '商品评分',
    goods_shop_id varchar(20)default '' not null comment '所属店铺',
    goods_picture_link varchar(256) default '' not null comment '跳转链接'
) comment 'goods';
