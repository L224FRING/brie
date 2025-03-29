package models

import (
    "github.com/google/uuid"
)


type User struct {
    ID  uuid.UUID
    Username string
    Password string
}

func CreateUser(userID uuid.UUID,username string,password string) User{
    return User{
        ID: userID,
        Username: username,
        Password: password,
    }
}

