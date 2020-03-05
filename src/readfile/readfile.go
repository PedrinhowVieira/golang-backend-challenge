package readfile

import (
	"encoding/csv"
	"fmt"
	"net/http"
)

func ReadFile(w http.ResponseWriter, r *http.Request) (records [][]string) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()
	records, err = csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	return
}