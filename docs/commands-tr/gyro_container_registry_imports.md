## gyro container registry imports

İmaj içe aktarmaları (ilerleme ve geçmiş)

### Synopsis

İmaj içe aktarmaları (ilerleme ve geçmiş)

`active` = süren içe aktarma (ya da null): `phase`, bayt ve katman ilerlemesi; `jobs` = son 20 içe aktarma (yeniden eskiye), başarısızsa `error` + `message`; `limits`.

API: GET /cloud/containers/{id}/registry/imports (operationId listContainerImageImports, scope services_read)

```
gyro container registry imports <ID> [flags]
```

### Options

```
  -h, --help   help for imports
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

