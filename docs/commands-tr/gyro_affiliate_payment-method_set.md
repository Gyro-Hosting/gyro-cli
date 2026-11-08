## gyro affiliate payment-method set

Ödeme Hedefini Kaydetme

### Synopsis

Ödeme Hedefini Kaydetme

Bir ödeme yolunun hedef bilgisini kaydeder ya da siler. — anahtarın sahibi

API: PUT /affiliate/payment-method (operationId coreSaveAffiliatePayoutMethod, scope Affiliate/SaveAffiliatePayoutMethod)

```
gyro affiliate payment-method set [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
      --gateway string           Ödeme yolu.
  -h, --help                     help for set
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --info string              Hedef bilgisi. Boş göndermek kaydı siler.
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

* [gyro affiliate payment-method](gyro_affiliate_payment-method.md)	 - Ödeme Hedefini Kaydetme

