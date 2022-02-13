# Go Commands:
**List all environment variables**
>go env  

**To run go main function including other functions with in main package but different files.**
>go run main.go, sample.go  

**To create an executable file. In windows it will create main.exe**
>go buil main.go sample.go 

**To format all the files under home directory. Three dots**
>go ../...


*GoLint*
**Install golint, it installs gofmt and go vet also**

> go get -u github.com/golang/lint/golint
> 
*reports poor coding style*
> golint ./...
> 

*gofmt: formats go code*
>gofmt ./...

*go vet: reports suspicious constructs*
>go vet ./...

*Run all test in a package say codepractice*

>cd codepractice 

> go test

*Run all the tests under a folder* 
> go test ./...

#Run documentation locally#
>godoc -http=:8080

*go commands help *
>go help test
> go help testflage
> go help tool

*Code test coverage on current package *
>go test -cover 

**To dump the coverge output to file**
> go test -coverprofile c.out 

**To view the html output of a package**
> go tool cover -html c.out
> go tool cover -help

*Remember to bet*
- Benchmark
- Example
- Test

```
BenchmarkCar(b *testing.B)
ExampleCat(t *testing.T)
TestCat(t *testing.T)
```

```
Commands: 
godoc  -http=:8080

go test
got -bench . 
go test -cover 
go test -coverprofile c.out
go tool cover -html c.out 
```
