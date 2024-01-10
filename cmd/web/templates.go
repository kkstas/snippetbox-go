package main

import "github.com/kkstas/snippetbox-go/internal/models"

type templateData struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
}
