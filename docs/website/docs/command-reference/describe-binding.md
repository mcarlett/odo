---
title: odo describe binding
---

## Description

`odo describe binding` command is useful for getting information about service bindings.

This command supports the service bindings added with the command `odo add binding` and bindings added manually to the Devfile, using a `ServiceBinding` resource from one of these apiVersion:
- `binding.operators.coreos.com/v1alpha1`
- `servicebinding.io/v1alpha3`

## Running the Command

There are 2 ways to describe a service binding:
- [Describe with access to Devfile](#describe-with-access-to-devfile)
- [Describe without access to Devfile](#describe-without-access-to-devfile)

### Describe with access to Devfile

This command returns information extracted from the Devfile and, if possible, from the cluster.

The command lists the Kubernetes resources declared in the Devfile as a Kubernetes component,
with the kind `ServiceBinding` and one of these apiVersion:
- `binding.operators.coreos.com/v1alpha1`
- `servicebinding.io/v1alpha3`

For each of these resources, the following information is displayed:
- the resource name,
- the list of the services to which the component is bound using this service binding,
- if the variables are bound as files or as environment variables,
- if the binding information is auto-detected.

When the service binding are not deployed yet to the cluster:

```shell
$ odo describe binding
ServiceBinding used by the current component:

Service Binding Name: my-nodejs-app-cluster-sample
Services:
 •  cluster-sample (Cluster.postgresql.k8s.enterprisedb.io)
Bind as files: false
Detect binding resources: true
Available binding information: unknown

Service Binding Name: my-nodejs-app-redis-standalone
Services:
 •  redis-standalone (Redis.redis.redis.opstreelabs.in)
Bind as files: false
Detect binding resources: true
Available binding information: unknown

Binding information for one or more ServiceBinding is not available because they don't exist on the cluster yet.
Start "odo dev" first to see binding information.
```

When the resources have been deployed to the cluster, the command also extracts information from the status of the resources to display information about the variables that can be used from the component.

```shell
$ odo describe binding 
ServiceBinding used by the current component:

Service Binding Name: my-nodejs-app-cluster-sample-2
Services:
 •  cluster-sample-2 (Cluster.postgresql.k8s.enterprisedb.io)
Bind as files: false
Detect binding resources: true
Available binding information:
 •  CLUSTER_PASSWORD
 •  CLUSTER_PROVIDER
 •  CLUSTER_TLS.CRT
 •  CLUSTER_TLS.KEY
 •  CLUSTER_USERNAME
 •  CLUSTER_CA.KEY
 •  CLUSTER_CLUSTERIP
 •  CLUSTER_HOST
 •  CLUSTER_PGPASS
 •  CLUSTER_TYPE
 •  CLUSTER_CA.CRT
 •  CLUSTER_DATABASE

Service Binding Name: my-nodejs-app-redis-standalone
Services:
 •  redis-standalone (Redis.redis.redis.opstreelabs.in)
Bind as files: false
Detect binding resources: true
Available binding information:
 •  REDIS_CLUSTERIP
 •  REDIS_HOST
 •  REDIS_PASSWORD
 •  REDIS_TYPE
```

### Describe without access to Devfile

```shell
odo describe binding --name <component_name>
```

The command extracts information from the cluster.

The command searches for a resource in the current namespace with the given name, the kind `ServiceBinding` and one of these apiVersion:
- `binding.operators.coreos.com/v1alpha1`
- `servicebinding.io/v1alpha3`

If a resource is found, it displays information about the service binding and the variables that can be used from the component.
