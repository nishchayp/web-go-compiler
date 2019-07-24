package app

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t := template.Must(template.ParseFiles("template/index.html"))
	t.Execute(w, nil)
}

func Compile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	type FormData struct {
		Code string
	}

	r.ParseForm()

	formData := FormData{
		strings.Join(r.Form["code"], ""),
	}

	tmpFile, err := ioutil.TempFile(os.TempDir(), "input-")
	if err != nil {
		log.Fatal("Cannot create temporary file", err)
	}

	// Clean up the file afterwards
	defer os.Remove(tmpFile.Name())

	log.Println("Created File: " + tmpFile.Name())

	// Writing to the file
	text := []byte(formData.Code)
	if _, err = tmpFile.Write(text); err != nil {
		log.Fatal("Failed to write to temporary file", err)
	}

	// Close the file
	if err := tmpFile.Close(); err != nil {
		log.Fatal(err)
	}

	cmdStr := "./go-compiler/gc" + " " + tmpFile.Name()
	log.Println("cmdStr: " + cmdStr)

	cmd := exec.Command("sh", "-c", cmdStr)
	output, _ := cmd.CombinedOutput()
	if string(output) == "" {
		output = []byte("OK, Syntax")
	}
	log.Println(string(output))

	type Data struct {
		Code   string
		Output string
	}

	data := Data{
		Code:   formData.Code,
		Output: string(output),
	}

	t := template.Must(template.ParseFiles("template/index.html"))
	t.Execute(w, data)

}
