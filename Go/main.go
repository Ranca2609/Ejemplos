package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
	comando()
}

func Obtain_Command_P(texto string) string {
	var tokens []string
	var comands []string
	var comands_together []string
	var comments []string
	if texto == "" {
		return ""
	}
	texto += " "
	var token string
	var r_comand string
	var char_comand string
	estate := 0
	for i := 0; i < len(texto); i++ {
		c := string(texto[i])
		if estate == 0 && c == "-" {
			estate = 1
		} else if estate == 0 && c == "#" {
			estate = 5
			continue
		} else if estate != 0 {
			if estate == 1 {
				if c == "=" {
					estate = 2
				} else if c == " " {
					continue
				}
			} else if estate == 2 {
				if c == " " {
					continue
				}
				if c == "\"" {
					estate = 3
					continue
				} else {
					estate = 4
				}
			} else if estate == 3 {
				if c == "\"" {
					estate = 4
					continue
				}
			} else if estate == 4 && c == "\"" {
				tokens = []string{}
				continue
			} else if estate == 4 && c == " " {
				estate = 0
				tokens = append(tokens, token)
				token = ""
				continue
			} else if estate == 5 {
				comments = append(comments, c)
				continue
			}
			token += c
		} else {
			if c != "" {
				estate = 0

				comands = append(comands, r_comand)
				r_comand = ""

			} else {
				continue

			}
			r_comand += c
		}
	}
	for i := 0; i < len(comands); i++ {
		char_comand += comands[i]
		comands_together = append(comands_together, char_comand)
	}
	longitud := len(comands_together)
	if longitud == 0 {
		fmt.Println("")
	} else {
		comando := comands_together[longitud-1]
		Identifier_Comand(comando, tokens)
	}
	return ""
}

var Path__ string
var Unit__ string
var Fit__ string
var Size__ string
var P string
var Type string
var Name string
var Info_Disco []string
var Discoo []string

func Identifier_Comand(comand string, tokens []string) {
	Comand_Compare := strings.ToLower(comand)
	if Comand_Compare == "mkdisk" || Comand_Compare == "mkdisk " {
		for i := 0; i < len(tokens); i++ {
			Params := strings.Split(tokens[i], "=")
			param_name := strings.ToLower(Params[0])
			if param_name == "size" {
				Size__ = Params[1]
				continue
			}
			if param_name == "fit" || param_name == "" {

				Fit__ = Params[1]
				continue
			}
			if param_name == "unit" {
				Unit__ = Params[1]
				continue
			}
			if param_name == "path" {
				Path__ = Params[1]
				continue
			}
		}
		general_disk(Size__, Fit__, Unit__, Path__)
		Info_Disco = append(Info_Disco, find_last_of(Path__)+"*"+Size__)

	} else if Comand_Compare == "rmdisk" || Comand_Compare == "rmdisk " {
		for i := 0; i < len(tokens); i++ {
			Params := strings.Split(tokens[i], "=")
			param_name := strings.ToLower(Params[0])
			if param_name == "path" {
				Path__ = Params[1]
				continue
			}
		}
		EliminarDisco(Path__)
	} else if Comand_Compare == "mkdir" || Comand_Compare == "mkdir " {
		for i := 0; i < len(tokens); i++ {
			Params := strings.Split(tokens[i], "=")
			param_name := strings.ToLower(Params[0])
			if param_name == "path" {
				Path__ = Params[1]
				continue
			}
			if param_name == "p" {
				P = Params[1]
				continue
			}
		}
		mkdir(Path__)
	} else if Comand_Compare == "fdisk" || Comand_Compare == "fdisk " {
		for i := 0; i < len(tokens); i++ {
			Params := strings.Split(tokens[i], "=")
			param_name := strings.ToLower(Params[0])
			if param_name == "path" {
				Path__ = Params[1]
				continue
			}
			if param_name == "type" {
				Type = Params[1]
				continue
			}
			if param_name == "unit" {
				Unit__ = Params[1]
				continue
			}
			if param_name == "name" {
				Name = Params[1]
				continue
			}
			if param_name == "size" {
				Size__ = Params[1]
				continue
			}
		}
		CrearParticiones(Type, Name, Size__, Path__)
	} else if Comand_Compare == "exec" || Comand_Compare == "exec " {
		Params := strings.Split(tokens[0], "=")
		Ruta := Params[1]

		file, err := os.Open(Ruta)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		var Lineas []string
		for scanner.Scan() {

			if scanner.Text() == "\n" || scanner.Text() == "" || scanner.Text() == " " {

			} else {
				time.Sleep(1 * time.Second)
				Obtain_Command_P(scanner.Text())
				Lineas = append(Lineas, scanner.Text())
			}
		}
	} else if Comand_Compare == "pause" || Comand_Compare == "pause " {
		var pause string
		fmt.Println("\033[1;33m" + "Presione entrer para continuar: " + "\033[0m")
		fmt.Scanln(&pause)
	} else {
		fmt.Println("Comando no reconocido.")
	}
}

//Estructures

type MBR struct {
	mbr_name_disk      string
	mbr_path           string
	mbr_tamano         int    //Tamaño total del disco en bytes
	mbr_fecha_creacion string //Fecha y hora de creación del disco
	mbr_dsk_signature  string //Número random, que identifica de forma única a cada disco
	dsk_fit            []byte //Tipo de ajuste de la partición. Tendrá los valores B (Best), F (First) o W (worst)
	mbr_partition      [4]Partition
}

type Partition struct {
	part_status string //Indica si la partición está activa o no
	part_type   string //Indica el tipo de partición, primaria o extendida. Tendrá los valores P o E
	part_fit    string //Tipo de ajuste de la partición. Tendrá los valores B (Best), F (First) o W (worst)
	part_start  int    //Indica en qué byte del disco inicia la partición
	part_size   int    //Contiene el tamaño total de la partición en bytes
	part_name   string //Nombre de la partición
}

func general_disk(size string, fit string, unit string, path string) {
	convertir_size, err := strconv.Atoi(size)
	if err != nil {
		log.Fatal(err)
	}
	if verificar_existencia(path) == true {
		verificar_param_obligatorios(convertir_size, fit, unit, path)
	} else {
		nombre := find_last_of(path)
		fmt.Println("\033[1;31m" + "Error: El disco con el nombre " + "\033[1;33m" + nombre + "\033[0m" + "\033[1;31m" + " ya existe, intente con otro nombre." + "\033[0m")
	}
}

func verificar_existencia(path string) bool {
	archivo := path
	if archivoExiste(archivo) {
		return false
	} else {
		return true
	}
}

func CreateDisk(size int, fit string, unit string, path string) {
	fmt.Println(path)
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Println("\033[1;31m" + "Ocurrio un error inesperado, intente denuevo." + "\033[0m")
	}
	var tmp int8 = 0
	s := &tmp
	var binario bytes.Buffer
	binary.Write(&binario, binary.BigEndian, s)

	if unit == "k" || unit == "K" {
		new_size := size * 1024
		for i := 0; i < new_size; i++ {
			escribirBytes(file, binario.Bytes())
		}
		_fit := identificar_fit(fit)
		dt := time.Now()

		var mbr_object MBR
		mbr_object.mbr_name_disk = find_last_of(path)
		mbr_object.mbr_path = path
		mbr_object.mbr_fecha_creacion = dt.String()
		mbr_object.mbr_tamano = size
		mbr_object.mbr_dsk_signature = strconv.Itoa(DiskSignature())
		mbr_object.dsk_fit = []byte(_fit)
		for i := 0; i < 4; i++ {
			mbr_object.mbr_partition[i].part_status = "0"
			mbr_object.mbr_partition[i].part_type = "P"
			mbr_object.mbr_partition[i].part_fit = fit
			mbr_object.mbr_partition[i].part_start = 0
			mbr_object.mbr_partition[i].part_size = 0
			mbr_object.mbr_partition[i].part_name = ""
		}
		fmt.Println("\033[1;32m" + "--------------DISCO CREADO CON EXITO :D-----------------" + "\033[0m")
		fmt.Println("\033[1;32m" + "Nombre del disco: " + "\033[0m" + "\033[1;33m" + find_last_of(path))
		fmt.Println("\033[1;32m" + "Ruta del disco: " + "\033[0m" + "\033[1;33m" + path + "\033[0m")
		fmt.Println("\033[1;32m" + "Fecha y hora de creacion: " + "\033[0m" + "\033[1;33m" + dt.String() + "\033[0m")
		fmt.Println("\033[1;32m" + "Tamaño: " + "\033[0m" + "\033[1;33m" + strconv.Itoa(size) + " kilobytes" + "\033[0m")
		fmt.Println("\033[1;32m" + "Signature del disco: " + "\033[0m" + "\033[1;33m" + strconv.Itoa(DiskSignature()) + "\033[0m")
		fmt.Println("\033[1;32m" + "Fit: " + "\033[0m" + "\033[1;33m" + _fit + "\033[0m")
		MBR_Graph(find_last_of(path), path, strconv.Itoa(size), strconv.Itoa(DiskSignature()), _fit, dt.String())
	} else if unit == "M" || unit == "m" {
		new_size := size * 1024 * 1024
		for i := 0; i < new_size; i++ {
			escribirBytes(file, binario.Bytes())
		}
		_fit := identificar_fit(fit)
		dt := time.Now()
		var mbr_object MBR
		mbr_object.mbr_name_disk = find_last_of(path)
		mbr_object.mbr_path = path
		mbr_object.mbr_fecha_creacion = dt.String()
		mbr_object.mbr_tamano = size
		mbr_object.mbr_dsk_signature = strconv.Itoa(DiskSignature())
		mbr_object.dsk_fit = []byte(_fit)
		for i := 0; i < 4; i++ {
			mbr_object.mbr_partition[i].part_status = "0"
			mbr_object.mbr_partition[i].part_type = "P"
			mbr_object.mbr_partition[i].part_fit = fit
			mbr_object.mbr_partition[i].part_start = 0
			mbr_object.mbr_partition[i].part_size = 0
			mbr_object.mbr_partition[i].part_name = ""
		}
		fmt.Println("\033[1;32m" + "--------------DISCO CREADO CON EXITO :D-----------------" + "\033[0m")
		fmt.Println("\033[1;32m" + "Nombre del disco: " + "\033[0m" + "\033[1;33m" + find_last_of(path))
		fmt.Println("\033[1;32m" + "Ruta del disco: " + "\033[0m" + "\033[1;33m" + path + "\033[0m")
		fmt.Println("\033[1;32m" + "Fecha y hora de creacion: " + "\033[0m" + "\033[1;33m" + dt.String() + "\033[0m")
		fmt.Println("\033[1;32m" + "Tamaño: " + "\033[0m" + "\033[1;33m" + strconv.Itoa(size) + " Megabytes" + "\033[0m")
		fmt.Println("\033[1;32m" + "Signature del disco: " + "\033[0m" + "\033[1;33m" + strconv.Itoa(DiskSignature()) + "\033[0m")
		fmt.Println("\033[1;32m" + "Fit: " + "\033[0m" + "\033[1;33m" + _fit + "\033[0m")
		MBR_Graph(find_last_of(path), path, strconv.Itoa(size), strconv.Itoa(DiskSignature()), _fit, dt.String())
	} else if unit == "B" || unit == "b" {
		for i := 0; i < size; i++ {
			escribirBytes(file, binario.Bytes())
		}
		_fit := identificar_fit(fit)
		dt := time.Now()
		var mbr_object MBR
		mbr_object.mbr_name_disk = find_last_of(path)
		mbr_object.mbr_path = path
		mbr_object.mbr_fecha_creacion = dt.String()
		mbr_object.mbr_tamano = size
		mbr_object.mbr_dsk_signature = strconv.Itoa(DiskSignature())
		mbr_object.dsk_fit = []byte(_fit)
		for i := 0; i < 4; i++ {
			mbr_object.mbr_partition[i].part_status = "0"
			mbr_object.mbr_partition[i].part_type = "P"
			mbr_object.mbr_partition[i].part_fit = fit
			mbr_object.mbr_partition[i].part_start = 0
			mbr_object.mbr_partition[i].part_size = 0
			mbr_object.mbr_partition[i].part_name = ""
		}
		fmt.Println("\033[1;32m" + "--------------DISCO CREADO CON EXITO :D-----------------" + "\033[0m")
		fmt.Println("\033[1;32m" + "Nombre del disco: " + "\033[0m" + "\033[1;33m" + find_last_of(path))
		fmt.Println("\033[1;32m" + "Ruta del disco: " + "\033[0m" + "\033[1;33m" + path + "\033[0m")
		fmt.Println("\033[1;32m" + "Fecha y hora de creacion: " + "\033[0m" + "\033[1;33m" + dt.String() + "\033[0m")
		fmt.Println("\033[1;32m" + "Tamaño: " + "\033[0m" + "\033[1;33m" + strconv.Itoa(size) + " Bytes" + "\033[0m")
		fmt.Println("\033[1;32m" + "Signature del disco: " + "\033[0m" + "\033[1;33m" + strconv.Itoa(DiskSignature()) + "\033[0m")
		fmt.Println("\033[1;32m" + "Fit: " + "\033[0m" + "\033[1;33m" + _fit + "\033[0m")
		MBR_Graph(find_last_of(path), path, strconv.Itoa(size), strconv.Itoa(DiskSignature()), _fit, dt.String())
	} else if unit == "" {
		new_size := size * 1024 * 1024
		for i := 0; i < new_size; i++ {
			escribirBytes(file, binario.Bytes())
		}
		_fit := identificar_fit(fit)
		dt := time.Now()
		fmt.Println("\033[1;32m" + "--------------DISCO CREADO CON EXITO :D-----------------" + "\033[0m")
		fmt.Println("\033[1;32m" + "Nombre del disco: " + "\033[0m" + "\033[1;33m" + find_last_of(path))
		fmt.Println("\033[1;32m" + "Ruta del disco: " + "\033[0m" + "\033[1;33m" + path + "\033[0m")
		fmt.Println("\033[1;32m" + "Fecha y hora de creacion: " + "\033[0m" + "\033[1;33m" + dt.String() + "\033[0m")
		fmt.Println("\033[1;32m" + "Tamaño: " + "\033[0m" + "\033[1;33m" + strconv.Itoa(size) + " Megabytes" + "\033[0m")
		fmt.Println("\033[1;32m" + "Signature del disco: " + "\033[0m" + "\033[1;33m" + strconv.Itoa(DiskSignature()) + "\033[0m")
		fmt.Println("\033[1;32m" + "Fit: " + "\033[0m" + "\033[1;33m" + _fit + "\033[0m")
		MBR_Graph(find_last_of(path), path, strconv.Itoa(size), strconv.Itoa(DiskSignature()), _fit, dt.String())
	} else {
		fmt.Println("\033[1;31m" + "Error: Caracter invalido, el caracter " + "\033[1;33m" + unit + "\033[0m" + "\033[1;31m" + " no pertenece al lenguaje." + "\033[0m\n")
	}
}

func verificar_param_obligatorios(size int, fit string, unit string, path string) {
	if path == "" {
		fmt.Println("\033[1;31m" + "Error: Falta el parametro obligatorio path" + "\033[0m")
	} else if size == 0 {
		fmt.Println("\033[1;31m" + "Error: El tamaño no puede ser 0 o nulo." + "\033[0m")

	} else if size < 0 {
		fmt.Println("\033[1;31m" + "Error: El tamaño no puede ser un numero negativo." + "\033[0m")
	} else {
		CreateDisk(size, fit, unit, path)
	}

}

func escribirBytes(file *os.File, bytes []byte) {
	_, err := file.Write(bytes)

	if err != nil {
		log.Fatal(err)
	}
}

func archivoExiste(ruta string) bool {
	if _, err := os.Stat(ruta); os.IsNotExist(err) {
		return false
	}
	return true
}

func DiskSignature() int {
	Signature_Disk := rand.Intn(600) * 1000
	return Signature_Disk
}

func find_last_of(path string) string {
	res1 := strings.Split(path, "/")
	longitud := len(res1)
	nombre_disco := res1[longitud-1]
	return nombre_disco

}

func identificar_fit(fit string) string {
	if fit == "ff" || fit == "wf" || fit == "bf" {
		return fit
	} else if fit == "" {
		new_fit := "ff"
		return new_fit
	} else {
		return "\033[1;31m" + "Error: Se ingreso un fit no perteneciente al Lenguaje." + "\033[0m"
	}
}

func EliminarDisco(path string) {
	nombre_disco := find_last_of(path)
	if path == "" {
		fmt.Println("\033[1;31m" + "Error: Falta el parametro obligatorio path." + "\033[0m")
	} else {
		fmt.Println("\033[1;34m" + "¿Desea eliminar el disco " + "\033[0m" + "\033[1;33m" + nombre_disco + "\033[0m" + "\033[1;34m" + " ?  [y/n]" + "\033[0m")
		var Opcion string
		fmt.Scanln(&Opcion)
		if Opcion == "Y" || Opcion == "y" {
			err := os.Remove(path)
			if err != nil {
				fmt.Println("\033[1;31m" + "Error: No existe el disco " + nombre_disco + " para ser eliminado" + "\033[0m")
			}
			fmt.Println("\033[1;32m" + "El disco " + "\033[0m" + "\033[1;33m" +
				nombre_disco + "\033[0m" + "\033[1;32m" +
				" fue eliminado exitosamente de la ruta " + "\033[0m" + "\033[1;33m" +
				path + "\033[0m")
		} else if Opcion == "N" || Opcion == "n" {
			fmt.Println("\033[1;32m" + "No se eliminara el disco " + "\033[0m" + "\033[1;33m" + nombre_disco + "\033[0m")
		}
	}
}

func crearDirectorioSiNoExiste(directorio string) {
	if _, err := os.Stat(directorio); os.IsNotExist(err) {
		err = os.Mkdir(directorio, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func mkdir(_path string) {

	nombre_directorio := find_last_of(_path)
	if len(nombre_directorio) < 12 {
		crearDirectorioSiNoExiste(_path)
	} else {
		fmt.Println("\033[1;31m" + "Error: El nombre " + "\033[0m" + "\033[1;33m" + nombre_directorio + "\033[0m" + "\033[1;31m" + " del directorio no puede excederse de 12 caracteres" + "\033[0m")
	}

}

var Atributos_Particion []string
var nombres_particiones []string

func CrearParticiones(type__ string, name string, size string, path string) {
	nombre_disco_comparar := find_last_of(path)
	for i := 0; i < len(Info_Disco); i++ {
		nuevo := strings.Split(Info_Disco[i], "*")
		if nuevo[0] == nombre_disco_comparar {
			if type__ == "E" || type__ == "E " {
				fmt.Println("\033[1;33m" + "Crearndo Particion Extendida" + "\033[0m")
			}
		}
		if type__ == "P" || type__ == "P " {
			fmt.Println("\033[1;33m" + "Crearndo Particion Primaria" + "\033[0m")
		}
		if type__ == "L" || type__ == "L " {
			fmt.Println("\033[1;33m" + "Crearndo Particion Logica" + "\033[0m")
		}
		nombres_particiones = append(nombres_particiones, name+"*"+type__)
		graph_disks(path, nombres_particiones)
	}
}

func comando() {
	iteraccion := 1
	for iteraccion < 1000 {
		var linea_comando string
		var segundo string
		var pause string
		fmt.Print("\033[1;33m" + ">>" + "\033[0m")
		fmt.Scanln(&linea_comando, &segundo)
		if linea_comando == "exit" {
			iteraccion = 1000
			fmt.Println("\033[1;33m" + "Se termino la ejecucion del programa" + "\033[0m")
		} else if linea_comando == "pause" {
			fmt.Println("\033[1;33m" + "Presione entrer para continuar: " + "\033[0m")
			fmt.Scanln(&pause)
		} else {
			Obtain_Command_P(linea_comando + " " + segundo)
			iteraccion += iteraccion
		}
	}
}

func graph_disks(path string, name []string) {
	nom_disco := find_last_of(path)
	nuevoo := strings.Split(nom_disco, ".")
	arch, _ := os.Create("archivo" + nuevoo[0] + ".dot")
	_, _ = arch.WriteString("digraph G{" + "\n")
	_, _ = arch.WriteString("node [shape=plaintext]" + "\n")
	_, _ = arch.WriteString("rankdir=T" + "\n")
	_, _ = arch.WriteString("A [label=<" + "\n")
	_, _ = arch.WriteString(`<TABLE BGCOLOR="#EE9B00" BORDER="2" COLOR="BLACK" CELLBORDER="1" CELLSPACING="0">` + "\n")
	_, _ = arch.WriteString("<TR>" + "\n")
	_, _ = arch.WriteString(`<TD ALIGN="LEFT" BGCOLOR="#CA6702" COLSPAN="2">` + "Nombre disco" + `</TD>` + "\n")
	_, _ = arch.WriteString(`</TR>` + "\n")
	_, _ = arch.WriteString(`<tr><td>Logica:+ ` + strconv.Itoa(rand.Intn(33)) + `</td><td>Extendida:` + strconv.Itoa(rand.Intn(66)) + `</td><td>Primaria:` + strconv.Itoa(rand.Intn(99)) + `</td></tr>` + "\n")
	_, _ = arch.WriteString(`</TABLE>` + "\n")
	_, _ = arch.WriteString(`>];` + "\n")
	_, _ = arch.WriteString("}" + "\n")
	//Crear Archivo
	app := "dot"
	arg0 := "-Tpng"
	arg1 := "./archivo" + nuevoo[0] + ".dot"
	arg2 := "-o"
	arg3 := "outfile" + nuevoo[0] + ".png"

	cmd := exec.Command(app, arg0, arg1, arg2, arg3)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(stdout))

}

func MBR_Graph(nombre string, path___ string, size string, signature string, fit string, fecha string) {

	arch, _ := os.Create("archivo" + size + ".dot")
	_, _ = arch.WriteString("digraph G{" + "\n")
	_, _ = arch.WriteString("node [shape=plaintext]" + "\n")
	_, _ = arch.WriteString("rankdir=T" + "\n")
	_, _ = arch.WriteString("A [label=<" + "\n")
	_, _ = arch.WriteString(`<TABLE BGCOLOR="#EE9B00" BORDER="2" COLOR="BLACK" CELLBORDER="1" CELLSPACING="0">` + "\n")
	_, _ = arch.WriteString("<TR>" + "\n")
	_, _ = arch.WriteString(`<TD ALIGN="LEFT" BGCOLOR="#CA6702" COLSPAN="2">` + "Nombre disco" + `</TD>` + "\n")
	_, _ = arch.WriteString(`<TD ALIGN="RIGHT" BGCOLOR="#CA6702" COLSPAN="2">` + nombre + `</TD>` + "\n")
	_, _ = arch.WriteString(`</TR>` + "\n")
	_, _ = arch.WriteString("<TR>" + "\n")
	_, _ = arch.WriteString(`<TD ALIGN="LEFT" BGCOLOR="#CA6702" COLSPAN="2">` + "Ruta" + `</TD>` + "\n")
	_, _ = arch.WriteString(`<TD ALIGN="RIGHT" BGCOLOR="#CA6702" COLSPAN="2">` + path___ + `</TD>` + "\n")
	_, _ = arch.WriteString(`</TR>` + "\n")
	_, _ = arch.WriteString("<TR>" + "\n")
	_, _ = arch.WriteString(`<TD ALIGN="LEFT" BGCOLOR="#CA6702" COLSPAN="2">` + "Tamaño" + `</TD>` + "\n")
	_, _ = arch.WriteString(`<TD ALIGN="RIGHT" BGCOLOR="#CA6702" COLSPAN="2">` + size + `</TD>` + "\n")
	_, _ = arch.WriteString(`</TR>` + "\n")
	_, _ = arch.WriteString("<TR>" + "\n")
	_, _ = arch.WriteString(`<TD ALIGN="LEFT" BGCOLOR="#CA6702" COLSPAN="2">` + "Signature" + `</TD>` + "\n")
	_, _ = arch.WriteString(`<TD ALIGN="RIGHT" BGCOLOR="#CA6702" COLSPAN="2">` + signature + `</TD>` + "\n")
	_, _ = arch.WriteString(`</TR>` + "\n")
	_, _ = arch.WriteString("<TR>" + "\n")
	_, _ = arch.WriteString(`<TD ALIGN="LEFT" BGCOLOR="#CA6702" COLSPAN="2">` + "Fit" + `</TD>` + "\n")
	_, _ = arch.WriteString(`<TD ALIGN="RIGHT" BGCOLOR="#CA6702" COLSPAN="2">` + fit + `</TD>` + "\n")
	_, _ = arch.WriteString(`</TR>` + "\n")
	_, _ = arch.WriteString("<TR>" + "\n")
	_, _ = arch.WriteString(`<TD ALIGN="LEFT" BGCOLOR="#CA6702" COLSPAN="2">` + "Fecha creacion" + `</TD>` + "\n")
	_, _ = arch.WriteString(`<TD ALIGN="RIGHT" BGCOLOR="#CA6702" COLSPAN="2">` + fecha + `</TD>` + "\n")
	_, _ = arch.WriteString(`</TR>` + "\n")
	_, _ = arch.WriteString(`</TABLE>` + "\n")
	_, _ = arch.WriteString(`>];` + "\n")
	_, _ = arch.WriteString("}" + "\n")
	//Crear Archivo
	app := "dot"
	arg0 := "-Tpng"
	arg1 := "./archivo" + size + ".dot"
	arg2 := "-o"
	arg3 := "outfile" + size + ".png"

	cmd := exec.Command(app, arg0, arg1, arg2, arg3)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(stdout))
}
