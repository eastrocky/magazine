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
	Region string
	Access struct {
		Key struct {
			ID string
		}
	}
	Secret struct {
		Access struct {
			Key string
		}
	}
	Session struct {
		Token string
	}
}

func main() {
	c := &Config{}
	magazine.Load("config.yml", c)

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(c.AWS.Region),
		Credentials: credentials.NewStaticCredentials(c.AWS.Access.Key.ID, c.AWS.Secret.Access.Key, c.AWS.Session.Token),
	})

	// do something...
	_, _ = sess, err
}
