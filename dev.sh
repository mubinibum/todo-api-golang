#!/bin/bash

# Warna
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${GREEN}🔧 Menjalankan go mod tidy...${NC}"
go mod tidy

echo -e "${GREEN}🧪 Menjalankan unit test...${NC}"
go test -v ./...

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✅ Semua test lulus.${NC}"
    
    echo -e "${GREEN}📦 Menambahkan perubahan ke Git...${NC}"
    git add .

    echo -e "${GREEN}📝 Silakan masukkan pesan commit:${NC}"
    read commit_msg
    git commit -m "$commit_msg"

    echo -e "${GREEN}🚀 Commit selesai. Siap untuk push.${NC}"
else
    echo -e "${RED}❌ Test gagal! Commit dibatalkan.${NC}"
fi
