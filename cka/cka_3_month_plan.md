To align with your **Phase 1: Foundations & Speed** schedule and your May 1st goal, I have refined your next three weeks into a high-intensity "Tactical Sprint." This plan shifts from basic imperative commands to the "Heavy Lifter" topics (Networking and Storage) while maintaining your **Traffic Light Mastery System**.

## 📅 3-Week CKA Tactical Sprint (Jan 26 – Feb 15)

The 2026 CKA curriculum places the highest weight on **Troubleshooting (30%)** and **Cluster Architecture (25%)**, followed by **Networking (20%)**.

---

### 🟢 Week 3 (Jan 26 – Feb 1): Configuration & Workloads

*Focus: Decoupling configuration and mastering imperative speed.*

* **Mon–Tue:** Practice creating **ConfigMaps** and **Secrets** from literal values and files. Compare the security of Secrets (base64) vs. ConfigMaps.
* **Wed–Thu:** Inject configurations into Pods via `env`, `envFrom`, and **Volumes**. Verify changes in real-time by mounting a Secret as a volume and reading the file inside the container.
* **Fri:** Perform **Rolling Updates** and **Rollbacks** on Deployments. Use `kubectl rollout history` and `--revision` to manage state.
* **Weekend Integration:** **Speed Drill.** Build a multi-tier app where the web pod pulls its DB password from a Secret and its UI theme from a ConfigMap in under 5 minutes.

---

### 🔵 Week 4 (Feb 2 – Feb 8): Scheduling & Cluster Maintenance

*Focus: Controlling where pods live and keeping the cluster alive.*

* **Mon–Tue:** Master **Taints & Tolerations** and **Node Affinity**. Ensure you can restrict pods to specific high-performance nodes.
* **Wed–Thu:** **Cluster Upgrades.** Practice the `kubeadm` upgrade workflow: Upgrade `kubeadm` -> `kubeadm upgrade plan` -> upgrade `kubelet`/`kubectl`.
* **Fri:** **ETCD Backup & Restore.** Use `etcdctl` to snapshot the keyspace and restore it to a new directory. This is a critical high-stakes exam task.
* **Weekend Integration:** Practice **Drain & Cordon**. Safely remove all workloads from a node before "maintenance" and then return it to service.

---

### 🟠 Week 5 (Feb 9 – Feb 15): Networking & Services

*Focus: Connectivity, Load Balancing, and Isolation.*

* **Mon–Tue:** Understand **Service Types**. Distinguish when to use `ClusterIP` (internal), `NodePort` (external access), and `LoadBalancer`.
* **Wed–Thu:** **Network Policies.** This is the "Zero-Trust" model. Practice creating policies that block all traffic except for specific ports from labeled pods.
* **Fri:** **CoreDNS & Ingress.** Troubleshoot DNS resolution using `nslookup` inside a pod. Configure an **Ingress Resource** to route traffic to different services based on URL paths.
* **Weekend Integration:** **The Connectivity Challenge.** Deploy two apps in different namespaces and write a NetworkPolicy that allows only the frontend pod to "talk" to the backend pod.

---

## 🛠️ Performance Protocols

* **The 2-Minute Rule:** If you feel resistance, just run `alias k=kubectl` and `k get nodes`. Once you see the cluster, the friction is gone.
* **Friction Audit:** If you struggle with YAML formatting, stop and bookmark the **Kubernetes Documentation "Cheat Sheet"**. Use the docs strategically during your labs.
* **Never Miss Twice:** If you skip a Monday lab, Tuesday is mandatory. Maintaining momentum is more important than perfect intensity.

### Mentor Challenge

Can you explain why a Pod in a **Pending** state might be related to a **Taint** on a node, and how you would investigate this using only one `kubectl` command?

[CKA Prep: Services and Networking Walkthrough](https://www.youtube.com/watch?v=rP-W3Tv3plw)
This video provides practical, hands-on walkthroughs for networking and service tasks that are essential for the CKA exam.