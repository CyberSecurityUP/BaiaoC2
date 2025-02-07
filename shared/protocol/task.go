package protocol

type TaskType int

const (
	TaskTypeNOP TaskType = iota
	TaskTypeExit
	TaskTypeSet
	TaskTypeFile
	TaskTypeDirectory
	TaskTypeWhoAmI
	TaskTypeProcess
	TaskTypeRegistry
	TaskTypeRPortFwd
	TaskTypeEnvironment
	TaskTypeSOCKS
	TaskTypeTokens
	TaskTypeRun
	TaskTypeItemStore
	TaskTypeLocalExec
	TaskTypePrintScreen
	TaskTypeRemoteExec
	TaskTypeLink
	TaskTypeUnlink
	TaskTypeP2P
)

type TaskCode int

const (
	TaskCodeNOP TaskCode = iota
	TaskCodeExit
	TaskCodeSetSleep
	TaskCodeSetSpawnTo
	TaskCodeSetBlockDLLs
	TaskCodeSetPPID
	TaskCodeFileCopy
	TaskCodeFileMove
	TaskCodeFileDelete
	TaskCodeFileUpload
	TaskCodeFileDownload
	TaskCodeDirectoryPrint
	TaskCodeDirectoryChange
	TaskCodeDirectoryCreate
	TaskCodeDirectoryCopy
	TaskCodeDirectoryMove
	TaskCodeDirectoryList
	TaskCodeDirectoryDelete
	TaskCodeWhoAmI
	TaskCodeProcessList
	TaskCodeProcessKill
	TaskCodeProcessInjectSpawn
	TaskCodeProcessInjectExplicit
	TaskCodeRegistryQuery
	TaskCodeRegistryAdd
	TaskCodeRegistryDelete
	TaskCodeRPortFwdBind
	TaskCodeRPortFwdRead
	TaskCodeRPortFwdWrite
	TaskCodeRPortFwdClose
	TaskCodeEnvironmentGet
	TaskCodeEnvironmentSet
	TaskCodeSOCKSConnect
	TaskCodeSOCKSRead
	TaskCodeSOCKSWrite
	TaskCodeSOCKSClose
	TaskCodeTokensList
	TaskCodeTokensMake
	TaskCodeTokensSteal
	TaskCodeTokensUse
	TaskCodeTokensRevert
	TaskCodeTokensDelete
	TaskCodeTokensPurge
	TaskCodeRun
	TaskCodeItemStoreList
	TaskCodeItemStoreAdd
	TaskCodeItemStoreDelete
	TaskCodeItemStorePurge
	TaskCodeLocalExecDotNet
	TaskCodeLocalExecBOF
	TaskCodeLocalExecManagedPowerShell
	TaskCodeLocalExecUnmanagedPowerShell
	TaskCodePrintScreen
	TaskCodeRemoteExecWinRM
	TaskCodeRemoteExecWMI
	TaskCodeRemoteExecPsExec
	TaskCodeRemoteExecSSH
	TaskCodeLinkSMB
	TaskCodeLinkTCP
	TaskCodeUnlink
	TaskCodeP2PAcknowledge
	TaskCodeP2PPassThru
)
