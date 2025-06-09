// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @title FamilyRegistry
 * @dev Contract for managing family relationships and child accounts
 */
contract FamilyRegistry {
    struct Child {
        address walletAddress;
        string name;
        uint8 age;
        bool active;
    }
    
    struct Family {
        uint256 id;
        address parent;
        string name;
        bool active;
        mapping(address => Child) children;
        address[] childAddresses;
    }

    // Mappings
    mapping(uint256 => Family) public families;
    mapping(address => uint256[]) public parentFamilies;
    mapping(address => uint256) public childToFamily;
    
    uint256 public familyCount = 0;
    
    // Events
    event FamilyCreated(uint256 indexed familyId, address indexed parent, string name);
    event ChildAdded(uint256 indexed familyId, address indexed childAddress, string name, uint8 age);
    event ChildRemoved(uint256 indexed familyId, address indexed childAddress);
    event FamilyUpdated(uint256 indexed familyId, string name);

    /**
     * @dev Creates a new family
     */
    function createFamily(string memory name) public returns (uint256) {
        familyCount++;
        
        Family storage newFamily = families[familyCount];
        newFamily.id = familyCount;
        newFamily.parent = msg.sender;
        newFamily.name = name;
        newFamily.active = true;
        
        parentFamilies[msg.sender].push(familyCount);
        
        emit FamilyCreated(familyCount, msg.sender, name);
        return familyCount;
    }
    
    /**
     * @dev Adds a child to a family
     */
    function addChild(uint256 familyId, address childAddress, string memory name, uint8 age) public {
        require(families[familyId].parent == msg.sender, "Only parent can add children");
        require(families[familyId].active, "Family is not active");
        require(childToFamily[childAddress] == 0, "Child already registered in a family");
        
        Child memory newChild = Child(childAddress, name, age, true);
        families[familyId].children[childAddress] = newChild;
        families[familyId].childAddresses.push(childAddress);
        
        childToFamily[childAddress] = familyId;
        
        emit ChildAdded(familyId, childAddress, name, age);
    }
    
    /**
     * @dev Removes a child from a family
     */
    function removeChild(uint256 familyId, address childAddress) public {
        require(families[familyId].parent == msg.sender, "Only parent can remove children");
        require(families[familyId].active, "Family is not active");
        require(families[familyId].children[childAddress].active, "Child is not active in this family");
        
        families[familyId].children[childAddress].active = false;
        childToFamily[childAddress] = 0;
        
        emit ChildRemoved(familyId, childAddress);
    }
    
    /**
     * @dev Updates family information
     */
    function updateFamily(uint256 familyId, string memory name) public {
        require(families[familyId].parent == msg.sender, "Only parent can update family");
        require(families[familyId].active, "Family is not active");
        
        families[familyId].name = name;
        
        emit FamilyUpdated(familyId, name);
    }
    
    /**
     * @dev Gets the number of children in a family
     */
    function getChildCount(uint256 familyId) public view returns (uint256) {
        return families[familyId].childAddresses.length;
    }
    
    /**
     * @dev Gets a child's information by address
     */
    function getChild(uint256 familyId, address childAddress) public view returns (
        address, string memory, uint8, bool
    ) {
        Child storage child = families[familyId].children[childAddress];
        return (child.walletAddress, child.name, child.age, child.active);
    }
    
    /**
     * @dev Gets child address by index
     */
    function getChildAddressByIndex(uint256 familyId, uint256 index) public view returns (address) {
        require(index < families[familyId].childAddresses.length, "Index out of bounds");
        return families[familyId].childAddresses[index];
    }
    
    /**
     * @dev Checks if an address is a parent in any family
     */
    function isParent(address account) public view returns (bool) {
        return parentFamilies[account].length > 0;
    }
    
    /**
     * @dev Checks if an address is a child in any family
     */
    function isChild(address account) public view returns (bool) {
        return childToFamily[account] != 0;
    }
} 