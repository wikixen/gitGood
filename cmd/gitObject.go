package main

type GitObject struct{}

func (g *GitObject) Init() {

}

func (g *GitObject) Serialize(repo *Repository) {}

func (g *GitObject) Deserialize(data []byte) {}

func objectRead(repo *Repository, sha string) {

}
