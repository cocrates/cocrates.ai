---
sidebar_position: 5
---
# EP5. Generasi Berbasis Arsitektur dalam Praktik

![Generasi Berbasis Arsitektur](/img/05-architecture-driven-generation-practice.png)

---

Di Ep4 kita mengalami aktivitas inti pertama Cocrates—**pipeline Learning**. Education → Knowledge Capture → Reflection. Cara belajar dengan AI.

Hari ini kita mempraktikkan aktivitas inti kedua: **pipeline generasi artefak berbasis arsitektur**.

Cocrates melakukan dua hal besar: **belajar** dan **membuat**. Ep4 adalah belajar; hari ini adalah **membuat**.

ADR → Spec → Generation → Verification. Kita akan menelusuri pembangunan perangkat lunak dengan AI menggunakan kasus nyata.

---

## 💬 "Bangun jsondb" — Tapi AI Memaksa Desain Dulu

Bayangkan Anda meminta Cocrates:

> *"Develop jsondb in examples/jsondb."*

AI biasa mungkin berkata *"jsondb? Siap!"* dan memuntahkan ratusan baris kode—tanpa tinjauan, tanpa desain, kode dulu.

Cocrates menjawab berbeda:

> *"Untuk apa Anda akan menggunakan jsondb ini? Saya akan merancang arsitekturnya dulu, mendapat tinjauan dan persetujuan Anda, lalu menghasilkan kode dengan aman."*

Itulah nilai inti Cocrates. Ingat Ep2? **"Output yang tidak diperiksa tidak layak dihasilkan."** Cocrates adalah AI yang menghormati prinsip itu. Tanpa desain, tanpa tinjauan—tidak ada kode.

Hari ini kita membangun jsondb—penyimpanan JSON lokal sederhana. Bahkan proyek kecil ini membutuhkan debat arsitektur nyata.

---

## 🏛️ ADR ① — Debat Model Penyimpanan: Perangkap "Generik"

**Situasi:** Bagaimana kita menyimpan data?

**Saran AI:** *"Struktur Collection-Document seperti MongoDB umum dan mudah."*

AI memimpin dengan *"generik."* Pola yang paling sering ia lihat dalam pelatihan.

Tapi *"generik"* berarti *"biasanya dilakukan begini"*—bukan **"tepat untuk situasi saya."**

Pengguna menolak:

> *"Tidak. Saya ingin penanganan berbasis path—`Set("episode/e1.json")` harus membuat `episode/e1.json` di disk."*

Itu penyimpanan **Path-Addressable**—berbeda dari Collection-Document.

Mereka memutuskan pemetaan path.

**💡 Pelajaran:** Generik tidak selalu tepat untuk kebutuhan Anda. AI menawarkan pola paling umum. Hanya **Anda** yang tahu konteks sebenarnya.

---

## 🏛️ ADR ② — Debat Arsitektur: Perangkap "Mudah dan Sederhana"

**Situasi:** Bagaimana kita mengirim jsondb?

**Saran AI:** *"Library engine + app paling mudah—tinggal import dan pakai."*

Kedengarannya praktis. Import dan jalan.

Pengguna bertanya:

> *"Bagaimana jika aplikasi lain perlu mengakses jsondb ini nanti?"*

Itu membalik segalanya. Library + app hanya bekerja in-process. Banyak aplikasi membutuhkan **Client-Server**.

Mereka pindah dari Library + app ke Client-Server. Jalur *"mudah"* tidak menahan kebutuhan masa depan.

**💡 Pelajaran:** Mudah dan sederhana tidak selalu skalabel. AI menyukai implementasi tercepat. Anda bertanya tentang masa depan untuk mencapai struktur yang lebih baik.

---

## 🏛️ ADR ③ — Debat Konkurensi: Perangkap "Kualitas"

**Situasi:** Permintaan konkuren—bagaimana kita menanganinya?

**Saran AI:** *"Satu RWMutex tingkat DB cukup. Sederhana."*

Secara teknis valid.

Pengguna mengantisipasi performa:

> *"Saat menulis satu file, apakah semua bacaan file lain harus menunggu?"*

Tepat. Satu RWMutex tingkat DB memblokir setiap bacaan saat ada file yang menulis. Dengan 100 atau 1000 file, itu bottleneck.

Mereka beralih ke **Per-File RWMutex Map**—akses konkuren ke file berbeda.

**💡 Pelajaran:** AI tidak salah secara teknis—tetapi **kualitas (performa, skala)** tidak sepenuhnya dipertimbangkan. Anda mempertanyakan performa untuk meningkatkan kualitas.

---

## 🎯 Pelajaran dari ADR — Tiga Alasan Saran AI Meleset

Tiga ADR, satu pelajaran penting:

**Saran AI bisa meleset karena banyak alasan—bukan hanya halusinasi.**

1. **Generik:** AI menawarkan pola paling umum. Generik ≠ tepat untuk Anda. (ADR ①: penyimpanan)
2. **Sederhana:** AI menyukai implementasi termudah. Mungkin tidak memenuhi kebutuhan masa depan. (ADR ②: arsitektur)
3. **Buta kualitas:** Secara teknis baik tetapi lemah pada performa atau skala. (ADR ③: konkurensi)

Mengenali ini membantu Anda meninjau proposal pertama AI lebih hati-hati.

---

## 📋 Spec — Mengintegrasikan Keputusan

Setelah tiga ADR disetujui, Cocrates menggabungkan setiap keputusan ke **satu dokumen mandiri**—**Spec**.

Spec saja harus sepenuhnya menjelaskan apa yang dibangun. Tidak perlu membuka ADR.

- Penyimpanan: Path-Addressable (key → path file langsung)
- Arsitektur: Client-Server (multi-proses)
- Konkurensi: Per-File RWMutex Map

---

## ⚙️ Generation — Kode dari Spec

Cocrates menghasilkan kode dari Spec:

- 7 file library Go
- Server REST API
- CLI
- Semua tes lulus

Karena generasi mengikuti Spec, kode sesuai keputusan—tanpa asumsi tambahan.

---

## 🔍 Verification — Cek Spec dan Temuan Tak Terduga

Bahkan generasi berbasis Spec bisa menyimpang dari Spec. Verifikasi menggunakan 72 item checklist:

- **71: lulus**
- **1: gagal** — double slash (`//`) di URL skema CLI

Temuan penting: **6 ASR Tidak Terdokumentasi.**

ASR tidak terdokumentasi adalah keputusan struktural dalam output yang tidak dinyatakan di Spec—6 dalam kasus ini:

- Pendekatan URL encoding
- Validasi SetPath yang hilang
- Empat pilihan desain implisit lainnya

**ASR tidak terdokumentasi bukan kegagalan.** Ini kesempatan untuk mempertajam Spec. Tinjau keenamnya dalam desain dan perbarui.

---

## 🔄 Living Cycle — Struktur Tidak Beku

Banyak orang mengira ini **Waterfall**.

Waterfall: desain → implementasi → verifikasi (selesai)

Kenyataannya adalah **siklus**:

```
ADR → Spec → Generation → Verification
  ↑                            │
  └──────── feedback ←─────────┘
```

Masalah dalam verifikasi mengembalikan Anda ke desain—tinjau ulang ADR, ubah jika perlu. Persyaratan berubah? Siklus yang sama: ADR → Spec → regenerasi → verifikasi ulang.

**ADR adalah dokumen hidup.** Keputusan hari ini mungkin tidak bertahan besok.

**Spec adalah dokumen hidup.** Ketika ADR berubah, Spec harus berubah.

**Struktur tidak ditetapkan sekali lalu dilupakan.** Struktur mengarahkan pengembangan; verifikasi dan kebutuhan yang berubah memperbaiki struktur.

> Pengembangan berbasis arsitektur bukan Waterfall. **Struktur memimpin pengembangan; pengalaman pengembangan memperbaiki struktur.**

---

## 📌 Poin Penting

1. **Tinjau saran AI.** Bukan hanya halusinasi—kegagalan Generik, Sederhana, atau Buta kualitas terjadi. Tinjauan dan persetujuan pengguna sangat penting.
2. **ADR dan Spec adalah dokumen hidup.** Verifikasi atau perubahan persyaratan mengembalikan Anda ke desain. Struktur berevolusi.
3. **Verifikasi adalah kualitas.** Verifikasi Spec tidak bisa menangkap semuanya—tetapi tanpa verifikasi Anda tidak menangkap apa pun. ASR tidak terdokumentasi adalah peluang perbaikan Spec.
4. **Pengembangan berbasis arsitektur adalah Living Cycle.** Struktur memimpin; pengalaman menyempurnakan. Bukan Waterfall.

---

**Tanyakan pada diri sendiri:**

- Bisakah Anda menyebutkan empat tahap pipeline Artifact Generation?
- Bisakah Anda menjelaskan tiga alasan saran AI meleset—dengan contoh?
- Apa itu ASR Tidak Terdokumentasi?
- Mengapa pengembangan berbasis arsitektur bukan Waterfall?

---

## 🎬 Selanjutnya

Hari ini kita mengalami **generasi berbasis arsitektur**—ADR, Spec, Generation, Verification—dalam proyek nyata.

Selanjutnya, Ep6 melangkah lebih jauh. Ep5 menggunakan skill yang ada; Ep6 kita **membangun skill baru dengan Cocrates**—skill untuk membuat skill.

Alih-alih *"tulis laporan,"* ajari AI *"cara menulis laporan."* Rancang dan bangun skill yang otomatis menghasilkan laporan dan dokumentasi.

Satu pertanyaan penutup:

> **"Apakah saya pernah menyalin-tempel proposal pertama AI tanpa berpikir kedua kalinya?"**

---

*Seri ini memperkenalkan framework Cocrates Harness. Cocrates adalah agent harness yang dirancang untuk dialog Sokratik agar pengguna tetap memegang kendali dan terus berkembang.*
