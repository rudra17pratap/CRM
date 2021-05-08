/*
	- InsertCustomer method () use to insert data
	- vardiac func = function in which we pass an infinite number of arguments, instead of passing them one at a time
	- Vardiac func is being implemneted in the main method ()
	- rest all the other method from the previous get_customer () section
*/
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Customer Class
type Customer struct {
	CustomerId   int
	CustomerName string
	SSN          string
}

//GetConnection method ret sql.DB
func GetConnection() (database *sql.DB) {
	databaseDriver := "mysql"
	databaseUser := "newuser"
	databasePass := "newuser"
	databaseName := "crm"
	database, error := sql.Open(databaseDriver, databaseUser+":"+databasePass+"@/"+databaseName)
	// stops the execution if there is any error in connecting
	if error != nil {
		panic(error.Error())
	}
	return database
}

//GetCustomers methos returns Cutomer Array
func GetCustomers() []Customer {
	var database *sql.DB
	database = GetConnection()

	// sql cmd to get the customers
	var error error
	var rows *sql.Rows
	rows, error = database.Query("SELECT * FROM  Customer ORDER BY Customerid DESC")
	if error != nil {
		panic(error.Error())
	}

	// define single customer object from struct
	var customer Customer
	customer = Customer{}
	// define array of customers objects from struct
	var customers []Customer
	customers = []Customer{}

	// Scan each rows
	for rows.Next() {
		var customerId int
		var customerName string
		var ssn string
		// retrive the val of the DB table using the addres of the value
		// scan the result into customer object
		error = rows.Scan(&customerId, &customerName, &ssn)
		if error != nil {
			panic(error.Error())
		}
		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.SSN = ssn
		// append the set of output in the array of customers struct define above
		customers = append(customers, customer)
	}
	//  defer function or method call arguments evaluate instantly, but they execute until the nearby functions returns
	defer database.Close()
	// return the customer array
	return customers
}

// Insert Customer in the table
func InsertCustomer(customer Customer) {
	var database *sql.DB
	database = GetConnection()

	var error error
	var insert *sql.Stmt
	insert, error = database.Prepare("INSERT INTO Customer(CustomerId, CustomerName, SSN) VALUES(?,?,?)")
	if error != nil {
		panic(error.Error())
	}
	insert.Exec(customer.CustomerId, customer.CustomerName, customer.SSN)
	defer database.Close()
}

// vardiac func implementation
func main() {
	var customers []Customer
	customers = GetCustomers()
	fmt.Println("Before Insert", customers)
	var customer Customer
	customer.CustomerId = 5
	customer.CustomerName = "Alok Kumar"
	customer.SSN = "2367343"
	InsertCustomer(customer)
	customers = GetCustomers()
	fmt.Println("After Insert", customers)
}

// go run insert_customer.go
// Before Insert [{4 Amit Kumar 2386343} {3 ARJUN KUMAR 234569} {2 ASHISH MOHANTY 67890} {1 RUDRA PRATAP 12345}]
// After Insert [{5 Alok Kumar 2367343} {4 Amit Kumar 2386343} {3 ARJUN KUMAR 234569} {2 ASHISH MOHANTY 67890} {1 RUDRA PRATAP 12345}]
