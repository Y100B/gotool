package controllers

import (
	"github.com/gin-gonic/gin"
	"gotool/api/auth"
	"gotool/api/database"
	"gotool/api/models"
	"gotool/api/repository"
	"gotool/api/repository/crud"
	"log"
	"net/http"

	_ "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/gin-swagger/swaggerFiles"
)

// @Tags 网站用户系统
// @Description 用户注册
// @Accept json
// @Produce json
// @Param   Email         path    string     true        "Email"
// @Param   PassWord      path    string     true        "PassWord"
// @Param   NickName      path    string     true        "NickName"
// @Param   AvatarUrl     path    string     true        "AvatarUrl"
// @Success 200 {string} string    "ok"
// @Router /v1/api/user/register [post]
func Register(c *gin.Context) {
	user := models.User{}
	if err:=c.ShouldBindJSON(&user);err==nil{
		if len(user.Email) ==0 || len(user.PassWord)==0{
			c.JSON(200, gin.H{"code": 201, "msg": "邮箱及密码不能为空！", "data": ""})
		}
		db := database.NewDb()
		repo := crud.NewRepositoryUsersCRUD(db)
		func(userRepository repository.UserRepository) {
			userDb, err := userRepository.Save(user)
			if err != nil {
				c.JSON(200, gin.H{"code": 201, "msg": "保存用户出错！", "data": ""})
				return
			}
			userDb.PassWord = ""
			c.JSON(200, gin.H{"code": 200, "msg": "注册成功！", "data": userDb})
		}(repo)
	}else {
		log.Println(err)
		c.JSON(200, gin.H{"code": 201, "msg": "表单解析错误！", "data": ""})
	}

}

// @Tags 网站用户系统
// @Description 用户登录
// @Accept json
// @Produce json
// @Param   Email        path    string     true        "Email"
// @Param   PassWord     path    string     true        "PassWord"
// @Success 200 {string} string    "ok"
// @Router /v1/api/user/login [post]
func Login(c *gin.Context) {
	user := models.User{}
	if err:=c.ShouldBindJSON(&user);err==nil{
		token, err := auth.Login(user.Email, user.PassWord)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "登录失败！", "data": ""})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "登录成功！", "Authorization": token})
	}else {
		log.Println(err)
		c.JSON(200, gin.H{"code": 201, "msg": "表单解析错误！", "data": ""})
	}

}

// @Tags 网站用户系统
// @Description 判断是否是登录用户
// @Accept x-www-form-urlencoded
// @Produce json
// @Success 200 {string} string    "ok"
// @Router /v1/api/user/data [post]
func Data(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		token = c.Request.Header.Get("Authorization")
	}
	if token == "" {
		token = c.PostForm("token")
	}
	res, err := auth.ParseToken(token)
	if err != nil {
		if err == auth.TokenExpired {
			newToken, err := auth.RefreshToken(token)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"code": 301, "msg": "刷新token失败！", "data": ""})
			} else {
				c.Header("Authorization", newToken)
				c.JSON(http.StatusOK, gin.H{"code": 300, "msg": "刷新token成功!", "data": res.User,"token":newToken})
				}
			} else {
				c.JSON(http.StatusOK, gin.H{"code": 403, "msg": "token错误！", "data": ""})
			}
		} else {
			c.Header("Authorization", token)
			c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "获取用户信息成功！", "data": res.User})
		}
	}
