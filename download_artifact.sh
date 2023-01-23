#!/usr/bin/env sh

TOKEN_FILE=.blog.github.token
TOKEN=$(cat "${TOKEN_FILE}")

ARCHIVE_NAME="blog-sewera-dev-dist-archive"
ARCHIVE_ZIP="${ARCHIVE_NAME}.zip"
ARCHIVE="dist.tar"

DL_URL=$(curl \
  -H "Accept: application/vnd.github+json" \
  -H "Authorization: Bearer ${TOKEN}" \
  -H "X-GitHub-Api-Version: 2022-11-28" \
  "https://api.github.com/repos/blazejsewera/blog/actions/artifacts?per_page=1&name=${ARCHIVE_NAME}" 2>/dev/null \
| jq -r '.artifacts[0].archive_download_url')

curl \
  -L \
  -H "Accept: application/vnd.github+json" \
  -H "Authorization: Bearer ${TOKEN}" \
  -H "X-GitHub-Api-Version: 2022-11-28" \
  -o "blog-sewera-dev-dist-archive.zip" \
  "${DL_URL}" 2>/dev/null

tar -xzf "${ARCHIVE_ZIP}" && tar -xf "${ARCHIVE}"

rm -f "${ARCHIVE_ZIP}" "${ARCHIVE}"
