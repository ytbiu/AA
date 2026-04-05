// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Test} from "forge-std/Test.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {MockUSDT} from "../src/MockUSDT.sol";
import {EIP7702DelegateAccount} from "../src/EIP7702DelegateAccount.sol";
import {USDTTokenPaymaster} from "../src/USDTTokenPaymaster.sol";
import {IEntryPoint} from "../src/interfaces/IEntryPoint.sol";
import {UserOperation} from "../src/interfaces/UserOperation.sol";
import {SimpleEntryPointMock} from "../src/mocks/SimpleEntryPointMock.sol";

contract USDTTokenPaymasterTest is Test {
    uint256 internal userPk = 0xA11CE;
    uint256 internal quoteSignerPk = 0xB0B;

    address internal user;
    address internal receiver;
    address internal quoteSigner;

    SimpleEntryPointMock internal entryPoint;
    MockUSDT internal usdt;
    EIP7702DelegateAccount internal accountLogic;
    USDTTokenPaymaster internal paymaster;

    function setUp() public {
        user = vm.addr(userPk);
        quoteSigner = vm.addr(quoteSignerPk);
        receiver = makeAddr("receiver");

        entryPoint = new SimpleEntryPointMock();
        usdt = new MockUSDT("Test USDT", "tUSDT");
        accountLogic = new EIP7702DelegateAccount(IEntryPoint(address(entryPoint)));

        // 模拟 EIP-7702：把 EOA 的代码替换为 delegation target 的运行时代码。
        vm.etch(user, address(accountLogic).code);

        paymaster = new USDTTokenPaymaster(
            IEntryPoint(address(entryPoint)),
            quoteSigner,
            500
        );
        paymaster.setTokenConfig(address(usdt), true, 600e6);

        usdt.mint(user, 10_000e6);
    }

    function testHandleUserOpWithApproveAndUSDTSettlement() public {
        uint256 transferAmount = 100e6;
        uint48 validUntil = uint48(block.timestamp + 5 minutes);

        bytes memory approveCall = abi.encodeWithSelector(IERC20.approve.selector, address(paymaster), type(uint256).max);
        bytes memory transferCall = abi.encodeWithSelector(IERC20.transfer.selector, receiver, transferAmount);
        address[] memory targets = new address[](2);
        targets[0] = address(usdt);
        targets[1] = address(usdt);

        uint256[] memory values = new uint256[](2);

        bytes[] memory batchData = new bytes[](2);
        batchData[0] = approveCall;
        batchData[1] = transferCall;

        bytes memory accountCall = abi.encodeWithSelector(
            EIP7702DelegateAccount.executeBatch.selector,
            targets,
            values,
            batchData
        );

        UserOperation memory op = UserOperation({
            sender: user,
            nonce: 0,
            initCode: bytes(""),
            callData: accountCall,
            callGasLimit: 300_000,
            verificationGasLimit: 250_000,
            preVerificationGas: 50_000,
            maxFeePerGas: 1 gwei,
            maxPriorityFeePerGas: 1 gwei,
            paymasterAndData: bytes(""),
            signature: bytes("")
        });

        uint256 maxCost = (op.callGasLimit + op.verificationGasLimit + op.preVerificationGas) * op.maxFeePerGas;
        uint256 tokenAmount = (maxCost * 600e6) / 1e18;

        bytes memory paymasterQuoteSignature = _signPaymasterQuote(
            quoteSignerPk,
            op,
            tokenAmount,
            validUntil,
            address(paymaster)
        );

        op.paymasterAndData = abi.encodePacked(address(paymaster), address(usdt), validUntil, tokenAmount, paymasterQuoteSignature);

        bytes32 userOpHash = entryPoint.getUserOpHash(op);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(userPk, _toEthSignedMessageHash(userOpHash));
        op.signature = abi.encodePacked(r, s, v);

        UserOperation[] memory ops = new UserOperation[](1);
        ops[0] = op;

        entryPoint.handleOps(ops, payable(address(0xBEEF)));

        assertEq(usdt.balanceOf(receiver), transferAmount, "receiver should get transfer");
        assertGt(usdt.balanceOf(address(paymaster)), 0, "paymaster should receive USDT settlement");
    }

    function _signPaymasterQuote(
        uint256 signerPk,
        UserOperation memory op,
        uint256 tokenAmount,
        uint48 validUntil,
        address paymasterAddr
    ) internal view returns (bytes memory) {
        bytes32 quoteHash = keccak256(
            abi.encode(
                block.chainid,
                paymasterAddr,
                op.sender,
                op.nonce,
                keccak256(op.callData),
                op.callGasLimit,
                op.verificationGasLimit,
                op.preVerificationGas,
                op.maxFeePerGas,
                op.maxPriorityFeePerGas,
                address(usdt),
                tokenAmount,
                validUntil
            )
        );

        (uint8 v, bytes32 r, bytes32 s) = vm.sign(signerPk, _toEthSignedMessageHash(quoteHash));
        return abi.encodePacked(r, s, v);
    }

    function _toEthSignedMessageHash(bytes32 hash) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", hash));
    }
}
