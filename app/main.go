package main

import (
		"fmt"
		"app/controller"
)

 func main (){
	 name := controller.Input("type your name")
	 fmt.Println("Hello,"+name + "!!")
 }

