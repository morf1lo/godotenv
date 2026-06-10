# godotenv
Go package that helps load environment variables from .env files

```bash
go get github.com/morf1lo/godotenv
```
## Usage
```go
import "github.com/morf1lo/godotenv"

func main() {
    godotenv.Load() // loads '.env' file by default
}
```

### if you have several files:
```go
import "github.com/morf1lo/godotenv"

func main() {
    godotenv.Load("1.env", "2.env") // loads '1.env' and '2.env' only
}
```

### or you can load a single `.env` file by just importing `autoload` package:
```go
import (
    "fmt"
    "os"

    _ "github.com/morf1lo/godotenv/autoload"
)

func main() {
    fmt.Println(os.Getenv("FOO")) // BAR
}
```
