package backend

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"versionctrls-desktop/internal/git"
	"versionctrls-desktop/internal/store"

	"github.com/radovskyb/watcher"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx      context.Context
	store    *store.Store
	repoMngr *git.RepositoryManager
	watchers map[string]*watcher.Watcher
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) Startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
	a.store = store.NewStore(os.Getenv("STORE_PATH"))
	a.watchers = make(map[string]*watcher.Watcher)

	err := a.store.OpenStore()
	if err != nil {
		fmt.Println("Error opening store:", err)
	}

	a.repoMngr = git.NewRepositoryManager()
	err = a.repoMngr.LoadFromDisk(a.store)
	if err != nil {
		fmt.Println("Error loading repositories from disk:", err)
	}
}

// domReady is called after front-end resources have been loaded
func (a App) DomReady(ctx context.Context) {

}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) BeforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) Shutdown(ctx context.Context) {
	// Perform your teardown here
}

func (a *App) ToggleWatcher(path string, watch bool) error {
	if watch {
		w := watcher.New()
		go func() {
			for {
				select {
				case event := <-w.Event:
					fmt.Println(event.Name())
					fmt.Println(event.Path)
					fmt.Println(event.Op)
				case err := <-w.Error:
					log.Fatalln(err)
				case <-w.Closed:
					return
				}
			}
		}()

		if err := w.AddRecursive(path); err != nil {
			runtime.LogFatal(a.ctx, err.Error())
		}

		a.watchers[path] = w
		if err := w.Start(time.Millisecond * 100); err != nil {
			delete(a.watchers, path)
			log.Fatalln(err)
		}

		return nil
	} else {
		// Check if watcher exists for this path
		w, ok := a.watchers[path]
		if ok {
			w.Close()
			delete(a.watchers, path)
		}
	}

	return nil
}

type OpenRepositoryResult struct {
	Path  string
	Error string
}

func (a *App) OpenRepository() OpenRepositoryResult {
	// Open the directory selection dialog
	path, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select a Repository",
	})

	if err != nil {
		fmt.Println("Error selecting folder:", err)
		return OpenRepositoryResult{
			Path:  "",
			Error: "Error selecting folder",
		}
	}

	if a.repoMngr.Exists(path) {
		return OpenRepositoryResult{
			Path:  path,
			Error: "repository already opened",
		}
	}

	err = a.repoMngr.Add(path)
	if err != nil {
		return OpenRepositoryResult{
			Path:  "",
			Error: err.Error(),
		}
	}

	a.store.StoreRepoPath(path)

	return OpenRepositoryResult{
		Path:  path,
		Error: "",
	}
}

func (a *App) ListRepos() []git.RepositoryInfo {
	return a.repoMngr.ListPaths()
}

func (a *App) VerifyIntegration(entry string) bool {
	repo := git.NewRepository(entry)
	repo.Open()
	return repo.HasIntegration()
}

func (a *App) RemoveRepository(path string) error {
	// Remove from store
	err := a.store.RemoveRepoPath(path)
	if err != nil {
		return err
	}

	// Remove from repository manager
	a.repoMngr.Delete(path)

	// Stop and remove any watchers
	if w, exists := a.watchers[path]; exists {
		w.Close()
		delete(a.watchers, path)
	}

	return nil
}
