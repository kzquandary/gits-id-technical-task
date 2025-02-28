## 1. A000124 Sequence

### **Deskripsi Soal**
Program ini menghasilkan deret angka berdasarkan rumus A000124 dari Sloane’s OEIS.

#### **Contoh Input:**
```
7
```

#### **Contoh Output:**
```
1-2-4-7-11-16-22
```

### **Jawaban**
Fungsi untuk menyelesaikan soal ini sudah ada di `function.go` dalam fungsi `GenerateA000124(n int)`. Fungsi ini bekerja dengan menambahkan bilangan sebelumnya dengan indeks saat ini + 1.

---

## 2. Dense Ranking

### **Deskripsi Soal**
GITS bermain game arcade dengan sistem peringkat *Dense Ranking*.

**Aturan:**
- Pemain dengan skor tertinggi mendapatkan peringkat 1.
- Pemain dengan skor yang sama memiliki peringkat yang sama.
- GITS ingin mengetahui peringkatnya setelah bermain.

#### **Contoh Input:**
```
7
100 100 50 40 40 20 10
4
5 25 50 120
```

#### **Contoh Output:**
```
6 4 2 1
```

### **Jawaban**
Fungsi untuk menyelesaikan soal ini sudah ada di `function.go` dalam fungsi `DenseRanking(scores []int, gitsScores []int)`. Fungsi ini menghapus skor duplikat, mengurutkan skor secara descending, dan mencari peringkat menggunakan `sort.Search`.

---

## 3. Balanced Bracket

### **Deskripsi Soal**
Program ini mengecek apakah tanda kurung dalam string seimbang atau tidak.

#### **Contoh Input & Output:**
| Input          | Output |
|---------------|--------|
| `{[()]}`      | YES    |
| `{[(])}`      | NO     |
| `{(([])[[]])}` | YES    |

**Aturan:**
1. Hanya tanda kurung `()`, `{}`, `[]` yang diperbolehkan.
2. Tanda kurung bisa dipisahkan dengan atau tanpa whitespace.
3. Program harus mengembalikan `YES` jika tanda kurung seimbang dan `NO` jika tidak.

### **Jawaban**
Fungsi untuk menyelesaikan soal ini sudah ada di `function.go` dalam fungsi `IsBalancedBracket(s string)`. Fungsi ini menggunakan *stack* untuk memastikan bahwa setiap tanda kurung pembuka memiliki pasangannya yang sesuai. Balanced Bracket Menggunakan satu loop untuk memproses string dan operasi *stack* **O(1) per operasi**, totalnya **O(n)**.


