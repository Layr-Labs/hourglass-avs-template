pragma solidity ^0.8.27;

import {BN254} from "@eigenlayer-middleware/src/libraries/BN254.sol";

import {IBN254CertificateVerifier} from "@hourglass-monorepo/src/interfaces/avs/l2/IBN254CertificateVerifier.sol";

contract BN254CertificateVerifier is IBN254CertificateVerifier {
    function maxOperatorTableStaleness() external pure returns (uint32) {
        return 86_400;
    }

    function verifyCertificate(
        BN254Certificate memory /*cert*/
    ) external pure returns (uint96[] memory signedStakes) {
        return new uint96[](0);
    }

    function verifyCertificateProportion(
        BN254Certificate memory, /*cert*/
        uint16[] memory /*totalStakeProportionThresholds*/
    ) external pure returns (bool) {
        return true;
    }

    function verifyCertificateNominal(
        BN254Certificate memory, /*cert*/
        uint96[] memory /*totalStakeNominalThresholds*/
    ) external pure returns (bool) {
        return true;
    }
}
