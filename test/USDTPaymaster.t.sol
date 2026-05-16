// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "../src/MockUSDT.sol";
import "../src/PriceOracle.sol";
import "../src/USDTPaymaster.sol";
import "../src/base/PaymasterBase.sol";

contract MockPancakeRouter is IPancakeRouter {
    address public wbnb;
    mapping(address => mapping(address => uint256)) private _prices; // from -> to -> price ratio

    constructor(address _wbnb) {
        wbnb = _wbnb;
    }

    function setPrice(address tokenA, address tokenB, uint256 priceA, uint256 priceB) external {
        // This sets the conversion rate: 1 tokenA = (priceB/priceA) tokenB
        // For example: if 1 BNB = 600 USDT, then priceBNB = 1e18, priceUSDT = 600e18
        _prices[tokenA][tokenB] = (priceB * 1e18) / priceA;
        _prices[tokenB][tokenA] = (priceA * 1e18) / priceB;
    }

    function getAmountsOut(uint256 amountIn, address[] calldata path) external view returns (uint256[] memory amounts) {
        amounts = new uint256[](path.length);
        amounts[0] = amountIn;

        for (uint256 i = 0; i < path.length - 1; i++) {
            address tokenIn = path[i];
            address tokenOut = path[i + 1];
            amounts[i + 1] = (amounts[i] * _prices[tokenIn][tokenOut]) / 1e18;
        }
    }

    function WBNB() external view returns (address) {
        return wbnb;
    }
}

contract USDTPaymasterTest is Test {
    USDTPaymaster paymaster;
    USDTPaymaster implementation;
    MockUSDT usdt;
    PriceOracle oracle;
    MockPancakeRouter mockRouter;

    address deployer;
    address relayer;
    address user;
    address recipient;

    function setUp() public {
        deployer = address(this);
        relayer = address(0x1);
        user = address(0x2);
        recipient = address(0x3);

        // Deploy MockUSDT
        usdt = new MockUSDT();
        usdt.mint(user, 1000 * 10 ** 18);

        // Deploy MockPancakeRouter
        address wbnbAddress = address(0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c); // BSC WBNB address
        mockRouter = new MockPancakeRouter(wbnbAddress);

        // Set price: 1 BNB = 600 USDT (approximately)
        mockRouter.setPrice(wbnbAddress, address(usdt), 1 ether, 600 * 1 ether);

        // Deploy PriceOracle
        oracle = new PriceOracle(address(mockRouter), wbnbAddress, address(usdt));

        // Deploy USDTPaymaster Implementation
        implementation = new USDTPaymaster();

        // Deploy UUPS Proxy
        ERC1967Proxy proxy = new ERC1967Proxy(
            address(implementation),
            abi.encodeCall(
                USDTPaymaster.initialize,
                (address(usdt), address(oracle), recipient, deployer)
            )
        );
        paymaster = USDTPaymaster(address(proxy));

        // Add Relayer
        paymaster.addRelayer(relayer);
    }

    function test_InitializeCorrectly() public view {
        assertEq(address(paymaster.usdtToken()), address(usdt));
        assertEq(address(paymaster.oracle()), address(oracle));
        assertEq(paymaster.feeRecipient(), recipient);
        assertEq(paymaster.owner(), deployer);
        assertEq(paymaster.feeRate(), 0);
    }

    function test_AddRelayerCorrectly() public view {
        assertTrue(paymaster.isRelayer(relayer));
    }

    function test_RemoveRelayerCorrectly() public {
        paymaster.removeRelayer(relayer);
        assertFalse(paymaster.isRelayer(relayer));
    }

    function test_NonOwnerCannotAddRelayer() public {
        vm.prank(user);
        vm.expectRevert();
        paymaster.addRelayer(user);
    }

    function test_SetFeeRateCorrectly() public {
        paymaster.setFeeRate(100); // 1%
        assertEq(paymaster.feeRate(), 100);
    }

    function test_SetFeeRecipientCorrectly() public {
        paymaster.setFeeRecipient(user);
        assertEq(paymaster.feeRecipient(), user);
    }

    function test_SetOracleCorrectly() public {
        paymaster.setOracle(user);
        assertEq(address(paymaster.oracle()), user);
    }

    function test_NonRelayerCannotExecuteBatch() public {
        IUSDTPaymaster.Call[] memory calls = new IUSDTPaymaster.Call[](1);
        calls[0] = IUSDTPaymaster.Call(address(usdt), abi.encodeWithSignature("faucet()"));

        IUSDTPaymaster.UserOperation memory userOp = IUSDTPaymaster.UserOperation(user, calls);
        bytes memory signature = new bytes(65);

        vm.prank(user);
        vm.expectRevert(PaymasterBase.NotRelayer.selector);
        paymaster.executeBatch(userOp, signature);
    }

    function test_UUPSUpgradeability() public view {
        // 验证 proxy 地址不为零
        assertNotEq(address(paymaster), address(0));
        // 验证 proxy 与 implementation 不同
        assertNotEq(address(paymaster), address(implementation));
    }

    function test_ImplementationDisabled() public {
        // 验证 implementation 已禁用初始化
        // 尝试调用 initialize 应该失败
        vm.expectRevert();
        implementation.initialize(
            address(usdt),
            address(oracle),
            recipient,
            deployer
        );
    }
}