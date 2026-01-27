apiVersion: v1
data:
  interval: 30s
kind: ConfigMap
metadata:
  name: web-logic
apiVersion: v1
data:
  file.txt: |
    gghost=hiroshima
kind: ConfigMap
metadata:
  name: web-logic1


Great job running the dry-run! You've successfully generated the YAML for your **Jan 26** task. To elaborate on this and provide the background you requested:

### **Anatomy of your ConfigMap**

The output you shared reveals the core structure of a Kubernetes configuration object:

* **`kind: ConfigMap`**: This tells the API server that this object is used to store non-confidential data in key-value pairs.
* **`metadata.name: web-logic`**: This is the unique identifier for your configuration within the namespace.
* **`data`**: This is the most important section for your lab. Unlike a Pod or Deployment, which uses a `spec` section to define a desired state, a ConfigMap uses `data` to store the actual variables (in this case, `interval: 30s`).

### **Background: Why this matters for the CKA**

1. **Decoupling Configuration**: By using this `web-logic` ConfigMap, you can change the `interval` of your application without ever touching the application's code or rebuilding the container image.
2. **Imperative Speed**: In the exam, you won't have time to type this YAML from scratch. Using `kubectl create configmap ... --dry-run=client -o yaml` allows you to generate the base file instantly and then make quick edits if needed.
3. **Storage in the Cluster**: When you eventually `apply` this, the Kube-API Server takes this YAML, validates it, and stores the `data` in **ETCD**, which acts as the cluster's distributed database.

---

### **📅 Looking at Jan 27 (Tuesday): Secrets**

Now that you've mastered the ConfigMap (Plain Text), today’s task is to create a **Secret**.

**The Challenge:**
Create a Secret for a database password and compare it to your ConfigMap:
`k create secret generic db-pass --from-literal=password=P@ssw0rd123 --dry-run=client -o yaml`

**Observation Task:**
Look at the `data` field in the Secret's output. You will notice the password is not readable because Kubernetes automatically encodes Secret data in **base64**.

### **🧠 Mentor Challenge**

Since you've seen that ConfigMaps use the `data` field for plain text, and Secrets encode that same field:
**Is a Secret more "secure" than a ConfigMap, or is it just "obfuscated"?** (Hint: Think about how easy it is to decode base64).

### **🚀 Pro-Tip**

For your **Speed Drills**, remember that you can create a ConfigMap from an entire directory of configuration files at once using `--from-file=path/to/dir/`. This is a massive time-saver during the Troubleshooting section of the exam.


