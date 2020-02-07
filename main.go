// Programmer name: Brooke Mason
// Date completed: 2/6/20
// Description: Lab 1.4.1

package main

import "fmt" 

func main() {
 //declare variable for favorite food and store your favorite food.
  var food = "Pizza"
  fmt.Println(food)
  //declare variables for name and age (make sure they are appropriate data types)
  var name string
  var age int
//ask the user to enter their answer for name and age.
fmt.Println(" Please enter your answer for your name and age.")
fmt.Scanln(&name, &age)
 //output a welcome statement using their name
 fmt.Println("Welcome", name,"!")
 //output a statement that says At their age you enjoyed the favorite food
 fmt.Println("At", age,"your favorite food was", food, "!")
}
 
