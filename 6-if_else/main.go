package main

import "fmt"

func main() {
	/*

	     age := 18
	     hasPermissions := false

	   	if age >= 18 && hasPermissions {
	   		fmt.Println("Person is an adult")
	   	} else if age >= 12 {
	   		fmt.Println("Person is a teenager")
	   	} else {
	   		fmt.Println("Person is a kid")
	   	}
	*/

	if age, hasPermissions := 15, false; age >= 15 && hasPermissions {
		fmt.Println("has permissions and person is an adult")
	} else if age >= 12 {
		fmt.Println("Person is a teenager")
	}

	// Go does not have ternary. :(

}
