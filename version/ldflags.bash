#!/bin/bash

set -eo pipefail

# Check if the current commit has a tag
tag=$(git describe --exact-match --tags 2>/dev/null || true)

if [ -n "$tag" ]; then
  # Current commit is tagged, so print the tag
  echo -n " -X github.com/sarvsav/iza/version.tag=$tag "
else
  # No exact tag on this commit; get the latest tag and the current commit hash
  last_tag=$(git describe --tags --abbrev=0)
  commit_hash=$(git rev-parse --short HEAD)
  echo -n " -X github.com/sarvsav/iza/version.tag=$last_tag-$commit_hash "
fi

# Add commit hash and date information
echo -n $(git show -s --format=' -X github.com/sarvsav/iza/version.commit=%H -X github.com/sarvsav/iza/version.date=%ct')

# Check for dirty state
git diff --quiet HEAD || echo ' -X github.com/sarvsav/iza/version.dirty=dirty'
