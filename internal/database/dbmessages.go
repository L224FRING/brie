package database

func (s *service) CreateMessage(sendername string ,recerivername string, message string ) error{
    //_,err=s.db.Exec("INSERT INTO users(id,username,password) VALUES(?,?,?)",id.String(),username,string(hashed_password))
	_,err:=s.db.Exec("INSERT INTO messagea(sender,receiver,message) VALUES(?,?,?)",sendername,recerivername,message);
	if err!=nil{
		return err
	}
    return nil
}

