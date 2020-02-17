### elogrus

Middleware elogrus is a [logrus](https://github.com/sirupsen/logrus) logger support for [echo](https://github.com/labstack/echo).

`v4.0` tag supports v4.

#### Install

```sh
go get -u github.com/Kamva/elogrus/v4
```

#### Usage

import package

```go
import "github.com/Kamva/elogrus/v4"
```

define new logrus

```go
e.Logger = elogrus.GetEchoLogger(logrus.New())
e.Use(elogrus.Hook())
```


#### TODO: 
- [ ] Write tests
- [ ] CI
- [ ] Set badges on README