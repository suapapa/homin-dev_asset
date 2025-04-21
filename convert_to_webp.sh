#!/bin/bash

set -e

# 변환 품질 설정 (원하는대로 조정 가능)
QUALITY=80

# PNG 파일 처리
find . -type f -iname "*.png" | while read -r file; do
  output="${file%.png}.webp"
  echo "Converting $file -> $output"
  cwebp -q $QUALITY "$file" -o "$output"
  rm $file
done

# JPG 및 JPEG 파일 처리
find . -type f \( -iname "*.jpg" -o -iname "*.jpeg" \) | while read -r file; do
  output="${file%.*}.webp"
  echo "Converting $file -> $output"
  cwebp -q $QUALITY "$file" -o "$output"
  rm $file
done
