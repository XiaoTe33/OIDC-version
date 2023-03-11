package utils

import (
	"fmt"
	"os"

	"winter-examination/src/dao"

	"github.com/gin-gonic/gin"
)

// Restart 删库跑路！！！
func Restart(c *gin.Context) {
	key := c.Param("yes")
	if key == "lly" || key == "lyq" {
		c.JSON(200, gin.H{
			"msg": "Restarted success!",
		})
		_, err := dao.Db.Exec(`create database if not exists winter_examination_database`)
		_, err = dao.Db.Exec(`use winter_examination_database`)
		_, err = dao.Db.Exec(`drop table if exists goods`)
		_, err = dao.Db.Exec(`create table if not exists goods
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
) comment 'goods'`)
		_, err = dao.Db.Exec(`drop table if exists orders`)
		_, err = dao.Db.Exec(`create table if not exists orders
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
) comment 'orders'`)
		_, err = dao.Db.Exec(`drop table if exists evaluations`)
		_, err = dao.Db.Exec(`create table if not exists evaluations
(
    e_id bigint not null auto_increment primary key comment '主键',
    e_u_id bigint not null comment '评价者',
    e_g_id bigint not null comment '评价的商品',
    e_text text not null comment '评价内容',
    e_score float not null comment '评分',
    e_picture varchar(256) not null default '' comment '图片',
    e_time varchar(30) not null comment '评价时间',
    e_is_deleted varchar(1) default '0' comment '是否删除'
)comment 'evaluations'`)
		_, err = dao.Db.Exec(`drop table if exists shops`)
		_, err = dao.Db.Exec(`create table if not exists shops
(
    shop_id bigint not null auto_increment comment '主键' primary key,
    shop_owner_id bigint not null comment '店长',
    shop_name varchar(30) not null comment '商店名',
    shop_is_delete varchar(3) default '0' not null comment '是否删除',
    shop_notice varchar(1024) default '' comment '店铺公告'
) comment 'shops'`)
		_, err = dao.Db.Exec(`drop table if exists stars`)
		_, err = dao.Db.Exec(`create table if not exists stars
(
    star_id bigint auto_increment primary key not null comment '主键',
    star_user_id bigint comment '用户',
    star_goods_id bigint comment '收藏的商品'
)comment 'stars'`)
		_, err = dao.Db.Exec(`drop table if exists users`)
		_, err = dao.Db.Exec(`create table if not exists users
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
) comment 'users'`)
		_, err = dao.Db.Exec(`drop table if exists coupons`)
		_, err = dao.Db.Exec(`create table if not exists coupons
(
    c_id bigint not null unique comment '主键',
    c_shop_id bigint not null comment '发放商id',
    c_name varchar(60) not null comment '优惠券名称',
    c_kind tinyint not null comment '优惠券种类',
    c_amount int not null comment '数量',
    c_discount varchar(60) not null comment '折扣详情',
    c_begin_at varchar(40) not null comment '开始时间',
    c_end_at varchar(40) not null comment '结束时间'
)`)
		_, err = dao.Db.Exec(`drop table if exists password_protect`)
		_, err = dao.Db.Exec(`create table if not exists password_protect
(
    uid bigint unique not null comment '用户',
    question varchar(100) not null comment '问题',
    answer varchar(100) not null comment '答案'
)`)
		os.RemoveAll("./src/static/evaluation/pictures")
		os.RemoveAll("./src/static/goods/pictures")
		os.RemoveAll("./src/static/user/photos")
		os.RemoveAll("./src/static/qr")
		os.MkdirAll("./src/static/evaluation/pictures", 777)
		os.MkdirAll("./src/static/goods/pictures", 777)
		os.MkdirAll("./src/static/user/photos", 777)
		os.MkdirAll("./src/static/qr", 777)
		fmt.Println(err)
	} else {
		c.JSON(200, gin.H{
			"msg": "删库跑路也是要密码的",
		})
	}
}
