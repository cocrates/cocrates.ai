---
sidebar_position: 11
---
# EP11. Skill Generasi Berbasis Arsitektur

![Skill Generasi Cocrates](/img/11-architecture-driven-generation-skills.png)

---

Dari Ep7 hingga Ep10 kita mengeksplorasi internal Cocrates—routing Agent, perilaku Skill, prinsip generasi berbasis arsitektur, Intent-To-Skill Routing.

Kita memahami **mengapa**.

Sekarang pertanyaan bergeser:

*"Saya paham prinsipnya. Skill apa yang sebenarnya dimiliki Cocrates, dan kapan masing-masing berjalan?"*

Hari ini kita menjawabnya. Ep11–Ep13 mengurai **skill inti pipeline generasi berbasis arsitektur**. Hari ini: **empat skill pilar** dan alur lengkap.

---

## 🎯 "Generate A" — Penalaran Internal Cocrates

Ketika Cocrates mendengar *"Generate A,"* apa yang terjadi di dalam?

**Pertanyaan pertama:** *"Apakah ada skill khusus untuk menghasilkan A?"*

*"Generate a blog series"* → `blog-series-authoring` ada. Gunakan. Sederhana.

**Pertanyaan kedua:** *"Jika tidak ada skill khusus?"*

Fallback ke **spec-driven-generation**—inti generik pipeline berbasis arsitektur.

Tapi spec-driven-generation bukan *"langsung bangun."* Ia mengajukan pertanyaan lebih halus.

**Pertanyaan ketiga:** *"Apakah Spec jelas?"*

Jika tidak—di sinilah benar-benar dimulai. Cocrates menelusuri:

1. **Identifikasi ASR** — apa yang secara struktural penting untuk artefak ini?
2. **ADR** — alternatif dan keputusan per ASR
3. **Spec** — gabungkan semua keputusan ADR ke satu dokumen mandiri
4. **Kesiapan Spec** — apakah Spec cukup konkret?
5. **Baru kemudian generasi** — setelah Spec lengkap

Itulah mengapa Cocrates tidak pernah mengizinkan *"bangun dulu, lihat nanti."*

---

## 🏛️ Empat Skill Pilar — Mesin Generasi

Pipeline berbasis arsitektur bertumpu pada empat skill:

### ① adr-writing — Rancang Keputusan

ADR (Architecture Decision Record) bukan log keputusan santai. **Satu Concern = satu ADR**—catat pilihan desain dengan audit trail.

Aturan inti:
- Bandingkan minimal 2–3 alternatif nyata
- Pro/kontra jelas per opsi
- Status: `proposed` → disetujui → `approved`
- Catat opsi yang ditolak juga—*"mengapa kita tidak memilih ini"* adalah aset masa depan

**Mengapa penting:** Tanpa ADR, alasan hilang. Nanti: *"Mengapa kita membangunnya begini?"*—tidak ada jawaban.

### ② spec-writing — Spesifikasi Hidup

Spec mengintegrasikan keputusan ADR ke dokumen **mandiri**.

Aturan inti:
- **Tanpa tautan ADR.** Spec saja harus menjelaskan semuanya. *"Lihat ADR 3"* dilarang.
- **Pernyataan dapat diverifikasi.** Bukan *"harus cepat"* tetapi *"harus menangani 1000 permintaan per detik"*—lulus/gagal jelas.
- **Dokumen hidup.** Persyaratan berubah → Spec berubah. Bukan tulis sekali buang.

### ③ spec-driven-generation — Spec Adalah Satu-satunya Sumber

Satu aturan besi dalam generasi: **Spec adalah satu-satunya otoritas.**

Jangan hasilkan apa yang tidak ada di Spec. Jangan tambah dari intuisi AI. Jika Spec tipis, perkuat Spec—jangan improvisasi.

Cocrates menjalankan **ASR Readiness Gate**—delapan kategori universal untuk memastikan Spec siap.

**Spec > prompt.** Seberapa pintar pun prompt, tidak bisa menggantikan Spec. Spec dapat diverifikasi, berversi, dapat direproduksi. Prompt menguap.

### ④ spec-driven-verification — Awal Siklus, Bukan Akhir

Verifikasi bukan hanya pengujian. **Dua tujuan:**

**Tujuan 1 — Deviasi:** Apakah output sesuai Spec? Temukan ketidaksesuaian.

**Tujuan 2 — ASR Tidak Terdokumentasi:** Keputusan struktural dalam output yang tidak ada di Spec—bukan kegagalan; kesempatan menyempurnakan Spec.

Setelah verifikasi, **gerbang tinjauan pengguna** tetap ada. AI memverifikasi belum selesai—Anda yang menyetujui.

---

## 💡 Wawasan Hari Ini: Skill sebagai Peran AI Khusus

Wawasan terbesar dari sistem skill Cocrates:

**Skill memberi AI peran khusus.**

- `adr-writing` → AI sebagai **arsitek** merancang keputusan
- `spec-writing` → AI sebagai **penulis spec** mendokumentasikan
- `spec-driven-generation` → AI sebagai **insinyur** mengimplementasikan ke spec
- `spec-driven-verification` → AI sebagai **insinyur QA** memverifikasi

Bukan satu AI memakai semua topi—**skill spesifik peran** mengikuti prosedur yang ditetapkan. Kualitas konsisten, hasil dapat diprediksi.

---

## 📌 Poin Penting

1. **Pipeline generasi memakai dua lapisan keputusan.** Skill khusus untuk A? → gunakan. Jika tidak spec-driven-generation → Spec jelas? → jika tidak, ASR → ADR → Spec dulu
2. **Empat skill pilar:** adr-writing, spec-writing, spec-driven-generation, spec-driven-verification
3. **Spec > prompt.** Spec dapat diverifikasi, berversi, dapat direproduksi; prompt volatil

---

**Tanyakan pada diri sendiri:**

- Bisakah Anda menyebutkan empat skill pilar Cocrates?
- Bisakah Anda menjelaskan alur internal spec-driven-generation saat fallback?
- Mengapa Spec melarang tautan ADR?

---

## 🎬 Selanjutnya

Hari ini kita meninjau **skill mana** yang membangun pipeline berbasis arsitektur.

Satu pertanyaan tetap ada:

*"Bagaimana jika skill yang saya butuhkan tidak ada?"*

Misalnya Anda butuh *"generator laporan code review rutin."* Tidak ada skill seperti itu di Cocrates. Lalu bagaimana?

**Ep12:** *"Cara membuat skill ketika tidak ada skill."*

Cocrates menawarkan **generating-skill-creation**—meta-skill. Lain kali kita mendalami skill yang membuat skill.

---

*Seri ini memperkenalkan framework Cocrates Harness. Cocrates adalah agent harness yang dirancang untuk dialog Sokratik agar pengguna tetap memegang kendali dan terus berkembang.*
