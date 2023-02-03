# Toolkit

## How to run

```bash
# frontend web server
fresh # w/ hot reload
go run main.go # w/o hot reload

# cmd line (maybe depractated)
go install
toolkit --help
```

## Testing

```bash
go test -v ./...
```

GRC can be used for better test output. See [here](https://stackoverflow.com/questions/27242652/colorizing-golang-test-run-output). Then run:

```bash
grc go test -v ./...
```

Or [richgo](https://github.com/kyoh86/richgo) can be used similarly

```bash
richgo test -v ./...
```
