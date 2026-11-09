## gyro service label set

Hizmetin takma adını ve/veya etiketlerini ayarla

### Synopsis

Hizmetin takma adını ve/veya etiketlerini ayarla

`nickname`, `tags` ya da ikisini gönderin; gönderilmeyen alan değişmez. `nickname` "" veya null temizler; `tags` tüm etiketlerin yerine geçer, `[]` temizler. Etiket `"ad"` ya da `{"name", "color"}`; bilinmeyen/eksik renk için sabit varsayılan atanır. Aynı adlı (büyük/küçük harf duyarsız) etiketler birleştirilir. Sınır: anahtar başına dakikada 30 etiket yazması (429 `rate_limited`). Alt hesap anahtarı `label:edit` eylemi ister. Hatalar: 422 `nothing_to_update`, `invalid_nickname`, `invalid_tags`; 404 `service_not_found` (başka hesabın hizmeti dahil) veya `feature_unavailable`.

API: PUT /cloud/services/{id}/label (operationId updateServiceLabel, scope services_write)

```
gyro service label set <ID> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for set
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --nickname string          "" veya null takma adı temizler. Kontrol karakteri içeremez.
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --tags strings             Tüm etiketlerin yerine geçer; [] veya null temizler. Öğe: etiket adı ya da {"name", "color"} (CLI: renk için --data kullanın).
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

* [gyro service label](gyro_service_label.md)	 - Hizmetin takma adını ve etiketlerini getir

