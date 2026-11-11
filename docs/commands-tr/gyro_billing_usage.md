## gyro billing usage

Saat saat kullanım (ya da günlük/aylık toplam)

```
gyro billing usage [flags]
```

### Examples

```
  gyro billing usage --from 2026-11-01 --to 2026-11-30
  gyro billing usage --service 17501 --daily
  gyro billing usage --monthly -o json
  gyro billing usage --csv kullanim.csv
```

### Options

```
      --all             tüm sayfaları getir
      --csv string      saatleri CSV kaydet (dosya, '.' = sunucunun adı, '-' = ekran)
      --daily           günlük toplamlar
      --from string     ilk gün YYYY-AA-GG (varsayılan: bu ayın 1'i)
  -h, --help            help for usage
      --monthly         aylık toplamlar
      --service int     yalnız bu hizmet no
      --status string   ucretli|tavan|kismi|odenmedi|ucretsiz
      --to string       son gün YYYY-AA-GG (varsayılan: bugün)
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

* [gyro billing](gyro_billing.md)	 - Ön ödemeli saatlik kullanım (İstanbul bulut, kapsayıcı): bakiye, kullanım, tahmin, döküm

