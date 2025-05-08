#!/bin/bash

# Fetch the latest information about remote branches
git fetch --prune

# Iterate over local branches
for branch in $(git branch --format '%(refname:short)'); do
    # Check if the branch exists on the remote
    if ! git show-ref --verify --quiet "refs/remotes/origin/$branch"; then
        # Delete the local branch if it doesn't exist on the remote
        git branch -d "$branch"
        echo "Deleted local branch: $branch"
    fi
done
