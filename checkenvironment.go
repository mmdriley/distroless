package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	fail := false

	// Fails unless one of:
	//  - /etc/passwd exists and has a matching UID
	//  - cgo is enabled
	//  - $USER or $HOME is set
	if currentUser, err := user.Current(); err == nil {
		fmt.Printf("Current user: %v\n", currentUser.Username)
	} else {
		fmt.Printf("Could not get current user: %v\n", err)
		fail = true
	}

	// Fails if /tmp does not exist
	if _, err := os.MkdirTemp("", "checkenvironment"); err != nil {
		fmt.Printf("Could not create temporary directory: %v\n", err)
		fail = true
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

	if fail {
		fmt.Printf("===\nOne or more tests failed.\n")
		os.Exit(0)
	} else {
		fmt.Printf("===\nEverything is fine.\n")
		os.Exit(1)
	}
}
