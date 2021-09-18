package manifests

const AppTopolvmWebhook = "topolvm-webhook.yaml"

const AppTopolvmWebhookYaml = `
apiVersion: admissionregistration.k8s.io/v1beta1
kind: MutatingWebhookConfiguration
metadata:
  name: topolvm-hook
  annotations:
    cert-manager.io/inject-ca-from: topolvm-system/topolvm-mutatingwebhook
  labels:
    app.kubernetes.io/version: "0.9.0"
    app.kubernetes.io/name: topolvm-hook
webhooks:
  - name: pvc-hook.topolvm.cybozu.com
    admissionReviewVersions:
      - "v1"
      - "v1beta1"
    namespaceSelector:
      matchExpressions:
        - key: topolvm.cybozu.com/webhook
          operator: NotIn
          values: ["ignore"]
    failurePolicy: Fail
    matchPolicy: Equivalent
    clientConfig:
      service:
        namespace: topolvm-system
        name: topolvm-controller
        path: /pvc/mutate
    rules:
      - operations: ["CREATE"]
        apiGroups: [""]
        apiVersions: ["v1"]
        resources: ["persistentvolumeclaims"]
    sideEffects: None
  - name: pod-hook.topolvm.cybozu.com
    admissionReviewVersions:
      - "v1"
      - "v1beta1"
    namespaceSelector:
      matchExpressions:
        - key: topolvm.cybozu.com/webhook
          operator: NotIn
          values: ["ignore"]
    failurePolicy: Fail
    matchPolicy: Equivalent
    clientConfig:
      service:
        namespace: topolvm-system
        name: topolvm-controller
        path: /pod/mutate
    rules:
      - operations: ["CREATE"]
        apiGroups: [""]
        apiVersions: ["v1"]
        resources: ["pods"]
    sideEffects: None 
---

# Source: topolvm/templates/controller/csidriver.yaml
apiVersion: storage.k8s.io/v1
kind: CSIDriver
metadata:
  name: topolvm.cybozu.com
  labels:
    app.kubernetes.io/name: topolvm
    app.kubernetes.io/version: "0.9.0"
spec:
  attachRequired: false
  podInfoOnMount: true
  volumeLifecycleModes:
    - Persistent
    - Ephemeral
---
# Source: topolvm/templates/storageclass.yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: topolvm-provisioner
  annotations:
  labels:
    app.kubernetes.io/name: topolvm
    app.kubernetes.io/version: "0.9.0"
provisioner: topolvm.cybozu.com
parameters:
  "csi.storage.k8s.io/fstype": "xfs"
volumeBindingMode: WaitForFirstConsumer
allowVolumeExpansion: true
`
