package git

import "github.com/go-git/go-git/v5"

// RepositoryManager handles multiple git repositories
type RepositoryManager struct {
	repositories map[string]*Repository
}

// NewRepositoryManager creates a new repository manager
func NewRepositoryManager() *RepositoryManager {
	return &RepositoryManager{
		repositories: make(map[string]*Repository),
	}
}

// AddRepository adds a repository to the manager
func (rm *RepositoryManager) Add(path string, repo *Repository) {
	rm.repositories[path] = repo
}

// DeleteRepository removes a repository from the manager
func (rm *RepositoryManager) Delete(path string) {
	delete(rm.repositories, path)
}

// GetRepository retrieves a repository by path
func (rm *RepositoryManager) Get(path string) (*Repository, bool) {
	repo, exists := rm.repositories[path]
	return repo, exists
}

type Repository struct {
	Path          string
	submodulePath string
	submoduleName string
	Repo          *git.Repository
}

func NewRepository(path string) *Repository {
	return &Repository{Path: path, submodulePath: "versionctrls-integration",
		submoduleName: "versionctrls-integration"}
}

func (r *Repository) Clone(url string, isBare bool, o *git.CloneOptions) (*Repository, error) {
	var err error
	r.Repo, err = git.PlainClone(url, isBare, o)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (r *Repository) Open(path string) (*Repository, error) {
	var err error
	r.Repo, err = git.PlainOpen(path)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (r *Repository) HasIntegration(path string) bool {
	// Get the worktree
	wt, err := r.Repo.Worktree()
	if err != nil {
		return false
	}

	// Get all submodules
	submodules, err := wt.Submodules()
	if err != nil {
		return false
	}

	// Look for the specific submodule
	for _, sub := range submodules {
		if sub.Config().Name == "versionctrls-integration" {
			return true
		}
	}

	return false
}
