Welcome to **Week 5**, the transition into the "Heavy Lifters" phase. Based on your roadmap, we are shifting our focus to **Services & Networking (ClusterIP, NodePort, DNS)**.

Since today is Wednesday, February 18th, 2026, we have a condensed window to master these high-priority topics before Sunday. Services and Networking account for **20% of your exam weight**.

---

## 🛰️ Week 5 Tactical Tasklist: Services & Networking

**Goal:** Master the "Service Plumbing" that connects your Pods to each other and the outside world.

### 🏋️ 1. The Daily "Service Speed Drill" (Repeated Daily)

*Goal: Successfully expose a deployment via all 3 major service types in under 3 minutes.*

* [ ] **Run 1: Internal Communication.** Create a deployment `app-v1` (image: `nginx`) and expose it via **ClusterIP** on port 80.
* [ ] **Run 2: External Access.** Expose the same deployment as a **NodePort** service on port 80, using a specific `nodePort` of 30080.
* [ ] **Run 3: Port Forwarding.** Practice accessing the service locally using `k port-forward svc/app-v1 8080:80`.
* [ ] **Run 4: DNS Verification.** Launch a temporary `busybox` pod and use `nslookup` to find the IP of your `app-v1` service.

### 📚 2. Core Concepts (The "Why" and "How")

*Study these until you can explain them in one sentence.*

* [ ] **ClusterIP (Default):** Provides an internal-only stable IP; use it for backend-to-frontend communication.
* [ ] **NodePort:** Opens a static port on *every* Node’s IP; use it for external traffic when no LoadBalancer is available.
* [ ] **LoadBalancer:** Requests an external IP from a cloud provider (standard for production).
* [ ] **Endpoint Objects:** These track the actual Pod IPs that match a Service's selector. If your Service isn't working, check `k get endpoints`.

### 🛠️ 3. Troubleshooting Lab (Friday & Saturday)

*Focus: Why is my service failing?*

1. **Selector Mismatch:** Create a service with a selector `app: nginx-v1` but label your Pods `app: nginx-v2`. Observe that the Service exists but has **No Endpoints**.
2. **Port Mapping:** Practice mapping different `port` (Service), `targetPort` (Pod), and `nodePort` (External Node) values.
3. **CoreDNS Check:** Check the logs of the CoreDNS pods in the `kube-system` namespace to ensure names are being resolved.

### 🧪 4. End-of-Week "Sunday Challenge"

*Challenge: Multi-Port Internal Connectivity*

1. Create a Deployment named `backend-api` with 3 replicas.
2. Expose it internally on port 8080 using the **Imperative Mindset** (`k expose ...`).
3. From a separate Pod in the **same namespace**, use `curl` to reach the `backend-api` using only its **DNS name**.
4. Change the service type to **NodePort** without deleting it (use `k edit svc ...` or `k patch ...`).

---

**Would you like to start with a practice scenario on "Selector Mismatches," or shall we jump straight into the YAML for a NodePort service?**

That YAML is **perfect**. You’ve successfully corrected the `targetPort` to **80** (matching the Nginx container) and manually assigned the `nodePort` to **30080**.

### 🔍 Validation Check

By setting `targetPort: 80`, the Service now knows exactly where to send traffic once it hits the Node.

* **External Request:** Hits any Node IP on port **30080**.
* **Service:** Routes that to the ClusterIP on port **80**.
* **Pod:** Receives the traffic on its container port **80**.

---

### 🏋️ Next Task: Run 4 — DNS Verification

As per your **Week 5 Tactical Tasklist**, you need to verify that your cluster's internal DNS (CoreDNS) can resolve this service name. In the CKA exam, verifying connectivity is just as important as creating the resource.

**Your Goal:** Launch a temporary `busybox` pod and use `nslookup` to find the IP of your `app-v1-np` service.

**Try this one-liner to verify:**

```bash
k run dns-test --image=busybox -i --rm --restart=Never -- nslookup app-v1-np

```

### 📚 Architect's Note on DNS

In Kubernetes, the standard DNS format for a service is:
`<service-name>.<namespace>.svc.cluster.local`.
Since you are in the `default` namespace, simply looking up `app-v1-np` should work.

**Does the output show the correct ClusterIP (10.107.80.87), or did it fail to resolve?**
