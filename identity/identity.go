// Package identity validates Turkish national identity numbers (TCKN) and
// tax identification numbers (VKN).
package identity

// ValidTCKN reports whether s is a valid Turkish national identity number.
// Rules: 11 digits, the first digit cannot be 0, and the 10th and 11th
// digits are check digits derived from the digit sums.
func ValidTCKN(s string) bool {
	if len(s) != 11 {
		return false
	}

	var d [11]int
	for i := 0; i < 11; i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return false
		}
		d[i] = int(c - '0')
	}
	if d[0] == 0 {
		return false
	}

	odd := d[0] + d[2] + d[4] + d[6] + d[8]
	even := d[1] + d[3] + d[5] + d[7]
	if (odd*7+even*9)%10 != d[9] {
		return false
	}

	var first10 int
	for i := 0; i < 10; i++ {
		first10 += d[i]
	}
	return first10%10 == d[10]
}

// ValidVKN reports whether s is a valid Turkish tax identification number.
// It is 10 digits long and uses the official Ministry of Finance algorithm.
func ValidVKN(s string) bool {
	if len(s) != 10 {
		return false
	}

	var d [10]int
	for i := 0; i < 10; i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return false
		}
		d[i] = int(c - '0')
	}

	var sum int
	for i := 0; i < 9; i++ {
		tmp := (d[i] + (9 - i)) % 10
		if tmp == 0 {
			sum += 9
			continue
		}
		p := (tmp * pow2(9-i)) % 9
		if p == 0 {
			p = 9
		}
		sum += p
	}

	check := (10 - sum%10) % 10
	return check == d[9]
}

func pow2(n int) int {
	r := 1
	for i := 0; i < n; i++ {
		r *= 2
	}
	return r
}
