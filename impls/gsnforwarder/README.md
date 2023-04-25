# GSN Forwarder Contract Code

### Helpful links
* https://github.com/opengsn/gsn/tree/master/packages/contracts/src/forwarder
* https://docs.opengsn.org/networks/ethereum/mainnet.html

### Solidity Code We Used
The github repo doesn't contain what's needed to create a byte-compatible
contract to what is deployed. We used etherscan's API's to download the source
solidity files validated by them Note: etherscan ony does partial validation
that does not include the metadata, so the actual source code used may differ
some if it generates the same byte code.

We built with the same solidity compiler (0.8.7), same optimization flags, and
same file paths and were able to get identical byte codes minus the contract
metadata at the end.
