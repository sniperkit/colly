package service

import (
	"errors"

	"github.com/sniperkit/colly/plugins/data/aggregate/core/graph"
)

type GraphService struct {
	Storage *graph.GraphStorage
}

func (s *GraphService) Name() string {
	return "graph-service"
}

func NewGraph() (*GraphService, error) {
	storage, err := graph.NewGraphStorage()
	if err != nil {
		return nil, errors.New("Cannot load graph storage")
	}

	return &GraphService{storage}, nil
}
