/*
	- CRM web application
		- Home Function
		- Create Function
		- Insert Function
		- Alter Function
		- Update Function
		- Delete Function
		- View Function
	- Pre-req
		- http://httpd.apache.org/download.cgi#apache24 OR https://www.apachefriends.org/index.html OR https://www.apachelounge.com/download/ (install this)
		- https://www.youtube.com/watch?v=oJnCEqeAsUk (check first comment if issue)
		- FOR FQDN issue = https://kb.wisc.edu/iam/page.php?id=39180 or Simply Change Computer Name , eg: rudra
		- To test apache - Open Chrome = http://localhost
		- To test the program - Open Chrome = http://localhost:8000
		- Set of templates needs to be created in the same location as per the structure
	- To test run this file only
		// go run crm_application.go crm_database_operations.go
		// and open web browser - http://localhost:8080
*/
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// for all the info for template package = https://golang.org/doc/articles/wiki/

var template_html = template.Must(template.ParseGlob("templates/*"))

// In every 7 function below we have two arguments :
// One with type http.ResponseWriter and its corresponding response stream, which is actually an interface type.
// The second is *http.Request and its corresponding HTTP request.
// Home function executes the "Home" template with writer param and customers array
func Home(writer http.ResponseWriter, request *http.Request) {
	var customers []Customer
	customers = GetCustomers()
	log.Println(customers)
	template_html.ExecuteTemplate(writer, "Home", customers)
}

// Create func takes "writer" and "request" param to render "Create" template
func Create(writer http.ResponseWriter, request *http.Request) {
	template_html.ExecuteTemplate(writer, "Create", nil)
}

// Insert function invokes GetCustomers() to get []customers
// Then renders the "Home" template with "writer" and "customers" arr as param by invking ExecuteTemplate()
func Insert(writer http.ResponseWriter, request *http.Request) {
	var customer Customer
	customer.CustomerName = request.FormValue("customername")
	customer.SSN = request.FormValue("ssn")
	InsertCustomer(customer)

	var customers []Customer
	customers = GetCustomers()
	template_html.ExecuteTemplate(writer, "Home", customers)
}

// Alter function renders the "Home" template with "writer" and "customers" arr as param
func Alter(writer http.ResponseWriter, request *http.Request) {
	var customer Customer
	var customerId int
	var customerIdStr string
	customerIdStr = request.FormValue("id")
	fmt.Sscanf(customerIdStr, "%d", &customerId)
	customer.CustomerId = customerId
	customer.CustomerName = request.FormValue("customername")
	customer.SSN = request.FormValue("ssn")
	UpdateCustomer(customer)
	var customers []Customer
	customers = GetCustomers()
	template_html.ExecuteTemplate(writer, "Home", customers)

}

// Update func invoke "ExecuteTemplate"  with "writer" and "customer" by "id" which renders "UPDATE" template
func Update(writer http.ResponseWriter, request *http.Request) {

	var customerId int
	var customerIdStr string
	customerIdStr = request.FormValue("id")
	fmt.Sscanf(customerIdStr, "%d", &customerId)
	var customer Customer
	customer = GetCustomerById(customerId)

	template_html.ExecuteTemplate(writer, "Update", customer)

}

// Delete() renders "Home" template after deleting with "CustomerByID"
func Delete(writer http.ResponseWriter, request *http.Request) {
	var customerId int
	var customerIdStr string
	customerIdStr = request.FormValue("id")
	fmt.Sscanf(customerIdStr, "%d", &customerId)
	var customer Customer
	customer = GetCustomerById(customerId)
	DeleteCustomer(customer)
	var customers []Customer
	customers = GetCustomers()
	template_html.ExecuteTemplate(writer, "Home", customers)

}

// View() rendesr "View" template by invoking "GetCustomerByID"
func View(writer http.ResponseWriter, request *http.Request) {
	var customerId int
	var customerIdStr string
	customerIdStr = request.FormValue("id")
	fmt.Sscanf(customerIdStr, "%d", &customerId)
	var customer Customer
	customer = GetCustomerById(customerId)
	fmt.Println(customer)
	var customers []Customer
	customers = []Customer{customer}
	//customers.append(customer)
	template_html.ExecuteTemplate(writer, "View", customers)

}

// main method handles Home, Alter, Create, Update, View, Insert, and Delete functions
func main() {
	// called the function http.HandleFunc from the package net/http to register another function to be the handle 7 functions
	log.Println("Server started on: http://localhost:8000")
	// accepts 2 args - is a string type pattern, which is the route you want to match and it’s the root path in the example.
	// The second argument is a function with the signature func(ResponseWriter, *Request)(https://gowalker.org/net/http#HandleFunc)
	http.HandleFunc("/", Home)
	http.HandleFunc("/alter", Alter)
	http.HandleFunc("/create", Create)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/view", View)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/delete", Delete)
	// http.ListenAndServe(":8000", nil) function to listen on localhost with port 8000
	http.ListenAndServe(":8000", nil)
}

// go run crm_application.go crm_database_operations.go
// and open web browser - http://localhost:8080
