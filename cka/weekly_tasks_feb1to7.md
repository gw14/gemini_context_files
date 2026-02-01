Based on your current progress and the **3-Week CKA Tactical Sprint**, here is your new **CKA Weekly Schedule** for **February 2 – February 8, 2026**.

I have moved your unfinished tasks from the previous week—specifically **Environment Variable injection**, **Secret Volumes**, and the **Speed Drill**—to the beginning of this week to ensure you have mastered **Configuration** before moving into **Scheduling & Maintenance**.

### 📅 CKA Weekly Schedule (Feb 2 – Feb 8, 2026)

**Theme:** *Scheduling, Maintenance, and Secure Configuration*

* **Feb 02 (Mon): Catch-up – Advanced Config Injection**
* **Task:** Inject a ConfigMap as individual **Environment Variables** and use `envFrom` to inject all keys at once. - DONE
* **Task:** Mount a Secret as a **Volume** in a Pod and read the "file" from within the container. - DONE


* **Feb 03 (Tue): Catch-up – Dynamic Updates & Speed Drill**
* **Task:** Update a ConfigMap and observe if/how the Pod's mounted volume updates in real-time.
* **Speed Drill:** Create a Pod that pulls its DB password from a Secret in under 2 minutes.


* **Feb 04 (Wed): Scheduling – Taints, Tolerations & Node Affinity**
* **Goal:** Control where pods live. Ensure you can restrict pods to specific high-performance nodes.
* **Lab:** Apply a taint to a node and observe a pod go into **Pending** state; then apply the correct toleration to fix it.


* **Feb 05 (Thu): Cluster Maintenance – Upgrades**
* **Goal:** Master the `kubeadm` upgrade workflow: Upgrade `kubeadm` -> `kubeadm upgrade plan` -> upgrade `kubelet`/`kubectl`.


* **Feb 06 (Fri): High-Stakes Task – ETCD Backup & Restore**
* **Goal:** Use `etcdctl` to snapshot the keyspace and restore it to a new directory.
* **Criticality:** This is a vital high-stakes exam task.


* **Feb 07 (Sat): Weekend Integration – Drain & Cordon**
* **Task:** Practice safely removing all workloads from a node before "maintenance" and then return it to service.


* **Feb 08 (Sun): Weekend Speed Drill – The Multi-Tier Challenge**
* **Task:** Build a multi-tier app where the web pod pulls its DB password from a Secret and its UI theme from a ConfigMap in under 5 minutes.



---

### 🧠 Performance Protocols

* **The 2-Minute Rule:** If you feel resistance, just run `alias k=kubectl` and `k get nodes`.
* **Friction Audit:** If you struggle with YAML, use the **Kubernetes Documentation "Cheat Sheet"** strategically.
* **Never Miss Twice:** Maintaining momentum is more important than perfect intensity.

**Mentor Challenge Reminder:** As you work on Tuesday's tasks, remember to investigate why a Pod might be **Pending** due to a **Taint** using only a single `kubectl` command (Hint: `k describe pod ...`).
