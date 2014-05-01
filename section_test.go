package main

import (
	"testing"

	assert "github.com/pilu/miniassert"
)

func TestNewSection(t *testing.T) {
	var s *Section

	s = newSection(".go,.tpl, .tmpl,    .html, , , ")
	assert.Equal(t, "", s.Name)
	assert.Equal(t, []string{".go", ".tpl", ".tmpl", ".html"}, s.Extensions)
	assert.Equal(t, 0, len(s.Commands))

	s = newSection("stylesheets: .css, .less")
	assert.Equal(t, "stylesheets", s.Name)
	assert.Equal(t, []string{".css", ".less"}, s.Extensions)
	assert.Equal(t, 0, len(s.Commands))
}

func TestSection_NewCommand(t *testing.T) {
	s := newSection("go")
	assert.Equal(t, 0, len(s.Commands))
	c := s.NewCommand("build", "./build")
	assert.Equal(t, 1, len(s.Commands))
	assert.Equal(t, c, s.Commands[0])
}