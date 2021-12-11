# Go Retry Test

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
