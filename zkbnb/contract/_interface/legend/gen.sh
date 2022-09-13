@echo off

abigen --abi=zkbnb/contract/_interface/legend/OldZkBNB.json --pkg=ZkBNB --out=zkbnb/core/legend/ZkBNBGen.go
abigen --abi=zkbnb/contract/_interface/legend/Governance.json --pkg=Governance --out=zkbnb/core/legend/GovernanceGen.go