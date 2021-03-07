package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/spf13/viper"
)

func EncryptPassword(origin string) string {
	secret := viper.GetString("md5_secret")
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(origin)))
}
