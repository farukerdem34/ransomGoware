# Legal Disclaimer
**Do not use** it in a way that could harm people or institutions! Unauthorized and damaging use of this software is **illegal**.

### Configurations
- `RANDOM=true` *Encryption Only*
- Unix OSes `ROOT='/path/to/root'` or Windows `ROOT='C:\path\to\root'`
- If you use your own key, *recommended*, instead of random generated AES encryption key. Change the key values. `YOUR-32-BIT-KEY-STRING`
## Interpreted Usage
```bash
go run encrypt.go
```
```bash
go run decrypt.go
```
## Compiled Usage
```bash
go build encrypt.go
./encrypt
```
```bash
go build decrypt.go
./decrypt
```
