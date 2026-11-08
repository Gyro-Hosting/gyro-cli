#!/bin/sh
# Gyro CLI installer (Linux / macOS).
#
#   curl -fsSL https://www.gyrohosting.com/cli/install.sh | sh
#   curl -fsSL https://www.gyrohosting.com/cli/install.sh | GYRO_VERSION=v1.0.0 GYRO_INSTALL_DIR=$HOME/.local/bin sh
#
# Downloads the release archive from GitHub, verifies its SHA-256 checksum (and the GPG signature of the
# checksum file when gpg is available) and installs the "gyro" binary. Windows: use Scoop or the zip.
set -eu

REPO="${GYRO_REPO:-Gyro-Hosting/gyro-cli}"
VERSION="${GYRO_VERSION:-latest}"
INSTALL_DIR="${GYRO_INSTALL_DIR:-}"
KEY_FPR="E96BE9635E77A0D9119AC8C3A7986589A154B990"
KEY_URL="${GYRO_KEY_URL:-https://raw.githubusercontent.com/$REPO/main/packaging/gyro-release-signing-public.asc}"

say() { printf '%s\n' "gyro-install: $*" >&2; }
die() { say "error: $*"; exit 1; }
need() { command -v "$1" >/dev/null 2>&1 || die "$1 is required"; }

need uname
need tar
if command -v curl >/dev/null 2>&1; then
  fetch() { curl -fsSL --retry 3 -o "$2" "$1"; }
  fetch_stdout() { curl -fsSL --retry 3 "$1"; }
elif command -v wget >/dev/null 2>&1; then
  fetch() { wget -q -O "$2" "$1"; }
  fetch_stdout() { wget -q -O - "$1"; }
else
  die "curl or wget is required"
fi

os=$(uname -s | tr '[:upper:]' '[:lower:]')
case "$os" in
  linux|darwin) ;;
  *) die "unsupported OS $os (Windows: scoop install gyro, or download the zip)" ;;
esac
arch=$(uname -m)
case "$arch" in
  x86_64|amd64) arch=amd64 ;;
  aarch64|arm64) arch=arm64 ;;
  *) die "unsupported architecture $arch" ;;
esac

if [ "$VERSION" = "latest" ]; then
  VERSION=$(fetch_stdout "https://api.github.com/repos/$REPO/releases/latest" | sed -n 's/.*"tag_name": *"\([^"]*\)".*/\1/p' | head -n1)
  [ -n "$VERSION" ] || die "could not find the latest release of $REPO"
fi
num=${VERSION#v}
archive="gyro_${num}_${os}_${arch}.tar.gz"
sums="gyro_${num}_SHA256SUMS"
base="https://github.com/$REPO/releases/download/$VERSION"

tmp=$(mktemp -d)
trap 'rm -rf "$tmp"' EXIT INT TERM
say "downloading gyro $VERSION ($os/$arch)"
fetch "$base/$archive" "$tmp/$archive" || die "download failed: $base/$archive"
fetch "$base/$sums" "$tmp/$sums" || die "download failed: $base/$sums"

if command -v gpg >/dev/null 2>&1 && fetch "$base/$sums.sig" "$tmp/$sums.sig" 2>/dev/null; then
  export GNUPGHOME="$tmp/gnupg"; mkdir -m 700 "$GNUPGHOME"
  if fetch "$KEY_URL" "$tmp/key.asc" && gpg --batch --quiet --import "$tmp/key.asc" 2>/dev/null; then
    if gpg --batch --status-fd 1 --verify "$tmp/$sums.sig" "$tmp/$sums" 2>/dev/null | grep -q "VALIDSIG $KEY_FPR"; then
      say "signature OK ($KEY_FPR)"
    else
      die "GPG signature of $sums is NOT valid"
    fi
  else
    say "warning: could not import the release key, skipping signature check"
  fi
fi

expected=$(grep " $archive\$" "$tmp/$sums" | awk '{print $1}')
[ -n "$expected" ] || die "$archive is not listed in $sums"
if command -v sha256sum >/dev/null 2>&1; then
  actual=$(sha256sum "$tmp/$archive" | awk '{print $1}')
else
  actual=$(shasum -a 256 "$tmp/$archive" | awk '{print $1}')
fi
[ "$expected" = "$actual" ] || die "checksum mismatch for $archive"
say "checksum OK"

tar -xzf "$tmp/$archive" -C "$tmp" gyro
if [ -z "$INSTALL_DIR" ]; then
  if [ -w /usr/local/bin ]; then INSTALL_DIR=/usr/local/bin
  elif command -v sudo >/dev/null 2>&1 && { [ -t 0 ] || [ -t 2 ]; }; then INSTALL_DIR=/usr/local/bin; SUDO=sudo
  else INSTALL_DIR="$HOME/.local/bin"
  fi
fi
mkdir -p "$INSTALL_DIR" 2>/dev/null || ${SUDO:-} mkdir -p "$INSTALL_DIR"
if [ -w "$INSTALL_DIR" ]; then SUDO=""; fi
${SUDO:-} install -m 0755 "$tmp/gyro" "$INSTALL_DIR/gyro"
say "installed $INSTALL_DIR/gyro"
case ":$PATH:" in *":$INSTALL_DIR:"*) ;; *) say "add $INSTALL_DIR to your PATH" ;; esac
"$INSTALL_DIR/gyro" version || true
say "next: gyro configure   (create an API key in Client area -> Account -> API)"
