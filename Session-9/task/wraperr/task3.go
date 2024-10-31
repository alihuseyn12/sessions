package wraperr

import (
	"fmt"
	"os"
)

func OpenFile(filename string) error {
	_, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("file not found", err)
	} else {
		return nil
	}

}
