//SPDX-License-Identifier: GPL-3.0-or-later

// Copyright (C) Moondance Labs Ltd.
// This file is part of Tanssi.
// Tanssi is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
// Tanssi is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
// You should have received a copy of the GNU General Public License
// along with Tanssi.  If not, see <http://www.gnu.org/licenses/>
pragma solidity ^0.8.0;

import {ParaID} from "../Types.sol";
import {IGateway} from "./IGateway.sol";

interface IOGateway is IGateway {
    // Emitted when operators data has been created
    event OperatorsDataCreated(uint256 indexed validatorsCount, bytes payload);

    // Emitted when the middleware contract address is changed by the owner.
    event MiddlewareChanged(address indexed previousMiddleware, address indexed newMiddleware);

    // Slash struct, used to decode slashes, which are identified by
    // operatorKey to be slashed
    // slashFraction to be applied as parts per billion
    // timestamp identifying when the slash happened
    struct Slash {
        bytes32 operatorKey;
        uint256 slashFraction;
        uint256 timestamp;
    }

    function sendOperatorsData(bytes32[] calldata data) external;
}
