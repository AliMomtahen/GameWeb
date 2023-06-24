package main

import (
	"database/sql"
	"errors"
	"fmt"
	
	
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	users      []User
	current_id int
	db  *sql.DB
}

func (this* Database) close(){
	this.db.Close()
}

func (this *Database) init_database() error {
	
	this.current_id = 3
	db , err := sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/test2")
	if err != nil{
		return nil
	}else{
		this.db = db;
	}
	return nil;
}

func (this *Database) find_user(name string) (*User , error)  {
	
	query_string := fmt.Sprintf("SELECT * FROM users WHERE username ='%s'" , name)
	result , err := this.db.Query(query_string)
	if err != nil {
		panic(err.Error())
	}
	
	var userr User	
	if result != nil{
		var _id int
		var _name , _pass string
		if result.Next(){
			err = result.Scan(&_id , &_name , &_pass)
		}else{
			return nil , nil
		}
		if err != nil{
			panic(err.Error())
			
		}
		userr.init(_id , _name , _pass)
	}else{
		return nil , nil
	}
	return &userr , nil
}

func (this *Database) add_user(name string, pass string) error {
	var new_user User
	
	user_find , _ := this.find_user(name)
	if user_find != nil {
		print("user alraedy exist\n")
		return errors.New("user already exist")
	}
	err := new_user.init(this.current_id, name, pass)
	if err == nil && user_find == nil {
		insert , errr := this.db.Query(fmt.Sprintf("insert into users values(%d ,'%s' , '%s' )" ,
					new_user.id , new_user.name , new_user.pass))
		if errr != nil{
			panic(errr.Error())
		}
		defer insert.Close()

	}
	return err

}