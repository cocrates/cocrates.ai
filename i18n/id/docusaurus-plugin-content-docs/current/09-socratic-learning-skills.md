---
sidebar_position: 9
---
# EP9. Skill Pembelajaran Sokratik

![Skill Pembelajaran Cocrates](/img/09-socratic-learning-skills.png)

---

Di Ep8 kita mempelajari filosofi Learning Pipeline—kebidanan Sokratik, Bloom, ZPD—dan mengapa Cocrates menjawab dengan pertanyaan.

Sekarang: **bagaimana itu diimplementasikan dalam sistem.**

Di bawah `.opencode/skills/` ada tiga file SKILL.md: education, knowledge-capture, reflection. Masing-masing adalah spesifikasi alur kerja yang dirancang cermat—bukan prompt sederhana.

Mari kita buka satu per satu.

---

## 🎓 Education — Pelatih Belajar Sokratik 1:1

### Tanpa Spoon-feeding

Aturan pertama skill education:

> **Jangan pernah memberi jawaban lengkap atau solusi penuh dalam satu kali.**

Bahkan jika pengguna bertanya *"Jelaskan DIP,"* Cocrates tidak akan mengajar. Ide inti dalam 1–3 kalimat, analogi sehari-hari—di bawah 20% respons.

Dan ia selalu **berhenti tidak lengkap**—contoh cacat, celah terbuka—menunggu pengguna mengisinya.

### Misi Berbasis Giliran — Tiga Blok

Setiap respons education punya tiga blok:

**💡 Concept Briefing:** Inti dalam 1–3 kalimat, analogi sehari-hari. Di bawah 20% balasan.

**💻 Thought Laboratory:** Contoh praktis cacat atau skenario lanjutan. Membuat pengguna berpikir *"Tunggu—mengapa dibangun begini?"*

**🔥 MISSION:** Tepat satu tugas untuk giliran berikutnya—bukan jawaban singkat; jelaskan dengan kata-kata sendiri dan sebutkan apa yang Anda pelajari dan apa yang tidak Anda ketahui.

Contoh: pengguna bertanya *"Apa itu DIP?"*

> **💡 [Concept Briefing]**
> DIP seperti standar colokan listrik. Modul tingkat tinggi tidak boleh bergantung pada detail tingkat rendah—mereka bergantung pada antarmuka.
>
> **💻 [Thought Lab]**
> ```typescript
> class OrderService {
>   constructor(private db: MySQLDatabase) {}
> }
> ```
>
> **🔥 [MISSION]**
> Jika Anda mengganti `MySQLDatabase` dengan `PostgresDatabase`, apa yang rusak? Temukan satu pelanggaran DIP dan jelaskan arah ketergantungan dengan kata-kata sendiri.

Tidak ada jawaban diberikan. Tidak ada pratinjau misi berikutnya. Ia menunggu balasan pengguna.

### Matriks 2D Bloom + Push & Pull

Skill education menggunakan Bloom sebagai matriks 2D. Cocrates **tidak menaikkan kedua sumbu sekaligus**—satu langkah kognitif pada jenis pengetahuan yang sama, atau satu dimensi pengetahuan pada level kognitif yang sama.

Ia memilih **Push (langkah dari bawah)** atau **Pull (tantangan dulu)** berdasarkan kondisi pengguna. Pemula atau beban kognitif berlebih → Push. Pengetahuan sebelumnya ada → Pull dengan tantangan lebih tinggi.

---

## 📝 Knowledge Capture — Catat Ketidaktahuan

Belajar dari Education tidak bisa hidup hanya di memori. knowledge-capture menyimpan ke `kb/` sebagai **esensial berorientasi ingat**.

### Esensial, Bukan Kuliah

Aturan kunci: **jangan simpan catatan kuliah panjang, penjelasan, atau kode penuh.**

Satu bullet = satu unit memori—maksimal 1–2 kalimat. Contoh format:

```markdown
# Dependency Inversion Principle

## One-line Definition
Modul tingkat tinggi tidak boleh bergantung langsung pada implementasi tingkat rendah—mereka bergantung pada abstraksi (antarmuka).

## Key Points
- Modul tingkat tinggi tidak boleh bergantung langsung pada implementasi tingkat rendah
- Bergantung pada antarmuka; terima implementasi via injeksi
- DIP bukan hanya "pakai antarmuka"—arah ketergantungan yang penting

## Wrong Assumptions / Gaps
- DIP = selalu buat file antarmuka baru (X)
  → Di batas legacy, Adapter untuk menghubungkan juga valid
```

**Bagian `Wrong Assumptions / Gaps` adalah pusat**—mencatat apa yang Anda salah pahami adalah alat belajar terkuat.

### Strategi Merge

knowledge-capture selalu mencari `kb/` sebelum menyimpan. Jika file topik ada, ia **merge** alih-alih membuat file baru.

- Jangan pernah menimpa KB yang ada—hanya tambahkan wawasan baru
- Catat tanggal dan penambahan di Update History

Saat pembelajaran menumpuk, KB tumbuh dari catatan menjadi **basis pengetahuan pribadi**—input inti untuk Reflection.

---

## 🔍 Reflection — Mode Pewawancara

Tahap terakhir dan krusial Learning Pipeline.

### Persona Pewawancara

Fitur khas reflection adalah **persona**. Cocrates **bukan pelatih atau guru di sini—pewawancara**.

Tujuannya bukan mengajar—melainkan tahu persis apa yang pengguna tahu dan tidak tahu. Agar mereka bisa berkata: *"Saya tahu bahwa saya tidak tahu apa-apa."*

### KB sebagai Kriteria Evaluasi

Reflection menggunakan KB tersimpan sebagai **kriteria evaluasi**. Key Points menjadi checklist; Wrong Assumptions / Gaps memeriksa apakah celah benar-benar ditutup.

Penting: **KB bukan kunci jawaban.** Jika pemahaman pengguna saat ini berbeda tetapi koheren, sarankan perbarui KB dulu.

### Kedalaman Pertanyaan

reflection tidak meminta definisi.

- *"Jelaskan prinsip ini dengan contoh dari domain lain."*
- *"Di mana kode ini melanggar prinsip?"*
- *"Apa yang rusak jika Anda melakukan sebaliknya?"*

Pewawancara menilai: **Solid**—menjelaskan dengan kata sendiri, menerapkan dengan benar. **Partial**—arah benar, batas kabur. **Gap**—hanya hafalan, kontradiksi.

Output bukan nilai—melainkan **observasi dan penemuan**.

> ✅ Solid — Menjelaskan arah ketergantungan DIP dengan benar dengan contoh domain lain
> ⚠️ Goyah — Batas antara DIP dan Interface Segregation; kebingungan di mana mereka tumpang tindih

### Ketika Celah Muncul?

reflection tidak beralih ke mode mengajar. Ia bertanya sekali lagi dari sudut lebih sempit; jika masih lemah, catat ⚠️ dan sarankan sesi terpisah: *"Apakah Anda ingin sesi education untuk bagian ini?"*

---

## 🔄 Bagaimana Tiga Skill Terhubung

Skill ini tidak terisolasi—mereka membentuk satu siklus belajar:

```
Education → Knowledge Capture → Reflection → (jika perlu) Education
```

**Kondisi keluar** setiap skill **terhubung ke masuk** skill berikutnya.

1. Ketika **Education** mengonfirmasi pemahaman → sarankan **Knowledge Capture**: *"Apakah saya rapikan ini di KB?"*
2. Ketika **Knowledge Capture** selesai → sarankan **Reflection**: *"Ingin cek pemahaman berdasarkan ini?"*
3. Ketika **Reflection** menemukan celah → sarankan **Education** lagi: *"Ingin belajar lebih tentang ini?"*

Siklus berjalan hanya ketika pengguna berpartisipasi aktif. Cocrates mengusulkan; pengguna bergerak.

---

## 📌 Poin Penting

1. **Education berpusat pada No Spoon-feeding.** Tiga blok (Concept Briefing → Thought Lab → MISSION) memimpin pengguna menemukan.
2. **Knowledge Capture mencatat ketidaktahuan.** Esensial ingat dan Wrong Assumptions/Gaps—bukan ringkasan. Merge topik yang sama.
3. **Reflection adalah mode pewawancara.** KB sebagai kriteria; hasil terbagi ✅ solid vs ⚠️ goyah. Celah kembali ke Education.

---

**Tanyakan pada diri sendiri:**

- Saat mencatat apa yang saya pelajari dari AI, apakah saya menyimpan "ringkasan detail" atau "apa yang tidak saya ketahui"?
- Saat memeriksa pemahaman, apakah saya jujur mengakui area ⚠️ alih-alih pura-pura?

---

## 🎬 Selanjutnya

Kita sudah masuk ke dalam Learning Pipeline. Selanjutnya: aktivitas inti kedua Cocrates.

**Artifact Generation Pipeline**—pembuatan artefak berbasis arsitektur.

Mengapa *"langsung tulis saja"* berbahaya, apa itu ASR, bagaimana ADR dan Spec menjamin kualitas—dengan analogi rumah pedesaan.

---

*Seri ini memperkenalkan framework Cocrates Harness. Cocrates adalah agent harness yang dirancang untuk dialog Sokratik agar pengguna tetap memegang kendali dan terus berkembang.*
