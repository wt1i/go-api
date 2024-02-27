# go-api
[![PkgGoDev](https://pkg.go.dev/badge/github.com/wt1i/go-api)](https://pkg.go.dev/github.com/wt1i/go-api)

采用依赖注入的方式，使用 gin 框架做的 go-api，支持 swag 进行 api 展示

## run
`make run`

## api
- http://localhost:8080/api/v1/swagger/index.html`

## prometheus metrics and pprof
- http://localhost:8090/metrics
- http://localhost:8090/debug/pprof/
go tool pprof http://localhost:8090/debug/pprof/heap
go tool pprof -seconds 5 http://localhost:8090/debug/pprof/profile