package render

func WordSplitter(line string) ([]string, error) {
	line += " "
	result := []string{}
	buf := []byte{}
	for i := 0; i < len(line)-1; i++ {
		if line[i] == byte(32) && len(buf) != 0 {
			result = append(result, string(buf))
			buf = []byte{}
		} else if (IsLetter(line[i]) && !IsLetter(line[i+1])) || (!IsLetter(line[i]) && IsLetter(line[i+1])) {
			buf = append(buf, line[i])
			result = append(result, string(buf))
			buf = []byte{}
		} else if !IsLetter(line[i]) && !IsLetter(line[i+1]) && line[i] != line[i+1] {
			buf = append(buf, line[i])
			result = append(result, string(buf))
			buf = []byte{}
		} else {
			buf = append(buf, line[i])
		}
	}
	return result, nil
}

func IsLetter(letter byte) bool {
	if letter >= byte(65) && letter <= byte(90) || letter >= byte(97) && letter <= byte(122) {
		return true
	} else {
		return false
	}
}
