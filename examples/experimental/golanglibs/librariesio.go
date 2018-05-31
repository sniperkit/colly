package main

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/fatih/structs"
	"github.com/hackebrot/go-librariesio/librariesio"
)

var (
	xLibrariesIO   *librariesio.Client
	requestTimeout = time.Duration(time.Second * 10)
	errEmptyToken  = errors.New("Empty libraries.io token provided, cannot instantiate api client...")
)

func newClientLibrariesIO(token string) (*librariesio.Client, error) {
	if token == "" {
		return nil, errEmptyToken
	}
	client := librariesio.NewClient(strings.TrimSpace(token))
	return client, nil
}

func getProject(owner, name string) (map[string]interface{}, error) {

	// Create a new context (with a timeout if you want)
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	// Request information about a project using the client
	project, _, err := xLibrariesIO.Project(ctx, owner, name)

	if err != nil {
		log.Errorln("error: ", err.Error())
		return nil, err
	}

	// All structs for API resources use pointer values.
	// If you expect fields to not be returned by the API
	// make sure to check for nil values before dereferencing.
	log.Printf("name: %v\n", *project.Name)
	log.Printf("version: %v\n", *project.LatestReleaseNumber)
	log.Printf("language: %v\n", *project.Language)

	return structs.Map(project), nil
}
