# Section 25 Docker

## Mempelajari beberapa hal seperti :

### What is a container ?
Container Bukan virtual machine.
Jawaban Singkat: Container adalah proses dengan isolasi sistem file.
Jawaban Panjang: Semua yang ada di Linux adalah file.
- /dev/sda = hardisk
- /dev/proc = proses
- /dev/usb
- /dev/cpu
- /dev/std (in/out)
- /bin/bash... (binary file)

### Docker Structure
Docker Structure terdiri dari tiga bagian, diantaranya adalah :
- Client
Terdapat command seperti docker run, pull, dan build.
- Docker Host
Terdiri dari dua bagian yaitu containers dan images.
- Registry
Contohnya adalah nginx

### What you can do with Docker ?
Build Dockerfile kedalam Docker Images, lalu mengupload dan publish images tersebut kedalam Docker Hub / Image Repository. Langkah terakhir adalah menjalankan image tersebut pada Docker Host.

