# Overview of phone utils library
Phone utils has a simple function:
1) IsEasyDial - Accepts a number as string and returns true if the number is
 easy dial else false
 
Quick execution of the util can be seen in go playground https://play.golang.org/p/dVn3YhFKmFN

## Assumptions:   
1) Phone util assumes a standard phone keyboard layout is used. Hence the keymap is calculated based on that layout

2) Phone number string is standardized by removing all non numerical characters.
     
3) Logging with log levels are helpful during debugging. A third party logrus logger has been used https://github.com
/Sirupsen/logrus

4) For unit test testify library has been used https://github.com/stretchr/testify.

### Install and Build
Requires Golang installed. Please follow the instruction from here https://golang.org/doc/install

This library is developed with go version 1.14.4

Download the library from https://github.com/gouthams/go-concurrency-util

Need access from github to resolve dependencies. Might need to run "go get" to
 resolve the package dependencies

To run a sample example, do the following
```shell script
cd solutionCheck
go run solution.go
Enter a phone number to check if it is easy dial
```

To execute unit tests, do the following
```shell script
cd phoneUtils
go test
```

To execute the unit test with the coverage profile, do the following
```shell script
cd phoneUtils
go test -coverprofile cp.out
go tool cover -html=cp.out
```

### Unit test
For assertion in unit test, this library is used https://github.com/stretchr/testify. 
This go module dependency should be resolved during the build time.  

### Future Consideration:
   1) Multiple keyboard layouts can be supported by just generating a keymap for each given keyboard layout. 
   