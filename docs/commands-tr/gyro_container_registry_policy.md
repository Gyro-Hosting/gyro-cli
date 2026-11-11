## gyro container registry policy

İmaj temizlik ayarı

### Synopsis

İmaj temizlik ayarı

`enabled`, `keep` 1-50 (varsayılan 10), `untagged_days` 1-90 (varsayılan 7).

API: PUT /cloud/containers/{id}/registry/policy (operationId typed_put_containers__id__registry_policy, scope services_write)

```
gyro container registry policy <ID> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
      --enabled                  enabled
  -h, --help                     help for policy
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --keep int                 keep
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --untagged-days int        untagged_days
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

* [gyro container registry](gyro_container_registry.md)	 - Ortamın Gyro imaj deposu

