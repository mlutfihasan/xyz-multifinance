package helper

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

func UploadFilePhoto(file multipart.File, handlerFile multipart.FileHeader, CustomerNik, fileType string) (string, error) {
	if file != nil {
		defer file.Close()

		var err = os.MkdirAll("./"+FileDirectory+"/"+FileDirectoryCustomer+"/"+fileType, 0700)
		if err != nil {
			return "", err
		}

		s := strings.Split(handlerFile.Filename, ".")
		sExtFile := s[len(s)-1]

		sLinkFile := "/" + FileDirectoryCustomer + "/" + fileType + "/" + CustomerNik + "-" + fileType + "." + sExtFile

		f, err := os.OpenFile("./"+FileDirectory+"/"+FileDirectoryCustomer+"/"+fileType+"/"+CustomerNik+"-"+fileType+"."+sExtFile, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			return "", err
		}
		defer f.Close()

		io.Copy(f, file)

		return sLinkFile, nil
	}
	return "", errors.New("Upload File " + fileType + " Failed")
}
