package utils

import (
	"errors"
	"mime/multipart"
	"net/http"
	"strings"
)

func GetImageFormat(file multipart.File) (string, error) {
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer)
	switch contentType {
	case "image/jpeg":
		return "jpeg", nil
	case "image/png":
		return "png", nil
	case "image/gif":
		return "gif", nil
	case "image/tiff":
		return "tiff", nil
	case "image/webp":
		return "webp", nil
	default:
		return "", errors.New("Invalid image format")
	}
}

func GetImageFormatFromURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	buffer := make([]byte, 512)
	_, err = resp.Body.Read(buffer)
	if err != nil {
		return "", err
	}
	return http.DetectContentType(buffer), nil
}

func DeleteSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}
