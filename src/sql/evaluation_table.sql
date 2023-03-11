-- `evaluations`
create database if not exists winter_examination_database;
use winter_examination_database;
drop table if exists evaluations;
create table if not exists evaluations
(
    e_id bigint not null auto_increment primary key comment '主键',
    e_u_id bigint not null comment '评价者',
    e_g_id bigint not null comment '评价的商品',
    e_text text not null comment '评价内容',
    e_score float not null comment '评分',
    e_picture varchar(256) not null default '' comment '图片',
    e_time varchar(30) not null comment '评价时间',
    e_is_deleted varchar(1) default '0' comment '是否删除'
)comment 'evaluations';