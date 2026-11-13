## gyro container pipeline create

Pipeline oluştur

### Synopsis

Pipeline oluştur

Depoyu bağlar. Yanıtta `webhook_url` ve `webhook_secret` (yalnız burada ve GET .../webhook) vardır; ikisini deponun webhook ayarına ekleyin. Özel depo için `token`.

API: POST /cloud/containers/{id}/pipelines (operationId pipeline_create, scope services_write)

```
gyro container pipeline create <ID> [flags]
```

### Options

```
      --approval                 Wait for approval before deploying (24 h)
      --auto                     Webhook pushes start runs
      --auto-rollback            Roll back when the new version is unhealthy
      --branch string            Push'u derlenen dallar (virgüllü glob, boş = yok)
      --build-args string        Dockerfile ARG values (visible in the image history) (JSON)
      --build-env commands       Variables for commands (JSON)
      --build-image commands     Image that runs commands (e.g. node:22-alpine)
      --commands strings         Shell commands run in the repository root before the Dockerfile (tests, npm ci ...); build secrets are environment variables here
      --context string           Build folder (empty = repository root)
      --cpu-m int                Build CPU (millicores, plan limit)
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dockerfile string        dockerfile
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
      --git-user string          User name for the token (default: x-access-token / oauth2 / git)
  -h, --help                     help for create
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --image string             Ortam deposundaki imaj adı
      --memory-mi int            Build memory (MiB, plan limit)
      --name string              Görünen ad (1-48)
      --notify strings           Notifications (channels: the environment's alert settings)
      --provider string          provider (github|gitlab|gitea|git)
      --repo string              https klon adresi, kimlik bilgisi yok
      --repo-file                Read gyro-pipeline.yml from the repository (GitHub, GitLab, Gitea)
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --tags string              Derlenen etiketler (glob, ör. v*)
      --targets strings          Yeni imaja geçecek servisler (boş = yalnız derle)
      --timeout-min int          Build time limit (minutes, plan limit)
      --token string             Özel depo için yalnız okuma erişim belirteci (şifreli saklanır, geri dönmez; güncellemede "-" siler)
      --verify-min int           Health wait after deploy (1-30 minutes)
```

### Options inherited from parent commands

```
      --api-url string     API taban adresi (ortam GYRO_API_URL)
      --debug              HTTP isteklerini stderr'e yaz (sırlar maskeli)
      --lang string        mesaj dili: tr ya da en (varsayılan yerel ayar)
      --no-color           renkleri kapat (NO_COLOR da olur)
      --no-header          tablo başlığını yazma
  -o, --output string      çıktı biçimi: table, json, yaml, text
  -p, --profile string     kullanılacak profil (ortam GYRO_PROFILE)
      --query string       sonuca uygulanacak JMESPath sorgusu
  -q, --quiet              yalnız kimlikleri yaz
      --timeout duration   toplam zaman aşımı, ör. 30s, 10m
      --wide               tüm tablo sütunlarını göster
  -y, --yes                onay sorma
```

### SEE ALSO

* [gyro container pipeline](gyro_container_pipeline.md)	 - Pipeline'ları listele

