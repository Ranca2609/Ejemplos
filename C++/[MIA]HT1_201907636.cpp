#include <iostream>
#include <fstream>
#include <stdio.h>
#include <conio.h>
#include <stdlib.h>
using namespace std;

    struct Profesor
    {
        int id_profesor;
        char  CUI[13];
        char   nombre[25];
        char curso[25];
    };

    struct Estudiante
    {
        int id_estudiante;
        char  CUI[13];
        char   nombre[25];
        char carnet[10];
    };

int main()
{

    Profesor profesor;
    Estudiante estudiante;

    int a, b, total;
    int seleccion;
    bool terminar = false;
    ofstream f;
    ifstream in;
    int tama = sizeof(Profesor);
    do
    {
        cout<<endl<<"Indica el n�mero de la operaci�n que desees realizar " <<endl;
        cout<<"Escribe 0 para terminar" << endl;
        cout<<"1. Registro Profesor" << endl;
        cout<<"2. Registro Estudiante" << endl;
        cout<<"3. Ver Registro" << endl;
        cout<<"0. Salir" << endl;
        cin>>seleccion;

        switch (seleccion) {
            case 1:
            {
                f.open("Archivo_Profesores.dat", ios::out|ios::app|ios::binary);
                if (f.fail())
                {
                cout << "Error al crear el fichero Archivo_Profesores.dat" << endl;
                system("pause");
                exit(1);
                }                
                string vacio;
                //ID
                cout << "Ingrese el id del Profesor: " <<endl;
                cin >> profesor.id_profesor;
                cout << "";
                getline(cin, vacio);
                //CUI
                string vacio2;
                cout << "Ingrese el CUI: " <<endl;
                cin.getline(profesor.CUI, 13);
                cout << "Presione enter para confirmar si escribio bien su numero de CUI " << endl;
                getline(cin, vacio2);
                //NOMBRE
                cout << "Ingrese el nombre del Profesor: " <<endl;
                cin.getline(profesor.nombre, 25);
                //CURSO
                cout << "Ingrese el curso del Profesor: " <<endl;
                cin.getline(profesor.curso, 25);
                cout <<"El id es: "<< profesor.id_profesor  << endl;
                cout <<"El nombre es: "<< profesor.nombre  << endl;
                cout <<"El CUI es: "<< profesor.CUI  << endl;
                cout <<"El curso es: "<< profesor.curso  << endl;
                f.write((char *)&profesor,tama);
                f.close();
                break;
            }
            case 2: {
                f.open("Archivo_Estudiantes.dat", ios::out|ios::app|ios::binary);
                if (f.fail())
                {
                cout << "Error al crear el ficheroArchivo_Estudiantes.dat" << endl;
                system("pause");
                exit(1);
                }                
                string vacio;
                //ID
                cout << "Ingrese el id del Estudiante: " <<endl;
                cin >> estudiante.id_estudiante;
                cout << "";
                getline(cin, vacio);
                //CUI
                string vacio2;
                cout << "Ingrese el CUI: " <<endl;
                cin.getline(estudiante.CUI, 13);
                cout << "Presione enter para confirmar si escribio bien su numero de CUI " << endl;
                getline(cin, vacio2);
                //NOMBRE
                cout << "Ingrese el nombre del Estudiante: " <<endl;
                cin.getline(estudiante.nombre, 25);
                //CARNET
                cout << "Ingrese el carnet del Estudiante: " <<endl;
                cin.getline(estudiante.carnet, 25);
                cout <<"El id es: "<< estudiante.id_estudiante  << endl;
                cout <<"El nombre es: "<< estudiante.nombre  << endl;
                cout <<"El CUI es: "<< estudiante.CUI  << endl;
                cout <<"El carnet es: "<< estudiante.carnet  << endl;
                f.write((char *)&estudiante,tama);
                f.close();
                break;
            }
            case 3: {
                cout << "-----------LISTADO DE ESTUDIANTES-----------" << endl;
                in.open("Archivo_Estudiantes.dat", ios::binary);
                 if(in.fail())
                {
                cout << "Error al abrir el fichero empleados.dat" << endl;
                system("pause");
                exit(1);
                }  
                in.read((char *) &estudiante, tama);
                int contador = 0;
                while(!in.eof())
                {
                cout << "Estudiante numero: " << contador << endl; 
                cout <<"El id es: "<< estudiante.id_estudiante  << endl;
                cout <<"El nombre es: "<< estudiante.nombre  << endl;
                cout <<"El CUI es: "<< estudiante.CUI  << endl;
                cout <<"El carnet es: "<< estudiante.carnet  << endl;
                cout << "*-----------------------------*" << endl;
                in.read((char *) &estudiante, tama);
                contador ++;
                }
                in.close();

                cout << "-----------LISTADO DE PROFESORES-----------" << endl;
                in.open("Archivo_Profesores.dat", ios::binary);
                 if(in.fail())
                {
                cout << "Error al abrir el fichero Archivo_Profesores.dat" << endl;
                system("pause");
                exit(1);
                }  
                in.read((char *) &profesor, tama);
                int contador2 = 0;
                while(!in.eof())
                {
                cout << "Profesor numero: " << contador << endl; 
                cout <<"El id es: "<< profesor.id_profesor  << endl;
                cout <<"El nombre es: "<< profesor.nombre  << endl;
                cout <<"El CUI es: "<< profesor.CUI  << endl;
                cout <<"El curso es: "<< profesor.curso  << endl;
                cout << "*-----------------------------*" << endl;
                in.read((char *) &profesor, tama);
                contador2 ++;
                }
                in.close();               
                break;
            }

            case 0: {
                terminar = true;
                break;
            }
            default: {
                cout << "Opci�n no v�lida, vuelve a intentarlo" << endl;
                break;
            }
        }
    }
    while (not terminar);
    
    return 0;
}
