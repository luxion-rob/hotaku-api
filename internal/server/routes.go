package server

// setupRoutes configures all application routes
func (s *Server) setupRoutes() {
	// Setup route groups
	s.router.GET("/health", s.healthController.HealthCheck)

	// Setup auth routes
	auth := s.router.Group("/api/v1/auth")
	{
		auth.POST("/register", s.authController.Register)
		auth.POST("/login", s.authController.Login)

		protected := auth.Group("")
		protected.Use(s.authMiddleware)
		{
			protected.GET("/profile", s.authController.Profile)
			protected.PUT("/profile", s.authController.UpdateProfile)
			protected.PUT("/change-password", s.authController.ChangePassword)
		}
	}

	// Setup author routes
	authors := s.router.Group("/api/v1/authors")
	{
		// Public read access
		authors.GET("/:author_id", s.authorController.GetAuthor)

		// Protected write operations
		protected := authors.Group("")
		protected.Use(s.authMiddleware)
		{
			protected.POST("", s.authorController.CreateAuthor)
			protected.PUT("/:author_id", s.authorController.UpdateAuthor)
			protected.DELETE("/:author_id", s.authorController.DeleteAuthor)
		}
	}

	// Setup upload routes
	upload := s.router.Group("/api/v1/upload")
	upload.Use(s.authMiddleware) // Require authentication for uploads
	{
		upload.POST("/manga/:manga_id/image", s.uploadController.UploadMangaImage)
		upload.POST("/manga/:manga_id/chapters/:chapter_id/pages", s.uploadController.UploadChapterPages)
		upload.PUT("/manga/:manga_id/chapters/:chapter_id/pages/:page", s.uploadController.ReplacePage)
		upload.DELETE("/files/*object_name", s.uploadController.DeleteFile)
		upload.GET("/files/*object_name", s.uploadController.GetFileInfo)
	}

	// Setup public image routes (no authentication required)
	images := s.router.Group("/api/v1/images")
	{
		images.GET("/*object_name", s.uploadController.GetImage)
	}
}
