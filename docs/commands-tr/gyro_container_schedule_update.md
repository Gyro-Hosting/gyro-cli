## gyro container schedule update

Zamanlanmış görevi değiştir (yalnız verilen seçenekler değişir)

```
gyro container schedule update <ENVIRONMENT_ID> <NAME|ID> [flags] [-- COMMAND...]
```

### Examples

```
  gyro container schedule update 17000 nightly-report --cron "30 2 * * *" --timezone UTC
  gyro container schedule update 17000 nightly-report --service api -- npm run report:full
  gyro container schedule update 17000 ping-app --header "Authorization: Bearer $NEW"
```

### Options

```
      --at string               bu zamanda bir kez çalıştır: 2026-12-01T03:00 (--timezone'a göre) ya da RFC 3339
      --body string             HTTP gövdesi (metin ya da @dosya)
      --clear-headers           saklanan HTTP başlıklarını sil
      --cpu string              çalışmanın vCPU'su (varsayılan: servisle aynı)
      --cron string             cron ifadesi: dakika saat gün ay haftanın-günü, ör. "0 3 * * *" (her gün 03:00), "*/15 * * * *", "0 9 * * 1-5"; @daily / @hourly da olur
      --description string      açıklama
      --every duration          cron yerine her N dakikada, ör. 15m, 2h (en az 1m; Sunucusuz 5m)
      --header stringArray      HTTP başlığı "Ad: değer" (tekrarlanabilir; değerler gizli saklanır)
  -h, --help                    help for update
      --http-timeout duration   HTTP çağrısının zaman sınırı (1s-120s, varsayılan 30s)
      --memory int              çalışmanın belleği MB (varsayılan: servisle aynı)
      --method string           HTTP yöntemi (varsayılan GET)
      --retries int             başarısız çalışmadan sonra tekrar deneme, 0-3 (30 sn × deneme arayla)
      --service string          "--" sonrasındaki komutu bu servisin imajı, değişkenleri ve gizli değerleriyle çalıştır
      --time-limit duration     bir çalışmayı bu süre sonunda durdur (10s-6h, varsayılan 30m)
      --timezone string         zamanlamanın IANA saat dilimi (varsayılan Europe/Istanbul), ör. UTC
      --url string              bunun yerine bu adresi çağır (http:// ya da https://)
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

* [gyro container schedule](gyro_container_schedule.md)	 - Zamanlanmış görevleri sonraki ve son çalışmalarıyla listele

