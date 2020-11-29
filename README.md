# Magazine

## Eject

Define arbitrary maps and eject them to yaml files.

**main.go**
```go
package main

import "github.com/eastrocky/magazine"

type Database struct {
	Hostname string
	Login    struct {
		Username string
		Password string
	}
	Tables []string
}

func main() {
	defaultConfig := &Database{
		Hostname: "localhost",
		Login: struct {
			Username string
			Password string
		}{
			Username: "admin",
		},
		Tables: []string{
			"users",
			"orders",
			"audit",
		},
	}

	magazine.Eject("config.yml", defaultConfig)
}

```

**config.yml**
```yml
hostname: localhost
login:
    username: admin
    password: ""
tables:
    - users
    - orders
    - audit

```

This records the shape, fields, and types that make up the ejected struct. Ejected files can serve as documentation and also be used to quickly swap configurations using the `Load` method.

## Load

Load binds previously ejected magazines back into structures.

**config.yml**
```yml
hostname: localhost
login:
    username: admin
    password: ""
tables:
    - users
    - orders
    - audit

```

**main.go**
```go
package main

import (
	"fmt"

	"github.com/eastrocky/magazine"
)

type Database struct {
	Hostname string
	Login    struct {
		Username string
		Password string
	}
	Tables []string
}

func main() {
	config := &Database{}

	magazine.Load("config.yml", config)
	fmt.Println("Username:", config.Login.Username)
}
```

**shell**
```shell
$ go run main.go
Username: admin
```

When structures are being loaded, Magazine can resolve values from the environment.

## Environment Variables

Environment variables can be used to override values loaded at particular key paths.

This can be useful for resolving secrets at runtime by discovering values set in the environment. This example loads a magazine with an empty password and sets its value using the environment instead.

**config.yml**
```yml
hostname: localhost
login:
    username: admin
    password: ""
tables:
    - users
    - orders
    - audit

```

**shell**
```shell
$ env
LOGIN_PASSWORD=hunter2
```

**main.go**
```go
package main

import (
	"fmt"

	"github.com/eastrocky/magazine"
)

type Database struct {
	Hostname string
	Login    struct {
		Username string
		Password string
	}
	Tables []string
}

func main() {
	config := &Database{}

	magazine.Load("config.yml", config)
    fmt.Println("Username:", config.Login.Username)
    fmt.Println("Password:", config.Login.Password)
}
```

**shell**
```shell
$ go run main.go
Username: admin
Password: hunter2
```

Magazine finds a variable matching the flattened key path `login.password` at `LOGIN_PASSWORD` and loads it into the structure instead.
