package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	//打开数据库
	//DSN数据源字符串：用户名:密码@协议(地址:端口)/数据库?参数=参数值
	db, err := sql.Open("mysql", "root:mytest@tcp(120.78.144.214:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
   // insert(db)

	//selectd(db)

	//updateUser(db)
	deletcUser(db)
}

//添加操作
func insert(db *sql.DB){

	stmt, err := db.Prepare("INSERT INTO user(name, age) VALUES(?, ?)")
	defer stmt.Close()

	if err != nil {
		log.Println(err)
		return
	}
	stmt.Exec("乔峰", "25")
	stmt.Exec("段誉", "25")
	stmt.Exec("虚竹","26")
}
//查询操作
func selectd(db *sql.DB){


	rows, err := db.Query("SELECT * FROM user")
	defer rows.Close()
	if err != nil {
		log.Print(err)
		return
	}

	for rows.Next() {

		var id int
		var name string
		var age string
		rows.Columns()
		err = rows.Scan(&id, &name, &age)

		fmt.Println("序号：",id)
		fmt.Println("名字：",name)
		fmt.Println("年龄：",age)
	}

}
//修改操作
func  updateUser(db *sql.DB)  {

	stmt, err := db.Prepare("UPDATE user SET name=?,age=? WHERE id=?")

	defer stmt.Close()
	if err != nil{
		log.Print(err)
		return
	}
	res, err := stmt.Exec("丐帮帮主", "79", 1)

	num, err := res.RowsAffected()

	fmt.Println(num)

}

//删除操作
func deletcUser(db *sql.DB){

	stmt, err := db.Prepare(`DELETE FROM user WHERE id=?`)
	defer stmt.Close()

	if err != nil{
		log.Print(err)
		return
	}
	res, err := stmt.Exec(1)

	num, err := res.RowsAffected()

	fmt.Println(num)


}