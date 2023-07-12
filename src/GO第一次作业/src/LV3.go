package main

import (
	"fmt"
)

type Hero interface {
	Say()                 //选择该英雄时的语音
	Attack(target Hero)   //普通攻击
	UseSkill(target Hero) //使用技能
	UseItem()             //使用物品
}

type Yasuo struct {
	Name      string
	HP        int
	Mana      int
	SkillDmg  int
	AttackDmg int
}

type GongBen struct {
	Name      string
	HP        int
	Mana      int
	SkillDmg  int
	AttackDmg int
}

func (Yasuo) Say() {
	fmt.Println("我是托儿索")
}

func (GongBen) Say() {
	fmt.Println("我是宫本")
}

func (y Yasuo) Attack(target Hero) {
	t := target.(GongBen)
	t.HP -= y.AttackDmg
	fmt.Printf("%s普通攻击攻击%s，造成%d点伤害\n", y.Name, target.(GongBen).GetName(), y.AttackDmg)
}

func (y Yasuo) UseSkill(target Hero) {
	if y.Mana >= 50 {
		y.Mana -= 50
		t := target.(GongBen)
		t.HP -= y.SkillDmg
		fmt.Printf("%s使用技能狂风绝息斩攻击%s，造成%d点伤害\n蓝量减少50，现在蓝量为：%d\n", y.Name, target.(GongBen).GetName(), y.SkillDmg, y.Mana)
	} else {
		fmt.Printf("%s法力值不足，无法使用技能\n", y.Name)
	}
}

func (y Yasuo) UseItem() {
	y.AttackDmg += 10
	fmt.Printf("%s使用物品，增加自身攻击力10点\n", y.Name)
}

func (g GongBen) Attack(target Hero) {
	t := target.(Yasuo)
	t.HP -= g.AttackDmg
	fmt.Printf("%s普通攻击攻击%s，造成%d点伤害\n", g.Name, target.(Yasuo).GetName(), g.AttackDmg)
}

func (g GongBen) UseSkill(target Hero) {
	if g.Mana >= 40 {
		g.Mana -= 40
		t := target.(Yasuo)
		t.HP -= g.SkillDmg
		fmt.Printf("%s使用技能一决生死攻击%s，造成%d点伤害\n蓝量减少40，现在蓝量为：%d\n", g.Name, target.(Yasuo).GetName(), g.SkillDmg, g.Mana)
	} else {
		fmt.Printf("%s法力值不足，无法使用技能\n", g.Name)
	}
}

func (g GongBen) UseItem() {
	g.AttackDmg += 15
	fmt.Printf("%s使用物品，增加自身攻击力15点\n", g.Name)
}

func (h Yasuo) GetName() string {
	return h.Name
}

func (h GongBen) GetName() string {
	return h.Name
}

func main() {
	var hero1 Hero = Yasuo{Name: "托儿索", HP: 300, Mana: 100, SkillDmg: 80, AttackDmg: 50}
	var hero2 Hero = GongBen{Name: "宫本", HP: 300, Mana: 100, SkillDmg: 90, AttackDmg: 40}
	var i, j, k int
	fmt.Println("请输入你的命令\n1.选择你的英雄\n2.退出游戏")
	for {
		fmt.Scan(&i)
		if i == 1 {
			fmt.Println("选择你的英雄:\n1", hero1.(Yasuo).Name, "\n2", hero2.(GongBen).Name)
			for {
				fmt.Scan(&j)
				if j == 1 {
					hero1.Say()
					fmt.Println("请输入你的命令\n1.普通攻击\n2.释放技能\n3.使用物品\n4.退出游戏")
					for {
						fmt.Scan(&k)
						if k == 1 {
							hero1.Attack(hero2)
						} else if k == 2 {
							hero1.UseSkill(hero2)
						} else if k == 3 {
							hero1.UseItem()
						} else if k == 4 {
							fmt.Println("玩家已退出游戏")
							return
						}
					}
				} else if j == 2 {
					hero2.Say()
					fmt.Println("请输入你的命令\n1.普通攻击\n2.释放技能\n3.使用物品\n4.退出游戏")
					for {
						fmt.Scan(&k)
						if k == 1 {
							hero2.Attack(hero1)
						} else if k == 2 {
							hero2.UseSkill(hero1)
						} else if k == 3 {
							hero2.UseItem()
						} else if k == 4 {
							fmt.Println("玩家已退出游戏")
							return
						}
					}
				} else {
					fmt.Println("请输入正确的序号")
				}
			}

		} else if i == 2 {
			fmt.Println("玩家已退出游戏")
			return
		}
	}
}
