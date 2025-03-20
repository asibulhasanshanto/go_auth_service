# Build your image if you haven't already
docker build -t authservice:v1.0.0 .

# Load the image into your Kind cluster
kind load docker-image authservice:v1.0.0 --name=kind

### Step 1: Install Helm
If you haven't already installed Helm, you can do so by following the official installation guide:
https://helm.sh/docs/intro/install/

### Step 2: Create a Helm Chart Locally
Open your terminal and navigate to the directory where you want to create the Helm chart.

#### Run the following command to create a new Helm chart:

```bash
helm create my-chart
```
Replace my-chart with the name of your chart. This will generate a directory with the following structure:


```
my-chart/
├── Chart.yaml          # Metadata about the chart
├── values.yaml         # Default configuration values
├── charts/             # Directory for dependency charts
├── templates/          # Directory for Kubernetes manifest templates
│   ├── deployment.yaml
│   ├── service.yaml
│   ├── ingress.yaml
│   ├── _helpers.tpl    # Helper templates
│   └── ...             # Other template files
└── tests/              # Test files
```
### Step 3: Customize the Chart
- Edit Chart.yaml: Update the metadata for your chart, such as the name, version, and description.

```yaml
apiVersion: v2
name: my-chart
description: A Helm chart for my application
version: 0.1.0
appVersion: "1.0"
```
- Edit values.yaml: Customize the default values for your chart. These values will be used in the templates.

```yaml
replicaCount: 1
image:
  repository: nginx
  tag: "latest"
  pullPolicy: IfNotPresent
service:
  type: ClusterIP
  port: 80
```
- Edit Templates: Modify the files in the templates/ directory to define your Kubernetes resources (e.g., Deployment, Service, Ingress, etc.).

Example templates/deployment.yaml:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ .Release.Name }}-deployment
spec:
  replicas: {{ .Values.replicaCount }}
  selector:
    matchLabels:
      app: {{ .Release.Name }}
  template:
    metadata:
      labels:
        app: {{ .Release.Name }}
    spec:
      containers:
        - name: {{ .Release.Name }}
          image: "{{ .Values.image.repository }}"
          ports:
            - containerPort: 80
  ```
### Step 4: Validate the Chart
Run the following command to lint and validate your chart:

```bash
helm lint my-chart
```
### Step 5: Package the Chart (Optional)
If you want to package your chart for distribution, run:

```bash
helm package my-chart
```
This will create a .tgz file in the current directory.

### Step 6: Deploy the Chart to Kubernetes
Ensure you have access to a Kubernetes cluster and kubectl is configured.

Install the chart using Helm:

```bash
helm install my-release ./my-chart
```
Replace my-release with the name of your release.

Verify the deployment:

```bash
kubectl get pods
kubectl get services
```
### Step 7: Upgrade or Uninstall the Chart
To upgrade the chart after making changes:

```bash
helm upgrade my-release ./my-chart
```
To uninstall the chart:

```bash
helm uninstall my-release
```
Additional Tips
Use helm template my-chart to render the templates locally without deploying them.

Use helm show values my-chart to view the default values of your chart.

Add dependencies to your chart by editing the Chart.yaml file and placing dependent charts in the charts/ directory.

By following these steps, you can create, customize, and deploy a Helm chart to Kubernetes. Let me know if you need further assistance!