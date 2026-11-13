#!/usr/bin/env bash
# Publishes the packages built by npm/build.mjs (platform packages first, launcher last).
# Already published versions are skipped, so re-running after a partial failure is safe.
#   npm/publish.sh [dist-dir]      (extra npm publish flags via NPM_PUBLISH_FLAGS, e.g. "--provenance")
set -euo pipefail
dist="$(cd "${1:-$(dirname "$0")/dist}" && pwd)"
while read -r d; do
  [ -n "$d" ] || continue
  name=$(node -p "require('$dist/$d/package.json').name")
  ver=$(node -p "require('$dist/$d/package.json').version")
  if npm view "$name@$ver" version >/dev/null 2>&1; then
    echo "skip $name@$ver (already published)"
    continue
  fi
  # shellcheck disable=SC2086
  (cd "$dist/$d" && npm publish --access public ${NPM_PUBLISH_FLAGS:-})
  echo "published $name@$ver"
done < "$dist/PUBLISH_ORDER"
