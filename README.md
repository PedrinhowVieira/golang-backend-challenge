# League Backend Challenge

This challenge consists in a basic web server written in GoLang. It accepts an input and execute one of the five different type of requests.

## The input

It must be a csv file containing a square matrix of integers.
> In mathematics, a square matrix is a matrix with the same number of rows and columns.
```
// input example
1,2,3
4,5,6
7,8,9
```
## The requests

#### Echo
Returns the matrix as a string in matrix format. The request is accessed with:
```
curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

// output example
1,2,3
4,5,6
7,8,9
``` 

#### Invert
Return the matrix as a string in matrix format where the columns and rows are inverted. The request is accessed with:
```
curl -F 'file=@/path/matrix.csv' "localhost:8080/invert"

// output example
1,4,7
2,5,8
3,6,9
``` 

#### Flatten
Return the matrix as a 1 line string, with values separated by commas. The request is accessed with:
```
curl -F 'file=@/path/matrix.csv' "localhost:8080/flatten"

// output example
1,2,3,4,5,6,7,8,9
``` 

#### Sum
Return the sum of the integers in the matrix. The request is accessed with:
```
curl -F 'file=@/path/matrix.csv' "localhost:8080/sum"

// output example
45
``` 

#### Multiply
Return the product of the integers in the matrix. The request is accessed with:
```
curl -F 'file=@/path/matrix.csv' "localhost:8080/multiply"

// output example
362880
``` 

## The web server

The wer server is started in the root of the project where the `main.go` file is located, with:
```
go run .
```

## Tests
To run tests on the operation package:
```
go test operation
```
#### Benchmark
To run benchmark tests:
```
go test operation -bench=.
```
#### Test Coverage
To run the test coverage
```
go test operation -cover
```

## Authors
Pedro de Moraes Vieira - pedro.dmvieira@gmail.com