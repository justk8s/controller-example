## CRD

CRD Design:


```yaml
apiVersion: study.justk8s.com/v1alpha1
kind: Student
metadata:
  name: mohamed
  namespace: insat
spec:
  phone: "1231234235"
  email: "me@justk8s.com"
  lastName: "Foulen"
  classRef:
    name: "dev-31"
    namespace: insat
```


```yaml
apiVersion: study.justk8s.com/v1alpha1
kind: Class
metadata:
  name: dev-31
  namespace: insat
spec:
  level: "A3"
```