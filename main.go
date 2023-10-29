package main

import (
    "crypto/ecdsa"
    "flag"
    "fmt"
    "github.com/ethereum/go-ethereum/common/hexutil"
    "github.com/ethereum/go-ethereum/crypto"
    "strings"
    "sync"
)

var cwgn sync.WaitGroup

func main() {
    // receive args
    num := flag.Int("c", 1, "cpu num")
    length := flag.Int("l", 3, "prefix length")
    flag.Parse()
    fmt.Println("cpu num:", *num, "prefix length:", *length)
    fmt.Println("you can start like this: ./main -c ", *num, " -l ", *length)
    // create wallet
    for i := 0; i < *num; i++ {
        cwgn.Add(1)
        go createWalletNew(*length)
    }

    cwgn.Wait()
}

func createWalletNew(length int) {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("createWallet panic:", err)
        }
        cwgn.Done()
    }()

    // create random wallet
    for {
        publicAddress, privateKey, err := CreateWalletRandom()
        if err != nil {
            fmt.Println("CreateWalletRandom err:", err)
            return
        }
        prefix := "0x"
        str := publicAddress[2:3]
        after := publicAddress[41:]
        strRepeat := strings.Repeat(after, length)
        for i := 0; i < length; i++ {
            prefix += str
        }

        if !strings.HasPrefix(publicAddress, prefix) && !strings.HasSuffix(publicAddress, strRepeat) {
            continue
        }

        fmt.Println("publicAddress:", publicAddress, "privateKey:", privateKey)
    }
}

func CreateWalletRandom() (publicAddress, privateKey string, err error) {
    // 生成新钱包
    pkey, err := crypto.GenerateKey()
    if err != nil {
        fmt.Println("crypto.GenerateKey err:", err)
        return "", "", err
    }
    privateKeyBytes := crypto.FromECDSA(pkey)
    privateKey = hexutil.Encode(privateKeyBytes)[2:]
    publicKey := pkey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        fmt.Println("error casting public key to ECDSA")
        return "", "", err
    }
    publicAddress = strings.ToLower(crypto.PubkeyToAddress(*publicKeyECDSA).Hex())

    return publicAddress, privateKey, nil
}
