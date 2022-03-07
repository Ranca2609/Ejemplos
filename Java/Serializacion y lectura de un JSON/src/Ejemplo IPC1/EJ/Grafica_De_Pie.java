/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */


import javax.swing.JFrame;
import javax.swing.JPanel;
import org.jfree.chart.ChartFactory;
import org.jfree.chart.ChartPanel;
import org.jfree.chart.JFreeChart;
import org.jfree.data.general.DefaultPieDataset;
import org.jfree.data.general.PieDataset;
import org.jfree.util.PublicCloneable;

/**
 *
 * @author Carlos Rangel
 */
public class Grafica_De_Pie extends JFrame {
    String titulo = "";
    
    public Grafica_De_Pie(String titulo, Object [] arreglo){
        super(titulo);
        
        this.titulo = titulo;
        
        if(arreglo instanceof Carro[]){
            System.out.println("Arreglo de carros");
        }else if(arreglo instanceof Vendedor[]){
            setContentPane(crearPanel());
            this.repaint();
            System.out.println("Arreglo de vendedores");
        }else{
            System.out.println("No se reconoce de que tipo es el arreglo, revisar la informacion ingresada.");
        }  
    }
    
    
    
   JPanel crearPanel(){
       JFreeChart grafica = crearChart(crearDataSet());
        return new ChartPanel(grafica);
   }
   
   
JFreeChart crearChart(PieDataset dataset){
 
        return ChartFactory.createPieChart(titulo, dataset, true, true, true);
}


    PieDataset crearDataSet(){
        int contaM = 0;
        int contaF =0; 
        DefaultPieDataset dataset = new DefaultPieDataset();
        for(Vendedor vendedor : Serializacion.vendedores){
            if(vendedor.getGenero() == 'M'){
                contaM++;
            }else{
                contaF++;
            }
        }
        dataset.setValue("Mujeres", contaF);
        dataset.setValue("Hombres", contaM);
        
        
        return dataset;
     }
}
