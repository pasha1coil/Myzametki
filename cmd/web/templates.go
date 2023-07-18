package main

import (
	"html/template"
	"path/filepath"
	"zametki/pkg/models"
)

type templateData struct {
	Zametki   *models.Zametki
	Zametkiss *[]models.Zametki
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {

	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	// Перебираем файл шаблона от каждой страницы.
	for _, page := range pages {

		name := filepath.Base(page)

		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}

	return cache, nil
}
