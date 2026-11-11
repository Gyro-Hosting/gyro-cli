## gyro container external-registry set

Harici özel depo bilgisi

### Synopsis

Harici özel depo bilgisi

Ortam başına bir özel depo (ör. ghcr.io + yalnız çekme izinli token). Token şifreli saklanır, geri dönmez.

API: PUT /cloud/containers/{id}/external-registry (operationId typed_put_containers__id__external_registry, scope services_write)

```
gyro container external-registry set <ID> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for set
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --password string          password
      --server string            server
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --username string          username
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

* [gyro container external-registry](gyro_container_external-registry.md)	 - Özel depo bilgisini sil

