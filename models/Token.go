package models

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"fmt"
)

var key ="wosuobianxiedekeynibiequfanyi123haoma"
func ReturnToken()(back string,err error){
	key := []byte(key)

	claims := &jwt.StandardClaims{
		NotBefore: int64(time.Now().Unix()),
		ExpiresAt: int64((time.Now().Add(20*time.Second)).Unix()),
		Issuer:    "token",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)
	fmt.Println("签名后的token信息:", ss)
	if err !=nil{
		return "",err
	}else {
		return ss,nil
	}
}
//func VerifyToken(token string)(zz string,err error){
//	t, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
//		return []byte(key), nil
//	})
//
//	if err != nil {
//		fmt.Println("parase with claims failed.", err)
//		return "",err
//	}else {
//		fmt.Println(t)
//		fmt.Println(t.Claims)
//		fmt.Println(t.Header)
//		fmt.Println(t.Method)
//		fmt.Println(t.Valid)
//		fmt.Println(t.Raw)
//		return t.Claims.Valid().Error(),nil
//	}
//
//
//}
func CheckToken(token string) (err error) {
	t, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	fmt.Println(t)
	if err != nil {
		fmt.Println("parase with claims failed.", err)
		return err
	}else {
		fmt.Println(t.Valid)
		return nil
	}

}
