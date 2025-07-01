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
}
