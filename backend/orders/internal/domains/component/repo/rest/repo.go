package component_repo_rest

type ComponentRepo struct {
	baseUrl string
}

func NewComponentRepo(baseUrl string) *ComponentRepo {
	return &ComponentRepo{
		baseUrl: baseUrl,
	}
}
