## gyro template list

Şablonlarınızı listele

### Synopsis

Şablonlarınızı listele

Şablonun gerçekten kullandığı alan üzerinden GB başına aylık 0,05 € (ön ödemeli bakiyeden saatlik; ilk şablonla açılan şablon deposu hizmeti üzerinden). Silince ücret durur. En çok 10 şablon; yalnız sizin hesabınızda görünür. `meta`: adet, sınır, toplam GB, fiyat.

API: GET /cloud/templates (operationId listTemplates, scope cloud_read)

```
gyro template list [flags]
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

* [gyro template](gyro_template.md)	 - Şablonlarınızı listele

