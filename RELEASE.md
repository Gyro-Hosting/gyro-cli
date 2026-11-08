# Releasing gyro-cli

1. `scripts/sync-spec.sh`, `go generate ./...`, `go vet ./...`, `go test ./...`.
2. Regenerate docs: `go run ./cmd/gyro gen-docs --dir docs/commands` (+ `GYRO_LANG=tr ... --dir docs/commands-tr`, `--format man --dir docs/man`).
3. `git tag vX.Y.Z && git push origin vX.Y.Z` → `release` workflow → GoReleaser: archives (with man pages + completions),
   `gyro_<v>_SHA256SUMS` + `.sig` (project GPG key, repository secrets `GPG_PRIVATE_KEY`, `PASSPHRASE`), deb/rpm/apk,
   multi-arch image `ghcr.io/gyro-hosting/gyro`; Homebrew cask / Scoop manifest when `HOMEBREW_TAP_TOKEN` / `SCOOP_BUCKET_TOKEN` are set.
4. The GitHub release is created as a **draft**: review the notes, then publish it.

Signing key fingerprint: `E96BE9635E77A0D9119AC8C3A7986589A154B990` (public key `packaging/gyro-release-signing-public.asc`).

## Automatic API sync (spec-sync)
`.github/workflows/spec-sync.yml` keeps the CLI in step with the Gyro Public API. It runs every 6 hours, on demand
(*Actions → spec-sync → Run workflow*) and on `repository_dispatch` event `spec-updated` (sent by the API build after a
spec change). It downloads the published OpenAPI document from `SPEC_URL` (repository secret or variable), and if it
differs from `internal/spec/openapi.yaml` it regenerates commands and docs, runs the tests, commits, tags
(new/removed operations → minor, description-only → patch) and calls `release.yml`, which publishes the release (not draft).
If tests or the release fail, an issue labelled `spec-sync` is opened instead.
