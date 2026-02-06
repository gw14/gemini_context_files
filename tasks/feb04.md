To help you master **Taints, Tolerations, and Node Affinity** for your CKA task today, here is a comprehensive guide on how these mechanisms control pod placement.

### 1. Taints and Tolerations

Taints are applied to **Nodes** to repel pods, while Tolerations are applied to **Pods** to allow them to be scheduled on tainted nodes.

* **Taints (Node Side):** A taint tells the node to only accept pods that have a matching toleration.
* **Command:** `kubectl taint nodes <node-name> key=value:effect`.
* **Effects:**
* `NoSchedule`: New pods won't be scheduled unless they tolerate the taint.
* `PreferNoSchedule`: The system tries to avoid placement but doesn't guarantee it.
* `NoExecute`: Evicts existing pods that do not have the matching toleration.




* **Tolerations (Pod Side):** Defined in the PodSpec to "ignore" or "match" a specific taint.
```yaml
tolerations:
- key: "key"
  operator: "Equal"
  value: "value"
  effect: "NoSchedule"

```



### 2. Node Affinity

Node Affinity is a property of **Pods** that attracts them to specific nodes based on labels. It is a more powerful evolution of `nodeSelector`.

* **RequiredDuringSchedulingIgnoredDuringExecution:** This is "Hard" affinity; the pod **must** meet the rule to be scheduled.
* **PreferredDuringSchedulingIgnoredDuringExecution:** This is "Soft" affinity; the scheduler will **try** to find a matching node, but will still schedule the pod elsewhere if no match exists.
* **Key Operators:** You can use logic like `In`, `NotIn`, `Exists`, `DoesNotExist`, `Gt`, and `Lt`.
```yaml
affinity:
  nodeAffinity:
    requiredDuringSchedulingIgnoredDuringExecution:
      nodeSelectorTerms:
      - matchExpressions:
        - key: disktype
          operator: In
          values:
          - ssd

```



### 3. Key Differences and Use Cases

* **Taints/Tolerations** are used to **keep pods away** from specific nodes (e.g., dedicated nodes for GPU or production workloads).
* **Node Affinity** is used to **pull pods** toward specific nodes (e.g., ensuring a pod lands on a node with high-performance storage).
* **Combined Use:** To ensure a pod lands *only* on a specific set of nodes and that *nothing else* lands there, you must use **both** a Taint (on the node) and Node Affinity (on the pod).

### Practice Exercise for Today

1. **Taint a Node:** `kubectl taint nodes <node-name> dedicated=experimental:NoSchedule`.
2. **Verify:** Try to run a standard pod; it should remain in a `Pending` state.
3. **Apply Toleration:** Create a pod with a matching toleration and verify it schedules on that specific node.
4. **Test Affinity:** Label a node (`kubectl label nodes <node-name> size=Large`) and use Node Affinity to force a pod to that node.
