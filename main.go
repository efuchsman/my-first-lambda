package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/aws/aws-lambda-go/lambda"
	log "github.com/sirupsen/logrus"
)

type Person struct {
	Name  string `json:"name"`
	Age   string `json:"age"`
	Sex   string `json:"sex"`
	State string `json:"state"`
}

func Handler(ctx context.Context) (*Person, error) {

	// Get the module path from GOMOD
	modulePath := os.Getenv("GOMOD")

	// Construct the absolute path to the 'person.json' file
	filePath := filepath.Join(modulePath, "data", "person.json")

	personData, err := ioutil.ReadFile(filePath)
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

	log.Infof("Read person data: %+v", person)
	return &person, nil
}

func main() {
	lambda.Start(Handler)
}
