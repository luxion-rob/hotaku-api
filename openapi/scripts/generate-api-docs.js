#!/usr/bin/env node

const fs = require('node:fs');
const path = require('node:path');
const yaml = require('js-yaml');

// Create openapi directory if it doesn't exist (parent directory)  
const openapiDir = path.join(__dirname, '..');  
if (!fs.existsSync(openapiDir)) {  
  fs.mkdirSync(openapiDir, { recursive: true });  
}  
  
// Create dist directory if it doesn't exist (for generated output)  
const distDir = path.join(openapiDir, 'dist');  
if (!fs.existsSync(distDir)) {  
  fs.mkdirSync(distDir, { recursive: true });  
}

// OpenAPI specification
const openApiSpec = {
  openapi: '3.0.3',
  info: {
    title: 'Hotaku API',
    description: 'Go Gin REST API Documentation',
    version: '1.0.0',
    contact: {
      name: 'API Support',
      email: 'support@hotaku.com'
    },
    license: {
      name: 'MIT',
      url: 'https://opensource.org/licenses/MIT'
    }
  },
  servers: [
    {
      url: 'https://api.hotaku.com',
      description: 'Production server'
    },
    {
      url: 'https://staging-api.hotaku.com',
      description: 'Staging server'
    }
  ],
  paths: {
    '/health': {
      get: {
        operationId: 'getHealthCheck',
        summary: 'Health Check',
        description: 'Returns API health status',
        tags: ['Health'],
        security: [],
        responses: {
          '200': {
            description: 'API is healthy',
            content: {
              'application/json': {
                schema: {
                  type: 'object',
                  properties: {
                    status: {
                      type: 'string',
                      example: 'healthy'
                    },
                    message: {
                      type: 'string',
                      example: 'API is running smoothly'
                    },
                    timestamp: {
                      type: 'integer',
                      example: 1640995200
                    },
                    version: {
                      type: 'string',
                      example: '1.0.0'
                    }
                  },
                  required: ['status', 'message', 'timestamp', 'version']
                }
              }
            }
          },
          '500': {
            description: 'Internal server error',
            content: {
              'application/json': {
                schema: {
                  $ref: '#/components/schemas/Error'
                }
              }
            }
          }
        }
      }
    },
    '/auth/login': {
      post: {
        operationId: 'loginUser',
        summary: 'User Login',
        description: 'Authenticate user and return JWT token',
        tags: ['Authentication'],
        security: [
          {
            bearerAuth: []
          }
        ],
        requestBody: {
          required: true,
          content: {
            'application/json': {
              schema: {
                type: 'object',
                properties: {
                  email: {
                    type: 'string',
                    format: 'email',
                    example: 'user@example.com'
                  },
                  password: {
                    type: 'string',
                    format: 'password',
                    example: 'password123'
                  }
                },
                required: ['email', 'password']
              }
            }
          }
        },
        responses: {
          '200': {
            description: 'Login successful',
            content: {
              'application/json': {
                schema: {
                  type: 'object',
                  properties: {
                    token: {
                      type: 'string',
                      example: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...'
                    },
                    user: {
                      $ref: '#/components/schemas/User'
                    }
                  }
                }
              }
            }
          },
          '400': {
            description: 'Bad request - Invalid input data',
            content: {
              'application/json': {
                schema: {
                  $ref: '#/components/schemas/Error'
                }
              }
            }
          },
          '401': {
            description: 'Invalid credentials',
            content: {
              'application/json': {
                schema: {
                  $ref: '#/components/schemas/Error'
                }
              }
            }
          },
          '422': {
            description: 'Validation error',
            content: {
              'application/json': {
                schema: {
                  $ref: '#/components/schemas/Error'
                }
              }
            }
          },
          '500': {
            description: 'Internal server error',
            content: {
              'application/json': {
                schema: {
                  $ref: '#/components/schemas/Error'
                }
              }
            }
          }
        }
      }
    },
    '/auth/register': {
      post: {
        operationId: 'registerUser',
        summary: 'User Registration',
        description: 'Register a new user account',
        tags: ['Authentication'],
        security: [
          {
            bearerAuth: []
          }
        ],
        requestBody: {
          required: true,
          content: {
            'application/json': {
              schema: {
                type: 'object',
                properties: {
                  name: {
                    type: 'string',
                    example: 'John Doe'
                  },
                  email: {
                    type: 'string',
                    format: 'email',
                    example: 'user@example.com'
                  },
                  password: {
                    type: 'string',
                    format: 'password',
                    example: 'password123'
                  }
                },
                required: ['name', 'email', 'password']
              }
            }
          }
        },
        responses: {
          '201': {
            description: 'User created successfully',
            content: {
              'application/json': {
                schema: {
                  type: 'object',
                  properties: {
                    message: {
                      type: 'string',
                      example: 'User created successfully'
                    },
                    user: {
                      $ref: '#/components/schemas/User'
                    }
                  }
                }
              }
            }
          },
          '400': {
            description: 'Bad request - Invalid input data',
            content: {
              'application/json': {
                schema: {
                  $ref: '#/components/schemas/Error'
                }
              }
            }
          },
          '409': {
            description: 'Conflict - User already exists',
            content: {
              'application/json': {
                schema: {
                  $ref: '#/components/schemas/Error'
                }
              }
            }
          },
          '422': {
            description: 'Validation error',
            content: {
              'application/json': {
                schema: {
                  $ref: '#/components/schemas/Error'
                }
              }
            }
          },
          '500': {
            description: 'Internal server error',
            content: {
              'application/json': {
                schema: {
                  $ref: '#/components/schemas/Error'
                }
              }
            }
          }
        }
      }
    }
  },
  components: {
    schemas: {
      User: {
        type: 'object',
        properties: {
          id: {
            type: 'integer',
            example: 1
          },
          name: {
            type: 'string',
            example: 'John Doe'
          },
          email: {
            type: 'string',
            format: 'email',
            example: 'user@example.com'
          },
          created_at: {
            type: 'string',
            format: 'date-time',
            example: '2024-01-01T00:00:00Z'
          },
          updated_at: {
            type: 'string',
            format: 'date-time',
            example: '2024-01-01T00:00:00Z'
          }
        },
        required: ['id', 'name', 'email']
      },
      Error: {
        type: 'object',
        properties: {
          error: {
            type: 'string',
            example: 'Bad Request'
          },
          message: {
            type: 'string',
            example: 'Invalid input data'
          },
          details: {
            type: 'object',
            additionalProperties: true
          }
        },
        required: ['error', 'message']
      }
    },
    securitySchemes: {
      bearerAuth: {
        type: 'http',
        scheme: 'bearer',
        bearerFormat: 'JWT'
      }
    }
  },
  security: [
    {
      bearerAuth: []
    }
  ],
  tags: [
    {
      name: 'Health',
      description: 'Health check endpoints'
    },
    {
      name: 'Authentication',
      description: 'Authentication endpoints'
    }
  ]
};

// Write the OpenAPI specification to file
const outputPath = path.join(distDir, 'openapi.yaml');
fs.writeFileSync(outputPath, yaml.dump(openApiSpec));

console.log(`✅ OpenAPI specification generated at: ${outputPath}`);

// Also generate a JSON version for compatibility
const jsonOutputPath = path.join(distDir, 'openapi.json');
fs.writeFileSync(jsonOutputPath, JSON.stringify(openApiSpec, null, 2));

console.log(`✅ OpenAPI JSON specification generated at: ${jsonOutputPath}`); 