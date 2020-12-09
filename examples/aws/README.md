# AWS Sessions
In this example, we want to define a structure that can hold values used to start a new SDK Session. Some of those variables are well documented like `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, and `AWS_SESSION_TOKEN`. Magazine normally expects the underscores to resemble a nested structure. Instead, we provide [field level tags](https://godoc.org/gopkg.in/yaml.v3#Marshal) to associate names with underscores which keeps our structure flat.

```go
type Config struct {
	AWS
}

type AWS struct {
	Region          string
	AccessKeyID     string `yaml:"ACCESS_KEY_ID"`
	SecretAccessKey string `yaml:"SECRET_ACCESS_KEY"`
	SessionToken    string `yaml:"SESSION_TOKEN"`
}
```

`AccessKeyID` is a string that we named `ACCESS_KEY_ID` inside our Magazine. Now there won't be any confusion about underscores. Since these fields are nested inside a structure called `AWS` the full key path to this field is `AWS_ACCESS_KEY_ID`.

```go
func main() {
	c := &Config{}
	magazine.Load("config.yml", c)

	session.NewSession(&aws.Config{
		Region:      aws.String(c.AWS.Region),
		Credentials: credentials.NewStaticCredentials(c.AWS.AccessKeyID, c.AWS.SecretAccessKey, c.AWS.SessionToken),
	})
}
```

## Default Region
We can set a default region when we "Eject" our Magazine.

When this `TestEjectConfig` test executes, it will write a Magazine with `us-west-2` as the default region. This region can be overridden at load time by setting the environment variable `AWS_REGION`.

```go
func TestEjectConfig(t *testing.T) {
	magazine.Eject("config.yml", Config{
		AWS: AWS{
			Region: "us-west-2", // default region
		},
	})
}
```

```yaml
aws:
    region: us-west-2
    ACCESS_KEY_ID: ""
    SECRET_ACCESS_KEY: ""
    SESSION_TOKEN: ""

```
