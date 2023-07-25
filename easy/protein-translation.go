package easy

func ProteinTranslation(RNA string) []string {
	res := []string{}
	mapProteins := map[string]string{
		"AUG": "Methionine",
		"UUU": "Phenylalanine",
		"UUC": "Phenylalanine",
		"UUA": "Leucine",
		"UUG": "Leucine",
		"UCU": "Serine",
		"UCC": "Serine",
		"UCA": "Serine",
		"UCG": "Serine",
		"UAU": "Tyrosine",
		"UAC": "Tyrosine",
		"UGU": "Cysteine",
		"UGC": "Cysteine",
		"UGG": "Tryptophan",
		"UAA": "Stop",
		"UAG": "Stop",
		"UGA": "Stop",
	}

	for i := range RNA {
		if i%3 == 0 && i+3 < len(RNA) {
			splitWord := RNA[i : i+3]
			if mapProteins[splitWord] == "Stop" {
				break
			} else {
				res = append(res, mapProteins[splitWord])
			}
		}
	}

	return res
}
