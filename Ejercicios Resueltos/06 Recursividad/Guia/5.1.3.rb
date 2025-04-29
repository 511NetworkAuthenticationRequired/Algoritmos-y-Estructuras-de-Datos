# Ejercicio 5.1.3
Funcion Potencia(a, b: entero): entero Es # a^b
	Si (b = 0) Entonces # Caso Base
		Potencia:= 1
	Sino # Caso Recursivo
		Potencia:= a * Potencia(a, b - 1)
	FinSi
FinFuncion

Funcion Potencia2(a, b: entero): entero Es # a^b
	Si (b = 1) Entonces # Caso Base
		Potencia:= a
	Sino # Caso Recursivo
		Potencia:= a * Potencia(a, b - 1)
	FinSi
FinFuncion
