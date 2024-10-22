package logs

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type LogMessage struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func CreateLog(message string, data interface{}) {
	msg := LogMessage{message, data}
	log_in_file(msg)

}

func log_in_file(msg LogMessage) {

	fileExist := check_file_exist()

	var (
		file *os.File
		err  error
	)

	if fileExist {
		file, err = open_file()
		fmt.Println(file.Read)
	} else {
		file, err = create_log_file()
		check_error(err)
	}

	defer file.Close()

	data, err := parse_to_json_string(&msg)
	fmt.Println(data)
	check_error(err)

	file.WriteString("Anthony\n")
	file.WriteString(data)
	file.Sync()
	w := bufio.NewWriter(file)
	w.Flush()

}
func parse_to_json_string(value *LogMessage) (string, error) {
	jsonData, err := json.Marshal(*value)
	if err != nil {
		return "", err
	}
	return string(jsonData) + "\n", nil

}
func create_log_file() (*os.File, error) {

	f, err := os.Create("/tmp/log-test.txt")
	check_error(err)
	return f, nil
}

func check_file_exist() bool {

	_, err := os.Stat("/tmp/log-test.txt")
	if err != nil {
		return false
	}
	return true

}

func open_file() (*os.File, error) {
	f, err := os.OpenFile("/tmp/log-test.txt", os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func check_error(er error) {

	if er != nil {
		log.Fatalf("Error ocorred %v", er)
	}
}
