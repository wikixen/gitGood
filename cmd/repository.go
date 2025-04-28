package main

import (
	"errors"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"path/filepath"
)

type Repository struct {
	worktree string
	gitDir   string
	config   *ini.File
}

// Init instantiates a Repository
func (r *Repository) Init(path string, force bool) error {
	r.worktree = path
	r.gitDir = filepath.Join(path, ".git")

	dirName, _ := os.Stat(r.gitDir)
	if !force || !dirName.IsDir() {
		return errors.New("git dir is not a directory")
	}

	r.config = ini.Empty()
	_, err := os.Create("config.ini")
	if err != nil {
		return err
	}

	if _, err = os.Stat("config.ini"); err == nil {
		r.config, err = ini.Load("config")
		if err != nil {
			return err
		}
	} else {
		return errors.New("config.ini does not exist")
	}

	if !force {
		version, err := r.config.Section("core").Key("repositoryformatversion").Int()
		if err != nil {
			return err
		}
		if version > 0 {
			return errors.New("unsupported git version")
		}
	}
	return nil
}

// createRepo creates a Repository in the directory
func createRepo(path string) error {
	repo := new(Repository)
	err := repo.Init(path, false)
	if err != nil {
		return err
	}

	dirName, err := os.Stat(repo.worktree)
	if err == nil {
		if !dirName.IsDir() {
			return errors.New("path is not a directory")
		}
		dirName, _ := os.Stat(repo.gitDir)
		if dirName.IsDir() && dirName.Size() > 0 {
			return errors.New("repository is not empty")
		}
	} else {
		err := os.MkdirAll(repo.worktree, 0755)
		if err != nil {
			return err
		}
	}

	makeDir(filepath.Join(repo.gitDir, "branches"))
	makeDir(filepath.Join(repo.gitDir, "objects"))
	makeDir(filepath.Join(repo.gitDir, "refs", "tags"))
	makeDir(filepath.Join(repo.gitDir, "refs", "heads"))

	err = os.WriteFile(filepath.Join(repo.gitDir, "description"), []byte("Unnamed repository; edit this file 'description' to name the repository.\\n"), 0755)
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath.Join(repo.gitDir, "HEAD"), []byte("ref: refs/heads/master\\n"), 0755)
	if err != nil {
		return err
	}

	_, err = os.Create(filepath.Join(repo.gitDir, "config"))
	if err != nil {
		return err
	}
	err = fillConfig()
	if err != nil {
		return err
	}

	return nil
}

func makeDir(path string) {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		log.Fatal(err)
	}
}

// fillConfig fills the .git/config with initial values
func fillConfig() error {
	conf, err := ini.Load("config")
	if err != nil {
		return err
	}
	core, err := conf.NewSection("core")
	if err != nil {
		return err
	}

	_, err = core.NewKey("repositoryformatversion", "0")
	if err != nil {
		return err
	}
	_, err = core.NewKey("filemode", "false")
	if err != nil {
		return err
	}
	_, err = core.NewKey("bare", "false")
	if err != nil {
		return err
	}

	return nil
}
