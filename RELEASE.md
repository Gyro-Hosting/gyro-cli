# Releasing gyro-cli

1. `scripts/sync-spec.sh`, `go generate ./...`, `go vet ./...`, `go test ./...`.
2. Regenerate docs: `go run ./cmd/gyro gen-docs --dir docs/commands` (+ `GYRO_LANG=tr ... --dir docs/commands-tr`, `--format man --dir docs/man`).
3. `git tag vX.Y.Z && git push origin vX.Y.Z` → `release` workflow → GoReleaser: archives (with man pages + completions),
   `gyro_<v>_SHA256SUMS` + `.sig` (project GPG key, repository secrets `GPG_PRIVATE_KEY`, `PASSPHRASE`), deb/rpm/apk,
   multi-arch image `ghcr.io/gyro-hosting/gyro`; Homebrew cask / Scoop manifest when `HOMEBREW_TAP_TOKEN` / `SCOOP_BUCKET_TOKEN` are set.
4. The GitHub release is created as a **draft**: review the notes, then publish it.

Signing key fingerprint: `E96BE9635E77A0D9119AC8C3A7986589A154B990` (public key `packaging/gyro-release-signing-public.asc`).
