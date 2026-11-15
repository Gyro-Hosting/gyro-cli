## gyro auth logout

Anahtarı bu profilden sil ve sunucuda iptal et

### Synopsis

'gyro auth login' ile oluşturulan anahtar sunucuda iptal edilir ve profilden silinir.
'gyro configure' ile yapıştırdığınız anahtar yalnız yerelden silinir (başka yerde kullanılıyor olabilir); onu da iptal etmek için --revoke.

```
gyro auth logout [flags]
```

### Options

```
  -h, --help       help for logout
      --keep-key   anahtarı yalnız yerelden sil, iptal etme
      --revoke     gyro auth login ile oluşturulmamış anahtarı da iptal et
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

* [gyro auth](gyro_auth.md)	 - Tarayıcıyla giriş yap, CLI oturumunu göster ya da kapat

