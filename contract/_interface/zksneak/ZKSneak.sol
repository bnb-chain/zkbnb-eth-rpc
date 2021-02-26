// SPDX-License-Identifier: GPL3.0
pragma solidity ^0.7.6;
pragma experimental ABIEncoderV2;


contract ZKSneak{

    bytes32 public accountRoot;

    event Deposit(bytes32 addr,uint256 amount);

    event NewBlock(bytes32 oldRoot,bytes32 newRoot);

    function deposit(bytes32 pk) public payable{
        require(msg.value > 0,"you should deposit some money");
        emit Deposit(pk, msg.value);
    }

    function verifyBlock(bytes32 newRoot) public{
        accountRoot = newRoot;
        emit NewBlock(accountRoot,newRoot);
    }

    function withdraw(address payable _addr,uint256 amount) public{
        require(address(this).balance > amount,"too much money");
        _addr.transfer(amount);
    }

}