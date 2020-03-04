package auth

import (
	"context"
	"doudou/pkg/auth/auther"
	jwtauth "doudou/pkg/auth/auther/jwt"
	buntdb "doudou/pkg/auth/auther/jwt/store"
	"doudou/pkg/errors"
	"doudou/pkg/utils"
	"log"

	"github.com/dgrijalva/jwt-go"
)

var (
	JWT_SIGNING_METHOD = utils.GetEnvironment("JWT_SIGNING_METHOD", "HS512")
	JWT_SIGNING_KEY    = utils.GetEnvironment("JWT_SIGNING_KEY", "doudou")
	JWT_EXPIRED        = utils.GetEnvironment("JWT_EXPIRED", "7200")
	JWT_STORE          = utils.GetEnvironment("JWT_STORE", "file")
	JWT_STORE_PATH     = utils.GetEnvironment("JWT_STORE_PATH", "./jwt_auth.db")
)

type AuthServer struct {
	auther auther.Auther
}

func DefaultAuthServer() *AuthServer {
	auther, _ := NewAuther()
	return &AuthServer{
		auther: auther,
	}
}
func NewAuthServer(a auther.Auther) (*AuthServer, error) {
	return &AuthServer{
		auther: a,
	}, nil
}

func (a *AuthServer) GenerateToken(ctx context.Context, userID string) (string, error) {
	log.Println("in GenerateToken")
	log.Println(userID)
	log.Println(a.auther)
	tokenInfo, err := a.auther.GenerateToken(userID)
	if err != nil {
		return "", errors.WithStack(err)
	}
	log.Println(tokenInfo)
	return tokenInfo.GetAccessToken(), nil
}

func (a *AuthServer) DestroyToken(ctx context.Context, token string) error {
	log.Println("in DestroyToken")
	log.Println(token)
	err := a.auther.DestroyToken(token)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (a *AuthServer) VertifyToken(ctx context.Context, token string) (string, error) {
	log.Println("in VertifyToken")
	log.Println(token)

	uid, err := a.auther.ParseUserID(token)

	if err != nil {
		if err == auther.ErrInvalidToken {
			return "", errors.ErrInvalidToken
		}
		return "", err
	}
	return uid, nil
}

func NewAuther() (auther.Auther, error) {
	exp, _ := utils.S(JWT_EXPIRED).Int()
	var opts []jwtauth.Option
	opts = append(opts, jwtauth.SetExpired(exp))
	opts = append(opts, jwtauth.SetSigningKey([]byte(JWT_SIGNING_KEY)))
	opts = append(opts, jwtauth.SetKeyfunc(func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, auther.ErrInvalidToken
		}
		return []byte(JWT_SIGNING_KEY), nil
	}))

	switch JWT_SIGNING_METHOD {
	case "HS256":
		opts = append(opts, jwtauth.SetSigningMethod(jwt.SigningMethodHS256))
	case "HS384":
		opts = append(opts, jwtauth.SetSigningMethod(jwt.SigningMethodHS384))
	case "HS512":
		opts = append(opts, jwtauth.SetSigningMethod(jwt.SigningMethodHS512))
	}

	var store jwtauth.Storer
	switch JWT_STORE {
	case "file":
		s, err := buntdb.NewStore(JWT_STORE_PATH)
		if err != nil {
			println(err)
			return nil, err
		}
		store = s
	}
	println(">>>> NewAuther <<<<")
	return jwtauth.New(store, opts...), nil
}
