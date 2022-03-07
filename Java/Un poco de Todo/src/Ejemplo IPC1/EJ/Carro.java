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
public class Carro implements Serializable{
    
    
    private int year, puertas;
    private String VIN, fabricante, modelo;
    private double precio;
    
    public Carro(){
    }

    public Carro(int year, int puertas, String VIN, String fabricante, String modelo, double precio) {
        this.year = year;
      
        this.VIN = VIN;
        this.fabricante = fabricante;
        this.modelo = modelo;
        this.precio = precio;
    }
    
    public String [] getArregloDatos(){
        String retorno [] =  {String.valueOf(this.VIN), String.valueOf(this.fabricante), String.valueOf(this.modelo),String.valueOf(this.year) ,String.valueOf(this.precio)};
        return retorno;
    }
    
    // Getters

    public int getYear() {
        return year;
    }



    public String getVIN() {
        return VIN;
    }

    public String getFabricante() {
        return fabricante;
    }

    public String getModelo() {
        return modelo;
    }

    public double getPrecio() {
        return precio;
    }

    
    // Setters 
    
    public void setYear(int year) {
        this.year = year;
    }

   

    public void setVIN(String VIN) {
        this.VIN = VIN;
    }

    public void setFabricante(String fabricante) {
        this.fabricante = fabricante;
    }

    public void setModelo(String modelo) {
        this.modelo = modelo;
    }

    public void setPrecio(double precio) {
        this.precio = precio;
    }
    


    
    @Override
    
    public String toString(){
        return"VIN: " +this.VIN +"\nModelo: " + this.modelo + "\nFabricante: " + this.fabricante + "\n" + "Precio: " +this.precio + "\nAño: " + this.year ;
    }
    
    
}
