package middleware

import (
	"Gf-Vben/internal/dao"
	"Gf-Vben/internal/model"
	"Gf-Vben/internal/model/entity"
	"context"
	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"time"
)

var (
	// The underlying JWT middleware.
	GfJWTMiddleware *jwt.GfJWTMiddleware
)

func Auth(r *ghttp.Request) {
	GfJWTMiddleware.MiddlewareFunc()(r)
	r.Middleware.Next()
}

func initJwt() {
	auth := jwt.New(&jwt.GfJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte("secret key"),
		Timeout:         time.Hour * 24 * 7,
		MaxRefresh:      time.Hour * 24 * 7,
		IdentityKey:     "uid",
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		Authenticator:   Authenticator,
		Unauthorized:    Unauthorized,
		PayloadFunc:     PayloadFunc,
		IdentityHandler: IdentityHandler,
	})
	GfJWTMiddleware = auth
}

// PayloadFunc is a callback function that will be called during login.
// Using this function it is possible to add additional payload data to the webtoken.
// The data is then made available during requests via c.Get("JWT_PAYLOAD").
// Note that the payload is not encrypted.
// The attributes mentioned on jwt.io can't be used as keys for the map.
// Optional, by default no additional data will be set.
func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

// IdentityHandler get the identity from JWT and set the identity for every request
// Using this function, by r.GetParam("id") get identity
func IdentityHandler(ctx context.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	return claims[GfJWTMiddleware.IdentityKey]
}

// Unauthorized is used to define customized Unauthorized callback function.
func Unauthorized(ctx context.Context, code int, message string) {
	r := g.RequestFromCtx(ctx)
	r.Response.WriteJson(g.Map{
		"code":    code,
		"message": message,
	})
	r.ExitAll()
}

// Authenticator is used to validate login parameters.
// It must return user data as user identifier, it will be stored in Claim Array.
// if your identityKey is 'id', your user data must have 'id'
// Check error (e) to determine the appropriate error message.
func Authenticator(ctx context.Context) (interface{}, error) {
	r := g.RequestFromCtx(ctx)
	req := new(model.LoginReq)
	if err := r.Parse(req); err != nil {
		return "", err
	}
	var u entity.User
	err := dao.User.Ctx(r.GetCtx()).Where("username", req.Username).Scan(&u)
	if err != nil {
		return nil, err
	}
	if u.Status == 0 {
		return nil, gerror.New("用户已禁用")
	}
	pw, err := gmd5.Encrypt(req.Password)
	if err != nil {
		return nil, err

	}
	if pw != u.Password {
		return nil, gerror.New("用户密码错误")
	}

	return g.Map{
		"username": u.Username,
		"uid":      u.Id,
		"roles":    "admin",
	}, nil

}
