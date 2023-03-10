#!/usr/bin/env bash

abigen --abi=ZkBNB.json --pkg=core --type ZkBNB --out=../core/ZkBNBGen.go
abigen --abi=Governance.json --pkg=core --type Governance --out=../core/GovernanceGen.go
