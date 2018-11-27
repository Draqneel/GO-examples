package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

//READ DICTIONARY.TXT !!!
func main(){
	if len(os.Args) < 4 {
		get_error("OS args less then 4")
	}
	check_mode(os.Args[1])
	hack_password(os.Args[1], os.Args[2], os.Args[3])
}

func check_mode(mode string){
	if mode != "md5" && mode != "sha256" && mode != "sha512"{
		get_error("crypto-hash func is not found")
	}
}


func hack_password(mode string, dictionary,static_hash string)  {
	file, err := os.Open(dictionary)
	check_error(err)
	defer file.Close()

	var reader *bufio.Reader = bufio.NewReader(file)
	var password, dynamic_hash string

	fmt.Println("----------START----------")
	for {
		password, _ = reader.ReadString('\n')
		password = strings.Replace(password,"\n","",-1)
		dynamic_hash = encrypt(mode, password)
		if static_hash == dynamic_hash {
			fmt.Println("----------(!!!)---------")
			fmt.Println("SUCCESS:", password)
			fmt.Println("----------(!!!)---------")
			break
		} else {
			fmt.Println("[FAILURE]:", password)
		}
	}
}

func encrypt(crypt string, text string) string {
	if crypt == "md5" {
		hash := md5.New()
		hash.Write([]byte(text))
		return hex.EncodeToString(hash.Sum(nil))
	} else if crypt == "sha256" {
		hash := sha256.New()
		hash.Write([]byte(text))
		return hex.EncodeToString(hash.Sum(nil))
	} else {
		hash := sha512.New()
		hash.Write([]byte(text))
		return hex.EncodeToString(hash.Sum(nil))
	}
}

func check_error(err error)  {
	if (err != nil) {
		get_error(err.Error())
	}
}

func get_error(err string)  {
	fmt.Println("error:",err)
	os.Exit(1)
}
