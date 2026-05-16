// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "./base/PaymasterBase.sol";

contract USDTPaymaster is Initializable, PaymasterBase, UUPSUpgradeable {
    address private _owner;

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(
        address usdtTokenAddr,
        address oracleAddr,
        address _feeRecipient,
        address initialOwner
    ) public initializer {
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

    function _authorizeUpgrade(address) internal override onlyAuthorized {}
}