package repos_rest_component

import (
	"sync"
)

type componentRepo struct {
	baseUrl string
	m       *sync.RWMutex
}

func NewComponentRepo(baseUrl string) *componentRepo {
	return &componentRepo{
		baseUrl: baseUrl,
		m:       &sync.RWMutex{},
	}
}
