package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	"hotaku-api/config"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// MinIOService handles MinIO operations
type MinIOService struct {
	client     *minio.Client
	bucketName string
}

// NewMinIOService creates a new MinIO service instance
func NewMinIOService(cfg *config.Config) (*MinIOService, error) {
	// Initialize MinIO client
	minioClient, err := minio.New(cfg.MinIO.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.MinIO.AccessKeyID, cfg.MinIO.SecretAccessKey, ""),
		Secure: cfg.MinIO.UseSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create MinIO client: %w", err)
	}

	service := &MinIOService{
		client:     minioClient,
		bucketName: cfg.MinIO.BucketName,
	}

	// Ensure bucket exists
	if err := service.ensureBucketExists(true); err != nil {
		return nil, fmt.Errorf("failed to ensure bucket exists: %w", err)
	}

	return service, nil
}

// ensureBucketExists creates the bucket if it doesn't exist
func (s *MinIOService) ensureBucketExists(isPublic bool) error {
	exists, err := s.client.BucketExists(context.Background(), s.bucketName)
	if err != nil {
		return err
	}

	if !exists {
		err = s.client.MakeBucket(context.Background(), s.bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}

		if isPublic {
			// Set bucket policy for public read access (optional)
			policy := `{
				"Version": "2012-10-17",
				"Statement": [
					{
						"Effect": "Allow",
						"Principal": {"AWS": "*"},
						"Action": ["s3:GetObject"],
						"Resource": ["arn:aws:s3:::` + s.bucketName + `/*"]
					}
				]
			}`

			err = s.client.SetBucketPolicy(context.Background(), s.bucketName, policy)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *MinIOService) constructFileURL(filename string) string {
	scheme := "http"
	if s.client.EndpointURL().Scheme == "https" {
		scheme = "https"
	}
	return fmt.Sprintf("%s://%s/%s/%s", scheme, s.client.EndpointURL().Host, s.bucketName, filename)
}

// UploadMangaImage uploads a manga image file to MinIO
func (s *MinIOService) UploadMangaImage(file *multipart.FileHeader, mangaID string) (string, error) {
	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer src.Close()

	// Generate unique filename
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("manga/%s/%s%s", mangaID, uuid.New().String(), ext)

	// Get content type
	contentType := file.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	// Upload to MinIO
	_, err = s.client.PutObject(
		context.Background(),
		s.bucketName,
		filename,
		src,
		file.Size,
		minio.PutObjectOptions{
			ContentType: contentType,
		},
	)
	if err != nil {
		return "", fmt.Errorf("failed to upload file to MinIO: %w", err)
	}

	return s.constructFileURL(filename), nil
}

// UploadChapterPage uploads a chapter page image to MinIO
func (s *MinIOService) UploadChapterPage(file *multipart.FileHeader, mangaID, chapterID string, pageNumber int) (string, error) {
	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer src.Close()

	// Generate unique filename
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("manga/%s/chapters/%s/page_%03d%s", mangaID, chapterID, pageNumber, ext)

	// Get content type
	contentType := file.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	// Upload to MinIO
	_, err = s.client.PutObject(
		context.Background(),
		s.bucketName,
		filename,
		src,
		file.Size,
		minio.PutObjectOptions{
			ContentType: contentType,
		},
	)
	if err != nil {
		return "", fmt.Errorf("failed to upload file to MinIO: %w", err)
	}

	return s.constructFileURL(filename), nil
}

// DeleteFile deletes a file from MinIO
func (s *MinIOService) DeleteFile(objectName string) error {
	err := s.client.RemoveObject(context.Background(), s.bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete file from MinIO: %w", err)
	}
	return nil
}

// GetFileURL generates a presigned URL for file access
func (s *MinIOService) GetFileURL(objectName string, expiry time.Duration) (string, error) {
	url, err := s.client.PresignedGetObject(context.Background(), s.bucketName, objectName, expiry, nil)
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL: %w", err)
	}
	return url.String(), nil
}

// ListFiles lists files in a directory
func (s *MinIOService) ListFiles(prefix string) ([]string, error) {
	var files []string

	opts := minio.ListObjectsOptions{
		Prefix:    prefix,
		Recursive: true,
	}

	for obj := range s.client.ListObjects(context.Background(), s.bucketName, opts) {
		if obj.Err != nil {
			return nil, fmt.Errorf("error listing objects: %w", obj.Err)
		}
		files = append(files, obj.Key)
	}

	return files, nil
}

// GetFileSize gets the size of a file
func (s *MinIOService) GetFileSize(objectName string) (int64, error) {
	objInfo, err := s.client.StatObject(context.Background(), s.bucketName, objectName, minio.StatObjectOptions{})
	if err != nil {
		return 0, fmt.Errorf("failed to get file info: %w", err)
	}
	return objInfo.Size, nil
}

// CopyFile copies a file within the same bucket
func (s *MinIOService) CopyFile(srcObject, dstObject string) error {
	src := minio.CopySrcOptions{
		Bucket: s.bucketName,
		Object: srcObject,
	}
	dst := minio.CopyDestOptions{
		Bucket: s.bucketName,
		Object: dstObject,
	}

	_, err := s.client.CopyObject(context.Background(), dst, src)
	if err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}

	return nil
}

// GetObject retrieves an object from MinIO
func (s *MinIOService) GetObject(objectName string) (*minio.Object, error) {
	obj, err := s.client.GetObject(context.Background(), s.bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get object: %w", err)
	}
	return obj, nil
}
