To help you master **Taints, Tolerations, and Node Affinity**, I have reconstructed this lab based on your CKA curriculum and tactical schedules.

### **Lab Objective**

Master the ability to control exactly where pods are placed in the cluster by manipulating node restrictions and pod preferences.

---

### **Part 1: Taints and Tolerations (The Restriction)**

*Goal: Prevent pods from running on a node unless they have "permission."*

1. **Identify your nodes:** Run `k get nodes` to see your available targets.
2. **Apply a Taint:** Pick a node (e.g., `node01`) and apply a taint that prevents any pod from being scheduled there unless it has a matching toleration.
* **Command:** `kubectl taint nodes node01 tier=gold:NoSchedule`.


3. **The Failure Test:** Create a simple Nginx pod. Observe its status using `k get pods`. It should remain in a **Pending** state because it cannot bypass the taint.
4. **The Fix (Toleration):** Create a YAML for a pod that "tolerates" the gold tier.
* **Requirement:** Add the `tolerations` section to the pod spec to match `key: "tier"`, `operator: "Equal"`, `value: "gold"`, and `effect: "NoSchedule"`.


5. **Verify:** Apply the YAML and ensure the pod now moves to **Running** on `node01`.

---

### **Part 2: Node Affinity (The Preference)**

*Goal: Force a pod to run on a specific node based on labels.*

1. **Label a Node:** Give a node a specific characteristic.
* **Command:** `kubectl label nodes node01 disk=ssd`.


2. **Define Affinity:** Create a deployment that **requires** (not just prefers) being placed on a node with an SSD.
* **Requirement:** Use `nodeAffinity` with `requiredDuringSchedulingIgnoredDuringExecution`.
* **Logic:** Match expressions where the key is `disk` and the operator is `In` with the value `ssd`.


3. **Deploy and Confirm:** Use `kubectl get pods -o wide` to verify the pod landed on the node labeled `disk=ssd`.

---

### **Mentor Speed Challenge**

To align with your **"Imperative Mindset"** goal:

* **The Task:** Can you perform the "Failure Test" from Part 1 and identify *exactly* why the pod is pending using **only one command**?
* **The Hint:** Use `k describe pod [pod-name]` and look at the **Events** section.

**When you finish, explain the difference between a Taint and Node Affinity in your own words to complete your "Green-Zone Check".**



controlplane:~$ k get nodes
NAME           STATUS   ROLES           AGE     VERSION
controlplane   Ready    control-plane   6d19h   v1.34.3
node01         Ready    <none>          6d18h   v1.34.3

controlplane:~$ k taint nodes node01 tier=gold:NoSchedule
node/node01 tainted

controlplane:~$ k get nodes node01 -o yaml | less

apiVersion: v1
kind: Node
metadata:
  annotations:
    flannel.alpha.coreos.com/backend-data: '{"VNI":1,"VtepMAC":"ea:f5:7a:06:3d:6c"}'
    flannel.alpha.coreos.com/backend-type: vxlan
    flannel.alpha.coreos.com/kube-subnet-manager: "true"
    flannel.alpha.coreos.com/public-ip: 172.30.2.2
    node.alpha.kubernetes.io/ttl: "0"
    projectcalico.org/IPv4Address: 172.30.2.2/24
    projectcalico.org/IPv4IPIPTunnelAddr: 192.168.1.1
    volumes.kubernetes.io/controller-managed-attach-detach: "true"
  creationTimestamp: "2026-01-31T23:24:16Z"
  labels:
    beta.kubernetes.io/arch: amd64
    beta.kubernetes.io/os: linux
    kubernetes.io/arch: amd64
    kubernetes.io/hostname: node01
    kubernetes.io/os: linux
  name: node01
  resourceVersion: "3385"
  uid: a235692f-8fa3-4fd6-b5b3-eac4e5b7598f
spec:
  podCIDR: 192.168.1.0/24
  podCIDRs:
  - 192.168.1.0/24
  taints:
  - effect: NoSchedule
    key: tier
    value: gold

controlplane:~$ k get pods --watch
NAME   READY   STATUS    RESTARTS   AGE
site   0/1     Pending   0          26s

  tolerations:
    - key: "tier"
      operator: "Equal"
      value: "gold"
      effect:  "NoSchedule"

controlplane:~$ k get pod site
NAME   READY   STATUS    RESTARTS   AGE
site   1/1     Running   0          14s

controlplane:~$ kubectl label nodes node01 disk=ssd
node/node01 labeled

controlplane:~$ k run site --image=nginx 
pod/site created

apiVersion: v1
kind: Pod
metadata:
  annotations:
    cni.projectcalico.org/containerID: e472ccff36efcb0c35f7b5594fbfe6c7f90cc88eedc46d80f5385e45f7f57c5e
    cni.projectcalico.org/podIP: 192.168.1.4/32
    cni.projectcalico.org/podIPs: 192.168.1.4/32
  creationTimestamp: "2026-02-08T06:13:51Z"
  generation: 1
  labels:
    run: site
  name: site
  namespace: default
  resourceVersion: "4055"
  uid: 9577b8be-ac7e-4379-ab73-51c84eb57560
spec:
  affinity:
    nodeAffinity:
      requiredDuringSchedulingIgnoredDuringExecution:
        nodeSelectorTerms:
        - matchExpressions:
          - key: disk
            operator: In
            values:
            - ssd
  containers:
  - image: nginx
    imagePullPolicy: Always
    name: site
    resources: {}
    terminationMessagePath: /dev/termination-log
    terminationMessagePolicy: File
    volumeMounts:
    - mountPath: /var/run/secrets/kubernetes.io/serviceaccount
      name: kube-api-access-9cg7m
      readOnly: true

controlplane:~$ k get pods 
NAME   READY   STATUS    RESTARTS   AGE
site   1/1     Running   0          17s
