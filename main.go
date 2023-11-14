package main

import (
	"fmt"

	"github.com/AmirMirzayi/arvancloud-interview/pkg"
)

func main() {
	processingQueue := pkg.NewMap()
	processingQueue.SetValue("amir", 30)
	fmt.Println(processingQueue.GetValue("amir"))
	fmt.Println(processingQueue.ExistedKey("amirs"))

}
