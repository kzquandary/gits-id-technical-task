package main

import (
	"fmt"
	"gist/helper"
)

func main() {
	for {
		fmt.Println("Pilih Soal:")
		fmt.Println("1. A000124 Sequence")
		fmt.Println("2. Dense Ranking")
		fmt.Println("3. Balanced Bracket")
		fmt.Println("4. Exit")
		fmt.Print("Masukkan pilihan: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Masukkan jumlah elemen: ")
			var n int
			fmt.Scan(&n)
			fmt.Println("Output:", helper.GenerateA000124(n))
		case 2:
			fmt.Print("Masukkan jumlah skor di leaderboard: ")
			var m int
			fmt.Scan(&m)
			fmt.Println("Masukkan skor-skor leaderboard:")
			scores := make([]int, m)
			for i := 0; i < m; i++ {
				fmt.Scan(&scores[i])
			}
			fmt.Print("Masukkan jumlah skor GITS: ")
			var p int
			fmt.Scan(&p)
			fmt.Println("Masukkan skor GITS:")
			gitsScores := make([]int, p)
			for i := 0; i < p; i++ {
				fmt.Scan(&gitsScores[i])
			}
			fmt.Println("Output:", helper.DenseRanking(scores, gitsScores))
		case 3:
			fmt.Print("Masukkan string brackets: ")
			var s string
			fmt.Scan(&s)
			if helper.IsBalancedBracket(s) {
				fmt.Println("Output: YES")
			} else {
				fmt.Println("Output: NO")
			}
		case 4:
			fmt.Println("Keluar...")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}
