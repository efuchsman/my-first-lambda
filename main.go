package main

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/aws/aws-lambda-go/lambda"
	log "github.com/sirupsen/logrus"
)

type Person struct {
	name  string
	age   string
	sex   string
	state string
}

func Handler(ctx context.Context) (*Person, error) {
	personData, err := ioutil.ReadFile("data/person.json")
	if err != nil {
		log.Errorf("ERROR READING FILE: %+v", err)
		return nil, err
	}

	var person Person
	err = json.Unmarshal(personData, &person)
	if err != nil {
		log.Errorf("ERROR UNMARSHALING JSON: %+v", err)
		return nil, err
	}
	return &person, nil
}

func main() {
	lambda.Start(Handler)
}
