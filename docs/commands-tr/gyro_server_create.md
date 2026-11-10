## gyro server create

Yeni bulut sunucu siparişi (bakiyeden ödenir)

### Synopsis

Bulut sunucu siparişi verir. Önce API'den fiyat teklifi gösterilir ve onay istenir; ödeme kart değil
hesap bakiyesinden yapılır. Komuta otomatik tekrar güvenliği anahtarı eklenir, tekrar denenen komut
iki kez sipariş vermez.

--plan: plan adı (standart-2-4, pro-2-8 ...) ya da ürün kimliği ('gyro plan list').
--location: bölge (eu-de-1, 'gyro location list') ya da şehir (falkenstein).

```
gyro server create [HOSTNAME] [flags]
```

### Examples

```
  gyro server create web-1 --plan pro-2-8 --location falkenstein --image ubuntu-24.04 --ssh-key mykey --wait
  gyro server create --plan standart-2-4 --location eu-de-1 --image debian-12 --ssh-key ~/.ssh/id_ed25519.pub --dry-run
  gyro server create db-1 --plan standart-2-4 --location eu-de-2 --image ubuntu-24.04 --user-data @cloud-init.yaml --yes -o json
```

### Options

```
      --app string               kurulacak tek tık uygulama (API v2 gerekir)
      --cycle string             monthly|quarterly|semiannually|annually (default "monthly")
      --data string              ek istek gövdesi JSON / @dosya
      --dry-run                  yalnız fiyat teklifini göster
  -h, --help                     help for create
      --hostname string          sunucu adı (ya da argüman olarak verin)
      --idempotency-key string   tekrar güvenliği anahtarı (varsayılan: rastgele)
      --image string             işletim sistemi imajı, ör. ubuntu-24.04 (zorunlu)
      --location string          bölge (eu-de-1) ya da şehir
      --plan string              plan adı ya da ürün kimliği (zorunlu)
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --ssh-key strings          SSH anahtarı kimliği, adı ya da .pub dosyası (yeniyse yüklenir)
      --user-data string         cloud-init verisi, metin ya da @dosya
      --wait                     sunucu çalışana kadar bekle
      --wait-timeout duration    --wait için en uzun süre (default 15m0s)
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

* [gyro server](gyro_server.md)	 - Bulut sunucular: oluşturma, güç, snapshot, SSH ...

