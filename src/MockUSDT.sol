// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract MockUSDT is ERC20 {
    uint256 public constant FAUCET_AMOUNT = 100 * 10 ** 18; // 100 USDT
    uint8 private constant DECIMALS = 18;

    constructor() ERC20("Mock USDT", "USDT") {
        _mint(msg.sender, 1000000 * 10 ** 18); // 初始 1M 给部署者
    }

    function decimals() public pure override returns (uint8) {
        return DECIMALS;
    }

    function faucet() external {
        _mint(msg.sender, FAUCET_AMOUNT);
    }

    function mint(address to, uint256 amount) external {
        _mint(to, amount);
    }
}
