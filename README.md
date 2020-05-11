# Remember to BET
• Benchmark
• Example
• Test

# Signature functions
```
BenchmarkCat(b *testing.B)
ExampleCat()
TestCat(t *testing.T)
```

# Commands
```
go test
go test -bench .
go test -cover
go test -coverprofile c.out
go tool cover -html=c.out

godoc -http=:8080
```
