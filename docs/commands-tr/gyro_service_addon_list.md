## gyro service addon list

Hizmet Eklentilerini Okuma

### Synopsis

Hizmet Eklentilerini Okuma

Hizmete bağlı ek hizmetleri döndürür. — anahtarın sahibi

API: GET /services/{id}/addons (operationId coreGetServiceAddons, scope Services/GetServiceAddons)

```
gyro service addon list <ID> [flags]
```

### Options

```
  -h, --help   help for list
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

* [gyro service addon](gyro_service_addon.md)	 - Hizmet Eklentilerini Okuma

