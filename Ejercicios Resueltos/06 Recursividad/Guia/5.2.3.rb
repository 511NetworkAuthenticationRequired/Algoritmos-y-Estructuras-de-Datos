# Ejercicio 5.3.4
# Es una modificación del planteo para el ejercicio anterior (se modifico el tipo de dato del campo dato del nodo), con actual:= prim (antes de invocar el procedimiento).
Procedimiento RevezLista(actual: puntero a nodo) Es # actual -> |3/proximo| --> |2/proximo| ●●● |10/proximo| --> null
    # Ambiente
    nodo = Registro
        numero: entero
        proximo: puntero a nodo
    FinRegistro
    # Proceso
    Si (actual.proximo = null) Entonces
      Escribir(^actual.nodo)
    Sino
        RevezLista(actual.proximo)
        Escribir(^actual.nodo)
    FinSi
FinProcedimiento
