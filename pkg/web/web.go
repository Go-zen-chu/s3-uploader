package web

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/pkg/errors"
)

var (
	templatesPath = ""
)

// GetTemplatesPath : get path of templates
func GetTemplatesPath() (string, error) {
	if templatesPath == "" {
		pwd, err := os.Getwd()
		if err != nil {
			return "", errors.Wrap(err, "Error while getting templates path")
		}
		templatesPath = filepath.Join(pwd, "../../web/templates")
	}
	return templatesPath, nil
}

// GetTemplate : get template with filepath
func GetTemplate(tmplName string) (*template.Template, error) {
	tmplPath, err := GetTemplatesPath()
	if err != nil {
		return nil, errors.Wrap(err, "Unable to get templates path")
	}
	idxTmplPath := filepath.Join(tmplPath, tmplName)
	tmpl, err := template.ParseFiles(idxTmplPath)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to parse template")
	}
	return tmpl, err
}

// HandleIndex : Handle function for index.html
func HandleIndex(w http.ResponseWriter) error {
	tmpl, err := GetTemplate("index.tmpl")
	if err != nil {
		return errors.Wrap(err, "Error while getting template")
	}
	data := map[string]interface{}{}
	if err := tmpl.Execute(w, data); err != nil {
		return errors.Wrap(err, "Unable to execute template")
	}
	return nil
}

// IndexHandler : Handler
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if err := HandleIndex(w); err != nil {
		log.Fatalln(err)
	}
}

func HandleFileUpload(w http.ResponseWriter, r *http.Request) (error, int) {
	if r.Method != "POST" {
		return errors.New("Allowed POST method only"), http.StatusMethodNotAllowed
	}

	err := r.ParseMultipartForm(32 << 20) // maxMemory
	if err != nil {
		return err, http.StatusInternalServerError
	}

	uploadedFile, handler, err := r.FormFile("upload")
	defer uploadedFile.Close()
	if err != nil {
		return err, http.StatusInternalServerError
	}
	log.Printf("Header : %v\n", handler.Header)

	s := session.New()
	os.Getenv("")
	object.PutFileToS3(s)

	// tmpFilePath := "/tmp/" + handler.Filename
	// saveFile, err := os.Create(tmpFilePath)
	// defer saveFile.Close()
	// if err != nil {
	// 	return err, http.StatusInternalServerError
	// }
	//
	// io.Copy(saveFile, uploadedFile)

	//TODO: upload tmp file to s3
	return nil, http.StatusOK
}

func FileUploadHandler(w http.ResponseWriter, r *http.Request) {
	err, code := HandleFileUpload(w, r)
	if err != nil {
		http.Error(w, err.Error(), code)
	}
	http.Redirect(w, r, "/uploaded", http.StatusFound)
}

// HandleShowUploaded : Handle function for uploaded images
func HandleShowUploaded(w http.ResponseWriter) error {
	tmpl, err := GetTemplate("uploaded.tmpl")
	if err != nil {
		return errors.Wrap(err, "Error while getting template")
	}
	data := map[string]interface{}{}
	if err := tmpl.Execute(w, data); err != nil {
		return errors.Wrap(err, "Unable to execute template")
	}
	return nil
}

func ShowUploadedHandler(w http.ResponseWriter, r *http.Request) {
	if err := HandleShowUploaded(w); err != nil {
		log.Fatalln(err)
	}
}
