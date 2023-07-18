package dao

import (
	"GO_3rd_job/model"
	"io/ioutil"
)

var database = map[string]string{
	"yxh": "123456",
	"wx":  "654321",
}

func AddUser(username, password string) {
	database[username] = password
}

func SelectUser(username string) bool {
	if database[username] == "" {
		return false
	}
	return true
}

func SelectPasswordFromUsername(username string) string {
	return database[username]
}

// 保存用户数据到文件中
func (db *FileDB) SaveUser(user User) error {

	// 将用户数据转为JSON格式
	userData, err := json.Marshal(user)
	if err != nil {
		return err
	}

	// 将数据写入文件
	err = ioutil.WriteFile(db.FilePath, userData, 0644)
	if err != nil {
		return err
	}

	return nil
}

// 从文件中获取用户数据
func (db *FileDB) GetUser() (User, error) {
	var user User

	// 读取文件中的数据
	userData, err := ioutil.ReadFile(db.FilePath)
	if err != nil {
		return user, err
	}

	// 解析JSON数据为User结构体
	err = json.Unmarshal(userData, &user)
	if err != nil {
		return user, err
	}

	return user, nil
}
