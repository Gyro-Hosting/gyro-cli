## gyro domain forwarding set

Adres Yönlendirmesi Kurma

### Synopsis

Adres Yönlendirmesi Kurma

Alan adına gelen ziyaretçiyi başka bir adrese gönderir. — ek hizmet gerekir

API: PUT /domains/{domain}/forwarding (operationId coreUpdateUrlForwarding, scope Domains/UpdateUrlForwarding)

```
gyro domain forwarding set <DOMAIN> [flags]
```

### Options

```
      --data string              istek gövdesi JSON, @dosya ya da @- (stdin)
      --dry-run                  hiçbir şeyi değiştirmeden teklifi / isteği göster
  -h, --help                     help for set
      --idempotency-key string   tekrar güvenliği anahtarı (para işlemlerinde otomatik)
      --method int               Yönlendirme türü. Kalıcı olan varsayılandır; başka bir değer ona düşer.
      --set stringArray          ek gövde alanı anahtar=değer (tekrarlanabilir)
      --url string               Tam hedef adres. Şeması yönlendirmenin şeması olur; yazılmazsa güvensiz şema kullanılır.
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

* [gyro domain forwarding](gyro_domain_forwarding.md)	 - Adres Yönlendirmesini Okuma

