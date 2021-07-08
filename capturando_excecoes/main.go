package main

import (
	"fmt"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func main() {
	employee, err := getInformation(1001)
	if err != nil {
		// Something is wrong. Do something.
	} else {
		fmt.Print(employee)
	}
}

//formas de Capturar erros

// 1
func getInformation(id int) (*Employee, error) {
	employee, err := apiCallEmployee(1000)
	if err != nil {
		return nil, err //Simply return the error to the caller.
	}
	return employee, nil
}

// 2
// func getInformation(id int) (*Employee, error) {
// 	employee, err := apiCallEmployee(1000)
// 	if err != nil {
// 		return nil, fmt.Errorf("Got an error when getting the employee information: %v", err)
// 	}
// 	return employee, nil
// }

// 3
// func getInformation(id int) (*Employee, error) {
//     for tries := 0; tries < 3; tries++ {
//         employee, err := apiCallEmployee(1000)
//         if err == nil {
//             return employee, nil
//         }

//         fmt.Println("Server is not responding, retrying ...")
//         time.Sleep(time.Second * 2)
//     }

//     return nil, fmt.Errorf("server has failed to respond to get the employee information")
// }

// 4
// var ErrNotFound = errors.New("Employee not found!")

// func getInformation(id int) (*Employee, error) {
//     if id != 1001 {
//         return nil, ErrNotFound
//     }

//     employee := Employee{LastName: "Doe", FirstName: "John"}
//     return &employee, nil
// }

func apiCallEmployee(id int) (*Employee, error) {
	employee := Employee{LastName: "Doe", FirstName: "John"}
	return &employee, nil
}
