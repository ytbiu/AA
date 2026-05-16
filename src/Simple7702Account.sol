// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "./interfaces/IUSDTPaymaster.sol";

interface EIP1271 {
    function isValidSignature(bytes32 hash, bytes calldata signature) external view returns (bytes4);
}

contract Simple7702Account is Initializable, OwnableUpgradeable, UUPSUpgradeable, EIP1271 {
    uint256 private constant MAX_BATCH_SIZE = 5;

    event BatchExecuted(address caller, uint256 callCount);

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(address _owner) public initializer {
        __Ownable_init(_owner);
    }

    // EIP-7702 只绑定代码，不绑定存储，所以不再存储 paymaster 地址
    // executeBatch 现在不限制调用者，任何人都可以调用
    // 但 Paymaster 会验证签名，所以只有用户自己签名才能通过
    function executeBatch(IUSDTPaymaster.Call[] calldata calls) external {
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

    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}
}
