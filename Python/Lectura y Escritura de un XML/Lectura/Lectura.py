from xml.etree import ElementTree as ET 
from  Constructor import Empleado


#----------Lectura de un XML----------------

tree = ET.parse('Archivo.xml')
root = tree.getroot()
Array = []
for Elemento in root:
    id = Elemento.attrib['id']
    for Subelemento in Elemento:
        if Subelemento.tag == 'nombre':
            Nombre = Subelemento.text.strip(' ')
        if Subelemento.tag == 'username':
            Username = Subelemento.text.strip(' ')
        if Subelemento.tag == 'password':
            Password = Subelemento.text.strip(' ')
    Array.append(Empleado(id, Nombre, Username, Password))

for i in Array:
    print('Id: '+i.id + '\nNombre: '+ i.nombre + "\nUsername: " + i.username
          + '\nPassword:' + i.password)

'''
----------Salida en Consola----------
Id: 1
Nombre: Angel Daniel
Username: AD
Password:123456
Id: 2
Nombre: Jose Pimentel
Username: JP
Password:ejemplo123@
'''
        
#----------Escritura de un XML---------------
# Se utiliza la libreria Minidoom para la escritura de un xml

   
        
    

    
    
