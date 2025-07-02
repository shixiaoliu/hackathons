// GetTask gets a task by ID
func (cm *ContractManager) GetTask(taskID uint64) (Task, error) {
	var task Task

	if cm.TaskRegistry == nil {
		return task, fmt.Errorf("task registry not initialized")
	}

	opts := &bind.CallOpts{Context: context.Background()}

	// 调用合约获取任务
	result, err := cm.TaskRegistry.GetTask(opts, big.NewInt(int64(taskID)))
	if err != nil {
		return task, fmt.Errorf("failed to get task: %v", err)
	}

	task = Task{
		ID:          result.Id.Uint64(),
		Creator:     result.Creator,
		AssignedTo:  result.AssignedTo,
		Title:       result.Title,
		Description: result.Description,
		Reward:      result.Reward,
		Completed:   result.Completed,
		Approved:    result.Approved,
	}

	return task, nil
} 