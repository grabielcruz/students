package handle_uploads

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"example.com/students/models"
	"example.com/students/utils"
	"github.com/julienschmidt/httprouter"
)

func UploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	file, header, err := r.FormFile("image")
	if err != nil {
		utils.SendInternalServerError(w, err)
		return
	}
	defer file.Close()
	fmt.Println(header.Filename)

	validImage := false
	extension := strings.ToLower(filepath.Ext(header.Filename))
	for _, v := range models.ImageTypes {
		if extension == v {
			validImage = true
			break
		}
	}
	if !validImage {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "El archivo del tipo %v no es una imagen reconocida", extension)
		return
	}

	name := fmt.Sprintf("*%v", extension)

	tempFile, err := ioutil.TempFile("public/photos", name)
	if err != nil {
		utils.SendInternalServerError(w, err)
		return
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, file)
	if err != nil {
		utils.SendInternalServerError(w, err)
		return
	}

	fileName := tempFile.Name()
	fmt.Fprint(w, fileName)
}