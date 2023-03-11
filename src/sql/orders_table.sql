-- `orders`
create database if not exists winter_examination_database;
use winter_examination_database;
drop table if exists orders;
create table if not exists orders
(
    order_id bigint not null  comment '主键' unique ,
    order_buyer_id bigint not null comment '客户',
    order_solder_id bigint not null comment '商店',
    order_goods_id bigint not null comment '商品',
    order_goods_amount int default 1 comment '商品数量',
    order_goods_style varchar(256) default '[]' comment '商品款式',
    order_time varchar(50) not null comment '下单时间',
    order_address varchar(100) not null comment '收货地址',
    order_status varchar(1) default '0' not null comment '订单状态',
    order_discount varchar(40) default '' not null comment '折扣方式',
    order_origin_price float default 0.00 not null comment '原价',
    order_actual_price float default 0.00 not null comment '实付金额'
) comment 'orders';