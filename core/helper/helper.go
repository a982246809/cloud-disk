package helper

import (
	"cloud-disk/core/define"
	"context"
	"crypto/md5"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
	"math/rand"
	"net/http"
	"net/smtp"
	"net/url"
	"path"
	"strconv"
)

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func GenerateToken(id int, identity, name string) (string, error) {
	//id
	//identity
	//name
	claim := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	//加密
	signedString, err := token.SignedString([]byte(define.JwtKey))

	return signedString, err
}
func AnalyzeToken(token string) (*define.UserClaim, error) {
	uc := new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, errors.New("token is not valid")
	}
	return uc, nil
}

func MailSendCode(toEmail, code string) error {
	e := email.NewEmail()
	e.From = "Jordan Wright <982246809@qq.com>"
	e.To = []string{toEmail}

	e.Subject = "Awesome Subject"

	e.HTML = []byte("<h1>验证码:" + code + "</h1>")
	err := e.SendWithTLS("smtp.qq.com:465", smtp.PlainAuth("", "982246809@qq.com", "hnoasrvcbqfgbgaj", "smtp.qq.com"), &tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
	if err != nil {
		return err
	}
	return nil
}

func RandCode() string {
	code := ""
	for i := 0; i < 6; i++ {
		code += strconv.Itoa(rand.Intn(10))
	}
	return code
}

func GetUUID() string {
	return uuid.NewV4().String()
}

func CosUpload(r *http.Request) (string, error) {
	// 将 examplebucket-1250000000 和 COS_REGION 修改为真实的信息
	// 存储桶名称，由 bucketname-appid 组成，appid 必须填入，可以在 COS 控制台查看存储桶名称。https://console.cloud.tencent.com/cos5/bucket
	// COS_REGION 可以在控制台查看，https://console.cloud.tencent.com/cos5/bucket, 关于地域的详情见 https://cloud.tencent.com/document/product/436/6224
	u, _ := url.Parse("https://1-1318162858.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,  // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			SecretKey: define.TencentSecretKey, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
		},
	})

	//_, err := c.Bucket.Put(context.Background(), nil)
	//if err != nil {
	//	panic(err)
	//}

	// 对象键（Key）是对象在存储桶中的唯一标识。
	// 例如，在对象的访问域名 `examplebucket-1250000000.cos.COS_REGION.myqcloud.com/test/objectPut.go` 中，对象键为 test/objectPut.go
	formFile, header, err := r.FormFile("file")

	name := "cloud-disk/" + GetUUID() + path.Ext(header.Filename)
	// 1.通过字符串上传对象
	//file, err := os.ReadFile("./mail_test.go")
	if err != nil {
		panic(err)
	}
	_, err = c.Object.Put(context.Background(), name, formFile, nil)
	if err != nil {
		panic(err)
	}
	// 2.通过本地文件上传对象
	//_, err := c.Object.PutFromFile(context.Background(), name, "./mail_test.go", nil)
	//if err != nil {
	//	panic(err)
	//}

	return "https://1-1318162858.cos.ap-nanjing.myqcloud.com" + "/" + name, nil

}
