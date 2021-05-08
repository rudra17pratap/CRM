/*
	- Database operations consist of CRUD methods such as CREATE, READ, UPDATE, and DELETE customer operations.
	- These DB operations will build in 1st Part
	- defer statement defers the execution of the function until the surrounding function returns.
	- The panic function stops the current flow and control. Deferred functions are executed normally after the panic call.
*/
package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Customer Class
type Customer struct {
	CustomerId   int
	CustomerName string
	SSN          string
}

//  GetConnection method  which returns sql.DB
func GetConnection() (database *sql.DB) {
	databaseDriver := "mysql"
	databaseUser := "newuser"
	databasePass := "newuser"
	databaseName := "crm"
	database, error := sql.Open(databaseDriver, databaseUser+":"+databasePass+"@/"+databaseName)
	if error != nil {
		panic(error.Error())
	}
	return database
}

//GetCustomerById with parameter customerId returns Customer
func GetCustomerById(customerId int) Customer {
	var database *sql.DB
	database = GetConnection()
	var error error
	var rows *sql.Rows
	rows, error = database.Query("SELECT * FROM Customer WHERE CustomerId=?", customerId)
	if error != nil {
		panic(error.Error())
	}
	var customer Customer
	customer = Customer{}
	for rows.Next() {
		var customerId int
		var customerName string
		var SSN string
		error = rows.Scan(&customerId, &customerName, &SSN)
		if error != nil {
			panic(error.Error())
		}
		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.SSN = SSN
	}
	defer database.Close()
	return customer
}

// GetCustomers method returns Customer Array
func GetCustomers() []Customer {
	var database *sql.DB
	database = GetConnection()
	var error error
	var rows *sql.Rows
	rows, error = database.Query("SELECT * FROM Customer ORDER BY Customerid DESC")
	if error != nil {
		panic(error.Error())
	}
	var customer Customer
	customer = Customer{}
	var customers []Customer
	customers = []Customer{}
	for rows.Next() {
		var customerId int
		var customerName string
		var ssn string
		error = rows.Scan(&customerId, &customerName, &ssn)
		if error != nil {
			panic(error.Error())
		}
		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.SSN = ssn
		customers = append(customers, customer)
	}
	defer database.Close()
	return customers
}

// InsertCustomer method with parameter customer
func InsertCustomer(customer Customer) {
	var database *sql.DB
	database = GetConnection()
	var error error
	var insert *sql.Stmt
	insert, error = database.Prepare("INSERT INTO CUSTOMER(CustomerName,SSN) VALUES(?,?)")
	if error != nil {
		panic(error.Error())
	}
	insert.Exec(customer.CustomerName, customer.SSN)
	defer database.Close()
}

// Update Customer method with parameter customer
func UpdateCustomer(customer Customer) {
	var database *sql.DB
	database = GetConnection()
	var error error
	var update *sql.Stmt
	update, error = database.Prepare("UPDATE CUSTOMER SET CustomerName=?, SSN=? WHERE CustomerId=?")
	if error != nil {
		panic(error.Error())
	}
	update.Exec(customer.CustomerName, customer.SSN, customer.CustomerId)
	defer database.Close()
}

// Delete Customer method with parameter customer
func DeleteCustomer(customer Customer) {
	var database *sql.DB
	database = GetConnection()
	var error error
	var delete *sql.Stmt
	delete, error = database.Prepare("DELETE FROM Customer WHERE Customerid=?")
	if error != nil {
		panic(error.Error())
	}
	delete.Exec(customer.CustomerId)
	defer database.Close()
}
