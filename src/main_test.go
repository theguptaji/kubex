package main

import "testing"

func Testkubex(t *testing.T) {
	str := "Implement me Senpai!!!"

	got := kubex()
	if got != str {
		t.Errorf("kubex() = %s; want %s", got, str)
	}
}
