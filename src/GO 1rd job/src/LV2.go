package main

import (
	"fmt"
)

type Movie struct {
	Name         string   //电影名
	Score        Score    //评分
	director     string   //导演
	strings      []string //主演
	Types        string   //类型
	language     string   //影片语言
	Release_date string   //上映日期
}
type Score struct {
	One_star     float32 //1星
	Two_star     float32
	Three_star   float32
	Four_star    float32
	Five_star    float32
	Douban_score float32 //豆瓣评分
	Number       int
}

func (m Movie) show() {
	fmt.Println("名字：", m.Name)
	fmt.Println("导演：", m.director)
	fmt.Println("豆瓣评分：", m.Score.Douban_score)
	fmt.Println("类型：", m.Types)
	fmt.Println("语言：", m.language)
	fmt.Println("上映日期：", m.Release_date)
}
func (x Score) Score_Details() {
	fmt.Println("参与人数:", x.Number)
	fmt.Println("一星：", x.One_star, "%")
	fmt.Println("两星：", x.Two_star, "%")
	fmt.Println("三星：", x.Three_star, "%")
	fmt.Println("四星：", x.Four_star, "%")
	fmt.Println("五星：", x.Five_star, "%")
	fmt.Println("豆瓣评分：", x.Douban_score)
}
func main() {
	m := Movie{Name: "西线无战事"}
	m.Score.Douban_score = 8.5
	m.Score.One_star = 0.4
	m.Score.Two_star = 1.3
	m.Score.Three_star = 13.1
	m.Score.Four_star = 44.6
	m.Score.Five_star = 40.7
	m.Score.Number = 148069
	m.director = "爱德华·贝尔格"
	m.strings = []string{"费利克斯·卡默雷尔", "阿尔布雷希特·舒赫", "亚伦·希尔默"}
	m.Release_date = "2022-09-12(多伦多电影节) / 2022-09-29(德国) / 2022-10-28(德国网络)"
	m.Types = " 剧情 / 动作 / 战争"
	m.language = "德语 / 法语 / 英语"
	fmt.Println("请输入你的命令\n1.获得名字\n2显示影片简介\n3获取评论详情\n4获取主演名单\n5退出程序")
	var option int
	for {
		fmt.Scanf("%d", &option)
		fmt.Scanln() // 添加此行消耗换行符
		if option == 1 {
			fmt.Println(m.Name)
		} else if option == 2 {
			m.show()
		} else if option == 3 {
			m.Score.Score_Details()
		} else if option == 4 {
			fmt.Println(m.strings)
		} else {
			return
		}
	}
}
