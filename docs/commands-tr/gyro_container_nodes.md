## gyro container nodes

Düğüm havuzu

### Synopsis

Düğüm havuzu

Size ayrılmış düğümler, havuz kapasitesi/kullanımı, kök alan adı A kaydı için IP'ler, özel depo (yalnız kullanıcı).

API: GET /cloud/containers/{id}/nodes (operationId typed_get_containers__id__nodes, scope services_read)

```
gyro container nodes <ID> [flags]
```

### Options

```
  -h, --help   help for nodes
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

* [gyro container](gyro_container.md)	 - Gyro Kapsayıcı (ECS), İstanbul: servisler, sürümler, loglar, alan adları

