package str

import (
	"testing"
)

func TestIntToStr(t *testing.T) {
	if IntToFormatStr(1234567890, ' ') != "1 234 567 890" {
		t.Errorf("IntToFormatStr(1234567890, ' ') != 1 234 567 890 (%s != 1 234 567 890)", IntToFormatStr(1234567890, ' '))
	}
	if IntToFormatStr(-1234567890, ' ') != "-1 234 567 890" {
		t.Errorf("IntToFormatStr(-1234567890, ' ') != -1 234 567 890 (%s != -1 234 567 890)", IntToFormatStr(-1234567890, ' '))
	}
	if IntToFormatStr(1234567890, 0) != "1234567890" {
		t.Errorf("IntToFormatStr(1234567890, 0) != 1234567890 (%s != 1234567890)", IntToFormatStr(1234567890, 0))
	}
	if IntToFormatStr(123456789, '_') != "123_456_789" {
		t.Errorf("IntToFormatStr(1234567890, '_') != 123_456_789 (%s != 123_456_789)", IntToFormatStr(1234567890, ' '))
	}
	if IntToFormatStr(-123456, '_') != "-123_456" {
		t.Errorf("IntToFormatStr(-123456, '_') != -123_456 (%s != -123_456)", IntToFormatStr(-123456, '_'))
	}
	if IntToFormatStr(0, ' ') != "0" {
		t.Errorf("IntToFormatStr(0, ' ') != 0 (%s != 0)", IntToFormatStr(0, ' '))
	}
	if IntToFormatStr(-1, ' ') != "-1" {
		t.Errorf("IntToFormatStr(-1, ' ') != -1 (%s != -1)", IntToFormatStr(-1, ' '))
	}
}
