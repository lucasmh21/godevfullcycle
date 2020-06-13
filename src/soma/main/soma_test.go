package main

import "testing"

func TestSoma(t *testing.T){
	resultado := soma(5,5)
	if resultado != 10{
		t.Error("Resultado incorreto!")
	}
}
