# mathg

A wrapper of [Golang's math package](https://pkg.go.dev/math) to support generic numbers. 

Note: Functions in this package just call the math function with the number casted to float64; this package is not well tested and may or may not function as you expect. 

## Usage
go get

```sh
go get github.com/stefafafan/mathg
```

You can use many of the math functions with float32, int64, etc.

```golang

import "github.com/stefafafan/mathg"

func main() {
    // int abs
    mathg.Abs(123)

    // float abs
    mathg.Abs(123.456)

    // ...
}
```

## Author

stefafafan ([GitHub](https://github.com/stefafafan), [Twitter](https://twitter.com/stefafafan))
