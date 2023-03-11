package app

import (
	"time"

	"winter-examination/src/conf"
	"winter-examination/src/utils"

	"github.com/gin-gonic/gin"
)

func InitRouters() {
	utils.InitCheckouts()
	go initBackEndRouters()
	gin.SetMode(gin.ReleaseMode)
	time.Sleep(1 * time.Second)
	go initFrontEndRouters()
	forever := make(chan int)
	<-forever
}

func initBackEndRouters() {

	r := gin.Default()
	r.MaxMultipartMemory = 3 << 20
	r.Use(Cors())
	r.LoadHTMLGlob("./src/static/github_login_page/*")
	u := r.Group("/user") //用户模块

	{
		{
			u.POST("/register", Register)      //注册新用户
			u.POST("/login", Login)            //用户名密码登录
			u.POST("/login/token", TokenLogin) //token登录
		}
		{
			u.PUT("/info", JWT(), UpdateMyInfo)     //修改我的基本信息
			u.GET("/info", JWT(), QueryMyInfo)      //获取我的基本信息
			u.GET("/info/other", QueryUserInfo)     //获取别人的基本信息
			u.POST("/photo", JWT(), AddUserPhoto)   //上传头像
			u.PUT("/money/:money", JWT(), AddMoney) //充值
		}

		{
			u.POST("/address", JWT(), AddUserAddress)      //新增收货地址
			u.GET("/address", JWT(), QueryMyAddressList)   //获取我的地址列表
			u.DELETE("/address/:id", JWT(), DeleteAddress) //删除地址
		}
		{
			u.GET("/car", JWT(), MyShoppingCarList)               //我的购物车
			u.PUT("/car/:goodsId", JWT(), AddToShoppingCar)       //加入购物车
			u.DELETE("/car/:numId", JWT(), DeleteFromShoppingCar) //移出购物车
		}
		{
			u.PUT("/star/:goodsId", JWT(), AddStar)       //收藏商品
			u.DELETE("/star/:goodsId", JWT(), DeleteStar) //取消收藏
			u.GET("/star", JWT(), QueryMyStars)           //我的收藏
		}

		u.PUT("/coupon/:couponId", JWT(), FetchCoupon) //抢购优惠券
		u.GET("/coupon", JWT(), MyCoupon)              //我的优惠券
		u.GET("/star/all", QueryAllStars)              //查看收藏表
		u.GET("/all", QueryAllUsers)                   //便于前端查表
		{
			u.PUT("/password", JWT(), UpdatePassword)           //修改密码
			u.POST("/password/protect", JWT(), AddPwdProtect)   //设置密保
			u.PUT("/password/protect", JudgePwdProtect)         //校验答案
			u.GET("/password/protect", QueryPwdProtectQuestion) //获取密保问题
		}
	}

	s := r.Group("/shop") //商店模块
	{
		s.POST("", JWT(), AddShop)                //新增商店
		s.DELETE("", JWT(), DeleteShop)           //删除商店
		s.PUT("/info", JWT(), UpdateShopInfo)     //更新商店信息
		s.PUT("/notice", JWT(), UpdateShopNotice) //修改商店公告
		s.GET("/info", JWT(), MyShopInfo)         //我的商店信息
		s.GET("/order", JWT(), MyShopOrders)      //本店订单
		s.GET("/all", QueryAllShops)              //查看当前所有商店信息
		s.GET("/info/other", QueryShop)           //id查商店

		g := s.Group("/goods") //商品模块(商家)
		{
			g.POST("", JWT(), AddGoods)                           //新增商品
			g.PUT("", JWT(), UpdateGoods)                         //修改商品信息
			g.PUT("/add/:goodsId/:amount", JWT(), AddGoodsAmount) //补货
			g.PUT("/cut/:goodsId/:amount", JWT(), CutGoodsAmount) //卸货
			g.DELETE("/:goodsId", JWT(), DeleteGoods)             //下架商品
			g.GET("", JWT(), MyShopGoods)                         //我的商品
			g.GET("/all", QueryAllGoodsWithoutMode)               //查看所有商品信息
		}
	}

	{
		r.GET("/goods", QueryGoods) //查找商品
	}

	c := r.Group("/coupon") //优惠券模块(商家)
	{
		c.POST("", JWT(), AddCoupon) //发放优惠券
		c.GET("/all", QueryAllCoupons)
	}

	o := r.Group("/order") //订单模块
	{
		o.POST("", JWT(), AddOrder)                  //新增订单
		o.PUT("/status", JWT(), UpdateOrderStatus)   //修改订单状态
		o.PUT("/address", JWT(), UpdateOrderAddress) //修改订单地址
		o.GET("", JWT(), MyOrder)                    //我的订单
		o.GET("/id/:id", QueryOrdersById)            //订单号查询订单
		o.GET("/shopId/:shopId", QueryOrderByShopId) //商店id查询订单
		o.GET("/all", QueryAllOrders)                //查看所有订单
	}

	e := r.Group("/evaluation") //评价模块
	{
		e.POST("", JWT(), AddEvaluation)           //新增评价
		e.DELETE("/:id", JWT(), DeleteEvaluations) //删除评价
		e.GET("", QueryGoodsEvaluations)           //查询某个商品的评价
		e.GET("/all", QueryAllEvaluations)         //查看所有
	}

	q := r.Group("/qr") //扫码
	{
		q.GET("", GetQR)                   //获取二维码
		q.GET("/key/:key", JudgeQR)        //用户扫码访问地址
		q.GET("/status/:key", GetQRStatus) //轮询地址
	}

	r.GET("/RESTART/:yes", utils.Restart) //数据重开

	gl := r.Group("/githubLogin")
	gl.GET("/homePage", homePage)
	gl.GET("/login", githubLogin)

	oidc := r.Group("/oidc")
	{
		oidc.GET("/GetClientId", JWT(), GetClientId)
		oidc.GET("/GetClientSecret", JWT(), GetClientSecret)
		oidc.GET("/GetCode", GetCode)
		oidc.GET("/JudgeCode", JudgeCode)
		oidc.GET("/JudgeAccessToken", JudgeAccessToken)

	}

	_ = r.Run(conf.BackEndPort)
}

func initFrontEndRouters() {
	r := gin.Default()
	r.Use(Cors())
	r.StaticFile("/favicon.ico", "./src/static/favicon/favicon.ico")
	r.Static("js/", "./templates/js")
	r.Static("css/", "./templates/css")
	r.Static("images/", "./templates/images")
	r.Static(conf.GinPathOfQR, "./src/static/qr")
	r.Static(conf.GinPathOfUserPhoto, "./src/static/user/photos")
	r.Static(conf.GinPathOfGoodsPicture, "./src/static/goods/pictures")
	r.Static(conf.GinPathOfEvaluationPictures, "./src/static/evaluation/pictures")
	//r.LoadHTMLGlob("templates/html/*")
	//
	//u := r.Group("/user")
	//u.Use(TokenMiddleware())
	//u.GET("/register", PageRegister)
	//u.GET("/login", PageLogin)
	//u.GET("/main", PageMain)

	_ = r.Run(conf.FrontEndPort)
}
