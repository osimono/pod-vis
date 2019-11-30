docker build . -t osimono/pod-vis:0.1
docker push osimono/pod-vis:0.1
kubectl apply -f k8s
kubectl delete po -l app=pod-vis
