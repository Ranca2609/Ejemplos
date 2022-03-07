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
import org.jfree.data.category.CategoryDataset;
import org.jfree.data.category.DefaultCategoryDataset;
import org.jfree.util.PublicCloneable;
/**
 *
 * @author Carlos Rangel
 */
public class GraficaBarras extends JFrame {
    String titulo;
    
    public GraficaBarras(String titulo) {
        
        
        super(titulo);
        
        this.titulo = titulo;
        
        setContentPane(crearPanel());
        
        
        this.repaint();
    }

    
    
    GraficaBarras() {
        
        
        throw new UnsupportedOperationException("Not supported yet."); //To change body of generated methods, choose Tools | Templates.
    }

    GraficaBarras(String barras, Carro[] carros) {
        
        throw new UnsupportedOperationException("Not supported yet."); //To change body of generated methods, choose Tools | Templates.
    }
    
    JPanel crearPanel(){
        JFreeChart grafica = crearChart(crearDataset());
        
        return new ChartPanel(grafica);
        
    }
    
    JFreeChart crearChart(CategoryDataset dataset){
        return ChartFactory.createBarChart(titulo, "", "Precios", dataset);
    }
    
    CategoryDataset crearDataset(){
        DefaultCategoryDataset dataset = new DefaultCategoryDataset();
        
        
        
        Carro top1 = Serializacion.carros[0];
        
        
        
        Carro top2 = Serializacion.carros[1];
        
        Carro top3 = Serializacion.carros[2];
        
        dataset.setValue(top1.getPrecio(), "Precio", top1.getFabricante() + " " + top1.getModelo() + " " + top1.getYear());
      
        
        dataset.setValue(top2.getPrecio(), "Precio", top2.getFabricante() + " " + top2.getModelo() + " " + top2.getYear());
    
        dataset.setValue(top3.getPrecio(), "Precio", top3.getFabricante() + " " + top3.getModelo() + " " + top3.getYear());
     
        return dataset;
    }
    
}
