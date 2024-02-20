package form

import ()

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}

type AccessDetails struct {
	AccessUUID string
	UserID     int64
}

type AuthForm struct {}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenForm struct {
	// AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refreshToken"`
}

func (f RefreshTokenForm) Validate() string {
	if f.RefreshToken == "" {
		return "refreshToken is required"
	}
	return ""
}
