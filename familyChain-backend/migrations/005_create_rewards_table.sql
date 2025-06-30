-- 创建奖品表
CREATE TABLE IF NOT EXISTS rewards (
    id SERIAL PRIMARY KEY,
    family_id INTEGER NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    image_url TEXT,
    token_price INTEGER NOT NULL,
    created_by INTEGER NOT NULL,
    active BOOLEAN DEFAULT TRUE,
    stock INTEGER DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (family_id) REFERENCES families(id) ON DELETE CASCADE,
    FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE CASCADE
);

-- 索引
CREATE INDEX idx_rewards_family_id ON rewards(family_id);
CREATE INDEX idx_rewards_active ON rewards(active); 