# Go time and its two clocks

> `https://dev.to/jingshao_chen_02a2352f476/go-time-and-its-two-clocks-3p0p`

## Time lapse

To calculate the time lapse in Go, you can use

```go
start := time.Now()
// long time consuming task
duration := time.Since(start)
```

Complete sample run:

```sh
go mod init timelapse

touch main.go

go run main.go 
```
