# go-ds-otelwrapper


Add opentelemetry to any [datastore](https://github.com/ipfs/go-datastore)

## Get started

Use the Wrap method to decorate your datastore.

```
myds := datastore.Datastore{}
newds := Wrap(&myds)

```

Every method starts a new span annotated by the method name
