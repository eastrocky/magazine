# magazine
Configuration loader for Go applications. Mimicks Spring Boot behaviors for loading YAML and overriding with environment variables.

## Loading a file
**config.yml**
```yml
author: eastrocky
application:
  name: magazine
  version: 1.0
```

**main.go**
```go
package main

import (
	"fmt"

    "github.com/eastrocky/magazine"
)

func main() {
	config, err := magazine.Load("config.yml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config["application.version"])
}
```

**output**
> 1.0

## Override Variables
**config.yml**
```yml
author: eastrocky
application:
  name: magazine
  version: 1.0
```

**shell**
```sh
export APPLICATION_VERSION=2.0
```

**main.go**
```go
package main

import (
	"fmt"

    "github.com/eastrocky/magazine"
)

func main() {
	config, err := magazine.Load("config.yml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config["application.version"])
}
```

**output**
> 2.0