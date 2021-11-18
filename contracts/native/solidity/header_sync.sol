pragma solidity >=0.5.0 <0.9.0;
pragma experimental ABIEncoderV2;

interface HeaderSync {
    event OKEpochSwitchInfoEvent(uint64 chainID, string BlockHash, uint64 Height, string NextValidatorsHash, string InfoChainID, uint64 BlockHeight);
    event syncHeader(uint64 chainID, uint64 height, string blockHash, uint256 BlockHeight);

    function syncBlockHeader(uint64 ChainID, address Address, bytes[] calldata Headers) external returns(bool success);
    function syncGenesisHeader(uint64 ChainID, bytes calldata GenesisHeader) external returns(bool success);
    function syncCrossChainMsg(uint64 ChainID, address Address, bytes[] calldata CrossChainMsgs) external returns(bool success);
}


