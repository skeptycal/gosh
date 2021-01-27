package refresh

import (
	"context"
	"os"
	"path/filepath"

	"github.com/skeptycal/util/gofile"
)

const (
	modeNewDir  = 0o755
	modeNewFile = 0o644
)

var refreshContext context.Context = context.Background()

// RefreshRepo performs various periodic repo functions for
// repo contained in dir.
//
// The default is the current directory.
//
// The init parameter determines whether a repo should be initialized.
//
// Init True
//
// - If the directory dir does not exist, it will be created.
//
// - If the path does not contain a Git repository, one will be created.
//
// - A default project directory structure will be created.
//
// - A new GitHub repository will be created and added as a remote.
//
// - Default .gitignore and other repository files will be created.
//
// Init False
//
// If the path does not exist or is not a repository, an error is returned.
func RefreshRepo(ctx context.Context, dir string, init bool) error {

	os.Chdir(dir)
	return nil
}

func CheckRepo(ctx context.Context, dir string, init bool) (string, error) {
	if dir == "" {
		dir = gofile.PWD()
	}

	dir, err := filepath.Abs(dir)
	if err != nil {
		if !os.IsNotExist(err) {
			return "", err
        }
        if !init {
            return "", err
        }
        err = os.MkdirAll(dir, modeNewDir)
        if err != nil {
            return "", err
		}
    }

    if gorepo.

	return dir, nil
}

func Template(ctx context.Context, dir string) error {

	return nil
}
