// SPDX-License-Identifier: LGPLv3
pragma solidity ^0.8.5;

import "./ERC2771Context.sol";

contract Mock is ERC2771Context {
	constructor(address trustedForwarder) ERC2771Context(trustedForwarder) {}

	receive() external payable {}

	// todo: express fee as a percentage, instead of an eth value?
	// would require more math on the solidity side
	function withdraw(uint256 value, uint256 fee) external {
		require(fee < value, "fee is higher than value");
		payable(msg.sender).transfer(fee);
		payable(_msgSender()).transfer(value-fee); // todo: safemath, i guess
	}
}