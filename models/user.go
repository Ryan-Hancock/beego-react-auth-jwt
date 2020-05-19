package models

import "github.com/astaxie/beego/orm"

type User struct {
	ID       int    `json:"id" orm:"auto"`
	Username string `json:"username" orm:"size(25);unique"`
	Password string `json:"password" orm:"size(25)"`
}

type UserStorage struct {
	DB orm.Ormer
}

func GetUserStorage() *UserStorage {
	o := GetORM()
	return &UserStorage{DB: o}
}

func (s *UserStorage) NewUser(u User) (id int64, err error) {
	return s.DB.Insert(&u)
}

func (s *UserStorage) GetUser(id int) (User, error) {
	u := User{ID: id}
	return u, s.DB.Read(&u)
}

func (s *UserStorage) CheckPassword(username, password string) (int, error) {
	var u User

	err := s.DB.QueryTable("user").
		Filter("username__exact", username).
		Filter("password__exact", password).
		One(&u)

	return u.ID, err
}
