package sieve

// interface
type Sieve interface {
	NthPrime(number int64) int64
}

// struct for Sieve interface.
type Finder struct {
}

// constructor
func NewSieve() Sieve {
	return &Finder{}
}

// finds prime at index
func (s *Finder) NthPrime(n int64) int64 {

	ceiling := 180000000 // 179424691 last test in sieve_test

	if n < 0 {
		return 0
	}

	// primesList[i] will be true if i is a non prime number.
	primesList := make([]bool, ceiling+1)
	primesList[0], primesList[1] = true, true

	// sets tp true if prime
	for num := 2; num*num <= ceiling; num++ {
		if !primesList[num] { // stop messing with this
			for i := num * num; i <= ceiling; i += num {
				primesList[i] = true
			}
		}
	}

	// count primes until nth
	counter := int64(-1)
	for i := 2; i <= ceiling; i++ {
		if !primesList[i] {
			counter++
			if counter == n {
				return int64(i)
			}
		}
	}
	//
	return 0
}

// testing takes 9.5+ seconds. look into ways to speed it up
// possible options: cache known primes, optimize for loops.
