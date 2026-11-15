## gyro container log-stats

Log özeti

### Synopsis

Log özeti

Zaman içinde satır sayısı (`series`, en çok 60 dilim) ve bir gruplamanın en çok görülen değerleri: `by=service` (varsayılan), `level` ya da JSON / logfmt satırların bir alanı (`by=status`). Aramayla aynı filtreler; aralık en çok 7 gün.

API: GET /cloud/containers/{id}/logs/stats (operationId containerLogStats, scope services_read)

```
gyro container log-stats <ID> [flags]
```

### Options

```
      --by string           service | level | <field>
      --end string          RFC3339 bitiş (varsayılan şimdi)
      --exclude string      satır bu metni içermez
      --filter string       satır bu metni içerir (büyük/küçük harf duyarsız, en çok 200)
  -h, --help                help for log-stats
      --level string        error | warn | info | debug (satırdaki seviye kelimesi)
      --parse where         json (where varsa varsayılan) | logfmt
      --service string      yalnız bu servis (varsayılan: ortamın tüm servisleri)
      --since-minutes int   şimdi biten aralık (varsayılan 60); paket saklama süresiyle sınırlı
      --start string        RFC3339 başlangıç; bir sorgu en çok 31 gün kapsar
      --top int             groups (default 10)
      --where ;             JSON / logfmt satırlarda alan koşulları, ; ile: `status>=500; path~/api; method=POST`
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

