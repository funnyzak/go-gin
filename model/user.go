package model

import (
	"errors"
	"go-gin/internal/gconfig"
	"go-gin/mappers"

	"github.com/twinj/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Common
	UserName           string `json:"username,omitempty"`
	Password           string `json:"password,omitempty"`
	ForgotPasswordCode string `json:"forgot_password_code,omitempty"`
	VerificationCode   string `json:"verification_code,omitempty"`

	// Optional
	Email     string `json:"email,omitempty"`
	Locked    bool   `json:"locked,omitempty"`
	Veryfied  bool   `json:"veryfied,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
	NickName  string `json:"nickname,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Blog      string `json:"blog,omitempty"`
	Bio       string `json:"bio,omitempty"`
}

var auth = new(Auth)

func (user User) Login(form mappers.LoginForm, db *gorm.DB, conf *gconfig.Config) (token Token, err error) {

	db.Model(&User{}).Where("username = ?", form.UserName).First(&user)

	bytePassword := []byte(form.Password)
	byteHashedPassword := []byte(user.Password)

	err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)

	if err != nil {
		return token, errors.New("invalid password")
	}

	tokenDetails, err := auth.CreateToken(user.UserName, conf)

	if err == nil {
		token.AccessToken = tokenDetails.AccessToken
		token.RefreshToken = tokenDetails.RefreshToken
	}

	return token, nil
}

func (u User) Register(form mappers.RegisterForm, db *gorm.DB, conf *gconfig.Config) (user User, err error) {
	err = db.Model(&User{}).Where("username = ?", form.UserName).First(&u).Error
	if err != nil {
		return user, err
	}

	bytePassword := []byte(form.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user.UserName = form.UserName
	user.Email = form.Email
	user.Password = string(hashedPassword)
	user.VerificationCode = uuid.NewV4().String()
	user.ForgotPasswordCode = uuid.NewV4().String()
	err = db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, err
}
