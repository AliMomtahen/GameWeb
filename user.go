package main

import (
	"errors"
	"strings"
)



type User struct {
	id int
	pass   string
	name string
	info  map[string]string
}




func (this *User) init(_id int ,name string ,_pass  string) (error){
	if _pass !=  strings.TrimSpace(_pass) || strings.IndexRune(_pass  , '=') != -1 || strings.IndexRune(_pass  , '\'') != -1{
		return errors.New("bad password")
	}
	this.pass = _pass;
	this.id = _id;
	this.name = name;
	return nil; 
}

func (this *User) add_info(key  string, value string)  {
	this.info[key] = value
}