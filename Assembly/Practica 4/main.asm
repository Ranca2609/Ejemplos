include func.asm

.model small
.stack
.data

;----------INFORMACION-------------
informacion db 0ah, 0dh, 'Universidad de San Carlos de Guatemala',
	       0ah, 0dh, 'Escuela de Ciencias y Sistemas',
	       0ah, 0dh, 'Arquitectura de Compiladores y ensamnbladores 1' , 
	       0ah, 0dh, 'Seccion B',0ah, 0dh, 'Carlos Roberto Rangel Castillo' , 
	       0ah, 0dh, '201907636',
'$'

salto db 0ah,0dh, '$' ,'$'

ent db 0ah, 0dh, 'Presione cualquier tecla para ver el menu: ','$'

division db 0ah, 0dh, '-----------------------------------------------------------','$'

bien db 0ah, 0dh, 'Bienvenido a la calculadora ','$'

carga_ruta db 0ah, 0dh, 'Introduzca la ruta ','$'

opciones db 0ah, 0dh, '1. Calculadora',
	    0ah, 0dh, '2. Archivo',
	    0ah, 0dh, '3. Salir' , 
	    0ah, 0dh, ' ' ,
'$'

;----------CODIGO-----------------
.code 

main proc

	Menu:
	     print informacion
	     print ent
	     entrada
	     cmp al,120
	     print opciones

	     entrada
	     cmp al,49
	     je CALCU
	     cmp al,50
	     je ARCHIVO
	     cmp al,51
	     je TERMINAR

	CALCU:
	      print bien
	      print division
	      JMP Menu

	ARCHIVO:
		print carga_ruta
		print division
		JMP Menu
	TERMINAR:
		salir
main endp
end main