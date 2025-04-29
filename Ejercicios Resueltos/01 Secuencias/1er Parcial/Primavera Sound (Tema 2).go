 ACCION PrimaveraSoundT2 ES
	AMBIENTE
		sec_fila: secuencia de caracteres
		ven_fila: carater
		sec_compras: secuencia de caracteres
		ven_compras: carater
		sec_salida: secuencia de caracteres
		hora: entero
		cont_abandono: entero
		PROCEDIMIENTO Inicilizar() ES
			ARR(sec_fila)
			AVZ(sec_fila, ven_fila)
			ARR(sec_compras)
			AVZ(sec_compras, ven_compras)
			CREAR(sec_salida)
			cont_abandono:= 0
		FP
		FUNCION ConvertirCAE(num: caracter):entero ES // Convierto un número en caracter a entero
			SEGUN (num) HACER
				="0": ConvertirCAE:= 0
				="1": ConvertirCAE:= 1
				="2": ConvertirCAE:= 2
				="3": ConvertirCAE:= 3
				="4": ConvertirCAE:= 4
				="5": ConvertirCAE:= 5
				="6": ConvertirCAE:= 6
				="7": ConvertirCAE:= 7
				="8": ConvertirCAE:= 8
				="9": ConvertirCAE:= 9
			FS
		FF
	PROCESO
		Inicilizar()
		MIENTRAS (NFDS(ven_compras)) HACER // Tienen correspondencia uno a uno asi que siempre se cumplría
			MIENTRAS (ven_fila <> "#") HACER
				PARA (i: 1 hasta 6) HACER
					AVZ(sec_compras, ven_compras)
				FP
				SI (ven_compras = "#") ENTONCES
					cont_abandono:= cont_abandono + 1
					AVZ(sec_compras, ven_compras)
					AVZ(sec_compras, ven_compras)
				SINO
					// SECUENCIA FILA
					hora:= ConvertirCAE(ven_fila)
					AVZ(sec_fila, ven_fila)
					hora:= (hora * 10) + ConvertirCAE(ven_fila)
					AVZ(sec_fila, ven_fila)
					AVZ(sec_fila, ven_fila)
					SI (hora < 10) ENTONCES
						PARA (x:= 1 hasta 6)
							AVZ(sec_fila, ven_fila)
							GRABAR(sec_salida, ven_fila)
						FP
						AVZ(sec_fila, ven_fila)
						AVZ(sec_fila, ven_fila)
						GRABAR(sec_salida, "+")
						// SECUENCIA COMPRAS
						MIENTRAS (ven_compras <> "?") HACER
							PARA (x:= 1 hasta 8) HACER
								AVZ(sec_compras, ven_compras)
								GRABAR(sec_salida, ven_compras)
							FP
							GRABAR(sec_salida, "+")
							MIENTRAS (ven_compras <> "." y ven_compras <> "?") HACER
								AVZ(sec_compras, ven_compras)
							FM
						FM
						AVZ(sec_compras, ven_compras)
						GRABAR(sec_salida, "#")
					CONTRARIO // Si fue a las/despues de las 10am
						PARA (i: 1 hasta 11) HACER// Avanzo todo ese usuario
							AVZ(sec_fila, ven_fila)
						FP
						MIENTRAS (sec_compras <> "?") HACER
							AVZ(sec_compras, ven_compras)
						FM
						AVZ(sec_compras, ven_compras)
					FS
					FS
				FS
			FM
		FM
		ESC("La cantidad de usuarios que abandonaron la fila es: ", cont_abandono)
		CERRAR(sec_fila)
		CERRAR(sec_compras)
		CERRAR(sec_salida)
FINPROCESO
