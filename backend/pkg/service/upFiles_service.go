package service

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/takwot/tech-strelka.git/pkg/database"
)

type UploadFilesService struct {
	repo database.Repository
}

func NewAlbumService(repo database.Repository) *UploadFilesService {
	return &UploadFilesService{repo: repo}
}

func (s *UploadFilesService) UploadMultipleFile(files []*multipart.FileHeader) ([]string, error) {

	filePaths := []string{}
	for _, file := range files {
		fileExt := filepath.Ext(file.Filename)
		originalFileName := strings.TrimSuffix(filepath.Base(file.Filename), filepath.Ext(file.Filename))
		now := time.Now()
		filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt
		// filePath := "http://localhost:5000/images/multiple/" + filename

		filePaths = append(filePaths, filename)
		out, err := os.Create("./public/multiple/" + filename)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		readerFile, _ := file.Open()
		_, err = io.Copy(out, readerFile)
		if err != nil {
			log.Fatal(err)
		}
	}

	return filePaths, nil

}
