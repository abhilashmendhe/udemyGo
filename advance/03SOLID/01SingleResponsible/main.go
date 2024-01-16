package main

/* S - Single Responsibility Principle (SRP)

This principle states that a struct should have only one reason to change,
meaning that a struct should have only one responsibility. This helps to keep
the code clean and maintainable, as changes to the struct only need to be made
in one place.*/

/*
type Employee struct {
	Name    string
	Salary  float64
	Address string
}
According to the SRP, each struct should have only one responsibility,
so in this case, it would be better to split the responsibilities of the
Employee struct into two separate structs: EmployeeInfo and EmployeeAddress.
*/

type EmployeeInfo struct {
	Name   string
	Salary float64
}

type EmployeeAddress struct {
	Address string
}

// Now we have separate function to handle the different responsibility
func (e EmployeeInfo) GetSalary() float64 {
	return e.Salary
}

func (e EmployeeAddress) GetAddress() string {
	return e.Address
}
