create database if not exists winter_examination_database;
use winter_examination_database;
drop table if exists password_protect;
create table if not exists password_protect
(
    uid bigint unique not null comment '用户',
    question varchar(100) not null comment '问题',
    answer varchar(100) not null comment '答案'
)