package main

import (
	"bufio"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/mr-tron/base58"
	wavesplatform "github.com/wavesplatform/go-lib-crypto"
	"github.com/wavesplatform/gowaves/pkg/crypto"
	"github.com/wavesplatform/gowaves/pkg/proto"
)

func generateSeed() (seed string, encoded string) {
	if len(words) == 0 {
		lines, err := urlToLines(SeedWordsURL)
		if err != nil {
			log.Println(err.Error())
		}

		// for _, line := range lines {
		// 	words = append(words, line)
		// }
		words = append(words, lines...)
	}
	seed = ""
	encoded = ""

	for i := 1; i <= 15; i++ {
		seed += words[getRandNum()]
		if i < 15 {
			seed += " "
		}
	}

	data := []byte(seed)
	encoded = base58.Encode(data)

	return seed, encoded
}

func getRandNum() int {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 2047
	rn := rand.Intn(max-min+1) + min
	return rn
}

func urlToLines(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return linesFromReader(resp.Body)
}

func linesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func getAddressFromSeed(seed string) string {
	as := ""

	c := wavesplatform.NewWavesCrypto()
	sd := wavesplatform.Seed(seed)
	pair := c.KeyPair(sd)

	pk := crypto.MustPublicKeyFromBase58(string(pair.PublicKey))
	a, err := proto.NewAddressFromPublicKey(55, pk)
	if err != nil {
		log.Println(err.Error())
	}
	as = a.String()

	return as
}
