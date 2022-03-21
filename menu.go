package main

import "fmt"

type Menu struct {
	id    int // 菜单id
	pid   int // 父菜单id
	name  string
	level int // 菜单层级
}

func main() {
	menuList := []Menu{
		{0, -1, "game", 0},
		{1, -1, "video", 0},
		{2, 0, "LOL", 1},
		{3, 0, "CF", 1},
		{4, 0, "DNF", 1},
		{5, 1, "TikTok", 1}}
	var level int = 0
	var pid int = -1
	fmt.Println("Choose MenuList(input id)")
	fmt.Println(menuList[0].id, menuList[0].name)
	fmt.Println(menuList[1].id, menuList[1].name)
	for true {
		var index int = 0
		var length int = 6
		fmt.Println("Input Menu id")
		fmt.Scanf("%d", &pid)
		fmt.Println("Choose MenuList(input id)")
		for index < length {
			if menuList[index].pid == pid {
				fmt.Println(menuList[index].id, menuList[index].name)
			}
			index++
		}
		level++
	}
}
