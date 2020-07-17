package revstr

func reverseWords(s string) string {
	// skipping all leading whitespace
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	// skipp all trailing whitespace
	j := len(s) - 1
	for j >= 0 && s[j] == ' ' {
		j--
	}
	if i > j {
		return ""
	}

	ss := make([]byte, j-i+1)
	r := i
	w := len(ss) - 1
	for i <= j {
		if s[i] == ' ' {
			w = w - (i - r - 1)
			t := w
			// copy a word
			for r < i {
				ss[w] = s[r]
				r++
				w++
			}
			// prefix the copied word with space
			w = t - 1
			ss[w] = ' '
			w--
			// skip space between words
			i++
			for s[i] == ' ' {
				i++
			}
			r = i
			continue
		}
		i++
	}
	// copy the last word
	w = w - (i - r - 1)
	t := w
	for r < i {
		ss[w] = s[r]
		r++
		w++
	}
	w = t
	return string(ss[w:])
}
