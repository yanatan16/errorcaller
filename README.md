errorcaller [![Build Status][1]][2]
===========

An error wrapper to attach caller (file:line) information to errors

## Install

```
go get github.com/yanatan16/errorcaller
```

## Use

```golang
import (
  "fmt"
  "github.com/yanatan16/errorcaller"
)

func DoSomethingErrorProne() error {
  if err := Oops(); err != nil {
    return errorcaller.Err(err)
  }

  if ok := OhNoes(); !ok {
    return errorcaller.New("Oh noes!")
  }
}

func CallDoSomethingErrorProne() {
  err := DoSomethingErrorProne()
  fmt.Println(err)
}
```

Will output something like this:

```
Oh noes! (on /Users/you/gopath/src/github.com/you/error-prone/errorprone.go:10)
```


Note that performing `errorcaller.Err(errorcaller.Err(err))` is the same as doing it once because it won't overwrite a caller in the first place.

## License

MIT License found in LICENSE file.


[1]: https://travis-ci.org/yanatan16/errorcaller.png?branch=master
[2]: http://travis-ci.org/yanatan16/errorcaller
