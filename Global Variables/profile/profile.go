package profile

import (
	"fmt"
	"global/variables/helpers"
)

func Profile() {
	user := helpers.User
	if len(user) == 0 {
		fmt.Println("No User Profile")
	} else {
		fmt.Println(user)
	}
}
