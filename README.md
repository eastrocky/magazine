# Magazine

Define arbitrary structures in Go and eject them to Yaml files.

📄 *main.go*
```go
magazine.Eject("config.yml", Config{
	Database{
		Username: "test-user",
	},
})
```

📄 *config.yml*
```yaml
database:
    username: test-user
    password: ""

```

The password string initializes to its default value since we did not initialize one. Later, we will demonstrate how Magazine can inject secrets from environment variables so that credentials are not stored in source code.

Stored in our source code instead are self-describing serializations of structures ejected from Magazine. Eject multiple structs to create profiles.

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

Quickly load empty structures with contents of ejected Magazine files. 

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

Magazine allows for the injection of secrets at load time through environment variables. First, initial values are provided inside a struct. Next, values at relative key paths in the Yaml object will be mapped to the `config` struct. Finally, Magazine searches for environment variables with matching key paths (all caps; underscore delimited;) and loads those into the `config` struct when set.

This table demonstrates the various locations, key paths, and precedence of our Config's "*database password*".

| Location    | Precedence | Relative Key Path          |
| ----------- | ---------- | -------------------------- |
| Go          | 3          | `config.Database.Password` |
| Yaml        | 2          | `database.password`        |
| Environment | 1          | `DATABASE_PASSWORD`        |

Setting the password via an environment variable called `DATABASE_PASSWORD` allows us to specify the password at runtime and outside source code. Environment variables will take the highest precedence if set.

Magazine only attempts to resolve variables when loading. Changing environment variables will not reflect in previously loaded structs until they have been loaded again.

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

$ DATABASE_PASSWORD=hunter2 go run main.go
Password: hunter2

```