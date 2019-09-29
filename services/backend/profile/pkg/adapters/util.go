// Copyright 2019 kemar
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package adapters

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const maxUploadSize = 20000 * 1024
const uploadPath = ".../../../assets"

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}

func randToken(len int, file multipart.File) string {
	a := sha1.New()
	io.Copy(a, file)
	b := make([]byte, len)
	rand.Read(b)
	ve := a.Sum(nil)
	val := append(ve[:], b[:]...)
	return fmt.Sprintf("%x", val)
}

// FileUpload ...
func FileUpload(w http.ResponseWriter, r *http.Request) (string, error) {

	// validate file size
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		renderError(w, "FILE_TOO_BIG", http.StatusBadRequest)
		return "", nil
	}

	// parse and validate file and profile parameters
	fileType := r.FormValue("type")
	file, fh, err := r.FormFile("contentPhoto")
	if err != nil {
		renderError(w, "INVALID_FILE", http.StatusBadRequest)
		return "", nil
	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		renderError(w, "INVALID_FILE", http.StatusBadRequest)
		return "", nil
	}

	ext := strings.Split(fh.Filename, ".")[1]

	// check file type, detectcontenttype only needs the first 512 bytes
	filetype := http.DetectContentType(fileBytes)
	fmt.Println("filetype: ", ext)
	switch filetype {
	case "image/jpeg", "image/jpg":
	case "image/gif", "image/png":
	case "application/pdf":
		break
	default:
		renderError(w, "INVALID_FILE_TYPE", http.StatusBadRequest)
		return "", nil
	}
	fileName := randToken(12, file)

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fileName = fileName + "." + ext
	newPath := filepath.Join(wd+uploadPath, "photos", fileName)
	fmt.Printf("FileType: %s, File: %s\n", fileType, newPath)

	// write file
	newFile, err := os.Create(newPath)
	if err != nil {
		renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
		return "", nil
	}
	defer newFile.Close() // idempotent, okay to call twice
	if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
		renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
		return "", nil
	}
	w.Write([]byte("SUCCESS"))

	return fileName, nil
}
