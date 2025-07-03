const { ethers } = require("hardhat");
const TaskRegistryABI = require("./abi/TaskRegistry.json");

async function testContract() {
    try {
        // 连接到 Sepolia 网络
        const provider = new ethers.providers.JsonRpcProvider(process.env.BLOCKCHAIN_RPC_URL);
        const wallet = new ethers.Wallet(process.env.BLOCKCHAIN_PRIVATE_KEY_PARENT, provider);
        
        const contractAddress = "0x11dB634CFD2f58967e472a179ebDbaF8AB067144";
        
        console.log("测试合约地址:", contractAddress);
        console.log("使用钱包地址:", wallet.address);
        
        // 创建合约实例
        const contract = new ethers.Contract(contractAddress, TaskRegistryABI, wallet);
        
        // 1. 测试只读方法
        console.log("\n=== 测试只读方法 ===");
        
        try {
            const owner = await contract.owner();
            console.log("合约所有者:", owner);
        } catch (error) {
            console.error("获取所有者失败:", error.message);
        }
        
        try {
            const taskCount = await contract.taskCount();
            console.log("当前任务数量:", taskCount.toString());
        } catch (error) {
            console.error("获取任务数量失败:", error.message);
        }
        
        // 2. 检查账户余额
        console.log("\n=== 检查账户状态 ===");
        const balance = await provider.getBalance(wallet.address);
        console.log("账户余额:", ethers.utils.formatEther(balance), "ETH");
        
        // 3. 测试创建任务（小金额）
        console.log("\n=== 测试创建任务 ===");
        const testReward = ethers.utils.parseEther("0.001"); // 0.001 ETH
        
        console.log("准备创建任务，奖励金额:", ethers.utils.formatEther(testReward), "ETH");
        
        // 估算 gas
        try {
            const gasEstimate = await contract.estimateGas.createTask(
                "测试任务",
                "这是一个测试任务",
                testReward,
                { value: testReward }
            );
            console.log("预估 gas:", gasEstimate.toString());
        } catch (error) {
            console.error("Gas 估算失败:", error.message);
            console.error("错误详情:", error);
            return;
        }
        
        // 实际创建任务
        try {
            const tx = await contract.createTask(
                "测试任务",
                "这是一个测试任务",
                testReward,
                { 
                    value: testReward,
                    gasLimit: 500000
                }
            );
            
            console.log("交易已发送:", tx.hash);
            const receipt = await tx.wait();
            console.log("交易已确认，区块号:", receipt.blockNumber);
            
            // 解析事件
            const events = receipt.logs.map(log => {
                try {
                    return contract.interface.parseLog(log);
                } catch {
                    return null;
                }
            }).filter(event => event !== null);
            
            console.log("事件:", events);
            
        } catch (error) {
            console.error("创建任务失败:", error.message);
            console.error("错误代码:", error.code);
            console.error("错误详情:", error);
        }
        
    } catch (error) {
        console.error("测试失败:", error);
    }
}

// 运行测试
testContract();