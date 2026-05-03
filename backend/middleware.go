package backend

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lestrrat-go/jwx/v4/jwt"
)

func simpleAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Checking the token
		token, err := jwt.ParseRequest(r, jwt.WithCookieKey("Http-Jwt"))
		if err != nil {
			//Invalid token or missing.
			w.Header().Set("WWW-Authenticate", `Bearer realm="User-restriction"`)
			/* At least 5 to 10 hours of reading,
			To learn more about authentication headers :
			https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Headers/WWW-Authenticate
			https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Headers/Authorization
			https://www.iana.org/assignments/http-authschemes/http-authschemes.xhtml
			*/
			http.Error(w, "Unauthorized: Missing or invalid token", http.StatusUnauthorized)
			//May replace with http.redirect later if it is anoying to get errors.
		}
		acc_id, err := ValidateJWT(token)
		if err == nil {
			r = r.WithContext(context.WithValue(
				r.Context(),
				"Account ID",
				acc_id,
			))
			//Use the context by doing r.Context().Value(key), here the key is the string "Account ID"
			next.ServeHTTP(w, r) // Token is valid, proceed to the next handler
		} else {
			//Invalid token.
			http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
		}
	})
}

func Accountmiddleware(next http.Handler) http.Handler {
	//This middleware will call the account handler no matter if the user is connected or not.
	//Depending on if the user is connected or not the page served will be different for the user.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Checking the token
		token, err := jwt.ParseRequest(r, jwt.WithCookieKey("Http-Jwt"), jwt.WithVerify(false))
		if err != nil {
			fmt.Println(err.Error())
			r = r.WithContext(context.WithValue(
				r.Context(),
				"Account ID",
				-1,
			))
		}
		acc_id, err := ValidateJWT(token)
		if err == nil {
			r = r.WithContext(context.WithValue(
				r.Context(),
				"Account ID",
				acc_id,
			))
			//Use the context by doing r.Context().Value(key), here the key is the string "Account ID"
		} else {
			fmt.Println(err.Error())
			r = r.WithContext(context.WithValue(
				r.Context(),
				"Account ID",
				-1,
			))
		}
		next.ServeHTTP(w, r)
	})
}
