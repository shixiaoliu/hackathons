-- 向奖品表添加区块链奖品ID字段
ALTER TABLE rewards ADD COLUMN IF NOT EXISTS contract_reward_id INTEGER;

-- 创建索引以加快查询速度
CREATE INDEX IF NOT EXISTS idx_rewards_contract_reward_id ON rewards(contract_reward_id);

-- 注释: 这个字段用于存储区块链上的奖品ID，这样应用程序就可以在调用合约时使用正确的ID 