## gyro container log-query

Loglarda ara (JSON gövde)

### Synopsis

Loglarda ara (JSON gövde)

GET aramayla aynı, parametreler gövdede (`where` liste). Loglar kapsayıcılarınızın stdout / stderr'e yazdıklarıdır. Her sorguyu platform ortamınız sabitlenmiş olarak kurar; ham sorgu dili yoktur. **Ücretsiz paket**: 7 gün, aylık 5 GB. Ücretli paketler (30 / 90 / 365 gün x aylık 10 / 50 / 100 / 300 GB, ön ödemeli saatlik; `GET /cloud/containers/{id}/logs/plans`) logları daha uzun saklar. Aylık hacmi aşınca ek ücret çıkmaz: saklanan hacim GB x gün / 30 ile sınırlıdır, en eski loglar daha erken silinir ve %80 / %100'de `log_quota` bildirimi gelir. Ortam iptal edilince logları hemen silinir. En iyi çaba hizmetidir, kaybolan log satırı için tazminat yoktur.

API: POST /cloud/containers/{id}/logs/query (operationId queryContainerLogs, scope services_write)

```
gyro container log-query <ID> [flags]
```

### Options

```
      --after tail_cursor        live tail cursor (tail_cursor / `cursor`): lines after it, oldest first; poll every 2-3 s
      --before next_cursor       page cursor (next_cursor of the previous page): older lines
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
      --end string               RFC3339 end (default now)
      --exclude string           line does not contain this text
      --filter string            line contains this text (case-insensitive, max 200)
  -h, --help                     help for log-query
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --level string             error | warn | info | debug (level word in the line)
      --param-limit int          lines (default 300, max 5000)
      --parse where              json (default when where is given) | logfmt
      --service string           only this service (default: all services of the environment)
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --since-minutes int        window ending now (default 60); limited to the package retention
      --start string             RFC3339 start (instead of since_minutes); one query spans at most 31 days
      --where ;                  field conditions on JSON / logfmt lines, separated by ;: `status>=500; path~/api; method=POST` (ops = != > >= < <= =~ ~)
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

