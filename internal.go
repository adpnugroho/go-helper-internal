package go_helper_internal

import "fmt"

func Internal() {
	fmt.Println("internal")
}

func CustomFunction(data string) string {
	return "Ini Datamu " + data
}
