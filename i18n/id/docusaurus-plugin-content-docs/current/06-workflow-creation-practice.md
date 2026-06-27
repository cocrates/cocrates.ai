---
sidebar_position: 6
---
# EP6. Pembuatan Skill Berbasis Arsitektur dalam Praktik

![Mengajari AI Membuat Skill Penulisan Laporan](/img/06-workflow-creation-practice.png)

---

Di Ep5 kita melihat apa yang terjadi ketika Anda meminta Cocrates "membangun jsondb"—ADR dan Spec sebelum generasi berbasis arsitektur. Kita belajar tidak mempercayai AI secara membabi buta dan merancang dengan tinjauan.

Hari ini kita melangkah lebih jauh. Ep5 adalah *"cara meminta AI membuat kode."* Ep6 adalah **"cara mengajari AI cara bekerja."**

Subjudulnya mengatakannya: ajari AI *"cara menulis laporan."* Melampaui permintaan sekali pakai—merancang perilaku AI itu sendiri.

---

## 🎣 "Tulis Laporan" vs "Cara Menulis Laporan"

Bandingkan dua pendekatan.

**Biasa: "Tulis laporan"**

Setiap permintaan menghasilkan kualitas berbeda. Bagus hari ini, buruk besok—struktur bervariasi bahkan pada topik yang sama. Kelelahan repetisi. Itulah batas **permintaan sekali pakai**.

**Cocrates: "Buat skill untuk menulis laporan"**

Bangun sekali, pakai selamanya. Jenis permintaan sama → alur kerja sama → kualitas konsisten.

Seperti meminta *"tangkapkan ikan"* setiap kali vs mengajari *"cara memancing."*

**Skill adalah standar perilaku AI—alur kerja yang dapat diulang.** Aturan agar AI menangani jenis permintaan yang sama secara konsisten.

Hari ini kita membangun **`document-authoring`**—alur kerja untuk laporan dan dokumen penjelasan.

---

## 🎯 Kernel — Satu Kalimat untuk Tujuan Inti

Setiap skill dimulai dengan **Kernel**—satu kalimat yang mendefinisikan apa yang dilakukan skill.

Cocrates mengajukan tiga pertanyaan:

1. *"Apa yang dihasilkan skill ini?"*
2. *"Siapa yang menggunakannya?"*
3. *"Bentuk apa outputnya?"*

Dari jawaban muncul Kernel:

> *"Skill ini membantu menghasilkan penjelasan atau laporan Markdown melalui tahap yang dapat ditinjau dalam situasi tipikal."*

Satu kalimat mendefinisikan skill. Kernel memandu semua yang mengikutinya.

---

## 🏗️ Frame — Membangun Kerangka

Dengan Kernel terdefinisi, **Frame** merancang hierarki dokumen dan tahap generasi.

Hierarki dokumen:
```
Meta → Outline → Body → Conclusion → Appendix
```

Tahap generasi (P1–P5):
```
P1: Analisis persyaratan → P2: Outline → P3: Konten bagian → P4: Integrasi & tinjauan → P5: Pemformatan akhir
```

Setiap tahap punya input, output, dan kondisi persetujuan yang jelas.

---

## 🎯 Outline — Momen Dramatis

**Hal paling penting terjadi di Outline.**

Cocrates mengusulkan *"Pendahuluan → Isi → Kesimpulan"*—tipikal dan aman.

Pengguna menolak:

> *"Tidak selalu pendahuluan → isi → kesimpulan. Beberapa laporan tidak butuh pendahuluan; memimpin dengan kesimpulan bisa lebih baik."*

Itu mematahkan **Silent Default** AI. Cocrates merancang **enam struktur dokumen yang dapat disesuaikan**:

1. **Standar:** Pendahuluan → Isi → Kesimpulan
2. **Kesimpulan dulu:** Kesimpulan → Isi → Ringkasan
3. **Masalah-solusi:** Masalah → Penyebab → Solusi
4. **Kronologis:** Latar belakang → Perkembangan → Hasil
5. **Komparatif:** Topik A → Topik B → Sintesis
6. **Bebas:** Pengguna mendefinisikan struktur

**💡 Pelajaran:** Proposal pertama AI tidak selalu benar. AI menawarkan pola paling umum. Konteks dan wawasan Anda membangun struktur yang lebih baik. *"Tidak selalu intro→isi→kesimpulan"* mematahkan ide tetap dan mengevolusi arsitektur yang lebih fleksibel.

---

## 📋 Spec — Menyelesaikan Desain

Setelah Outline ditetapkan, Cocrates menggabungkan keputusan ke **Spec**.

**Lima prinsip inti:**
1. Gerbang persetujuan tahap—tinjauan pengguna di setiap langkah
2. Perluasan bertahap Snowflake—mulai kecil, perbaiki secara inkremental
3. Mandiri—Spec saja harus cukup untuk menghasilkan
4. Dapat diverifikasi—setiap item lulus/gagal
5. Dipimpin pengguna—pengguna membuat keputusan akhir, bukan AI

**Enam struktur dokumen:** seperti di atas

**Tujuh larangan:**
1. Tidak ada prosa bertele-tele
2. Tidak maju tahap tanpa konfirmasi pengguna
3. Tidak ada konten di luar Spec
4. ...

**Lima kriteria penyelesaian:**
1. Persetujuan pengguna di semua tahap
2. Setiap bagian sesuai Spec
3. ...

---

## 📄 Pembuatan Skill — Dan Verifikasi

Setelah Spec, Cocrates membuat `.opencode/skills/document-authoring/SKILL.md`.

Ia memverifikasi skill dengan **7 item checklist**.

**Hasil: semua PASS ✅**

| Cek | Hasil |
|-------|--------|
| Kernel terdefinisi jelas? | ✅ |
| Frame dirancang? | ✅ |
| Outline terperinci? | ✅ |
| Spec mandiri? | ✅ |
| Gerbang persetujuan di setiap tahap? | ✅ |
| Larangan jelas? | ✅ |
| Kriteria penyelesaian dapat diverifikasi? | ✅ |

---

## 🔄 Setelah Skill — Apa yang Berubah

Setelah skill ada, pengalaman berubah total.

**Sebelum:** *"Tulis laporan"* → AI menghasilkan liar → pengguna menerima pasif

**Sesudah:** *"Tulis retrospektif proyek"* → `document-authoring` aktif → P1 melalui tahap dengan kolaborasi pengguna → tinjau dan setujui setiap langkah → kualitas konsisten

Anda tidak lagi menunggu pasif output AI. Anda menjadi **desainer aktif** yang membentuk dan mengendalikan proses.

---

## ❄️ Pola Snowflake Bersama

Ep4 (belajar), Ep5 (generasi), Ep6 (pembuatan skill)—semuanya berbagi pola **Snowflake**.

| Tahap | Ep4 (Belajar) | Ep5 (Generasi) | Ep6 (Skill) |
|-------|----------------|------------------|-------------|
| Mulai kecil | Briefing konsep | Klarifikasi persyaratan | Definisikan Kernel |
| Perluas bertahap | 3 misi | 3 ADR | Frame→Outline→Spec |
| Tinjau & setujui | Setelah setiap misi | Setiap ADR disetujui | Setiap tahap disetujui |
| Simpan ke file | `kb/bloom-taxonomy.md` | `spec/jsondb.md` | `skills/document-authoring/SKILL.md` |

Pola ini berlaku konsisten di seluruh aktivitas Cocrates.

---

## 📌 Poin Penting

1. **Skill merancang cara AI bekerja.** Melampaui sekali pakai—definisikan perilaku AI. Bangun sekali, pakai ulang sebagai aset abadi.
2. **Jangan terima proposal AI tanpa tinjauan.** AI cenderung ke pilihan mudah dan generik. Gunakan U→A→E→A untuk kecocokan yang lebih baik.
3. **Snowflake adalah pola bersama.** Mulai kecil, perluas, tinjau dan setujui—belajar, generasi, skill.
4. **Dengan skill, Anda menjadi desainer aktif.** Tidak lagi menunggu pasif—Anda merancang dan mengendalikan.

---

**Tanyakan pada diri sendiri:**

- Bisakah Anda menjelaskan permintaan sekali pakai vs skill?
- Bisakah Anda menyebutkan enam tahap pembuatan skill?
- Bagaimana Snowflake berlaku di Ep4, Ep5, dan Ep6?

---

## 🎬 Selanjutnya

Tiga episode praktik mencakup aktivitas inti Cocrates.

Ep4: **Belajar** — Education, Knowledge Capture, Reflection.
Ep5: **Generasi** — ADR, Spec, Generation, Verification.
Ep6: **Pembuatan skill** — Kernel, Frame, Outline, Spec, Skill, Verification.

Anda sudah mengalami ketiganya.

Selanjutnya, Ep7: **struktur dan prinsip framework Cocrates—deep dive arsitektur.**

Anda sudah belajar **menggunakan** Cocrates; Ep7 membuka hood—Agent dan Skills, desain harness, Intent-To-Skill Routing.

> **"Apakah Anda akan meminta ikan pada AI—atau mengajarinya cara memancing?"**

Apa pun pilihan Anda, Cocrates berjalan bersama Anda.

---

*Seri ini memperkenalkan framework Cocrates Harness. Cocrates adalah agent harness yang dirancang untuk dialog Sokratik agar pengguna tetap memegang kendali dan terus berkembang.*
