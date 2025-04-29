# Ejercicio 5.2.1: 234 --> 432
Procedimiento ImprimirCifras(numero: entero) Es # Con numero ≥ 0
    Si (numero < 10) Entonces
        Escribir(numero)
    Sino
       Escribir(numero DIV 10)
       ImprimirCifras(numero MOD 10)
    FinSi
FinProcedimiento
