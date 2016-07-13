# Garland

Toy to control a LED garland by producing and consuming messages into Kafka \o/.

## Getting started

Setup the environment variables B, K and T required by the [Qli](https://github.com/thbkrkr/qli/blob/master/client/client.go#L47) client.

To turn on the 10th LED in red, produce a message:
```
echo -n '{"10":"0,50,0"}' | oq
```