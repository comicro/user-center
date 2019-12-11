package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"user-center/conf"
	pb "user-center/proto/user"
)

type CustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}

type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

type TokenService struct {
}

/**
 * 解密token的信息
 */
func (srv *TokenService) Decode(tokenString string) (*CustomClaims, error) {

	fmt.Println("decode token")

	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return conf.TokenKey(), nil
	})

	if err != nil {
		return nil, err
	}

	// 验证 token 并且返回自定义的 claims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

/**
 * 通过用户生成token
 */
func (srv *TokenService) Encode(user *pb.User) (string, error) {

	fmt.Println("encode json")

	// 过期时间检测
	expireTime := time.Now().Add(time.Hour * 72).Unix()

	// 生成claims
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireTime,
			Issuer:    "go.micro.srv.user",
		},
	}
	// 生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(conf.TokenKey())

	return tokenString, err
}
