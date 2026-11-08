## gyro dns record update

DNS Kaydını Düzenleme

### Synopsis

DNS Kaydını Düzenleme

Var olan bir DNS kaydını değiştirir. — numara ile eşleşir

API: PUT /domains/{domain}/dns-records (operationId coreUpdateDnsRecord, scope Domains/UpdateDnsRecord)

```
gyro dns record update <DOMAIN> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for update
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --identity string          Listeden gelen kayıt numarası. Liste bir numara döndürdüyse mutlaka gönderin.
      --name string              Kaydın adı.
      --priority int             Öncelik. Verilmezse kayıt, türü değişmediği sürece mevcut önceliğini korur.
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --ttl int                  Yaşam süresi. Verilmezse kayıt mevcut yaşam süresini korur.
      --type string              Kayıt türü.
      --value string             Yeni içerik.
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

* [gyro dns record](gyro_dns_record.md)	 - DNS kayıtlarını listele, ekle, değiştir, sil

