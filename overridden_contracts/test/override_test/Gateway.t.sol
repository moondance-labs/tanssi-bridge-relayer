// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2023 Snowfork <hello@snowfork.com>
pragma solidity 0.8.25;

import {Test, console2, Vm} from "forge-std/Test.sol";
import {
    AgentExecuteCommand,
    InboundMessage,
    OperatingMode,
    ParaID,
    ChannelID,
    Command,
    multiAddressFromBytes32,
    multiAddressFromBytes20
} from "../../src/Types.sol";
import {IGateway} from "../../src/interfaces/IGateway.sol";
import {IMiddlewareBasic} from "../../src/interfaces/IMiddlewareBasic.sol";
import {MockGateway} from "../mocks/MockGateway.sol";
import {
    CreateAgentParams,
    CreateChannelParams,
    SetOperatingModeParams,
    RegisterForeignTokenParams
} from "../../src/Params.sol";
import {OperatingMode, ParaID, Command} from "../../src/Types.sol";
import {GatewayProxy} from "../../src/GatewayProxy.sol";
import {MultiAddress} from "../../src/MultiAddress.sol";
import {AgentExecutor} from "../../src/AgentExecutor.sol";
import {Verification} from "../../src/Verification.sol";

import {Strings} from "openzeppelin/utils/Strings.sol";

import {Gateway} from "../../src/Gateway.sol";
import {IOGateway} from "../../src/interfaces/IOGateway.sol";
import {Operators} from "../../src/Operators.sol";
import {Assets} from "../../src/Assets.sol";
import {MockOGateway} from "../mocks/MockOGateway.sol";
import {Token} from "../../src/Token.sol";
//NEW
import {WETH9} from "canonical-weth/WETH9.sol";
import {UD60x18, ud60x18, convert} from "prb/math/src/UD60x18.sol";

contract GatewayTest is Test {
    // Emitted when token minted/burnt/transfered
    event Transfer(address indexed from, address indexed to, uint256 value);

    ParaID public bridgeHubParaID = ParaID.wrap(1013);
    bytes32 public bridgeHubAgentID = 0x03170a2e7597b7b7e3d84c05391d139a62b157e78786d8c082f29dcf4c111314;
    address public bridgeHubAgent;

    ParaID public assetHubParaID = ParaID.wrap(1000);
    bytes32 public assetHubAgentID = 0x81c5ab2571199e3188135178f3c2c8e2d268be1313d029b30f534fa579b69b79;
    address public assetHubAgent;

    address public relayer;

    bytes32[] public proof = [bytes32(0x2f9ee6cfdf244060dc28aa46347c5219e303fc95062dd672b4e406ca5c29764b)];
    bytes public parachainHeaderProof = bytes("validProof");

    MockOGateway public gatewayLogic;
    GatewayProxy public gateway;

    WETH9 public token;

    address public account1;
    address public account2;
    address public middleware = makeAddr("middleware");

    uint64 public maxDispatchGas = 500_000;
    uint256 public maxRefund = 1 ether;
    uint256 public reward = 1 ether;
    bytes32 public messageID = keccak256("cabbage");

    // remote fees in DOT
    uint128 public outboundFee = 1e10;
    uint128 public registerTokenFee = 0;
    uint128 public sendTokenFee = 1e10;
    uint128 public createTokenFee = 1e10;
    uint128 public maxDestinationFee = 1e11;

    MultiAddress public recipientAddress32;
    MultiAddress public recipientAddress20;

    // For DOT
    uint8 public foreignTokenDecimals = 10;

    // ETH/DOT exchange rate
    UD60x18 public exchangeRate = ud60x18(0.0025e18);
    UD60x18 public multiplier = ud60x18(1e18);

    // tokenID for DOT
    bytes32 public dotTokenID;

    uint256 public constant SLASH_FRACTION = 500_000;
    uint256 public constant ONE_DAY = 86400;

    ChannelID internal constant PRIMARY_GOVERNANCE_CHANNEL_ID = ChannelID.wrap(bytes32(uint256(1)));
    ChannelID internal constant SECONDARY_GOVERNANCE_CHANNEL_ID = ChannelID.wrap(bytes32(uint256(2)));

    // It's generated using VALIDATORS_DATA using scaleCodec. See OSubstrateTypes.EncodedOperatorsData in OSubstrateTypes.sol
    bytes private constant FINAL_VALIDATORS_PAYLOAD =
        hex"7015003800000cd43593c715fdd31c61141abd04a99fd6822c8558854ccde39a5684e7a56da27d90b5ab205c6974c9ea841be688864633dc9ca8a357843eeacf2314649965fe228eaf04151687736326c9fea17e25fc5287613693c912909cb226aa4794f26a480100000000000000";

    bytes32[] private VALIDATORS_DATA = [
        bytes32(0xd43593c715fdd31c61141abd04a99fd6822c8558854ccde39a5684e7a56da27d),
        bytes32(0x90b5ab205c6974c9ea841be688864633dc9ca8a357843eeacf2314649965fe22),
        bytes32(0x8eaf04151687736326c9fea17e25fc5287613693c912909cb226aa4794f26a48)
    ];

    // Test vector generated by: https://github.com/moondance-labs/tanssi/blob/242196324a37ac0020a7c7955bffe09670f63751/primitives/bridge/src/tests.rs#L84
    bytes private constant TEST_VECTOR_SLASH_DATA =
        hex"0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000002A000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000030404040404040404040404040404040404040404040404040404040404040404000000000000000000000000000000000000000000000000000000000000138800000000000000000000000000000000000000000000000000000000000001F405050505050505050505050505050505050505050505050505050505050505050000000000000000000000000000000000000000000000000000000000000FA0000000000000000000000000000000000000000000000000000000000000019006060606060606060606060606060606060606060606060606060606060606060000000000000000000000000000000000000000000000000000000000000BB8000000000000000000000000000000000000000000000000000000000000012C";

    bytes private constant TEST_VECTOR_REWARDS_DATA =
        hex"00000000000000000000000000000000000000000000000000000000075BCD15000000000000000000000000000000000000000000000000000000000000002A00000000000000000000000000000000000000000000000000007048860DDF79000000000000000000000000000000000000000000000000000000E5F4C8F3CAb6e16d27ac5ab427a7f68900ac5559ce272dc6c37c82b3e052246c82244c50e40101010101010101010101010101010101010101010101010101010101010101";

    function setUp() public {
        AgentExecutor executor = new AgentExecutor();
        gatewayLogic = new MockOGateway(
            address(0), address(executor), bridgeHubParaID, bridgeHubAgentID, foreignTokenDecimals, maxDestinationFee
        );
        Gateway.Config memory config = Gateway.Config({
            mode: OperatingMode.Normal,
            deliveryCost: outboundFee,
            registerTokenFee: registerTokenFee,
            assetHubParaID: assetHubParaID,
            assetHubAgentID: assetHubAgentID,
            assetHubCreateAssetFee: createTokenFee,
            assetHubReserveTransferFee: sendTokenFee,
            exchangeRate: exchangeRate,
            multiplier: multiplier,
            rescueOperator: 0x4B8a782D4F03ffcB7CE1e95C5cfe5BFCb2C8e967
        });
        gateway = new GatewayProxy(address(gatewayLogic), abi.encode(config));
        MockGateway(address(gateway)).setCommitmentsAreVerified(true);

        SetOperatingModeParams memory params = SetOperatingModeParams({mode: OperatingMode.Normal});
        MockGateway(address(gateway)).setOperatingModePublic(abi.encode(params));

        bridgeHubAgent = IGateway(address(gateway)).agentOf(bridgeHubAgentID);
        assetHubAgent = IGateway(address(gateway)).agentOf(assetHubAgentID);

        // fund the message relayer account
        relayer = makeAddr("relayer");

        // Features

        token = new WETH9();

        account1 = makeAddr("account1");
        account2 = makeAddr("account2");

        // create tokens for account 1
        hoax(account1);
        token.deposit{value: 500}();

        // create tokens for account 2
        token.deposit{value: 500}();

        recipientAddress32 = multiAddressFromBytes32(keccak256("recipient"));
        recipientAddress20 = multiAddressFromBytes20(bytes20(keccak256("recipient")));

        dotTokenID = bytes32(uint256(1));
    }

    function _makeReportSlashesCommand(uint256 slashFraction) public pure returns (Command, bytes memory) {
        IOGateway.Slash[] memory slashes = new IOGateway.Slash[](1);
        slashes[0] = IOGateway.Slash({operatorKey: bytes32(uint256(1)), slashFraction: slashFraction, epoch: 1});
        uint256 eraIndex = 1;
        return (Command.ReportSlashes, abi.encode(IOGateway.SlashParams({eraIndex: eraIndex, slashes: slashes})));
    }

    function _makeReportRewardsCommand() public returns (Command, bytes memory, address) {
        uint256 epoch = 0;
        uint256 eraIndex = 1;
        uint256 totalPointsToken = 1 ether;
        uint256 tokensInflatedToken = 1 ether;
        bytes32 rewardsRoot = bytes32(uint256(1));
        bytes32 foreignTokenId = bytes32(uint256(1));

        RegisterForeignTokenParams memory params =
            RegisterForeignTokenParams({foreignTokenID: foreignTokenId, name: "Test", symbol: "TST", decimals: 10});

        vm.expectEmit(true, true, false, false);
        emit IGateway.ForeignTokenRegistered(foreignTokenId, address(0));
        MockGateway(address(gateway)).registerForeignTokenPublic(abi.encode(params));

        address tokenAddress = MockGateway(address(gateway)).tokenAddressOf(foreignTokenId);

        return (
            Command.ReportRewards,
            abi.encode(epoch, eraIndex, totalPointsToken, tokensInflatedToken, rewardsRoot, foreignTokenId),
            tokenAddress
        );
    }

    function makeMockProof() public pure returns (Verification.Proof memory) {
        return Verification.Proof({
            leafPartial: Verification.MMRLeafPartial({
                version: 0,
                parentNumber: 0,
                parentHash: bytes32(0),
                nextAuthoritySetID: 0,
                nextAuthoritySetLen: 0,
                nextAuthoritySetRoot: 0
            }),
            leafProof: new bytes32[](0),
            leafProofOrder: 0,
            parachainHeadsRoot: bytes32(0)
        });
    }

    function createLongOperatorsData() public view returns (bytes32[] memory) {
        bytes32[] memory result = new bytes32[](1001);

        for (uint256 i = 0; i <= 1000; i++) {
            result[i] = VALIDATORS_DATA[i % 3];
        }

        return result;
    }

    function _createParaIDAndAgent() public returns (ParaID) {
        ParaID paraID = ParaID.wrap(1);
        bytes32 agentID = keccak256("1");

        MockGateway(address(gateway)).createAgentPublic(abi.encode(CreateAgentParams({agentID: agentID})));

        CreateChannelParams memory params =
            CreateChannelParams({channelID: paraID.into(), agentID: agentID, mode: OperatingMode.Normal});

        MockGateway(address(gateway)).createChannelPublic(abi.encode(params));
        return paraID;
    }

    function testSendOperatorsData() public {
        // FINAL_VALIDATORS_PAYLOAD has been encoded with epoch 1.
        uint48 epoch = 1;
        IOGateway(address(gateway)).setMiddleware(middleware);
        // Create mock agent and paraID
        vm.prank(middleware);
        vm.expectEmit(true, false, false, true);
        emit IGateway.OutboundMessageAccepted(PRIMARY_GOVERNANCE_CHANNEL_ID, 1, messageID, FINAL_VALIDATORS_PAYLOAD);

        IOGateway(address(gateway)).sendOperatorsData(VALIDATORS_DATA, epoch);
    }

    function testShouldNotSendOperatorsDataBecauseOperatorsTooLong() public {
        bytes32[] memory longOperatorsData = createLongOperatorsData();
        uint48 epoch = 42;
        IOGateway(address(gateway)).setMiddleware(middleware);
        vm.prank(middleware);
        vm.expectRevert(Operators.Operators__OperatorsLengthTooLong.selector);
        IOGateway(address(gateway)).sendOperatorsData(longOperatorsData, epoch);
    }

    function testSendOperatorsDataWith50Entries() public {
        string memory root = vm.projectRoot();
        string memory path = string.concat(root, "/test/snowbridge-data/test_vector_message_validator_50.json");
        string memory json = vm.readFile(path);

        // Get payload
        bytes memory final_payload = vm.parseJsonBytes(json, "$.payload");

        // Get accounts array
        bytes32[] memory accounts = abi.decode(vm.parseJson(json, "$.accounts"), (bytes32[]));

        uint48 epoch = abi.decode(vm.parseJson(json, "$.epoch"), (uint48));

        IOGateway(address(gateway)).setMiddleware(middleware);

        vm.prank(middleware);
        vm.expectEmit(true, false, false, true);
        emit IGateway.OutboundMessageAccepted(PRIMARY_GOVERNANCE_CHANNEL_ID, 1, messageID, final_payload);

        IOGateway(address(gateway)).sendOperatorsData(accounts, epoch);
    }

    function testSendOperatorsDataWith400Entries() public {
        string memory root = vm.projectRoot();
        string memory path = string.concat(root, "/test/snowbridge-data/test_vector_message_validator_400.json");
        string memory json = vm.readFile(path);

        // Get payload
        bytes memory final_payload = vm.parseJsonBytes(json, "$.payload");

        // Get accounts array
        bytes32[] memory accounts = abi.decode(vm.parseJson(json, "$.accounts"), (bytes32[]));
        uint48 epoch = abi.decode(vm.parseJson(json, "$.epoch"), (uint48));

        IOGateway(address(gateway)).setMiddleware(middleware);

        vm.prank(middleware);
        vm.expectEmit(true, false, false, true);
        emit IGateway.OutboundMessageAccepted(PRIMARY_GOVERNANCE_CHANNEL_ID, 1, messageID, final_payload);

        IOGateway(address(gateway)).sendOperatorsData(accounts, epoch);
    }

    function testSendOperatorsDataWith1000Entries() public {
        string memory root = vm.projectRoot();
        string memory path = string.concat(root, "/test/snowbridge-data/test_vector_message_validator_1000.json");
        string memory json = vm.readFile(path);

        // Get payload
        bytes memory final_payload = vm.parseJsonBytes(json, "$.payload");

        // Get accounts array
        bytes32[] memory accounts = abi.decode(vm.parseJson(json, "$.accounts"), (bytes32[]));
        uint48 epoch = abi.decode(vm.parseJson(json, "$.epoch"), (uint48));

        IOGateway(address(gateway)).setMiddleware(middleware);

        vm.prank(middleware);
        vm.expectEmit(true, false, false, true);
        emit IGateway.OutboundMessageAccepted(PRIMARY_GOVERNANCE_CHANNEL_ID, 1, messageID, final_payload);

        IOGateway(address(gateway)).sendOperatorsData(accounts, epoch);
    }

    function testOwnerCanChangeMiddleware() public {
        vm.expectEmit(true, true, false, false);
        emit IOGateway.MiddlewareChanged(address(0), 0x0123456789012345678901234567890123456789);

        IOGateway(address(gateway)).setMiddleware(0x0123456789012345678901234567890123456789);

        require(IOGateway(address(gateway)).s_middleware() == 0x0123456789012345678901234567890123456789);
    }

    function testNonOwnerCantChangeMiddleware() public {
        address notOwner = makeAddr("notOwner");
        vm.prank(notOwner);
        vm.expectRevert(abi.encodeWithSelector(Gateway.Unauthorized.selector));
        IOGateway(address(gateway)).setMiddleware(0x9876543210987654321098765432109876543210);
    }

    function testDecodeSlashes() public {
        uint256 eraIndex = 42;
        IOGateway.Slash[] memory slashes = new IOGateway.Slash[](3);
        bytes32 alice = 0x0404040404040404040404040404040404040404040404040404040404040404;
        bytes32 bob = 0x0505050505050505050505050505050505050505050505050505050505050505;
        bytes32 charlie = 0x0606060606060606060606060606060606060606060606060606060606060606;

        slashes[0] = IOGateway.Slash({operatorKey: alice, slashFraction: 5_000, epoch: 500});
        slashes[1] = IOGateway.Slash({operatorKey: bob, slashFraction: 4_000, epoch: 400});
        slashes[2] = IOGateway.Slash({operatorKey: charlie, slashFraction: 3_000, epoch: 300});

        assertEq(abi.encode(IOGateway.SlashParams({eraIndex: eraIndex, slashes: slashes})), TEST_VECTOR_SLASH_DATA);
    }

    // middleware not set, should not be able to process slash
    function testSubmitSlashesWithoutMiddleware() public {
        deal(assetHubAgent, 50 ether);

        (Command command, bytes memory params) = _makeReportSlashesCommand(SLASH_FRACTION);

        vm.expectEmit(true, true, true, true);
        emit IOGateway.UnableToProcessSlashMessageB(abi.encodeWithSelector(Gateway.MiddlewareNotSet.selector));
        // Expect the gateway to emit `InboundMessageDispatched`
        vm.expectEmit(true, true, true, true);
        emit IGateway.InboundMessageDispatched(assetHubParaID.into(), 1, messageID, false);

        hoax(relayer, 1 ether);
        IGateway(address(gateway)).submitV1(
            InboundMessage(assetHubParaID.into(), 1, command, params, maxDispatchGas, maxRefund, reward, messageID),
            proof,
            makeMockProof()
        );
    }

    // middleware set, but not complying with the interface, should not process slash
    function testSubmitSlashesWithMiddlewareNotComplyingInterface() public {
        deal(assetHubAgent, 50 ether);

        (Command command, bytes memory params) = _makeReportSlashesCommand(SLASH_FRACTION);

        IOGateway(address(gateway)).setMiddleware(0x0123456789012345678901234567890123456789);

        bytes memory empty;
        // Expect the gateway to emit `InboundMessageDispatched`
        // For some reason when you are loading an address not complying an interface, you get an empty message
        // It still serves us to know that this is the reason
        vm.expectEmit(true, true, true, true);
        emit IOGateway.UnableToProcessSlashMessageB(empty);
        vm.expectEmit(true, true, true, true);
        emit IGateway.InboundMessageDispatched(assetHubParaID.into(), 1, messageID, false);

        hoax(relayer, 1 ether);
        IGateway(address(gateway)).submitV1(
            InboundMessage(assetHubParaID.into(), 1, command, params, maxDispatchGas, maxRefund, reward, messageID),
            proof,
            makeMockProof()
        );
    }

    // middleware set, complying interface but slash reverts
    function testSubmitSlashesWithMiddlewareComplyingInterfaceAndSlashRevert() public {
        deal(assetHubAgent, 50 ether);

        (Command command, bytes memory params) = _makeReportSlashesCommand(SLASH_FRACTION);

        bytes memory expectedError = bytes("no process slash");

        // We mock the call so that it reverts
        vm.mockCallRevert(
            address(middleware), abi.encodeWithSelector(IMiddlewareBasic.slash.selector), "no process slash"
        );

        IOGateway(address(gateway)).setMiddleware(address(middleware));

        IOGateway.Slash memory expectedSlash =
            IOGateway.Slash({operatorKey: bytes32(uint256(1)), slashFraction: SLASH_FRACTION, epoch: 1});

        vm.expectEmit(true, true, true, true);
        emit IOGateway.UnableToProcessIndividualSlashB(
            expectedSlash.operatorKey, expectedSlash.slashFraction, expectedSlash.epoch, expectedError
        );
        vm.expectEmit(true, true, true, true);
        emit IGateway.InboundMessageDispatched(assetHubParaID.into(), 1, messageID, true);

        hoax(relayer, 1 ether);
        IGateway(address(gateway)).submitV1(
            InboundMessage(assetHubParaID.into(), 1, command, params, maxDispatchGas, maxRefund, reward, messageID),
            proof,
            makeMockProof()
        );
    }

    // middleware set, complying interface and slash processed
    function testSubmitSlashesWithMiddlewareComplyingInterfaceAndSlashProcessed() public {
        deal(assetHubAgent, 50 ether);

        (Command command, bytes memory params) = _makeReportSlashesCommand(SLASH_FRACTION);

        // We mock the call so that it does not revert
        vm.mockCall(address(middleware), abi.encodeWithSelector(IMiddlewareBasic.slash.selector), abi.encode(10));

        IOGateway(address(gateway)).setMiddleware(address(middleware));

        // Since we are asserting all fields, the last one is a true, therefore meaning
        // that the dispatch went through correctly

        vm.expectEmit(true, true, true, true);
        emit IGateway.InboundMessageDispatched(assetHubParaID.into(), 1, messageID, true);

        hoax(relayer, 1 ether);
        vm.recordLogs();
        IGateway(address(gateway)).submitV1(
            InboundMessage(assetHubParaID.into(), 1, command, params, maxDispatchGas, maxRefund, reward, messageID),
            proof,
            makeMockProof()
        );

        Vm.Log[] memory entries = vm.getRecordedLogs();
        // We assert none of the slash error events has been emitted
        for (uint256 i = 0; i < entries.length; i++) {
            assertNotEq(entries[i].topics[0], IOGateway.UnableToProcessIndividualSlashB.selector);
            assertNotEq(entries[i].topics[0], IOGateway.UnableToProcessIndividualSlashS.selector);
        }
    }

    function testDecodeRewards() public {
        uint256 epoch = 123_456_789;
        uint256 eraIndex = 42;
        uint256 totalPointsToken = 123_456_789_012_345;
        uint256 tokensInflatedToken = 987_654_321_098;
        bytes32 rewardsRoot = 0xb6e16d27ac5ab427a7f68900ac5559ce272dc6c37c82b3e052246c82244c50e4;
        bytes32 foreignTokenId = 0x0101010101010101010101010101010101010101010101010101010101010101;

        assertEq(
            abi.encode(epoch, eraIndex, totalPointsToken, tokensInflatedToken, rewardsRoot, foreignTokenId),
            TEST_VECTOR_REWARDS_DATA
        );
    }

    function testSubmitRewards() public {
        deal(assetHubAgent, 50 ether);

        (Command command, bytes memory params, address tokenAddress) = _makeReportRewardsCommand();

        // We mock the call so that it does not revert
        vm.mockCall(
            address(middleware), abi.encodeWithSelector(IMiddlewareBasic.distributeRewards.selector), abi.encode(true)
        );

        IOGateway(address(gateway)).setMiddleware(address(middleware));

        // Expect the gateway to emit `InboundMessageDispatched`
        vm.expectEmit(true, true, true, true);
        emit IGateway.InboundMessageDispatched(assetHubParaID.into(), 1, messageID, true);

        hoax(relayer, 1 ether);
        IGateway(address(gateway)).submitV1(
            InboundMessage(assetHubParaID.into(), 1, command, params, maxDispatchGas, maxRefund, reward, messageID),
            proof,
            makeMockProof()
        );

        assert(Token(tokenAddress).balanceOf(address(middleware)) > 0);
    }

    function testSubmitRewardsWithoutMiddleware() public {
        deal(assetHubAgent, 50 ether);

        (Command command, bytes memory params, address tokenAddress) = _makeReportRewardsCommand();

        vm.expectEmit(true, true, true, true);
        emit IOGateway.UnableToProcessRewardsMessageB(abi.encodeWithSelector(Gateway.MiddlewareNotSet.selector));
        // Expect the gateway to emit `InboundMessageDispatched`
        vm.expectEmit(true, true, true, true);
        emit IGateway.InboundMessageDispatched(assetHubParaID.into(), 1, messageID, false);

        hoax(relayer, 1 ether);
        IGateway(address(gateway)).submitV1(
            InboundMessage(assetHubParaID.into(), 1, command, params, maxDispatchGas, maxRefund, reward, messageID),
            proof,
            makeMockProof()
        );
    }

    // middleware set, but not complying with the interface, should not process rewards
    function testSubmitRewardsWithMiddlewareNotComplyingInterface() public {
        deal(assetHubAgent, 50 ether);

        (Command command, bytes memory params, address tokenAddress) = _makeReportRewardsCommand();

        IOGateway(address(gateway)).setMiddleware(0x0123456789012345678901234567890123456789);

        bytes memory empty;
        vm.expectEmit(true, true, true, true);
        emit IOGateway.UnableToProcessRewardsMessageB(empty);
        vm.expectEmit(true, true, true, true);
        emit IGateway.InboundMessageDispatched(assetHubParaID.into(), 1, messageID, false);

        hoax(relayer, 1 ether);
        IGateway(address(gateway)).submitV1(
            InboundMessage(assetHubParaID.into(), 1, command, params, maxDispatchGas, maxRefund, reward, messageID),
            proof,
            makeMockProof()
        );
    }

    // middleware set, complying interface but rewards reverts
    function testSubmitRewardsWithMiddlewareComplyingInterfaceAndRewardsRevert() public {
        deal(assetHubAgent, 50 ether);

        (Command command, bytes memory params, address tokenAddress) = _makeReportRewardsCommand();

        bytes memory expectedError = bytes("can't process rewards"); //This should actually come from IODefaultOperatorRewards

        // We mock the call so that it reverts
        vm.mockCallRevert(
            address(middleware),
            abi.encodeWithSelector(IMiddlewareBasic.distributeRewards.selector),
            "can't process rewards"
        );

        IOGateway(address(gateway)).setMiddleware(address(middleware));

        uint256 expectedEpoch = 0;
        uint256 expectedEraIndex = 1;
        uint256 expectedTotalPointsToken = 1 ether;
        uint256 expectedTotalTokensInflated = 1 ether;
        bytes32 expectedRewardsRoot = bytes32(uint256(1));
        bytes32 expectedForeignTokenId = bytes32(uint256(1));

        address expectedTokenAddress = MockGateway(address(gateway)).tokenAddressOf(expectedForeignTokenId);
        bytes memory expectedBytes = abi.encodeWithSelector(
            Gateway.EUnableToProcessRewardsB.selector,
            expectedEpoch,
            expectedEraIndex,
            expectedTokenAddress,
            expectedTotalPointsToken,
            expectedTotalTokensInflated,
            expectedRewardsRoot,
            expectedError
        );

        vm.expectEmit(true, true, true, true);
        emit IOGateway.UnableToProcessRewardsMessageB(expectedBytes);
        vm.expectEmit(true, true, true, true);
        emit IGateway.InboundMessageDispatched(assetHubParaID.into(), 1, messageID, false);

        hoax(relayer, 1 ether);
        IGateway(address(gateway)).submitV1(
            InboundMessage(assetHubParaID.into(), 1, command, params, maxDispatchGas, maxRefund, reward, messageID),
            proof,
            makeMockProof()
        );
    }

    // middleware set, complying interface and rewards processed
    function testSubmitRewardsWithMiddlewareComplyingInterfaceAndRewardsProcessed() public {
        deal(assetHubAgent, 50 ether);

        (Command command, bytes memory params, address tokenAddress) = _makeReportRewardsCommand();

        // We mock the call so that it does not revert
        vm.mockCall(
            address(middleware), abi.encodeWithSelector(IMiddlewareBasic.distributeRewards.selector), abi.encode(true)
        );

        IOGateway(address(gateway)).setMiddleware(address(middleware));

        vm.expectEmit(true, true, true, true);
        emit IGateway.InboundMessageDispatched(assetHubParaID.into(), 1, messageID, true);

        hoax(relayer, 1 ether);
        vm.recordLogs();
        IGateway(address(gateway)).submitV1(
            InboundMessage(assetHubParaID.into(), 1, command, params, maxDispatchGas, maxRefund, reward, messageID),
            proof,
            makeMockProof()
        );

        Vm.Log[] memory entries = vm.getRecordedLogs();
        // We assert none of the rewards error events has been emitted
        for (uint256 i = 0; i < entries.length; i++) {
            assertNotEq(entries[i].topics[0], IOGateway.UnableToProcessRewardsMessageB.selector);
            assertNotEq(entries[i].topics[0], IOGateway.UnableToProcessRewardsMessageS.selector);
        }
    }
}
