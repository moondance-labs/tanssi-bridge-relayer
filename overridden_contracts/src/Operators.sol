// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2023 Snowfork <hello@snowfork.com>
pragma solidity 0.8.25;

import {console2} from "forge-std/console2.sol";
import {BeefyClient} from "./BeefyClient.sol";
import {ScaleCodec} from "./utils/ScaleCodec.sol";
import {OSubstrateTypes} from "./libraries/OSubstrateTypes.sol";
import {MultiAddress, Ticket, Costs, ParaID} from "./Types.sol";

import {IOGateway} from "./interfaces/IOGateway.sol";

library Operators {
    error Operators__OperatorsLengthTooLong();
    error Operators__OperatorsKeysCannotBeEmpty();

    uint16 private constant MAX_OPERATORS = 1000;

    function encodeOperatorsData(bytes32[] calldata operatorsKeys, uint48 epoch)
        internal
        returns (Ticket memory ticket)
    {
        if (operatorsKeys.length == 0) {
            revert Operators__OperatorsKeysCannotBeEmpty();
        }
        uint256 validatorsKeysLength = operatorsKeys.length;

        if (validatorsKeysLength > MAX_OPERATORS) {
            revert Operators__OperatorsLengthTooLong();
        }

        // TODO: This is a type from Snowbridge, do we want our own simplified Ticket type?
        ticket.dest = ParaID.wrap(0);
        // TODO For now mock it to 0
        ticket.costs = Costs(0, 0);

        ticket.payload = OSubstrateTypes.EncodedOperatorsData(operatorsKeys, uint32(validatorsKeysLength), epoch);
        emit IOGateway.OperatorsDataCreated(validatorsKeysLength, ticket.payload);
    }
}
