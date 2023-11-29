package codewars

func DNAStrand(dna string) string {
	result := ""
	for _, v := range dna {
		switch {
		case v == 'A':
			result += "T"
		case v == 'T':
			result += "A"
		case v == 'G':
			result += "C"
		case v == 'C':
			result += "G"
		}
	}
	return result
}
