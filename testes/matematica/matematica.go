package matematica

import (
	"fmt"
	"strconv"
)

// Media é responsável por calcular o que vc já sabe.... :D
func Media(numeros ...float64) float64 {
	ttl := 0.0
	for _, num := range numeros {
		ttl += num
	}
	media := ttl / float64(len(numeros))
	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	return mediaArredondada
}
