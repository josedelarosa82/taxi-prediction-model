package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("data/feb2023.csv") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)
	count := 0
	for fileScanner.Scan() {
		count += 1
		if count < 10 {
			/*line = strings.Replace(fileScanner.Text(), `"N`, ``, -1)
			arr_line = strings.SplitAfter(line, ",")*/
			arr_line := strings.Split(fileScanner.Text(), ",")
			for k, v := range arr_line {
				if v == `"N` {
					arr_line[k] = ""

				}
				arr_line[k] = strings.Replace(arr_line[k], `"`, ``, -1)
			}

			fmt.Println(arr_line)
			fmt.Println("---------------------------------------")
			//fmt.Println(servicio.Id + " " + servicio.Telefono_origen + " " + servicio.Latitud + " " + servicio.Estado + " " + servicio.Fecha_completa)
		} else {
			break
		}
	}

	file.Close()

}
