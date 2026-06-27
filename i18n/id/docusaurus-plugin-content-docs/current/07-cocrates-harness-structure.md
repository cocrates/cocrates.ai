---
sidebar_position: 7
---
# EP7. Struktur Cocrates Harness

![Arsitektur Cocrates Harness](/img/07-cocrates-harness-structure.png)

---

Di Ep6 kita belajar mengajari AI "cara menulis laporan" alih-alih "tulis laporan."

Tapi Anda mungkin bertanya-tanya:

*"Bagaimana Cocrates dibangun agar bisa melakukan semua hal berbeda ini?"*

Ep4 belajar, Ep5 membuat, Ep6 pembuatan skill—waktunya melihat ke dalam. Seperti membuka hood setelah Anda belajar mengemudi.

---

## 🔪 "Bisakah Satu Pisau Melakukan Semua Tugas Dapur?"

Bayangkan dapur. Tomat untuk diiris, roti untuk dipotong, ikan untuk difilet.

Bisakah **satu pisau** melakukan semuanya?

Mungkin—tetapi tomat hancur, roti sobek, sisik ikan tetap menempel. Dapur profesional punya pisau terpisah karena alasan tertentu.

Cocrates sama.

Seseorang bilang "ajari saya Bloom's Taxonomy," seseorang bilang "tulis laporan," seseorang bilang "buat slide." Bisakah **satu prompt** menangani semua itu?

**Tidak.**

Setiap deliverable membutuhkan pendekatan struktural berbeda. Laporan butuh outline dan alur logika; slide butuh pesan utama per halaman; belajar butuh misi berbasis giliran. Satu prompt raksasa menjadi tidak terkelola.

Di sinilah filosofi desain Cocrates muncul.

---

## 🏛️ Cocrates = Agent + Skills

Cocrates bukan satu prompt besar. Ini **harness** dengan dua lapisan.

**Cocrates Agent** — prinsip bersama dan kontrol. Membaca intent, memilih skill, mengelola alur. Seperti komandan mengerahkan unit yang tepat.

**Skills** — prosedur konkret per tugas. Setiap skill hidup di file sendiri (`.opencode/skills/*/SKILL.md`), dapat ditambah dan diedit kapan saja. Seperti tim spesialis.

Kuncinya adalah **pemisahan**.

- Prinsip umum di Agent; prosedur per tugas di Skills
- Skills independen—ubah satu tanpa merusak yang lain
- Jenis permintaan baru? Tambah Skills. Agent tetap

---

## 🔍 Cocrates Agent — Enam Bagian

`cocrates.md` adalah prompt dengan enam bagian.

**1. Persona:** Bagaimana agen mendefinisikan dirinya. *"Entitas yang mengubah ketidakpastian menjadi penyelidikan sistematis dan memandu desain berbasis struktur, tinjauan, dan persetujuan hingga pengguna sepenuhnya memahami deliverable."* Itulah identitas Cocrates.

**2. Principle:** Paling penting—**Harness Ignorance**. Jangan biarkan pengguna menerima output yang tidak mereka pahami; gunakan pertanyaan untuk menampakkan asumsi dan celah.

**3. Harness Architecture:** *"Cocrates adalah harness agen inti plus skills."* Agent menangani prinsip, intent, pemilihan skill, manajemen tugas, guardrail; Skills menangani alur kerja per tugas.

**4. Request Handling:** Infer intent, muat skill yang tepat, jalankan alur kerjanya. Menggunakan **Intent-To-Skill Routing**—belajar → education, capture → knowledge-capture, evaluasi → reflection, keputusan → adr-writing, dll. Delapan intent dipetakan ke delapan skill.

**5. Core Activities:** Dua pipeline:

- **Generation Pipeline:** Desain (ADR → Spec) → generasi berbasis Spec → verifikasi berbasis Spec
- **Learning Pipeline:** Education → Knowledge Capture → Reflection

**6. Success Criteria:** Ketika percakapan berakhir—apakah pengguna lebih jelas tahu apa yang mereka ketahui dan tidak ketahui? Bisakah mereka menjelaskan struktur deliverable? Apakah mereka berpartisipasi aktif di setiap langkah?

---

## 💡 "Gunakan → Pahami → Evolusi"

Sifat terpenting Cocrates: **ini bukan harness yang sudah jadi.**

**Pengguna mengevolusinya.**

Mulai dengan skill bawaan. Gunakan dan pikirkan *"Alur kerja ini tidak cocok untuk kasus saya."* Edit skill. Butuh jenis tugas baru? Bangun skill.

Putaran itu adalah nilai sebenarnya Cocrates.

```
Use → Understand → Evolve
```

Pertahankan prinsip dan persona Agent; perluas melalui Skills. Satu prompt raksasa terlihat sederhana di awal tetapi semakin sulit dirawat. Agent + Skills menyelesaikan itu dari awal.

---

## 📌 Poin Penting

1. **Satu prompt tidak bisa mencakup setiap deliverable.** Struktur berbeda per output; satu prompt besar tidak terkelola.
2. **Cocrates = Cocrates Agent + Skills.** Agent untuk prinsip dan kontrol; Skills untuk prosedur per tugas. Pemisahan adalah intinya.
3. **Cocrates tidak jadi—ia berevolusi.** Gunakan → pahami → evolusi; pengguna menumbuhkan sistem.

---

**Tanyakan pada diri sendiri:**

- Apakah alat AI Anda satu prompt atau Agent + Skills?
- Ketika jenis permintaan baru muncul, seberapa mudah Anda memperluasnya?

---

## 🎬 Selanjutnya

Hari ini kita meninjau struktur Cocrates. Agent + Skills, dua pipeline—Generation dan Learning.

Selanjutnya kita mendalami setiap pipeline, dimulai dengan **Learning**.

Mengapa *"ceritakan"* berbahaya, mengapa Cocrates menjawab dengan pertanyaan—kita mengurai filosofi dan mekanismenya.

---

*Seri ini memperkenalkan framework Cocrates Harness. Cocrates adalah agent harness yang dirancang untuk dialog Sokratik agar pengguna tetap memegang kendali dan terus berkembang.*
