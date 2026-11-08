## gyro completion

Kabuk tamamlama betiği üret

### Synopsis

Tamamlamayı yükleyin:
  bash:  source <(gyro completion bash)        # ya da: gyro completion bash > /etc/bash_completion.d/gyro
  zsh:   gyro completion zsh > "${fpath[1]}/_gyro"
  fish:  gyro completion fish > ~/.config/fish/completions/gyro.fish
  pwsh:  gyro completion powershell | Out-String | Invoke-Expression
Sunucu, SSH anahtarı ve diğer kimlikler API'den tamamlanır.

```
gyro completion bash|zsh|fish|powershell [flags]
```

### Options

```
  -h, --help   help for completion
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

* [gyro](gyro.md)	 - Gyro Hosting komut satırı aracı

