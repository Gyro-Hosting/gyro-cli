## gyro container run

Servisin imajı, değişkenleri ve gizli değerleriyle tek seferlik görev çalıştır (migration, seed)

### Synopsis

Mevcut bir servisin imajı (çalışan sürümün digest'ine sabitlenir), değişkenleri ve gizli değerleriyle bir komutu
bir kez çalıştırır, sonra durur. "--" sonrası komuttur; imajın ENTRYPOINT'inin yerine geçer ve liste olarak gider
(kabuk yok: kabuk özellikleri için sh -c "a && b" yazın). Komut yoksa imajın kendi komutu çalışır.
--follow çıktıyı akıtır; --wait sessizce bekler. İkisinde de CLI görevin çıkış koduyla biter (124 = zaman sınırı
doldu, 130 = durduruldu). Düğümlü paketlerde ücretsizdir (düğümlerinizi kullanır); Sunucusuz'da görev vCPU-saniye /
GB-saniye ile görev (task) gibi ücretlenir (fatura satırı "<servis>-job"). Ortam başına eşzamanlı görev sayısı
pakete göre 2-3; bkz. "gyro limits".

```
gyro container run <ENVIRONMENT_ID> --service <NAME> [flags] [-- COMMAND [ARG...]]
```

### Examples

```
  gyro container run 17000 --service api --follow -- npm run migrate
  gyro container run 17000 --service web --wait --time-limit 10m -- sh -c "php artisan migrate --force && php artisan db:seed --force"
  gyro container run 17000 --service api --cpu 1 --memory 1024 -- python manage.py migrate
```

### Options

```
      --cpu string            görevin vCPU'su (varsayılan: servisle aynı)
      --follow                görev bitene kadar çıktıyı akıt (--wait içerir)
  -h, --help                  help for run
      --memory int            görevin belleği MB (varsayılan: servisle aynı)
      --service string        imajı, değişkenleri ve gizli değerleri kullanılan servis (zorunlu)
      --time-limit duration   görevi bu süre sonunda durdur (10s-6h, varsayılan 30m)
      --wait                  görev bitene kadar bekle; onun çıkış koduyla bit
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

* [gyro container](gyro_container.md)	 - Gyro Kapsayıcı (ECS), İstanbul: servisler, sürümler, loglar, alan adları

