package keyvault

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

        // CertificatePolicyAction enumerates the values for certificate policy action.
    type CertificatePolicyAction string

    const (
                // CertificatePolicyActionAutoRenew ...
        CertificatePolicyActionAutoRenew CertificatePolicyAction = "AutoRenew"
                // CertificatePolicyActionEmailContacts ...
        CertificatePolicyActionEmailContacts CertificatePolicyAction = "EmailContacts"
            )
    // PossibleCertificatePolicyActionValues returns an array of possible values for the CertificatePolicyAction const type.
    func PossibleCertificatePolicyActionValues() []CertificatePolicyAction {
        return []CertificatePolicyAction{CertificatePolicyActionAutoRenew,CertificatePolicyActionEmailContacts}
    }

        // DataAction enumerates the values for data action.
    type DataAction string

    const (
            // DataActionBackupHsmKeys Backup HSM keys.
        DataActionBackupHsmKeys DataAction = "Microsoft.KeyVault/managedHsm/keys/backup/action"
            // DataActionCreateHsmKey Create an HSM key.
        DataActionCreateHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/create"
            // DataActionDecryptHsmKey Decrypt using an HSM key.
        DataActionDecryptHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/decrypt/action"
            // DataActionDeleteHsmKey Delete an HSM key.
        DataActionDeleteHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/delete"
            // DataActionDeleteRoleAssignment Delete role assignment.
        DataActionDeleteRoleAssignment DataAction = "Microsoft.KeyVault/managedHsm/roleAssignments/delete/action"
            // DataActionDeleteRoleDefinition Delete role definition.
        DataActionDeleteRoleDefinition DataAction = "Microsoft.KeyVault/managedHsm/roleDefinitions/delete/action"
            // DataActionDownloadHsmSecurityDomain Download an HSM security domain.
        DataActionDownloadHsmSecurityDomain DataAction = "Microsoft.KeyVault/managedHsm/securitydomain/download/action"
            // DataActionDownloadHsmSecurityDomainStatus Check status of HSM security domain download.
        DataActionDownloadHsmSecurityDomainStatus DataAction = "Microsoft.KeyVault/managedHsm/securitydomain/download/read"
            // DataActionEncryptHsmKey Encrypt using an HSM key.
        DataActionEncryptHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/encrypt/action"
            // DataActionExportHsmKey Export an HSM key.
        DataActionExportHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/export/action"
            // DataActionGetRoleAssignment Get role assignment.
        DataActionGetRoleAssignment DataAction = "Microsoft.KeyVault/managedHsm/roleAssignments/read/action"
            // DataActionImportHsmKey Import an HSM key.
        DataActionImportHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/import/action"
            // DataActionPurgeDeletedHsmKey Purge a deleted HSM key.
        DataActionPurgeDeletedHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/deletedKeys/delete"
            // DataActionRandomNumbersGenerate Generate random numbers.
        DataActionRandomNumbersGenerate DataAction = "Microsoft.KeyVault/managedHsm/rng/action"
            // DataActionReadDeletedHsmKey Read deleted HSM key.
        DataActionReadDeletedHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/deletedKeys/read/action"
            // DataActionReadHsmBackupStatus Read an HSM backup status.
        DataActionReadHsmBackupStatus DataAction = "Microsoft.KeyVault/managedHsm/backup/status/action"
            // DataActionReadHsmKey Read HSM key metadata.
        DataActionReadHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/read/action"
            // DataActionReadHsmRestoreStatus Read an HSM restore status.
        DataActionReadHsmRestoreStatus DataAction = "Microsoft.KeyVault/managedHsm/restore/status/action"
            // DataActionReadHsmSecurityDomainStatus Check the status of the HSM security domain exchange file.
        DataActionReadHsmSecurityDomainStatus DataAction = "Microsoft.KeyVault/managedHsm/securitydomain/upload/read"
            // DataActionReadHsmSecurityDomainTransferKey Download an HSM security domain transfer key.
        DataActionReadHsmSecurityDomainTransferKey DataAction = "Microsoft.KeyVault/managedHsm/securitydomain/transferkey/read"
            // DataActionReadRoleDefinition Get role definition.
        DataActionReadRoleDefinition DataAction = "Microsoft.KeyVault/managedHsm/roleDefinitions/read/action"
            // DataActionRecoverDeletedHsmKey Recover deleted HSM key.
        DataActionRecoverDeletedHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/deletedKeys/recover/action"
            // DataActionReleaseKey Release an HSM key using Secure Key Release.
        DataActionReleaseKey DataAction = "Microsoft.KeyVault/managedHsm/keys/release/action"
            // DataActionRestoreHsmKeys Restore HSM keys.
        DataActionRestoreHsmKeys DataAction = "Microsoft.KeyVault/managedHsm/keys/restore/action"
            // DataActionSignHsmKey Sign using an HSM key.
        DataActionSignHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/sign/action"
            // DataActionStartHsmBackup Start an HSM backup.
        DataActionStartHsmBackup DataAction = "Microsoft.KeyVault/managedHsm/backup/start/action"
            // DataActionStartHsmRestore Start an HSM restore.
        DataActionStartHsmRestore DataAction = "Microsoft.KeyVault/managedHsm/restore/start/action"
            // DataActionUnwrapHsmKey Unwrap using an HSM key.
        DataActionUnwrapHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/unwrap/action"
            // DataActionUploadHsmSecurityDomain Upload an HSM security domain.
        DataActionUploadHsmSecurityDomain DataAction = "Microsoft.KeyVault/managedHsm/securitydomain/upload/action"
            // DataActionVerifyHsmKey Verify using an HSM key.
        DataActionVerifyHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/verify/action"
            // DataActionWrapHsmKey Wrap using an HSM key.
        DataActionWrapHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/wrap/action"
            // DataActionWriteHsmKey Update an HSM key.
        DataActionWriteHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/write/action"
            // DataActionWriteRoleAssignment Create or update role assignment.
        DataActionWriteRoleAssignment DataAction = "Microsoft.KeyVault/managedHsm/roleAssignments/write/action"
            // DataActionWriteRoleDefinition Create or update role definition.
        DataActionWriteRoleDefinition DataAction = "Microsoft.KeyVault/managedHsm/roleDefinitions/write/action"
            )
    // PossibleDataActionValues returns an array of possible values for the DataAction const type.
    func PossibleDataActionValues() []DataAction {
        return []DataAction{DataActionBackupHsmKeys,DataActionCreateHsmKey,DataActionDecryptHsmKey,DataActionDeleteHsmKey,DataActionDeleteRoleAssignment,DataActionDeleteRoleDefinition,DataActionDownloadHsmSecurityDomain,DataActionDownloadHsmSecurityDomainStatus,DataActionEncryptHsmKey,DataActionExportHsmKey,DataActionGetRoleAssignment,DataActionImportHsmKey,DataActionPurgeDeletedHsmKey,DataActionRandomNumbersGenerate,DataActionReadDeletedHsmKey,DataActionReadHsmBackupStatus,DataActionReadHsmKey,DataActionReadHsmRestoreStatus,DataActionReadHsmSecurityDomainStatus,DataActionReadHsmSecurityDomainTransferKey,DataActionReadRoleDefinition,DataActionRecoverDeletedHsmKey,DataActionReleaseKey,DataActionRestoreHsmKeys,DataActionSignHsmKey,DataActionStartHsmBackup,DataActionStartHsmRestore,DataActionUnwrapHsmKey,DataActionUploadHsmSecurityDomain,DataActionVerifyHsmKey,DataActionWrapHsmKey,DataActionWriteHsmKey,DataActionWriteRoleAssignment,DataActionWriteRoleDefinition}
    }

        // DeletionRecoveryLevel enumerates the values for deletion recovery level.
    type DeletionRecoveryLevel string

    const (
            // DeletionRecoveryLevelCustomizedRecoverable Denotes a vault state in which deletion is recoverable
            // without the possibility for immediate and permanent deletion (i.e. purge when 7<=
            // SoftDeleteRetentionInDays < 90).This level guarantees the recoverability of the deleted entity during
            // the retention interval and while the subscription is still available.
        DeletionRecoveryLevelCustomizedRecoverable DeletionRecoveryLevel = "CustomizedRecoverable"
            // DeletionRecoveryLevelCustomizedRecoverableProtectedSubscription Denotes a vault and subscription state
            // in which deletion is recoverable, immediate and permanent deletion (i.e. purge) is not permitted, and in
            // which the subscription itself cannot be permanently canceled when 7<= SoftDeleteRetentionInDays < 90.
            // This level guarantees the recoverability of the deleted entity during the retention interval, and also
            // reflects the fact that the subscription itself cannot be cancelled.
        DeletionRecoveryLevelCustomizedRecoverableProtectedSubscription DeletionRecoveryLevel = "CustomizedRecoverable+ProtectedSubscription"
            // DeletionRecoveryLevelCustomizedRecoverablePurgeable Denotes a vault state in which deletion is
            // recoverable, and which also permits immediate and permanent deletion (i.e. purge when 7<=
            // SoftDeleteRetentionInDays < 90). This level guarantees the recoverability of the deleted entity during
            // the retention interval, unless a Purge operation is requested, or the subscription is cancelled.
        DeletionRecoveryLevelCustomizedRecoverablePurgeable DeletionRecoveryLevel = "CustomizedRecoverable+Purgeable"
            // DeletionRecoveryLevelPurgeable Denotes a vault state in which deletion is an irreversible operation,
            // without the possibility for recovery. This level corresponds to no protection being available against a
            // Delete operation; the data is irretrievably lost upon accepting a Delete operation at the entity level
            // or higher (vault, resource group, subscription etc.)
        DeletionRecoveryLevelPurgeable DeletionRecoveryLevel = "Purgeable"
            // DeletionRecoveryLevelRecoverable Denotes a vault state in which deletion is recoverable without the
            // possibility for immediate and permanent deletion (i.e. purge). This level guarantees the recoverability
            // of the deleted entity during the retention interval(90 days) and while the subscription is still
            // available. System wil permanently delete it after 90 days, if not recovered
        DeletionRecoveryLevelRecoverable DeletionRecoveryLevel = "Recoverable"
            // DeletionRecoveryLevelRecoverableProtectedSubscription Denotes a vault and subscription state in which
            // deletion is recoverable within retention interval (90 days), immediate and permanent deletion (i.e.
            // purge) is not permitted, and in which the subscription itself  cannot be permanently canceled. System
            // wil permanently delete it after 90 days, if not recovered
        DeletionRecoveryLevelRecoverableProtectedSubscription DeletionRecoveryLevel = "Recoverable+ProtectedSubscription"
            // DeletionRecoveryLevelRecoverablePurgeable Denotes a vault state in which deletion is recoverable, and
            // which also permits immediate and permanent deletion (i.e. purge). This level guarantees the
            // recoverability of the deleted entity during the retention interval (90 days), unless a Purge operation
            // is requested, or the subscription is cancelled. System wil permanently delete it after 90 days, if not
            // recovered
        DeletionRecoveryLevelRecoverablePurgeable DeletionRecoveryLevel = "Recoverable+Purgeable"
            )
    // PossibleDeletionRecoveryLevelValues returns an array of possible values for the DeletionRecoveryLevel const type.
    func PossibleDeletionRecoveryLevelValues() []DeletionRecoveryLevel {
        return []DeletionRecoveryLevel{DeletionRecoveryLevelCustomizedRecoverable,DeletionRecoveryLevelCustomizedRecoverableProtectedSubscription,DeletionRecoveryLevelCustomizedRecoverablePurgeable,DeletionRecoveryLevelPurgeable,DeletionRecoveryLevelRecoverable,DeletionRecoveryLevelRecoverableProtectedSubscription,DeletionRecoveryLevelRecoverablePurgeable}
    }

        // JSONWebKeyCurveName enumerates the values for json web key curve name.
    type JSONWebKeyCurveName string

    const (
                // JSONWebKeyCurveNameP256 ...
        JSONWebKeyCurveNameP256 JSONWebKeyCurveName = "P-256"
                // JSONWebKeyCurveNameP256K ...
        JSONWebKeyCurveNameP256K JSONWebKeyCurveName = "P-256K"
                // JSONWebKeyCurveNameP384 ...
        JSONWebKeyCurveNameP384 JSONWebKeyCurveName = "P-384"
                // JSONWebKeyCurveNameP521 ...
        JSONWebKeyCurveNameP521 JSONWebKeyCurveName = "P-521"
            )
    // PossibleJSONWebKeyCurveNameValues returns an array of possible values for the JSONWebKeyCurveName const type.
    func PossibleJSONWebKeyCurveNameValues() []JSONWebKeyCurveName {
        return []JSONWebKeyCurveName{JSONWebKeyCurveNameP256,JSONWebKeyCurveNameP256K,JSONWebKeyCurveNameP384,JSONWebKeyCurveNameP521}
    }

        // JSONWebKeyEncryptionAlgorithm enumerates the values for json web key encryption algorithm.
    type JSONWebKeyEncryptionAlgorithm string

    const (
                // JSONWebKeyEncryptionAlgorithmA128CBC ...
        JSONWebKeyEncryptionAlgorithmA128CBC JSONWebKeyEncryptionAlgorithm = "A128CBC"
                // JSONWebKeyEncryptionAlgorithmA128CBCPAD ...
        JSONWebKeyEncryptionAlgorithmA128CBCPAD JSONWebKeyEncryptionAlgorithm = "A128CBCPAD"
                // JSONWebKeyEncryptionAlgorithmA128GCM ...
        JSONWebKeyEncryptionAlgorithmA128GCM JSONWebKeyEncryptionAlgorithm = "A128GCM"
                // JSONWebKeyEncryptionAlgorithmA128KW ...
        JSONWebKeyEncryptionAlgorithmA128KW JSONWebKeyEncryptionAlgorithm = "A128KW"
                // JSONWebKeyEncryptionAlgorithmA192CBC ...
        JSONWebKeyEncryptionAlgorithmA192CBC JSONWebKeyEncryptionAlgorithm = "A192CBC"
                // JSONWebKeyEncryptionAlgorithmA192CBCPAD ...
        JSONWebKeyEncryptionAlgorithmA192CBCPAD JSONWebKeyEncryptionAlgorithm = "A192CBCPAD"
                // JSONWebKeyEncryptionAlgorithmA192GCM ...
        JSONWebKeyEncryptionAlgorithmA192GCM JSONWebKeyEncryptionAlgorithm = "A192GCM"
                // JSONWebKeyEncryptionAlgorithmA192KW ...
        JSONWebKeyEncryptionAlgorithmA192KW JSONWebKeyEncryptionAlgorithm = "A192KW"
                // JSONWebKeyEncryptionAlgorithmA256CBC ...
        JSONWebKeyEncryptionAlgorithmA256CBC JSONWebKeyEncryptionAlgorithm = "A256CBC"
                // JSONWebKeyEncryptionAlgorithmA256CBCPAD ...
        JSONWebKeyEncryptionAlgorithmA256CBCPAD JSONWebKeyEncryptionAlgorithm = "A256CBCPAD"
                // JSONWebKeyEncryptionAlgorithmA256GCM ...
        JSONWebKeyEncryptionAlgorithmA256GCM JSONWebKeyEncryptionAlgorithm = "A256GCM"
                // JSONWebKeyEncryptionAlgorithmA256KW ...
        JSONWebKeyEncryptionAlgorithmA256KW JSONWebKeyEncryptionAlgorithm = "A256KW"
                // JSONWebKeyEncryptionAlgorithmRSA15 ...
        JSONWebKeyEncryptionAlgorithmRSA15 JSONWebKeyEncryptionAlgorithm = "RSA1_5"
                // JSONWebKeyEncryptionAlgorithmRSAOAEP ...
        JSONWebKeyEncryptionAlgorithmRSAOAEP JSONWebKeyEncryptionAlgorithm = "RSA-OAEP"
                // JSONWebKeyEncryptionAlgorithmRSAOAEP256 ...
        JSONWebKeyEncryptionAlgorithmRSAOAEP256 JSONWebKeyEncryptionAlgorithm = "RSA-OAEP-256"
            )
    // PossibleJSONWebKeyEncryptionAlgorithmValues returns an array of possible values for the JSONWebKeyEncryptionAlgorithm const type.
    func PossibleJSONWebKeyEncryptionAlgorithmValues() []JSONWebKeyEncryptionAlgorithm {
        return []JSONWebKeyEncryptionAlgorithm{JSONWebKeyEncryptionAlgorithmA128CBC,JSONWebKeyEncryptionAlgorithmA128CBCPAD,JSONWebKeyEncryptionAlgorithmA128GCM,JSONWebKeyEncryptionAlgorithmA128KW,JSONWebKeyEncryptionAlgorithmA192CBC,JSONWebKeyEncryptionAlgorithmA192CBCPAD,JSONWebKeyEncryptionAlgorithmA192GCM,JSONWebKeyEncryptionAlgorithmA192KW,JSONWebKeyEncryptionAlgorithmA256CBC,JSONWebKeyEncryptionAlgorithmA256CBCPAD,JSONWebKeyEncryptionAlgorithmA256GCM,JSONWebKeyEncryptionAlgorithmA256KW,JSONWebKeyEncryptionAlgorithmRSA15,JSONWebKeyEncryptionAlgorithmRSAOAEP,JSONWebKeyEncryptionAlgorithmRSAOAEP256}
    }

        // JSONWebKeyOperation enumerates the values for json web key operation.
    type JSONWebKeyOperation string

    const (
                // JSONWebKeyOperationDecrypt ...
        JSONWebKeyOperationDecrypt JSONWebKeyOperation = "decrypt"
                // JSONWebKeyOperationEncrypt ...
        JSONWebKeyOperationEncrypt JSONWebKeyOperation = "encrypt"
                // JSONWebKeyOperationExport ...
        JSONWebKeyOperationExport JSONWebKeyOperation = "export"
                // JSONWebKeyOperationImport ...
        JSONWebKeyOperationImport JSONWebKeyOperation = "import"
                // JSONWebKeyOperationSign ...
        JSONWebKeyOperationSign JSONWebKeyOperation = "sign"
                // JSONWebKeyOperationUnwrapKey ...
        JSONWebKeyOperationUnwrapKey JSONWebKeyOperation = "unwrapKey"
                // JSONWebKeyOperationVerify ...
        JSONWebKeyOperationVerify JSONWebKeyOperation = "verify"
                // JSONWebKeyOperationWrapKey ...
        JSONWebKeyOperationWrapKey JSONWebKeyOperation = "wrapKey"
            )
    // PossibleJSONWebKeyOperationValues returns an array of possible values for the JSONWebKeyOperation const type.
    func PossibleJSONWebKeyOperationValues() []JSONWebKeyOperation {
        return []JSONWebKeyOperation{JSONWebKeyOperationDecrypt,JSONWebKeyOperationEncrypt,JSONWebKeyOperationExport,JSONWebKeyOperationImport,JSONWebKeyOperationSign,JSONWebKeyOperationUnwrapKey,JSONWebKeyOperationVerify,JSONWebKeyOperationWrapKey}
    }

        // JSONWebKeySignatureAlgorithm enumerates the values for json web key signature algorithm.
    type JSONWebKeySignatureAlgorithm string

    const (
            // JSONWebKeySignatureAlgorithmES256 ECDSA using P-256 and SHA-256, as described in
            // https://tools.ietf.org/html/rfc7518.
        JSONWebKeySignatureAlgorithmES256 JSONWebKeySignatureAlgorithm = "ES256"
            // JSONWebKeySignatureAlgorithmES256K ECDSA using P-256K and SHA-256, as described in
            // https://tools.ietf.org/html/rfc7518
        JSONWebKeySignatureAlgorithmES256K JSONWebKeySignatureAlgorithm = "ES256K"
            // JSONWebKeySignatureAlgorithmES384 ECDSA using P-384 and SHA-384, as described in
            // https://tools.ietf.org/html/rfc7518
        JSONWebKeySignatureAlgorithmES384 JSONWebKeySignatureAlgorithm = "ES384"
            // JSONWebKeySignatureAlgorithmES512 ECDSA using P-521 and SHA-512, as described in
            // https://tools.ietf.org/html/rfc7518
        JSONWebKeySignatureAlgorithmES512 JSONWebKeySignatureAlgorithm = "ES512"
            // JSONWebKeySignatureAlgorithmPS256 RSASSA-PSS using SHA-256 and MGF1 with SHA-256, as described in
            // https://tools.ietf.org/html/rfc7518
        JSONWebKeySignatureAlgorithmPS256 JSONWebKeySignatureAlgorithm = "PS256"
            // JSONWebKeySignatureAlgorithmPS384 RSASSA-PSS using SHA-384 and MGF1 with SHA-384, as described in
            // https://tools.ietf.org/html/rfc7518
        JSONWebKeySignatureAlgorithmPS384 JSONWebKeySignatureAlgorithm = "PS384"
            // JSONWebKeySignatureAlgorithmPS512 RSASSA-PSS using SHA-512 and MGF1 with SHA-512, as described in
            // https://tools.ietf.org/html/rfc7518
        JSONWebKeySignatureAlgorithmPS512 JSONWebKeySignatureAlgorithm = "PS512"
            // JSONWebKeySignatureAlgorithmRS256 RSASSA-PKCS1-v1_5 using SHA-256, as described in
            // https://tools.ietf.org/html/rfc7518
        JSONWebKeySignatureAlgorithmRS256 JSONWebKeySignatureAlgorithm = "RS256"
            // JSONWebKeySignatureAlgorithmRS384 RSASSA-PKCS1-v1_5 using SHA-384, as described in
            // https://tools.ietf.org/html/rfc7518
        JSONWebKeySignatureAlgorithmRS384 JSONWebKeySignatureAlgorithm = "RS384"
            // JSONWebKeySignatureAlgorithmRS512 RSASSA-PKCS1-v1_5 using SHA-512, as described in
            // https://tools.ietf.org/html/rfc7518
        JSONWebKeySignatureAlgorithmRS512 JSONWebKeySignatureAlgorithm = "RS512"
            // JSONWebKeySignatureAlgorithmRSNULL Reserved
        JSONWebKeySignatureAlgorithmRSNULL JSONWebKeySignatureAlgorithm = "RSNULL"
            )
    // PossibleJSONWebKeySignatureAlgorithmValues returns an array of possible values for the JSONWebKeySignatureAlgorithm const type.
    func PossibleJSONWebKeySignatureAlgorithmValues() []JSONWebKeySignatureAlgorithm {
        return []JSONWebKeySignatureAlgorithm{JSONWebKeySignatureAlgorithmES256,JSONWebKeySignatureAlgorithmES256K,JSONWebKeySignatureAlgorithmES384,JSONWebKeySignatureAlgorithmES512,JSONWebKeySignatureAlgorithmPS256,JSONWebKeySignatureAlgorithmPS384,JSONWebKeySignatureAlgorithmPS512,JSONWebKeySignatureAlgorithmRS256,JSONWebKeySignatureAlgorithmRS384,JSONWebKeySignatureAlgorithmRS512,JSONWebKeySignatureAlgorithmRSNULL}
    }

        // JSONWebKeyType enumerates the values for json web key type.
    type JSONWebKeyType string

    const (
                // JSONWebKeyTypeEC ...
        JSONWebKeyTypeEC JSONWebKeyType = "EC"
                // JSONWebKeyTypeECHSM ...
        JSONWebKeyTypeECHSM JSONWebKeyType = "EC-HSM"
                // JSONWebKeyTypeOct ...
        JSONWebKeyTypeOct JSONWebKeyType = "oct"
                // JSONWebKeyTypeOctHSM ...
        JSONWebKeyTypeOctHSM JSONWebKeyType = "oct-HSM"
                // JSONWebKeyTypeRSA ...
        JSONWebKeyTypeRSA JSONWebKeyType = "RSA"
                // JSONWebKeyTypeRSAHSM ...
        JSONWebKeyTypeRSAHSM JSONWebKeyType = "RSA-HSM"
            )
    // PossibleJSONWebKeyTypeValues returns an array of possible values for the JSONWebKeyType const type.
    func PossibleJSONWebKeyTypeValues() []JSONWebKeyType {
        return []JSONWebKeyType{JSONWebKeyTypeEC,JSONWebKeyTypeECHSM,JSONWebKeyTypeOct,JSONWebKeyTypeOctHSM,JSONWebKeyTypeRSA,JSONWebKeyTypeRSAHSM}
    }

        // KeyEncryptionAlgorithm enumerates the values for key encryption algorithm.
    type KeyEncryptionAlgorithm string

    const (
                // KeyEncryptionAlgorithmCKMRSAAESKEYWRAP ...
        KeyEncryptionAlgorithmCKMRSAAESKEYWRAP KeyEncryptionAlgorithm = "CKM_RSA_AES_KEY_WRAP"
                // KeyEncryptionAlgorithmRSAAESKEYWRAP256 ...
        KeyEncryptionAlgorithmRSAAESKEYWRAP256 KeyEncryptionAlgorithm = "RSA_AES_KEY_WRAP_256"
                // KeyEncryptionAlgorithmRSAAESKEYWRAP384 ...
        KeyEncryptionAlgorithmRSAAESKEYWRAP384 KeyEncryptionAlgorithm = "RSA_AES_KEY_WRAP_384"
            )
    // PossibleKeyEncryptionAlgorithmValues returns an array of possible values for the KeyEncryptionAlgorithm const type.
    func PossibleKeyEncryptionAlgorithmValues() []KeyEncryptionAlgorithm {
        return []KeyEncryptionAlgorithm{KeyEncryptionAlgorithmCKMRSAAESKEYWRAP,KeyEncryptionAlgorithmRSAAESKEYWRAP256,KeyEncryptionAlgorithmRSAAESKEYWRAP384}
    }

        // KeyRotationPolicyAction enumerates the values for key rotation policy action.
    type KeyRotationPolicyAction string

    const (
            // KeyRotationPolicyActionNotify Trigger event grid events. For preview, the notification time is not
            // configurable and it is default to 30 days before expiry.
        KeyRotationPolicyActionNotify KeyRotationPolicyAction = "notify"
            // KeyRotationPolicyActionRotate Rotate the key based on the key policy.
        KeyRotationPolicyActionRotate KeyRotationPolicyAction = "rotate"
            )
    // PossibleKeyRotationPolicyActionValues returns an array of possible values for the KeyRotationPolicyAction const type.
    func PossibleKeyRotationPolicyActionValues() []KeyRotationPolicyAction {
        return []KeyRotationPolicyAction{KeyRotationPolicyActionNotify,KeyRotationPolicyActionRotate}
    }

        // KeyUsageType enumerates the values for key usage type.
    type KeyUsageType string

    const (
                // KeyUsageTypeCRLSign ...
        KeyUsageTypeCRLSign KeyUsageType = "cRLSign"
                // KeyUsageTypeDataEncipherment ...
        KeyUsageTypeDataEncipherment KeyUsageType = "dataEncipherment"
                // KeyUsageTypeDecipherOnly ...
        KeyUsageTypeDecipherOnly KeyUsageType = "decipherOnly"
                // KeyUsageTypeDigitalSignature ...
        KeyUsageTypeDigitalSignature KeyUsageType = "digitalSignature"
                // KeyUsageTypeEncipherOnly ...
        KeyUsageTypeEncipherOnly KeyUsageType = "encipherOnly"
                // KeyUsageTypeKeyAgreement ...
        KeyUsageTypeKeyAgreement KeyUsageType = "keyAgreement"
                // KeyUsageTypeKeyCertSign ...
        KeyUsageTypeKeyCertSign KeyUsageType = "keyCertSign"
                // KeyUsageTypeKeyEncipherment ...
        KeyUsageTypeKeyEncipherment KeyUsageType = "keyEncipherment"
                // KeyUsageTypeNonRepudiation ...
        KeyUsageTypeNonRepudiation KeyUsageType = "nonRepudiation"
            )
    // PossibleKeyUsageTypeValues returns an array of possible values for the KeyUsageType const type.
    func PossibleKeyUsageTypeValues() []KeyUsageType {
        return []KeyUsageType{KeyUsageTypeCRLSign,KeyUsageTypeDataEncipherment,KeyUsageTypeDecipherOnly,KeyUsageTypeDigitalSignature,KeyUsageTypeEncipherOnly,KeyUsageTypeKeyAgreement,KeyUsageTypeKeyCertSign,KeyUsageTypeKeyEncipherment,KeyUsageTypeNonRepudiation}
    }

        // OperationStatus enumerates the values for operation status.
    type OperationStatus string

    const (
                // OperationStatusFailed ...
        OperationStatusFailed OperationStatus = "Failed"
                // OperationStatusInProgress ...
        OperationStatusInProgress OperationStatus = "InProgress"
                // OperationStatusSuccess ...
        OperationStatusSuccess OperationStatus = "Success"
            )
    // PossibleOperationStatusValues returns an array of possible values for the OperationStatus const type.
    func PossibleOperationStatusValues() []OperationStatus {
        return []OperationStatus{OperationStatusFailed,OperationStatusInProgress,OperationStatusSuccess}
    }

        // RoleDefinitionType enumerates the values for role definition type.
    type RoleDefinitionType string

    const (
                // RoleDefinitionTypeMicrosoftAuthorizationroleDefinitions ...
        RoleDefinitionTypeMicrosoftAuthorizationroleDefinitions RoleDefinitionType = "Microsoft.Authorization/roleDefinitions"
            )
    // PossibleRoleDefinitionTypeValues returns an array of possible values for the RoleDefinitionType const type.
    func PossibleRoleDefinitionTypeValues() []RoleDefinitionType {
        return []RoleDefinitionType{RoleDefinitionTypeMicrosoftAuthorizationroleDefinitions}
    }

        // RoleScope enumerates the values for role scope.
    type RoleScope string

    const (
            // RoleScopeGlobal Global scope
        RoleScopeGlobal RoleScope = "/"
            // RoleScopeKeys Keys scope
        RoleScopeKeys RoleScope = "/keys"
            )
    // PossibleRoleScopeValues returns an array of possible values for the RoleScope const type.
    func PossibleRoleScopeValues() []RoleScope {
        return []RoleScope{RoleScopeGlobal,RoleScopeKeys}
    }

        // RoleType enumerates the values for role type.
    type RoleType string

    const (
            // RoleTypeBuiltInRole Built in role.
        RoleTypeBuiltInRole RoleType = "AKVBuiltInRole"
            // RoleTypeCustomRole Custom role.
        RoleTypeCustomRole RoleType = "CustomRole"
            )
    // PossibleRoleTypeValues returns an array of possible values for the RoleType const type.
    func PossibleRoleTypeValues() []RoleType {
        return []RoleType{RoleTypeBuiltInRole,RoleTypeCustomRole}
    }

        // SasTokenType enumerates the values for sas token type.
    type SasTokenType string

    const (
                // SasTokenTypeAccount ...
        SasTokenTypeAccount SasTokenType = "account"
                // SasTokenTypeService ...
        SasTokenTypeService SasTokenType = "service"
            )
    // PossibleSasTokenTypeValues returns an array of possible values for the SasTokenType const type.
    func PossibleSasTokenTypeValues() []SasTokenType {
        return []SasTokenType{SasTokenTypeAccount,SasTokenTypeService}
    }

        // SettingTypeEnum enumerates the values for setting type enum.
    type SettingTypeEnum string

    const (
                // SettingTypeEnumBoolean ...
        SettingTypeEnumBoolean SettingTypeEnum = "boolean"
            )
    // PossibleSettingTypeEnumValues returns an array of possible values for the SettingTypeEnum const type.
    func PossibleSettingTypeEnumValues() []SettingTypeEnum {
        return []SettingTypeEnum{SettingTypeEnumBoolean}
    }

