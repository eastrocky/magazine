package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/eastrocky/magazine"
)

type Config struct {
	AWS
}

type AWS struct {
	Region          string
	AccessKeyID     string `yaml:"ACCESS_KEY_ID"`
	SecretAccessKey string `yaml:"SECRET_ACCESS_KEY"`
	SessionToken    string `yaml:"SESSION_TOKEN"`
}

func main() {
	c := &Config{}
	magazine.Load("config.yml", c)

	session.NewSession(&aws.Config{
		Region:      aws.String(c.AWS.Region),
		Credentials: credentials.NewStaticCredentials(c.AWS.AccessKeyID, c.AWS.SecretAccessKey, c.AWS.SessionToken),
	})
}
