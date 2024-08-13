package main

import "fmt"

type employee struct {
	employeeID    string
	employeeeName string
	phone         string
}

func main() {

	employeeList := []employee{}

	employee1 := employee{
		employeeID:    "101",
		employeeeName: "Pradoo",
		phone:         "00000000",
	}
	employee2 := employee{
		employeeID:    "102",
		employeeeName: "Prayad",
		phone:         "00000001",
	}
	employee3 := employee{
		employeeID:    "103",
		employeeeName: "Pranee",
		phone:         "00000002",
	}
	employeeList = append(employeeList, employee1)
	employeeList = append(employeeList, employee2)
	employeeList = append(employeeList, employee3)

	fmt.Println("employee = ", employeeList)

}
