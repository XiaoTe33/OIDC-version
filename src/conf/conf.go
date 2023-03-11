package conf

const (
	ClientId     = "115ba89949f592d08082"
	ClientSecret = "a87619ccc43d4f5db73abe479c74eec4c40af896"
	RedirectUri  = "http://localhost:9090/githubLogin/login"
)

// 本地配置
const (
	FrontEndPort = ":9091"
	BackEndPort  = ":9090"

	Database = "mysql"
	MysqlDNS = "root:root@tcp(127.0.0.1:3306)/winter_examination_database"

	JWTLastTime = 3600 //jwt持续时间（秒）

	OKMsg = "200" //成功返回消息

	LocalSavePathOfGoodsPictures      = "./src/static/goods/pictures/"      //商品图片保存路径
	LocalSavePathOfUserPhotos         = "./src/static/user/photos/"         //用户头像保存地址
	LocalSavePathOfEvaluationPictures = "./src/static/evaluation/pictures/" //评价图片保存地址
	LocalSavePathOfQR                 = "./src/static/qr/"                  //二维码保存地址

	OrderIdLeftShiftNumber = 32 //生成全局OrderId左移位数
	GoodsIdLeftShiftNumber = 20 //生成全局GoodsId左移位数

	OrderIdBaseTimeStamp int64 = 1672531200 //2023-1-1 00:00:00
	GoodsIdBaseTimeStamp int64 = 1672531200 //2023-1-1 00:00:00

)

const (
	IP = "http://localhost"

	GinPathOfUserPhoto          = "user/photo/" //路由地址
	GinPathOfGoodsPicture       = "goods/picture/"
	GinPathOfEvaluationPictures = "evaluation/picture/"
	GinPathOfQR                 = "qr/"

	WebLinkPathOfUserPhoto          = IP + FrontEndPort + "/" + GinPathOfUserPhoto //链接地址
	WebLinkPathOfGoodsPicture       = IP + FrontEndPort + "/" + GinPathOfGoodsPicture
	WebLinkPathOfEvaluationPictures = IP + FrontEndPort + "/" + GinPathOfEvaluationPictures
	WebLinkPathOfQR                 = IP + FrontEndPort + "/" + GinPathOfQR
)

// 部署配置
//const (
//	FrontEndPort = ":9091"
//	BackEndPort  = ":9090"
//
//	Database = "mysql"
//	MysqlDNS = "root:root@tcp(127.0.0.1:3307)/winter_examination_database"
//
//	JWTLastTime = 3600 //jwt持续时间（秒）
//
//	OKMsg = "200" //成功返回消息
//
//	LocalSavePathOfGoodsPictures      = "./src/static/goods/pictures/"      //商品图片保存路径
//	LocalSavePathOfUserPhotos         = "./src/static/user/photos/"         //用户头像保存地址
//	LocalSavePathOfEvaluationPictures = "./src/static/evaluation/pictures/" //评价图片保存地址
//	LocalSavePathOfQR                 = "./src/static/qr/"                  //二维码保存地址
//
//	OrderIdLeftShiftNumber = 32 //生成全局OrderId左移位数
//	GoodsIdLeftShiftNumber = 20 //生成全局GoodsId左移位数
//
//	OrderIdBaseTimeStamp int64 = 1672531200 //2023-1-1 00:00:00
//	GoodsIdBaseTimeStamp int64 = 1672531200 //2023-1-1 00:00:00
//)
//
//const (
//	IP = "http://39.101.72.18"
//
//	GinPathOfUserPhoto          = "user/photo/" //路由地址
//	GinPathOfGoodsPicture       = "goods/picture/"
//	GinPathOfEvaluationPictures = "evaluation/picture/"
//	GinPathOfQR                 = "qr/"
//
//	WebLinkPathOfUserPhoto          = IP + FrontEndPort + "/" + GinPathOfUserPhoto //链接地址
//	WebLinkPathOfGoodsPicture       = IP + FrontEndPort + "/" + GinPathOfGoodsPicture
//	WebLinkPathOfEvaluationPictures = IP + FrontEndPort + "/" + GinPathOfEvaluationPictures
//	WebLinkPathOfQR                 = IP + FrontEndPort + "/" + GinPathOfQR
//)
