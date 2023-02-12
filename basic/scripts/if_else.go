package main

import ("fmt")

func main() {
	_, errors := fmt.Println("Hello, World!")
	if errors != nil{
		fmt.Println("O seguinte erro foi identificado:\n",errors)
	} else {
		fmt.Println("Nenhum erro foi detectado!")
	}
}
