// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {ERC20Permit} from "@openzeppelin/contracts/token/ERC20/extensions/ERC20Permit.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

contract MockUSDT is ERC20, ERC20Permit, Ownable {
    constructor(string memory name_, string memory symbol_) ERC20(name_, symbol_) ERC20Permit(name_) Ownable(msg.sender) {}

    function decimals() public pure override returns (uint8) {
        return 6;
    }

    function mint(address to, uint256 amount) external onlyOwner {
        _mint(to, amount);
    }

    /// @notice 任意地址可领取任意数量（仅用于测试网/演示）
    function faucetMint(address to, uint256 amount) external {
        _mint(to, amount);
    }

    /// @notice 调用者给自己领币（仅用于测试网/演示）
    function faucetMint(uint256 amount) external {
        _mint(msg.sender, amount);
    }
}
