---
sidebar_position: 10
---
# EP10. Prinsip Generasi Berbasis Arsitektur

![Prinsip Generasi Cocrates](/img/10-architecture-driven-generation-activity.png)

---

Di Ep8 dan Ep9 kita mengeksplorasi filosofi dan skill Learning Pipeline. Education → Knowledge Capture → Reflection—bertanya, catat ketidaktahuan, verifikasi sebagai pewawancara.

Sekarang aktivitas inti kedua Cocrates: **Artifact Generation Pipeline**—pembuatan artefak berbasis arsitektur.

Ingat: **AI yang mengerjakan.** Anda adalah **Insinyur AI-native (ketua tim)** yang meninjau dan memutuskan apa yang dihasilkan AI. Di era AI mengerjakan segalanya untuk Anda, nilai Anda adalah seberapa baik Anda **meninjau, menilai, dan memutuskan** pekerjaan AI.

Jika Learning adalah "cara mengenali ketidaktahuan dan menumbuhkan pengetahuan," Generation adalah **"cara Anda meninjau dan menyetujui struktur yang diusulkan AI, lalu membangun dari struktur itu."**

---

## 🚨 Perangkap "Langsung Tulis Saja"

Anda meminta AI: *"Tulis laporan."* Tiga puluh halaman dalam lima detik. Anda membacanya—logika tipis, kedalaman tidak merata per bagian. Sulit minta ulang; Anda menambal canggung dan kehilangan waktu.

Apa penyebabnya?

**Generasi tanpa struktur.**

Output tanpa struktur punya tiga masalah:

**Pertama, konsistensi dan logika lemah.** Tidak ada benang merah—AI mendaftar apa yang terdengar masuk akal saat itu. Kedalaman bervariasi; klaim tidak selaras.

**Kedua, Anda tidak bisa memahami atau menjelaskan deliverable.** *"Kelihatannya baik"* dan lanjut—tetapi *"Mengapa bagian ini distrukturkan begini?"* membuat Anda buntu. AI yang memutuskan struktur, bukan Anda.

**Ketiga, kotak hitam yang tidak bisa ditinjau.** Tinjauan butuh kriteria. Tidak ada struktur? Hanya *"apakah terlihat masuk akal?"*

Itulah *"langsung tulis saja"*—efektif **"tulis sembarang cara."**

---

## 🏗️ ASR — Apa Porselen Halusnya?

*"Oke—saya perlu merancang struktur dulu."*

Tapi "struktur" abstrak—seperti *"hidup baik"* tanpa langkah konkret.

Masuk **ASR (Architecturally Significant Requirement)**—persyaratan atau keputusan desain yang **secara material memengaruhi struktur, komposisi, dan kualitas** deliverable akhir. Dari arsitektur perangkat lunak, tetapi berlaku untuk **setiap jenis deliverable**—dokumen, slide, seri blog.

ASR penting karena tanpa meninjau ASR, AI mengisi celah dengan **Silent Defaults**—pilihan *"cukup masuk akal"* yang mungkin tidak sesuai niat Anda, dan Anda mungkin tidak tahu mereka dipilih.

---

## 🏠 Analogi Rumah Pedesaan — Tragedi Silent Defaults

Anda pindah. Kotak untuk dikemas.

Anda bilang ke pengangkut **"kemas dengan baik."** Mereka mengemas dengan cara mereka. Porselen halus di bawah buku berat.

Bukan niat buruk—hanya default mereka vs ide Anda tentang "kemas dengan baik."

Itulah **Silent Default**. Jika Anda tidak menentukan, AI mengisi dengan defaultnya—dan bisa berbeda dari niat Anda.

Struktur adalah memutuskan apa ke mana dan bagaimana. **ASR adalah mengidentifikasi mana porselen halusnya.**

Bayangkan membangun rumah pedesaan.

Anda bilang ke arsitek: *"Rancang rumah yang paling cocok dengan gaya hidup keluarga kami."*

Keputusan: satu lantai atau dua? Atap, teras atap, ruang loteng?

Masing-masing adalah **Concern** dan **ASR**—dampak nyata pada struktur dan kegunaan.

**Kasus 1 — Pengguna memutuskan jelas:**
> *"Bangun ruang loteng lantai dua."*
Keputusan dibuat. Tidak perlu ADR—catat di Spec.

**Kasus 2 — Pengguna mendelegasikan:**
> *"Pilih yang paling cocok dengan gaya hidup kami."*
**Perlu ADR.** Arsitek membandingkan opsi; pengguna meninjau dan memutuskan.

---

## 🔄 Pipeline Empat Tahap

Dalam Artifact Generation Pipeline Cocrates:

```
Identifikasi ASR → ADR → Spec → Generation & Verification
```

**Tahap 1 — Identifikasi ASR:** AI menampakkan ASR. Tugas Anda: **discernment**—mana yang benar-benar memengaruhi struktur vs noise. Tanpa ini, AI memperlakukan semuanya penting atau memutuskan semuanya sendiri.
> Keterampilan AI-native: **discernment** antara ASR penting dan persyaratan yang lebih rendah

**Tahap 2 — ADR (alternatif & keputusan):** AI menganalisis opsi per Concern. Tugas Anda: **judgment**—apakah analisis ini cukup? Opsi lain? Apa yang AI lewatkan? ADR mengaudit *mengapa Anda memutuskan begini*.
> Keterampilan AI-native: **judgment** untuk mengevaluasi analisis AI dan memutuskan optimal

**Tahap 3 — Spec (integrasikan keputusan):** AI menggabungkan ADR yang disetujui ke satu dokumen. Tugas Anda: **insight**—ada yang hilang? Seperti apa output akhir dari Spec ini? Celah mengembalikan Anda ke ADR. Spec harus **mandiri**—satu bacaan menjelaskan semuanya.
> Keterampilan AI-native: **insight** untuk memverifikasi kelengkapan Spec dan mengantisipasi output

**Tahap 4 — Generation & Verification:** AI menghasilkan dari Spec. Tugas Anda: **sikap verifikasi**—jangan percaya membabi buta. Periksa setiap item Spec terhadap output. Temukan **ASR Tidak Terdokumentasi**—pilihan struktural dalam output yang tidak ada di Spec. Kirim kembali ke ADR untuk tinjauan.
> Keterampilan AI-native: **sikap verifikasi**—Anda memiliki deliverable

---

## 💡 Apa yang Diberikan Pipeline Ini

Intinya: **Anda berpartisipasi aktif di setiap tahap.**

AI tidak meninjau untuk Anda. Anda memutuskan, meninjau, menyetujui. AI membantu.

Generasi berbasis arsitektur dimulai dengan **mengakui struktur kabur**. Identifikasi apa yang penting, analisis opsi, putuskan, integrasikan ke Spec, hasilkan dan verifikasi terhadap Spec.

Itulah mengapa Cocrates membuat "langsung buat sesuatu" terlihat kompleks—agar Anda tidak pernah menerima output yang tidak Anda pahami.

---

## 📌 Poin Penting

1. **AI bekerja; Anda (Insinyur AI-native) meninjau dan memutuskan.** Nilai Anda adalah seberapa baik Anda mengevaluasi proposal AI dan membuat keputusan akhir.
2. **Tahap ASR butuh discernment**—pisahkan persyaratan benar-benar struktural dari noise.
3. **Tahap ADR butuh judgment**—evaluasi apakah analisis dan proposal AI optimal; minta perbaiki keputusan.
4. **Tahap Spec butuh insight**—periksa kelengkapan; bayangkan output yang dihasilkan.
5. **Generation/Verification butuh sikap verifikasi**—jangan percaya membabi buta; Anda bertanggung jawab.

---

**Tanyakan pada diri sendiri:**

- Saat meminta dari AI, apakah saya meninjau struktur sebagai ketua tim—atau menerima proposal pertama?
- Ketika AI bilang "ini penting," apakah semuanya sama penting? Apakah saya membedakan apa yang penting?

---

## 🎬 Selanjutnya

Hari ini kita mempelajari prinsip Artifact Generation Pipeline—ASR, ADR, Spec, Verification—dan mengapa setiap tahap ada.

Selanjutnya: empat skill—adr-writing, spec-writing, spec-driven-generation, spec-driven-verification—bagaimana mereka bekerja di dalam, seolah membuka setiap SKILL.md.

---

*Seri ini memperkenalkan framework Cocrates Harness. Cocrates adalah agent harness yang dirancang untuk dialog Sokratik agar pengguna tetap memegang kendali dan terus berkembang.*
