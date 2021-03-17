# app to show k8s
app | port
----|-------
sea | 5000
pirates | 6000

```
# Create secret for docker creds to push image to registry
kubectl create secret generic jenkins-docker-creds \
    --from-file=.dockerconfigjson=/home/<your profile>/.docker/config.json \
    --type=kubernetes.io/dockerconfigjson

# Create service account for jenkins deploy stage
kubectl create serviceaccount deployer -n jenkins

kubectl create rolebinding jenkins-deployer-edit --clusterrole=edit --serviceaccount=jenkins:deployer -n jenkins

kubectl create clusterrolebinding jenkins-deployer-admin --clusterrole=cluster-admin --serviceaccount=jenkins:deployer
```