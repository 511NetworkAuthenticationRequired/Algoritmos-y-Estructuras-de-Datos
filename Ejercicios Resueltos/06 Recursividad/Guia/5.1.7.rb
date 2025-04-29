# Ejercicio 5.1.7
Funcion Euclides(m, n: entero) Es # Con m > n, m∈ℕ, n∈ℕ
	Si (n = 0) Entonces
		Euclides:= m
	Sino
		Euclides:= Euclides(n, m MOD n)
	FinSi
FinFuncion
