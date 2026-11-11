## gyro container registry get

Ortamın Gyro imaj deposu

### Synopsis

Ortamın Gyro imaj deposu

Depo adresi, imaj öneki `<host>/<kod>/`, kullanıcı adı, kullanım (2 GB ücretsiz, üst sınır 50 GB), gönderim izni ve temizlik ayarı. Yalnız bu ortam kendi depolarına erişir. 2 GB üstü GB başına aylık 0,05 € (saatlik bakiyeden). Parola yalnız POST .../registry/credentials ile döner.

API: GET /cloud/containers/{id}/registry (operationId typed_get_containers__id__registry, scope services_read)

```
gyro container registry get <ID> [flags]
```

### Options

```
  -h, --help   help for get
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

* [gyro container registry](gyro_container_registry.md)	 - Ortamın Gyro imaj deposu

