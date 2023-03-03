package git

import "fmt"

func BuscaGitTag(tag, donoDoRepositorio, repo string) {
	fmt.Println("Buscando commits pertencentes a git tag...")
	// Formatação de Print
	fmt.Printf("%s %s %s\n", tag, donoDoRepositorio, repo)
}
