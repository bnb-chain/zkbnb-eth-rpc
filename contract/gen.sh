#!/usr/bin/env bash

abigen --abi=ZkBNB.json --pkg=core --type ZkBNB --out=../core/ZkBNBGen.go
abigen --abi=Governance.json --pkg=core --type Governance --out=../core/GovernanceGen.go
abigen --abi=AssetGovernance.json --pkg=core --type AssetGovernance --out=../core/AssetGovernanceGen.go
abigen --abi=./ERC20.json --bin ./ERC20.bin --pkg=core --type ERC20 --out=../core/ERC20Gen.go
abigen --abi=./ERC721A.json --bin ./ERC721A.bin --pkg=core --type ERC721 --out=../core/ERC721Gen.go
