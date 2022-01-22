/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package lectura_y_insercion_de_una_matriz;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.util.Scanner;

/**
 *
 * @author Carlos Rangel
 */
public class Lectura_y_insercion_de_una_matriz {
    static int numero_de_filas;
    static int numero_de_columnas;
    static Scanner sca = new Scanner(System.in);
    static int[][] Matriz;
    private static String Ruta;
    static String linea;
    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
      File archivo = null;
      FileReader fr = null;
      BufferedReader br = null;
      try {
         System.out.println("Ingrese la ruta de la matriz: ");
         String ruta = sca.nextLine();
         archivo = new File (ruta);
         Ruta = ruta;
         fr = new FileReader (archivo);
         br = new BufferedReader(fr);
         String[] numeros = null;
         int Count = 0; 
         while((linea=br.readLine())!=null){
            numeros = linea.split(",");
            Count++;
         }
         numero_de_filas = Count;
         numero_de_columnas = numeros.length;
         Matriz = new int [numero_de_filas][numero_de_columnas];
         System.out.println("El numero de filas es: " + numero_de_filas + "\n" +"El numero de columnas es: "+ numero_de_columnas);
      }
      catch(Exception e){
         e.printStackTrace();
      }finally{
         try{                    
            if( null != fr ){   
               fr.close();     
            }                  
         }catch (Exception e2){ 
            e2.printStackTrace();
         }
      }        
        try {
            archivo = new File(Ruta);
            fr = new FileReader(archivo);
            br = new BufferedReader(fr);
            String linea;

            int numero_filas = 0;

            while ((linea = br.readLine()) != null) {

                String[] numeros = linea.split(",");
                for (int i = 0; i < numeros.length; i++) {
                    Matriz[numero_filas][i] = Integer.parseInt(numeros[i]);
                }
                numero_filas++;
            }
        } catch (Exception e) {
            e.printStackTrace();
        } finally {
            try {
                if (null != fr) {
                    fr.close();
                }
            } catch (Exception e2) {
                e2.printStackTrace();
            }
        }        
        System.out.println("Matriz Ingresada");
        
        for (int i = 0;i< numero_de_filas ; i++){
            for(int j = 0; j<numero_de_columnas; j++){
                System.out.print( "["+Matriz[i][j]+ "]");
            }
            System.out.println("");               
        }
        
    }
    
}

