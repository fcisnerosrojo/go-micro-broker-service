package main

import (
	"testing"
)

func TestReadJSON(t *testing.T) {
	// TODO
	app := &Config{}
	_ = app.readJSON(nil, nil, nil)
}

func TestWriteJSON(t *testing.T) {
	// TODO
	app := &Config{}
	_ = app.writeJSON(nil, 0, nil)
}

func TestErrorJSON(t *testing.T) {
	// TODO
	app := &Config{}
	_ = app.errorJSON(nil, nil)
}
