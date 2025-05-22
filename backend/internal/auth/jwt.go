package auth

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

var key []byte

func init() {
    err:=godotenv.Load()
    if err!=nil{
        log.Fatalf("Error loading .env file: %v",err)
    }
    key=[]byte(os.Getenv("JWTKEY"))
    
}


func CreateJWT(userid uuid.UUID, username string) (string,error){
    claims := jwt.MapClaims{
        "user_id" : userid.String(),
        "exp": time.Now().Add(time.Hour*24*365).Unix(),
        "iat": time.Now().Unix(),
        "username": username,
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
    signedToken,err:=token.SignedString(key)
    if err!=nil{
        return "", err
    }
    return signedToken,nil
}

