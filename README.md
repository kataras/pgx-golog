# pgx-golog

[![build status](https://img.shields.io/github/actions/workflow/status/kataras/pgx-golog/ci.yml?style=for-the-badge)](https://github.com/kataras/pgx-golog/actions) [![report card](https://img.shields.io/badge/report%20card-a%2B-ff3333.svg?style=for-the-badge)](https://goreportcard.com/report/github.com/kataras/pgx-golog) [![godocs](https://img.shields.io/badge/go-%20docs-488AC7.svg?style=for-the-badge)](https://pkg.go.dev/github.com/kataras/pgx-golog/)

Free [golog](https://github.com/kataras/golog) and [pgx](https://github.com/jackc/pgx) logging integration.

```go
import "github.com/jackc/pgx/v5/tracelog"
import pgxgolog "github.com/kataras/pgx-golog"

// [...]

logger := pgxgolog.NewLogger(yourGologLoggerInstanceHere)
tracer := &tracelog.TraceLog{
    Logger:   logger,
    LogLevel: tracelog.LogLevelTrace,
}
```

## 📖 Learning pgx-golog

### Installation

The only requirement is the [Go Programming Language](https://go.dev/dl/).

#### Create a new project

```sh
$ mkdir myapp
$ cd myapp
$ go mod init myapp
$ go get github.com/kataras/pgx-golog
```

<details><summary>Install on existing project</summary>

```sh
$ cd myapp
$ go get github.com/kataras/pgx-golog
```

**Run**

```sh
$ go mod tidy
$ go run .
```

</details>

<br/>

## 📝 License

This project is licensed under the [MIT License](LICENSE).
