alidns-webhook
=================

Cert-manager ACME DNS webhook provider for alidns.

## Running the test suite

All DNS providers **must** run the DNS01 provider conformance testing suite,
else they will have undetermined behaviour when used with cert-manager.

**It is essential that you configure and run the test suite when creating a
DNS01 webhook.**

An example Go test file has been provided in [main_test.go]().

You can run the test suite with:

```bash
$ TEST_ZONE_NAME=example.com go test .
```

