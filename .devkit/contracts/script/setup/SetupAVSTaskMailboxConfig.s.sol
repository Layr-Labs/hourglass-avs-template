// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {Script, console} from "forge-std/Script.sol";
import {stdJson} from "forge-std/StdJson.sol";

import {IKeyRegistrarTypes} from "@eigenlayer-contracts/src/contracts/interfaces/IKeyRegistrar.sol";
import {OperatorSet, OperatorSetLib} from "@eigenlayer-contracts/src/contracts/libraries/OperatorSetLib.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {ITaskMailbox, ITaskMailboxTypes} from "@eigenlayer-contracts/src/contracts/interfaces/ITaskMailbox.sol";
import {IAVSTaskHook} from "@eigenlayer-contracts/src/contracts/interfaces/IAVSTaskHook.sol";
import {IECDSACertificateVerifier} from "@eigenlayer-contracts/src/contracts/interfaces/IECDSACertificateVerifier.sol";
import {IBN254CertificateVerifier} from "@eigenlayer-contracts/src/contracts/interfaces/IBN254CertificateVerifier.sol";

contract SetupAVSTaskMailboxConfig is Script {
    using stdJson for string;

    struct Context {
        address avs;
        uint256 avsPrivateKey;
        uint256 deployerPrivateKey;
        IBN254CertificateVerifier certificateVerifier;
        IECDSACertificateVerifier ecdsaCertificateVerifier;
        ITaskMailbox taskMailbox;
        IAVSTaskHook taskHook;
    }

    function run(
        string memory environment,
        uint32 executorOperatorSetId,
        uint96 taskSLA,
        uint8 curveType,
        string memory _context
    ) public {
        // Read the context
        Context memory context = _readContext(environment, _context);

        // Load the private key from the environment variable
        uint256 avsPrivateKey = vm.envUint("PRIVATE_KEY_AVS");
        address avs = vm.addr(avsPrivateKey);

        vm.startBroadcast(avsPrivateKey);
        console.log("AVS address:", avs);

        // Set the Executor Operator Set Task Config
        ITaskMailboxTypes.ExecutorOperatorSetTaskConfig memory executorOperatorSetTaskConfig = ITaskMailboxTypes
            .ExecutorOperatorSetTaskConfig({
            taskHook: context.taskHook,
            taskSLA: taskSLA,
            feeToken: IERC20(address(0)),
            curveType: IKeyRegistrarTypes.CurveType(curveType),
            feeCollector: address(0),
            consensus: ITaskMailboxTypes.Consensus({
                consensusType: ITaskMailboxTypes.ConsensusType.STAKE_PROPORTION_THRESHOLD,
                value: abi.encode(10_000)
            }),
            taskMetadata: bytes("")
        });
        context.taskMailbox.setExecutorOperatorSetTaskConfig(
            OperatorSet(avs, executorOperatorSetId), executorOperatorSetTaskConfig
        );
        ITaskMailboxTypes.ExecutorOperatorSetTaskConfig memory executorOperatorSetTaskConfigStored =
            context.taskMailbox.getExecutorOperatorSetTaskConfig(OperatorSet(avs, executorOperatorSetId));
        console.log(
            "Executor Operator Set Task Config set with curve type:",
            uint8(executorOperatorSetTaskConfigStored.curveType),
            address(executorOperatorSetTaskConfigStored.taskHook)
        );

        vm.stopBroadcast();
    }

    function _readContext(string memory environment, string memory _context) internal view returns (Context memory) {
        // Parse the context
        Context memory context;
        context.avs = stdJson.readAddress(_context, ".context.avs.address");
        context.avsPrivateKey = uint256(stdJson.readBytes32(_context, ".context.avs.avs_private_key"));
        context.deployerPrivateKey = uint256(stdJson.readBytes32(_context, ".context.deployer_private_key"));
        context.certificateVerifier = IBN254CertificateVerifier(
            stdJson.readAddress(_context, ".context.eigenlayer.l2.bn254_certificate_verifier")
        );
        context.ecdsaCertificateVerifier = IECDSACertificateVerifier(
            stdJson.readAddress(_context, ".context.eigenlayer.l2.ecdsa_certificate_verifier")
        );
        context.taskMailbox = ITaskMailbox(_readHourglassConfigAddress(environment, "taskMailbox"));
        context.taskHook = IAVSTaskHook(_readAVSL2ConfigAddress(environment, "avsTaskHook"));

        return context;
    }

    function _readHourglassConfigAddress(
        string memory environment,
        string memory key
    ) internal view returns (address) {
        // Load the Hourglass config file
        string memory hourglassConfigFile =
            string.concat("script/", environment, "/output/deploy_hourglass_core_output.json");
        string memory hourglassConfig = vm.readFile(hourglassConfigFile);

        // Parse and return the address
        return stdJson.readAddress(hourglassConfig, string.concat(".addresses.", key));
    }

    function _readAVSL2ConfigAddress(string memory environment, string memory key) internal view returns (address) {
        // Load the AVS L2 config file
        string memory avsL2ConfigFile = string.concat("script/", environment, "/output/deploy_avs_l2_output.json");
        string memory avsL2Config = vm.readFile(avsL2ConfigFile);

        // Parse and return the address
        return stdJson.readAddress(avsL2Config, string.concat(".addresses.", key));
    }
}
