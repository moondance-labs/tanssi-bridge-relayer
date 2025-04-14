// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2023 Snowfork <hello@snowfork.com>
pragma solidity 0.8.25;

library GatewayCoreStorage {
    struct Layout {
        // Owner of the gateway for configuration purposes.
        address owner;
        // Address of the Symbiotic middleware to properly execute messages.
        address middleware;
    }

    bytes32 internal constant SLOT = keccak256("tanssi-bridge-relayer.gateway.core");

    function layout() internal pure returns (Layout storage ptr) {
        bytes32 slot = SLOT;
        assembly {
            ptr.slot := slot
        }
    }
}
