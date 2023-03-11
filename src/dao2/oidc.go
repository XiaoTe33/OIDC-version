package dao2

import "winter-examination/src/model"

func QueryByClientSecretAndClientId(clientId string, clientSecret string) error {
	c := model.Client{}
	return GetDB().Where("client_id = ?", clientId).Where("client_secret = ?", clientSecret).Find(&c).Error
}

func AddClientSecret(userId, clientId, clientSecret string) error {

	return GetDB().Create(&model.Client{
		UserId:       userId,
		ClientId:     clientId,
		ClientSecret: clientSecret,
	}).Error
}

func QueryClientByClientId(clientId string) (model.Client, error) {
	c := model.Client{}
	e := GetDB().Where("client_id = ?", clientId).Find(&c).Error
	return c, e
}
