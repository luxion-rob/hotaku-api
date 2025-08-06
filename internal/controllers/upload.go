package controllers

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"hotaku-api/internal/domain/response"
	"hotaku-api/internal/service"

	"github.com/gin-gonic/gin"
)

const (
	MaxFileSize = 10 * 1024 * 1024 // 10MB
)

// UploadController handles file upload operations
type UploadController struct {
	minioService *service.MinIOService
}

// NewUploadController creates a new upload controller
func NewUploadController(minioService *service.MinIOService) *UploadController {
	return &UploadController{
		minioService: minioService,
	}
}

// isValidImageFile checks if the file is a valid image
func isValidImageFile(filename string) bool {
	validExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".webp": true,
	}

	ext := strings.ToLower(filepath.Ext(filename))
	return validExtensions[ext]
}

// getImageContentType determines the MIME type based on file extension
func getImageContentType(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	case ".webp":
		return "image/webp"
	case ".svg":
		return "image/svg+xml"
	case ".bmp":
		return "image/bmp"
	case ".ico":
		return "image/x-icon"
	default:
		return "application/octet-stream"
	}
}

func (c *UploadController) validateImageFile(file *multipart.FileHeader) error {
	if !isValidImageFile(file.Filename) {
		return fmt.Errorf("invalid file type. Only image files (jpg, jpeg, png, gif, webp) are allowed")
	}

	if file.Size > MaxFileSize {
		return fmt.Errorf("file size %d exceeds maximum allowed size %d", file.Size, MaxFileSize)
	}
	return nil
}

// UploadMangaImage handles manga image upload
func (c *UploadController) UploadMangaImage(ctx *gin.Context) {
	mangaID := ctx.Param("manga_id")
	if mangaID == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Manga ID is required", nil))
		return
	}

	// Get the uploaded file
	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Failed to get uploaded file", nil))
		return
	}

	if err := c.validateImageFile(file); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error(), nil))
		return
	}

	// Upload to MinIO
	fileURL, err := c.minioService.UploadMangaImage(file, mangaID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Failed to upload file", nil))
		return
	}

	ctx.JSON(http.StatusOK, response.UploadResponse{
		URL:      fileURL,
		Filename: file.Filename,
		Size:     file.Size,
	})
}

// UploadChapterPages handles chapter pages upload
func (c *UploadController) UploadChapterPages(ctx *gin.Context) {
	mangaID := ctx.Param("manga_id")
	chapterID := ctx.Param("chapter_id")

	if mangaID == "" || chapterID == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Manga ID and Chapter ID are required", nil))
		return
	}

	// Get the uploaded files
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Failed to get uploaded files", nil))
		return
	}

	files := form.File["pages"]
	if len(files) == 0 {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "No files uploaded", nil))
		return
	}

	// List existing files in chapter to determine max page number
	prefix := fmt.Sprintf("manga/%s/chapters/%s/", mangaID, chapterID)
	existingFiles, err := c.minioService.ListFiles(prefix)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Failed to list existing files", nil))
		return
	}

	// Find max existing page number
	maxPage := 0
	_, fileName := filepath.Split(existingFiles[len(existingFiles)-1])
	baseName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	var pageNum int
	if n, err := fmt.Sscanf(baseName, "page_%03d", &pageNum); err == nil && n == 1 && pageNum > maxPage {
		maxPage = pageNum
	}

	var uploadResponses []response.UploadResponse

	// Upload each new file, continuing the page numbering
	for i, file := range files {
		pageNumber := maxPage + i + 1

		if err := c.validateImageFile(file); err != nil {
			ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error(), nil))
			return
		}

		fileURL, err := c.minioService.UploadChapterPage(file, mangaID, chapterID, pageNumber)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Failed to upload file "+file.Filename, nil))
			return
		}

		uploadResponses = append(uploadResponses, response.UploadResponse{
			URL:      fileURL,
			Filename: file.Filename,
			Size:     file.Size,
		})
	}

	ctx.JSON(http.StatusOK, uploadResponses)
}

// ReplacePage handles replace specific page
func (c *UploadController) ReplacePage(ctx *gin.Context) {
	mangaID := ctx.Param("manga_id")
	chapterID := ctx.Param("chapter_id")

	if mangaID == "" || chapterID == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Manga ID and Chapter ID are required", nil))
		return
	}

	page, err := strconv.Atoi(ctx.Param("page"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Page must be integer", nil))
		return
	}

	if page <= 0 {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Page must be a positive integer", nil))
		return
	}

	if page > 999 {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Page number exceeds maximum allowed value", nil))
		return
	}

	// Get the uploaded file
	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "No image file found in the request. Key 'image' required.", nil))
		return
	}

	if err := c.validateImageFile(file); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error(), nil))
		return
	}

	// Upload to MinIO
	fileURL, err := c.minioService.UploadChapterPage(file, mangaID, chapterID, page)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Failed to upload file", nil))
		return
	}

	ctx.JSON(http.StatusOK, response.UploadResponse{
		URL:      fileURL,
		Filename: file.Filename,
		Size:     file.Size,
	})
}

// DeleteFile handles file deletion
func (c *UploadController) DeleteFile(ctx *gin.Context) {
	objectName := ctx.Param("object_name")
	if objectName == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Object name is required", nil))
		return
	}

	err := c.minioService.DeleteFile(objectName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Failed to delete file", nil))
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetFileInfo gets file information
func (c *UploadController) GetFileInfo(ctx *gin.Context) {
	objectName := ctx.Param("object_name")
	if objectName == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Object name is required", nil))
		return
	}

	// Remove /info suffix to get the actual object name
	objectName = strings.TrimSuffix(objectName, "/info")

	// Trim leading slash for wildcard parameter
	if len(objectName) > 0 && objectName[0] == '/' {
		objectName = objectName[1:]
	}

	size, err := c.minioService.GetFileSize(objectName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Failed to get file info", nil))
		return
	}

	ctx.JSON(http.StatusOK, response.FileInfoResponse{
		ObjectName: objectName,
		Size:       size,
	})
}

// GetImage retrieves and serves an image file
func (c *UploadController) GetImage(ctx *gin.Context) {
	objectName := ctx.Param("object_name")
	if objectName == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Object name is required", nil))
		return
	}

	// Validate that the requested file is an image
	if !isValidImageFile(objectName) {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid file type. Only image files (jpg, jpeg, png, gif, webp) are allowed", nil))
		return
	}

	// Get the object from MinIO
	obj, err := c.minioService.GetObject(objectName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Failed to retrieve image", nil))
		return
	}
	defer obj.Close()

	// Get object info for content type and size
	objInfo, err := obj.Stat()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Failed to get image info", nil))
		return
	}

	// Determine content type based on file extension
	contentType := objInfo.ContentType
	if contentType == "application/octet-stream" {
		contentType = getImageContentType(objectName)
	}

	// Set appropriate headers
	ctx.Header("Content-Type", contentType)
	ctx.Header("Content-Length", fmt.Sprintf("%d", objInfo.Size))
	ctx.Header("Cache-Control", "public, max-age=31536000") // Cache for 1 year
	ctx.Header("Access-Control-Allow-Origin", "*")

	// Stream the file to the response
	ctx.DataFromReader(http.StatusOK, objInfo.Size, contentType, obj, nil)
}
