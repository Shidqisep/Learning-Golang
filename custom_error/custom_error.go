package main

import "fmt"

type validationError struct {
	Message string
}

func (e *validationError) Error() string {
	return e.Message
}

type notFoundError struct {
	Message string
}

func (e *notFoundError) Error() string {
	return e.Message
}

func saveData(id string, data any) error {
	if id == "" {
		return &validationError{"id tidak boleh kosong"}
	}

	if id != "Doni" {
		return &notFoundError{"id not found"}
	}
	return nil
}

func main() {

	// ini pake if
	err := saveData("Doni", nil)
	if err != nil {
	// 	if validationError, ok := err.(*validationError); ok{
	// 		fmt.Println("Validation Error:", validationError.Message)
	// 	} else if notFoundError, ok := err.(*notFoundError); ok{
	// 		fmt.Println("Not Found Error:", notFoundError.Message)
	// 	} else {
	// 		fmt.Println("Unknown Error:", err.Error())
	// 	}
	// }

		switch err.(type) {
		case *validationError:
			fmt.Println("Validation Error:", err.Error())
		case *notFoundError:
			fmt.Println("Not Found Error:", err.Error())
		default:
			fmt.Println("unknown error", err.Error())
		}
	}else{
		fmt.Println("success")
	}
}