# Command
## Start
Terminal:
```
go mod init <name>
```
Install dependency
```
go mod tidy
```
Run
```
go run .
```
## Test
```
go test
```
`-v` flag to get verbose output that lists all of the tests and their results
```
go test -v
```

## Build
```
go build
```
Discover the Go install path
```
go list -f '{{.Target}}'
```
