package gjwt

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jeffotoni/go.standard.project.layout/project.web/standard.libray/crud.user.singleton/pkg/cert"
	"github.com/jeffotoni/go.standard.project.layout/project.web/standard.libray/crud.user.singleton/pkg/models/jwt"
	"github.com/jeffotoni/go.standard.project.layout/project.web/standard.libray/crud.user.singleton/pkg/pkg/zerolog"
)

// jwt init
func init() {
	var errx error
	privateByte := []byte(cert.RSA_PRIVATE)
	PrivateKey, errx = jwt.ParseRSAPrivateKeyFromPEM(privateByte)
	if errx != nil {
		zerolog.Error(
			"1.0.0",
			"gjwt.go",
			25,
			"api.crud.user.com.br",
			"init jwt.ParseRSAPrivateKeyFromPEM(privateByte)",
			errx.Error())
		return
	}

	publicByte := []byte(cert.RSA_PUBLIC)
	PublicKey, errx = jwt.ParseRSAPublicKeyFromPEM(publicByte)
	if errx != nil {
		zerolog.Error(
			"1.0.0",
			"gjwt.go",
			38,
			"api.crud.user.com.br",
			"init jwt.ParseRSAPublicKeyFromPEM(publicByte)",
			errx.Error())
		return
	}
}