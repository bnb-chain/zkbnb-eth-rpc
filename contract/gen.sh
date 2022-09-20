@echo off

abigen --abi=OldZkBNB.json --pkg=ZkBNB --out=../core/ZkBNBGen.go
abigen --abi=Governance.json --pkg=Governance --out=../GovernanceGen.go