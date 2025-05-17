package main

import (
	"log"

	"github.com/diother/hintermann-golang/internal/builder"
	"github.com/diother/hintermann-golang/internal/fs"
)

func main() {
	if err := fs.CleanDist(); err != nil {
		log.Fatal(err)
	}
	tmpl, err := builder.LoadTemplates()
	if err != nil {
		log.Fatal(err)
	}
	projectMetaList, err := builder.LoadProjectMetaList()
	if err != nil {
		log.Fatal(err)
	}
	if err := builder.RenderProjects(tmpl); err != nil {
		log.Fatal(err)
	}
	staticPageList, err := builder.LoadStaticPageList()
	if err != nil {
		log.Fatal(err)
	}
	if err := builder.RenderStaticPages(projectMetaList, staticPageList, tmpl); err != nil {
		log.Fatal(err)
	}
	if err := builder.CopyGlobalAssets(); err != nil {
		log.Fatal(err)
	}
	if err := builder.RenderSeoFiles(projectMetaList, staticPageList); err != nil {
		log.Fatal(err)
	}
}
