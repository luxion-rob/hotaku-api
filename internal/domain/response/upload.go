package response

// UploadResponse represents the response for file uploads
type UploadResponse struct {
	URL      string `json:"url"`
	Filename string `json:"filename"`
	Size     int64  `json:"size"`
}

// FileInfoResponse represents file information
type FileInfoResponse struct {
	ObjectName string `json:"object_name"`
	Size       int64  `json:"size"`
}
