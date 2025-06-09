// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

/**
 * @title RewardToken
 * @dev ERC20 token that can be minted by contract owner and used for task rewards
 */
contract RewardToken is ERC20, Ownable {
    mapping(address => bool) public authorizedMinters;

    event MinterAdded(address indexed minter);
    event MinterRemoved(address indexed minter);
    event TokensMinted(address indexed to, uint256 amount);
    event TokensBurned(address indexed from, uint256 amount);

    /**
     * @dev Constructor that gives the msg.sender all of existing tokens.
     */
    constructor(string memory name, string memory symbol) ERC20(name, symbol) Ownable(msg.sender) {
        authorizedMinters[msg.sender] = true;
    }

    /**
     * @dev Modifier to check if the sender is authorized to mint tokens
     */
    modifier onlyMinter() {
        require(authorizedMinters[msg.sender], "Caller is not an authorized minter");
        _;
    }

    /**
     * @dev Adds a new authorized minter
     */
    function addMinter(address minter) public onlyOwner {
        require(minter != address(0), "Invalid minter address");
        authorizedMinters[minter] = true;
        emit MinterAdded(minter);
    }

    /**
     * @dev Removes an authorized minter
     */
    function removeMinter(address minter) public onlyOwner {
        require(authorizedMinters[minter], "Address is not a minter");
        authorizedMinters[minter] = false;
        emit MinterRemoved(minter);
    }

    /**
     * @dev Creates `amount` tokens and assigns them to `account`
     */
    function mint(address to, uint256 amount) public onlyMinter {
        require(to != address(0), "ERC20: mint to the zero address");
        _mint(to, amount);
        emit TokensMinted(to, amount);
    }

    /**
     * @dev Destroys `amount` tokens from `account`
     */
    function burn(address from, uint256 amount) public onlyMinter {
        require(from != address(0), "ERC20: burn from the zero address");
        require(balanceOf(from) >= amount, "ERC20: burn amount exceeds balance");
        _burn(from, amount);
        emit TokensBurned(from, amount);
    }
    
    /**
     * @dev Transfers rewards from contract owner to child
     */
    function transferReward(address to, uint256 amount) public onlyMinter {
        require(to != address(0), "ERC20: transfer to the zero address");
        _transfer(owner(), to, amount);
    }
} 