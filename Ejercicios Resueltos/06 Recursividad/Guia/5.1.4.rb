# Ejercicio 5.1.4
Funcion DetectorDigitos(numero: entero): boooleano Es
	Si ((numero MOD 2) <> 0) Entonces # Caso Base
		DetectorDigitos:= falso
	Sino # Caso Recursivo
		Si (numero < 10) Entonces
			DetectorDigitos:= verdadero
		Sino
			DetectorDigitos:= DetectorDigitos(numero DIV 10)
		FinSi
	FinSi
FinFuncion
