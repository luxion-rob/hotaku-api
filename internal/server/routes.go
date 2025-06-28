package server

// setupHealthRoutes configures health check routes
func (s *Server) setupHealthRoutes() {
	s.router.GET("/health", s.healthController.HealthCheck)
}

// setupAuthRoutes configures authentication routes
func (s *Server) setupAuthRoutes() {
	auth := s.router.Group("/auth")
	{
		// Public routes (no authentication required)
		auth.POST("/register", s.authController.Register)
		auth.POST("/login", s.authController.Login)

		// Protected routes (authentication required)
		protected := auth.Group("")
		protected.Use(s.authMiddleware)
		{
			protected.GET("/profile", s.authController.Profile)
			protected.PUT("/profile", s.authController.UpdateProfile)
			protected.PUT("/change-password", s.authController.ChangePassword)
		}
	}
}

// setupRoutes configures all application routes
func (s *Server) setupRoutes() {
	// Setup route groups
	s.setupHealthRoutes()
	s.setupAuthRoutes()
}
