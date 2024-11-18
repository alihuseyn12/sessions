package json

import (
	"encoding/json"
	"fmt"
)

type ConfigFile struct {
	App_name   string   `json:"app_name"`
	Version    string   `json:"version"`
	Debug_mode bool     `json:"debug_mode"`
	Services   []string `json:"services"`
}

func CheckConfigFile(f []byte) {
	var config ConfigFile
	err1 := json.Unmarshal(f, &config)
	if err1 != nil {
		fmt.Println("Unmarshal err:", err1)
	}
	fmt.Println("Configuration Details:")
	fmt.Println("- App Name:", config.App_name)
	fmt.Println("- Version:", config.Version)
	if config.Debug_mode {
		fmt.Println("- Debug Mode: Enabled")
	} else {
		fmt.Println("- Debug Mode: Disabled")
	}
	fmt.Println("- Services:", len(config.Services))
	var s []string
	for _, service := range config.Services {
		//fmt.Println("- Service:", service)
		s = append(s, service)
	}
	fmt.Printf("- Service: %s,%s", s[0], s[1])

}
