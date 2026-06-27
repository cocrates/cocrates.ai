---
sidebar_position: 12
---
# EP12. Skill Pembuatan Alur Kerja Berbasis Arsitektur

![Prinsip Pembuatan Skill Cocrates](/img/12-workflow-creation-skill.png)

---

Sejauh ini kita mempelajari dua mode generasi Cocrates.

**Pertama:** Jika skill khusus ada untuk A, gunakan. *"Generate a blog series"* → `blog-series-authoring`. Cepat dan sederhana.

**Kedua:** Jika tidak, fallback ke **spec-driven-generation**. ASR → ADR → Spec → Generation → Verification.

Tapi setelah Ep11, Anda mungkin bertanya-tanya:

*"Bagaimana jika skill yang saya butuhkan sama sekali tidak ada?"*

Misalnya Anda butuh *"hasilkan laporan code review rutin."* Tidak ada skill seperti itu. Anda bisa fallback setiap kali—atau **bangun skill sekali** jika pola akan diulang.

Hari ini: **generating-skill-creation**—skill untuk membuat skill.

---

## 🔍 Meta-Skill: Skill untuk Membuat Skill

generating-skill-creation adalah **meta-skill** Cocrates. Berbeda dari skill normal yang menghasilkan artefak, ia **merancang dan membuat skill itu sendiri**.

Bandingkan dengan analogi:

- Skill normal: *"Tangkapkan ikan"* → teknik memancing
- Meta-skill: *"Ajari saya memancing"* → mengajarkan memancing

generating-skill-creation berjalan pada **lima Artifact Components**:

**Lima Artifact Components:**

1. **Assignment & Constraints:** Masalah dan batasan apa yang diselesaikan skill ini?
2. **Context & Rules:** Konteks dan aturan apa yang mengatur skill?
3. **Entities:** Konsep dan objek apa yang ditangani skill?
4. **Space & Placement:** Di mana output hidup dan bagaimana disusun?
5. **Structure & Flow:** Apa alur dan struktur tahap demi tahap?

Mendefinisikan kelimanya adalah langkah pertama pembuatan skill.

---

## ❄️ Lima Tahap Snowflake — Konkretisasi Bertahap

Setelah Components teridentifikasi, Cocrates menyempurnakan skill melalui **lima tahap Snowflake**:

**Tahap 1 — Define:** *"Apa yang dilakukan skill ini?"* Definisikan Kernel dalam satu kalimat.

**Tahap 2 — Plan:** *"Dalam urutan apa?"* Rencanakan tahap keseluruhan dan urutan.

**Tahap 3 — Architecture Design:** *"Bagaimana komponen saling berhubungan?"* Rancang Entities, Space, Flow.

**Tahap 4 — Detail Design:** *"Apa yang persis terjadi setiap tahap?"* Input, output, kondisi persetujuan, larangan.

**Tahap 5 — Generation:** *"Hasilkan SKILL.md."* Hanya setelah semua desain selesai.

Satu aturan krusial:

---

## 🚫 Aturan Mutlak — Tidak Ada Generasi Sebelum Detail Design Selesai

Di antara prinsip generasi Cocrates, ini ditegakkan paling ketat:

> **"Jangan pindah ke Generation sampai Detail Design sepenuhnya selesai."**

Mengapa? Prinsip Ep2: *"The unexamined code is not worth generating."* Skill secara permanen mendefinisikan perilaku AI. Skill buruk berulang kali menghasilkan hasil buruk.

Jadi dalam Detail Design, Cocrates memeriksa secara ketat:

- Apakah input dan output jelas per tahap?
- Gerbang persetujuan di setiap tahap?
- Apakah larangan konkret?
- Bagaimana pengecualian ditangani?

Tidak satu baris SKILL.md sampai setiap pertanyaan terjawab.

---

## 🔄 Tujuh Tahap Meta Snowflake — Gambaran Lebih Besar

Lima tahap Snowflake adalah proses desain skill dasar. generating-skill-creation melukiskan gambaran lebih besar: **tujuh tahap Meta Snowflake**.

```
Kernel → Components → Frame → Outline → Spec → Skill → Verification
```

Secara berurutan:

**Kernel:** Satu kalimat mendefinisikan tujuan skill.

**Components:** Identifikasi dan definisikan lima Artifact Components.

**Frame:** Struktur dan kerangka keseluruhan—tahap mana yang ada.

**Outline:** Detail Frame—sub-item dan alur per tahap.

**Spec:** Spec mandiri mengintegrasikan semua keputusan. Siapa pun yang hanya membaca Spec memahami perilaku penuh.

**Skill:** Hasilkan SKILL.md dari Spec.

**Verification:** Apakah skill sesuai Spec? Ada yang hilang?

Bukan urutan murni atau paralel—setiap tahap bergantung pada sebelumnya; umpan balik verifikasi bisa kembali ke tahap lebih awal.

---

## 🎯 Penyempurnaan Per Tahap — Strategi Lazy Refinement

Satu prinsip penting dalam tujuh tahap Meta Snowflake:

**Components tidak terkonkretisasi dengan kecepatan sama.**

Beberapa sangat spesifik di awal; beberapa baru jelas di akhir. Cocrates menyebut ini **Lazy Refinement**.

Contoh: *Entities* mendapat bentuk kasar di Kernel/Components, detail di Outline/Spec. *Structure & Flow* butuh detail dari Frame ke depan.

Mengapa? **Hindari keputusan prematur.** Penguncian awal membuat opsi lebih baik nanti sulit diadopsi. Lazy refinement menunda keputusan untuk menjaga fleksibilitas.

---

## 💡 Wawasan Hari Ini: Skill adalah SOP AI

Skill adalah memberi AI **SOP (Standard Operating Procedure)**.

Ketika karyawan baru bergabung, Anda ajarkan alur kerja: *"Ketika permintaan ini datang, cek A, tinjau B, lalu lakukan C."*

Skill melakukan itu untuk AI: *"Untuk jenis permintaan ini, bekerja begini."*

Setelah didefinisikan, skill adalah aset yang dapat dipakai ulang—hari ini, minggu depan, setahun dari sekarang, kualitas sama.

**Permintaan sekali pakai adalah konsumsi. Skill adalah aset.**

---

## 📌 Poin Penting

1. **generating-skill-creation adalah meta-skill** berbasis lima Artifact Components.
2. **Lima tahap Snowflake:** Define → Plan → Architecture Design → Detail Design → Generation—dengan gerbang persetujuan antar tahap.
3. **Aturan mutlak: tidak ada Generation sebelum Detail Design selesai.** Skill buruk melipatgandakan output buruk.
4. **Tujuh tahap Meta Snowflake:** Kernel → Components → Frame → Outline → Spec → Skill → Verification.
5. **Lazy refinement:** jangan konkretisasi semuanya sekaligus—hanya saat dibutuhkan.

---

**Tanyakan pada diri sendiri:**

- Bisakah Anda menjelaskan Snowflake lima vs Meta Snowflake tujuh?
- Bisakah Anda menyebutkan dan menjelaskan lima Artifact Components?
- Mengapa Lazy Refinement penting?

---

## 🎬 Selanjutnya

Kita sudah mencakup seluruh Cocrates.

Ep1–2 filosofi dan prinsip, Ep3 instalasi, Ep4–6 praktik, Ep7–10 internal, Ep11–12 skill inti dan meta.

**Kita siap menggunakan Cocrates.**

Tapi satu pertanyaan penting tetap ada:

*"Apakah itu akhirnya? Apakah 'menggunakan Cocrates dengan baik' cukup?"*

Tidak. Bagian sebenarnya dimulai sekarang. Menggunakan Cocrates saja tidak cukup.

**Ep13:** Pengguna harus berevolusi; Cocrates Harness juga harus berevolusi—dan keduanya membentuk loop umpan balik yang saling memperkuat.

*"Menggunakan Cocrates saja tidak cukup—Anda harus mengevolusikan Cocrates."*

---

*Seri ini memperkenalkan framework Cocrates Harness. Cocrates adalah agent harness yang dirancang untuk dialog Sokratik agar pengguna tetap memegang kendali dan terus berkembang.*
