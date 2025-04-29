ACCION SupermercadoT1() ES
	AMBIENTE
		sec_stock: secuencia de caracteres
		ven_stock: carater
		sec_ventas: secuencia de caracteres // |M|M|D|D|FP|FE|UV|
		ven_ventas: carater
		sec_salida: secuencia de caracteres
		mes: alfanumerico
		mes_sec: entero
		uv_sucursal: entero
		uv_domicilio: entero
		stock: entero
		PROCEDIMIENTO Inicilizar() ES
			ARR(sec_stock)
			ARR(sec_ventas)
			AVZ(sec_ventas, ven_ventas)
			AVZ(sec_stock, ven_stock)
			CREAR(sec_salida)
			uv_domicilio_sucursal:= 0
			uv_domicilio:= 0
			stock:= 0
		FP
		FUNCION ConvertirMes(mes: caracter):entero ES
			SEGUN (mes) HACER
				="enero": ConvertirMes:= 1
				="febrero": ConvertirMes:= 2
				="marzo": ConvertirMes:= 3
				="abril": ConvertirMes:= 4
				="mayo":  ConvertirMes:= 5
				="junio": ConvertirMes:= 6
				="julio": ConvertirMes:= 7
				="agosto": ConvertirMes:= 8
				="septiembre": ConvertirMes:= 9
				="octubre": ConvertirMes:= 10
				="noviembre": ConvertirMes:= 11
				="diciembre": ConvertirMes:= 12
			FS
		FF
		FUNCION ConvertirCAE(var: caracter):entero ES
			SEGUN (var) HACER
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
		ESC("Para generar un informe, ingrese un mes (Ej: enero, diciembre): ") // ESC("Para generar un informe, ingrese un mes en formato numerico (Ej: Enero = 1, ..., Diciembre = 12)")
		LEER(mes)
		MIENTRAS (NFDS(sec_stock)) HACER // Tienen correspondencia uno a uno asi que cuando termina una secuencia, termina la otra también
			MIENTRAS (ven_ventas <> "#") HACER
				// SECUENCIA VENTAS | Guardo el mes del producto actual y verifica que tipo de envio fue 
				AVZ(sec_ventas, ven_ventas)
				AVZ(sec_ventas, ven_ventas)
				SI (ven_ventas = "0") ENTONCES
					AVZ(sec_ventas, ven_ventas)
					mes_sec:= ConvertirCAE(ven_ventas)
				CONTRARIO
					AVZ(sec_ventas, ven_ventas)
					mes_sec:= 10 + ConvertirCAE(ven_ventas)
				FS
				AVZ(sec_ventas, ven_ventas)
				AVZ(sec_ventas, ven_ventas)
				SI (ven_ventas = "S") ENTONCES
					uv_sucursal:= uv_sucursal + ConvertirCAE(ven_ventas) * 10
					AVZ(sec_ventas, ven_ventas)
					uv_sucursal:= uv_sucursal + ConvertirCAE(ven_ventas)
				CONTRARIO
					uv_domicilio:= uv_domicilio + ConvertirCAE(ven_ventas) * 10
					AVZ(sec_ventas, ven_ventas)
					uv_domicilio:= uv_domicilio + ConvertirCAE(ven_ventas)
				FS
			FM
			// SECUENCIA STOCK
			PARA (i: 1 hasta 6) HACER
				AVZ(sec_stock ven_stock)
			FP
			stock:= ConvertirCAE(ven_stock) * 100
			AVZ(sec_stock, ven_stock)
			stock:= stock + (ConvertirCAE(ven_stock) * 10)
			AVZ(sec_stock, ven_stock)
			stock:= stock + ConvertirCAE(ven_stock)

			SI (ConvertirMes(ven_stock) = mes_sec) ENTONCES
				MIENTRAS (ven_stock <> "&") HACER
					AVZ(sec_stock, ven_stock)
					ESC(ven_stock)
					SI ((stock - uv_sucursal) = 0) ENTONCES
						GRABAR(sec_salida, ven_stock)
					FS
				FM
				ESC(uv_domicilio)
				ESC(uv_sucursal)
			SINO
			MIENTRAS (ven_stock <> "&") HACER
				AVZ(sec_stock, ven_stock)
				SI ((stock - uv_sucursal) = 0) ENTONCES
					GRABAR(sec_salida, ven_stock)
				FS
			FM
		FM
		CERRAR(sec_ventas)
		CERRAR(sec_ventas)
		CERRAR(sec_salida)
FINACCION
