To create a roadmap that transforms you from "Zero to Hero," we will merge the technical rigor of the **FFmpeg Architect** with the metacognitive strategies of the **Skill Tutor**.

This plan is specifically designed to balance your **CKA study goals** (targeting May 1st) with your project to write a **Go Mastery book**, ensuring your "CPU" doesn't overheat.

---

## 📅 The 3-Month Architecture

### Month 1: The Foundation (Metacognition & Video Physics)

**Focus:** Building the mental infrastructure and understanding raw data.

* **Week 1-2: Physics of Video (Module 1).** Master the "Envelope vs. Letter" (Container vs. Codec) distinction. You will learn why a `.mp4` is just a box and how H.264/HEVC actually compresses pixels.
* **Week 3-4: The Syntax & Stream Selection (Module 2).** Moving away from "magic snippets". We will master the `-map` flag to explicitly control every stream, preventing the unpredictable defaults that plague junior engineers.
* **Skill Tutor Strategy:** **The 2-Minute Rule.** On days when you are exhausted from Go or CKA study, your only task is to open the terminal and check `ffmpeg -version`.

### Month 2: The Logic (Filtergraphs & Compression Science)

**Focus:** Transforming data and understanding the trade-offs of the "Heavy Lifters" phase.

* **Week 5-6: The Visual Filtergraph (Module 3).** You will learn to "program" video using nodes. We will use ASCII art to visualize stream flows: `[0:v] -> [scale] -> [out]`.
* **Week 7-8: Compression Science (Module 4).** Understanding the "Why" behind `-crf` and `-preset`. You will learn how to balance CPU time against file size and visual quality.
* **Skill Tutor Strategy:** **Socratic Interrogation.** I will ask you to explain why a specific bitrate was chosen, forcing you to use the **Feynman Technique** to expose any "fluency illusions".

### Month 3: The Professional (Audio & Production Pipelines)

**Focus:** Production-grade engineering and future-proofing your career.

* **Week 9-10: Audio & Sync (Module 5).** Mastering EBU R128 loudness normalization and fixing the dreaded A/V drift.
* **Week 11-12: Scaling & Production (Module 6).** Learning HLS/DASH segmentation for modern web streaming.
* **Skill Tutor Strategy:** **Cross-Pollination.** We will brainstorm how to write a **Go wrapper** for FFmpeg, potentially creating a new chapter for your **Go Mastery book** or a unique freelance offering.

---

## 🛠️ The Daily System

To ensure you reach the May 1st finish line for your CKA while doing this, we use **Habit Stacking**:

> "After I finish dinner with **Ron, Liv, and Niv**, I will sit at the terminal for **10 minutes** of FFmpeg Speed Drills."

---

### Your First Terminal Challenge (Module 1 Assessment)

Before we begin the technical deep-dive, you must demonstrate "First Principles" thinking.

**In your own words, explain the difference between a Container (The Envelope) and a Codec (The Letter).** How does this distinction change your approach when you need to change a file extension without re-encoding the video?

Would you like me to provide a visualization of the **FFmpeg Loop** (Input -> Decode -> Filter -> Encode -> Mux) to help ground your explanation?
