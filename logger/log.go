package Logger

import (
	"api-spammer/rand"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const (
	Default = "\033[0m"
	Error   = "\033[31m"
	Success = "\033[32m"
	Warning = "\033[33m"
)

var Styles = map[string]string{
	"default": Default,
	"error":   Error,
	"success": Success,
	"warning": Warning,
}

var loggerFolder = fmt.Sprintf("./logs/%s", time.Now().Format("02.01.2006 15_04_05"))

func Init() {
	os.MkdirAll(loggerFolder, os.ModePerm)
}

func Print(style string, messages ...any) {
	if _, ok := Styles[style]; !ok {
		style = "default"
	}

	fmt.Print(Styles[style])
	fmt.Println(messages...)
	fmt.Print(Default)
}

func WriteLog(data interface{}) {
	path := filepath.Join(loggerFolder, fmt.Sprintf("%s.json", rand.String(10)))

	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		fmt.Println(err)
		return
	}
}
