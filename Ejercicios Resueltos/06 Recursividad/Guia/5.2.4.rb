# Ejercicio 5.2.4
Procedimiento CambioBase(numero: entero) Es
    Si (numero < 2) Entonces
        Escribir(numero)
    Sino
        CambioBase(numero DIV 2)
        Escribir(numero MOD 2)
    FinSi
FinProcedimiento
