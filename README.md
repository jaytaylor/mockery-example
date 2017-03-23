# Mockery-example

Example case for [mockery issue #128](https://github.com/vektra/mockery/issues/128) filed with the golang tool "[mockery](https://github.com/vektra/mockery)".

## Get it

```bash
go get github.com/jaytaylor/mockery-example
```

## Run it

```bash
go run main.go
```

### Sample output:

```bash
$ go run main.go
common prefix: {
  Prefix: "2017-01-01"
}
content: {
  Key: "foo-object"
}
```

## Run tests

```bash
go test -v ./...
```

### Sample output:

```bash
$ go test -v ./...
=== RUN   TestS3Mock
--- PASS: TestS3Mock (0.00s)
    main_test.go:41: common prefix: {
          Prefix: "2017-01-01"
        }
PASS
ok      github.com/jaytaylor/mockery-example    0.222s
?       github.com/jaytaylor/mockery-example/mocks  [no test files]
```

