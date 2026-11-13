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

# Verify: every package version must be visible on the public registry, otherwise fail the job
# (a new version can take a few minutes to appear; wait up to NPM_VERIFY_TIMEOUT seconds, default 900).
deadline=$(( $(date +%s) + ${NPM_VERIFY_TIMEOUT:-900} ))
while read -r d; do
  [ -n "$d" ] || continue
  name=$(node -p "require('$dist/$d/package.json').name")
  ver=$(node -p "require('$dist/$d/package.json').version")
  until [ "$(npm view "$name@$ver" version --registry=https://registry.npmjs.org/ 2>/dev/null)" = "$ver" ]; do
    if [ "$(date +%s)" -ge "$deadline" ]; then
      echo "::error::$name@$ver is not visible on registry.npmjs.org (publish did not take effect)"
      exit 1
    fi
    sleep 20
  done
  echo "verified $name@$ver on registry.npmjs.org"
done < "$dist/PUBLISH_ORDER"
