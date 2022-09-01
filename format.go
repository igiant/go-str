package str

import "strings"

func IntToFormatStr(n int, sep byte) string {
	if n == 0 {
		return "0"
	}
	minus := false
	if n < 0 {
		minus = true
		n = -n
	}
	t := make([]uint8, 0)
	for i := 1; n > 0; i++ {
		m := n % 10
		t = append(t, uint8(m))
		n = n / 10
		if sep != 0 && n != 0 && i%3 == 0 {
			t = append(t, sep)
		}
	}
	s := strings.Builder{}
	if minus {
		s.WriteByte('-')
	}
	for i := len(t) - 1; i >= 0; i-- {
		if t[i] <= 9 {
			s.WriteByte(t[i] + '0')
		} else {
			s.WriteByte(t[i])
		}
	}
	return s.String()
}
