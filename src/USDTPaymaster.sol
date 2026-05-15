// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import "./interfaces/IUSDTPaymaster.sol";
import "./interfaces/IPriceOracle.sol";

contract USDTPaymaster is
    Initializable,
    OwnableUpgradeable,
    ReentrancyGuard,
    UUPSUpgradeable,
    IUSDTPaymaster
{
    using ECDSA for bytes32;
    using MessageHashUtils for bytes32;

    IERC20 private _usdtToken;
    IPriceOracle private _oracle;
    uint256 public feeRate; // 百分比，10000 = 100% (100 = 1%)
    address public feeRecipient;
    mapping(address => bool) private _relayers;

    event BatchExecuted(address indexed user, uint256 gasUsed, uint256 compensation);
    event RelayerAdded(address indexed relayer);
    event RelayerRemoved(address indexed relayer);
    event FeeRateUpdated(uint256 rate);
    event FeeRecipientUpdated(address recipient);
    event OracleUpdated(address oracle);

    error NotRelayer();
    error InvalidSignature();
    error CallFailed();
    error TransferFailed();

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(
        address usdtTokenAddr,
        address oracleAddr,
        address _feeRecipient,
        address _owner
    ) public initializer {
        __Ownable_init(_owner);
        // Note: UUPSUpgradeable doesn't require explicit initialization

        _usdtToken = IERC20(usdtTokenAddr);
        _oracle = IPriceOracle(oracleAddr);
        feeRecipient = _feeRecipient;
        feeRate = 0;
    }

    function usdtToken() external view returns (address) {
        return address(_usdtToken);
    }

    function oracle() external view returns (address) {
        return address(_oracle);
    }

    modifier onlyRelayer() {
        if (!_relayers[msg.sender]) revert NotRelayer();
        _;
    }

    function executeBatch(UserOperation calldata userOp, bytes calldata signature)
        external
        onlyRelayer
        nonReentrant
    {
        // 1. 验证签名
        bytes32 hash = keccak256(abi.encode(userOp));
        bytes32 ethSignedHash = hash.toEthSignedMessageHash();
        address signer = ethSignedHash.recover(signature);
        if (signer != userOp.user) revert InvalidSignature();

        // 2. 执行 batch
        uint256 gasBefore = gasleft();
        for (uint256 i = 0; i < userOp.calls.length; i++) {
            (bool success, ) = userOp.calls[i].to.call(userOp.calls[i].data);
            if (!success) revert CallFailed();
        }
        uint256 gasUsed = gasBefore - gasleft();

        // 3. 计算补偿
        uint256 bnbCost = gasUsed * tx.gasprice;
        uint256 bnbPriceInUsdt = _oracle.getBNBPriceInUSDT();
        uint256 compensation = (bnbCost * bnbPriceInUsdt) / 1e18;

        // 4. 加上手续费
        uint256 fee = (compensation * feeRate) / 10000;
        uint256 totalCompensation = compensation + fee;

        // 5. 从用户 USDT 转给 Relayer 和手续费归属
        if (!_usdtToken.transferFrom(userOp.user, msg.sender, compensation))
            revert TransferFailed();
        if (fee > 0) {
            if (!_usdtToken.transferFrom(userOp.user, feeRecipient, fee))
                revert TransferFailed();
        }

        emit BatchExecuted(userOp.user, gasUsed, totalCompensation);
    }

    function addRelayer(address relayer) external onlyOwner {
        _relayers[relayer] = true;
        emit RelayerAdded(relayer);
    }

    function removeRelayer(address relayer) external onlyOwner {
        _relayers[relayer] = false;
        emit RelayerRemoved(relayer);
    }

    function setFeeRate(uint256 rate) external onlyOwner {
        feeRate = rate;
        emit FeeRateUpdated(rate);
    }

    function setFeeRecipient(address recipient) external onlyOwner {
        feeRecipient = recipient;
        emit FeeRecipientUpdated(recipient);
    }

    function setOracle(address _oracleAddr) external onlyOwner {
        _oracle = IPriceOracle(_oracleAddr);
        emit OracleUpdated(_oracleAddr);
    }

    function isRelayer(address relayer) external view returns (bool) {
        return _relayers[relayer];
    }

    function _authorizeUpgrade(address newImplementation)
        internal
        override
        onlyOwner
    {}
}