package dao

import (
	"database/sql"
	"fmt"
	"log"
	"picking/model"
)

func GetUserById(userId string) (*model.User, error) {
	user := &model.User{}
	err := db.QueryRow("select user_id, nick_name, role, login_time from users where user_id=?", userId).Scan(&user.UserId, &user.NickName, &user.Role, &user.LoginTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Printf("get user failed, err:%v\n", err)
		return nil, err
	}
	return user, nil
}

func AddUser(user model.User) error {
	sqlStr := "insert into users (user_id, nick_name, role, login_time) values (?, ?, ?, ?)"
	_, err := db.Exec(sqlStr, user.UserId, user.NickName, user.Role, user.LoginTime)
	if err != nil {
		log.Printf("insert user failed, err:%v\n", err)
		return err
	}

	return nil
}

func CheckUserExist(userId string) error {
	sqlStr := "SELECT EXISTS (SELECT * FROM users WHERE user_id = ?)"
	var exist bool
	err := db.QueryRow(sqlStr, userId).Scan(&exist)
	if err != nil {
		log.Printf("check user failed, err:%d\n", err)
		return err
	}

	if exist {
		log.Printf("user already exist")
		return fmt.Errorf("user already exist, check your userId")
	}

	return nil
}

func DelUser(userId string) error {
	sqlStr := "delete from users where user_id=?"
	_, err := db.Exec(sqlStr, userId)
	if err != nil {
		log.Printf("delete user failed, err:%v\n", err)
		return err
	}

	return nil
}

func GetUserList() ([]model.User, error) {
	sqlStr := "select * from users"
	rows, err := db.Query(sqlStr)
	if err != nil {
		log.Printf("query user list failed, err:%v\n", err)
		return nil, err
	}

	users := make([]model.User, 0)
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.UserId, &user.NickName, &user.Role, &user.LoginTime)
		if err != nil {
			log.Printf("scan user list failed, err:%v\n", err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
