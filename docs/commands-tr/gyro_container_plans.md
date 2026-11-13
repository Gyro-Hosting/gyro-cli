## gyro container plans

Kapsayıcı ortam paketlerini listele

### Synopsis

Kapsayıcı ortam paketlerini listele

Sipariş edilebilir Gyro Kapsayıcı (ECS) ortam paketleri ve ön ödemeli fiyatlar. Paket = 1 ortam (size ayrılmış düğüm havuzu); `max_services` tanımlanabilecek en fazla uygulama, `max_replicas` servis başına en fazla kopya; hepsi havuzu paylaşır.

API: GET /cloud/containers/plans (operationId listContainerPlans, scope cloud_read)

```
gyro container plans [flags]
```

### Options

```
  -h, --help   help for plans
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

