/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

import java.io.Serializable;

/**
 *
 * @author Carlos Rangel
 */
public class Vendedor implements Serializable{
   
    private int dpi, ventas;
    private String nombre, correo, password;
    private char genero;
    
    public Vendedor(){
    }

    
    public boolean validacion(String correo, String password){
        return this.correo.equals(correo) && this.password.equals(password);
    }
    
    
    public Vendedor(int dpi, int ventas, String nombre, String correo, String password, char genero) {
        this.dpi = dpi;
        this.ventas = ventas;
        this.nombre = nombre;
        this.correo = correo;
        this.password = password;
        this.genero = genero;
    }

   // Vendedor() {
   //     throw new UnsupportedOperationException("Not supported yet."); //To change body of generated methods, choose Tools | Templates.
   // }

    public int getDpi() {
        return dpi;
    }

    public int getVentas() {
        return ventas;
    }

    public String getNombre() {
        return nombre;
    }

    public String getCorreo() {
        return correo;
    }

    public String getPassword() {
        return password;
    }

    public char getGenero() {
        return genero;
    }

    public void setDpi(int dpi) {
        this.dpi = dpi;
    }

    public void setVentas(int ventas) {
        this.ventas = ventas;
    }

    public void setNombre(String nombre) {
        this.nombre = nombre;
    }

    public void setCorreo(String correo) {
        this.correo = correo;
    }

    public void setPassword(String password) {
        this.password = password;
    }

    public void setGenero(char genero) {
        this.genero = genero;
    }
    
    @Override
    public String toString(){
        return "\nDPI: " + this.dpi + "\nNombre: " + this.nombre + "\nVentas: " + this.ventas + "\nGenero: " + this.genero + "\nCorreo: " + this.correo + "\nPassword: " + this.password 
    ;
    }
    
}
