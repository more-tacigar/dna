# dna - A Simple Go Framework for API Server #

イッツ、シンプル。

```
go get github.com/more-tacigar/dna
```

```go
import (
    "github.com/more-tacigar/dna"
)

func main() {
    e := dna.NewEngine()
    r := e.Router
    r.GET("/", func(c *dna.Context) {
        m := map[string]string {
            "hello": "world"
        }
        if err := dna.EncodeJson(m, c.Writer); err != nil {
            c.Status(500)
            c.Abort()
            return
        }
        c.Status(200)
    })
}
```
