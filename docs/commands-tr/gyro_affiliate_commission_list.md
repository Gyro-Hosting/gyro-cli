## gyro affiliate commission list

Komisyonları Listeleme

### Synopsis

Komisyonları Listeleme

Kazanılan komisyonları en yeniden başlayarak döndürür. — anahtarın sahibi

API: GET /affiliate/commissions (operationId coreGetAffiliateCommissions, scope Affiliate/GetAffiliateCommissions)

```
gyro affiliate commission list [flags]
```

### Options

```
      --all         tüm sayfaları getir
  -h, --help        help for list
      --limit int   sayfa başına kayıt (en çok 100)
      --page int    sayfa numarası (default 1)
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

* [gyro affiliate commission](gyro_affiliate_commission.md)	 - Komisyonları Listeleme

