Go Project
-----------------------
A Go project have a folder structure like followings:
```
goproject/          
    main.go             - The main entry code 
    helper/             - The sub package folder 
        helper.go       - Helper functions 
        constants.go    - Constant Definition      
```

## Initiate the Go Module 
```
go mod init goproject
```
After the execution, a file `go.mod` is created.

Reference: https://go.dev/doc/tutorial/create-module

## Notes 
- Sub-package helps code files easy to maintain.
- Every sub-package is holded by file folders.
- Within the same sub-package, cannot have two variable or function have the same name.

## Problems
- Showing 'undeclare name' in VSCode if using the variable name in the different file. 