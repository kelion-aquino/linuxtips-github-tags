package main

import (
	"fmt"

	"github.com/kelion-aquino/linuxtips-github-tags/git"
)

func main() {
	fmt.Println("#VAAAAAAAAAAAAAI")

	tag := "v1.26.2"
	donoDoRepositorio := "kubernetes"
	repo := "kubernetes"

	// camelCase -> public
	// PascalCase -> private
	git.BuscaGitTag(tag, donoDoRepositorio, repo)
}
