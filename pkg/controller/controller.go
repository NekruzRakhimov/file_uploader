package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("name")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	timeSign := fmt.Sprintf("%d", time.Now().UnixNano())

	filePath := fmt.Sprintf("%s_%s", timeSign, file.Filename)
	filePath = strings.Replace(filePath, " ", "", 111)

	if err = c.SaveUploadedFile(file, fmt.Sprintf("./files/layouts/%s", filePath)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"file_path": filePath})
}

func DownloadFile(c *gin.Context) {
	filePath := c.Query("file_path")

	filePath = fmt.Sprintf("./files/layouts/%s", filePath)
	fmt.Println("<filepath>: ", filePath)
	//c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")

	contentType, err := GetFileContentType(filePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	c.Writer.Header().Set("Content-Type", contentType)
	c.File(filePath)
	//c.File("files/contracts/edit-document.docx")
}

func GetFileContentType(filePath string) (contentType string, err error) {
	if strings.HasSuffix(filePath, "pdf") {
		return "application/pdf", nil
	} else if strings.HasSuffix(filePath, "docx") {
		return "application/vnd.openxmlformats-officedocument.wordprocessingml.document", nil
	} else if strings.HasSuffix(filePath, "xlsx") {
		return "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", nil
	} else if strings.HasSuffix(filePath, "xls") {
		return "application/vnd.ms-excel", nil
	} else if strings.HasSuffix(filePath, "txt") {
		return "text/plain", nil
	} else if strings.HasSuffix(filePath, "png") {
		return "image/png", nil
	} else if strings.HasSuffix(filePath, "jpeg") {
		return "image/jpeg", nil
	} else if strings.HasSuffix(filePath, "jpg") {
		return "image/jpg", nil
	} else {
		return "", errors.New("file extension not found")
	}
}
