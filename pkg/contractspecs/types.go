package contractspecs

import (
	"sync"
)

const (
	// To be kept up to date with supported contracts
	// TODO(PN): Build some awesome automated way to generate this based on
	// generate/contract code.

	// InvalidContractType is the enum value for an invalid contract type.
	InvalidContractType ContractType = iota

	// CivilTcrContractType is the enum value for the Civil TCR type
	CivilTcrContractType

	// NewsroomContractType is the enum value for the Newsroom type
	NewsroomContractType

	// CivilPLCRVotingContractType is the enum value for the Civil PLCR Voting type
	CivilPLCRVotingContractType

	// CVLTokenContractType is the enum value for the CVL token type
	CVLTokenContractType

	// CivilParameterizerContractType is the enum value for the Civil Parameterizer type
	CivilParameterizerContractType

	// CivilGovernmentContractType is the enum value for the Civil Government type
	CivilGovernmentContractType

	// MultiSigWalletContractType is the enum value for the Gnosis Multi Sig Wallet type
	MultiSigWalletContractType

	// MultiSigWalletFactoryContractType is the enum value for the Multi Sig Factory type
	MultiSigWalletFactoryContractType

	// NewsroomFactoryType is the enum value for the Newsroom Factory type
	NewsroomFactoryType
)

// ContractSpecs specifies metadata around a smart contract to be used in the
// crawler.
type ContractSpecs struct {
	name        string
	simpleName  string
	abiStr      string
	importPath  string
	typePackage string
}

// Name returns the full contract name
func (c *ContractSpecs) Name() string {
	return c.name
}

// SimpleName returns the short contract name
func (c *ContractSpecs) SimpleName() string {
	return c.simpleName
}

// AbiStr returns the contract ABI string
func (c *ContractSpecs) AbiStr() string {
	return c.abiStr
}

// ImportPath returns the import path to the contract
func (c *ContractSpecs) ImportPath() string {
	return c.importPath
}

// TypePackage returns the package of the smart contract
func (c *ContractSpecs) TypePackage() string {
	return c.typePackage
}

// CSpecs is a struct that contains a map from ContractType to
// contractSpecs
type CSpecs struct {
	specs map[ContractType]*ContractSpecs
}

// Get returns the contract specs for a given ContractType
func (c *CSpecs) Get(t ContractType) (*ContractSpecs, bool) {
	specs, ok := c.specs[t]
	return specs, ok
}

// Types returns a list of available types in CSpecs specs
func (c *CSpecs) Types() []ContractType {
	types := make([]ContractType, len(c.specs))
	index := 0
	for k := range c.specs {
		types[index] = k
		index++
	}
	return types
}

// ContractType is an enum for the Civil contract type
type ContractType int

// CTypes is a struct that contains a map of readable name to a
// ContractType enum value
type CTypes struct {
	types            map[string]ContractType
	simpleNameToName map[string]string
	mutex            sync.Mutex
}

// Get returns the contract type for a given contract simple name
func (c *CTypes) Get(name string) (ContractType, bool) {
	if c.types == nil || len(c.types) == 0 {
		c.build()
	}
	_type, ok := c.types[name]
	return _type, ok
}

// Names returns a list of the names in NameToContractType
func (c *CTypes) Names() []string {
	if c.types == nil || len(c.types) == 0 {
		c.build()
	}
	keys := make([]string, len(c.types))
	keyIndex := 0
	for k := range c.types {
		keys[keyIndex] = k
		keyIndex++
	}
	return keys
}

// GetFromContractName returns the contract type for a given contract name
func (c *CTypes) GetFromContractName(name string) (ContractType, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.types == nil || len(c.types) == 0 {
		c.build()
	}
	if c.simpleNameToName == nil || len(c.simpleNameToName) == 0 {
		c.buildSimpleNameToName()
	}

	simpleName := c.simpleNameToName[name]
	_type, ok := c.types[simpleName]
	return _type, ok
}

func (c *CTypes) buildSimpleNameToName() {
	c.simpleNameToName = make(map[string]string, len(ContractTypeToSpecs.specs))
	for _, spec := range ContractTypeToSpecs.specs {
		c.simpleNameToName[spec.name] = spec.simpleName
	}
}

func (c *CTypes) build() {
	c.types = make(map[string]ContractType, len(ContractTypeToSpecs.specs))
	for _type, spec := range ContractTypeToSpecs.specs {
		c.types[spec.simpleName] = _type
	}
}
