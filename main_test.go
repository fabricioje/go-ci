package main

import "testing"

func TestCalcular(t *testing.T) {

    if Calcular(5, 5) != 10 {
		t. Error("O resultado deveria ser 10")
	}

}