# ykt
Use a Yubikey as TOTP.

```

./ go build -ldflags "-s -w" -o ykt main.go

./ykt list
./ykt get name
./ykt set name secret
./ykt del name

```

### License

released under the GPL 3 license.
