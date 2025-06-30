-- 创建兑换记录表
CREATE TABLE IF NOT EXISTS exchanges (
    id SERIAL PRIMARY KEY,
    reward_id INTEGER NOT NULL,
    child_id INTEGER NOT NULL,
    token_amount INTEGER NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'pending', -- pending, completed, cancelled
    exchange_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    completed_date TIMESTAMP,
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (reward_id) REFERENCES rewards(id) ON DELETE CASCADE,
    FOREIGN KEY (child_id) REFERENCES children(id) ON DELETE CASCADE
);

-- 索引
CREATE INDEX idx_exchanges_child_id ON exchanges(child_id);
CREATE INDEX idx_exchanges_reward_id ON exchanges(reward_id);
CREATE INDEX idx_exchanges_status ON exchanges(status); 