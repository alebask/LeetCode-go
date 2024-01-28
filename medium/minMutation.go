package medium

func minMutation(startGene string, endGene string, bank []string) int {
	if startGene == endGene {
		return 0
	}
	if len(bank) == 0 {
		return -1
	}

	bank_ht := make(map[string]bool)
	for i := range bank {
		bank_ht[bank[i]] = true
	}
	if _, ok := bank_ht[endGene]; !ok {
		return -1
	}

	choices := "ACGT"

	start_ht := make(map[string]struct{})
	end_ht := make(map[string]struct{})
	processed_ht := make(map[string]struct{})

	start_ht[startGene] = struct{}{}
	end_ht[endGene] = struct{}{}

	mutationsCount := 0

	for len(start_ht) > 0 {
		new_ht := make(map[string]struct{})
		for word := range start_ht {
			for i := range word {
				for j := range choices {
					if choices[j] != word[i] {
						s := word[:i] + string(choices[j]) + word[i+1:]
						if _, ok := end_ht[s]; ok {
							return mutationsCount + 1
						}
						_, processed := processed_ht[s]
						_, inBank := bank_ht[s]
						if !processed && inBank {
							processed_ht[s] = struct{}{}
							new_ht[s] = struct{}{}
						}
					}
				}
			}
		}
		start_ht = new_ht
		if len(start_ht) > len(end_ht) {
			start_ht, end_ht = end_ht, start_ht
		}
		mutationsCount++
	}

	return -1
}
