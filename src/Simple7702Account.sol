// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "./interfaces/IUSDTPaymaster.sol";

interface EIP1271 {
    function isValidSignature(bytes32 hash, bytes calldata signature) external view returns (bytes4);
}

contract Simple7702Account is Initializable, OwnableUpgradeable, UUPSUpgradeable, EIP1271 {
    IUSDTPaymaster public paymaster;
    uint256 private constant MAX_BATCH_SIZE = 5;

    event BatchExecuted(address caller, uint256 callCount);

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(address _paymaster, address _owner) public initializer {
        __Ownable_init(_owner);
        paymaster = IUSDTPaymaster(_paymaster);
    }

    modifier onlyPaymaster() {
        require(msg.sender == address(paymaster), "Simple7702Account: only paymaster");
        _;
    }

    function executeBatch(IUSDTPaymaster.Call[] calldata calls) external onlyPaymaster {
        require(calls.length <= MAX_BATCH_SIZE, "Simple7702Account: batch too large");

        for (uint256 i = 0; i < calls.length; i++) {
            (bool success, ) = calls[i].to.call(calls[i].data);
            require(success, "Simple7702Account: call failed");
        }

        emit BatchExecuted(msg.sender, calls.length);
    }

    function isValidSignature(bytes32 hash, bytes calldata signature) external view returns (bytes4) {
        // 使用 ECDSA 验证签名
        address signer = recoverSigner(hash, signature);
        if (signer == owner()) {
            return 0x1626ba7e; // EIP-1271 magic value
        }
        return 0xffffffff; // 无效
    }

    function recoverSigner(bytes32 ethSignedHash, bytes calldata signature) internal pure returns (address) {
        bytes32 r;
        bytes32 s;
        uint8 v;

        if (signature.length == 65) {
            assembly {
                r := calldataload(signature.offset)
                s := calldataload(add(signature.offset, 32))
                v := byte(0, calldataload(add(signature.offset, 64)))
            }
            return ecrecover(ethSignedHash, v, r, s);
        }
        return address(0);
    }

    function setPaymaster(address _paymaster) external onlyOwner {
        paymaster = IUSDTPaymaster(_paymaster);
    }

    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}
}
