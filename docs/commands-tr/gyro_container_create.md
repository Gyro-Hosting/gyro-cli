## gyro container create

Kapsayıcı ortamı oluştur (toplu)

### Synopsis

Kapsayıcı ortamı oluştur (toplu)

Bir paketten 1..N ortamı TEK siparişte açar (`quantity`, varsayılan 1, istek başına en çok 25). Ön ödemeli: sipariş 0'dır, kullanım saatlik bakiyeden düşer; bakiye her ortamın ilk 24 saatini karşılamalıdır, yeni hesaplar ilk günlerde sınırlıdır. `name` (`{n}` ile) ve `tags` her ortamı adlandırır; `project_id` projeye koyar. Teklif için `"dry_run": true`; gerçek sipariş `"confirm": true` ve `Idempotency-Key` başlığı ister.

API: POST /cloud/containers (operationId createContainerEnvironments, scope cloud_order)

```
gyro container create [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for create
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --name string              ad (takma ad); toplu oluşturmada {n} = sıra no
      --plan string              GET /cloud/containers/plans paketi
      --product-id int           plan yerine ürün kimliği
      --project-id int           proje (0 = Genel)
      --quantity int             tek siparişte açılacak ortam (--count), varsayılan 1, en çok 25
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --tags strings             etiketler (en çok 10)
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

