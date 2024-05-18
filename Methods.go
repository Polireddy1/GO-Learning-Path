package main

import "fmt"

// Author structure
type author struct {
	name	string
	branch  string
	articles  int
	salary  int                                     //METHODS
}

// Method with a receiver
// of author type
func (a author )  show() {

	fmt.Println("Author's Name: ", a.name)
	fmt.Println("Branch Name: ", a.branch)
	fmt.Println("Published articles: ", a.articles)
	fmt.Println("Salary: ", a.salary)
}

// Main function
func main() {

	// Initializing the values
	// of the author structure
	res := author{
		name:	 "Sona",
		branch: "CSE",
		articles: 203,
		salary: 34000,
	}

	// Calling the method
	res.show()
}
