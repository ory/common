# ory-libs/env

Adds defaults to `os.GetEnv()` and saves you 3 lines of code:

```go
import "github.com/ory/common/env"

func main() {
  port := env.Getenv("PORT", "80")
}
```

versus

```go
import "os"

func main() {
  port := os.Getenv("PORT")
  if port == "" {
    port = "80"
  }
}
```
