package main

import (
	"fmt"
)

type employee struct {  
    salary  int
    country string
}

func main() {

	//Declaracação de Map
	employeeSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	newEmp := "joe"
	value, ok := employeeSalary[newEmp]
	if ok {
		fmt.Println("Salary of", newEmp, "is", value)
		return
	}
	fmt.Println(newEmp, "not found")

	fmt.Println("map before deletion", employeeSalary)
	//Deletar item do Map
	delete(employeeSalary, "steve")
	fmt.Println("map after deletion", employeeSalary)

	emp1 := employee{
        salary:  12000,
        country: "USA",
    }
    emp2 := employee{
        salary:  14000,
        country: "Canada",
    }
    emp3 := employee{
        salary:  13000,
        country: "India",
    }

	employeeInfo := map[string]employee{
        "Steve": emp1,
        "Jamie": emp2,
        "Mike":  emp3,
    }

	for name, info := range employeeInfo {
        fmt.Printf("Employee: %s Salary:$%d  Country: %s\n", name, info.salary, info.country)
    }

}
