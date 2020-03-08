# League Backend Challenge

This challenge consists in a basic web server written in GoLang. It accepts an input and execute one of the five 
different types of request.

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

The wer server is started at the root of the project where the `main.go` file is located, with:
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

## Concurrency
Two different algorithms, one with concurrency (`'operationName'Concurrency`) and one without (`'operationName'Default`), were develop to improve the performance of the higher 
cost operations, like Sum, Multiply and Invert. The threshold that chooses which one of these two algorithms will be 
used was based on a benchmark test.
> Sum and Multiply threshold: 2500x2500 matrix; Invert threshold: 350x350 matrix.

| Functions                    | Time/Operation |
|------------------------------|----------------|
| BenchmarkInvertDefault       | 67101319 ns/op |
| BenchmarkInvertConcurrency   | 52792818 ns/op |
| BenchmarkMultiplyDefault     | 57223414 ns/op |
| BenchmarkMultiplyConcurrency | 56262858 ns/op |
| BenchmarkSumDefault          | 56098063 ns/op |
| BenchmarkSumConcurrency      | 55602143 ns/op |

In fact It was developed a third algorithm (`'operationName'ConversionIn`) that instead of converting all the matrix before 
executing the operation, it converts in the same loop that the operation runs. It shows improvements over the other two 
functions previously presented.
> Same thresholds were used.

| Functions                     | Time/Operation |
|-------------------------------|----------------|
| BenchmarkInvertConversionIn   | 42739369 ns/op |
| BenchmarkMultiplyConversionIn | 40981590 ns/op |
| BenchmarkSumConversionIn      | 39374232 ns/op |

Although the performance improvement of this algorithm, I could not generate a concurrency version of this. So it was kept 
out of the main operation functions so the comparison of concurrency and non-concurrency could be made.
## Authors
Pedro de Moraes Vieira - pedro.dmvieira@gmail.com