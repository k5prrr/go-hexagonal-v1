# sh scripts/createFileForGPT.sh

OUTPUT_FILE=~/projectForGPT.txt

echo "Дерево файлов:" >> "$OUTPUT_FILE"
tree >> "$OUTPUT_FILE"

find . -type d -name .git -prune -o -type f -exec sh -c 'echo "Содержимое файла: $1:"; cat "$1"' _ {} \; >> "$OUTPUT_FILE"


