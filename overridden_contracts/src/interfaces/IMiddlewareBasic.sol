// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2023 Snowfork <hello@snowfork.com>
pragma solidity ^0.8.0;

interface IMiddlewareBasic {
    /**
     * @notice Distribute rewards for a specific era contained in an epoch by providing a Merkle root, total points, total amount of tokens and the token address of the rewards.
     * @param epoch network epoch of the middleware
     * @param eraIndex era index of Starlight's rewards distribution
     * @param totalPointsToken total amount of points for the reward distribution
     * @param amount amount of tokens to distribute
     * @param root Merkle root of the reward distribution
     * @param tokenAddress The token address of the rewards
     * @dev This function is called by the gateway only
     * @dev Emit DistributeRewards event.
     */
    function distributeRewards(
        uint256 epoch,
        uint256 eraIndex,
        uint256 totalPointsToken,
        uint256 amount,
        bytes32 root,
        address tokenAddress
    ) external;

    /**
     * @notice Slashes an operator's stake
     * @dev Only the owner can call this function
     * @dev This function first updates the stake cache for the target epoch
     * @param epoch The epoch number
     * @param operatorKey The operator key to slash
     * @param percentage Percentage to slash, represented as parts per billion.
     */
    function slash(uint48 epoch, bytes32 operatorKey, uint256 percentage) external;
    /**
     * @notice Determines which epoch a timestamp belongs to
     * @param timestamp The timestamp to check
     * @return epoch The corresponding epoch number
     */
    function getEpochAtTs(uint48 timestamp) external view returns (uint48 epoch);
}
