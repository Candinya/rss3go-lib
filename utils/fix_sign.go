package utils

func FixSign(oldSign string) string {
	var fixed string
	fixed = oldSign
	if fixed[:2] == "0x" {
		fixed = fixed[2:]
	}
	if fixed[128:] == "1b" {
		fixed = fixed[:128] + "00"
	} else if fixed[128:] == "1c" {
		fixed = fixed[:128] + "01"
	}
	return fixed
}
