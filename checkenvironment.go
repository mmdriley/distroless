package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	if _, err := user.Current(); err != nil {
		fmt.Printf("Could not get current user: %v\n", err)
		return
	}

	// Fails if /tmp does not exist
	if _, err := os.MkdirTemp("", "checkenvironment"); err != nil {
		fmt.Printf("Could not create temporary directory: %v\n", err)
		return
	}

	// if homedir, err := os.UserHomeDir(); err != nil {
	// 	fmt.Printf("Could not get user homedir: %v\n", err)
	// 	return
	// } else {
	// 	fmt.Printf("User homedir is %v\n", homedir)
	// }

	// if groups, err := os.Getgroups(); err != nil {
	// 	fmt.Printf("Could not get user groups: %v\n", err)
	// } else {
	// 	fmt.Printf("User groups are: %v\n", groups)
	// }

	// if _, err := http.Get("http://www.neverssl.com"); err != nil {
	// 	fmt.Printf("Can't retrieve HTTP URL: %v\n", err)
	// 	return
	// }

	// if _, err := http.Get("https://www.github.com"); err != nil {
	// 	fmt.Printf("Can't retrieve HTTPS URL: %v\n", err)
	// 	return
	// }

	fmt.Printf("Everything is fine\n")
}
