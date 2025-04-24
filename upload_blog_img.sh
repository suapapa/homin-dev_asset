#!/bin/bash

set -e

files=$@

for file in $files; do
    file=$(echo "$file" | xargs)  # trim space

    if [[ $file == *.webp ]]; then
        dest_file=./asset/blog/img/$(basename $file)
    else
        dest_file=./asset/blog/img/$(basename $file).webp
    fi
    echo "$file -> $dest_file"

    magick $file -resize 640x640 $dest_file
    git add $dest_file

    # trim prefix ./asset/
    path=${dest_file#./asset/}
    echo "- ![img-00](https://homin.dev/asset/$path)"
done

git commit -m "update blog img"
git push

./update_blog_img.sh latest