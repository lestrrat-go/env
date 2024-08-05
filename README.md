# env

Utilities for accessing environment variables

# SYNOPSIS

```go
import "github.com/lestrrast-go/env"

env.SetPrefix("MYAPP")

value, ok := env.Lookup("FOO") // looks up MYAPP_FOO

env.Value("FOO", "defaultValue") // return defaultValue if MYAPP_FOO is not set

env.SetValue("FOO", "newValue") // set MYAPP_FOO
```