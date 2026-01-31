To maintain your momentum toward the **May 1st CKA deadline** and your **Go Mastery book** project, we will apply the **Systems Over Goals** approach. This week focuses on **Module 1: The Physics of Video** while integrating your existing commitments to your family—Ron, Liv, and Niv.

---

### 📅 Week 1: The Foundation (Jan 31 - Feb 6)

| Day | FFmpeg Architect Task (15-20 min) | Skill Tutor "Habit Stack" |
| --- | --- | --- |
| **Sun** | **The Envelope vs. The Letter:** Research the internal structure of an MP4 container versus the H.264 codec. | **Weekly Planning:** Align this list with your Sunday weekly planning session. |
| **Mon** | **The Physics of I-Frames:** Understand why I-frames are the "anchors" of video compression. | **2-Minute Rule:** Just open your terminal and type `ffmpeg -formats` after the kids' showers. |
| **Tue** | **Chroma Subsampling:** Identify the difference between `yuv420p` and `yuv444p`. | **CKA Integration:** Spend 1 hour on CKA labs, then use FFmpeg for 10 mins as a "brain break". |
| **Wed** | **Bit Depth & Color:** Research how 8-bit vs. 10-bit color impacts file size and "banding." | **Implementation Intention:** "After dinner with Ron, Liv, and Niv, I will sit at the terminal". |
| **Thu** | **The Demuxer Stage:** Learn how FFmpeg "unwraps" a container before it even sees the video data. | **Friction Reduction:** Ensure your `alias ff='ffmpeg'` is set up in your `.zshrc` or `.bashrc`. |
| **Fri** | **Transcoding vs. Remuxing:** Practice the logic of `-c copy` (The "No-Touch" rule). | **Financial Management:** Complete your daily financial check, then reward yourself with 10 mins of Go/FFmpeg. |
| **Sat** | **Weekly Review:** Consolidate your notes on Containers vs. Codecs for your knowledge base. | **Rubber Ducking:** Explain the difference between a codec and a container to a "junior dev" (or the kids!). |

---

### 🛠️ The "Architect's" Daily Command

To prevent "magic snippet" culture, your goal this week isn't to run complex filters. It is to use the **Information Tool**.

**Your Daily Drill:**
`ffprobe -v error -show_format -show_streams input.mp4`

* **The Goal:** Identify which stream is the "Letter" (Codec) and which is the "Envelope" (Format).
* **The Trade-off:** Using `ffprobe` instead of `ffmpeg` prevents accidental encoding while you are still in the "Observation" phase.

---



