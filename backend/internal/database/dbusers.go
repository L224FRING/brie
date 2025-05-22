package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) CreateUser(id uuid.UUID ,username string, password string ) error{
    hashed_password, err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
    if err!=nil{
        return err
    }
    _,err=s.db.Exec("INSERT INTO users(id,username,password) VALUES(?,?,?)",id.String(),username,string(hashed_password))
    if err!=nil{
        return err
    }
    return nil
}

func (s *service) VerifyUser(username string, password string) (bool,error){
    var hashedPassword string
    err:=s.db.QueryRow("SELECT password FROM users WHERE username = ?",username).Scan(&hashedPassword)
    if err!=nil{
        if errors.Is(err,sql.ErrNoRows){
            return false,nil
        }
        return false,err
    }

    err=bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(password))
    if err!=nil{
        return false,nil
    }
    return true,nil
}

func (s *service) GetUserID(username string) (uuid.UUID,error) {
    var idstr string
    err:=s.db.QueryRow("SELECT id FROM users WHERE username = ?",username).Scan(&idstr)
    if err!=nil{
        if errors.Is(err,sql.ErrNoRows){
            return uuid.Nil, fmt.Errorf("User ID Not Found")
        }
        return uuid.Nil,err
    }
    id,err:=uuid.Parse(idstr)
    if err!=nil{
        return uuid.Nil,fmt.Errorf("Invalid UUID Format: %v",err)
    }
    return id,nil
}


