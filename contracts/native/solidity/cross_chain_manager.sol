pragma solidity >=0.5.0 <0.9.0;

interface CrossChainManager {
    event makeProof(string merkleValueHex, uint64 BlockHeight, string key);

    function importOuterTransfer(uint64 SourceChainID, uint32 Height, bytes calldata Proof, bytes calldata RelayerAddress, bytes calldata Extra, bytes calldata HeaderOrCrossChainMsg) external returns(bool success);

}


