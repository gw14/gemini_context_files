apiVersion: v1
data:
  password: UEBzc3cwcmQxMjM=
kind: Secret
metadata:
  name: db-pass

Excellent work on your **Jan 27 (Tuesday)** Secret lab! You have successfully created a **Generic Secret** and captured the YAML output.

### **The Anatomy of your Secret**

In your `cka_notes_jan27.md` file, you can see the key difference between a ConfigMap and a Secret:

* **`kind: Secret`**: Tells Kubernetes to handle this object with extra care.
* **`data.password: UEBzc3cwcmQxMjM=`**: This is the **base64 encoded** version of your password.

### **Under the Hood: The "Base64" Truth**

As your CKA Mentor, I want to emphasize that your password `P@ssw0rd123` is now **obfuscated**, not encrypted.

* **Decryption Test**: If you run `echo "UEBzc3cwcmQxMjM=" | base64 --decode`, you will see your plain-text password immediately.
* **Exam Tip**: In the CKA exam, if you are asked to "create a secret," you must use the imperative command (`k create secret generic ...`) because it handles this encoding for you automatically.

### **📅 Looking at Jan 28 (Wednesday): Injection**

Now that you have your ConfigMap (`web-logic`) and your Secret (`db-pass`) created, tomorrow's task is to **inject** them into a Pod. You will practice:

1. Passing the ConfigMap as **Environment Variables**.
2. Observing how the application inside the Pod "sees" the configuration.

### **🧠 Mentor Challenge**

Look at the YAML you uploaded. Why is the `apiVersion` for a Secret still `v1`, the same as a ConfigMap? (Hint: Both are core, stable resources in the Kubernetes API).
