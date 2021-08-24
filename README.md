# HW kubernetes part 2
### Output requested uri:
```bash
curl -H 'Host: arch.homework' "http://$(minikube ip)/hello-world/123/532"
# Output: Requested url: /hello-world/123/532
```
### Health check:
```bash
curl -H 'Host: arch.homework' "http://$(minikube ip)/health"
# Output: {"status": "ok"}
```
### Rewrite target:
```bash
curl -H 'Host: arch.homework' "http://$(minikube ip)/otusapp/alikhan murzayev/hello"
# Output: Requested url: /alikhan%20murzayev/hello
```
