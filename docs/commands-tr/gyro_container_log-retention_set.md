## gyro container log-retention set

Log paketini değiştir

### Synopsis

Log paketini değiştir

`days` 7 | 30 | 90 | 365 ve `gb` (aylık hacim) `GET .../logs/plans` listesinden; `{"free": true}` ücretsiz pakete döner. Mevcuttan KISA saklama süresi eski logları kalıcı olarak siler (yanıtta `data_loss`, `deletes_logs_older_than_days` ve `warning` metni; `GET .../logs/plans` bu paketleri `data_loss` ile işaretler). Ücretsizden ücretli pakete her geçiş en az 6 saat ücretlenir (`minimum_hours`); daha önce ücretsize dönülürse kalan saatler hemen düşülür (`minimum_completion_hours`, saatlik ayrıntı ve CSV'de "asgari süre tamamlaması", ek IPv4 gibi). Ücretli paket ön ödemeli bakiyeden saatlik düşer (fatura satırı "Log saklama") ve `"confirm": true` ister. Platformda paket için yer yoksa hata döner. Loglar kapsayıcılarınızın stdout / stderr'e yazdıklarıdır. Her sorguyu platform ortamınız sabitlenmiş olarak kurar; ham sorgu dili yoktur. **Ücretsiz paket**: 7 gün, aylık 5 GB. Ücretli paketler (30 / 90 / 365 gün x aylık 10 / 50 / 100 / 300 GB, ön ödemeli saatlik; `GET /cloud/containers/{id}/logs/plans`) logları daha uzun saklar. Aylık hacmi aşınca ek ücret çıkmaz: saklanan hacim GB x gün / 30 ile sınırlıdır, en eski loglar daha erken silinir ve %80 / %100'de `log_quota` bildirimi gelir. Ortam iptal edilince logları hemen silinir. En iyi çaba hizmetidir, kaybolan log satırı için tazminat yoktur.

API: PUT /cloud/containers/{id}/logs/retention (operationId setContainerLogPackage, scope services_write)

```
gyro container log-retention set <ID> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --days int                 days (7|30|90|365)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
      --free                     free
      --gb int                   gb
  -h, --help                     help for set
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
```

### Options inherited from parent commands

```
      --account gyro account list   panel kullanıcısı olduğunuz başka bir hesapta işlem yap (gyro account list id'si; "self" = kendi hesabınız; ortam GYRO_ACCOUNT, profil varsayılanı: `gyro account use`)
      --api-url string              API taban adresi (ortam GYRO_API_URL)
      --debug                       HTTP isteklerini stderr'e yaz (sırlar maskeli)
      --lang string                 mesaj dili: tr ya da en (varsayılan yerel ayar)
      --no-color                    renkleri kapat (NO_COLOR da olur)
      --no-header                   tablo başlığını yazma
  -o, --output string               çıktı biçimi: table, json, yaml, text
  -p, --profile string              kullanılacak profil (ortam GYRO_PROFILE)
      --query string                sonuca uygulanacak JMESPath sorgusu
  -q, --quiet                       yalnız kimlikleri yaz
      --timeout duration            toplam zaman aşımı, ör. 30s, 10m
      --wide                        tüm tablo sütunlarını göster
  -y, --yes                         onay sorma
```

### SEE ALSO

* [gyro container log-retention](gyro_container_log-retention.md)	 - Log paketini değiştir

