## gyro domain whois-profile delete

Profili Silme

### Synopsis

Profili Silme

Kayıtlı bir kişi profilini siler. — hesap geneli

API: DELETE /domains/whois-profiles/{pid} (operationId coreDeleteWhoisProfile, scope Domains/DeleteWhoisProfile)

```
gyro domain whois-profile delete <PID> [flags]
```

### Options

```
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for delete
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
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

* [gyro domain whois-profile](gyro_domain_whois-profile.md)	 - Kayıtlı Profilleri Listeleme

