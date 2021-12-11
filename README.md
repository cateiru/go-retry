# Go Retry Test

[![Go Reference](https://pkg.go.dev/badge/github.com/cateiru/go-retry.svg)](https://pkg.go.dev/github.com/cateiru/go-retry)

```text
go get -u github.com/cateiru/go-retry
```

## Example

```go
    import (
        "testing"

        "github.com/cateiru/go-retry"
    )

    func TestExample(t *testing.T) {
        Retry(t, func() bool {
            result := Example()
            if result == "ok" {
                return true
            }
            return false
        }, "Status message")
    }
```

## LICENSE

[MIT](./LICENSE)
