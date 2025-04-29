# Ejercicio 5.1.6
Funcion DetectorPar(numero: entero): boooleano Es
	Si (numero < 2) Entonces
		Si (numero = 0) Entonces
			DetectorPar:= verdadero
		Sino
			DetectorPar:= falso
		FinSi
	Sino
		DetectorPar:= DetectorPar(numero - 2)
	FinSi
FinFuncion
