#!/usr/bin/env bash

git config -f .gitmodules --get-regexp '^submodule\.' | grep '\.path ' | while read -r key path; do
    # Extract submodule name from the key (everything between 'submodule.' and '.path')
    submodule_name=$(echo "$key" | sed 's/^submodule\.\(.*\)\.path$/\1/')

    url=$(git config -f .gitmodules --get "submodule.$submodule_name.url")
    branch=$(git config -f .gitmodules --get "submodule.$submodule_name.branch" 2>/dev/null)

    if [ -n "$branch" ]; then
        echo "running git submodule add for $url at $path with branch $branch"
    else
        echo "running git submodule add for $url at $path (no branch specified)"
    fi
done
