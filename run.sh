#!/bin/bash
for i in {1..20}; do
  go run cmd/channels/main.go select-ready
done