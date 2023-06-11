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
	ProcessFiles()

}

func ProcessFiles() {
	fmt.Println("***Start process**")
	name_unified := "../data/SERVICIO_UNIFICADO_2023.csv"
	files := []string{"../data/Octubre-2022.csv", "../data/Noviembre-2022.csv", "../data/Diciembre-2022.csv", "../data/Enero-2023.csv", "../data/Febrero-2023.csv", "../data/Marzo-Abril-2023.csv"}

	file_unified := [][]string{}
	for _, name := range files {
		csvLines, err := readFile(name)
		if err != nil {
			fmt.Println(err)
		}
		file_unified = append(file_unified, csvLines)
	}
	writerFile(name_unified, file_unified)
	fmt.Println("***Finish process**")

}

func readFile(name string) ([]string, error) {
	const format_date_complete = "2006-01-02 15:04:05"
	const format_date = "2006-01-02"
	sw := true
	file, err := os.Open(name) // For read access.
	line := ""
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file_array := []string{}
	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	if name == "../data/Octubre-2022.csv" {
		line = "ID,CITY,USER,LATITUDEORI,LATITUDEDEST,LONGITUDEORI,LONGITUDEDEST,STATUS,CHANNEL,COMPLETEDATE,DATE,DAYOFWEEK,HOUR,MONTH"
		file_array = append(file_array, line)
	}
	fmt.Println("***************(", name, ")**********")
	for fileScanner.Scan() {
		arr_line := strings.Split(fileScanner.Text(), ",")
		if len(arr_line) > 8 {
			for k, v := range arr_line {
				if v == `"N` {
					arr_line[k] = ""
				}
				arr_line[k] = strings.Replace(arr_line[k], `"`, ``, -1)
			}
			//Id
			line = arr_line[0] + ","
			//City
			line += arr_line[1] + ","
			//User
			line += arr_line[2] + ","
			//Latitude_origin
			line += arr_line[3] + ","
			//Latitude_destination
			line += arr_line[4] + ","
			//Longitude_origin
			line += arr_line[5] + ","
			//Longitude_destination
			line += arr_line[6] + ","
			//Estatus
			line += arr_line[7] + ","
			//Channel
			line += arr_line[8] + ","
			date, error := time.Parse(format_date_complete, arr_line[9])
			if error != nil {
				fmt.Println(error)
				fmt.Println(fileScanner.Text())
			} else {
				if sw {
					fmt.Println("Before time (", date, ")")
				}
				//Se convierte a horario de colombia
				date := date.Add(time.Hour * -5)
				if sw {
					fmt.Println("After time (", date, ")")
					sw = false
				}
				//Complite_date
				line += date.Format(format_date_complete) + ","
				//Date
				line += date.Format(format_date) + ","
				//Dia_de_semana
				line += strconv.Itoa(changeDayOfWeek(int(date.Weekday())+1)) + ","
				//Hour
				line += strconv.Itoa(date.Hour()) + ","
				//Month
				line += strconv.Itoa(int(date.Month()))
				file_array = append(file_array, line)
			}
		}
	}
	fmt.Printf("File %v procesed successfully!\n", name)
	return file_array, nil
}

func changeDayOfWeek(week int) int {
	switch day := week; day {
	case 1:
		return 7
	case 2:
		return 1
	case 3:
		return 2
	case 4:
		return 3
	case 5:
		return 4
	case 6:
		return 5
	default:
		return 6
	}
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
	fmt.Printf("File %v created successfully!\n", name)
}
