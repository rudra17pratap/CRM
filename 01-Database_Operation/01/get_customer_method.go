/*
	- GetCustomers retrives "customer" data from database
	- Table name - "customer" , with = Customerid, CustomerName and SSN attributes
	- GetConnection retrives the database connection
	- returns the rows from database which queried

	- Installations pre-req:
		- https://www.python.org/ftp/python/3.9.5/python-3.9.5-amd64.exe   OR  directly type python3 on terminal and download from MS store
		- pip install mysql-connector-python
		- https://dev.mysql.com/downloads/windows/
		- https://github.com/go-sql-driver/mysql
			- https://mathaywardhill.com/2017/04/27/get-started-with-golang-and-sql-server-in-visual-studio-code/#:~:text=Installation%20is%20straightforward%20just%20double,shortcut%20Ctrl%2BShift%2BX.
		- You can do the Create table process from the Workbench also:
			- https://dev.mysql.com/doc/workbench/en/wb-getting-started-tutorial-creating-a-model.html
		- OR
		- CREATE TABLE `crm`.`customer` (
			`CustomerId` INT NOT NULL,
			`CustomerName` VARCHAR(45) NOT NULL,
			`SSN` VARCHAR(45) NOT NULL,
			PRIMARY KEY (`CustomerId`));
		- OR
		- CRETE, UPDATE, DELETE (WORKBENCH) = https://www.youtube.com/watch?v=qb7abQ6ROy4

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

// main method
func main() {
	// create the array of customer var
	var customers []Customer
	// ret the customer array
	customers = GetCustomers()
	fmt.Println("Customers", customers)
}

// go run get_customer_method.go
// Customers [{3 ARJUN KUMAR 234569} {2 ASHISH MOHANTY 67890} {1 RUDRA PRATAP 12345}]
