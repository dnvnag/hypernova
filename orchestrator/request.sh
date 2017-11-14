#!/bin/bash
curl -d '{"cloud":"hybrid cloud", "cloudtype":"amazon", "zone":"west"}' -H "content-Type:application/json" -X POST http://127.0.0.1:8080/infra/create
