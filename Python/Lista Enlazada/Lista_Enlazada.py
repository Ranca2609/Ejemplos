from Nodo import Persona

class Lista:
    def __init__(self):
        self.raiz = None
        
    def insertar(self, nombre, apellido, edad):
        
        Nodo_nuevo = Persona(nombre, apellido, edad)
        
        if self.raiz is None:
            self.raiz = Nodo_nuevo
        else:
            temporal = self.raiz
            while temporal.siguiente is not None:
                temporal = temporal.siguiente
            temporal.siguiente = Nodo_nuevo

    def imprimir(self):
        temporal = self.raiz
        while temporal != None:
            print('Nombre: '+ temporal.nombre +
                  '\nApellido: '+ temporal.apellido+
                  '\nEdad: '+ str(temporal.edad))
            temporal = temporal.siguiente

Prueba = Lista()
Prueba.insertar('Carlos', 'Rangel', 21)
Prueba.insertar('Luis', 'Castellanos', 20)
Prueba.imprimir()