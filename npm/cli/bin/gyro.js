#!/usr/bin/env node
// npm launcher for the gyro CLI. The binary itself lives in a per-platform package
// (@gyro-hosting/cli-<os>-<arch>) that npm installs through optionalDependencies.
// Nothing is downloaded at install time; this file only finds and runs that binary.
"use strict";

const { spawnSync } = require("node:child_process");
const fs = require("node:fs");
const path = require("node:path");

const PLATFORMS = {
  "linux-x64": "@gyro-hosting/cli-linux-x64",
  "linux-arm64": "@gyro-hosting/cli-linux-arm64",
  "darwin-x64": "@gyro-hosting/cli-darwin-x64",
  "darwin-arm64": "@gyro-hosting/cli-darwin-arm64",
  "win32-x64": "@gyro-hosting/cli-win32-x64",
  "win32-arm64": "@gyro-hosting/cli-win32-arm64",
};

function binaryPath() {
  if (process.env.GYRO_BINARY_PATH) return process.env.GYRO_BINARY_PATH;
  const key = `${process.platform}-${process.arch}`;
  const pkg = PLATFORMS[key];
  if (!pkg) {
    fail(
      `gyro: this platform (${key}) has no prebuilt binary.\n` +
        `Supported: ${Object.keys(PLATFORMS).join(", ")}.\n` +
        "Other install options: https://github.com/Gyro-Hosting/gyro-cli#readme"
    );
  }
  const exe = process.platform === "win32" ? "gyro.exe" : "gyro";
  let dir;
  try {
    dir = path.dirname(require.resolve(`${pkg}/package.json`));
  } catch (_) {
    fail(
      `gyro: the platform package ${pkg} is not installed.\n` +
        "It is an optional dependency; it is skipped when npm runs with --omit=optional / --no-optional,\n" +
        "when the lock file was created on another platform, or when the registry mirror lacks it.\n" +
        `Fix: npm install -g @gyro-hosting/cli (without --omit=optional), or npm install ${pkg}.\n` +
        "Other install options: https://github.com/Gyro-Hosting/gyro-cli#readme"
    );
  }
  const bin = path.join(dir, "bin", exe);
  if (!fs.existsSync(bin)) fail(`gyro: binary missing in ${pkg} (${bin}). Reinstall the package.`);
  return bin;
}

function fail(msg) {
  process.stderr.write(msg + "\n");
  process.exit(1);
}

const r = spawnSync(binaryPath(), process.argv.slice(2), { stdio: "inherit", windowsHide: false });
if (r.error) fail(`gyro: could not start the binary: ${r.error.message}`);
if (r.signal) process.kill(process.pid, r.signal);
process.exit(r.status === null ? 1 : r.status);
