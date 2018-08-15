package matematica

import "testing"

const erroPadrao = "Valor esperado %v, mas o resultado foi %v"

func TestMedia(t *testing.T) {
	t.Parallel()
	valorEsperado := 6.45
	valor := Media(7.3, 9.2, 3.4, 5.9)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}
