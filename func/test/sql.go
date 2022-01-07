package test

import (

    "fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"unittest/model"

)


func connMysql()(db *sql.DB){
	db, err := sql.Open("mysql", "rd:test1234@tcp(mysql:3306)/member")
	if err != nil {
		panic(err)
	} else{
		fmt.Printf("connect success !")
		return db
	}
	
}

func GetUserById(id string)(data model.User){
	db := connMysql()
	rows, err := db.Query("SELECT id, username, password, create_time FROM users WHERE id = ? ", id)
	if err != nil {
		panic(err)
	}

	//切記用完都要做 Close
	defer rows.Close()
	//var array = []model.User{}
	var u model.User
	for rows.Next() {
		err := rows.Scan(&u.Id, &u.Username, &u.Password, &u.Create_time)
		if err != nil {
			panic(err)
		}
		fmt.Printf("id is %s, username is %s, password is %s, create_date is %s \n", u.Id,  u.Username, u.Password,  u.Create_time)
		//array = append(array, u)
	}
	return u
}


func GetUserList()(data []model.User){
	db := connMysql()
	rows, err := db.Query("SELECT * FROM users ");
	if err != nil {
		panic(err)
	}

	//切記用完都要做 Close
	defer rows.Close()
	var array = []model.User{}
	for rows.Next() {
		var u model.User
		err := rows.Scan(&u.Id, &u.Username, &u.Password, &u.Phone, &u.Email, &u.Create_time)
		if err != nil {
			panic(err)
		}
		fmt.Printf("id is %s, username is %s, password is %s, create_date is %s \n", u.Id,  u.Username, u.Password,  u.Create_time)
		array = append(array, u)
	}
	return array
}

