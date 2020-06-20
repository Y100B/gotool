package auth

import (
	"gotool/api/database"
	"gotool/api/models"
	"gotool/api/security"
	"gotool/api/utils/channels"
	"github.com/jinzhu/gorm"
)

// SignIn method
func Login(email, PassWord string) (string, error) {
	user := models.User{}
	var err error
	var db *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		db= database.NewDb()
		if err != nil {
			ch <- false
			return
		}
		err = db.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
		if err != nil {
			ch <- false
			return
		}
		err = security.VerifyPassWord(user.PassWord, PassWord)
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		user.PassWord = ""
		return GenerateJWT(user)
	}

	return "", err
}

