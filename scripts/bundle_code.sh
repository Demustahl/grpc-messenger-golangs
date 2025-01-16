#!/bin/bash

# Файл, куда будет записан весь код
OUTPUT_FILE="project_code_bundle.txt"

# Очистим файл, если он уже существует
> "$OUTPUT_FILE"

# Рекурсивно пройти по каталогу и обработать все файлы
# find . -type f -name "*.go" | while read -r file; do
# find . -type f \( -name "*.go" -o -name "*.proto" -o -name "*.html" -o -name "*.css" -o -name "*.js" \) | while read -r file; do
find . -type f -name "*.go" \
    -not -path "./api/generated/*" | while read -r file; do
    # Записать название файла и его путь
    echo "========== $file ==========" >> "$OUTPUT_FILE"
    echo "" >> "$OUTPUT_FILE"
    
    # Записать содержимое файла
    cat "$file" >> "$OUTPUT_FILE"
    echo "" >> "$OUTPUT_FILE"
    echo "" >> "$OUTPUT_FILE"
done

echo "Сборка завершена! Код записан в файл $OUTPUT_FILE"
