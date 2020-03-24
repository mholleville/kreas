# Helm project by Kreas

The following elements constitute the project :

```
.
├── charts/                 # A directory containing any charts upon which this chart depends.
├── crds/                   # Custom Resource Definitions
├── templates/              # A directory of templates that, when combined with values, will generate valid Kubernetes manifest files.
├── .helmignore             # Source files (alternatively `lib` or `app`)
├── Chart.yaml              # A YAML file containing information about the chart
├── values.yaml             # The default configuration values for this chart
├── values-sample.yaml      # The sample configuration values for this chart
├── LICENSE                 # # OPTIONAL: A plain text file containing the license for the chart
└── README.md               # OPTIONAL: A human-readable README file
```

### Usage

```shell script
$ helm install . --name ${RELEASE_NAME} --namespace ${RELEASE_NAMESPACE} -f values-sample.yaml
```

NOTE : This command will deploy an `helm` chart with the `sample` values. The `sample` values will override the `default` values for `replicaCount` only.

### Usefull links

- [helm](https://helm.sh/)
- [charts](https://helm.sh/docs/topics/charts/)
- [tests](https://helm.sh/docs/topics/chart_tests/)
- [best practices](https://helm.sh/docs/chart_best_practices/conventions/)
- [helmenv](https://github.com/yuya-takeyama/helmenv)

