// SPDX-License-Identifier: Apache-2.0
pragma solidity 0.8.25;

import {IMiddlewareBasic} from "../../src/interfaces/IMiddlewareBasic.sol";

contract MockOMiddlewareReverter is IMiddlewareBasic {
    function distributeRewards(
        uint256 epoch,
        uint256 eraIndex,
        uint256 totalPointsToken,
        uint256 tokensInflatedToken,
        bytes32 rewardsRoot
    ) external {
        revert("no distribute rewards");
    }

    function slash(uint48 epoch, bytes32 operatorKey, uint256 percentage) external {
        revert("no process slash");
    }

    function getEpochAtTs(uint48 timestamp) external view returns (uint48 epoch) {
        return timestamp;
    }
}
