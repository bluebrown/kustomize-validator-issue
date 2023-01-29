# Kustomize Validator Issue

Running the validator via ustomize `validators` property, results in errors.

```console
$ make
Error: validator shouldn't modify the resource map: nodes unequal: 
 -- {"apiVersion":"v1","data":{"key":"value"},"kind":"ConfigMap","metadata":{"name":"myapp"}},
 -- {"apiVersion":"v1","data":{"key":"value"},"kind":"ConfigMap","metadata":{"annotations":{"config.kubernetes.io/path":"configmap_myapp.yaml","internal.config.kubernetes.io/path":"configmap_myapp.yaml"},"name":"myapp"}}

--
&resource.Resource{RNode:yaml.RNode{fieldPath:[]string(nil), value:(*yaml.Node)(0xc00177fea0), Match:[]string(nil)}, refVarNames:[]string(nil)}
------
&resource.Resource{RNode:yaml.RNode{fieldPath:[]string(nil), value:(*yaml.Node)(0xc001a234a0), Match:[]string(nil)}, refVarNames:[]string(nil)}
```
