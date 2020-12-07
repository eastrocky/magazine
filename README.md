# Magazine
Magazine is a utility for serializing and deserializing native structures to and from YAML with field level injection.

## Eject
Creating a new Magazine is as easy as passing your structure to Eject.

📄 *main.go*

```go
type Config struct {
	Username string
	Password string
}

magazine.Eject("config.yml", Config{
	Database{
		Username: "test-user",
	},
})
```

Magazine will serialize the shape and values stored in your structures to a YAML representation. In this example, notice how the value at the key path `database.password` serializes as the [zero value](https://tour.golang.org/basics/12) for a `string` because we did not initialize the value ourselves:

📄 *config.yml*

```yaml
database:
    username: test-user
    password: ""

```

Eject multiple structures to act as profiles for particular configurations or environments.

📄 *main.go*

```go
magazine.Eject("development.yml", Config{
	Database{
		Username: "development-user",
	},
	Memory{
		GB: 1,
	},
	CPU{
		Cores: 4,
	},
})

magazine.Eject("production.yml", Config{
	Database{
		Username: "production-user",
	},
	Memory{
		GB: 128,
	},
	CPU{
		Cores: 64,
	},
})
```

📄 *development.yml*

```yaml
database:
    username: development-user
    password: ""
memory:
    gb: 1
cpu:
    cores: 4

```

📄 *production.yml*

```yaml
database:
    username: production-user
    password: ""
memory:
    gb: 128
cpu:
    cores: 64

```

Later we will see how we can inject sensitive values like passwords using environment variables.

## Load
Magazine can bind matching values from YAML files into pre-initialized structures. Those values become accessible to the rest of your code from the structure itself.

📄 *main.go*
```go
config := &Config{}
magazine.Load("production.yml", config)
fmt.Println(config.CPU.Cores)
```

📄 *standard out*
```shell
$ go run main.go
64
```

The ejected files can serve as documentation and presets.

## Inject

Magazine allows you to apply overrides to individual fields using environment variables. This enables you to provide the value only when needed instead of storing it in source code or YAML file.

Values from the environment are resolved by the variable's key path. If the environment contains a variable named after a matching YAML key path, that value is used instead. The environment variable must also use the common naming convention of capital letters delimited by underscores.

This table demonstrates the various locations, key paths, and load precedence of our Config's "*database password*".

| Location    | Precedence | Relative Key Path          |
| ----------- | ---------- | -------------------------- |
| Go          | 3          | `config.Database.Password` |
| Yaml        | 2          | `database.password`        |
| Environment | 1          | `DATABASE_PASSWORD`        |

📄 *main.go*
```go
config := &Config{}
magazine.Load("production.yml", config)
fmt.Println("Password:", config.Database.Password)
```

📄 *standard out*
```shell
$ go run main.go
Password: 
```

Before we set the password using the environment, it is empty. Using an environment variable called `DATABASE_PASSWORD` allows us to inject the password into our structure at load time:

```shell
$ DATABASE_PASSWORD=hunter2 go run main.go
Password: hunter2
```
