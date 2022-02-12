class Nodo{
    constructor(id, nombre){
        this.id = id;
        this.nombre = nombre;
        this.siguiente = null;
    }
}

class ListaEnlazada{
    constructor(){
    this.raiz = null;
    }

    insertar(id, nombre){
        let nodo_nuevo = new Nodo(id, nombre)
        if (this.raiz == null){
            this.raiz = nodo_nuevo;
        }else{
            let temporal = this.raiz;
            while(temporal.siguiente != null){
                temporal = temporal.siguiente;
            }
            temporal.siguiente = nodo_nuevo;
        }
    }
 
    imprimir(){
        let temporal = this.first
        while(temporal != null){
            console.log('Id: ' + temporal.id +
                        '\nNombre: ' + temporal.nombre)
            temporal = temporal.siguiente
        }

    }



}
Prueba = new ListaEnlazada()
Prueba.insertar(1, 'Carlos')
Prueba.insertar(1, 'Katherine')
Prueba.imprimir()