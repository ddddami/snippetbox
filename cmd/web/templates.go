package main

import "snippetbox.damilola.dev/internal/models"

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
