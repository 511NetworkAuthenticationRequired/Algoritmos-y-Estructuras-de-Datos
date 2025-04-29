# Ejercicio 5.1.8
Funcion SumaVector(Vector: Arreglo de [1...10] de entero, x: entero) Es
	Si (x = 10) Entonces
		SumaVector:= Vector[x]
	Sino
		SumaVector:= Vector[x] + SumaVector(Vector, x + 1)
	FinSi
FinFuncion
