package logger

import (
	util "api-spammer/internal/utils"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const (
	ColorDefault Color = "\033[0m"
	ColorError   Color = "\033[31m"
	ColorSuccess Color = "\033[32m"
	ColorWarning Color = "\033[33m"
)

type Color string

var loggerFolder = fmt.Sprintf("./logs/%s", time.Now().Format("02.01.2006 15_04_05"))

func Init() {
	os.MkdirAll(loggerFolder, os.ModePerm)
}
func Log(color Color, messages ...any) {
	fmt.Print(color)
	fmt.Println(messages...)
	fmt.Print(ColorDefault)
}

func WriteLog(data interface{}) {
	path := filepath.Join(loggerFolder, fmt.Sprintf("%s.json", util.RandString(10)))

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
