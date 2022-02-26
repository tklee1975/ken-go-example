Hello World of Go 
-----------------------

To run the Go program: 
Go to the terminal: 
```
go run helloworld.go
```

Notes:
It's important to make a file with the package `package main` and the function `func main`.  

If not, it'll show:
```
package command-line-arguments is not a main package
```

The name of the main package go file is not important. Some service may require a specific name. For example: AWS beanstalk will the name as `application.go`.