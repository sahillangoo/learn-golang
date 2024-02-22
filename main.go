package main

import (
	"fmt"
	"go/calculator/util"
	"os"
)

func main() {
	// var operation string
	// var num1, num2 int = 0, 0
	fmt.Println("calculator v0.01")
	// fmt.Println("=================")
	// fmt.Println("Which operation do you want to perform? (sum, sub, mul, div)")
	// fmt.Scanf("%s", &operation)
	// fmt.Println("Enter the first number")
	// fmt.Println("Enter the second number")
	// fmt.Scanf("%d", &num1)
	// fmt.Scanf("%d", &num2)

	// switch operation {
	// case "sum":
	// 	fmt.Println("The result is:", num1+num2)
	// case "sub":
	// 	fmt.Println("The result is:", num1-num2)
	// case "mul":
	// 	fmt.Println("The result is:", num1*num2)
	// case "div":
	// 	fmt.Println("The result is:", num1/num2)
	// }
	rootPath, _ := os.Getwd()
	filePath := rootPath + "/data/file.txt"
	c, err := util.ReadTextFile(filePath)
	if err == nil {
		fmt.Println(c)

		newContent := fmt.Sprintf("The content of the file is: %v\n", c)
		util.WriteTextFile(filePath, newContent)
	} else {
		fmt.Printf("error panic!! %v\n", err)
	}

}
