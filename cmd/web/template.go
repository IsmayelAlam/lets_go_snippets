package main

import "lets_go.snippetbox.ismayel/internal/models"

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
