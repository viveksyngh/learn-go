package main 

import "fmt"
import "errors"

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Can't work with 42")
	}

	return arg + 3, nil
}

//Custom Errors

type argError struct {
	value int
	message string
}

func (err *argError) Error() string {
	return fmt.Sprintf("%d - %s", err.value, err.message) 
}


func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "Can't work with it"}
	}
	return arg + 3, nil
}


func main() {

	for _, i := range [] int{7, 42} {
		if val, err := f1(i); err != nil {
			fmt.Println("f1 failed: ", err)
			} else {
				fmt.Println("f1 worked: ", val)
			}
	}

	for _, i := range[] int{7, 42} {
		if val, err := f1(i); err != nil {
			fmt.Println("f2 failed : ", err)
		} else {
			fmt.Println("f2 failed : ", val)
		}
	}

}


