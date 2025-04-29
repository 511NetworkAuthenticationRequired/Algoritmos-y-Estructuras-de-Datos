# Ejercicio 5.2.2: como el ejercicio daba de ejemplo Venezuela (que no es un palíndromo), voy a hacer algoritmos que impriman el contenido de una estructura de datos al revéz.
# Usando una secuencia de caracters: Argentina --> anitnegrA
Procedimiento RevezSecuencia(palabra: secuencia de caracteres, letra: caracter) Es # |A|r|g|e|n|t|i|n|a| --> |a|n|i|t|n|e|g|r|A|
    Si (FDS(palabra)) Entonces
        Escribir(letra)
    Sino
        Avanzar(palabra, letra)
        RevezSecuencia(palabra, letra)
        Escribir(letra)
    FinSi
FinProcedimiento

# Usando una lista simple, con actual:= prim (antes de invocar el procedimiento).
Procedimiento RevezLista(actual: puntero a palabra) Es # actual -> |A/proximo| --> |r/proximo| ●●● |a/proximo| --> null
    # Ambiente
    palabra = Registro
        letra: caracter
        proximo: puntero a palabra
    FinRegistro
    # Proceso
    Si (actual.proximo = null) Entonces
      Escribir(^actual.letra)
    Sino
        RevezLista(actual.proximo)
        Escribir(^actual.letra)
    FinSi
FinProcedimiento

# Con arreglos, es requisito acotarlo, por lo que se asume el peor caso, donde "Pneumonoultramicroscopicsilicovolcanoconiosis" sea la palabra en cuestión. La posición 46 siempre va a estar vacía, y actúa como una marca de fin para el peor caso.
Procedimiento RevezArreglo(vector: Arreglo de [1...46] de caracteres, x: entero) Es
    Si ((vector[x + 1] = "")) Entonces
        Escribir(vector[x])
    Sino
        RevezArreglo(vector, x + 1)
        Escribir(vector[x])
    FinSi
FinProcedimiento
