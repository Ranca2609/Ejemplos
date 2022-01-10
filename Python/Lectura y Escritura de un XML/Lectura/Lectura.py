from xml.etree import ElementTree as ET 
from  Constructor import Empleado
from xml.dom import minidom

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
# Se utiliza la libreria Minidom para la escritura de un xml

document = minidom.Document()
root = document.createElement('empresa')
document.appendChild(root)
for i in Array: 
    Empleado = document.createElement('empleado')  
    Empleado.setAttribute("id", i.id) 
    root.appendChild(Empleado)
    #----------------------------------------------------------
    Nombre = document.createElement('nombre')
    Nombre.appendChild(document.createTextNode(i.nombre))  
    Empleado.appendChild(Nombre)
    #----------------------------------------------------------
    Username = document.createElement('username')
    Username.appendChild(document.createTextNode(i.username))  
    Empleado.appendChild(Username)
    #----------------------------------------------------------
    Password = document.createElement('password')
    Password.appendChild(document.createTextNode(i.password))  
    Empleado.appendChild(Password)
    
xml = document.toprettyxml(indent='\t', newl='\n', encoding='utf-8') 
xml = xml.decode('utf-8')
Salida = open('C:/Users/Carlos Rangel/Documents/GitHub/Ejemplos/Python/Lectura y Escritura de un XML/Lectura/Salida.xml', 'w')
Salida.write(xml)
Salida.close()

    
    
