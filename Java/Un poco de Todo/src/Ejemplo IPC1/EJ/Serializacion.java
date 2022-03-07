/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

import com.google.gson.Gson;
import java.io.BufferedReader;
import java.io.File;
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.FileReader;
import java.io.IOException;
import java.io.ObjectInputStream;
import java.io.ObjectOutputStream;

/**
 *
 * @author Carlos Rangel
 */
public class Serializacion {
    
    public static Carro carros[];
    public static Vendedor vendedores[];
    static String ruta_carros;

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) throws IOException {
        
//       SerializarDefault();
       cargarDatosSerializados();
       InicioInterfaz();
       
       //Lectura Carros
//       String textoJSON = leerArchivo(ruta);
//       System.out.println(textoJSON); 
       //carros = (Carro [])lerrJSON(textoJSON, new Carro());
       //for(Carro carro: carros){
       //    System.out.println(carro + "\n");
       //}
       
       
       
       //Lectura Vendedores
//       String textoVendedores = leerArchivo(ruta);
//       vendedores = (Vendedor[]) lerrJSON(textoVendedores, new Vendedor());
//       System.out.println(textoVendedores);
       //for(Vendedor vendedor: vendedores){
        //   System.out.println(vendedor);
      // }
    }
    
    

    static void SerializarDefault(String ruta_carros) throws IOException{
         String textoCarros =  leerArchivo(ruta_carros);
         carros =(Carro[]) lerrJSON(textoCarros, new Carro());
         serializarArreglo(carros, "Carros.dat");
    }
    
    static void InicioInterfaz(){
       Interfaz principa = new Interfaz();
       principa.setVisible(true);
    }
    
    static void cargarDatosSerializados() throws IOException{

       File archivoCarros, archivoVendedores;
       archivoCarros = new File("Carros.dat");
       archivoVendedores = new File("Vendedores.dat");
       
       if(archivoCarros.exists() && !archivoCarros.isDirectory()){
           carros =(Carro[]) leerArregloSerializado("Carros.dat");
           System.out.println("Se realizo la carga serializada de los carros correctamente");
           
       }
       if(archivoVendedores.exists() && !archivoVendedores.isDirectory()){
           vendedores = (Vendedor[])leerArregloSerializado("Vendedores.dat");
           System.out.println("Se realizo la carga serializada de los vendedores correctamente");
       }
    }
    
    static Object[] leerArregloSerializado(String ruta){
        try{
            ObjectInputStream leer_archivo = new ObjectInputStream(new FileInputStream(ruta));
            Object arreglo_recuperado [] = (Object[]) leer_archivo.readObject();
            leer_archivo.close();
            System.out.println("Se recupero la informacion " + ruta + " correctamente");
            return arreglo_recuperado;
        }catch(Exception e){
            System.out.println("Ocurrio un error inesperado  \n");
            e.printStackTrace();
        }
        
        return null;
    }

    static void serializarArreglo(Object arreglo [], String ruta){
   try{
        ObjectOutputStream escribir_archivo = new ObjectOutputStream(new FileOutputStream(ruta));
        escribir_archivo.writeObject(arreglo);
        escribir_archivo.close();
        System.out.println("Se serializo el arrelo correctamente");
   }catch(Exception e){
        System.out.println("Ocurrio un error inesperado");
        e.printStackTrace();
    }
    }
    
    static String leerArchivo(String ruta) throws IOException {
        String texto = "";
        BufferedReader lector = null;
        try {
            File archivo = new File(ruta);
            lector = new BufferedReader(new FileReader(archivo));

            String linea = lector.readLine();

            while (linea != null) {
                texto += linea;
                linea = lector.readLine();
            }
        } catch (Exception e) {
            System.out.println("Ocurrió un error");
            e.printStackTrace();
        } finally {
            if (lector != null) {
                lector.close();
            }
        }
        return texto;
    }
    
    static Object[] lerrJSON(String textoJSON, Object obj){
        Gson gson = new Gson();   
//        Carro carros [] = gson.fromJson(textoJSON, Carro[].class);    
        if (obj instanceof Carro) {
            return gson.fromJson(textoJSON, Carro[].class);
        }else if(obj instanceof Vendedor){
            return gson.fromJson(textoJSON, Vendedor[].class);
            
        }
        return gson.fromJson(textoJSON, Cliente[].class);
    }
}

