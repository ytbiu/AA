# AA 学习手册（基于本项目）

这套文档按“先理解，再动手，再排障”的顺序组织，建议按编号阅读。

## 阅读顺序

1. [01-账户抽象基础](/Users/yintao/devlop/solidity/AA/docs/learning/01-account-abstraction-foundations.md)
2. [02-EIP-7702 升级流程](/Users/yintao/devlop/solidity/AA/docs/learning/02-eip-7702-upgrade-flow.md)
3. [03-EIP-4337 UserOperation 流程](/Users/yintao/devlop/solidity/AA/docs/learning/03-eip-4337-userop-flow.md)
4. [04-项目实现拆解](/Users/yintao/devlop/solidity/AA/docs/learning/04-project-implementation-walkthrough.md)
5. [05-完整交易流程（本项目）](/Users/yintao/devlop/solidity/AA/docs/learning/05-end-to-end-transaction-flow.md)
6. [06-常见错误与排障](/Users/yintao/devlop/solidity/AA/docs/learning/06-common-errors-and-debugging.md)
7. [07-实验与进阶练习](/Users/yintao/devlop/solidity/AA/docs/learning/07-hands-on-experiments.md)
8. [08-自研与第三方边界](/Users/yintao/devlop/solidity/AA/docs/learning/08-what-we-build-vs-third-party.md)

## 图示

- 时序图: [tx-flow-sequence.mmd](/Users/yintao/devlop/solidity/AA/docs/diagrams/tx-flow-sequence.mmd)
- 架构图: [tx-flow-architecture.mmd](/Users/yintao/devlop/solidity/AA/docs/diagrams/tx-flow-architecture.mmd)

## 你在本项目里需要牢记的三个点

1. 7702 负责“把 EOA 变成可执行账户逻辑”，4337 负责“把账户操作标准化并交给 bundler 打包”。
2. Paymaster 用 BNB 先垫付，之后再在 `postOp` 阶段扣用户的 USDT。
3. 生产环境通常要求 bundler 路径可用，因此 paymaster 代码必须满足 bundler 的模拟规则（例如避免禁用 opcode）。
