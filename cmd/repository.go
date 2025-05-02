package main

import (
	"errors"
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"path/filepath"
)

// Might revisit this later & combine Init and createRepo

type Repository struct {
	worktree string
	gitDir   string
	config   *ini.File
}

// Init instantiates a Repository
func (r *Repository) Init(path string, force bool) error {
	path, err := filepath.Localize(path)
	if err != nil {
		return errors.New("error localizing path")
	}
	r.worktree = path
	r.gitDir = filepath.Join(".", path, ".git")

	dirName, _ := os.Stat(r.gitDir)
	if !(force || dirName.IsDir()) {
		return errors.New("git dir is not a directory")
	}

	err = os.MkdirAll(r.gitDir, 0755)
	if err != nil {
		return errors.New("error creating .git directory")
	}
	configPath := filepath.Join(".", r.gitDir, "config.ini")
	_, err = os.Create(configPath)
	if err != nil {
		return errors.New("error creating config.ini")
	}

	if _, err = os.Stat(configPath); err == nil {
		r.config, err = ini.Load(configPath)
		if err != nil {
			return errors.New("error loading config.ini")
		}
	} else {
		return errors.New("config.ini does not exist")
	}

	if !force {
		version, err := r.config.Section("core").Key("repositoryformatversion").Int()
		if err != nil {
			return errors.New("error reading repositoryformatversion")
		}
		if version > 0 {
			return errors.New("unsupported git version")
		}
	}
	return nil
}

// createRepo creates a Repository in the directory
func createRepo(path string) error {
	fmt.Println("Initializing git repository...")
	repo := new(Repository)
	err := repo.Init(path, true)
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

	err = fillConfig(repo.gitDir)
	if err != nil {
		return err
	}

	fmt.Println("Initialized empty Git repository in " + repo.worktree)

	return nil
}

func makeDir(path string) {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		log.Fatal(err)
	}
}

// fillConfig fills the .git/config with initial values
func fillConfig(path string) error {
	path = filepath.Join(".", path, "config.ini")
	conf, err := ini.Load(path)
	if err != nil {
		return errors.New("error loading config.ini")
	}

	core, err := conf.NewSection("core")
	if err != nil {
		return errors.New("error creating core section")
	}

	_, err = core.NewKey("repositoryformatversion", "0")
	if err != nil {
		return errors.New("error creating repositoryformatversion key")
	}

	_, err = core.NewKey("filemode", "false")
	if err != nil {
		return errors.New("error creating filemode key")
	}

	_, err = core.NewKey("bare", "false")
	if err != nil {
		return errors.New("error creating bare key")
	}

	err = conf.SaveTo(filepath.Join(path))
	if err != nil {
		return errors.New("error saving config.ini")
	}

	return nil
}
