package cli

import (
	"fmt"
	"strings"

	"github.com/Gyro-Hosting/gyro-cli/internal/i18n"
	"github.com/Gyro-Hosting/gyro-cli/internal/spec"
)

// 1.29.3 (#3892): `container pipeline create` used to demand --repo for every source because the spec listed it as required,
// while an archive pipeline (--provider archive) must have NO repository (the API refuses one with invalid_repo). The rule now
// depends on the source: git providers (github, gitlab, gitea, git; default github) need --repo, archive must not have it.
// `container pipeline update` refuses --repo together with --provider archive.
func pipelineSourceCheck(op *spec.Op, body map[string]any, flag func(string) string) error {
	if op.ID != "pipeline_create" && op.ID != "pipeline_update" {
		return nil
	}
	prov, _ := body["provider"].(string)
	prov = strings.ToLower(strings.TrimSpace(prov))
	if prov == "" && op.ID == "pipeline_create" {
		prov = "github" // server default
	}
	repoRaw, repoSet := body["repo"]
	repo := ""
	if s, ok := repoRaw.(string); ok {
		repo = strings.TrimSpace(s)
	}
	if prov == "archive" {
		if repo != "" {
			return &ExitErr{ExitUsage, fmt.Errorf("%s", i18n.F(
				"--%s is not used with --provider archive: an archive pipeline has no repository, it builds the archives you upload (gyro container pipeline upload). Remove --%s, or use a git provider (github, gitlab, gitea, git).",
				"--%s, --provider archive ile kullanılmaz: arşiv pipeline'ının deposu yoktur, yüklediğiniz arşivleri derler (gyro container pipeline upload). --%s'yi kaldırın ya da bir git sağlayıcısı seçin (github, gitlab, gitea, git).",
				flag("repo"), flag("repo")))}
		}
		if repoSet {
			delete(body, "repo") // empty value: nothing to send
		}
		return nil
	}
	if op.ID == "pipeline_create" && repo == "" {
		return &ExitErr{ExitUsage, fmt.Errorf("%s", i18n.F(
			"--%s is required for --provider %s (https clone URL or git@host:org/repo.git). To build uploaded archives without a repository use --provider archive.",
			"--provider %[2]s için --%[1]s zorunlu (https klon adresi ya da git@sunucu:kurum/depo.git). Depo olmadan yüklenen arşivleri derlemek için --provider archive kullanın.",
			flag("repo"), prov))}
	}
	return nil
}
