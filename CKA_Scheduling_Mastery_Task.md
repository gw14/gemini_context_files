# CKA Task: Master Pod Placement Control

**Objective:** Prove you can control exactly where pods are placed by using taints, tolerations, and node affinity.

---

### **Part 1: The Repulsion Lab (Taints and Tolerations)**

**Goal:** Use a taint to keep a node exclusive and a toleration to grant access.

1.  **Select a Target:** Choose one of your worker nodes to be your "special" node for this exercise.
2.  **Apply a Taint:**
    *   Use the command `kubectl taint nodes <your-node-name> gpu=true:NoSchedule`.
    *   This "repels" any pod that doesn't have a specific permission slip.
3.  **Verify Repulsion:**
    *   Create a generic `nginx` pod: `kubectl run generic-pod --image=nginx`.
    *   Check its status (`kubectl get pods`). Observe that it is `Pending`.
    *   Use `kubectl describe pod generic-pod` and look at the Events to see the scheduler's reason for not scheduling it (it will mention the taint).
4.  **Grant Permission:**
    *   Create a file `special-pod.yaml`.
    *   Generate the YAML for a pod (`kubectl run special-pod --image=nginx --dry-run=client -o yaml`).
    *   Manually add the following `toleration` to the pod's `spec` section to match the taint you created:
        ```yaml
        tolerations:
        - key: "gpu"
          operator: "Equal"
          value: "true"
          effect: "NoSchedule"
        ```
5.  **Verify Access:**
    *   Apply the manifest: `kubectl apply -f special-pod.yaml`.
    *   Verify that `special-pod` becomes `Running` while `generic-pod` remains `Pending`.
6.  **Cleanup:**
    *   Delete the pods: `kubectl delete pod generic-pod special-pod`.
    *   Remove the taint from the node: `kubectl taint nodes <your-node-name> gpu=true:NoSchedule-`.

---

### **Part 2: The Attraction Lab (Node Affinity)**

**Goal:** Use a label and node affinity to attract a pod to a specific node.

1.  **Apply a Label:**
    *   Choose a node (it can be the same one as before) and label it: `kubectl label nodes <your-node-name> disk=ssd`.
2.  **Create an Affinity Rule:**
    *   Create a file `affinity-pod.yaml`.
    *   Generate a basic pod YAML and add the following `affinity` block to the pod's `spec`:
        ```yaml
        affinity:
          nodeAffinity:
            requiredDuringSchedulingIgnoredDuringExecution:
              nodeSelectorTerms:
              - matchExpressions:
                - key: disk
                  operator: In
                  values:
                  - ssd
        ```
3.  **Verify Attraction:**
    *   Apply the manifest: `kubectl apply -f affinity-pod.yaml`.
    *   Check where the pod was scheduled: `kubectl get pods -o wide`. Confirm it landed on the node you labeled.
4.  **Cleanup:**
    *   Delete the pod: `kubectl delete pod affinity-pod`.
    *   Remove the label: `kubectl label nodes <your-node-name> disk-`.

---

### **Part 3: The Mastery Challenge (Combine Both)**

**Scenario:** You have a GPU node. You must ensure **only** GPU-enabled pods can run on it, and **no other** pods are allowed.

1.  **Isolate the Node:**
    *   Taint your chosen node to repel general workloads: `kubectl taint nodes <your-node-name> dedicated=gpu:NoSchedule`.
    *   Label the same node so it can be identified by special workloads: `kubectl label nodes <your-node-name> hardware=gpu`.
2.  **Deploy and Verify:**
    *   Deploy a generic pod (`kubectl run no-access --image=redis`). It should remain `Pending`.
    *   Create a final manifest, `gpu-workload.yaml`, for a pod that has **both**:
        *   A `toleration` for `dedicated=gpu:NoSchedule`.
        *   A `nodeAffinity` rule that requires the `hardware=gpu` label.
    *   Apply the `gpu-workload.yaml` manifest.
3.  **Final Confirmation:**
    *   Confirm that `no-access` is `Pending` and `gpu-workload` is `Running` **on the correct node**.
4.  **Final Cleanup:**
    *   Delete the pods.
    *   Remove the taint.
    *   Remove the label.

Completing this entire sequence demonstrates full control over pod placement.