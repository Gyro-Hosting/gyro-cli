# @gyro-hosting/cli — the `gyro` command for Gyro Hosting

`gyro` is the official command-line interface for [Gyro Hosting](https://www.gyrohosting.com).
Every operation of the Gyro Public API is a command: cloud servers, containers, snapshots, SSH keys, load balancers,
domains and DNS, S3 object storage, corporate mail, uptime monitoring, invoices, balance, support tickets and more.

This npm package ships the same signed binary as the GitHub release (Linux, macOS, Windows; x64 and arm64).
No install scripts run: npm picks the matching `@gyro-hosting/cli-<os>-<arch>` package and a small Node launcher starts it.

## Install

```sh
npm install -g @gyro-hosting/cli     # then: gyro --help
npx @gyro-hosting/cli --version       # run once without installing
```

Other ways to install:

| Method | Command |
|---|---|
| Script (Linux/macOS) | `curl -fsSL https://www.gyrohosting.com/cli/install.sh \| sh` |
| Scoop (Windows) | `scoop bucket add gyro https://github.com/Gyro-Hosting/scoop-bucket && scoop install gyro` |
| Debian / RHEL | `.deb` / `.rpm` from the [release page](https://github.com/Gyro-Hosting/gyro-cli/releases) |
| Docker | `docker run --rm -e GYRO_TOKEN ghcr.io/gyro-hosting/gyro server list` |

## Getting started

```sh
gyro configure            # API key: Client area → account menu (top right) → API Credentials (wck_...)
gyro whoami
gyro plan list
gyro server create web-1 --plan standart-2-4 --region eu-de-1 --image ubuntu-24.04 \
     --ssh-key ~/.ssh/id_ed25519.pub --wait
gyro server list
```

In CI, set `GYRO_TOKEN` instead of running `gyro configure`. Output formats: `-o table|json|yaml|text`, `--query` (JMESPath).
Money and destructive actions ask for confirmation (`--yes` in scripts); `--dry-run` shows the price or the request without changing anything.

## Türkçe

```sh
npm install -g @gyro-hosting/cli
gyro configure   # API anahtarı: Müşteri paneli → hesap menüsü → API Kimlik Bilgileri
gyro server list
```

## Links

- Documentation and API reference: <https://www.gyrohosting.com/developers>
- Source, releases and checksums: <https://github.com/Gyro-Hosting/gyro-cli>
- Terraform provider: [`gyro-hosting/gyro`](https://registry.terraform.io/providers/gyro-hosting/gyro)

## Integrity

Each platform binary is byte-identical to the archive in the GitHub release. The release checksum list
`gyro_<version>_SHA256SUMS` is signed with the Gyro release key
`E96B E963 5E77 A0D9 119A C8C3 A798 6589 A154 B990`, and the npm packages are published from GitHub Actions
with npm provenance. Set `GYRO_BINARY_PATH` to run a binary of your own instead.

License: MPL-2.0
