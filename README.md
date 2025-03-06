# go-must

Standard assertions.

## Dependencies
None

## Installation
```shell
go get github.com/go-stdlib/go-must
```

## Usage

The package offers standard assertions for panics.

```go
package main

import (
    "os"
    "github.com/go-stdlib/go-must"
    "strconv"
)

func main() {
    // Use `V0` or `NotErr` for calls that just return an error.
    must.V0(os.Remove("file.txt"))
    // Use `V1` for calls that return a single value + error.
    f := must.V1(os.Open("file.txt"))
    // Use `V2` for calls that return two values + error.
    rf, wf := must.V2(os.Pipe())

    // Use `Alias` for type aliases.
    v64 := must.Alias[int64](1024)

    // Use `Fn` or `Fn1` for functions that return a single value + error.
    val := must.Fn(func() (int64, error) {
        return strconv.ParseInt(os.Getenv("NUM_THREADS"), 10, 32)
    })
    
    // Use `NotZero` for nil/zero checking.
    var file *os.File
    valid := must.NotZero(file)
}
```

## License

[Apache 2.0](../LICENSE)

