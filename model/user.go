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
	UserName           string `json:"username,omitempty" gorm:"unique;column:username"`
	Password           string `json:"password,omitempty" gorm:"column:password"`
	ForgotPasswordCode string `json:"forgot_password_code,omitempty" gorm:"column:forgot_password_code"`
	VerificationCode   string `json:"verification_code,omitempty" gorm:"column:verification_code"`

	// Optional
	Email     *string `json:"email,omitempty" gorm:"column:email"`
	Locked    bool    `json:"locked,omitempty" gorm:"column:locked"`
	Veryfied  bool    `json:"veryfied,omitempty" gorm:"column:veryfied"`
	AvatarURL *string `json:"avatar_url,omitempty" gorm:"column:avatar_url"`
	NickName  *string `json:"nickname,omitempty" gorm:"column:nickname"`
	Phone     *string `json:"phone,omitempty" gorm:"column:phone"`
	Blog      *string `json:"blog,omitempty" gorm:"column:blog"`
	Bio       *string `json:"bio,omitempty" gorm:"column:bio"`
}

var auth = new(Auth)

func (u *User) Login(form mappers.LoginForm, db *gorm.DB, conf *gconfig.Config) (token *Token, err error) {

	db.Model(&User{}).Where("username = ?", form.UserName).First(&u)

	bytePassword := []byte(form.Password)
	byteHashedPassword := []byte(u.Password)

	err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)

	if err != nil {
		return token, errors.New("invalid password")
	}

	td, err := auth.CreateToken(u.UserName, conf)
	if err != nil {
		return token, err
	}
	token = &Token{
		AccessToken:  td.AccessToken,
		RefreshToken: td.RefreshToken,
	}
	return token, err
}

func (u *User) Register(form mappers.RegisterForm, db *gorm.DB, conf *gconfig.Config) (err error) {
	err = db.Model(&User{}).Where("username = ?", form.UserName).First(&u).Error
	if err == nil {
		return errors.New("username already exists")
	}

	bytePassword := []byte(form.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	u.UserName = form.UserName
	u.Email = form.Email
	u.Password = string(hashedPassword)
	u.VerificationCode = uuid.NewV4().String()
	u.ForgotPasswordCode = uuid.NewV4().String()
	err = db.Create(&u).Error
	return err
}

func (u *User) GetByUsername(username string, db *gorm.DB) (err error) {
	err = db.Model(&User{}).Where("username = ?", username).First(&u).Error
	return err
}

func (u *User) GetByEmail(email string, db *gorm.DB) (err error) {
	err = db.Model(&User{}).Where("email = ?", email).First(&u).Error
	return err
}

func (u *User) GetByVerificationCode(verificationCode string, db *gorm.DB) (err error) {
	err = db.Model(&User{}).Where("verification_code = ?", verificationCode).First(&u).Error
	return err
}

func (u *User) GetByForgotPasswordCode(forgotPasswordCode string, db *gorm.DB) (err error) {
	err = db.Model(&User{}).Where("forgot_password_code = ?", forgotPasswordCode).First(&u).Error
	return err
}

func (u *User) UpdateVerificationCode(username string, db *gorm.DB) (err error) {
	err = db.Model(&User{}).Where("username = ?", username).First(&u).Error
	if err != nil {
		return err
	}

	u.VerificationCode = uuid.NewV4().String()
	err = db.Save(&u).Error
	return err
}

func (u *User) UpdateForgotPasswordCode(username string, db *gorm.DB) (err error) {
	err = db.Model(&User{}).Where("username = ?", username).First(&u).Error
	if err != nil {
		return err
	}

	u.ForgotPasswordCode = uuid.NewV4().String()
	err = db.Save(&u).Error
	return err
}

func (u *User) UpdatePassword(username string, password string, db *gorm.DB) (err error) {
	err = db.Model(&User{}).Where("username = ?", username).First(&u).Error
	if err != nil {
		return err
	}

	bytePassword := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	u.Password = string(hashedPassword)
	err = db.Save(&u).Error
	return err
}

func (u *User) UpdateEmail(username string, email *string, db *gorm.DB) (err error) {
	err = db.Model(&User{}).Where("username = ?", username).First(&u).Error
	if err != nil {
		return err
	}

	u.Email = email
	err = db.Save(&u).Error
	return err
}

func (u *User) UpdateAvatarURL(username string, avatarURL *string, db *gorm.DB) (err error) {
	err = db.Model(&User{}).Where("username = ?", username).First(&u).Error
	if err != nil {
		return err
	}

	u.AvatarURL = avatarURL
	err = db.Save(&u).Error
	return err
}
