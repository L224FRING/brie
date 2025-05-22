package database

func (s *service) CreateMessage(sendername string ,receivername string, message string ) error{
    //_,err=s.db.Exec("INSERT INTO users(id,username,password) VALUES(?,?,?)",id.String(),username,string(hashed_password))
	_,err:=s.db.Exec("INSERT INTO messages(sender,receiver,message) VALUES(?,?,?)",sendername,receivername,message);
	if err!=nil{
		return err
	}
    return nil
}

