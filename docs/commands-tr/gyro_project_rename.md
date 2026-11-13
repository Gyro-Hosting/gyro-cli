## gyro project rename

Projeyi yeniden adlandır / renklendir

### Synopsis

Projeyi yeniden adlandır / renklendir

name, description, color, icon alanlarından istediğinizi gönderin. Varsayılan proje değiştirilemez (409 `default_project`). Yazma: anahtar başına dakikada 60 (429 `rate_limited`). Alt hesap anahtarları proje oluşturamaz, yeniden adlandıramaz, silemez (403 `subkey_forbidden`).

API: PATCH /cloud/projects/{id} (operationId updateProject, scope services_write)

```
gyro project rename <ID> [flags]
```

### Options

```
      --color string             color (gray|blue|green|red|orange|yellow|purple|pink|teal|indigo)
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --description string       description
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for rename
      --icon string              icon (folder|briefcase|building|cart3|cloud|code-slash|controller|globe2|heart|house|kanban|lightning|people|rocket-takeoff|shield-lock|shop|star|tools|truck|box-seam|mortarboard|camera)
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --name string              Letter or digit first; then letters, digits, space and . _ - / + # & ( ). Unique per account (case-insensitive). Reserved: Genel, General, Default.
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
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

* [gyro project](gyro_project.md)	 - Projeleri listele

