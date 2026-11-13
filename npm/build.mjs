#!/usr/bin/env node
// Builds the npm packages for a published gyro release.
//   node npm/build.mjs --version 1.26.0 [--out npm/dist] [--assets <dir>] [--repo Gyro-Hosting/gyro-cli]
//                      [--gpg-key packaging/gyro-release-signing-public.asc] [--require-signature]
// Steps: download the release archives + gyro_<v>_SHA256SUMS (+ .sig), verify the GPG signature of the
// checksum list (when gpg and the key are available; mandatory with --require-signature), verify every
// archive against the list, extract the binary and write one package per platform plus the
// @gyro-hosting/cli launcher package. Nothing in the produced packages runs at install time.
import { createHash } from "node:crypto";
import { execFileSync } from "node:child_process";
import fs from "node:fs";
import os from "node:os";
import path from "node:path";
import { fileURLToPath } from "node:url";

const here = path.dirname(fileURLToPath(import.meta.url));
const root = path.resolve(here, "..");
const args = Object.fromEntries(
  process.argv.slice(2).reduce((acc, a, i, all) => {
    if (a.startsWith("--")) acc.push([a.slice(2), all[i + 1] && !all[i + 1].startsWith("--") ? all[i + 1] : true]);
    return acc;
  }, [])
);
const version = String(args.version || "").replace(/^v/, "");
if (!/^\d+\.\d+\.\d+(-[0-9A-Za-z.-]+)?$/.test(version)) die("usage: build.mjs --version X.Y.Z");
const repo = args.repo || "Gyro-Hosting/gyro-cli";
const out = path.resolve(args.out || path.join(here, "dist"));
const assets = path.resolve(args.assets || path.join(os.tmpdir(), `gyro-npm-assets-${version}`));
const gpgKey = path.resolve(args["gpg-key"] || path.join(root, "packaging", "gyro-release-signing-public.asc"));
const SCOPE = "@gyro-hosting";
const META = {
  license: "MPL-2.0",
  author: "Gyro Hosting (https://www.gyrohosting.com)",
  homepage: "https://www.gyrohosting.com/developers",
  repository: { type: "git", url: "git+https://github.com/Gyro-Hosting/gyro-cli.git" },
  bugs: { url: "https://github.com/Gyro-Hosting/gyro-cli/issues" },
};
// npm os/cpu -> goreleaser os/arch
const TARGETS = [
  { os: "linux", cpu: "x64", gos: "linux", garch: "amd64", ext: "tar.gz" },
  { os: "linux", cpu: "arm64", gos: "linux", garch: "arm64", ext: "tar.gz" },
  { os: "darwin", cpu: "x64", gos: "darwin", garch: "amd64", ext: "tar.gz" },
  { os: "darwin", cpu: "arm64", gos: "darwin", garch: "arm64", ext: "tar.gz" },
  { os: "win32", cpu: "x64", gos: "windows", garch: "amd64", ext: "zip" },
  { os: "win32", cpu: "arm64", gos: "windows", garch: "arm64", ext: "zip" },
];

function die(m) { console.error(m); process.exit(1); }
const sha256 = (f) => createHash("sha256").update(fs.readFileSync(f)).digest("hex");
const url = (name) => `https://github.com/${repo}/releases/download/v${version}/${name}`;

async function fetchTo(name) {
  const dest = path.join(assets, name);
  if (fs.existsSync(dest) && fs.statSync(dest).size > 0) return dest;
  for (let attempt = 1; ; attempt++) {
    const r = await fetch(url(name), { redirect: "follow" });
    if (r.ok) { fs.writeFileSync(dest, Buffer.from(await r.arrayBuffer())); return dest; }
    if (attempt >= 4 || r.status === 404) die(`download failed: ${name}: HTTP ${r.status}`);
    await new Promise((s) => setTimeout(s, 2000 * attempt));
  }
}

function haveGpg() { try { execFileSync("gpg", ["--version"], { stdio: "ignore" }); return true; } catch { return false; } }

function verifySignature(sums, sig) {
  if (!fs.existsSync(gpgKey) || !haveGpg()) {
    if (args["require-signature"]) die("signature check required but gpg or the key is missing");
    console.warn("WARNING: gpg/key unavailable, SHA256SUMS signature NOT checked (checksums still verified)");
    return false;
  }
  const home = fs.mkdtempSync(path.join(os.tmpdir(), "gyro-gpg-"));
  try {
    execFileSync("gpg", ["--batch", "--homedir", home, "--import", gpgKey], { stdio: "ignore" });
    execFileSync("gpg", ["--batch", "--homedir", home, "--verify", sig, sums], { stdio: "ignore" });
    const fpr = execFileSync("gpg", ["--batch", "--homedir", home, "--with-colons", "--fingerprint"]).toString();
    if (!fpr.includes("E96BE9635E77A0D9119AC8C3A7986589A154B990")) die("unexpected signing key");
    console.log("signature OK (E96B E963 ... A154 B990)");
    return true;
  } catch (e) { die("SHA256SUMS signature verification FAILED"); }
  finally { fs.rmSync(home, { recursive: true, force: true }); }
}

function writeJson(f, o) { fs.writeFileSync(f, JSON.stringify(o, null, 2) + "\n"); }

const platformReadme = (name, t) => `# ${name}

The \`gyro\` command-line tool for Gyro Hosting, prebuilt for ${t.os} ${t.cpu}.

Do not install this package directly. Install [\`${SCOPE}/cli\`](https://www.npmjs.com/package/${SCOPE}/cli),
which selects the right platform package automatically:

\`\`\`sh
npm install -g ${SCOPE}/cli
\`\`\`

The binary is taken unchanged from the signed GitHub release v${version}
(<https://github.com/Gyro-Hosting/gyro-cli/releases/tag/v${version}>); its SHA-256 matches \`gyro_${version}_SHA256SUMS\`.
`;

async function main() {
  fs.mkdirSync(assets, { recursive: true });
  fs.rmSync(out, { recursive: true, force: true });
  fs.mkdirSync(out, { recursive: true });

  const sumsName = `gyro_${version}_SHA256SUMS`;
  const sums = await fetchTo(sumsName);
  const sig = await fetchTo(`${sumsName}.sig`);
  const signed = verifySignature(sums, sig);
  const expected = Object.fromEntries(
    fs.readFileSync(sums, "utf8").trim().split("\n").map((l) => { const [h, n] = l.trim().split(/\s+/); return [n.replace(/^\*/, ""), h]; })
  );

  const manifest = [];
  const optionalDependencies = {};
  for (const t of TARGETS) {
    const archive = `gyro_${version}_${t.gos}_${t.garch}.${t.ext}`;
    if (!expected[archive]) die(`${archive} is not listed in ${sumsName}`);
    const file = await fetchTo(archive);
    const got = sha256(file);
    if (got !== expected[archive]) die(`CHECKSUM MISMATCH ${archive}: ${got} != ${expected[archive]}`);

    const name = `${SCOPE}/cli-${t.os}-${t.cpu}`;
    const dir = path.join(out, `cli-${t.os}-${t.cpu}`);
    const bindir = path.join(dir, "bin");
    fs.mkdirSync(bindir, { recursive: true });
    const exe = t.os === "win32" ? "gyro.exe" : "gyro";
    const tmp = fs.mkdtempSync(path.join(os.tmpdir(), "gyro-x-"));
    if (t.ext === "zip") execFileSync("unzip", ["-q", "-o", file, exe, "-d", tmp]);
    else execFileSync("tar", ["-xzf", file, "-C", tmp, exe]);
    fs.copyFileSync(path.join(tmp, exe), path.join(bindir, exe));
    fs.chmodSync(path.join(bindir, exe), 0o755);
    fs.rmSync(tmp, { recursive: true, force: true });

    writeJson(path.join(dir, "package.json"), {
      name, version,
      description: `gyro CLI binary for ${t.os}-${t.cpu} (installed automatically by ${SCOPE}/cli)`,
      ...META,
      os: [t.os], cpu: [t.cpu],
      files: ["bin/", "README.md", "LICENSE"],
      preferUnplugged: true,
      publishConfig: { access: "public" },
    });
    fs.writeFileSync(path.join(dir, "README.md"), platformReadme(name, t));
    fs.copyFileSync(path.join(root, "LICENSE"), path.join(dir, "LICENSE"));
    optionalDependencies[name] = version;
    manifest.push(`${sha256(path.join(bindir, exe))}  ${name}/bin/${exe}  (from ${archive} ${got})`);
  }

  const main = path.join(out, "cli");
  fs.mkdirSync(path.join(main, "bin"), { recursive: true });
  fs.copyFileSync(path.join(here, "cli", "bin", "gyro.js"), path.join(main, "bin", "gyro.js"));
  fs.chmodSync(path.join(main, "bin", "gyro.js"), 0o755);
  fs.copyFileSync(path.join(here, "cli", "README.md"), path.join(main, "README.md"));
  fs.copyFileSync(path.join(root, "LICENSE"), path.join(main, "LICENSE"));
  writeJson(path.join(main, "package.json"), {
    name: `${SCOPE}/cli`, version,
    description: "Official command-line interface for Gyro Hosting (cloud servers, containers, DNS, S3 storage, mail, billing).",
    keywords: ["gyro", "gyrohosting", "cli", "cloud", "vps", "dns", "s3", "devops", "hosting"],
    ...META,
    bin: { gyro: "bin/gyro.js" },
    files: ["bin/", "README.md", "LICENSE"],
    engines: { node: ">=16" },
    optionalDependencies,
    publishConfig: { access: "public" },
  });

  fs.writeFileSync(path.join(out, "BINARIES.sha256"), manifest.join("\n") + "\n");
  fs.writeFileSync(path.join(out, "PUBLISH_ORDER"),
    TARGETS.map((t) => `cli-${t.os}-${t.cpu}`).concat("cli").join("\n") + "\n");
  console.log(`built ${TARGETS.length + 1} packages for v${version} in ${out} (signature ${signed ? "verified" : "not checked"})`);

  const hostKey = `${process.platform}-${process.arch}`;
  const host = TARGETS.find((t) => `${t.os}-${t.cpu}` === hostKey);
  if (host) {
    const b = path.join(out, `cli-${hostKey}`, "bin", host.os === "win32" ? "gyro.exe" : "gyro");
    const v = execFileSync(b, ["--version"]).toString().trim();
    if (!v.includes(version)) die(`host binary reports "${v}", expected ${version}`);
    console.log(`host binary OK: ${v}`);
  }
}
main().catch((e) => die(e.stack || String(e)));
