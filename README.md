# wrand - Weighted Random


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
