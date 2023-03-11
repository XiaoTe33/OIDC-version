-- `users`
create database if not exists winter_examination_database;
use winter_examination_database;
drop table if exists users;
create table if not exists users
(
    user_id bigint not null auto_increment comment 'id' primary key,
    username varchar(20) not null comment '用户名',
    password varchar(256) not null comment 'password',
    phone varchar(11) default '' comment '手机号',
    email varchar(256) default '' comment '邮箱',
    money varchar(256) default '0.00' not null comment '余额',
    photo varchar(256) default '' comment '头像',
    shopping_car varchar(4096) default '[]' comment '购物车',
    address varchar(4096) default '[]' comment '收货地址',
    coupon varchar(4096) default '[]' comment '优惠券'
) comment 'users';



