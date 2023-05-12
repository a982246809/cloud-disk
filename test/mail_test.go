package test

import (
	"bytes"
	"cloud-disk/core/define"
	"cloud-disk/core/helper"
	"context"
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/smtp"
	"net/url"
	"os"
	"testing"
)

func TestMailTest(t *testing.T) {
	e := email.NewEmail()
	e.From = "Jordan Wright <982246809@qq.com>"
	e.To = []string{"982246809@qq.com"}

	e.Subject = "Awesome Subject"

	e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
	e.SendWithTLS("smtp.qq.com:465", smtp.PlainAuth("", "982246809@qq.com", "hnoasrvcbqfgbgaj", "smtp.qq.com"), &tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
}

func TestRandCode(t *testing.T) {
	//t, err := time.Parse("Mon Jan 02 2006 15:04:05 GMT-0700 (MST)", "Mon May 01 2023 00:00:00 GMT+0800 (中国标准时间)")
	println(helper.RandCode())
}

func TestUUID(t *testing.T) {
	v4 := uuid.NewV4()
	fmt.Println(v4)

}
func TestCos(t *testing.T) {
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
	name := "cloud-disk/mail.test.go"
	// 1.通过字符串上传对象
	file, err := os.ReadFile("./mail_test.go")
	if err != nil {
		panic(err)
	}
	_, err = c.Object.Put(context.Background(), name, bytes.NewReader(file), nil)
	if err != nil {
		panic(err)
	}
	// 2.通过本地文件上传对象
	//_, err := c.Object.PutFromFile(context.Background(), name, "./mail_test.go", nil)
	//if err != nil {
	//	panic(err)
	//}

}

func Test22(t *testing.T) {
	//int
	println(string(int64(65)))

}
