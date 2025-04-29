# Ejercicio 5.1.1
Funcion Factorial(numero: entero): entero Es # Con numero∈ℕ
	Si (numero = 0) Entonces # Caso Base
		Factorial:= 1
	Sino # Caso Recursivo (numero > 0)
		Factorial:= numero * Factorial(numero - 1)
	FinSi
FinFuncion
