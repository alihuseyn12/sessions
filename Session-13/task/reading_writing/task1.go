package reading_writing

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func PassOfStudent(file *os.File) {
	csvReader := csv.NewReader(file)
	studens, err1 := csvReader.ReadAll()

	if err1 != nil {
		fmt.Println("csvReader error is :", err1)

	}
	fmt.Println("Students with passing grades:")
	for i := 1; i < len(studens); i++ {
		in, err2 := strconv.Atoi(studens[i][2])
		if err2 != nil {
			fmt.Println("Atoi error is :", err2)
			return
		}
		if in > 60 {
			fmt.Println("-", studens[i][0])
		}

	}
}
