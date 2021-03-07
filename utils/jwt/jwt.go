package jwt

import (
	"GinScaffold/settings"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

//!!!go的全局变量初始化是在main函数前的 所以此处会导致空指针异常
//go语言初始化顺序 const var init main
//var jwtSecret = settings.Conf.JwtConfig.Secret

//jwt-go中自带的jwt.StandardClaims只包含了官方字段
type MyClaims struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenToken(userID int64, username string) (string, error) {
	c := MyClaims{
		UserId:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(viper.GetInt("jwt.expire_duration")) * time.Hour).Unix(),
			Issuer:    viper.GetString("name"),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	//需要穿byte数组！！！
	return token.SignedString([]byte(settings.Conf.JwtConfig.Secret))
}

func ParseToken(token string) (*MyClaims, error) {
	mc := new(MyClaims)
	t, err := jwt.ParseWithClaims(token, mc, func(token *jwt.Token) (interface{}, error) {
		return settings.Conf.JwtConfig.Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if t.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
