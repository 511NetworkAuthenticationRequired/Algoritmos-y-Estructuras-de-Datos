Este repositorio lo fui armando mientras cursaba, y está basado principalmente en las clases de teoría. Si entendí o anoté mal algo, existe una posibilidad de que también esté mal acá. Así que si notas algún fallo, no dudes en hacérmelo saber directamente a mí o bien, mediante el repositorio en si (issues, discussions, pull requests).

Recomiendo arrancar por la [introducción](Pseudocodigo/Introduccion), y luego ir viendo las [estructuras de control](Pseudocodigo/Estructuras%20de%20Control). Luego revisar las estructuras de datos que te interesen, pero el siguiente orden es el recomendado:
1. Secuencias
2. Archivos Secuenciales
3. Archivos Indexados
4. Arreglos
5. Árboles
6. Listas
   
Una vez que conoces las estructuras, te fijás bien [complejidad y eficiencia](Pseudocodigo/Complejidad%20Algoritmica). Recursividad recomiendo ver junto con árboles porque la misma naturaleza de la estructura permite visualizar "gráficamente", aunque por ahí si se te dan las matemáticas, algunos ejercicios iterativos te van a parecer más fáciles de resolver con recursión. Sin embargo, conviene que sigas el plan de estudio.

> [!NOTE]
> Aún me faltan subir los ejercicios resueltos, pero por ahora está -casi- toda la guía de recursividad. De todos modos, si vas a [ejercicios pascal](https://github.com/511NetworkAuthenticationRequired/Pascal), encontrarás más ejercicios de Pascal que siguen una lógica bastante parecida al pseudocódigo, y podrían ser útiles, sobre todo para recursividad.

## CONSEJOS

<details>
  <summary>hacé click acá para ver consejos para resolución</summary>
   
#### PARA RESOLUCIÓN
1. **Entender el enunciado y no iniciar el `PROCESO` de inmediato:**
   - Lee bien qué te pide el ejercicio y qué estructuras vas a utilizar. Asegúrate de comprender todos los detalles antes de comenzar a escribir.
   - No comiences el `PROCESO` hasta que tengas claro cómo vas a resolver el problema. Es fundamental saber qué vas a hacer antes de empezar a escribir el código.
2. **Plantea dudas:**
   - Si algo no te queda claro, pregúntalo hasta que lo entiendas completamente. No dudes en hacer tantas preguntas como necesites, siempre que sean sobre el enunciado.
3. **Empieza con el `AMBIENTE`:**
   - Una vez que tengas claro el enunciado, comienza a escribir el `AMBIENTE`. Esto te ayudará a organizar la información necesaria y a visualizar mejor el problema.
4. **Hacer bosquejos para secuencias:**
   - Si el ejercicio involucra secuencias, haz un bosquejo de cómo serían varias subsecuencias y cómo recorrer los ciclos.
5. **Visualizar arreglos:**
   - Para problemas con arreglos, especialmente matrices y estructuras de más de dos dimensiones, haz un gráfico o diagrama. Esto te va a ayudar a ver claramente la estructura de los datos.
6. **Verificar punteros en listas:**
   - Si el ejercicio involucra listas, asegúrate de que la asignación o reasignación de punteros esté correcta. Podés usar diagramas para visualizar cómo se realiza la manipulación de los punteros.

#### PARA RECURSIVIDAD EN CONCRETO
Recomiendo plantear un caso base, aunque sea inicial. Luego, recién el caso general y de ahí ir haciendo pruebas de escritorio.  
Si falla, se ve en qué, y se busca solucionar. Si tratando de arreglar termina no funcionando, es posible que el caso base sea el error.

En problemas que involucren matemática, plasmar eso en alguna ecuación o planteo matemático puede ayudar mucho a realizar la subacción recursiva.

Considero que la mejor forma de probar si funciona es traduciendo, o trabajando directamente, sobre algún lenguaje de programación.  
Yo recomiendo Pascal porque es muy similar, y más aún si te restringes siguiendo las pautas del pseudocódigo.  
Es fácil ver ahí si hay overflow (generalmente el compilador se va a detener y tira un runtime error), o si no funciona como debería.  
Si bien esto varía de persona a persona, se puede ir probando y, por fuerza bruta, entender lo que está pasando.

De hecho, muchos problemas matemáticos que "requieren" recursividad tienen una estructura similar a cómo se resuelven en pseudocódigo o en lenguajes de programación.  
En matemáticas, la recursividad se utiliza para resolver problemas dividiendo el problema grande en subproblemas más pequeños, lo cual también es lo que hace la recursividad en programación. Por ejemplo, problemas como el cálculo de Fibonacci, el cálculo de factoriales, o incluso algoritmos de búsqueda y ordenamiento como QuickSort, siguen una estructura recursiva tanto en su forma matemática como en su implementación en código.  
Esta relación entre matemática y programación hace que, si entiendes cómo resolver problemas recursivos en matemáticas, seas capaz de aplicar ese conocimiento directamente en programación.

</details>

## CONTRIBUCIONES/APORTES

<details>
  <summary>hacé click acá para ver como contribuir</summary>
   
#### 1. **ISSUES**
Los **Issues** son para discutir problemas, sugerencias o dudas. Si encontrás un error o tenés una propuesta para mejorar el proyecto, abrí un **Issue**. También podés participar en **Issues** existentes para discutir posibles soluciones.

Pasos para crear un **Issue**:
1. Ir a la sección de [Issues](../../issues).
2. Describir el problema o la sugerencia.
3. Participar en issues abiertos si tienes más información.

#### 2. **PULL REQUESTS**
Si ya tienes una solución lista o querés agregar una nueva característica, abrí un **Pull Request**. Los cambios deben estar listos y probados antes de hacer el pull. Este es el flujo ideal después de discutir cualquier mejora o corrección.

Pasos para crear un **Pull Request**:
1. Ir a la sección de [Pull Requests](../../pulls).
2. Crear una rama y realizar los cambios.
3. Crear un pull request con una descripción de los cambios.

#### 3. **DISCUSSIONS**
Si preferís una discusión más abierta o tienes preguntas generales sobre el proyecto, puedes abrir una **Discussion**. Esto es útil para conversar ideas o aclarar dudas sin necesidad de un **Issue** o **Pull Request**.

### REGLAS:
1. **Diagrama de Árboles o Listas:** Cuando quieras representar un árbol o lista, usar **Mermaid** para crear diagramas. Si por alguna razón no puedes usar **Mermaid**, podés utilizar **ASCII** con code blocks **SCSS**. Para previsualizar recomiendo usar el [sitio web de Mermaid](https://mermaid.js.org).
2. **Estructura del Texto:**
- Usa títulos con `#` y subtítulos de manera consistente en todo el documento. Esto facilita la lectura y mantiene una estructura clara.
- Si tenes links o "referencias" importantes, **siempre usa links relativos** para mantener la consistencia dentro del repositorio. Esto también asegura que las referencias sigan funcionando incluso si el repositorio se mueve o cambia.
3. **Orden en los Documentos:**
Cada archivo debe tener una estructura clara y ordenada. Seguí un flujo lógico, y usa títulos y subtítulos de manera apropiada para dividir los temas y secciones.
4. **Bloques de Código:**
Aunque no afecta funcionalmente, por convención, los bloques de código deben estar en formato **JS**. Esto ayuda a la consistencia y a la legibilidad del código dentro del repositorio.
5. **Recursos:** Ante la duda podés consultar [documentación GitHub](https://docs.github.com/en/get-started/writing-on-github/getting-started-with-writing-and-formatting-on-github/basic-writing-and-formatting-syntax) o [Stack Edit](https://stackedit.io/app), es último es útil para previsualizar.
  
</details>

## INDICE
<details>
  <summary>hacé click acá para ver el índice completo.</summary>

  ### INTRODUCCIÓN
  - [Operadores](Pseudocodigo/Introduccion/0.%20Operadores.md)
  - [Datos Simples](Pseudocodigo/Introduccion/1.%20Datos%20Simples.md)
  - [Tipos de Datos](Pseudocodigo/Introduccion/2.%20Tipos%20de%20Datos.md)
  - [Subacciones](Pseudocodigo/Introduccion/3.%20Subacciones.md)

  ### COMPLEJIDAD ALGORÍTMICA
  - [Complejidad Temporal & Espacial](Pseudocodigo/Complejidad%20Algoritmica/0.%20Complejidad%20Temporal%20%26%20Espacial.md)
  - [Operaciones Elementales](Pseudocodigo/Complejidad%20Algoritmica/1.%20Operaciones%20Elementales.md)

  #### ESTRUCTURAS DE CONTROL
  - [Secuenciales & Funciones Predefinidas](Pseudocodigo/Estructuras%20de%20Control/0.%20Secuenciales%20%26%20Funciones%20Predefinidas.md)
  - [Condicionales](Pseudocodigo/Estructuras%20de%20Control/1.%20Condicionales.md)
  - [Iterativas](Pseudocodigo/Estructuras%20de%20Control/2.%20Iterativas.md)

  ### ESTRUCTURAS DE DATOS

  #### ÁRBOLES
  - [Grafos & Árboles](Pseudocodigo/Estructuras%20de%20Datos/Arboles/0.%20Grafos%20%26%20Arboles.md)
  - [Binarios](Pseudocodigo/Estructuras%20de%20Datos/Arboles/1.%20Binarios.md)
  - [Recorridos](Pseudocodigo/Estructuras%20de%20Datos/Arboles/2.%20Recorridos.md)
  - [AVL, ABB & Expresión](Pseudocodigo/Estructuras%20de%20Datos/Arboles/3.%20AVL%2C%20ABB%20%26%20Expresion.md)

  #### ARCHIVOS
  - [Archivos & Registros](Pseudocodigo/Estructuras%20de%20Datos/Archivos/0.%20Archivos%20%26%20Registros.md)
  - [Genéricos & Emisión](Pseudocodigo/Estructuras%20de%20Datos/Archivos/1.%20Genericos%20%26%20Emision.md)
  - [Corte de Control](Pseudocodigo/Estructuras%20de%20Datos/Archivos/2.%20Corte%20de%20control.md)
  - [Mezcla](Pseudocodigo/Estructuras%20de%20Datos/Archivos/3.%20Mezcla.md)
  - [Actualización](Pseudocodigo/Estructuras%20de%20Datos/Archivos/4.%20Actualizacion.md)
  - [Tablas Comparativas](Pseudocodigo/Estructuras%20de%20Datos/Archivos/5.%20Tablas%20Comparativas.md)

  #### ARREGLOS
  - [Vectores & Matrices](Pseudocodigo/Estructuras%20de%20Datos/Arreglos/0.%20Vectores%20%26%20Matrices.md)
  - [Métodos de Ordenamiento](Pseudocodigo/Estructuras%20de%20Datos/Arreglos/1.%20Metodos%20de%20Ordenamiento.md)
  - [Métodos de Búsqueda](Pseudocodigo/Estructuras%20de%20Datos/Arreglos/2.%20Metodos%20de%20Busqueda.md)
  - [Procesos Estadísticos](Pseudocodigo/Estructuras%20de%20Datos/Arreglos/3.%20Procesos%20Estadisticos.md)

  #### LISTAS
  - [Particularizadas & Generalizadas](Pseudocodigo/Estructuras%20de%20Datos/Listas/0.%20Particularizadas%20%26%20Generalizadas.md)
  - [Tipos de Carga](Pseudocodigo/Estructuras%20de%20Datos/Listas/1.%20Tipos%20de%20Carga.md)
  - [Lista Simple](Pseudocodigo/Estructuras%20de%20Datos/Listas/2.%20Lista%20Simple.md)
  - [Lista Doble](Pseudocodigo/Estructuras%20de%20Datos/Listas/3.%20Lista%20Doble.md)
  - [Lista Circular](Pseudocodigo/Estructuras%20de%20Datos/Listas/4.%20Lista%20Circular.md)
  - [Lista Circular Doble](Pseudocodigo/Estructuras%20de%20Datos/Listas/5.%20Lista%20Circular%20Doble.md)

  #### SECUENCIAS
  - [Secuencias](Pseudocodigo/Estructuras%20de%20Datos/Secuencias/0.%20Secuencias.md)
  - [Subsecuencias](Pseudocodigo/Estructuras%20de%20Datos/Secuencias/1.%20Subsecuencias.md)

  #### RECURSIVIDAD
  - [Casos Base & General](Pseudocodigo/Recursividad/0.%20Casos%20Base%20%26%20General.md)
  - [Iteración vs Recursividad](Pseudocodigo/Recursividad/1.%20Iteracion%20VS%20Recursividad.md)
  - [Tipos de Recursividad](Pseudocodigo/Recursividad/2.%20Tipos%20de%20Recursividad.md)

</details>
