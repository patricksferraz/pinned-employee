# Kubernetes Deployment Guide

This guide provides step-by-step instructions for deploying the Pinned Employee application to a Kubernetes cluster.

## 📋 Prerequisites

- Kubernetes cluster up and running
- `kubectl` configured and connected to your cluster
- Docker registry access (if using private images)
- Basic understanding of Kubernetes concepts

## 🔐 Setting Up Secrets

### 1. Application Secrets

1. Navigate to the `k8s` directory:
```bash
cd k8s
```

2. Create your environment file:
```bash
cp .env.example .env
```

3. Edit the `.env` file with your configuration values

4. Create the Kubernetes secret:
```bash
kubectl create secret generic employee-secret \
  --from-env-file .env \
  --namespace <your-namespace>
```

### 2. Docker Registry Secrets (Optional)

If you're using a private Docker registry, create a registry secret:

```bash
kubectl create secret docker-registry regsecret \
  --docker-server=$DOCKER_REGISTRY_SERVER \
  --docker-username=$DOCKER_USER \
  --docker-password=$DOCKER_PASSWORD \
  --docker-email=$DOCKER_EMAIL \
  --namespace <your-namespace>
```

Required environment variables:
- `DOCKER_REGISTRY_SERVER`: Registry URL (e.g., docker.io, gcr.io)
- `DOCKER_USER`: Registry username
- `DOCKER_PASSWORD`: Registry password
- `DOCKER_EMAIL`: (Optional) Your email address

## 🚀 Deployment

### 1. Verify Configuration

Before deploying, verify your Kubernetes configurations:
```bash
kubectl apply -f ./k8s --dry-run=client
```

### 2. Deploy Application

Deploy all resources:
```bash
kubectl apply -f ./k8s
```

### 3. Verify Deployment

Check the status of your deployment:
```bash
kubectl get all -n <your-namespace>
```

## 🔍 Troubleshooting

### Common Issues

1. **Secret Creation Fails**
   - Ensure all required environment variables are set
   - Verify you have the necessary permissions in the namespace

2. **Pod Fails to Start**
   - Check pod logs: `kubectl logs <pod-name> -n <your-namespace>`
   - Verify secrets are properly mounted: `kubectl describe pod <pod-name> -n <your-namespace>`

3. **Image Pull Errors**
   - Verify registry credentials
   - Check if the image exists in the registry
   - Ensure the image pull policy is correct

## 🧹 Cleanup

To remove all deployed resources:
```bash
kubectl delete -f ./k8s
```

To remove secrets:
```bash
kubectl delete secret employee-secret -n <your-namespace>
kubectl delete secret regsecret -n <your-namespace>
```

## 📚 Additional Resources

- [Kubernetes Documentation](https://kubernetes.io/docs/home/)
- [Kubernetes Secrets Management](https://kubernetes.io/docs/concepts/configuration/secret/)
- [Docker Registry Authentication](https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/)
