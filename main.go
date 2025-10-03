package main

import (
	"log"
	"strings"
)

var words []string

func main() {
	found := false
	s, _ := generateSeed()
	a := getAddressFromSeed(s)
	counter := 0

	suffix := "Prag"

	go func() {
		for !found {
			s, _ = generateSeed()
			a = getAddressFromSeed(s)

			// if strings.HasSuffix(strings.ToLower(a), suffix) {
			if strings.HasSuffix(a, suffix) {
				found = true

				log.Println(a)
				log.Println(s)
			}

			counter++
			log.Println(counter)
		}
	}()

	go func() {
		for !found {
			s, _ = generateSeed()
			a = getAddressFromSeed(s)

			// if strings.HasSuffix(strings.ToLower(a), suffix) {
			if strings.HasSuffix(a, suffix) {
				found = true

				log.Println(a)
				log.Println(s)
			}

			counter++
			log.Println(counter)
		}
	}()

	go func() {
		for !found {
			s, _ = generateSeed()
			a = getAddressFromSeed(s)

			// if strings.HasSuffix(strings.ToLower(a), suffix) {
			if strings.HasSuffix(a, suffix) {
				found = true

				log.Println(a)
				log.Println(s)
			}

			counter++
			log.Println(counter)
		}
	}()

	go func() {
		for !found {
			s, _ = generateSeed()
			a = getAddressFromSeed(s)

			// if strings.HasSuffix(strings.ToLower(a), suffix) {
			if strings.HasSuffix(a, suffix) {
				found = true

				log.Println(a)
				log.Println(s)
			}

			counter++
			log.Println(counter)
		}
	}()

	go func() {
		for !found {
			s, _ = generateSeed()
			a = getAddressFromSeed(s)

			// if strings.HasSuffix(strings.ToLower(a), suffix) {
			if strings.HasSuffix(a, suffix) {
				found = true

				log.Println(a)
				log.Println(s)
			}

			counter++
			log.Println(counter)
		}
	}()

	go func() {
		for !found {
			s, _ = generateSeed()
			a = getAddressFromSeed(s)

			// if strings.HasSuffix(strings.ToLower(a), suffix) {
			if strings.HasSuffix(a, suffix) {
				found = true

				log.Println(a)
				log.Println(s)
			}

			counter++
			log.Println(counter)
		}
	}()

	go func() {
		for !found {
			s, _ = generateSeed()
			a = getAddressFromSeed(s)

			// if strings.HasSuffix(strings.ToLower(a), suffix) {
			if strings.HasSuffix(a, suffix) {
				found = true

				log.Println(a)
				log.Println(s)
			}

			counter++
			log.Println(counter)
		}
	}()

	go func() {
		for !found {
			s, _ = generateSeed()
			a = getAddressFromSeed(s)

			// if strings.HasSuffix(strings.ToLower(a), suffix) {
			if strings.HasSuffix(a, suffix) {
				found = true

				log.Println(a)
				log.Println(s)
			}

			counter++
			log.Println(counter)
		}
	}()

	go func() {
		for !found {
			s, _ = generateSeed()
			a = getAddressFromSeed(s)

			// if strings.HasSuffix(strings.ToLower(a), suffix) {
			if strings.HasSuffix(a, suffix) {
				found = true

				log.Println(a)
				log.Println(s)
			}

			counter++
			log.Println(counter)
		}
	}()

	func() {
		for !found {
			s, _ = generateSeed()
			a = getAddressFromSeed(s)

			// if strings.HasSuffix(strings.ToLower(a), suffix) {
			if strings.HasSuffix(a, suffix) {
				found = true

				log.Println(a)
				log.Println(s)
			}

			counter++
			log.Println(counter)
		}
	}()

	// for !strings.HasSuffix(strings.ToLower(a), "satsh") {

}
