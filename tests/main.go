package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"go.dedis.ch/kyber/v3/encrypt/ecies"
	"go.dedis.ch/kyber/v3/group/edwards25519"
	"go.dedis.ch/kyber/v3/util/random"
)

//options for what to do enum
const (
	EXEC_DECRYPT   = 0
	EXEC_ENCRYPT   = 1
	EXEC_SPEEDTEST = 2
)

//options for where to put output enum
const (
	IO_TERMINAL = 0
	IO_CSV      = 1
	IO_FOLDER   = 2
)

func main() {

	argCount := len(os.Args[1:])
	k1 := 4
	n1 := 5
	var s []string
	var ciphertext string

	toDo := -1
	filePath := ""
	filetype := IO_TERMINAL

	numspeedtest := 100
	for i := 1; i < argCount+1; i++ {
		if os.Args[i] == "-encrypt" {
			toDo = EXEC_ENCRYPT
		}
		if os.Args[i] == "-decrypt" {
			toDo = EXEC_DECRYPT
		}
		if len(os.Args[i]) > 2 {
			if os.Args[i][:2] == "-k" {
				k1, _ = strconv.Atoi(os.Args[i][2:len(os.Args[i])])
			}
			if os.Args[i][:2] == "-n" {
				n1, _ = strconv.Atoi(os.Args[i][2:len(os.Args[i])])
			}
			if os.Args[i][:2] == "-s" {
				str := os.Args[i][2:len(os.Args[i])]
				str = strings.Trim(str, "'")
				s = append(s, str)
			}
			if os.Args[i][:2] == "-c" {
				str := os.Args[i][2:len(os.Args[i])]
				str = strings.Trim(str, "'")

				bytes, _ := hex.DecodeString(str)
				ciphertext = string(bytes[:])
			}
		}
		if len(os.Args[i]) > 4 && os.Args[i][:4] == "-csv" {
			filePath = os.Args[i][4:len(os.Args[i])]
			filePath = strings.Trim(filePath, "'")
			filetype = IO_CSV
		}
		if len(os.Args[i]) > 7 && os.Args[i][:7] == "-folder" {
			filePath = os.Args[i][7:len(os.Args[i])]
			filePath = strings.Trim(filePath, "'")
			filetype = IO_FOLDER

		}

		if len(os.Args[i]) > 8 && os.Args[i][:10] == "-speedtest" {
			toDo = EXEC_SPEEDTEST
			numspeedtest, _ = strconv.Atoi(os.Args[i][10:len(os.Args[i])])

		}

	}
	if toDo == EXEC_SPEEDTEST {
		DoSpeedTests(numspeedtest, n1, k1)
		os.Exit(0)
	}
	if len(s) == 0 && (filetype != IO_FOLDER && filePath != "") {
		fmt.Printf("Cannot execute, no input. Pass strings to program with the format -sFoo")
		os.Exit(0)
	}
	if toDo == -1 {
		fmt.Printf("No command given. Pass either -encrypt or -decrypt")
		os.Exit(0)
	}
	if k1 > n1 {
		fmt.Printf("Cannot do this, as k greater than n\n")
		os.Exit(0)
	}

	suite := edwards25519.NewBlakeSHA256Ed25519()
	private := suite.Scalar().Pick(random.New())
	var shares []string

	if toDo == EXEC_ENCRYPT {
		keystruct := suite.Scalar()
		_, ciphertext, keystruct = eceis([]byte(s[0]), private)
		structMarshalled, _ := keystruct.MarshalBinary()
		shares = SSSaaS_CreateShares(structMarshalled, n1, k1)

		if filetype == IO_TERMINAL {
			fmt.Printf("Ciphertext:\n'%x'\n", ciphertext)

			fmt.Printf("Shares:\n")
			for i := 0; i < len(shares); i++ {
				fmt.Printf("'%s',\n", shares[i])
			}
			fmt.Printf("\n")

		}
		if filetype == IO_FOLDER {
			AddToFolder(filePath, ciphertext, shares)
		}

		if filePath != "" && filetype == IO_CSV {
			AddToCSV(filePath, ciphertext, shares)
		}
	}
	if toDo == EXEC_DECRYPT {
		if filePath != "" && filetype == IO_FOLDER {
			ciphertext, s = TakeFromFolder(filePath)
		}

		recovered := SSSaaS_CombineShares(s)
		rec := suite.Scalar().SetBytes(([]byte(recovered)))
		//marshalled := rec.SetBytes([]byte(keyst))
		plaintext, err := ecies.Decrypt(suite, rec, []byte(ciphertext), suite.Hash)
		if err != nil {
			fmt.Printf("\n%s\n", err)
		}
		fmt.Printf("plaintext: %s\n", plaintext)

	}
}
func DoSpeedTests(numTests int, n1 int, k1 int) {
	fmt.Printf("Doing %d of tests with %d keys which require %d to recombine\n", numTests, n1, k1)
	var data [][]byte

	for i := 0; i < numTests; i++ {
		bytes := make([]byte, 64)
		rand.Read(bytes)

		data = append(data, bytes)
	}
	times := make(map[string]time.Duration)
	fmt.Printf("\n\n\n")
	times["codahale"] = test_codahale(data, n1, k1)
	fmt.Printf("\n\n\n")
	times["sprenkels"] = test_sprenkels(data, n1, k1)
	fmt.Printf("\n\n\n")
	times["SSSaas"] = test_SSSaaS(data, n1, k1)
	fmt.Printf("\n\n\n")
	times["kryptology"] = test_Kryptology(data, n1, k1)
	fmt.Printf("\n\n\n")

	fmt.Printf("Total elapsed: \n")

	for k, v := range times {
		fmt.Printf("	%s: %s\n", k, v)
	}

}
func AddToCSV(filepath string, cipher string, share []string) {
	fi, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		fmt.Printf("Could not open file")
		return
	}
	w := bufio.NewWriter(fi)
	var str string
	cipher = fmt.Sprintf("%x", cipher)
	str += cipher + ","
	for i := 0; i < len(share); i++ {
		str += share[i] + ","
	}
	str = str[:len(str)-2]
	str += "\n"

	w.WriteString((str))
	w.Flush()
	fi.Close()
	fmt.Printf("Written")

}
func AddToFolder(filepath string, cipher string, shares []string) {

	os.Mkdir(filepath, 0777)
	cipher = fmt.Sprintf("%x", cipher)

	for i := 0; i < len(shares); i++ {
		filepathAppend := filepath + "/" + strconv.Itoa(i)
		fmt.Printf("%s\n", filepathAppend)

		fi, err := os.OpenFile(filepathAppend, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0660)
		w := bufio.NewWriter(fi)

		if err != nil {
			fmt.Printf("Could not write to file")
			return
		}

		appendedSecret := cipher + "," + shares[i]
		w.WriteString(appendedSecret)
		w.Flush()
		fi.Close()
	}
}
func TakeFromFolder(filepath string) (string, []string) {
	dir, _ := os.Open(filepath)
	files, _ := dir.ReadDir(0)

	var shares []string
	var cipher string

	for i := range files {
		if files[i].Name() == ".DS_Store" {
			continue
		}
		if files[i].IsDir() {
			continue
		}
		fullPath := filepath + "/" + files[i].Name()

		fi, errFile := os.OpenFile(fullPath, os.O_RDONLY, 444)
		if errFile != nil {
			fmt.Printf("Could not open file: %s", errFile)
		}
		r := bufio.NewReader(fi)

		str, _ := r.ReadString('\n')

		split := strings.Split(str, ",")

		shares = append(shares, split[1])
		cipher = split[0]
	}

	bytes, _ := hex.DecodeString(cipher)
	cipher = string(bytes[:])

	return cipher, shares
}
