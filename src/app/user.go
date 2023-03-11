package app

import (
	"winter-examination/src/conf"
	"winter-examination/src/model"
	"winter-examination/src/service"
	"winter-examination/src/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user = model.UserLoginReq{}
	err := c.ShouldBind(&user)
	if handleBindingError(c, err, &user) {
		return
	}

	if handleError(c, service.Login(user)) {
		return
	}
	jsonToken(c, user.Username)

}
func Register(c *gin.Context) {
	var user model.UserRegisterReq
	err := c.ShouldBind(&user)
	if handleBindingError(c, err, &user) {
		return
	}
	service.Register(user)
	jsonSuccess(c)
}

//username := c.PostForm("username")
//password := c.PostForm("password")
//password2 := c.PostForm("password2")
//email := c.PostForm("email")
//phone := c.PostForm("phone")
//fmt.Println(username, password, password2, email, phone)
//msg := service.Register(username, password, password2, email, phone)
//c.JSON(200, gin.H{
//	"msg": msg,
//})

func TokenLogin(c *gin.Context) {
	token := c.PostForm("token")
	if !utils.IsValidJWT(token) {
		jsonError(c, "token无效")
		return
	}
	jsonToken(c, utils.GetUsernameByToken(token))
}

func QueryMyInfo(c *gin.Context) {
	userId := c.GetString("userId")
	data := service.QueryMyInfo(userId)
	jsonData(c, data)
}

func QueryUserInfo(c *gin.Context) {
	id := c.Query("id")
	username := c.Query("username")
	phone := c.Query("phone")
	email := c.Query("email")
	msg, user := service.QueryUserInfo(id, username, phone, email)
	if msg != conf.OKMsg {
		jsonError(c, msg)
	} else {
		jsonData(c, user)
	}

	//token := c.PostForm("token")
	//id := c.PostForm("id")
	//username := c.PostForm("username")
	//phone := c.PostForm("phone")
	//email := c.PostForm("email")
	//msg, user := service.QueryUser(token, id, username, phone, email)
	//if msg != conf.OKMsg {
	//	c.JSON(200, gin.H{
	//		"msg": msg,
	//	})
	//	return
	//}
	//if token != "" {
	//	c.JSON(200, gin.H{
	//		"msg":             msg,
	//		"data":            user,
	//		"refreshed_token": utils.RefreshToken(token),
	//	})
	//	return
	//}
	//c.JSON(200, gin.H{
	//	"msg":  conf.OKMsg,
	//	"data": user,
	//})
}

func UpdateMyInfo(c *gin.Context) {
	userId := c.GetString("userId")
	var req = model.MyInfoUpdateReq{}
	err := c.ShouldBind(&req)
	if handleBindingError(c, err, &req) {
		return
	}
	service.UpdateUserInfo(userId, req)
	jsonSuccess(c)
}

//func UpdateUser(c *gin.Context) {
//	token := c.PostForm("token")
//	password := c.PostForm("password")
//	phone := c.PostForm("phone")
//	email := c.PostForm("email")
//	newUsername := c.PostForm("username")
//	username := utils.GetUsernameByToken(token)
//	msg := service.UpdateUser(username, newUsername, password, phone, email)
//	if newUsername != "" {
//		c.JSON(200, gin.H{
//			"msg":             msg,
//			"refreshed_token": utils.CreateJWT(newUsername),
//		})
//		return
//	}
//	c.JSON(200, gin.H{
//		"msg":             msg,
//		"refreshed_token": utils.RefreshToken(token),
//	})
//}

func QueryAllUsers(c *gin.Context) {
	users := service.QueryAllUsers()
	jsonData(c, users)
}

func AddUserPhoto(c *gin.Context) {
	photo, err := c.FormFile("photo")
	if handleError(c, err) {
		return
	}
	ok, style := utils.IsValidPictureFile(photo.Filename)
	if !ok {
		jsonError(c, "仅支持jpg,png,jfif格式的图片")
		return
	}
	err = c.SaveUploadedFile(photo, conf.LocalSavePathOfUserPhotos+c.GetString("userId")+style)
	if handleError(c, err) {
		return
	}
	service.AddUserPhoto(c.GetString("userId"))
	jsonSuccess(c)
}

//func AddUserPhoto(c *gin.Context) {
//	token := c.PostForm("token")
//	photo, err := c.FormFile("photo")
//	if err != nil {
//		c.JSON(200, gin.H{
//			"msg":             "解析文件出错",
//			"refreshed_token": utils.RefreshToken(token),
//		})
//		return
//	}
//	ok, style := utils.IsValidPictureFile(photo.Filename)
//	if !ok {
//		c.JSON(200, gin.H{
//			"msg":             "仅支持jpg,png,jfif格式的图片",
//			"refreshed_token": utils.RefreshToken(token),
//		})
//		return
//	}
//	err = c.SaveUploadedFile(photo, conf.LocalSavePathOfUserPhotos+utils.GetUserIdByToken(token)+style)
//	if err != nil {
//		c.JSON(200, gin.H{
//			"msg":             "保存文件出错",
//			"refreshed_token": utils.RefreshToken(token),
//		})
//		return
//	}
//	msg := service.AddUserPhoto(token)
//	c.JSON(200, gin.H{
//		"msg":             msg,
//		"refreshed_token": utils.RefreshToken(token),
//	})
//}

func QueryMyPhoto(c *gin.Context) {
	userId := c.GetString("userId")
	photo := service.QueryMyPhoto(userId)
	jsonData(c, gin.H{
		"photo": photo,
	})
}

func QueryPhoto(c *gin.Context) {
	userId := c.Param("id")
	photo := service.QueryMyPhoto(userId)
	jsonData(c, gin.H{
		"photo": photo,
	})
}

func AddUserAddress(c *gin.Context) {
	var addr = model.UserAddressReq{}
	userId := c.GetString("userId")
	err := c.ShouldBind(&addr)
	if handleBindingError(c, err, &addr) {
		return
	}
	err = service.AddUserAddress(userId, addr)
	if handleError(c, err) {
		return
	}
	jsonSuccess(c)
}

func QueryMyAddressList(c *gin.Context) {
	userId := c.GetString("userId")
	data, err := service.QueryMyAddressList(userId)
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}

func DeleteAddress(c *gin.Context) {
	userId := c.GetString("userId")
	id := c.Param("id")
	err := service.DeleteAddress(userId, id)
	if handleError(c, err) {
		return
	}
	jsonSuccess(c)

}

func AddToShoppingCar(c *gin.Context) {
	userId := c.GetString("userId")
	goodsId := c.Param("goodsId")
	err := service.AddToShoppingCar(userId, goodsId)
	if handleError(c, err) {
		return
	}
	jsonSuccess(c)
}

func DeleteFromShoppingCar(c *gin.Context) {
	userId := c.GetString("userId")
	numId := c.Param("numId")
	err := service.DeleteFromShoppingCar(userId, numId)
	if handleError(c, err) {
		return
	}
	jsonSuccess(c)
}

func MyShoppingCarList(c *gin.Context) {
	userId := c.GetString("userId")
	data, err := service.ShowShoppingCarList(userId)
	if handleError(c, err) {
		return
	}
	jsonData(c, data)

}

func AddMoney(c *gin.Context) {
	userId := c.GetString("userId")
	var req = model.AddMoneyReq{}
	if handleBindingError(c, c.ShouldBindUri(&req), &req) {
		return
	}
	service.AddMoney(req, userId)
	jsonSuccess(c)

}

func UpdatePassword(c *gin.Context) {
	userId := c.GetString("userId")
	password := c.PostForm("password")
	newPassword := c.PostForm("newPassword")
	if handleError(c, service.UpdatePassword(userId, password, newPassword)) {
		return
	}
	jsonSuccess(c)
}
