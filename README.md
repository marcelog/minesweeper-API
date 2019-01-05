# minesweeper-API

## Build

Install [golint](https://github.com/golang/lint) like:
```bash
go get -u golang.org/x/lint/golint
```

Then just:

```bash
make
```

It will run tools like [vet](https://golang.org/cmd/vet/), and
[golint](https://github.com/golang/lint), compile, run tests, format the code,
and open the coverage report.

## Running
After building the server (see above), just run either of:

* build/minesweeper-API-linux
* build/minesweeper-API-darwin

## API Docs
See [spec.yml](https://github.com/marcelog/minesweeper-API/blob/master/spec.yml)
which is in [Swagger](https://swagger.io) format.