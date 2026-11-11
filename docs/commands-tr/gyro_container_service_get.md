## gyro container service get

Kapsayıcı servisini getir

### Synopsis

Kapsayıcı servisini getir

Tanım, durum, sorun kodu, örnekler, sürüm, alan adları ve gizli değer adları (değerler asla dönmez).

API: GET /cloud/containers/{id}/services/{name} (operationId typed_get_containers__id__services__name, scope services_read)

```
gyro container service get <ID> <NAME> [flags]
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

* [gyro container service](gyro_container_service.md)	 - Kapsayıcı servislerini listele

