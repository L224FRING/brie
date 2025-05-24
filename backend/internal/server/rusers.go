package server

import (

	"encoding/json"
	"net/http"
	"time"
	"github.com/google/uuid"
	"brie/internal/auth"
)


type Credentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

// TODO abstract this shit
func (s *Server) LoginUser(w http.ResponseWriter, r *http.Request){
    // Getting Sign In Credentials 
    var creds Credentials
    err:=json.NewDecoder(r.Body).Decode(&creds) 
    if err!=nil{
        http.Error(w,"Invalid Request Body",http.StatusBadRequest)
        return
    }
    if creds.Username=="" || creds.Password==""{
        http.Error(w,"Missing Credentials",http.StatusUnprocessableEntity)
        return
    }

    ok, err:= s.db.VerifyUser(creds.Username,creds.Password)
    if err!=nil{
        http.Error(w,"Error Verifying User",500)
    }
    if !ok{
        http.Error(w,"Your username or password is invalid",http.StatusUnauthorized)
        return
    }

    userid, err:=s.db.GetUserID(creds.Username)
    if err!=nil{
        http.Error(w,"Error Fetching userId",http.StatusInternalServerError)
        return
    }
    // Generating JWT Token
    token, err:=auth.CreateJWT(userid,creds.Username)
    if err!=nil{
        http.Error(w,err.Error(),500)
    }
    // Setting up the Cookie
    http.SetCookie(w,&http.Cookie{
        Name: "jwt_token",
        Value: token,
        Expires: time.Now().Add(time.Hour*24*365),
        HttpOnly: true,
        Secure: false,
        SameSite: http.SameSiteLaxMode,
        Path: "/",
    })

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK) 

    jsonResponse := map[string]string{
        "message": "User Logged in successfully",
    }
    json.NewEncoder(w).Encode(jsonResponse)

}


func (s *Server) CreateUser(w http.ResponseWriter,r *http.Request){
    // Getting Sign In Credentials 
    var creds Credentials
    err:=json.NewDecoder(r.Body).Decode(&creds) 
    if err!=nil{
        http.Error(w,"Invalid Request Body",http.StatusBadRequest)
        return
    }
    if creds.Username=="" || creds.Password==""{
        http.Error(w,"Missing Credentials",http.StatusUnprocessableEntity)
        return
    }

    // Adding User Row to Database
    userid:=uuid.New()
    err=s.db.CreateUser(userid,creds.Username,creds.Password)
    if err!=nil{
        http.Error(w,err.Error(),500)
        return
    }

    // Generating JWT Token
    token, err:=auth.CreateJWT(userid,creds.Username)
    if err!=nil{
        http.Error(w,err.Error(),500)
    }
    // Setting up the Cookie
    http.SetCookie(w,&http.Cookie{
        Name: "jwt_token",
        Value: token,
        Expires: time.Now().Add(time.Hour*24*365),
        HttpOnly: true,
        Secure: true,
        SameSite: http.SameSiteLaxMode,
        Path: "/",
    })

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)  // Use 201 for resource creation

    jsonResponse := map[string]string{
        "message": "User created successfully",
    }
    json.NewEncoder(w).Encode(jsonResponse)
}

func (s *Server) LogoutUser(w http.ResponseWriter, r *http.Request){

    http.SetCookie(w,&http.Cookie{
        Name: "jwt_token",
        Value: "",
        Expires: time.Now().Add(-time.Hour),
        HttpOnly: true,
        Secure: true,
        SameSite: http.SameSiteLaxMode,
        Path: "/",
    })

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    jsonResponse := map[string]string{
        "message": "User Logged Out successfully",
    }
    json.NewEncoder(w).Encode(jsonResponse)
}
