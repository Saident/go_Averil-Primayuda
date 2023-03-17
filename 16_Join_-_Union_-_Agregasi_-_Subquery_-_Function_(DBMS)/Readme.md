# Section 16 Join – Union – Agregasi – Subquery – Function (DBMS)

## Mempelajari beberapa hal seperti :

### Join
Sebuah klausa untuk mengkombinasikan record dari dua atau lebih tabel. Jenis join ada 3, yaitu:
- Inner Join
Inner join akan mengembalikan baris-baris dari dua tabel atau lebih yang memenuhi syarat.
- Left Join
Left join akan mengembalikan seluruh baris dari tabel disebelah kiri yang dikenai kondisi ON dan hanya baris dari tabel disebelah kanan yang memenuhi kondisi join.
- Right Join
Right join akan mengembalikan semua baris dari tabel sebelah kanan yang dikenai kondisi ON dengan data dari tabel sebelah kiri yang memenuhi kondisi join. Teknik ini merupakan kebalikan dari left join.

### Union
Fungsi UNION pada SQL digunakan untuk menggabungkan dua tabel dalam bentuk baris baru ke bawah dimana field yang di-SELECT antara tabel satu dan lainnya adalah harus sama. Atau sederhananya yaitu untuk menempatkan baris dari kueri satu sama lain dan nilainya distinct/unik.

### Agregasi
Agregasi adalah di mana nilai beberapa baris dikelompokkan bersama untuk membentuk nilai ringkasan tunggal. Fungsi - fungsi pada agregasi sebagai berikut :
- Min
- Max
- Sum
- Avg
- Count
- Having

### Subquery
Subquery atau Inner query atau Nested query adalah query di dalam query SQL lain