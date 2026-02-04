Here is your daily task list for **CKA**, **Golang**, and **PowerShell** for the week of February 2 – February 8, 2026, presented in bullet points:

* **Monday, Feb 02**
* **CKA:** Advanced Config Injection—Inject a ConfigMap as individual Environment Variables and use `envFrom` to inject all keys at once. - DONE
* **Golang:** Final Interface Review & Type Assertions—Finish the Empty Interface `interface{}` lab and practice the `v, ok := i.(T)` pattern.
* **PowerShell:** Refactoring "Disk Space Cleaner"—Convert your script into a reusable function with `-WhatIf` and `-Confirm` support using `SupportsShouldProcess=$true`.


* **Tuesday, Feb 03**
* **CKA:** Dynamic Updates & Speed Drill—Mount a Secret as a Volume and verify real-time updates; then practice creating a Secret-pulling Pod in under 2 minutes. - DONE
* **Golang:** Integration Lab—Build a Pluggable Logger that writes to multiple destinations (console, file, etc.) using the `io.Writer` interface.
* **PowerShell:** Robust Path Validation—Finalize a script that uses `[ValidateScript({Test-Path $_})]` to validate file paths before execution.


* **Wednesday, Feb 04**
* **CKA:** Scheduling—Master Taints, Tolerations, and Node Affinity to control exactly where pods are placed in the cluster.
* **Golang:** Pointers & Memory Management—Write functions to compare "Passing by Value" vs. "Passing by Reference" using pointer receivers.
* **PowerShell:** Manifest-Based Modules—Create your first Script Module (`.psm1`) and corresponding Module Manifest (`.psd1`).


* **Thursday, Feb 05**
* **CKA:** Cluster Maintenance—Master the `kubeadm` upgrade workflow, including upgrading `kubeadm`, running the upgrade plan, and updating `kubelet`/`kubectl`.
* **Golang:** Advanced Error Handling—Create a custom struct that implements the `Error()` interface and use `errors.Is()` and `errors.As()` for logic handling.
* **PowerShell:** Managing Module Exports—Experiment with "Private" vs. "Public" functions using `Export-ModuleMember`.


* **Friday, Feb 06**
* **CKA:** High-Stakes Task—Use `etcdctl` to perform an ETCD backup (snapshot) and a full restore to a new directory.
* **Golang:** Goroutines—Introduce the `go` keyword and use `sync.WaitGroup` to ensure the main function waits for concurrent tasks to finish.
* **PowerShell:** Advanced Functions (Proxying)—Create a proxy for `Get-Process` that adds custom default filtering behaviors.


* **Saturday, Feb 07**
* **CKA:** Weekend Integration—Practice the `drain` and `cordon` commands to safely remove workloads from a node and return it to service.
* **Golang:** Memory & Concurrency Review—Refactor your Pluggable Logger from Tuesday to utilize pointer receivers for its methods.
* **PowerShell:** Module Troubleshooting—Practice loading and unloading modules using `Import-Module -Force` and identifying session conflicts.


* **Sunday, Feb 08**
* **CKA:** Weekend Speed Drill—Build a multi-tier application where pods pull credentials from Secrets and UI themes from ConfigMaps in under 5 minutes.
* **Golang:** Integration Lab—Build a concurrent file scanner that searches multiple files for a keyword using goroutines and pointers.
* **PowerShell:** Integration Lab—Combine all of your Kubernetes-related functions into a single, cohesive module called `KubeToolbox`.
