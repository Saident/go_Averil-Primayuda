# Section 12 Concurrent Programming

## Mempelajari beberapa hal seperti :

### Different Sequential, Parallel, and Concurrent <br>
- Sequential merupakan program yang logikanya berjalan secara berurutan
- Dalam program Parallel, beberapa task dapat dijalankan bersamaan
- Dalam program Concurrent, beberapa task dapat dijalankan independen dan muncul bersamaan

### Goroutines <br>
Goroutine adalah fungsi atau metode yang dijalankan secara bersamaan (independen) dengan fungsi atau metode lain.

### Channels and Select <br>
- Channel adalah objek komunikasi yang digunakan goroutine untuk berkomunikasi satu sama lain.
- Select membuat program lebih mudah untuk mengontrol komunikasi data melalui satu atau banyak saluran.

### Race Condition <br>
Race condition adalah dimana 2 thread mengakses memori secara bersamaan, salah satunya sedang melakukan write. Kondisi race terjadi karena akses yang tidak sinkron ke memori.

