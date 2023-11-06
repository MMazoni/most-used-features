# MUF - Most Used Features

```
curl -LOk https://github.com/MMazoni/most-used-features/archive/refs/heads/main.zip
unzip main.zip
cd most-used-features-main
chmod +x ./bin/muf
./bin/muf --help
```

### Get the most used features using the apache logs

```
./bin/muf access
```

### Catalog the CSRF errors

```
./bin/muf csrf
```

### Build

```
go build -o bin/muf
```
