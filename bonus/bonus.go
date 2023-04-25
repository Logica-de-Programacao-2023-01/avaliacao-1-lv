package bonus

//Joãozinho ganhou um kit de construção de torres. O kit é composto por várias barras de madeira, e seus comprimentos são
//conhecidos. As barras podem ser empilhadas umas sobre as outras, desde que seus comprimentos sejam iguais.
//
//Joãozinho quer construir o menor número possível de torres com as barras que tem. Você deve ajudar Joãozinho a usar as
//barras da melhor maneira possível, determinando a altura da torre mais alta e quantas torres podem ser construídas.

import "sort"

func CalculateTowers(barLengths []int) (int, int) {
	sort.Ints(barLengths)
	unicobarra := make(map[int]bool)

	for_, length := range barLengths {
		unicobarra[length] = true
	}
	contador := make(map[int]int)
	for_, length := range barLengths {
		contador[length]++,
	}
	numerotorre := len(unicobarra)
	var maxaltura int

	for_, contar := range contador {
		if contar > maxaltura {
		maxaltura = contar
	}
	}
	return maxaltura, numerotorre
}
