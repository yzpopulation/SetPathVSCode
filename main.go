package main

import (
	"fmt"
	"os"

	"golang.org/x/sys/windows/registry"
)

var CodePath = ``

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Args must 1")
		return
	}
	CodePath = os.Args[1]
	exist, _ := PathExists(CodePath)
	if !exist {
		fmt.Println("File is not exists! ==>" + CodePath)
		return
	}
	SetPath("*\\shell\\VSCode")
	SetPath("Directory\\Background\\shell\\VSCode")
	SetPath("Directory\\shell\\VSCode")
	SetPath("Drive\\shell\\VSCode")
}
func SetPath(basePath string) {
	key, exist, err := registry.CreateKey(registry.CURRENT_USER, "SOFTWARE\\Classes\\"+basePath, registry.ALL_ACCESS)
	if err != nil {
		fmt.Println(err)
	}
	defer key.Close()
	if exist {
		fmt.Println(`键已存在`)
	} else {
		fmt.Println(`新建注册表键`)
		key.SetExpandStringValue("", "通过 Code 打开")
		key.SetExpandStringValue("Icon", CodePath)
		key2, exist2, err := registry.CreateKey(registry.CURRENT_USER, "SOFTWARE\\Classes\\"+basePath+"\\command", registry.ALL_ACCESS)
		if err != nil {
			fmt.Println(err)
		}
		defer key2.Close()
		if exist2 {
			fmt.Println(`command键已存在`)
		} else {
			fmt.Println(`新建注册表command键已存在键`)
			key2.SetExpandStringValue("", "\""+CodePath+"\" \"%V\"")
		}
	}

}
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
