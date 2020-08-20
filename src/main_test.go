package main

import (
	"os/exec"
	"testing"
)

func TestLoadProject(t *testing.T) {
	expected := "Loading project: hoge\n"
	out, err := exec.Command("go", "run", "main.go", "load", "hoge").Output()
	if err != nil {
		t.Errorf("error while running command: %s", err)
	}
	if string(out) != expected {
		t.Errorf("expected %s ; getting %s", expected, string(out))
	}
	return
}

func TestUnloadProject(t *testing.T) {
	expected := "Unloading project: hoge\n"
	out, err := exec.Command("go", "run", "main.go", "unload", "hoge").Output()
	if err != nil {
		t.Errorf("error while running command: %s", err)
	}
	if string(out) != expected {
		t.Errorf("expected %s ; getting %s", expected, string(out))
	}
	return
}

func TestBuildProject(t *testing.T) {
	expected := "Building project: hoge\n"
	out, err := exec.Command("go", "run", "main.go", "build", "hoge").Output()
	if err != nil {
		t.Errorf("error while running command: %s", err)
	}
	if string(out) != expected {
		t.Errorf("expected %s ; getting %s", expected, string(out))
	}
	return
}

func TestShowProject(t *testing.T) {
	expected := "Showing project: hoge\n"
	out, err := exec.Command("go", "run", "main.go", "show", "hoge").Output()
	if err != nil {
		t.Errorf("error while running command: %s", err)
	}
	if string(out) != expected {
		t.Errorf("expected %s ; getting %s", expected, string(out))
	}
	return
}