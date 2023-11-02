package client

import (
	"bytes"
	"log"
	"os"
)

// readFromFile метод чтения из файла
func readFromFile(filename string) (string, error) {
	// Открываем файл только для чтения
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	defer file.Close()

	// Читаем содержимое файла в буфер
	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(file)
	if err != nil {
		return "", err
	}

	// Возвращаем содержимое файла в виде строки
	return buffer.String(), nil
}

// readFromFiles метод для чтения из нескольких файлов
func readFromFiles() []string {
	fileNames := []string{"textUser1.txt", "textUser2.txt", "textUser3.txt", "textUser4.txt", "textUser5.txt"}
	reqs := make([]string, len(fileNames))
	var err error
	for i, fName := range fileNames {
		reqs[i], err = readFromFile(fName)
		if err != nil {
			log.Fatal(err)
		}
	}
	return reqs
}

// writeToFile записывает ответ от чат гпт в файл
func writeToFile(filename string, content string) error {
	// Открываем файл для записи, если он не существует, то создаем его
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	// Записываем содержимое в файл
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
