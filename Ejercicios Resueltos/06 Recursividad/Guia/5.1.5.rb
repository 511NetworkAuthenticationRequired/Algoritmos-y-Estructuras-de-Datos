# Ejercicio 5.1.5
Funcion Division(a, b: entero): entero Es
	Si (a < b) Entonces
		Division:= 0
	Sino
		Division:= 1 + Division(a - b, b)
	FinSi
FinFuncion
