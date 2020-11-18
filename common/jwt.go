package common

import (
	"crypto/rsa"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
	"management/log"
)

type Jwt struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func NewJwt() (j *Jwt) {
	//var cfg struct {
	//	JwtSecret string `json:"jwt_secret" toml:"jwt_secret"`
	//}

	//_, err := toml.DecodeFile("./config/application.toml", &cfg)
	//if err != nil {
	//	panic(err)
	//}
	sk, err := jwt.ParseRSAPrivateKeyFromPEM([]byte("-----BEGIN RSA PRIVATE KEY-----\nMIIEpgIBAAKCAQEA0CXaTQE4ifUm2+tlH1oBLif6yWT7vRnnCqEvX8JUcmLQMrVu\n9WVMZbKDHphGMt6dFtQpCaXcrjPyPLqouel3lbro8vgIuPNSoeE1OZVbVWElvZ2V\npEvSA1rPe0bZ5F6A99Ts0Rz6uulVQkIkMQoxc4KFgrmBk+3a3VALsZLcEHSCsnnZ\nhnr6jYYSdh2DQvhIFqNaptFkEtb/xLLP9f0Xv2iVIq7wpMWByficDjzmuWxVeor+\nAUIy9fGcBepxQCgKq1yXhY/0zDUbgUUB/lxIGoQCDEKLRwJyQw/BbNeY5NFlZGqA\n3uPOoKhyMv5Pi/g5E4opSBx5W/lv3gKpB1cfpQIDAQABAoIBAQDGm2ep1EFjeXSj\noP8zJAk+Rk2IPv/pFr8aqGPwphc3sctgpzgBlK+J1gRAfCF3RmxzrOqfVxCzc8Nu\naNi30+oUB21g8IQ6HYo6Bg5oLHgihnihbaysQOBZ7RtOUHN18Spzz0pL2a/wCtYc\nS8oGtOgshFzqOCFIykrsowUVYcDzPM7E7KTaAoGCOhoL02Nw2iqtfHWkK/CuiSEG\nDoV4yLh6cSlSemKmN+kkcXICtbvAfLQFuKr1jUAC/3xW0MbnZtfkbbhb/0OBI1ri\nnYKgZbe5D9ne8nuZrcMH7D5FetoS5Ax7SIIZWYxmXGx74mxbeQRyfTLwLXdOGRGC\nup8VdHshAoGBAPpTH9q28SAgYKcQmCX+B3Aee/x1zmhfUPhYnGTcFQxMcfkUdTfx\noOJswWE+XRERkLxpRnIhpzgtJp2+pWhJsE6fFNMxML5fK5yeEF6J1EbCyGEwOtst\nvvigLs5oR9MlbF7nzopRMViN43e5BjkUD22yI5EMUZW4ubqa8mIlsN3NAoGBANTd\n73hywAubUnmlw4c2uoUzmINOxX+aUID87ONn0+fzY3Ug9LtJ6J0R88P4qY1wgdrU\nQjCMFgzmCH9q1Am8W1FqJkjGS/tWgVOIrUyfBRJO+EIlMIellR6n9DXCw9LGeO2V\nI7m6U1NWZGa8mtl/0S54QhNsfsp8p5twWV1tB7E5AoGBAKDzNXYRTnRTnRGOD+XN\nscabMykeLfrZ3lvvzY7kGvxvYpC+YKf5ynILb0MxL/G7k44xOkRD8xqhnUSrwfqN\n9rh2fJNV+3tMAeSPlQLUKBLfRquGsTEf9rwxcibw0c2nMEjNTvWMQugnQuxFoQSu\nK0Vi1o96ljJoNbMP0Wzdwxy5AoGBAK1kuxRaJKVPuDbvF/6kTfsCtFEBcU8n3Du1\nyyDSCoL+dx2J4tBMu/Z2ESKpAzP7WUtvaxswgSWwm2tvEZl8nMYMuXK+VFY/eMka\npE+tmOv497CpqoZUEswN85d3NxwSH58nxRoc9JMF5HLrXxecTkCUJP69eepm8ABl\n2+WGUqXBAoGBANoU+NjYsC5XVdk37iGjjLlFdHVX6G9AB1n2DSFGxS1qzJrwYk0C\n0kY+pDLzz9QLw+T7Sp6aochrgxZY97suGRnzw/jf4lVG48lg8d2dqXiqJGMmKLXh\n9vN4F6/II8r+aEDkHJU07VCTXvOnE6zJaSbsJiGWaKOqFTw3GnTQYiok\n-----END RSA PRIVATE KEY-----"))
	if err != nil {
		err = errors.New("jwt parse private key error: " + err.Error())
		return
	}
	j = &Jwt{PrivateKey: sk, PublicKey: sk.Public().(*rsa.PublicKey)}
	return
}

func (t *Jwt) CreateToken(username string) (accessToken string, err error) {
	// jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{"username": username})
	// sign the jwt token
	accessToken, err = token.SignedString(t.PrivateKey)
	return
}

func (t *Jwt) DecodeToken(token string) (uid string, err error) {
	// get map claims
	claims, err := t.claimsFromToken(token)
	if err != nil {
		return
	}
	if _, ok := claims["username"]; !ok {
		err = errors.New("token is not expected")
		return
	}
	return claims["username"].(string), nil
}

func (t *Jwt) claimsFromToken(tokenString string) (jwt.MapClaims, error) {
	// parse token
	jwtToken, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			log.Logger.Error("", zap.String("token", "token解析错误"))
			return
		}
		return t.PublicKey, nil
	})

	// get claims
	var claims jwt.MapClaims
	if jwtToken == nil || jwtToken.Claims == nil {
		return claims, errors.New("jwtToken error")
	}

	claims = jwtToken.Claims.(jwt.MapClaims)
	return claims, err
}
