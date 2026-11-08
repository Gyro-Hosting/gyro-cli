## gyro reseller report get

Uç-müşteri başına kullanım ve fatura raporu

### Synopsis

Uç-müşteri başına kullanım ve fatura raporu

API: GET /cloud/reseller/report (operationId resellerReport, scope reseller)

```
gyro reseller report get [flags]
```

### Options

```
      --from string   
  -h, --help          help for get
      --to string     
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

* [gyro reseller report](gyro_reseller_report.md)	 - Uç-müşteri başına kullanım ve fatura raporu

