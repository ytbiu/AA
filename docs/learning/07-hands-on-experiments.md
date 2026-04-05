# 07. 实验与进阶练习

## 目标

通过“改参数 + 观察结果”的方式，把 7702/4337/Paymaster 的行为吃透。

## 实验 1：观察 7702 升级前后 code 变化

1. 升级前执行 `cast code <addr>`，记录结果。
2. 触发 `/api/v1/7702/upgrade`。
3. 升级后再次读取 code，对比 delegation 前缀。

你会得到：7702 不改地址，但改变地址的执行行为。

## 实验 2：故意让报价过低，复现 AA33

1. 临时降低后端 `tokenPerNative` 或减少安全系数。
2. 发起 UserOp。
3. 观察 bundler 的模拟报错。

你会理解：AA33 常是“验证阶段预算或校验失败”。

## 实验 3：复现并理解 AA50

1. 用户 USDT 留很少余额。
2. 发起转账 + 代付。
3. 观察 `postOp` 扣款失败。

你会理解：validate 过了不代表 settlement 一定成功。

## 实验 4：调 verificationGasLimit，观察效率阈值

1. 把 `verificationGasLimit` 设很大。
2. 观察 `Verification gas limit efficiency too low`。
3. 回调到合理区间，比较成功率。

你会理解：bundler 不只看能不能跑，还看“参数是否合理”。

## 实验 5：一次 approve，多次交易（进阶课题）

当前实现在每笔 UserOp 中都执行 `approve + transfer`。

进阶可尝试：

1. 首次单独做较大额度 approve，后续 UserOp 只做 transfer。
2. 对比“每笔都 approve”与“复用 allowance”的 gas 与体验差异。
3. 评估安全边界（额度控制、撤销策略、地址白名单）。

## 实验记录模板

建议每次实验记录：

1. 输入参数（gas、amount、allowance、nonce）。
2. 结果（成功/失败 hash）。
3. 关键日志或错误码。
4. 结论与下一步调整。

持续 5~10 次迭代后，你会对 AA 的工程调参非常熟悉。
