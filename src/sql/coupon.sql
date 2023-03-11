create database if not exists winter_examination_database;
use winter_examination_database;
drop table if exists coupons;
create table if not exists coupons
(
    c_id bigint not null unique comment '主键',
    c_shop_id bigint not null comment '发放商id',
    c_name varchar(60) not null comment '优惠券名称',
    c_kind tinyint not null comment '优惠券种类',
    c_amount int not null comment '数量',
    c_discount varchar(60) not null comment '折扣详情',
    c_begin_at varchar(40) not null comment '开始时间',
    c_end_at varchar(40) not null comment '结束时间'
)