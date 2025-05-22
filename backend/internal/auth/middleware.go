package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ClaimsKey string

const ClaimsContextKey ClaimsKey = "claims"


func JWTMiddleware(next http.Handler) http.Handler{
    return http.HandlerFunc(func (w http.ResponseWriter,r *http.Request){
        // Get JWT
        cookie,err := r.Cookie("jwt_token")
        if err!=nil{
            http.Error(w,"Unauthorised: Missing Token",http.StatusUnauthorized)
            return
        }
        // Parse JWT
        tokenStr:=cookie.Value
        claims:=jwt.MapClaims{}

        token, err := jwt.ParseWithClaims(tokenStr,claims,func(token *jwt.Token) (interface{}, error) {
            return key,nil
        })
        if err!=nil || !token.Valid{
            http.Error(w,"Unauthorised: Invalid Token",http.StatusUnauthorized)
            return
        }
        exp, ok:=claims["exp"].(float64);
        if ok {
            if time.Now().Unix()>int64(exp) {
                http.Error(w ,"Token expired",http.StatusUnauthorized)
                return
            }
        } else {
            http.Error(w,"Unauthorised: Invalid Token Claims",http.StatusUnauthorized)
            return
        }
        ctx:=context.WithValue(r.Context(),ClaimsContextKey,claims)
        next.ServeHTTP(w,r.WithContext(ctx))
    })
}

