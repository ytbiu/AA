// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "./base/PaymasterBase.sol";

contract USDTPaymasterNonProxy is PaymasterBase {
    address private _owner;

    constructor(
        address usdtTokenAddr,
        address oracleAddr,
        address _feeRecipient,
        address initialOwner
    ) {
        _owner = initialOwner;
        _usdtToken = IERC20(usdtTokenAddr);
        _oracle = IPriceOracle(oracleAddr);
        feeRecipient = _feeRecipient;
        feeRate = 0;
    }

    function owner() public view returns (address) {
        return _owner;
    }

    function _checkOwner() internal view override {
        require(msg.sender == _owner, "NotOwner");
    }

    function transferOwnership(address newOwner) external onlyAuthorized {
        require(newOwner != address(0), "Invalid owner");
        _owner = newOwner;
    }
}