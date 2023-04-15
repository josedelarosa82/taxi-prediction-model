package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type ServicioData struct {
	Id               string
	Usuario          string
	Telefono_origen  string
	Idciudad         string
	Latitud          string
	Latitud_destino  string
	Longitud         string
	Longitud_destino string
	Estado           string
	Origen           string
	Fecha_completa   string
	Dia_de_semana    string
	Hora             string
}

func main() {
	name_unified := "data/SERVICIO_UNIFICADO_2022.csv"
	fil := []string{"data/SERVICIO_2022-05-29.csv",
		"data/HIS_SERVICIO_04_2022-05-29.csv",
		"data/HIS_SERVICIO_03_2022-05-29.csv",
		"data/HIS_SERVICIO_02_2022-05-29.csv",
		"data/HIS_SERVICIO_01_2022-05-29.csv"}
	file_unified := [][]string{}
	for _, name := range fil {
		fmt.Println("Procesando el file: ", name)
		csvFile, err := os.Open(name)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Successfully Opened CSV file")
		defer csvFile.Close()

		csvLines, err := csv.NewReader(csvFile).ReadAll()
		if err != nil {
			fmt.Println(err)
		}
		file_unified = append(file_unified, csvLines...)
	}
	print("Tamaño a procesar: ", len(file_unified))
	writer(name_unified, file_unified)

	fmt.Println("Procesando el file unificado: ", name_unified)
	csvFile, err := os.Open(name_unified)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	/*csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	count := 0
	for _, line := range csvLines {
		servicio := ServicioData{
			Id:               line[0],
			Idciudad:         line[1],
			Usuario:          line[2],
			Telefono_origen:  line[3],
			Latitud:          line[4],
			Latitud_destino:  line[5],
			Longitud:         line[6],
			Longitud_destino: line[7],
			Estado:           line[8],
			Origen:           line[9],
			Fecha_completa:   line[10],
			Dia_de_semana:    line[11],
			Hora:             line[12],
		}
		count += 1
		if count < 5 {
			fmt.Println(servicio.Id + " " + servicio.Telefono_origen + " " + servicio.Latitud + " " + servicio.Estado + " " + servicio.Fecha_completa)
		}
	}*/

}

func writer(name string, serv [][]string) {
	csvFile, err := os.Create(name)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)

	for _, servRow := range serv {
		_ = csvwriter.Write(servRow)
	}
	csvwriter.Flush()
	csvFile.Close()
}
