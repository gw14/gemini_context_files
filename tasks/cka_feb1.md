Here are the official documentation links and specific "Must-Read" pages to support your tasks for **Sunday, February 1, 2026**:

### **1. Kubernetes (CKA) Documentation**

* **Official Home:** [Kubernetes Documentation](https://kubernetes.io/docs/home/)
* **The "Cheat Sheet":** [Kubectl Cheat Sheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/) — This is your best friend for the **5-minute Speed Drill**.
* **Secrets & ConfigMaps:**
* [Distribute Credentials Securely Using Secrets](https://kubernetes.io/docs/concepts/configuration/secret/)
* [Configure a Pod to Use a ConfigMap](https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/)

1. CKA: Weekend Speed Drill
The Goal: Build a multi-tier application with dynamic configuration in under 5 minutes.

Scenario: You are deploying a "Holiday Theme" web portal. The application requires database credentials (stored securely) and a UI "Primary Color" (stored as a configuration variable).

Implementation Steps:

Secret: Create a secret named db-creds containing db_user=admin and db_password=complexpassword123.

ConfigMap: Create a ConfigMap named ui-config with the key THEME_COLOR=festive-red.

Pod/Deployment: Define a pod that:

Uses envFrom to pull the db-creds Secret as environment variables.

Mounts the ui-config ConfigMap as a volume at /etc/config.

Verification: Run kubectl exec into the pod and verify echo $db_user returns "admin" and cat /etc/config/THEME_COLOR returns "festive-red".

k create secret generic db-creds --from-literal=db_user=admin --from-literal=db_password=complexpassword123 --dry_run=client -o yaml                                         
apiVersion: v1
data:
  db_password: Y29tcGxleHBhc3N3b3JkMTIz
  db_user: YWRtaW4=
kind: Secret
metadata:
 k create configmap ui-config --from-literal=THEME_COLOR=festive-red --dry-run=client -o yaml                                                                                 
apiVersion: v1
data:
  THEME_COLOR: festive-red
kind: ConfigMap
metadata:
  name: ui-config
  name: db-creds

https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/#populate-a-volume-with-data-stored-in-a-configmap

Populate a Volume with data stored in a ConfigMap 
Add the ConfigMap name under the volumes section of the Pod specification. This adds the ConfigMap data to the directory specified as volumeMounts.mountPath (in this case, /etc/config). The command section lists directory files with names that match the keys in ConfigMap.


https://kubernetes.io/docs/tasks/inject-data-application/distribute-credentials-secure/#define-a-container-environment-variable-with-data-from-a-single-secret

Define a container environment variable with data from a single Secret 
Define an environment variable as a key-value pair in a Secret:

kubectl create secret generic backend-user --from-literal=backend-username='backend-admin'

Assign the backend-username value defined in the Secret to the SECRET_USERNAME environment variable in the Pod specification.

apiVersion: v1
kind: Pod
metadata:
  name: env-single-secret
spec:
  containers:
  - name: envars-test-container
    image: nginx
    env:
    - name: SECRET_USERNAME
      valueFrom:
        secretKeyRef:
          name: backend-user
          key: backend-username



apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: front-back-end
  name: front-back-end
spec:
  replicas: 1
  selector:
    matchLabels:
      app: front-back-end
  strategy: {}
  template:
    metadata:
      labels:
        app: front-back-end
    spec:
      containers:
      - image: nginx
        name: nginx
        env:
        - name: db_user
          valueFrom:
            secretKeyRef: 
              name: db-creds
              key: db_user
        - name: db_password
          valueFrom:
            secretKeyRef: 
              name: db-creds
              key: db_password
        volumeMounts:
        - name: config-volume
          mountPath: /etc/config
        resources: {}
      volumes:
        - name: config-volume
          configMap:
            name: ui-config
status: {}


