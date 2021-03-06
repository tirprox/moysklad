package moysklad

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
)

func ToFile(v interface{}, filename string) {

	buf := make([]byte, 0)
	out := bytes.NewBuffer(buf)

	enc := json.NewEncoder(out)
	enc.SetEscapeHTML(false)

	err := enc.Encode(&v)

	newFile, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	out.WriteTo(newFile)

	defer newFile.Close()
}
