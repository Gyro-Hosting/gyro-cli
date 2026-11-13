## gyro project move

Hizmetleri projeye taşı

### Synopsis

Hizmetleri projeye taşı

Tüm numaralar anahtarın kapsamındaki hizmetleriniz olmalı, yoksa 404 `service_not_found` ve hiçbir şey taşınmaz. Çağrı başına en çok 500; id 0 = varsayılan proje. Alt hesap anahtarı `project:move` eylemi ister.

API: POST /cloud/projects/{id}/services (operationId moveServicesToProject, scope services_write)

```
gyro project move <ID> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for move
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --service-ids int64Slice   service_ids (default [])
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

