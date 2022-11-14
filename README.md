# vivid-bridge
**vivid-bridge** is a trustless crosschain bridge for Vivid blockchain for transefering ERC20-VID from Ethereum to Vivid.

Bridges as applications 'lock' assets on one chain in exchange for wrapped assets on another chain. The wrapped assets can then be exchanged for the original 'backing' asset. VivdBridge locks the ERC20VID on Ethereum Mainnet and issues WrappedVID on Vivid Network. The Wrapped VID is exchanged with Native VID. This will maintain the total supply of ERC20 VID i.e. the total NativeVID on Vivid Network is always equal to locked ERC20 VID,

# Comparison with Bridges

A bridge between two blockchains that requires either a trusted intermediary, committee or an honest majority assumption to ensure that funds can’t be stolen. A trusted bridge can be exploited more easily than a trust-minimized bridge because it provides weaker guarantees for its security.

Examples of trusted bridges are the bridge for WBTC which is run by BitGo, a committee based bridge like Axelar, or the IBC bridge between Osmosis and the Cosmos Hub.

Cross-chain communication between application-specific blockchains in Cosmos creates the potential for high horizontal scaleability with transaction finality. These design features provide convincing solutions to well-known problems that plague other platforms, such as transaction costs, network capacity, and transaction confirmation finality.

## Comparsion with Trusted Bridges
--TODO--

## Comparison with Polygon-PoS Bridge
--TODO--

## Comparison eith Solana Wormhole bridge
--TODO--

---
# Operation

![Overview of VividBridge](architecture/images/bridge-overview.drawio.svg)
## Ethereum Mainnet to Vivid Network
* An account transfers ERC20-VID to VividBridgeBank contract on Ethereum (D1, D2)
* Retrieve the resulting TokenToNativeMessage from the VividBridge Network(D3)
* Call ExchangeContract on Vivid suppling TokenToNativeMessage(D4).
* WrappedVID is locked and recepient account is deposited with Native-VID

## Vivid Network to Ethereum Mainnet
* An account transfers Native-VID to VividBridgeBank on Vivid(W1, W2)
* Retrieve the resulting NativeToTokenMessage from the VividBridge Network(W3).
* Call VividBridgeBank Contract on Ethereum to process NativeToTokenMessage(W4).
* WrappedVID is locked in the bank and VID Token is deposites in the ERC20-VID for the account(W5, W6).

```
BridgeMessage(
    byte        version,
    u8          len_signatures,
    [][66]byte  signatures,

    u32         timestamp,
    int         nonce,
    u16         emitter_chain, 
    [32]byte    emitter_address,

    address     token,
    uint256     amount,
    uint16      recipientChain,
    bytes32     recipient,
    uint256     arbiterFee,
    bytes32     txn_hash
) returns int sequenceNumber

```
## References:

Wormhole Relayer
https://book.wormhole.com/wormhole/6_relayers.html
