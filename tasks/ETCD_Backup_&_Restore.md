To master the **ETCD Backup & Restore** workflow, you must treat the [Kubernetes Documentation](https://kubernetes.io/docs/home/) as your primary manual during the exam.

Here is the exact technical workflow, mapped to the official documentation links for your verification:

### **1. Preparation: Identify ETCD Control Plane Details**

Before running any commands, you must identify the certificates and endpoints used by the ETCD pod.

* **Action:** Run `kubectl describe pod etcd-controlplane -n kube-system`.
* **Documentation Reference:** [Operating etcd clusters for Kubernetes](https://kubernetes.io/docs/tasks/administer-cluster/configure-upgrade-etcd/)

### **2. The Snapshot (Backup)**

You will use the `etcdctl` command-line tool. Remember that in the exam environment, you may need to set `ETCDCTL_API=3`.

* **Action:** Execute the snapshot save command using the `--cacert`, `--cert`, and `--key` flags identified in step 1.
* **Command Template:** `ETCDCTL_API=3 etcdctl --endpoints=https://127.0.0.1:2379 --cacert=<ca-file> --cert=<cert-file> --key=<key-file> snapshot save /opt/snapshot-pre-bootcamp.db`
* **Documentation Reference:** [Backing up an etcd cluster](https://www.google.com/search?q=https://kubernetes.io/docs/tasks/administer-cluster/configure-upgrade-etcd/%23backing-up-an-etcd-cluster)

### **3. Verification**

Always verify that the snapshot is not empty or corrupted before proceeding.

* **Action:** Run `ETCDCTL_API=3 etcdctl --write-out=table snapshot status /opt/snapshot-pre-bootcamp.db`.
* **Documentation Reference:** This is found within the same [Backing up an etcd cluster](https://www.google.com/search?q=https://kubernetes.io/docs/tasks/administer-cluster/configure-upgrade-etcd/%23backing-up-an-etcd-cluster) section.

### **4. The Restore (The High-Stakes Step)**

This is where you restore the data to a **new** directory. Using a new directory prevents data corruption and allows for a rollback if the restore fails.

* **Action:** Use the `snapshot restore` command and specify a new `--data-dir`.
* **Command Template:**
`ETCDCTL_API=3 etcdctl --data-dir=/var/lib/etcd-from-backup snapshot restore /opt/snapshot-pre-bootcamp.db`
* **Documentation Reference:** [Restoring an etcd cluster](https://www.google.com/search?q=https://kubernetes.io/docs/tasks/administer-cluster/configure-upgrade-etcd/%23restoring-an-etcd-cluster)

### **5. Update the Static Pod Manifest**

To point the cluster to the restored data, you must edit the ETCD manifest file located at `/etc/kubernetes/manifests/etcd.yaml` on the control plane.

* **Action:** Update the `hostPath` for the `etcd-data` volume to point to your new directory (e.g., `/var/lib/etcd-from-backup`).
* **Documentation Reference:** [Static Pods](https://kubernetes.io/docs/tasks/configure-pod-container/static-pod/) (This explains how Kubernetes automatically restarts the component when the manifest is changed).

---

### **Green-Zone Check**

To confirm you are ready to execute: **Why is it critical to check the ETCD Pod description (Step 1) even if you think you know the certificate paths?**


---

k describe pod etcd-cka-labs-control-plane -n kube-system | less

```bash
      etcd
      --advertise-client-urls=https://172.19.0.2:2379
      --cert-file=/etc/kubernetes/pki/etcd/server.crt
      --client-cert-auth=true
      --data-dir=/var/lib/etcd
      --feature-gates=InitialCorruptCheck=true
      --initial-advertise-peer-urls=https://172.19.0.2:2380
      --initial-cluster=cka-labs-control-plane=https://172.19.0.2:2380
      --key-file=/etc/kubernetes/pki/etcd/server.key
      --listen-client-urls=https://127.0.0.1:2379,https://172.19.0.2:2379
      --listen-metrics-urls=http://127.0.0.1:2381
      --listen-peer-urls=https://172.19.0.2:2380
      --name=cka-labs-control-plane
      --peer-cert-file=/etc/kubernetes/pki/etcd/peer.crt
      --peer-client-cert-auth=true
      --peer-key-file=/etc/kubernetes/pki/etcd/peer.key
      --peer-trusted-ca-file=/etc/kubernetes/pki/etcd/ca.crt
      --snapshot-count=10000
      --trusted-ca-file=/etc/kubernetes/pki/etcd/ca.crt
      --watch-progress-notify-interval=5s
```

etcdctl and etcdutl are command-line tools used to interact with etcd clusters, but they serve different purposes:

etcdctl: This is the primary command-line client for interacting with etcd over a network. It is used for day-to-day operations such as managing keys and values, administering the cluster, checking health, and more.

etcdutl: This is an administration utility designed to operate directly on etcd data files, including migrating data between etcd versions, defragmenting the database, restoring snapshots, and validating data consistency. For network operations, etcdctl should be used.

https://kubernetes.io/docs/tasks/administer-cluster/configure-upgrade-etcd/#backing-up-an-etcd-cluster

```bash
#template:
ETCDCTL_API=3 ; ./etcdctl --endpoints=https://172.19.0.2:2379 \
  --cacert=ca.crt --cert=server.crt --key=server.key \
  snapshot save ./etcd-snapshot.db

guy@Lenovo-ideapad-330:~/Desktop/gemini_context_files/cka/etcd-v3.6.7-linux-amd64$ ETCDCTL_API=3 ; ./etcdutl --write-out=table snapshot status ./etcd-snapshot.db 
+----------+----------+------------+------------+---------+
|   HASH   | REVISION | TOTAL KEYS | TOTAL SIZE | VERSION |
+----------+----------+------------+------------+---------+
| 5600c7f9 |  1457477 |        302 |     8.1 MB |   3.6.0 |
+----------+----------+------------+------------+---------+

```



