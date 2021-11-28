#!/bin/bash

VERSION=$1

RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m'


if [ "$VERSION" == "" ]
then
    echo -e "${RED}usage: ./fetch-libwebp.sh v1.2.0${NC}"
    exit 1
fi

if [[ -n `git status --porcelain` ]]
then
    echo -e "${RED}git status - is not clean:${NC}\n"
    git status --long
    exit 1
fi

if [ -d "internal/source" ]
then
  echo -e "${BLUE}Update libwebp sources to {$VERSION}...${NC}"
  git subtree pull \
    --prefix internal/source \
    --squash \
    --message "update libwebp source to $VERSION" \
    https://chromium.googlesource.com/webm/libwebp $VERSION
else
  echo -e "${BLUE}Add libwebp ${VERSION}...${NC}"
  git subtree add \
    --prefix internal/source \
    --squash \
    --message "add libwebp $VERSION" \
    https://chromium.googlesource.com/webm/libwebp $VERSION
fi

echo -e "${BLUE}Cleaning old links...${NC}"
mkdir -p internal/libwebp
rm -f internal/libwebp/*.{c,h}

echo -e "${BLUE}Create new links...${NC}"
for i in $(find internal/source/src -type f \( -iname \*.h -o -iname \*.c \)); do
  echo "#include \"../${i#internal/}\"" > "internal/libwebp/$(basename -- $i)"
done

echo -e "${BLUE}Add link files to git${NC}"
git add internal/libwebp/*

echo -e "${GREEN}libwebp updated to $VERSION${NC}"
