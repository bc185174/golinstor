// This file was autogenerated by genconsts.py

// LINSTOR - management of distributed storage/DRBD9 resources
// Copyright (C) 2017 - 2021  LINBIT HA-Solutions GmbH
// All Rights Reserved.
// Author: Robert Altnoeder, Roland Kammerer, Gabor Hernadi, Rene Peinthor
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package linstor

// ## Bits 62 - 63 (most significant 2) are reserved for the message type masks (error, warning, info)
// Bits 25 - 26 are reserved for the operation type masks (create, modify, delete)
// Bits 18 - 22 are reserved for the object type masks (node, resource, resource definition, ...)
// Bits 0  - 14 are reserved for codes ###
const MaskBitsType = 0xC000000000000000
const MaskError = 0xC000000000000000
const MaskWarn = 0x8000000000000000
const MaskInfo = 0x4000000000000000
const MaskSuccess = 0x0000000000000000

// ## Operation type masks ###
const MaskBitsOp = 0x0000000003000000
const MaskCrt = 0x0000000001000000
const MaskMod = 0x0000000002000000
const MaskDel = 0x0000000003000000

// ## Type masks (Node, ResDfn, Res, VolDfn, Vol, NetInterface, ...) ###
const MaskBitsObj = 0x00000000007C0000
const MaskExtFiles = 0x0000000000500000
const MaskPhysicalDevice = 0x00000000004C0000
const MaskVlmGrp = 0x0000000000480000
const MaskRscGrp = 0x0000000000440000
const MaskKvs = 0x0000000000400000
const MaskNode = 0x00000000003C0000
const MaskRscDfn = 0x0000000000380000
const MaskRsc = 0x0000000000340000
const MaskVlmDfn = 0x0000000000300000
const MaskVlm = 0x00000000002C0000
const MaskNodeConn = 0x0000000000280000
const MaskRscConn = 0x0000000000240000
const MaskVlmConn = 0x0000000000200000
const MaskNetIf = 0x00000000001C0000
const MaskStorPoolDfn = 0x0000000000180000
const MaskStorPool = 0x0000000000140000
const MaskCtrlConf = 0x0000000000100000
const MaskSnapshot = 0x00000000000C0000

// ## Codes ###
const MaskBitsCode = 0x0000000000007FFF

// ## Codes 1-9: success ###
const Created = (1 | MaskSuccess)
const Deleted = (2 | MaskSuccess)
const Modified = (3 | MaskSuccess)
const PassphraseAccepted = (4 | MaskSuccess)

// ## Codes 100 - 999: failures ###

// ## Codes 100 - 199: sql failures ###
const FailSql = (100 | MaskError)
const FailSqlRollback = (101 | MaskError)

// ## Codes 200-299: invalid * failures ###
const FailInvldNodeName = (200 | MaskError)
const FailInvldNodeType = (201 | MaskError)
const FailInvldRscName = (202 | MaskError)
const FailInvldRscPort = (203 | MaskError)
const FailInvldNodeId = (204 | MaskError)
const FailInvldVlmNr = (205 | MaskError)
const FailInvldVlmSize = (206 | MaskError)
const FailInvldMinorNr = (207 | MaskError)
const FailInvldStorPoolName = (208 | MaskError)
const FailInvldNetName = (209 | MaskError)
const FailInvldNetAddr = (210 | MaskError)
const FailInvldNetPort = (211 | MaskError)
const FailInvldNetType = (212 | MaskError)
const FailInvldProp = (213 | MaskError)
const FailInvldTransportType = (214 | MaskError)
const FailInvldTcpPort = (215 | MaskError)
const FailInvldCryptPassphrase = (216 | MaskError)
const FailInvldEncryptType = (217 | MaskError)
const FailInvldSnapshotName = (218 | MaskError)
const FailInvldPlaceCount = (219 | MaskError)
const FailInvldFreeSpaceMgrName = (220 | MaskError)
const FailInvldStorDriver = (221 | MaskError)
const FailInvldDrbdProxyCompressionType = (222 | MaskError)
const FailInvldKvsName = (223 | MaskError)
const FailInvldLayerKind = (224 | MaskError)
const FailInvldLayerStack = (225 | MaskError)
const FailInvldExtName = (226 | MaskError)
const FailInvldProvider = (227 | MaskError)
const FailInvldVlmSizes = (228 | MaskError)
const FailInvldVlmCount = (229 | MaskError)
const FailInvldConf = (230 | MaskError)
const FailInvldSnapshotShippingSource = (231 | MaskError)
const FailInvldSnapshotShippingTarget = (232 | MaskError)
const FailNodeHasUsedRsc = (233 | MaskError)
const FailInvldRscGrpName = (234 | MaskError)
const FailInvldExtFileName = (235 | MaskError)
const FailInvldExtFile = (236 | MaskError)

// ## Codes 300-399: dependency not found failures ###
const FailNotFoundNode = (300 | MaskError)
const FailNotFoundRscDfn = (301 | MaskError)
const FailNotFoundRsc = (302 | MaskError)
const FailNotFoundVlmDfn = (303 | MaskError)
const FailNotFoundVlm = (304 | MaskError)
const FailNotFoundNetIf = (305 | MaskError)
const FailNotFoundNodeConn = (306 | MaskError)
const FailNotFoundRscConn = (307 | MaskError)
const FailNotFoundVlmConn = (308 | MaskError)
const FailNotFoundStorPoolDfn = (309 | MaskError)
const FailNotFoundStorPool = (310 | MaskError)
const FailNotFoundDfltStorPool = (311 | MaskError)
const FailNotFoundCryptKey = (312 | MaskError)
const FailNotFoundSnapshotDfn = (313 | MaskError)
const FailNotFoundSnapshotVlmDfn = (314 | MaskError)
const FailNotFoundSnapshot = (315 | MaskError)
const FailNotFoundKvs = (316 | MaskError)
const FailNotFoundRscGrp = (317 | MaskError)
const FailNotFoundVlmGrp = (318 | MaskError)
const FailNotFoundExosEnclosure = (319 | MaskError)
const FailNotFoundExtFile = (320 | MaskError)

// ## Codes 400-499: access denied failures ###
const FailAccDeniedNode = (400 | MaskError)
const FailAccDeniedRscDfn = (401 | MaskError)
const FailAccDeniedRsc = (402 | MaskError)
const FailAccDeniedVlmDfn = (403 | MaskError)
const FailAccDeniedVlm = (404 | MaskError)
const FailAccDeniedStorPoolDfn = (405 | MaskError)
const FailAccDeniedStorPool = (406 | MaskError)
const FailAccDeniedNodeConn = (407 | MaskError)
const FailAccDeniedRscConn = (408 | MaskError)
const FailAccDeniedVlmConn = (409 | MaskError)
const FailAccDeniedStltConn = (410 | MaskError)
const FailAccDeniedCtrlCfg = (411 | MaskError)
const FailAccDeniedCommand = (412 | MaskError)
const FailAccDeniedWatch = (413 | MaskError)
const FailAccDeniedSnapshotDfn = (414 | MaskError)
const FailAccDeniedSnapshot = (415 | MaskError)
const FailAccDeniedSnapshotVlmDfn = (416 | MaskError)
const FailAccDeniedFreeSpaceMgr = (417 | MaskError)
const FailAccDeniedKvs = (418 | MaskError)
const FailAccDeniedRscGrp = (419 | MaskError)
const FailAccDeniedVlmGrp = (420 | MaskError)
const FailAccDeniedSnapDfn = (421 | MaskError)
const FailAccDeniedExtFile = (422 | MaskError)

// ## Codes 500-599: data already exists failures ###
const FailExistsNode = (500 | MaskError)
const FailExistsRscDfn = (501 | MaskError)
const FailExistsRsc = (502 | MaskError)
const FailExistsVlmDfn = (503 | MaskError)
const FailExistsVlm = (504 | MaskError)
const FailExistsNetIf = (505 | MaskError)
const FailExistsNodeConn = (506 | MaskError)
const FailExistsRscConn = (507 | MaskError)
const FailExistsVlmConn = (508 | MaskError)
const FailExistsStorPoolDfn = (509 | MaskError)
const FailExistsStorPool = (510 | MaskError)
const FailExistsStltConn = (511 | MaskError)
const FailExistsCryptPassphrase = (512 | MaskError)
const FailExistsWatch = (513 | MaskError)
const FailExistsSnapshotDfn = (514 | MaskError)
const FailExistsSnapshot = (516 | MaskError)
const FailExistsExtName = (517 | MaskError)
const FailExistsNvmeTargetPerRscDfn = (518 | MaskError)
const FailExistsNvmeInitiatorPerRscDfn = (519 | MaskError)
const FailLostStorPool = (521 | MaskError)
const FailExistsRscGrp = (522 | MaskError)
const FailExistsVlmGrp = (523 | MaskError)
const FailExistsOpenflexTargetPerRscDfn = (524 | MaskError)
const FailExistsSnapshotShipping = (525 | MaskError)
const FailExistsExosEnclosure = (526 | MaskError)

// ## Codes 600-699: data missing failures ###
const FailMissingProps = (600 | MaskError)
const FailMissingPropsNetcomType = (601 | MaskError)
const FailMissingPropsNetcomPort = (602 | MaskError)
const FailMissingNetcom = (603 | MaskError)
const FailMissingPropsNetifName = (604 | MaskError)
const FailMissingStltConn = (605 | MaskError)
const FailMissingExtName = (606 | MaskError)
const FailMissingNvmeTarget = (608 | MaskError)
const FailNoStltConnDefined = (609 | MaskError)
const FailMissingOpenflexTarget = (610 | MaskError)

// ## Codes 700-799: uuid mismatch failures ###
const FailUuidNode = (700 | MaskError)
const FailUuidRscDfn = (701 | MaskError)
const FailUuidRsc = (702 | MaskError)
const FailUuidVlmDfn = (703 | MaskError)
const FailUuidVlm = (704 | MaskError)
const FailUuidNetIf = (705 | MaskError)
const FailUuidNodeConn = (706 | MaskError)
const FailUuidRscConn = (707 | MaskError)
const FailUuidVlmConn = (708 | MaskError)
const FailUuidStorPoolDfn = (709 | MaskError)
const FailUuidStorPool = (710 | MaskError)
const FailUuidKvs = (711 | MaskError)

// ## Codes 800-899: number pools exhausted ###
const FailPoolExhaustedVlmNr = (800 | MaskError)
const FailPoolExhaustedMinorNr = (801 | MaskError)
const FailPoolExhaustedTcpPort = (802 | MaskError)
const FailPoolExhaustedNodeId = (803 | MaskError)
const FailPoolExhaustedRscLayerId = (804 | MaskError)
const FailPoolExhaustedSpecialSatellteTcpPort = (805 | MaskError)
const FailPoolExhaustedSnapshotShippingTcpPort = (806 | MaskError)

// ## Other failures ###
const FailNotEnoughFreeSpace = (980 | MaskError)
const FailOnlyOneActRscPerSharedStorPoolAllowed = (981 | MaskError)
const FailCryptInit = (982 | MaskError)
const FailSnapshotShippingNotSupported = (983 | MaskError)
const FailSnapshotShippingInProgress = (984 | MaskError)
const FailUndecidableAutoplacment = (985 | MaskError)
const FailPreSelectScriptDidNotTerminate = (986 | MaskError)
const FailLinstorManagedSatelliteDidNotStartProperly = (987 | MaskError)
const FailStltDoesNotSupportLayer = (988 | MaskError)
const FailStltDoesNotSupportProvider = (989 | MaskError)
const FailStorPoolConfigurationError = (990 | MaskError)
const FailInsufficientReplicaCount = (991 | MaskError)
const FailRscBusy = (992 | MaskError)
const FailInsufficientPeerSlots = (993 | MaskError)
const FailSnapshotsNotSupported = (994 | MaskError)
const FailNotConnected = (995 | MaskError)
const FailNotEnoughNodes = (996 | MaskError)
const FailInUse = (997 | MaskError)
const FailUnknownError = (998 | MaskError)
const FailImplError = (999 | MaskError)

// ## Codes 1000-9999: warnings ###
const WarnInvldOptPropNetcomEnabled = (1001 | MaskWarn)
const WarnNotConnected = (1002 | MaskWarn)
const WarnStltNotUpdated = (1003 | MaskWarn)
const WarnNoStltConnDefined = (1004 | MaskWarn)
const WarnDelUnsetProp = (1005 | MaskWarn)
const WarnRscAlreadyDeployed = (1006 | MaskWarn)
const WarnRscAlreadyHasDisk = (1007 | MaskWarn)
const WarnRscAlreadyDiskless = (1008 | MaskWarn)
const WarnAllDiskless = (1009 | MaskWarn)
const WarnStorageError = (1010 | MaskWarn)
const WarnNotFoundCryptKey = (1011 | MaskWarn)
const WarnStorageKindAdded = (1012 | MaskWarn)
const WarnNotEnoughNodesForTieBreaker = (1013 | MaskWarn)
const WarnMixedPmemAndNonPmem = (1014 | MaskWarn)
const WarnUneffectiveProp = (1015 | MaskWarn)
const WarnInvldSnapshotShippingPrefix = (1016 | MaskWarn)
const WarnNodeEvicted = (1017 | MaskWarn)
const WarnRscDeactivated = (1018 | MaskWarn)
const WarnNotFound = (3000 | MaskWarn)

// ## Codes 10000-19999: info ###
const InfoNoRscSpawned = (10000 | MaskInfo)
const InfoNodeNameMismatch = (10001 | MaskInfo)
const InfoPropSet = (10002 | MaskInfo)
const InfoTieBreakerCreated = (10003 | MaskInfo)
const InfoTieBreakerDeleting = (10004 | MaskInfo)
const InfoTieBreakerTakeover = (10006 | MaskInfo)
const InfoPropRemoved = (10005 | MaskInfo)
const InfoAutoDrbdProxyCreated = (10007 | MaskInfo)
const InfoNoop = (10007 | MaskInfo)
const InfoRscAlreadyExists = (10008 | MaskInfo)

// ## Special codes ###
const UnknownApiCall = (0x0FFFFFFFFFFFFFFF | MaskError)
const ApiCallAuthReq = (0x0FFFFFFFFFFFFFFE | MaskError)
const ApiCallParseError = (0x0FFFFFFFFFFFFFFD | MaskError)

// ## SignIn codes ###
const SuccessSignIn = (10000 | MaskSuccess)
const FailSignIn = (10000 | MaskError)
const FailSignInMissingCredentials = (10001 | MaskError)

// ## Special answer message content types ###
// Textual MsgApiCallResponse responses
const ApiReply = "Reply"

// Indicates that the immediate answers to the API call are complete
const ApiEndOfImmediateAnswers = "EndOfImmediateAnswers"

// ## Create object APIs ###
const ApiCrtNode = "CrtNode"
const ApiCrtRsc = "CrtRsc"
const ApiCrtRscDfn = "CrtRscDfn"
const ApiCrtNetIf = "CrtNetIf"
const ApiCrtVlmDfn = "CrtVlmDfn"
const ApiCrtSnapshot = "CrtSnapshot"
const ApiCrtStorPoolDfn = "CrtStorPoolDfn"
const ApiCrtStorPool = "CrtStorPool"
const ApiCrtNodeConn = "CrtNodeConn"
const ApiCrtRscConn = "CrtRscConn"
const ApiCrtVlmConn = "CrtVlmConn"
const ApiAutoPlaceRsc = "AutoPlaceRsc"
const ApiCrtCryptPass = "CrtCryptPass"
const ApiCrtOfTargetNode = "CrtOfTargetNode"
const ApiRestoreVlmDfn = "RestoreVlmDfn"
const ApiRestoreSnapshot = "RestoreSnapshot"
const ApiCrtRscGrp = "CrtRscGrp"
const ApiCrtVlmGrp = "CrtVlmGrp"
const ApiSpawnRscDfn = "SpawnRscDfn"
const ApiCreateDevicePool = "CreateDevicePool"
const ApiMakeRscAvail = "MakeRscAvail"
const ApiCrtExosEnclosure = "CrtExosEnclosure"

// ## Modify object APIs ###
const ApiModNode = "ModNode"
const ApiModNodeConn = "ModNodeConn"
const ApiModRsc = "ModRsc"
const ApiToggleDisk = "ToggleDisk"
const ApiModRscConn = "ModRscConn"
const ApiModRscDfn = "ModRscDfn"
const ApiModNetIf = "ModNetIf"
const ApiModStorPool = "ModStorPool"
const ApiModStorPoolDfn = "ModStorPoolDfn"
const ApiModVlmDfn = "ModVlmDfn"
const ApiModVlm = "ModVlm"
const ApiModVlmConn = "ModVlmConn"
const ApiModSnapshot = "ModSnapshot"
const ApiModCryptPass = "ModCryptPass"
const ApiEnableDrbdProxy = "EnableDrbdProxy"
const ApiDisableDrbdProxy = "DisableDrbdProxy"
const ApiModDrbdProxy = "ModifyDrbdProxy"
const ApiRollbackSnapshot = "RollbackSnapshot"
const ApiShipSnapshot = "ShipSnapshot"
const ApiModKvs = "ModifyKvs"
const ApiModRscGrp = "ModifyRscGrp"
const ApiModVlmGrp = "ModifyVlmGrp"
const ApiActivateRsc = "ActivateRsc"
const ApiDeactivateRsc = "DeactivateRsc"
const ApiModExosDflts = "ModifyExosDefaults"
const ApiModExosEnclosure = "ModExosEnclosure"

// ## Delete object APIs ###
const ApiDelNode = "DelNode"
const ApiDelRsc = "DelRsc"
const ApiDelRscDfn = "DelRscDfn"
const ApiDelNetIf = "DelNetIf"
const ApiDelVlmDfn = "DelVlmDfn"
const ApiDelStorPoolDfn = "DelStorPoolDfn"
const ApiDelStorPool = "DelStorPool"
const ApiDelNodeConn = "DelNodeConn"
const ApiDelRscConn = "DelRscConn"
const ApiDelVlmConn = "DelVlmConn"
const ApiDelSnapshot = "DelSnapshot"
const ApiDelKvs = "DelKvs"
const ApiDelRscGrp = "DelRscGrp"
const ApiDelVlmGrp = "DelVlmGrp"
const ApiLostNode = "LostNode"
const ApiLostStorPool = "LostStorPool"
const ApiDelExosEnclosure = "DelExosEnclosure"

// ## Authentication APIs ###
const ApiSignIn = "SignIn"
const ApiVersion = "Version"

// ## Debug APIs ###
const ApiCrtDbgCnsl = "CrtDbgCnsl"
const ApiDstrDbgCnsl = "DstrDbgCnsl"

// ## Command APIs ###
const ApiControlCtrl = "ControlCtrl"
const ApiCmdShutdown = "Shutdown"
const ApiNodeReconnect = "NodeReconnect"
const ApiNodeRestore = "NodeRestore"
const ApiNodeEvict = "NodeEvict"

// ## List object APIs ###
const ApiLstNode = "LstNode"
const ApiLstRsc = "LstRsc"
const ApiLstRscDfn = "LstRscDfn"
const ApiLstNetIf = "LstNetIf"
const ApiLstVlmDfn = "LstVlmDfn"
const ApiLstVlm = "LstVlm"
const ApiLstSnapshotDfn = "LstSnapshotDfn"
const ApiLstStorPool = "LstStorPool"
const ApiLstStorPoolDfn = "LstStorPoolDfn"
const ApiLstErrorReports = "LstErrorReports"
const ApiReqErrorReports = "ReqErrorReports"
const ApiDelErrorReport = "DelErrorReport"
const ApiDelErrorReports = "DelErrorReports"
const ApiReqSosReport = "ReqSosReport"
const ApiReqRscConnList = "ReqRscConnList"
const ApiLstRscConn = "LstRscConn"
const ApiLstKvs = "LstKvs"
const ApiLstRscGrp = "LstRscGrp"
const ApiLstVlmGrp = "LstVlmGrp"
const ApiLstPhysStor = "LstPhysicalStorage"
const ApiLstSnapshotShippings = "LstSnapShips"
const ApiLstPropsInfo = "LstPropsInfo"
const ApiLstExosDflts = "LstExosDefaults"
const ApiLstExosEnclosures = "LstExosEnclosures"
const ApiExosEnclosureEvents = "ExosEvents"
const ApiExosExec = "ExosExec"
const ApiExosMap = "ExosMap"
const ApiLstExtFiles = "LstExtFiles"

// ## Query APIs ###
const ApiQryMaxVlmSize = "QryMaxVlmSize"
const ApiRspMaxVlmSize = "RspMaxVlmSize"

// ## Event APIs ###
const ApiCrtWatch = "CrtWatch"
const ApiDelWatch = "DelWatch"
const ApiEvent = "Event"
const ApiRptSpc = "RptSpc"
const ApiPing = "Ping"
const ApiPong = "Pong"
const ApiModInf = "ModInf"
const ApiVsnInf = "VsnInf"
const ApiSetCtrlProp = "SetCtrlProp"
const ApiDelCtrlProp = "DelCtrlProp"
const ApiLstCtrlProps = "LstCtrlProps"

// ## Encryption APIs ###
const ApiEnterCryptPass = "EnterCryptPass"

// ## External files APIs ###
const ApiSetExtFile = "SetExtFile"
const ApiDelExtFile = "DeleteExtFile"
const ApiDeployExtFile = "DeployExtFile"
const ApiUndeployExtFile = "UndeployExtFile"

// ## DRBD property keys ###
const KeyUuid = "UUID"
const KeyDrbdCurrentGi = "DrbdCurrentGi"
const KeyDmstats = "DMStats"
const KeyDrbdAutoQuorum = "auto-quorum"
const KeyDrbdAutoAddQuorumTiebreaker = "auto-add-quorum-tiebreaker"
const KeyMinorNrAutoRange = "MinorNrAutoRange"
const KeyDrbdAutoDiskful = "auto-diskful"
const KeyDrbdAutoDiskfulAllowCleanup = "auto-diskful-allow-cleanup"
const KeyDrbdDisableAutoVerifyAlgo = "auto-verify-algo-disable"
const KeyDrbdAutoVerifyAlgoAllowedUser = "auto-verify-algo-allowed-user-list"

// ## Node property keys ###
const KeyNode = "Node"
const Key1StNode = "FirstNode"
const Key2NdNode = "SecondNode"
const KeyCurStltConnName = "CurStltConnName"

// ## Resource property keys ###
const KeyRscDfn = "RscDfn"
const KeyRscGrp = "RscGrp"
const KeyTcpPortAutoRange = "TcpPortAutoRange"
const KeyPeerSlotsNewResource = "PeerSlotsNewResource"
const KeyPeerSlots = "PeerSlots"
const KeyRscRollbackTarget = "RollbackTarget"
const KeyRscMigrateFrom = "MigrateFrom"

// ## Volume property keys ###
const KeyVlmGrp = "VlmGrp"
const KeyVlmNr = "VlmNr"
const KeyVlmRestoreFromResource = "RestoreFromResource"
const KeyVlmRestoreFromSnapshot = "RestoreFromSnapshot"

// ## ldap property keys ###
const KeySearchDomain = "SearchDomain"

// ## nvme property keys ###
const KeyTrType = "TRType"

// ## Snapshot property keys ###
const KeySnapshot = "Snapshot"
const KeySnapshotDfnSequenceNumber = "SequenceNumber"

// ## Network Interface property keys ###
const KeyPort = "Port"
const KeyDisableHttpMetrics = "disable-http-metrics"

// ## Writecache property keys ###
const KeyWritecacheBlocksize = "Blocksize"
const KeyWritecachePoolName = "PoolName"
const KeyWritecacheSize = "Size"
const KeyWritecacheOptionHighWatermark = "HighWatermark"
const KeyWritecacheOptionLowWatermark = "LowWatermark"
const KeyWritecacheOptionStartSector = "StartSector"
const KeyWritecacheOptionWritebackJobs = "WritebackJobs"
const KeyWritecacheOptionAutocommitBlocks = "AutocommitBlocks"
const KeyWritecacheOptionAutocommitTime = "AutocommitTime"
const KeyWritecacheOptionFua = "Fua"
const KeyWritecacheOptionAdditional = "Additional"

// ## Cache property keys ###
const KeyCacheOperatingMode = "OpMode"
const KeyCacheMetaPoolName = "MetaPool"
const KeyCacheCachePoolName = "CachePool"
const KeyCacheMetaSize = "Metasize"
const KeyCacheCacheSize = "Cachesize"
const KeyCacheBlockSize = "Blocksize"
const KeyCachePolicy = "Policy"
const KeyUpdateCacheInterval = "UpdateCacheInterval"

// ## BCache property keys ###
const KeyBcachePoolName = "PoolName"
const KeyBcacheSize = "Size"
const KeyBcacheBlocksize = "Blocksize"
const KeyBcacheBucketsize = "Bucketsize"
const KeyBcacheDataOffset = "DataOffset"
const KeyBcacheWriteback = "Writeback"
const KeyBcacheDiscard = "Discard"
const KeyBcacheCacheReplacementPolicy = "CacheReplacementPolicy"

// ## Autoplace property keys ###
const KeyAutoplaceStratWeightMaxFreespace = "MaxFreeSpace"
const KeyAutoplaceStratWeightMinReservedSpace = "MinReservedSpace"
const KeyAutoplaceStratWeightMinRscCount = "MinRscCount"
const KeyAutoplacePreSelectFileName = "PreSelectScript"
const KeyAutoplacePreSelectScriptTimeout = "PreSelectScriptTimeout"
const KeyAutoplaceMaxThroughput = "MaxThroughput"
const KeySite = "Site"

// ## Auto-Evict property keys ###
const KeyAutoEvictMinReplicaCount = "AutoEvictMinReplicaCount"
const KeyAutoEvictAfterTime = "AutoEvictAfterTime"
const KeyAutoEvictMaxDisconnectedNodes = "AutoEvictMaxDisconnectedNodes"
const KeyAutoEvictAllowEviction = "AutoEvictAllowEviction"

// ## Snapshot shipping property keys ###
const KeySnapshotShippingPrefix = "SnapshotShippingPrefix"
const KeyTargetNode = "TargetNode"
const KeySourceNode = "SourceNode"
const KeyRunEvery = "RunEvery"
const KeyAutoSnapshotPrefix = "Prefix"
const KeyKeep = "Keep"
const KeyAutoSnapshotNextId = "NextAutoId"
const KeyTcpPortRange = "TcpPortRange"

// ## Property namespaces ###
const NamespcNetcom = "NetCom"
const NamespcDflt = "Default"
const NamespcLogging = "Logging"
const NamespcAlloc = "Allocation"
const NamespcNetif = "NetIf"
const NamespcStlt = "Satellite"
const NamespcNode = "Node"
const NamespcStorageDriver = "StorDriver"
const NamespcDrbdProxy = "DrbdProxy"
const NamespcAuxiliary = "Aux"
const NamespcDrbdOptions = "DrbdOptions"
const NamespcDrbdNetOptions = "DrbdOptions/Net"
const NamespcDrbdDiskOptions = "DrbdOptions/Disk"
const NamespcDrbdResourceOptions = "DrbdOptions/Resource"
const NamespcDrbdPeerDeviceOptions = "DrbdOptions/PeerDevice"
const NamespcDrbdProxyOptions = "DrbdOptions/Proxy"
const NamespcDrbdProxyCompressionOptions = "DrbdOptions/ProxyCompression"
const NamespcDrbdHandlerOptions = "DrbdOptions/Handlers"
const NamespcConnectionPaths = "Paths"
const NamespcRest = "REST"
const NamespcFilesystem = "FileSystem"
const NamespcNvme = "NVMe"
const NamespcSysFs = "sys/fs"
const NamespcWritecache = "Writecache"
const NamespcWritecacheOptions = "Writecache/Options"
const NamespcCache = "Cache"
const NamespcCacheFeatures = "Cache/Features"
const NamespcCachePolicyArgs = "Cache/Policy"
const NamespcBcache = "BCache"
const NamespcAutoplacer = "Autoplacer"
const NamespcAutoplacerWeights = "Autoplacer/Weights"
const NamespcSnapshotShipping = "SnapshotShipping"
const NamespcAutoSnapshot = "AutoSnapshot"
const NamespcStltDevSymlinks = "Satellite/Device/Symlinks"
const NamespcExos = "StorDriver/Exos"

// ## Storage pool property keys ###
const KeyStorPoolDfn = "StorPoolDfn"
const KeyStorPoolName = "StorPoolName"
const KeyStorPoolDrbdMetaName = "StorPoolNameDrbdMeta"
const KeyStorPoolVolumeGroup = "LvmVg"
const KeyStorPoolLvcreateType = "LvcreateType"
const KeyStorPoolLvcreateOptions = "LvcreateOptions"
const KeyStorPoolThinPool = "ThinPool"
const KeyStorPoolZpool = "ZPool"
const KeyStorPoolZpoolthin = "ZPoolThin"
const KeyStorPoolZfsCreateOptions = "ZfscreateOptions"
const KeyStorPoolFileDirectory = "FileDir"
const KeyStorPoolPrefNic = "PrefNic"
const KeyStorPoolCryptPasswd = "CryptPasswd"
const KeyStorPoolOverrideVlmId = "OverrideVlmId"
const KeyStorPoolDfnMaxOversubscriptionRatio = "MaxOversubscriptionRatio"
const KeyStorPoolWaitTimeoutAfterCreate = "WaitTimeoutAfterCreate"
const KeyStorPoolOpenflexApiHost = "Openflex/ApiHost"
const KeyStorPoolOpenflexApiPort = "Openflex/ApiPort"
const KeyStorPoolOpenflexStorDev = "Openflex/StorDev"
const KeyStorPoolOpenflexStorDevHost = "Openflex/StorDevHost"
const KeyStorPoolOpenflexStorPool = "Openflex/StorPool"
const KeyStorPoolOpenflexUserName = "Openflex/UserName"
const KeyStorPoolOpenflexUserPw = "Openflex/UserPassword"
const KeyStorPoolOpenflexJobWaitMaxCount = "Openflex/JobWaitMaxCount"
const KeyStorPoolOpenflexJobWaitDelay = "Openflex/JobWaitDelay"
const KeyOfTargetPortAutoRange = "OpenflexTargetPortAutoRange"
const KeyStorPoolExosApiIp = "IP"
const KeyStorPoolExosApiIpEnv = "IPEnv"
const KeyStorPoolExosApiPort = "Port"
const KeyStorPoolExosApiUser = "Username"
const KeyStorPoolExosApiUserEnv = "UsernameEnv"
const KeyStorPoolExosApiPassword = "Password"
const KeyStorPoolExosApiPasswordEnv = "PasswordEnv"
const KeyStorPoolExosVlmType = "VolumeType"
const KeyStorPoolExosCreateVolumeOptions = "CreateVolumeOptions"
const KeyStorPoolExosEnclosure = "Enclosure"
const KeyStorPoolExosPoolSn = "PoolSN"
const KeyPrefNic = "PrefNic"

// ## Storage pool traits keys ###
const KeyStorPoolSupportsSnapshots = "SupportsSnapshots"
const KeyStorPoolProvisioning = "Provisioning"

// Unit of smallest allocation. The size in KiB as a decimal number.
const KeyStorPoolAllocationUnit = "AllocationUnit"

// ## Storage pool traits values ###
const ValStorPoolProvisioningFat = "Fat"
const ValStorPoolProvisioningThin = "Thin"
const ValStorPoolDrbdMetaInternal = ".internal"

// ## DRBD Proxy keys (other than 'options') ###
const KeyDrbdProxyCompressionType = "CompressionType"
const KeyDrbdProxyAutoEnable = "AutoEnable"

// ## File system property keys ###
const KeyFsType = "Type"
const KeyFsMkfsparameters = "MkfsParams"
const ValFsTypeExt4 = "ext4"
const ValFsTypeXfs = "xfs"

// ## sys/fs property keys ###
const KeySysFsBlkioThrottleRead = "blkio_throttle_read"
const KeySysFsBlkioThrottleWrite = "blkio_throttle_write"
const KeySysFsBlkioThrottleReadIops = "blkio_throttle_read_iops"
const KeySysFsBlkioThrottleWriteIops = "blkio_throttle_write_iops"

// ## Property values ###
const ValNetcomTypeSsl = "SSL"
const ValNetcomTypePlain = "Plain"
const ValSslProtoTlsv1 = "TLSv1"

// ## DRBD related property values ###
const ValDrbdProxyCompressionNone = "none"
const ValDrbdProxyCompressionZstd = "zstd"
const ValDrbdProxyCompressionZlib = "zlib"
const ValDrbdProxyCompressionLzma = "lzma"
const ValDrbdProxyCompressionLz4 = "lz4"
const ValDrbdAutoQuorumDisabled = "disabled"
const ValDrbdAutoQuorumIoError = "io-error"
const ValDrbdAutoQuorumSuspendIo = "suspend-io"

// ## Node Type values ###
const ValNodeTypeCtrl = "Controller"
const ValNodeTypeStlt = "Satellite"
const ValNodeTypeCmbd = "Combined"
const ValNodeTypeAux = "Auxiliary"
const ValNodeTypeOpenflexTarget = "Openflex_Target"
const ValNodeTypeExosTarget = "Exos_Target"

// ## Writecache option values ###
const ValWritecacheFuaOn = "On"
const ValWritecacheFuaOff = "Off"

// ## Net interface Type values ###
const ValNetifTypeIp = "IP"
const ValNetifTypeRdma = "RDMA"
const ValNetifTypeRoce = "RoCE"

// ## Authentication keys ###
const KeySecIdentity = "SecIdentity"
const KeySecRole = "SecRole"
const KeySecType = "SecType"
const KeySecDomain = "SecDomain"
const KeySecPassword = "SecPassword"
const KeyPoolName = "PoolName"

// ## External commands keys ###
const KeyExtCmdWaitTo = "ExtCmdWaitTimeout"

// ## External files keys ###
const KeyExtFile = "ExtFile"

// ## Default ports ###
const DfltCtrlPortSsl = 3377
const DfltCtrlPortPlain = 3376
const DfltStltPortSsl = 3367
const DfltStltPortPlain = 3366

// ## Boolean values ###
const ValTrue = "True"
const ValFalse = "False"
const ValYes = "Yes"
const ValNo = "No"

// ## Snapshot-shipping values ###
const ValSnapShipName = "SnapshotShipping"

// enum generated in package -> "golinstor/snapshotshipstatus"
// snapshotshipstatus.Running = "Running"
// snapshotshipstatus.Complete = "Complete"

// ## Flag string values ###
const FlagClean = "CLEAN"
const FlagEvicted = "EVICTED"
const FlagDelete = "DELETE"
const FlagDiskless = "DISKLESS"
const FlagQignore = "QIGNORE"
const FlagEncrypted = "ENCRYPTED"
const FlagGrossSize = "GROSS_SIZE"
const FlagSuccessful = "SUCCESSFUL"
const FlagFailedDeployment = "FAILED_DEPLOYMENT"
const FlagFailedDisconnect = "FAILED_DISCONNECT"
const FlagResize = "RESIZE"
const FlagDiskAdding = "DISK_ADDING"
const FlagDiskAddRequested = "DISK_ADD_REQUESTED"
const FlagDiskRemoving = "DISK_REMOVING"
const FlagDiskRemoveRequested = "DISK_REMOVE_REQUESTED"
const FlagTieBreaker = "TIE_BREAKER"
const FlagDrbdDiskless = "DRBD_DISKLESS"
const FlagNvmeInitiator = "NVME_INITIATOR"
const FlagRscInactive = "INACTIVE"

// ## Device layer kinds ###
// enum generated in package -> "golinstor/devicelayerkind"
// devicelayerkind.Drbd = "DRBD"
// devicelayerkind.Luks = "LUKS"
// devicelayerkind.Storage = "STORAGE"
// devicelayerkind.Nvme = "NVME"
// devicelayerkind.Openflex = "OPENFLEX"
// devicelayerkind.Exos = "EXOS"
// devicelayerkind.Writecache = "WRITECACHE"
// devicelayerkind.Cache = "CACHE"

// ## Satellite connection statuses ###
// enum generated in package -> "golinstor/connectionstatus"
// connectionstatus.Offline = 0
// connectionstatus.Connected = 1
// connectionstatus.Online = 2
// connectionstatus.VersionMismatch = 3
// connectionstatus.FullSyncFailed = 4
// connectionstatus.AuthenticationError = 5
// connectionstatus.Unknown = 6
// connectionstatus.HostnameMismatch = 7
// connectionstatus.OtherController = 8
// connectionstatus.Authenticated = 9
// connectionstatus.NoStltConn = 10

// ## Default names ###
const DefaultNetif = "default"
const DfltSnapshotShippingPrefix = "ship"

// ## Default values ###
const DfltAutoSnapshotKeep = "10"
const DfltShippedSnapshotKeep = "10"
