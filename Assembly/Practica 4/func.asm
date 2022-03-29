
;--------Imprimir cadenas-----------------------------
print macro buffer
push ax
push dx

	mov ax, @data
	mov ds,ax
	mov ah,09h
	mov dx,offset buffer
	int 21h
pop dx
pop ax
endm

;--------Terminar la ejecucion del programa----------------
salir macro
	.exit
endm

;---------Imprimir-----------------------------
Imprimir macro dato
	mov al, dato
	mov var, al
	print var
endm

;--------Obtener ruta---------------------------------
Obtener_Ruta macro buffer

	LOCAL Obtener_char, FIN
	MOV si,09
	Obtener_char:
	entrada
	cmp al, 0dh
	je FIN
	mov buffer[si],al
	inc si
	jmp Obtener_char
	FIN:
	mov al, 00h
	mov buffer[si],al
endm

;---------Obtener caracter de entrada------------------
entrada macro

	mov ah, 01h
	int 21h
endm
;----------Abre un archivo por su ruta----------------
Open macro ruta,handle

mov ah,3dh
mov al,00h
lea dx,ruta
int 21h
mov handle,ax
jc Error
endm

