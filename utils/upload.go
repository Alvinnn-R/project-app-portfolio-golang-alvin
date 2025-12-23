package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// AllowedImageExtensions contains allowed image file extensions
var AllowedImageExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
	".webp": true,
}

// MaxFileSize is the maximum file size (5MB)
const MaxFileSize = 5 * 1024 * 1024

// UploadFile handles file upload and returns the file path
func UploadFile(file multipart.File, header *multipart.FileHeader, uploadDir string) (string, error) {
	// Check file size
	if header.Size > MaxFileSize {
		return "", fmt.Errorf("file size exceeds maximum allowed size (5MB)")
	}

	// Check file extension
	ext := strings.ToLower(filepath.Ext(header.Filename))
	if !AllowedImageExtensions[ext] {
		return "", fmt.Errorf("file type not allowed. Allowed types: jpg, jpeg, png, gif, webp")
	}

	// Create upload directory if not exists
	uploadPath := filepath.Join("public", "assets", uploadDir)
	if err := os.MkdirAll(uploadPath, 0755); err != nil {
		return "", fmt.Errorf("failed to create upload directory: %v", err)
	}

	// Generate unique filename
	filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), sanitizeFilename(header.Filename))
	filePath := filepath.Join(uploadPath, filename)

	// Create destination file
	dst, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %v", err)
	}
	defer dst.Close()

	// Copy file content
	if _, err := io.Copy(dst, file); err != nil {
		return "", fmt.Errorf("failed to save file: %v", err)
	}

	// Return relative path for storage in database (URL format)
	return "/" + filepath.ToSlash(filePath), nil
}

// DeleteFile removes a file from the filesystem
func DeleteFile(filePath string) error {
	if filePath == "" {
		return nil
	}

	// Remove leading slash and convert to OS path
	cleanPath := strings.TrimPrefix(filePath, "/")
	cleanPath = filepath.FromSlash(cleanPath)

	// Check if file exists
	if _, err := os.Stat(cleanPath); os.IsNotExist(err) {
		return nil // File doesn't exist, nothing to delete
	}

	return os.Remove(cleanPath)
}

// sanitizeFilename removes special characters from filename
func sanitizeFilename(filename string) string {
	// Get extension
	ext := filepath.Ext(filename)
	name := strings.TrimSuffix(filename, ext)

	// Replace spaces and special characters
	name = strings.ReplaceAll(name, " ", "_")
	name = strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_' || r == '-' {
			return r
		}
		return '_'
	}, name)

	return name + ext
}
