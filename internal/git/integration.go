package git

import (
	"fmt"
	"path/filepath"
	"versionctrls-desktop/internal/store"

	"github.com/go-git/go-git/v5"
)

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

func (rm *RepositoryManager) Exists(path string) bool {
	_, ok := rm.repositories[path]
	return ok
}

func (rm *RepositoryManager) LoadFromDisk(store *store.Store) error {
	for _, path := range store.GetRepoPaths() {
		err := rm.Add(path)
		if err != nil {
			return err
		}
	}

	return nil
}

// AddRepository adds a repository to the manager
func (rm *RepositoryManager) Add(path string) error {
	repo := NewRepository(path)
	repo, err := repo.Open()
	if err != nil {
		fmt.Println("Error opening repository:", err)
		return fmt.Errorf("error opening repository. might not be a valid repository")
	}
	rm.repositories[path] = repo
	return nil
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

func (rm *RepositoryManager) GetAll() map[string]*Repository {
	return rm.repositories
}

type RepositoryInfo struct {
	Path string
	Name string
}

func (rm *RepositoryManager) ListPaths() []RepositoryInfo {
	infos := make([]RepositoryInfo, 0, len(rm.repositories))
	for k := range rm.repositories {
		infos = append(infos, RepositoryInfo{
			Path: k,
			Name: filepath.Base(k),
		})
	}
	return infos
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

func (r *Repository) Open() (*Repository, error) {
	var err error
	r.Repo, err = git.PlainOpen(r.Path)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (r *Repository) HasIntegration() bool {
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

func (r *Repository) IsRepository() bool {
	_, err := git.PlainOpen(r.Path)
	return err == nil
}
