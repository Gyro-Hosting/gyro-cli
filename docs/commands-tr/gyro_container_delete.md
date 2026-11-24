## gyro container delete

Kapsayıcı ortamını sil (ön ödemeli)

### Synopsis

Kapsayıcı ortamını sil (ön ödemeli)

Paneldeki Harcamalar > Sil ile aynı: saatlik ücret hemen durur, ortam tüm servis, düğüm, imaj deposu, pipeline ve özel alan adı bağlarıyla hemen silinir (saklama süresi yok). `"confirm": true` (gövde ya da sorgu) ve `Idempotency-Key` başlığı gerekir; tekrar çağrı 202 `deleted` döner. Faturalı ortamlar 409 `not_prepaid`. Terraform: `gyro_container_environment` silinince bu çağrılır.

API: DELETE /cloud/containers/{id} (operationId deleteContainerEnvironment, scope cloud_delete)

```
gyro container delete <ID> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for delete
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
```

### Options inherited from parent commands

```
      --account gyro account list   panel kullanıcısı olduğunuz başka bir hesapta işlem yap (gyro account list id'si; "self" = kendi hesabınız; ortam GYRO_ACCOUNT, profil varsayılanı: `gyro account use`)
      --api-url string              API taban adresi (ortam GYRO_API_URL)
      --debug                       HTTP isteklerini stderr'e yaz (sırlar maskeli)
      --lang string                 mesaj dili: tr ya da en (varsayılan yerel ayar)
      --no-color                    renkleri kapat (NO_COLOR da olur)
      --no-header                   tablo başlığını yazma
  -o, --output string               çıktı biçimi: table, json, yaml, text
  -p, --profile string              kullanılacak profil (ortam GYRO_PROFILE)
      --query string                sonuca uygulanacak JMESPath sorgusu
  -q, --quiet                       yalnız kimlikleri yaz
      --timeout duration            toplam zaman aşımı, ör. 30s, 10m
      --wide                        tüm tablo sütunlarını göster
  -y, --yes                         onay sorma
```

### SEE ALSO

* [gyro container](gyro_container.md)	 - Gyro Kapsayıcı (ECS), İstanbul: servisler, sürümler, loglar, alan adları

