package main

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
)

type Repository struct {
	worktree string
	gitdir   string
	config   *ini.File
}

func (r *Repository) Init(path string, force ...bool) {
	if force == nil {
		force = []bool{false}
	}

	r.worktree = path
	r.gitdir = path + "/.git"

	dirName, err := os.Stat(path)
	if err != nil {
		log.Panicf("Failed to find path %s", path)
	}
	if !(force[0] || dirName.IsDir()) {
		log.Panicf("%s is not a Git repository", path)
	}

	r.config, err = ini.Load("config")
	if err != nil {
		log.Panic("Failed to load config file")
	}

	cf, ok := repoFile(*r, false, "config")
	_, err = os.Stat(cf)
	if !ok && os.IsNotExist(err) {
		log.Panicf("Config file %s does not exist", cf)
	}

	if !force[0] {
		version := r.config.Section("core").Key("version").MustInt(0)
		if version != 0 {
			log.Panicf("Git version %d is not supported", version)
		}
	}
}

// repoFile creates a file with the provided path
func repoFile(repo Repository, mkdir bool, path ...string) (string, bool) {
	if repoDir(repo, mkdir, path[:len(path)-1]...) != "" {
		return repoPath(repo, path...), true
	}
	return "", false
}

// repoPath converts the strings provided in paths into a single path
func repoPath(repo Repository, paths ...string) string {
	path := repo.gitdir
	for p := range path {
		path += "/" + paths[p]
	}
	return path
}

// repoDir is the same as repoPath, but it makes a dir
func repoDir(repo Repository, mkdir bool, paths ...string) string {
	path := repoPath(repo, paths...)

	dirName, err := os.Stat(path)
	if os.IsExist(err) {
		if dirName.IsDir() {
			return path
		} else {
			log.Panicf("%s is not a directory", path)
		}
	}

	if mkdir {
		err = os.Mkdir(path, 0755)
		if err != nil {
			log.Panicf("Failed to create directory %s", path)
		}
		return path
	}
	return ""
}

func repoCreate(path string) Repository {
	repo := Repository{}
	repo.Init(path, true)

	dirName, err := os.Stat(repo.worktree)
	if os.IsExist(err) {
		if !dirName.IsDir() {
			log.Panicf("%s is not a directory", path)
		}
		_, err := os.Stat(repo.gitdir)
		_, e := os.ReadDir(repo.gitdir)
		if os.IsExist(err) && e == nil {
			log.Panicf("%s is not empty", path)
		}
	} else {
		err := os.Mkdir(repo.worktree, 0755)
		if err != nil {
			return Repository{}
		}
	}

	if repoDir(repo, true, "branches") == "" {
		log.Panicf("An error occurred while creating the repository")
	}
	if repoDir(repo, true, "objects") == "" {
		log.Panicf("An error occurred while creating the repository")
	}

	return repo
}
