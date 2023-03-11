package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"winter-examination/src/conf"
	"winter-examination/src/dao"
	"winter-examination/src/model"
	"winter-examination/src/utils"
)

func Login(req model.UserLoginReq) error {
	u := dao.QueryUserByUsername(req.Username)
	if u.Password != utils.SHA256Secret(req.Password) {
		return errors.New("密码错误")
	}
	return nil
}

func Register(u model.UserRegisterReq) {
	dao.AddUser(model.User{
		Username: u.Username,
		Password: utils.SHA256Secret(u.Password),
		Phone:    u.Phone,
		Email:    u.Email,
	})
}

func QueryMyInfo(userId string) model.MyInfoRsp {
	u := dao.QueryUserById(userId)
	return model.MyInfoRsp{
		Username: u.Username,
		Phone:    u.Phone,
		Email:    u.Email,
		Money:    u.Money,
		Photo:    u.Photo,
	}
}

func QueryUserInfo(id string, username string, phone string, email string) (msg string, user model.UserInfoRsp) {
	var u = model.User{}
	if id != "" {
		if u = dao.QueryUserById(id); u != (model.User{}) {
			goto RETURN
		}
	}
	if username != "" {
		if u = dao.QueryUserByUsername(username); u != (model.User{}) {
			goto RETURN
		}
	}
	if phone != "" {
		if u = dao.QueryUserByPhone(phone); u != (model.User{}) {
			goto RETURN
		}
	}
	if email != "" {
		if u = dao.QueryUserByEmail(email); u != (model.User{}) {
			goto RETURN
		}
	}
	return "没有找到符合条件的用户", model.UserInfoRsp{}
RETURN:
	return conf.OKMsg, model.UserInfoRsp{
		Username: u.Username,
		Phone:    u.Phone,
		Email:    u.Email,
		Photo:    u.Photo,
	}

}

func UpdateUserInfo(myId string, u model.MyInfoUpdateReq) {
	user := dao.QueryUserById(myId)
	user.Username = u.Username
	user.Phone = u.Phone
	user.Email = u.Email
	dao.UpdateUser(user)
}

func UpdatePassword(myId string, pwd string, newPwd string) error {
	user := dao.QueryUserById(myId)
	if user.Password != utils.SHA256Secret(pwd) {
		return errors.New("密码不正确")
	}
	user.Password = utils.SHA256Secret(newPwd)
	dao.UpdateUser(user)
	return nil
}

func QueryMyPhoto(myId string) string {
	return dao.QueryUserById(myId).Photo
}

func QueryUserPhoto() {

}

func AddUserAddress(userId string, addr model.UserAddressReq) error {
	u := dao.QueryUserById(userId)
	var addrArr []model.UserAddressReq
	err := json.Unmarshal([]byte(u.Address), &addrArr)
	if err != nil {
		fmt.Println("AddUserAddress json.Unmarshal failed ...", err)
		return err
	}
	addrArr = append(addrArr, addr)
	newAddr, err := json.Marshal(addrArr)
	if err != nil {
		fmt.Println("AddUserAddress json.Marshal failed ...", err)
		return err
	}
	u.Address = string(newAddr)
	dao.UpdateUser(u)
	return nil
}

//func UpdateUser(username string, newUsername string, password string, phone string, email string) (msg string) {
//	user := dao.QueryUserByUsername(username)
//	if username != "" {
//		if !utils.IsValidUsername(newUsername) {
//			return "用户名需在1~20个字符之间"
//		}
//		if utils.IsRegisteredUsername(newUsername) {
//			return "用户名已存在"
//		}
//		user.Username = newUsername
//	}
//	if password != "" {
//		if !utils.IsValidPassword(password) {
//			return "密码由6~20位数字、字母和部分特殊字符组成"
//		}
//		user.Password = utils.SHA256Secret(password)
//	}
//	if phone != "" {
//		if !utils.IsValidPhone(phone) {
//			return "请输入正确的手机号"
//		}
//		if utils.IsRegisteredPhone(phone) {
//			return "手机号已被注册"
//		}
//		user.Phone = phone
//	}
//	if email != "" {
//		if !utils.IsValidEmail(email) {
//			return "请输入正确的邮箱"
//		}
//		if utils.IsRegisteredEmail(email) {
//			return "邮箱已被注册"
//		}
//		user.Email = email
//	}
//	dao.UpdateUser(user)
//	return conf.OKMsg
//}

func QueryAllUsers() (users []model.User) {
	return dao.QueryAllUsers()
}

func AddUserPhoto(userId string) {
	user := dao.QueryUserById(userId)
	user.Photo = conf.WebLinkPathOfUserPhoto + userId + ".jpg"
	dao.UpdateUser(user)
}

func QueryMyAddressList(myId string) ([]model.UserAddressRsp, error) {
	u := dao.QueryUserById(myId)
	var addrArr []model.UserAddressReq
	err := json.Unmarshal([]byte(u.Address), &addrArr)
	if err != nil {
		fmt.Println("QueryMyAddressList json.Unmarshal failed ...", err)
		return nil, err
	}
	var addrList []model.UserAddressRsp
	for i, addr := range addrArr {
		addrList = append(addrList, model.UserAddressRsp{
			Id:       i,
			Province: addr.Province,
			City:     addr.City,
			County:   addr.County,
			Detail:   addr.Detail,
			Tag:      addr.Tag,
		})
	}
	return addrList, nil
}

func DeleteAddress(myId string, addrId string) error {
	u := dao.QueryUserById(myId)
	var addrArr []model.UserAddressReq
	err := json.Unmarshal([]byte(u.Address), &addrArr)
	if err != nil {
		fmt.Println("DeleteAddress json.Unmarshal failed ...", err)
		return err
	}
	num, err := strconv.Atoi(addrId)
	if err != nil {
		fmt.Println("DeleteAddress strconv failed ...", err)
		return err
	}
	var after []model.UserAddressReq
	for i, addr := range addrArr {
		if i == num {
			continue
		}
		after = append(after, model.UserAddressReq{
			Province: addr.Province,
			City:     addr.City,
			County:   addr.County,
			Detail:   addr.Detail,
			Tag:      addr.Tag,
		})
	}
	newAddr, err := json.Marshal(after)
	if err != nil {
		fmt.Println("DeleteAddress json.Marshal failed ...", err)
		return err
	}
	u.Address = string(newAddr)
	dao.UpdateUser(u)
	return nil
}

func AddToShoppingCar(myId string, goodsId string) error {
	u := dao.QueryUserById(myId)
	var car []string
	err := json.Unmarshal([]byte(u.ShoppingCar), &car)
	if err != nil {
		fmt.Println("AddToShoppingCar json.Unmarshal failed ...", err)
		return err
	}
	car = append(car, goodsId)
	bytes, err := json.Marshal(car)
	if err != nil {
		fmt.Println("AddToShoppingCar json.Marshal failed ...", err)
		return err
	}
	u.ShoppingCar = string(bytes)
	dao.UpdateUser(u)
	return nil
}

func ShowShoppingCarList(myId string) ([]model.ShoppingCarRsp, error) {
	u := dao.QueryUserById(myId)
	var car []string
	err := json.Unmarshal([]byte(u.ShoppingCar), &car)
	if err != nil {
		fmt.Println("ShowShoppingCarList json.Unmarshal failed ...", err)
		return nil, err
	}
	var carList []model.ShoppingCarRsp
	for i := 0; i < len(car); i++ {
		goods := dao.QueryGoodsById(car[i])
		if goods == (model.Goods{}) {
			continue
		}
		shop := dao.QueryShopById(goods.ShopId)
		carList = append(carList, model.ShoppingCarRsp{
			Id:          i,
			GoodsId:     goods.Id,
			Name:        goods.Name,
			ShopName:    shop.Name,
			Kind:        goods.Kind,
			Price:       goods.Price,
			SoldAmount:  goods.SoldAmount,
			Score:       goods.Score,
			PictureLink: goods.PictureLink,
		})
	}
	return carList, nil
}

func DeleteFromShoppingCar(myId string, numId string) error {
	u := dao.QueryUserById(myId)
	var car []string
	err := json.Unmarshal([]byte(u.ShoppingCar), &car)
	if err != nil {
		fmt.Println("DeleteFromShoppingCar json.Unmarshal failed ...", err)
		return err
	}
	num, err := strconv.Atoi(numId)
	if err != nil {
		fmt.Println("DeleteFromShoppingCar strconv failed ...", err)
		return err
	}
	var newCar []string
	for i := 0; i < len(car); i++ {
		if i == num {
			continue
		}
		newCar = append(newCar, car[i])
	}
	bytes, err := json.Marshal(newCar)
	if err != nil {
		fmt.Println("DeleteFromShoppingCar json.Marshal failed ...", err)
		return err
	}
	u.ShoppingCar = string(bytes)
	dao.UpdateUser(u)
	return nil
}

func AddMoney(req model.AddMoneyReq, userId string) {
	u := dao.QueryUserById(userId)
	money, _ := strconv.ParseFloat(req.Money, 64)
	local, _ := strconv.ParseFloat(u.Money, 64)
	u.Money = fmt.Sprintf("%.2f", local+money)
	dao.UpdateUser(u)
}
