# Ejercicio 5.1.2
Funcion Fibonacci(numero: entero): entero Es
	Si ((numero = 1) o (numero = 2)) Entonces # Caso Base
		Fibonacci:= 1
	Sino # Caso Recursivo
		Fibonacci:= Fibonacci(numero - 1) + Fibonacci(numero - 2)
	FinSi
FinFuncion
