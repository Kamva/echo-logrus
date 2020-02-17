### echo-logrus

Middleware elogrus is a [logrus](https://github.com/sirupsen/logrus) logger support for [echo](https://github.com/labstack/echo).

`v4.0` tag supports v4.

#### Install

```sh
go get -u github.com/Kamva/elogrus
```

#### Usage

import package

```go
"github.com/Kamva/elogrus"
```

define new logrus

```go
elogrus.Logger = logrus.New()
e.Logger = elogrus.GetEchoLogger()
e.Use(elogrus.Hook())
```


#### TODO: 
- [ ] Write tests
- [ ] CI
- [ ] Set badges on README