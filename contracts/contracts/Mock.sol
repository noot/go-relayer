// SPDX-License-Identifier: LGPLv3
pragma solidity ^0.8.5;

import "./ERC2771Context.sol";

// Mock is a mocked contract that implements ERC2771 recipient functionality.
contract Mock is ERC2771Context {
	event Withdraw(address recipient, address relayer, uint256 value, uint256 fee);

	event FallbackCalled();

	constructor(address trustedForwarder) ERC2771Context(trustedForwarder) {}

	receive() external payable {}

	fallback() external payable {
		emit FallbackCalled();
	}

	// todo: express fee as a percentage, instead of an eth value?
	// would require more math on the solidity side
	function withdraw(uint256 value, uint256 fee) external {
		require(isTrustedForwarder(msg.sender), "function must be called by forwarder");
		require(address(this).balance >= value, "value to transfer is higher than contract balance");
		require(fee < value, "fee is higher than value");
		payable(tx.origin).transfer(fee);
		payable(_msgSender()).transfer(value-fee); // todo: safemath, i guess
		emit Withdraw(_msgSender(), tx.origin, value-fee, fee);
	}
}