#!/bin/bash

# Test MinIO Integration Script
# This script tests the MinIO setup and API endpoints

set -e

echo "🧪 Testing MinIO Integration"
echo "============================"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    if [ "$1" -eq 0 ]; then
        echo -e "${GREEN}✅ $2${NC}"
    else
        echo -e "${RED}❌ $2${NC}"
        exit 1
    fi
}

# Check if MinIO container is running
echo "1. Checking MinIO container status..."
if docker ps | grep -q "hotaku-minio"; then
    print_status 0 "MinIO container is running"
else
    print_status 1 "MinIO container is not running"
fi

# Check if MinIO is accessible
echo "2. Testing MinIO connectivity..."
if curl -s http://localhost:9000/minio/health/live > /dev/null; then
    print_status 0 "MinIO is accessible"
else
    print_status 1 "MinIO is not accessible"
fi

# Check if bucket exists
echo "3. Checking MinIO bucket..."
if docker exec hotaku-minio mc ls myminio/manga-images > /dev/null 2>&1; then
    print_status 0 "Manga-images bucket exists"
else
    echo -e "${YELLOW}⚠️  Manga-images bucket not found, creating it...${NC}"
    if docker exec hotaku-minio mc mb myminio/manga-images --ignore-existing && \
       docker exec hotaku-minio mc policy set download myminio/manga-images; then
        print_status 0 "Manga-images bucket created"
    else
        print_status 1 "Failed to create manga-images bucket"
    fi
fi

# Test API health endpoint
echo "4. Testing API health endpoint..."
if curl -s http://localhost:3000/health | grep -q "status"; then
    print_status 0 "API is running and healthy"
else
    print_status 1 "API is not responding"
fi

# Test authentication (if you have a test user)
echo "5. Testing authentication..."
LOGIN_RESPONSE=$(curl -s -X POST http://localhost:3000/api/v1/auth/login \
    -H "Content-Type: application/json" \
    -d '{"email":"test@example.com","password":"testpassword"}' || echo "{}")

if echo "$LOGIN_RESPONSE" | grep -q "token"; then
    TOKEN=$(echo "$LOGIN_RESPONSE" | grep -o '"token":"[^"]*"' | cut -d'"' -f4)
    print_status 0 "Authentication successful"
    
    # Test upload endpoint (without actual file)
    echo "6. Testing upload endpoint structure..."
        echo -n -e '\x89PNG\r\n\x1a\n\x00\x00\x00\rIHDR\x00\x00\x00\x01\x00\x00\x00\x01\x08\x06\x00\x00\x00\x1f\x15\xc4\x89\x00\x00\x00\rIDATx\x9cc\xf8\x0f\x00\x00\x01\x01\x00\x05\x00\x01\x00\x00\x00\x00IEND\xae\x42\x60\x82' > /tmp/test.png  
        UPLOAD_RESPONSE=$(curl -s -X POST http://localhost:3000/api/v1/upload/manga/test-manga/image \
            -H "Authorization: Bearer $TOKEN" \
            -F "image=@/tmp/test.png" || echo "{}")  
        echo "$UPLOAD_RESPONSE"
        rm -f /tmp/test.png  
    if echo "$UPLOAD_RESPONSE" | grep -q "message"; then
        print_status 0 "Upload endpoint is accessible"
    else
        print_status 1 "Upload endpoint is not accessible"
    fi
else
    echo -e "${YELLOW}⚠️  Authentication test skipped (no test user)${NC}"
fi

echo ""
echo -e "${GREEN}🎉 MinIO Integration Test Completed!${NC}"
echo ""
echo "📋 Next Steps:"
echo "1. Access MinIO Console: http://localhost:9001"
echo "2. Login with: minioadmin / minioadmin"
echo "3. Test file uploads via API endpoints"
echo "4. Check uploaded files in the manga-images bucket"
echo ""
echo "🔗 API Endpoints:"
echo "- POST /api/v1/upload/manga/{manga_id}/image"
echo "- POST /api/v1/upload/manga/{manga_id}/chapters/{chapter_id}/pages"
echo "- DELETE /api/v1/upload/files/{object_name}"
echo "- GET /api/v1/upload/files/{object_name}/info" 