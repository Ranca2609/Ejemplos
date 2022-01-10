package Analizador;
import java_cup.runtime.Symbol;




%%

%class lexico
%public 
%line
%char
%cup
%unicode
%ignorecase

%init{
    yyline = 1; 
    yychar = 1; 
%init}


BLANCOS=[ \r\t]+
D=[0-9]+
DD=[0-9]+("."[ |0-9]+)?
CADENA= [\"]([^\"\n]|(\\\"))*[\"]
ID=[A-Za-z]+["_"0-9A-Za-z]*


%%


";"                          { return new Symbol(sym.ptcoma, yyline, yychar, yytext()); }
"("                          { return new Symbol(sym.parizq, yyline, yychar, yytext()); }
")"                          { return new Symbol(sym.parder, yyline, yychar, yytext()); }
"+"                          { return new Symbol(sym.mas, yyline, yychar, yytext()); }
"-"                          { return new Symbol(sym.menos, yyline, yychar, yytext()); }
"*"                          { return new Symbol(sym.por, yyline, yychar, yytext()); }
"/"                          { return new Symbol(sym.dividido, yyline, yychar, yytext()); }
"{"                          { return new Symbol(sym.llavea, yyline, yychar, yytext()); }
"}"                          { return new Symbol(sym.llavec, yyline, yychar, yytext()); }
"="                          { return new Symbol(sym.igual, yyline, yychar, yytext()); }
"["                          { return new Symbol(sym.corchetea, yyline, yychar, yytext()); }
"]"                          { return new Symbol(sym.corchetec, yyline, yychar, yytext()); }
":"                          { return new Symbol(sym.dospuntos, yyline, yychar, yytext()); }
","                          { return new Symbol(sym.coma, yyline, yychar, yytext()); }
"if"                         { return new Symbol(sym.iff, yyline, yychar, yytext()); }
"else"                       { return new Symbol(sym.elsee, yyline, yychar, yytext()); }
">"                          { return new Symbol(sym.mayorque, yyline, yychar, yytext()); }
"<"                          { return new Symbol(sym.menorque, yyline, yychar, yytext()); }
"=="                         { return new Symbol(sym.igualar, yyline, yychar, yytext()); }
"!="                         { return new Symbol(sym.noigual, yyline, yychar, yytext()); }
"<="                         { return new Symbol(sym.menorigualque, yyline, yychar, yytext()); }
">="                         { return new Symbol(sym.mayorigualque, yyline, yychar, yytext()); }
"&&"                         { return new Symbol(sym.annd, yyline, yychar, yytext()); }
"||"                         { return new Symbol(sym.oor, yyline, yychar, yytext()); }
"!"                          { return new Symbol(sym.noot, yyline, yychar, yytext()); }
"**"                         { return new Symbol(sym.elevado, yyline, yychar, yytext()); }
"%"                          { return new Symbol(sym.modulo, yyline, yychar, yytext()); }
"print"                      { return new Symbol(sym.print, yyline, yychar, yytext()); }
"string"                     { return new Symbol(sym.str, yyline, yychar, yytext()); }
"int"                        { return new Symbol(sym.intt, yyline, yychar, yytext()); }
\n                           { yychar = 1; }
{BLANCOS}                  { }
{D}                        { return new Symbol(sym.entero, yyline, yychar, yytext()); }
{DD}                       { return new Symbol(sym.decimal, yyline, yychar, yytext()); }
{CADENA}                   { return new Symbol(sym.cadena, yyline, yychar, yytext()); }
{ID}                       { return new Symbol(sym.id, yyline, yychar, yytext()); }

. {
    System.out.println("Este es un error lexico: " +yytext()+", en la linea: "+yyline+", en la columna: "+yychar);
}



