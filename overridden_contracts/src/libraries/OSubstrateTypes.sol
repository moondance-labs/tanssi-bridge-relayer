// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2023 Snowfork <hello@snowfork.com>
pragma solidity 0.8.25;

import {ScaleCodec} from "../utils/ScaleCodec.sol";
import {ParaID} from "../Types.sol";

library OSubstrateTypes {
    enum Message {
        V0
    }

    enum OutboundCommandV1 {
        ReceiveValidators
    }

    function EncodedOperatorsData(bytes32[] calldata operatorsKeys, uint32 operatorsCount, uint48 epoch)
        internal
        view
        returns (bytes memory)
    {
        bytes memory operatorsFlattened = new bytes(operatorsCount * 32);
        for (uint32 i = 0; i < operatorsCount; i++) {
            for (uint32 j = 0; j < 32; j++) {
                operatorsFlattened[i * 32 + j] = operatorsKeys[i][j];
            }
        }

        return bytes.concat(
            bytes4(0x70150038),
            bytes1(uint8(Message.V0)),
            bytes1(uint8(OutboundCommandV1.ReceiveValidators)),
            ScaleCodec.encodeCompactU32(operatorsCount),
            operatorsFlattened,
            ScaleCodec.encodeU64(uint64(epoch))
        );
    }
}
