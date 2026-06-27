---
sidebar_position: 1
---
# EP1. LLM Sama, Hasil Berbeda

![Insinyur Berbantuan AI vs Insinyur AI-native](/img/01-same-llm-different-results.png)

---

"ChatGPT, buatkan modul login untuk saya!"

Jika Anda seorang pengembang, Anda mungkin sudah melemparkan permintaan seperti ini berkali-kali sehari—menggunakan Claude yang sama, GPT yang sama. Namun hasilnya bisa sangat berbeda.

Ada orang yang menyalin-tempel kode hasil AI hingga begadang semalaman. Ada pula yang dengan tenang mengarahkan pekerjaan seolah memimpin tim sepuluh insinyur brilian. Alatnya identik—lalu mengapa celahnya ada?

Menemukan jawaban atas pertanyaan itu adalah awal cerita hari ini.

---

## 🚨 Hari Pengembang A — "Apa pun Kata AI"

Pengembang junior A yang sibuk kembali berpacu dengan tenggat waktu. Terburu-buru, ia meminta AI:

> *"Tambahkan modul login."*

AI dengan cepat menghasilkan kode yang tampak masuk akal. A menempelkannya ke proyek tanpa banyak berpikir. *"AI pasti benar, kan."* Untuk sementara berjalan—lega.

**Tapi masalah sebenarnya muncul seminggu kemudian.**

Tinjauan keamanan menandai kerentanan manajemen sesi—atau ketidakcocokan antarmuka dengan modul lain. A, yang menempelkan kode yang tidak pernah dipahaminya, akhirnya menulis ulang semuanya dari awal melalui malam tanpa tidur lagi.

Basis kode tumbuh, tetapi kepala A dipenuhi hal-hal yang **tidak** ia ketahui—ketidaktahuan yang terus membesar.

Apakah ini terdengar familiar? *"Ya, kadang saya juga begitu..."*

---

## 🤝 Pendekatan Pengembang B — "Berpikir Bersama AI"

Pengembang B, masa kerja sama, juga membutuhkan modul login. Tapi hubungannya dengan AI sama sekali berbeda.

> *"Proyek kami akan memiliki sekitar 100.000 pengguna dan keamanan sangat penting. Pertama, sarankan pendekatan autentikasi yang bisa kita gunakan dan bandingkan trade-off masing-masing."*

Ketika AI menawarkan autentikasi berbasis sesi, JWT, OAuth 2.0, dan lainnya, B berdiskusi bolak-balik dengannya—merancang struktur terbaik untuk masa depan proyek (skalabilitas aplikasi seluler, dll.).

Baru setelah satu jam diskusi intens dan arsitektur yang terkunci, B meminta kode. Hasilnya mungkin terlihat mirip dengan A di permukaan—tetapi B **sepenuhnya memahami mengapa dirancang begini, mengapa penanganan error mengikuti pola ini, dan setiap alasan desain di baliknya**.

Apa bedanya? Teknologi? Versi LLM? Bukan. Ini soal **sikap**.

---

## 🔍 Insinyur Berbantuan AI vs Insinyur AI-native

Celah antara A dan B bukan sekadar kesenjangan keterampilan—ini perbedaan dalam **cara Anda berhubungan dengan AI**.

| | Insinyur Berbantuan AI | Insinyur AI-native |
|---|---|---|
| **Peran AI** | **Alat** yang membantu saya (seperti kalkulator) | **Rekan tim** yang mengerjakan pekerjaan |
| **Peran Anda** | Mengerjakan sendiri dengan bantuan AI | Mengarahkan, **meninjau**, dan menyetujui |
| **Gaya tinjauan** | "AI pasti benar" → lanjut | "Apakah saya memahami ini?" → tinjau, lalu setujui |
| **Hasil** | Lebih banyak kode↑, lebih sedikit pemahaman↓, lebih banyak ketidaktahuan↑ | Lebih banyak kode↑, lebih banyak pemahaman↑, lebih banyak kapabilitas↑ |

A adalah **Insinyur Berbantuan AI**. AI hanyalah alat—pembantu seperti kalkulator. Gunakan saat dibutuhkan dan lanjut.

B adalah **Insinyur AI-native**. AI yang mengerjakan; Anda memimpin sebagai ketua tim. Anda mengarahkan, AI mengeksekusi, dan Anda meninjau, memahami, menganalisis, serta mengevaluasi sebelum menyetujui. AI bukan alat—ia **rekan tim** yang bekerja atas nama Anda.

Apakah Anda melihat AI sebagai alat atau rekan tim—itu garis antara Insinyur Berbantuan AI dan Insinyur AI-native.

---

## 💡 "Apakah Itu Benar-benar Mungkin?" — OpenAI Membuktikannya

Banyak insinyur menolak di sini:

*"Tunggu—Anda bilang saya mendelegasikan pekerjaan ke AI dan hanya meninjau? Apakah itu realistis?"*

Itu pertanyaan yang wajar. Dan ada jawaban yang jelas.

**Ya. Sudah terbukti dalam praktik.**

Pada April tahun ini, [insinyur OpenAI Ryan Lopopolo membagikan](https://finance.biggo.com/news/308f3eaef3245166) hal berikut:

- **Tiga insinyur** membangun **satu juta baris** kode produksi dalam **lima bulan**.
- Dan yang krusial—mereka **tidak menulis satu baris pun dengan tangan**.
- Mereka bergerak **10× lebih cepat** daripada pengembangan tradisional.

Tim menjalankan beberapa agen pengkodean AI otonom melalui sistem bernama Symphony. Apa peran manusia? **Arsitek & Penjaga**—merancang sistem dan meninjau apa yang dihasilkan agen.

Lopopolo menggambarkan peran mereka begini:

> *"Group tech lead untuk organisasi 500 orang"*

Tiga orang menghasilkan output organisasi 500 orang. Itulah Insinyur AI-native dalam praktik.

---

## 🎯 Berapa Banyak Rekan Tim AI yang Anda Miliki?

Kita sudah menetapkan bahwa pendekatan AI-native **memungkinkan**. Jadi pertanyaannya berubah.

Bukan lagi *"Apakah memungkinkan?"* Melainkan ini:

> **"Berapa banyak rekan tim AI yang Anda miliki?"**

Bekerja dengan cara AI-native berarti memperlakukan AI sebagai rekan tim.

- **Satu?** — Anda bisa mendelegasikan pekerjaan setara satu orang ke AI dan meninjaunya.
- **Empat?** — Anda bisa mendelegasikan pekerjaan setara empat orang.
- **Sepuluh? Dua puluh?** — Anda bisa menjalankan rekan tim AI seperti tim kecil yang Anda pimpin.

Berapa banyak rekan tim AI yang Anda pimpin sekarang? Atau lebih tepatnya—berapa banyak yang **bisa** Anda pimpin?

**Pada akhirnya, kapabilitas Anda yang penting.** Sama seperti keterampilan manajer menentukan ukuran tim, keterampilan Insinyur AI-native menentukan berapa banyak rekan tim AI yang bisa mereka jalankan.

Lalu bagaimana membangun kapabilitas itu?

---

## 🦉 Cocrates Harness — Hubungan Sokratik Timbal Balik

**Cocrates** adalah agent harness yang membantu Anda menjadi Insinyur AI-native.

Cocrates adalah **Co + Socrates**—**hubungan Sokratik timbal balik**.

**Anda adalah Socrates bagi AI.**

AI hanya menghasilkan jawaban paling masuk akal secara probabilistik. Anda bertanya, meninjau, dan menilai agar ia tidak menghasilkan yang salah. Itulah keterampilan inti Insinyur AI-native.

**AI juga Socrates bagi Anda.**

Ia menghentikan Anda membuat permintaan samar atau menyetujui output yang tidak Anda pahami. Ia tidak membiarkan Anda mengabaikan apa yang hanya Anda ketahui setengahnya. Seperti Socrates, ia menampakkan ketidaktahuan Anda dan membantu Anda mengembangkan kapabilitas sendiri.

Cocrates adalah **agent harness yang membuat Anda menggunakan AI dalam hubungan Sokratik timbal balik ini**—alat untuk memimpin rekan tim AI sebagai Insinyur AI-native, dan cara yang benar menggunakan AI.

---

## 📌 Poin Penting

1. **LLM sama, hasil berbeda.** Bukan performa AI—**pendekatan Anda** yang membentuk hasil.
2. **Insinyur AI-native.** Seseorang yang **mengarahkan** pekerjaan AI, **meninjau**nya, dan **menyetujuinya**. Tinjauan dan persetujuan memisahkan Insinyur Berbantuan AI dari Insinyur AI-native.
3. **"Berapa banyak rekan tim AI yang Anda miliki?"** Jawabannya tergantung kapabilitas Anda—satu, sepuluh, atau lebih. Yang penting adalah **kapabilitas Anda**.

Luangkan waktu sejenak untuk bertanya pada diri sendiri:

- *"Apakah saya Insinyur AI-native sekarang—atau masih Insinyur Berbantuan AI?"*
- *"Dari permintaan yang saya buat ke AI hari ini, berapa banyak yang saya tinjau dengan cara AI-native?"*

Langkah pertama menuju AI-native adalah **kesadaran diri**.

---

## 🎬 Selanjutnya

Hari ini kita mempelajari *apa* masalahnya: **sikap**. Celah antara Insinyur Berbantuan AI dan Insinyur AI-native bukan teknologi—melainkan sikap.

Lalu apa keterampilan inti yang harus dimiliki Insinyur AI-native?

**Tinjauan.** Memahami, menganalisis, mengevaluasi, dan menyetujui apa yang dihasilkan AI. Tindakan tinjauan itulah pekerjaan terpenting Insinyur AI-native.

**Episode berikutnya**, kita menyatakan prinsip yang mengalir di seluruh seri ini:

> **"The unexamined code is not worth generating."**

Kita akan mengurai mengapa satu baris ini adalah senjata bagi Insinyur AI-native—dan apa arti **pemeriksaan** yang sebenarnya.

---

*Seri ini memperkenalkan framework Cocrates Harness. Cocrates adalah agent harness yang dirancang untuk dialog Sokratik agar pengguna tetap memegang kendali dan terus berkembang.*
