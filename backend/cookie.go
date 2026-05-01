package backend

import (
	"fmt"
	"net/http"
	"time"

	"github.com/lestrrat-go/jwx/v4/jwa"
	"github.com/lestrrat-go/jwx/v4/jwk"
	"github.com/lestrrat-go/jwx/v4/jwt"
)

type Server struct {
	alg       jwa.SignatureAlgorithm
	signKey   jwk.Key
	verifyKey jwk.Key
}

var server_info Server

func Setup() {
	//HS256 is more efficient than RS256.
	//So we will use the HS256 algorithm even though it is a little less secure to make less calculations.
	server_info.alg = jwa.HS256()
	symKey, err := jwk.Import[jwk.SymmetricKey]([]byte("To put a secure string with more than 32 characters!"))
	if err != nil {
		// handle error
	}
	server_info.signKey = symKey
	server_info.verifyKey = server_info.signKey
}

func GenerateJWT(id int) []byte {

	token, err := jwt.NewBuilder().
		Issuer(`NOTESHAREEFREI`).
		Audience([]string{`users`}).
		Expiration(time.Now().Add(24*time.Hour)).
		Claim("acc", id). //acc for account and the id is the one stored inside the database.
		IssuedAt(time.Now()).
		Build()
	if err != nil {
		// handle errors
		fmt.Print(err.Error())
		panic("could not build the token")
	}

	signed, err := jwt.Sign(token, jwt.WithKey(server_info.alg, server_info.signKey))
	if err != nil {
		// handle errors
		fmt.Print(err.Error())
		panic("Failed generating the token with the key.")
	}

	return signed
}

func ValidateJWT(token jwt.Token) (int, error) {
	err := jwt.Validate(token, jwt.WithIssuer(`NOTESHAREEFREI`), jwt.WithRequiredClaim("acc"))
	if err != nil {
		//Error is handled outside
		return 0, err
	}

	account_id, err := jwt.Get[float64](token, "acc")
	if err != nil {
		//Error is handled outside
		return 0, err
	}
	fmt.Println("account id in float 64 : ", account_id, int(account_id))

	return int(account_id), nil
}

func GenerateCookieWithJWT(JWT []byte) http.Cookie {
	jwtcookie := http.Cookie{
		Name:  "__Host-Http-Jwt",
		Value: string(JWT),

		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
		MaxAge:   86400, //Equals 1 day in seconds.
	}
	return jwtcookie
}
