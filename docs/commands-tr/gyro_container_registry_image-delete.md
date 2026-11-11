## gyro container registry image-delete

İmaj etiketini sil

### Synopsis

İmaj etiketini sil

Sorgu `repo` ve `tag`. Bir servisin kullandığı imaj silinemez. `"confirm": true` gerekir.

API: DELETE /cloud/containers/{id}/registry/images (operationId typed_delete_containers__id__registry_images, scope services_write)

```
gyro container registry image-delete <ID> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for image-delete
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --repo string              Repository, e.g. web
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --tag string               Tag, e.g. 1.4
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

