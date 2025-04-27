package main

import (
	"errors"
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"path/filepath"
)

const CreateErr = "error creating directory"

type Repository struct {
	worktree string
	gitDir   string
	config   *ini.File
}

func (r *Repository) Init(path string, force ...bool) error {
	if force == nil {
		force = []bool{false}
	}

	r.worktree = path
	r.gitDir = filepath.Join(path, ".git")

	dirName, _ := os.Stat(r.gitDir)
	if !force[0] || !dirName.IsDir() {
		return errors.New(fmt.Sprintf("%s isn't a git directory", dirName))
	}

	file, err := os.Create(filepath.Join(r.gitDir, "config.ini"))
	if err != nil {
		return errors.New("error creating config file")
	}
	defer file.Close()

	r.config, err = ini.Load(filepath.Join(r.gitDir, "config.ini"))
	if err != nil {
		return errors.New("error reading config file")
	}

	if !force[0] {
		vers := 0 // Need to change this
		if vers != 0 {
			return errors.New(fmt.Sprintf("unsupported repositoryformatversion: %d", vers))
		}
	}

	return nil
}

func CreateRepo(path string) error {
	repo := Repository{}
	err := repo.Init(path)
	if err != nil {
		log.Fatalln(err)
	}

	err = os.MkdirAll(repo.worktree, os.ModePerm)
	if err != nil {
		return err
	}
	d, err := os.Open(repo.worktree)
	defer d.Close()
	filenames, _ := d.Readdirnames(1)
	if len(filenames) > 0 {
		return errors.New("directory isn't empty")
	}

	err = os.MkdirAll(filepath.Join(repo.gitDir, "branches"), os.ModePerm)
	if err != nil {
		return errors.New(CreateErr)
	}
	err = os.MkdirAll(filepath.Join(repo.gitDir, "objects"), os.ModePerm)
	if err != nil {
		return errors.New(CreateErr)
	}
	err = os.MkdirAll(filepath.Join(repo.gitDir, "refs", "tags"), os.ModePerm)
	if err != nil {
		return errors.New(CreateErr)
	}
	err = os.MkdirAll(filepath.Join(repo.gitDir, "refs", "tags"), os.ModePerm)
	if err != nil {
		return errors.New(CreateErr)
	}
	return nil
}

func createDir(r Repository, path ...string) error {
	err := os.MkdirAll(filepath.Join(r.gitDir, path...), os.ModePerm)
	if err != nil {
		return errors.New(CreateErr)
	}
	return nil
}
