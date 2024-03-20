package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"validation error"}
	}

	if id != "ucup" {
		return &notFoundError{"data not found"}
	}

	return nil
}

func main() {
	err := SaveData("ucup", nil)
	if err != nil {
		// terjadi error
		// if validationErr, ok := err.(*validationError); ok {
		// 	fmt.Println("validation error: ", validationErr.Error())
		// 	fmt.Println("validation error: ", validationErr.Error())
		// } else if notfoundErr, ok := err.(*notFoundError); ok {
		// 	fmt.Println("not found error: ", notfoundErr.Error())
		// } else {
		// 	fmt.Println("unknown error: ", err.Error())
		// }

		// menggunakan switch case
		switch finalErr := err.(type) {
		case *validationError:
			fmt.Println("validation error: ", finalErr.Error())
		case *notFoundError:
			fmt.Println("not found error: ", finalErr.Error())
		default:
			fmt.Println("unknown error: ", err.Error())
		}
	} else {
		// success
		fmt.Println("Success")
	}
}
