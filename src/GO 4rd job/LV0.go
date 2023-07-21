package main

import (
	"GO_4rd_job/modle"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func initDB() {

	var err error
	dsn := "root:zlt18716333190@tcp(127.0.0.1:3306)/school?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := dsn.DSN
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {

		log.Println(err)
	}

	log.Println("connect success:")
}
func main() {
	initDB()
	var sts = []modle.Student{{Name: "老王1", Sex: "男", Age: "15"}, {Name: "老王2", Sex: "男", Age: "15"},
		{Name: "老王3", Sex: "男", Age: "15"}, {Name: "老王4", Sex: "男", Age: "15"}, {Name: "老王5", Sex: "男", Age: "15"},
		{Name: "老王6", Sex: "男", Age: "15"}, {Name: "老王7", Sex: "男", Age: "15"}, {Name: "老王8", Sex: "男", Age: "15"},
		{Name: "老王9", Sex: "男", Age: "15"}, {Name: "老王10", Sex: "男", Age: "15"}}
	db.CreateInBatches(sts, 10)
	for i := 1; i <= 10; i++ {
		st := modle.Student{}
		//查询第一条记录
		db.First(&st, i)
		log.Println("success find:", st)
	}
}
