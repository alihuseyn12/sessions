package xml

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Optionss struct {
	AutoBackup     bool `xml:"auto_backup"`
	MaxConnections int  `xml:"max_connections"`
}
type Config struct {
	Database string   `xml:"database"`
	Username string   `xml:"username"`
	Password string   `xml:"password"`
	Options  Optionss `xml:"options"`
}

func EncodeFile(f *os.File) {
	config := Config{
		Database: "mydb",
		Username: "admin",
		Password: "secret",
		Options: Optionss{
			AutoBackup:     true,
			MaxConnections: 100,
		},
	}
	encoder := xml.NewEncoder(f)
	err1 := encoder.Encode(config)

	if err1 != nil {
		fmt.Println("writestrin err", err1)
	}

}
