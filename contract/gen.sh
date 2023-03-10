#!/usr/bin/env bash

abigen --abi=ZkBNB.json --pkg=core --type ZkBNB --out=../core/ZkBNBGen.go
abigen --abi=Governance.json --pkg=core --type Governance --out=../core/GovernanceGen.go
abigen --abi=ERC20.json --pkg=core --type ERC20 --out=../core/ERC20Gen.go
abigen --abi=ERC721.json --pkg=core --type ERC721 --out=../core/ERC721Gen.go
