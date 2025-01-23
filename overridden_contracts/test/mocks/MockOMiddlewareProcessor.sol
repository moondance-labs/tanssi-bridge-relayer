// SPDX-License-Identifier: Apache-2.0
pragma solidity 0.8.25;

import {IMiddlewareBasic} from "../../src/interfaces/IMiddlewareBasic.sol";

contract MockOMiddlewareProcessor is IMiddlewareBasic {
    event RewardProcessed();
    event SlashProcessed();

    function distributeRewards(
        uint256 epoch,
        uint256 eraIndex,
        uint256 totalPointsToken,
        uint256 tokensInflatedToken,
        bytes32 rewardsRoot
    ) external {
        emit RewardProcessed();
    }

    function slash(uint48 epoch, bytes32 operatorKey, uint256 percentage) external {
        emit SlashProcessed();
    }

    function getEpochAtTs(uint48 timestamp) external view returns (uint48 epoch) {
        return timestamp;
    }
}
