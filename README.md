# wrand - Weighted Random for Go

[![Test](https://github.com/tebeka/wrand/actions/workflows/go.yml/badge.svg)](https://github.com/tebeka/wrand/actions/workflows/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/tebeka/wrand.svg)](https://pkg.go.dev/github.com/tebeka/wrand)


## Example

```go
weights := map[string]int{
    "A": 20,
    "B": 30,
    "C": 50,
}

rw, err := wrand.New(weights, nil)
if err != nil {
    fmt.Printf("error: %s", err)
    return
}

for i := 0; i < 5; i++ {
    fmt.Println(rw.Rand())
}
```
