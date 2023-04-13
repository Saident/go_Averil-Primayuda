# Section 24 Clean Architecture

## Mempelajari beberapa hal seperti :

### Clean and Hexagonal Architecture
- Independent of Frameworks
Ini memungkinkan Anda untuk menggunakan frameworks seperti alat, daripada memiliki batasan.
- Testable
Business rules dapat diuji tanpa hal lain.
- Independent of Ul
Ul dapat berubah dengan mudah, tanpa mengubah sistem lainnya.
- Independent of Database
Business rules tidak terikat dengan database, Anda dapat mengganti database dengan mudah.
- Independent of any external

### Context Golang
Apa itu context ?
Package context atau kita sebut saja context, adalah suatu package yang membawa deadline, cancellation signal, atau value lain pada request/permintaan API

### MVC to Clean Code
The goal ketika kita menggunakan clean code adalah kode kita lebih modular, scallabel, dan maintainable.
- Modular dalam artian kita bisa dengan mudah mengganti dipendensi satu ke dependensi lain
- Scallabel dalam artian kita dapat dengan mudah menambahkan feature baru dan lain sebagainya
- Maintainable dalam artian kita dapat dengan mudah memperbaiki issue bilamana terdapat issue pada kode kita

