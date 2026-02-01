## 🗓️ CKA-First Weekly Schedule: Feb 2 – Feb 8, 2026

This schedule implements the **Priority Tier System**: CKA is your morning engine, Golang is your logic bridge, and all other subjects are compressed into low-friction evening blocks.

---

### 🧱 Priority 1: CKA (The Heavy Lifter)

*Morning Window: 60–90 Minutes*

* **Feb 02 (Mon):** **Networking Deep-Dive.** Focus on Service types (ClusterIP vs. NodePort) and how Kube-Proxy manages iptables.
* **Feb 03 (Tue):** **Storage Mechanics.** Debug why a `PersistentVolumeClaim` is stuck in "Pending" status.
* **Feb 04 (Wed):** **Troubleshooting Drills.** Practice node-side fixes (Kubelet failure, swap issues).
* **Feb 05 (Thu):** **Security Logic.** RBAC Role vs. ClusterRole binding and ServiceAccount permissions.
* **Feb 06 (Fri):** **Workload Management.** Rolling updates and rollbacks using `kubectl rollout`.
* **Feb 07 (Sat):** **Speed Run.** Use a `kind` cluster to deploy a multi-container pod in under 3 minutes without copy-pasting YAML.
* **Feb 08 (Sun):** **Weekly Review.** Rubber ducking session: Explain the "Life of a Packet" in K8s to Ron, Liv, or Niv.

---

### 🥈 Priority 2: Golang (The Logic Layer)

*After-CKA Window: 30–45 Minutes*

* **Goal:** Move from **Phase 1 (Foundations)** into **Phase 2 (Concurrency)**.
* **Tasks:** * **Mon-Tue:** Complete the **Empty Interface** `interface{}` and **Pluggable Logger** labs.
* **Wed-Thu:** Transition to **Pointers** and advanced error handling using `errors.Is()` and `errors.As()`.
* **Fri-Sun:** Launch your first **Goroutines** and use `sync.WaitGroup` for a concurrent file scanner.



---

### 🥉 Priority 3: Supporting Stack (Background Tasks)

*Evening/Post-Dinner: 15–20 Minutes total (Combined)*

* **Philosophy:** These are 2-minute "Maintenance Sprints" to keep your tools sharp.
* **The "One-Command" Drill:**
* **FFmpeg:** Run `ffprobe` to identify a container vs. a codec.
* **PowerShell:** Create a `.psd1` manifest for your "Kube-Toolbox."
* **SQLite:** Define a 3-table schema for your CLI combat game.
* **Lua/Python:** Convert one global variable to a local/protocol-based structure in your config files.



---

### ⚠️ Resilience Check

If you are too tired after the **CKA morning engine**, do NOT force the Priority 3 tasks. Your "Minimum Viable Day" is **CKA only**.

**Would you like me to generate a specific CKA speed-drill lab for tomorrow morning to start the week with high momentum?**
