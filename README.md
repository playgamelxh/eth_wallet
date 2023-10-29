# eth_wallet
Generate wallet addresses with consecutive prefixes or suffixes
```
Usage of ./main:
  -c int
        The number of CPU used to generate (default 1)
  -l int
        The length of the prefix or suffix (default 3)
        
./main -c 2 -l 4
./main.exe -c 2 -l 4
```

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
# Donation address
0x79032bFC86a75E58e265B046cFe0d25555555555