package helper

import (
	"sort"
)

// GenerateA000124 menghasilkan deret bilangan berdasarkan pola A000124.
// Rumus A000124: a(n) = a(n-1) + n
// Parameter: n (int) - jumlah angka dalam deret
// Return: []int - daftar bilangan dalam deret
func GenerateA000124(n int) []int {
	sequence := make([]int, n)
	sequence[0] = 1
	for i := 1; i < n; i++ {
		sequence[i] = sequence[i-1] + i
	}
	return sequence
}

// DenseRanking menghitung peringkat berdasarkan metode Dense Ranking.
// Parameter:
//   - scores ([]int): daftar skor di papan peringkat (leaderboard)
//   - gitsScores ([]int): skor yang diperoleh GITS
//
// Return: []int - daftar peringkat yang diperoleh GITS
func DenseRanking(scores []int, gitsScores []int) []int {
	// Menghapus duplikasi skor dan membuat daftar peringkat unik
	rankedScores := []int{}
	seen := map[int]bool{}
	for _, score := range scores {
		if !seen[score] {
			seen[score] = true
			rankedScores = append(rankedScores, score)
		}
	}

	// Mengurutkan skor dalam urutan menurun
	sort.Sort(sort.Reverse(sort.IntSlice(rankedScores)))

	// Menentukan peringkat untuk setiap skor GITS
	result := make([]int, len(gitsScores))
	for i, gScore := range gitsScores {
		pos := sort.Search(len(rankedScores), func(j int) bool { return rankedScores[j] <= gScore })
		result[i] = pos + 1
	}

	return result
}

// IsBalancedBracket memeriksa apakah string berisi tanda kurung yang seimbang.
// Parameter: s (string) - string yang berisi tanda kurung
// Return: bool - true jika tanda kurung seimbang, false jika tidak
func IsBalancedBracket(s string) bool {
	stack := []rune{}
	bracketMap := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, char := range s {
		switch char {
		case '(', '[', '{':
			// Menambahkan tanda kurung pembuka ke dalam stack
			stack = append(stack, char)
		case ')', ']', '}':
			// Memeriksa apakah stack kosong atau pasangan tidak cocok
			if len(stack) == 0 || stack[len(stack)-1] != bracketMap[char] {
				return false
			}
			// Menghapus tanda kurung yang sudah dipasangkan dari stack
			stack = stack[:len(stack)-1]
		}
	}
	// Jika stack kosong, berarti tanda kurung seimbang
	return len(stack) == 0
}
