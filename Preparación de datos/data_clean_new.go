package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	name_unified := "data/SERVICIO_UNIFICADO_2023.csv"
	files := []string{"data/feb2023.csv"}

	file_unified := [][]string{}
	for _, name := range files {
		csvLines, err := readFile(name)
		if err != nil {
			fmt.Println(err)
		}
		file_unified = append(file_unified, csvLines)
	}
	writerFile(name_unified, file_unified)

}

func readFile(name string) ([]string, error) {
	const format_date = "2006-01-02 15:04:05"
	file, err := os.Open(name) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file_array := []string{}
	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	line := "ID,IDCIUDAD,USUARIO,TELEFONOORIGEN,LATITUD,LATITUDDESTINO,LONGITUD,LONGITUDDESTINO,ESTADO,ORIGEN,FECHACOMPLETA,DIADESEMANA,HORA"
	file_array = append(file_array, line)
	for fileScanner.Scan() {
		arr_line := strings.Split(fileScanner.Text(), ",")
		if len(arr_line) > 41 && arr_line[18] == "11001" {
			for k, v := range arr_line {
				if v == `"N` {
					arr_line[k] = ""
				}
				arr_line[k] = strings.Replace(arr_line[k], `"`, ``, -1)
			}
			//Id
			line = arr_line[0] + ","
			//ciudad
			line += arr_line[18] + ","
			//usuario
			line += arr_line[41] + ","
			//telefono origen
			line += arr_line[37] + ","
			//Latitud
			line += arr_line[22] + ","
			//Latitud_destino
			line += ","
			//Longitud
			line += arr_line[24] + ","
			//Longitud_destino
			line += ","
			//Estado
			line += arr_line[14] + ","
			//Origen
			line += arr_line[30] + ","

			date, error := time.Parse(format_date, arr_line[16])

			if error != nil {
				fmt.Println(error)
				fmt.Println(fileScanner.Text())
				//Fecha_completa
				line += ","
				//Dia_de_semana
				line += ","
				//Hora
				line += ""

			}

			//Fecha_completa
			line += date.Format(format_date) + ","
			//Dia_de_semana
			line += strconv.Itoa(int(date.Weekday())) + ","
			//Hora
			line += strconv.Itoa(date.Hour())

			file_array = append(file_array, line)
			//line := strings.Join(arr_line, ",")

		}
	}

	fmt.Printf("File %d procesed successfully!\n", name)

	return file_array, nil
}

func writerFile(name string, serv [][]string) {
	csvFile, err := os.Create(name)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer csvFile.Close()

	buffer := bufio.NewWriter(csvFile)

	for _, file_arr := range serv {
		fmt.Println("Amount to proseced in lines: ", len(file_arr))
		for _, servRow := range file_arr {
			_, err = buffer.WriteString(servRow + "\n")
			if err != nil {
				log.Fatal(err)
			}
		}
		//csvwriter.WriteString(servRow)
	}
	if err := buffer.Flush(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File %d created successfully!\n", name)
}
