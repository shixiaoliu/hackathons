// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./RewardToken.sol";

/**
 * @title RewardRegistry
 * @dev Contract for managing physical rewards and token exchanges
 */
contract RewardRegistry is Ownable {
    // 奖品结构
    struct Reward {
        uint256 id;
        address creator;      // 创建者地址（家长）
        uint256 familyId;     // 所属家庭ID
        string name;          // 奖品名称
        string description;   // 奖品描述
        string imageURI;      // 奖品图片URI
        uint256 tokenPrice;   // 代币价格
        uint256 stock;        // 库存数量
        bool active;          // 是否激活
    }

    // 兑换记录结构
    struct Exchange {
        uint256 id;
        uint256 rewardId;     // 奖品ID
        address child;        // 兑换的孩子地址
        uint256 tokenAmount;  // 代币数量
        uint256 exchangeDate; // 兑换日期
        bool fulfilled;       // 是否已发放
    }

    // 合约状态变量
    mapping(uint256 => Reward) public rewards;
    mapping(uint256 => Exchange) public exchanges;
    uint256 public rewardCount = 0;
    uint256 public exchangeCount = 0;
    
    // 家庭对应的奖品ID列表
    mapping(uint256 => uint256[]) public familyRewards;
    // 孩子对应的兑换记录ID列表
    mapping(address => uint256[]) public childExchanges;
    
    // 代币合约地址
    RewardToken public tokenContract;

    // 事件
    event RewardCreated(uint256 indexed rewardId, address indexed creator, uint256 indexed familyId, string name, uint256 tokenPrice);
    event RewardUpdated(uint256 indexed rewardId, string name, uint256 tokenPrice, bool active);
    event RewardExchanged(uint256 indexed exchangeId, uint256 indexed rewardId, address indexed child, uint256 tokenAmount);
    event ExchangeFulfilled(uint256 indexed exchangeId, address indexed parent);
    
    /**
     * @dev 构造函数，设置代币合约地址
     * @param _tokenAddress 代币合约地址
     */
    constructor(address _tokenAddress) Ownable(msg.sender) {
        tokenContract = RewardToken(_tokenAddress);
    }
    
    /**
     * @dev 创建新的实物奖励
     * @param _familyId 家庭ID
     * @param _name 奖品名称
     * @param _description 奖品描述
     * @param _imageURI 奖品图片URI
     * @param _tokenPrice 代币价格
     * @param _stock 库存数量
     */
    function createReward(
        uint256 _familyId, 
        string memory _name, 
        string memory _description, 
        string memory _imageURI, 
        uint256 _tokenPrice, 
        uint256 _stock
    ) public returns (uint256) {
        require(_tokenPrice > 0, "Token price must be greater than zero");
        require(_stock > 0, "Stock must be greater than zero");
        
        rewardCount++;
        
        rewards[rewardCount] = Reward(
            rewardCount,
            msg.sender,
            _familyId,
            _name,
            _description,
            _imageURI,
            _tokenPrice,
            _stock,
            true
        );
        
        // 添加到家庭奖品列表
        familyRewards[_familyId].push(rewardCount);
        
        emit RewardCreated(rewardCount, msg.sender, _familyId, _name, _tokenPrice);
        return rewardCount;
    }
    
    /**
     * @dev 更新奖品信息
     * @param _rewardId 奖品ID
     * @param _name 奖品名称
     * @param _description 奖品描述
     * @param _imageURI 奖品图片URI
     * @param _tokenPrice 代币价格
     * @param _stock 库存数量
     * @param _active 是否激活
     */
    function updateReward(
        uint256 _rewardId,
        string memory _name,
        string memory _description,
        string memory _imageURI,
        uint256 _tokenPrice,
        uint256 _stock,
        bool _active
    ) public {
        require(rewards[_rewardId].creator == msg.sender, "Only creator can update the reward");
        require(_tokenPrice > 0, "Token price must be greater than zero");
        
        Reward storage reward = rewards[_rewardId];
        reward.name = _name;
        reward.description = _description;
        reward.imageURI = _imageURI;
        reward.tokenPrice = _tokenPrice;
        reward.stock = _stock;
        reward.active = _active;
        
        emit RewardUpdated(_rewardId, _name, _tokenPrice, _active);
    }
    
    /**
     * @dev 兑换奖品
     * @param _rewardId 奖品ID
     */
    function exchangeReward(uint256 _rewardId) public returns (uint256) {
        Reward storage reward = rewards[_rewardId];
        
        require(reward.active, "Reward is not active");
        require(reward.stock > 0, "Reward out of stock");
        
        uint256 tokenPrice = reward.tokenPrice;
        
        // 检查代币余额
        require(tokenContract.balanceOf(msg.sender) >= tokenPrice, "Insufficient token balance");
        
        // 扣除代币
        tokenContract.burn(msg.sender, tokenPrice);
        
        // 减少库存
        reward.stock--;
        
        // 创建兑换记录
        exchangeCount++;
        exchanges[exchangeCount] = Exchange(
            exchangeCount,
            _rewardId,
            msg.sender,
            tokenPrice,
            block.timestamp,
            false
        );
        
        // 添加到孩子的兑换记录
        childExchanges[msg.sender].push(exchangeCount);
        
        emit RewardExchanged(exchangeCount, _rewardId, msg.sender, tokenPrice);
        return exchangeCount;
    }
    
    /**
     * @dev 标记兑换记录为已发放
     * @param _exchangeId 兑换记录ID
     */
    function fulfillExchange(uint256 _exchangeId) public {
        Exchange storage exchange = exchanges[_exchangeId];
        Reward storage reward = rewards[exchange.rewardId];
        
        require(reward.creator == msg.sender, "Only reward creator can fulfill the exchange");
        require(!exchange.fulfilled, "Exchange already fulfilled");
        
        exchange.fulfilled = true;
        
        emit ExchangeFulfilled(_exchangeId, msg.sender);
    }
    
    /**
     * @dev 获取家庭的奖品数量
     * @param _familyId 家庭ID
     */
    function getFamilyRewardCount(uint256 _familyId) public view returns (uint256) {
        return familyRewards[_familyId].length;
    }
    
    /**
     * @dev 获取家庭的奖品ID
     * @param _familyId 家庭ID
     * @param _index 索引
     */
    function getFamilyRewardId(uint256 _familyId, uint256 _index) public view returns (uint256) {
        require(_index < familyRewards[_familyId].length, "Index out of bounds");
        return familyRewards[_familyId][_index];
    }
    
    /**
     * @dev 获取孩子的兑换记录数量
     * @param _child 孩子地址
     */
    function getChildExchangeCount(address _child) public view returns (uint256) {
        return childExchanges[_child].length;
    }
    
    /**
     * @dev 获取孩子的兑换记录ID
     * @param _child 孩子地址
     * @param _index 索引
     */
    function getChildExchangeId(address _child, uint256 _index) public view returns (uint256) {
        require(_index < childExchanges[_child].length, "Index out of bounds");
        return childExchanges[_child][_index];
    }
    
    /**
     * @dev 获取奖品信息
     * @param _rewardId 奖品ID
     */
    function getReward(uint256 _rewardId) public view returns (
        uint256, address, uint256, string memory, string memory, string memory, uint256, uint256, bool
    ) {
        Reward memory reward = rewards[_rewardId];
        return (
            reward.id,
            reward.creator,
            reward.familyId,
            reward.name,
            reward.description,
            reward.imageURI,
            reward.tokenPrice,
            reward.stock,
            reward.active
        );
    }
    
    /**
     * @dev 获取兑换记录信息
     * @param _exchangeId 兑换记录ID
     */
    function getExchange(uint256 _exchangeId) public view returns (
        uint256, uint256, address, uint256, uint256, bool
    ) {
        Exchange memory exchange = exchanges[_exchangeId];
        return (
            exchange.id,
            exchange.rewardId,
            exchange.child,
            exchange.tokenAmount,
            exchange.exchangeDate,
            exchange.fulfilled
        );
    }
} 