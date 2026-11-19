## gyro account card update

Kartı otomatik ödemede kullan / varsayılan yap

### Synopsis

Kartı otomatik ödemede kullan / varsayılan yap

`auto_pay` true kartı otomatik ödeme sırasının sonuna ekler (yenileme faturaları önce bakiyeden, sonra sırayla bu kartlardan ödenir), false çıkarır. `default: true` varsayılan kart yapar. Süresi dolmuş kart reddedilir.

API: PATCH /cloud/account/cards/{id} (operationId updateSavedCard, scope billing_pay)

```
gyro account card update <ID> [flags]
```

### Options

```
      --auto-pay                 auto_pay
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --default                  default
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for update
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
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

* [gyro account card](gyro_account_card.md)	 - Kayıtlı kartları listele

