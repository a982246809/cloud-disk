package define

import (
	"github.com/dgrijalva/jwt-go"
)

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

// 用于加密的ksy
var JwtKey = "cloud-disc-key"

// TencentSecretKey 腾讯云对象存储
var TencentSecretKey = "83mIclhRyJQmgMo22nKdkeJqJYYJsCI9"
var TencentSecretID = "AKIDhFdpngMuKXKwLi095pMipiPuDPPJ4d5K"
var CosBucket = "https://getcharzp-1256268070.cos.ap-chengdu.myqcloud.com"
