# HW kubernetes part 2
### Health check
```bash
curl -H 'Host: arch.homework' "http://$(minikube ip)/health"
# Output: {"status": "ok"}
```
### Rewrite target
```bash
curl -H 'Host: arch.homework' "http://$(minikube ip)/otusapp/alikhan murzayev/hello"
# Output: Requested url: /alikhan%20murzayev/hello
```
