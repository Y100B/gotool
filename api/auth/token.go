package auth

import (
	"gotool/api/models"
	"gotool/config"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)


/*
 * 非登录用户拦截器
 */
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			token = c.Request.Header.Get("Authorization")
		}
		claims, err := ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				if token, err = RefreshToken(token); err == nil {
					c.Header("Authorization", token)
					c.JSON(http.StatusOK, gin.H{"code": 403, "message": "refresh token", "token": token})
					c.AbortWithStatus(http.StatusBadRequest)
					return
				}
			}
			c.JSON(http.StatusUnauthorized, gin.H{"code": 403, "message": err.Error()})
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.Set("claims", claims)
	}
}

var (
	TokenExpired = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed = errors.New("That's not even a token")
	TokenInvalid = errors.New("Couldn't handle this token:")
	SignKey = "test"
)

/*
 * 新建JWT数据
 */
func GenerateJWT(user models.User) (string, error) {
	claim := models.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "Alfredo Mendoza",
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(config.SECRETKEY)
}

/*
 * 解析token数据
 */
func ParseToken(tokenString string) (*models.Claim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.Claim{}, func(token *jwt.Token) (interface{}, error) {
		return config.SECRETKEY, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*models.Claim); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

/*
 * 刷新token数据
 */
func RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &models.Claim{}, func(token *jwt.Token) (interface{}, error) {
		return config.SECRETKEY, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*models.Claim); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return GenerateJWT(claims.User)
	}
	return "", TokenInvalid
}