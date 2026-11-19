## gyro account auto-topup set

Otomatik yüklemeyi aç/kapat

### Synopsis

Otomatik yüklemeyi aç/kapat

Açmak için `consent: true` gerekir (çekimlere onay). `amount` en az 250 TL, en çok `limits.max_single`. Gönderilmeyen alan değişmez. `card_id` 0 = otomatik ödeme kartı.

API: PUT /cloud/account/auto-topup (operationId setAutoTopup, scope billing_pay)

```
gyro account auto-topup set [flags]
```

### Options

```
      --amount float             amount
      --card-id int              card_id
      --consent                  consent
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
      --enabled                  enabled
  -h, --help                     help for set
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --threshold float          threshold
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

* [gyro account auto-topup](gyro_account_auto-topup.md)	 - Otomatik bakiye yükleme ayarı

