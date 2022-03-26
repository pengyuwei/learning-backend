package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type cmdPara struct {
	src  string
	dest string
	key  string
	mode bool
}

func readFileAll(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	// fmt.Println(string(content))
	return content, err
}

func writeFile(content []byte, filename string) {
	err := ioutil.WriteFile(filename, content, 0644)
	if err != nil {
		panic(err)
	}
}

func RSAEncrypt(src []byte, filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)

	// 从数据中找出pem格式的块
	block, _ := pem.Decode(buf)
	if block == nil {
		return nil, err
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		log.Fatalf("got unexpected key type: %T", publicKey)
	}

	// 公钥加密:EncryptPKCS1v15用来加密密钥，其他信息使用EncryptOAEP加密，长度有限制
	label := []byte("orders")
	result, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, rsaKey, src, label)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from encryption: %s\n", err)
	}
	return result, err
}

func RSADecrypt(src []byte, filename string) ([]byte, error) {
	// 根据文件名读出内容
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)

	// 从数据中解析出pem块
	block, _ := pem.Decode(buf)
	if block == nil {
		return nil, err
	}

	// privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// rsaKey, ok := privateKey.(*rsa.PrivateKey)
	// if !ok {
	// 	log.Fatalf("got unexpected key type: %T", privateKey)
	// }

	// 私钥解密
	// fmt.Println("Decrypt content: ", src)
	label := []byte("orders")
	result, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, src, label)
	if err != nil {
		fmt.Println("DecryptOAEP error", err)
		return nil, err
	}
	return result, nil
}

/*
openssl genrsa -out privatekey.pem 2048
openssl rsa -in privatekey.pem -out publickey.pem -pubout
./crypt -e -s src.log -d dest.log -k publickey.pem
./crypt -s dest.log -d output.log -k privatekey.pem
*/
func main() {
	var para cmdPara
	flag.StringVar(&para.src, "s", "src.log", "source file")
	flag.StringVar(&para.dest, "d", "dest.log", "dest file")
	flag.StringVar(&para.key, "k", "publickey.pem", "public key or private key.pem")
	flag.BoolVar(&para.mode, "e", false, "encrypt or decrypt")
	flag.Parse()
	if para.mode {
		content, _ := readFileAll(para.src)
		cipherText, err := RSAEncrypt(content, para.key)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(string(cipherText))
		writeFile(cipherText, para.dest)
		fmt.Println(para.dest, " done.")
	} else {
		content, _ := readFileAll(para.src)
		plainText, err := RSADecrypt(content, para.key)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(string(plainText))
		writeFile(plainText, para.dest)
		fmt.Println(para.dest, " done.")
	}
}
