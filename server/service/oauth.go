package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"

	"github.com/hoon3051/TilltheCop/server/form"
	"github.com/hoon3051/TilltheCop/server/initializer"
	"github.com/hoon3051/TilltheCop/server/model"
	"gorm.io/gorm"

	"golang.org/x/oauth2"
)

type OauthService struct{}

func (svc OauthService) GenerateStateString() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func (svc OauthService) GetGoogleOauthURL(state string) string {
	//Create oauthConfig
	GoogleOauthURL := form.GetGoogleOauthConfig()
	url := GoogleOauthURL.AuthCodeURL(state)
	return url
}

func (svc OauthService) GetOauthToken(code string) (*oauth2.Token, error) {
	GoogleOauthConfig := form.GetGoogleOauthConfig()
	token, err := GoogleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (svc OauthService) GetOauthUser(token *oauth2.Token) (form.OauthUser, error) {
	GoogleOauthConfig := form.GetGoogleOauthConfig()
	client := GoogleOauthConfig.Client(context.Background(), token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return form.OauthUser{}, err
	}
	defer response.Body.Close() // 응답 본문을 함수 종료 시 자동으로 닫도록 함

	userInfo := form.OauthUser{}
	if err := json.NewDecoder(response.Body).Decode(&userInfo); err != nil {
		return form.OauthUser{}, err
	}

	return userInfo, nil
}

func (svc OauthService) FindUserExists(userInfo form.OauthUser) bool {
	var userModel = model.User{}
	initializer.DB.First(&userModel, "email=?", userInfo.Email)

	if userModel.ID != 0 {
		return true
	} else {
		return false
	}
}

func (svc OauthService) SaveOauthUser(oauthToken *oauth2.Token, userInfo form.OauthUser) (form.OauthUser, form.Token, error) {
	var userModel = model.User{}
	var oauthModel = model.Oauth{}

	// get user from DB
	initializer.DB.First(&userModel, "email=?", userInfo.Email)

	// update oauth from DB
	initializer.DB.First(&oauthModel, "user_id=?", userModel.ID)
	oauthModel.AccessToken = oauthToken.AccessToken
	oauthModel.RefreshToken = oauthToken.RefreshToken
	oauthModel.Expiry = oauthToken.Expiry
	initializer.DB.Save(&oauthModel)

	// return user
	user := form.OauthUser{
		ID:    userInfo.ID,
		Email: userModel.Email,
	}

	// Create token
	var authService = AuthService{}
	td, err := authService.CreateToken(int64(userModel.ID))
	if err != nil {
		return user, form.Token{}, err
	}

	// Save token
	err = authService.SaveToken(int64(userModel.ID), td)
	if err != nil {
		return user, form.Token{}, err
	}

	// Set return value (token)
	token := form.Token{
		AccessToken:  td.AccessToken,
		RefreshToken: td.RefreshToken,
	}

	return user, token, nil
}

func (svc OauthService) Register(tx *gorm.DB, oauthToken form.OauthToken, userId uint) error {
	var oauthModel = model.Oauth{}
	oauthModel.Provider = "google"
	oauthModel.AccessToken = oauthToken.AccessToken
	oauthModel.RefreshToken = oauthToken.RefreshToken
	oauthModel.Expiry = oauthToken.Expiry
	oauthModel.User_id = userId
	result := tx.Create(&oauthModel)
	if result.Error != nil {
		return result.Error
	}

	return nil

}
