---
sidebar_position: 3
---
# EP3. Menginstal Cocrates Harness

![Prosedur Instalasi Cocrates](/img/03-installation.png)

---

Di Ep2 kita mendeklarasikan prinsip yang kuat:

> "The unexamined code is not worth generating."
> Output yang belum ditinjau tidak layak dibuat.

Prinsipnya jelas. Semua yang dibuat AI harus lulus tinjauan (U→A→E→A), dan Anda tidak boleh membiarkan ketidaktahuan tanpa perhatian.

Tapi satu pertanyaan praktis tetap ada:

*"Saya paham prinsipnya. Bagaimana saya benar-benar memulai?"*

Hari ini kita menjawabnya. Instalasi memakan waktu sekitar lima menit. Tapi ingat—awal sebenarnya adalah **setelah** instalasi.

---

## 🎭 Apa Sebenarnya Cocrates Harness

Pertama, mari kita jelas tentang apa itu Cocrates Harness.

Cocrates Harness bukan program mandiri. Ini **plugin** yang berjalan di platform bernama **opencode**.

Bayangkan begini: opencode adalah **panggung**; Cocrates Harness adalah **aktor** di panggung itu. Tanpa panggung, tidak ada tempat bagi aktor untuk berdiri.

Jadi penyiapan Cocrates punya tiga langkah:

**Satu,** instal opencode.
**Dua,** konfigurasikan plugin Cocrates.
**Tiga,** siapkan file skill.

Ingat ketiganya. Penyiapan lebih sederhana dari yang Anda kira.

---

## 🖥️ Tiga Opsi UI

Ada tiga cara menggunakan opencode. Pilih yang sesuai lingkungan Anda.

### 1) opencode Terminal (TUI)
Penggunaan berpusat pada keyboard di terminal. Bagus jika Anda nyaman dengan CLI.

### 2) opencode Desktop (Beta)
Aplikasi GUI yang diinstal. Masih beta—fitur dan stabilitas mungkin belum matang.

### 3) VS Code atau Cursor + Ekstensi OpenChamber ⭐ (Direkomendasikan)
Jika Anda mengenal VS Code atau Cursor, Anda bisa menggunakan Cocrates senatural Copilot atau Cursor. Catatan: VS Code + OpenChamber tetap memerlukan opencode terinstal.

---

## ⚙️ Langkah 1 — Instal opencode

### macOS / Linux

Buka terminal dan jalankan satu baris:

```bash
curl -fsSL https://opencode.ai/install | bash
```

curl mengunduh skrip instalasi dengan aman; bash menjalankannya. Biner opencode terinstal otomatis.

Setelah selesai, verifikasi:

```bash
opencode --version
```

Jika versi tercetak, Anda siap. Jika muncul *"command not found"*, tambahkan `~/.local/bin` ke PATH Anda.

### Windows

Windows punya curl bawaan, tetapi tidak bash secara default. Gunakan **Chocolatey**, manajer paket Windows.

Buka PowerShell **sebagai Administrator** dan instal Chocolatey:

```powershell
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))
```

Dengan Chocolatey terinstal, opencode cukup satu baris:

```powershell
choco install opencode
```

Lalu periksa dengan `opencode --version`.

### VS Code + Ekstensi OpenChamber

CLI opencode selesai. Jika UI terminal terasa canggung, instal VS Code + ekstensi **OpenChamber**.

1. Jika Anda belum punya VS Code, unduh dari [code.visualstudio.com](https://code.visualstudio.com) dan instal.
2. Buka VS Code, klik Extensions (Ctrl+Shift+X), cari **"OpenChamber"**, dan instal.

Dengan OpenChamber, Anda bisa menggunakan Cocrates di dalam VS Code seperti Copilot atau Cursor.

![Ekstensi OpenChamber](/img/openchamber-extension.png)

---

## 🔧 Langkah 2-3 — Konfigurasi Plugin + File Skill

Setelah opencode terinstal, cara termudah adalah meminta AI menginstal Cocrates Harness.

Jalankan Cocrates dan minta:

> *"Install Cocrates Harness based on https://cocrates.ai/install.md."*

AI akan menanganinya—mengunduh skill dari repo GitHub, menambahkan konfigurasi plugin ke opencode.jsonc, dan menyelesaikan penyiapan.

Jika sudah terinstal, Anda bisa mengulangi permintaan yang sama. AI menanganinya sebagai **upgrade** mengikuti [install.md](https://cocrates.ai/install.md).

- **Plugin:** Membandingkan versi terinstal dengan versi terbaru. Memperbarui jika ada versi baru dan melaporkan nomor versi lama→baru. Jika sudah terbaru, memberitahu bahwa tidak perlu pembaruan.
- **Skill:** Membandingkan setiap skill antara repo resmi dan file lokal. Melewati skill yang identik; jika ada perbedaan, menampilkan ringkasan perubahan dan meminta pilihan **mengganti dengan versi resmi atau mempertahankan salinan lokal**.

Tidak perlu menghafal perintah "upgrade" terpisah—cukup minta instalasi seperti biasa.

![Cocrates Harness Terinstal](/img/cocrates-harness-installed.png)

**Tapi. Itu bukan akhirnya.**

Jika AI menginstalnya dan Anda lanjut tanpa melihat—itu masih Insinyur Berbantuan AI. **Meninjau dan menyetujui apa yang diinstal AI**—itulah Insinyur AI-native.

### Verifikasi Instalasi

Cocrates Harness punya dua bagian: **Plugin** dan **Skill.** Periksa keduanya untuk memastikan penyiapan.

**Cek Plugin.** Buka `~/.config/opencode/opencode.jsonc`. Apakah `"plugin": ["@cocrates/cocrates-harness"]` ada di sana? Jika ya, lulus.

**Cek Skill.** Buka `~/.config/opencode/skills/`. Apakah ada file skill seperti adr-writing, spec-writing, education, dll.? Jika ya, lulus.

Dua cek itu cukup. Jika AI melakukannya dengan benar, Anda verifikasi dan setujui.

Jika keduanya terlihat baik, restart opencode dan pilih Cocrates Agent. Anda aktif.

---

## 🦉 Percakapan Pertama — Bagian Sebenarnya Dimulai Sekarang

Instalasi selesai! Anda siap. Waktunya merasakan hal sebenarnya.

Tanyakan pertanyaan pertama Anda kepada Cocrates:

> **"Tell me about Bloom's Taxonomy."**

Satu peringatan: Cocrates tidak akan berperilaku seperti AI biasa.

Kebanyakan AI memuntahkan penjelasan panjang. Cocrates tidak. Ia akan menyelipkan contoh yang sedikit cacat dan memberi Anda misi pertama.

**Instalasi bukan akhir. Bagian sebenarnya dimulai sekarang.**

---

## 📌 Poin Penting

1. **Cocrates adalah plugin opencode.** Ia berjalan di opencode dan terdiri dari Plugin + Skill.
2. **Penyiapan yang direkomendasikan: opencode + VS Code + OpenChamber.** Penggunaan GUI Cocrates di VS Code.
3. **Minta AI menginstal, lalu verifikasi.** Satu baris: "Install based on install.md." Permintaan yang sama juga meng-upgrade instalasi yang sudah ada. Setelah instalasi atau upgrade, pastikan Plugin dan Skill ada. Bangun kebiasaan memeriksa dan meninjau pekerjaan AI.

---

**Tanyakan pada diri sendiri:**

- Apakah opencode terinstal?
- Apakah Cocrates Agent aktif?
- Apakah Anda siap mengajukan pertanyaan pertama?

Jika ketiganya ya, Anda sudah mengambil langkah pertama sebagai Insinyur AI-native.

---

## 🎬 Selanjutnya

Hari ini kita menyiapkan alat. Kita menginstal Cocrates Harness, memeriksa Plugin dan Skill, dan bersiap untuk pertanyaan pertama.

Tapi mengetahui prinsip dan mempraktikkannya berbeda. Memiliki alat dan menggunakannya juga berbeda.

Selanjutnya, Ep4: **percakapan nyata pertama** Anda dengan Cocrates.

Ketika Anda bertanya "Tell me about Bloom's Taxonomy," bagaimana Cocrates merespons? Kita akan menelusuri pipeline pembelajaran tiga langkah—Education → Knowledge Capture → Reflection—dengan dialog nyata.

Pertemuan pertama Anda dengan AI yang memprovokasi pemikiran. Layak dinantikan.

---

*Seri ini memperkenalkan framework Cocrates Harness. Cocrates adalah agent harness yang dirancang untuk dialog Sokratik agar pengguna tetap memegang kendali dan terus berkembang.*
