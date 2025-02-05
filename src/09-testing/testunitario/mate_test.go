package testunitario

import "testing"

func TestSuma(t *testing.T){
	total := Suma(5, 5)

	if total !=10{
		t.Errorf("Suma incorrecta, tiene %d. Se esperaba %d", total, 10)
	}

}