package reading_writing

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func LineCount(file *os.File) {
	rd := bufio.NewReader(file)
	conunt := 0
	for {

		_, err := rd.ReadString('\n')
		conunt++
		if err == io.EOF {
			break
		}

	}
	fmt.Println("Total lines in file:", conunt)
}
