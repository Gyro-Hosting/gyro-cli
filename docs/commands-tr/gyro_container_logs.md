## gyro container logs

Ortamın ya da bir servisin loglarında ara, canlı izle

### Synopsis

Kapsayıcılarınızın stdout / stderr'e yazdıklarını eskiden yeniye gösterir. Servis verilmezse ortamın tüm servisleri
aranır. Filtreler: --filter (satır içerir, büyük/küçük harf duyarsız), --exclude, --level (error|warn|info|debug) ve
JSON / logfmt satırlarda --where alan koşulları (status>=500, method=POST, path~/api; birden fazlası için bayrağı tekrarlayın).
--follow yeni satırları basmayı sürdürür (2 sn'de bir, Ctrl-C durdurur). --all tüm aralığı geriye doğru sayfalar
(en çok 100.000 satır). Ne kadar geriye gidebileceğiniz log paketinize bağlıdır (ücretsiz: 7 gün; gyro container log-usage).

```
gyro container logs <ENVIRONMENT_ID> [SERVICE] [flags]
```

### Examples

```
  gyro container logs 17000 api --since 30m
  gyro container logs 17000 --level error --since 24h
  gyro container logs 17000 api --where 'status>=500' --where 'path~/api' --since 2h
  gyro container logs 17000 api -f
```

### Options

```
      --all                 tüm aralığı geriye doğru sayfala (en çok 100.000 satır)
      --end string          bitiş (RFC3339 ya da geçen süre)
      --exclude string      satır bu metni içermez
      --filter string       satır bu metni içerir (büyük/küçük harf duyarsız)
  -f, --follow              yeni satırları basmayı sürdür (Ctrl-C durdurur)
  -h, --help                help for logs
      --level string        error | warn | info | debug
      --limit int           sayfa başına satır (varsayılan 300, en çok 5000)
      --parse string        json | logfmt
      --previous            servisin son çöken kopyasının çıktısı
      --service string      yalnız bu servis (ikinci argümanla aynı)
      --since string        şimdi biten aralık: 30m, 6h, 7d (varsayılan 1h; --follow ile 15m)
      --start string        başlangıç (RFC3339 ya da geçen süre, ör. 3d); bir sorgu en çok 31 gün
      --where stringArray   JSON / logfmt satırlarda alan koşulu: status>=500, method=POST, path~/api (tekrarlanabilir)
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

