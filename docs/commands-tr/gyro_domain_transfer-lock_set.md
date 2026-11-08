## gyro domain transfer-lock set

Transfer Kilidini Değiştirme

### Synopsis

Transfer Kilidini Değiştirme

Alan adını transfere kapatır ya da açar. — sağlayıcıya yazar

API: PUT /domains/{domain}/transfer-lock (operationId coreUpdateTransferLock, scope Domains/UpdateTransferLock)

```
gyro domain transfer-lock set <DOMAIN> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for set
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --locked                   Yeni kilit durumu. Gerçek bir mantıksal değer olmalıdır; metin ve sayı reddedilir.
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

* [gyro domain transfer-lock](gyro_domain_transfer-lock.md)	 - Transfer Kilidini Okuma

