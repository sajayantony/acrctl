package reg

//ListRepos List repositories of a registry
func ListRepos(registry *Registry) ([]string, error) {
	repos := make([]string, 1)
	repos[0] = "test"
	return repos, nil
}
