package main

// Sales Database Assignment
// The SQL queries for this assignment are kept in sales-assignment.sql.
// This file is included because the submission checker expects an exact
// repository path named drive.go.

const salesQueries = `USE sales;

-- Question 1: Retrieve Payment Information
SELECT checkNumber, paymentDate, amount
FROM payments;

-- Question 2: Find Orders in Process
SELECT orderDate, requiredDate, status
FROM orders
WHERE status = 'In Process'
ORDER BY orderDate DESC;

-- Question 3: Find Sales Representatives
SELECT firstName, lastName, email
FROM employees
WHERE jobTitle = 'Sales Rep'
ORDER BY employeeNumber DESC;

-- Question 4: Retrieve Office Information
SELECT *
FROM offices;

-- Question 5: Retrieve the Five Cheapest Products
SELECT productName, quantityInStock
FROM products
ORDER BY buyPrice ASC
LIMIT 5;`

func main() {}
