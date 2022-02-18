%{
# include <cerrno>
#include <stdio.h>
#include <string>
#include "driver.h"
#include "parser.tab.hh"
#undef yywrap
#define yywrap() 1
static yy::location loc;
%}
%option noyywrap nounput batch debug noinput
%option case-insensitive
%option outfile="scanner.cpp"

DIGITO [0-9]
LETRA [a-zA-ZñÑ]
NUMERO {DIGITO}+("."{DIGITO}+)?
IDENTIFICADOR {LETRA}({LETRA}|{DIGITO}|"_")*
DELIMITADOR [\t\r\f" "]
SPACE " "+
RUTA \/.[^" "]+|{IDENTIFICADOR}"."{IDENTIFICADOR}
%%
"="	                {return yy::calculadora_parser::make_IGUAL(yytext,loc);}
"rep"	              {return yy::calculadora_parser::make_REP(yytext,loc);}
"-path"	            {return yy::calculadora_parser::make_PATH(yytext,loc);}
"exec"	            {return yy::calculadora_parser::make_EXEC(yytext,loc);}
"mkdisk"	          {return yy::calculadora_parser::make_MKDISK(yytext,loc);}
{NUMERO}            {return yy::calculadora_parser::make_NUMERO(strtol (yytext, NULL, 10),loc);}
{IDENTIFICADOR}	    {return yy::calculadora_parser::make_IDENTIFICADOR(yytext,loc);}
{RUTA}	            {return yy::calculadora_parser::make_RUTA(yytext,loc);}
{SPACE}             {}
{DELIMITADOR}       {}

.               {printf("ERROR LEXICO %s\n",yytext);}
<<EOF>>		{return yy::calculadora_parser::make_FIN(loc);}
%%
void calculadora_driver::iniciarScanner()
{
  yy_flex_debug = false;
  if(file.empty() || file == "-"){
    printf("No se encuentra el archivo de entrada");
    exit(1);
  }else{
    yyin = fopen (file.c_str (), "r");
  }
}
void calculadora_driver::terminarScanner()
{
  fclose(yyin);
}