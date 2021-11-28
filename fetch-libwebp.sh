#!/bin/bash

VERSION=$1

if [ "$VERSION" == "" ]; then
    echo "usage: ./fetch-libwebp.sh v1.2.0"
    exit 1
fi

if [ -d "internal/source" ]
then
  echo "Update libwebp sources to $VERSION..."
  git subtree pull \
    --prefix internal/source \
    --squash \
    --message "update libwebp source to $VERSION" \
    https://chromium.googlesource.com/webm/libwebp $VERSION
else
  echo "Add libwebp $VERSION..."
  git subtree add \
    --prefix internal/source \
    --squash \
    --message "add libwebp $VERSION" \
    https://chromium.googlesource.com/webm/libwebp $VERSION
fi

echo "Cleaning old links..."
mkdir -p internal/libwebp
rm -f internal/libwebp/*.{c,h}

echo "Create new links..."
for i in $(find internal/source/src -type f \( -iname \*.h -o -iname \*.c \)); do
  echo "#include \"../${i#internal/}\"" > "internal/libwebp/$(basename -- $i)"
done

echo "Add link files to git"
git add internal/libwebp/*

echo "libwebp updated to $VERSION"
