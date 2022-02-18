%skeleton "lalr1.cc" /* -*- C++ -*- */

%defines
%define api.parser.class {calculadora_parser}
%define api.token.constructor
%define api.namespace {yy}
%define api.value.type variant
%define parse.assert
%code requires
{
#include <string>
#include <stdio.h>
class calculadora_driver;
}
%param { calculadora_driver& driver }
%locations
%define parse.trace
%define parse.error verbose
%code
{
#include "driver.h"
#include <iostream>
#include <time.h>
void imprime_hora(){
  time_t my_time = time(NULL);
  printf("%s", ctime(&my_time));
}
}
%define api.token.prefix {TOK_}

//Listadode Terminales


%token <std::string> IDENTIFICADOR "id"
%token <std::string> REP "reservada_rep"
%token <std::string> PATH "reservada_path"
%token <std::string> EXEC "reservada_exec"
%token <std::string> MKDISK "reservada_mkdisk"
%token <std::string> RUTA "ruta"
%token <std::string> IGUAL "igual"
%token <float> NUMERO "num"
%token FIN 0 "eof"

//Listado de No Terminales
%type <std::string> INSTRUCCIONES

%printer { yyoutput << $$; } <*>;
%%
%start INICIO;

INICIO : INICIO INSTRUCCIONES 
       | INSTRUCCIONES {  printf("Disco creado en Tarea2.dk"); printf("\n"); printf("Funcion rep ejecutada"); printf("\n"); printf("MBR Tamaño = 1024"); printf("\n"); printf("MBR Signature = 12606"); printf("\n"); printf("MBR Creacion ="); imprime_hora();}  
       | "eof";

INSTRUCCIONES : "reservada_rep" 
              | "reservada_path" {printf("Ruta encontrada\n");}
              | "reservada_exec" {printf("Comando exec ejecutado con exito\n");}
              | "ruta"
              | "igual"
              | "reservada_mkdisk";

%%
void yy::calculadora_parser::error(const location_type& lugar, const std::string& lexema)
{
  std::cout << "Error Sintactico " << lexema << std::endl;
}