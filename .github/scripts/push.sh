#!/bin/sh

set -e

# Check if any files were changed, if not we don't need to do anything
status=$(git status --porcelain | grep "^ M" | wc -l)
if [ $status = 0 ]; then
  exit 0
fi

git config user.email "git@estrato.co.uk"
git config user.name "Estrato GitHub Bot"

branch=$(jq --raw-output .pull_request.head.ref "$GITHUB_EVENT_PATH")
git checkout $branch
git commit -m"$1" .
git push
