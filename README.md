# eth_wallet
Generate wallet addresses with consecutive prefixes or suffixes

# install go environment
```
# Mac OS or Linux 
# Download 
wget https://go.dev/dl/go1.21.3.darwin-amd64.tar.gz # (Interl or Amd) or
wget https://go.dev/dl/go1.21.3.darwin-arm64.tar.gz #(Apple ARM) or
wget https://go.dev/dl/go1.21.3.linux-amd64.tar.gz #(Linux)
# install
tar -C /usr/local -xzf go1.21.3.darwin-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

# Windows 
# download and install, address https://go.dev/dl/go1.21.3.windows-amd64.msi
```

# direct run
```
go run main.go -c 2 -l 4
```

# build
```
# Mac OS or Linux
export GOOS=darwin
export GOARCH=amd64
go build main.go

unset GOOS
unset GOARCH

# Windows
go build main.go -o main.exe
```

# Execute executable file
```
./main -c 2 -l 4
or
./main.exe -c 2 -l 4
```