package error 

import "fmt"

func Print_Error(error_message string) {
	fmt.Println("golb: ", error_message)
	fmt.Println("Try 'golb --help' for more information.")
}

func Check(e error) {
	if e != nil {
		Print_Error(e.Error())
		panic(e)
	}
}
