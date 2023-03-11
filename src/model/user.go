package model

type User struct {
	Id          string `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Money       string `json:"money"`
	Photo       string `json:"photo"`
	ShoppingCar string `json:"shoppingCar"`
	Address     string `json:"address"`
	Coupon      string `json:"coupon"`
}

type UserLoginReq struct {
	Username string `json:"username" form:"username" binding:"required,IsRegisteredUsername" err:"用户名不存在"`
	Password string `json:"password" form:"password" binding:"required"                      err:"请输入密码"`
}

type UserRegisterReq struct {
	Username   string `json:"username"   binding:"required,max=20,IsValidUsername"       form:"username"   err:"用户名过长或已被注册"`
	Password   string `json:"password"   binding:"required,min=6,max=20,IsValidPassword" form:"password"   err:"密码长度应为6~20"`
	RePassword string `json:"rePassword" binding:"required,eqfield=Password"             form:"rePassword" err:"两次密码不一致"`
	Email      string `json:"email"      binding:"required,IsValidEmail"                 form:"email"      err:"邮箱格式不正确或已被注册"`
	Phone      string `json:"phone"      binding:"required,IsValidPhone"                 form:"phone"      err:"手机号格式不正确或已被注册"`
}

type MyInfoUpdateReq struct {
	Username string `json:"username" binding:"required,max=20,IsValidUsername" form:"username" err:"用户名过长或已被注册"`
	Email    string `json:"email"    binding:"required,IsValidEmail"           form:"email"    err:"邮箱格式不正确或已被注册"`
	Phone    string `json:"phone"    binding:"required,IsValidPhone"           form:"phone"    err:"手机号格式不正确或已被注册"`
}

type MyInfoRsp struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Money    string `json:"money"`
	Photo    string `json:"photo"`
}

type UserInfoRsp struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Photo    string `json:"photo"`
}

type UserAddressReq struct {
	Province string `json:"province" form:"province" binding:"required" msg:"请填写省"`
	City     string `json:"city"     form:"city"     binding:"required" msg:"请填写市"`
	County   string `json:"county"   form:"county"   binding:"required" msg:"请填写区"`
	Detail   string `json:"detail"   form:"detail"   binding:"required" msg:"请填写详细地址"`
	Tag      string `json:"tag"      form:"tag"      binding:"required" msg:"请填写标签"`
}

type UserAddressRsp struct {
	Id       int    `json:"id"`
	Province string `json:"province"`
	City     string `json:"city"`
	County   string `json:"county"`
	Detail   string `json:"detail"`
	Tag      string `json:"tag"`
}

type ShoppingCarRsp struct {
	Id          int    `json:"id"`
	GoodsId     string `json:"goodsId"`
	Name        string `json:"name"`
	ShopName    string `json:"shopName"`
	Kind        string `json:"kind"`
	Price       string `json:"price"`
	SoldAmount  string `json:"soldAmount"`
	Score       string `json:"score"`
	PictureLink string `json:"pictureLink"`
}

type AddMoneyReq struct {
	Money string `json:"money" uri:"money" binding:"required,IsValidGoodsPrice" err:"充值价格格式不正确"`
}
