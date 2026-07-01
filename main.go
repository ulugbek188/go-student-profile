package main

import "fmt"

// Application constants
const (
	University = "PDP University"
	Country    = "Uzbekistan"
	StudentID  = 250022
)

// returns basic student information
func studentInfo() (string, int, string) {
	name := "Ulug'bek"
	age := 18
	major := "Backend Engineering"
	return name, age, major
}

// returns body information
func studentBody() (float64, float64) {
	height := 1.78
	weight := 61.9
	return height, weight
}

// returns academic information
func studentScore() float64 {
	GPA := 4.8
	return GPA
}

func studentContact() (int, string) {
	phone := +998777777777
	email := "zokirovu779@gmail.com"
	return phone, email
}

func main() {

	// Get student information
	name, age, major := studentInfo()

	// Get body information
	height, weight := studentBody()

	// Get average score
	GPA := studentScore()

	// Get contact
	phone, email := studentContact()

	// Print student profile
	fmt.Println("==============================")
	fmt.Println("      STUDENT PROFILE")
	fmt.Println("==============================")

	fmt.Println("StudentID : ", StudentID)
	fmt.Println("Name      : ", name)
	fmt.Println("Age       : ", age)
	fmt.Println("Country   : ", Country)
	fmt.Println("University: ", University)
	fmt.Println("Major     : ", major)
	fmt.Println("Height    : ", height)
	fmt.Println("Weight    : ", weight)
	fmt.Println("GPA       : ", GPA)
	fmt.Println("Phone     : ", phone)
	fmt.Println("Email     : ", email)

	fmt.Println("==============================")
}
