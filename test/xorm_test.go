package test

import (
	"bytes"
	"cloud-disk/core/models"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"xorm.io/xorm"
)

func TestXormTest(t *testing.T) {
	println("123213")
	//192.168.0.111
	engine, err := xorm.NewEngine("mysql", "root:123456@tcp(localhost:3306)/cloud-disk?charset=utf8mb4&parseTime=True")
	if err != nil {
		fmt.Println(err)
		t.Fatal()
	}

	data := make([]*models.UserBasic, 0)
	fmt.Println(data)
	err = engine.Find(&data)
	if err != nil {
		fmt.Println(err)
		t.Fatal()
	}
	fmt.Println(data)

	b, _ := json.Marshal(data)
	fmt.Println(b)
	dst := new(bytes.Buffer)
	fmt.Println(dst)
	json.Indent(dst, b, "", "")
	fmt.Println(dst.String())
}
