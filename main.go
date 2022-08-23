package main

//go package note
//https://linguinecode.com/post/how-to-import-local-files-packages-in-golang

import (
	"github.com/KelvinWu602/Forus/core/fragment"
)

func main() {
	var happy string = "kufeilarigueliurgvliuegviusjyhvyjuijy"
	fragments := fragment.Breakdown_bytes([]byte(happy), 50)
	fragment.Print_fragments(fragments)
}
