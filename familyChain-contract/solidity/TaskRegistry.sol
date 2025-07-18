// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @title TaskRegistry
 * @dev Contract for managing child tasks and rewards
 */
contract TaskRegistry {
    struct Task {
        uint256 id;
        address creator;
        address assignedTo;
        string title;
        string description;
        uint256 reward;
        bool completed;
        bool approved;
        bool rejected;
    }

    mapping(uint256 => Task) public tasks;
    uint256 public taskCount = 0;
    
    address public owner;

    constructor() {
        owner = msg.sender;
    }

    // Events
    event TaskCreated(uint256 indexed taskId, address indexed creator, string title, uint256 reward);
    event TaskAssigned(uint256 indexed taskId, address indexed assignedTo);
    event TaskCompleted(uint256 indexed taskId, address indexed completedBy);
    event TaskApproved(uint256 indexed taskId, address indexed approvedBy);
    event RewardTransferred(uint256 indexed taskId, address indexed recipient, uint256 amount);
    event TaskRejected(uint256 indexed taskId, address indexed rejectedBy);

    /**
     * @dev Creates a new task
     */
    function createTask(string memory title, string memory description, uint256 reward) public payable returns (uint256) {
        require(msg.value == reward, "msg.value must equal reward");
        taskCount++;
        tasks[taskCount] = Task(
            taskCount,
            msg.sender,
            address(0),
            title,
            description,
            reward,
            false,
            false,
            false
        );
        
        emit TaskCreated(taskCount, msg.sender, title, reward);
        return taskCount;
    }

    /**
     * @dev Assigns a task to a child's address
     */
    function assignTask(uint256 taskId, address childAddress) public {
        require(tasks[taskId].creator == msg.sender, "Only task creator can assign the task");
        require(tasks[taskId].assignedTo == address(0), "Task already assigned");
        
        tasks[taskId].assignedTo = childAddress;
        
        emit TaskAssigned(taskId, childAddress);
    }

    /**
     * @dev Marks a task as completed by the assigned child
     */
    function completeTask(uint256 taskId) public {
        require(tasks[taskId].assignedTo == msg.sender, "Only assigned child can complete the task");
        require(!tasks[taskId].completed, "Task already completed");
        
        tasks[taskId].completed = true;
        
        emit TaskCompleted(taskId, tasks[taskId].assignedTo);
    }

    /**
     * @dev Approves a completed task and transfers the reward
     */
    function approveTask(uint256 taskId) public {
        require(tasks[taskId].creator == msg.sender, "Only task creator can approve the task");
        require(tasks[taskId].completed, "Task not completed yet");
        require(!tasks[taskId].approved, "Task already approved");
        
        tasks[taskId].approved = true;
        
        // Log the assigned address and reward amount
        address assignedAddress = tasks[taskId].assignedTo;
        uint256 rewardAmount = tasks[taskId].reward;
        
        // For debugging purposes
        emit TaskApproved(taskId, msg.sender);
        emit RewardTransferred(taskId, assignedAddress, rewardAmount);
        
        // Transfer reward to child
        address payable recipient = payable(assignedAddress);
        recipient.transfer(rewardAmount);
    }

    /**
     * @dev Rejects a completed task and refunds the reward to the creator
     */
    function rejectTask(uint256 taskId) public {
        require(tasks[taskId].creator == msg.sender, "Only task creator can reject the task");
        require(tasks[taskId].completed, "Task not completed yet");
        require(!tasks[taskId].approved, "Task already approved");
        require(!tasks[taskId].rejected, "Task already rejected");
        
        tasks[taskId].rejected = true;
        
        // Refund reward to creator
        address payable creatorAddress = payable(tasks[taskId].creator);
        creatorAddress.transfer(tasks[taskId].reward);
        
        emit TaskRejected(taskId, msg.sender);
        emit RewardTransferred(taskId, creatorAddress, tasks[taskId].reward);
    }

    /**
     * @dev Retrieves a task by ID
     */
    function getTask(uint256 taskId) public view returns (
        uint256, address, address, string memory, string memory, uint256, bool, bool, bool
    ) {
        Task memory task = tasks[taskId];
        return (
            task.id,
            task.creator,
            task.assignedTo,
            task.title,
            task.description,
            task.reward,
            task.completed,
            task.approved,
            task.rejected
        );
    }

    function withdraw() public {
        require(msg.sender == owner, "Only owner can withdraw");
        payable(owner).transfer(address(this).balance);
    }

    receive() external payable {}
} 