package core

import "github.com/dotbitHQ/das-lib/common"

type Env struct {
	THQCodeHash      string
	ContractArgs     string
	ContractCodeHash string
	MapContract      map[common.DasContractName]string
	MapSoScript      map[common.SoScriptType]string
}

var EnvMainNet = Env{
	THQCodeHash:      "0x9e537bf5b8ec044ca3f53355e879f3fd8832217e4a9b41d9994cf0c547241a79",
	ContractArgs:     "0x0e65e63058c5b76b44ab13c15e9245fb8c301e04",
	ContractCodeHash: "0x00000000000000000000000000000000000000000000000000545950455f4944",
	MapContract: map[common.DasContractName]string{
		common.DasContractNameConfigCellType:            "0x3775c65aabe8b79980c4933dd2f4347fa5ef03611cef64328685618aa7535794",
		common.DasContractNameAccountCellType:           "0x96dc231bbbee6aa474076468640f9e0ad27cf13b1343716a7ce04b116ea18ba8",
		common.DasContractNameBalanceCellType:           "0xbdc8f42643ccad23e8df3d2e8dbdea9201812cd1b7f84c46e69b020529629822",
		common.DasContractNameDispatchCellType:          "0xda22fd296682488687a6035b5fc97c269b72d7de128034389bd03041b40309c0",
		common.DasContractNameIncomeCellType:            "0x108fba6a9b9f2898b4cdf11383ba2a6ed3da951b458c48e5f5de0353bbca2d46",
		common.DasContractNameAccountSaleCellType:       "0xb782d5f4e24603340997494871ba8f7d175e6920e63ead6b8137170b2e370469",
		common.DasContractNameAlwaysSuccess:             "0xca5016f232830f8a73e6827b5e1108aca68e7cf8baea4847ac40ef1da43c4c50",
		common.DasContractNameApplyRegisterCellType:     "0xf18c3eab9fd28adbb793c38be9a59864989c1f739c22d2b6dc3f4284f047a69d",
		common.DasContractNamePreAccountCellType:        "0xf6574955079797010689a22cd172ce55b52d2c34d1e9bc6596e97babc2906f7e",
		common.DasContractNameProposalCellType:          "0xd7b779b1b30f86a77db6b292c9492906f2437b7d88a8a5994e722619bb1d41c8",
		common.DasContractNameReverseRecordCellType:     "0x000f3e1a89d85d268ed6d36578d474ecf91d8809f4f696dd2e5b97fe67b84a2e",
		common.DASContractNameOfferCellType:             "0x3ffc0f8b0ce4bc09f700ca84355a092447d79fc5224a6fbd64af95af840af91b",
		common.DASContractNameSubAccountCellType:        "0x97b19f14184f24d55b1247596a5d7637f133c7bb7735f0ae962dc709c5fc1e2e",
		common.DASContractNameEip712LibCellType:         "0xa71fc12ca4ce4127dc911dbe8006f1907c850957849e282ef1f26e3ee50ba7bf",
		common.DasContractNameReverseRecordRootCellType: "0x3bd4491a15b3c84de133f4d619f5f7317fe7df7098a5c8dfd0ff073e448321a7",
		common.DasKeyListCellType:                       "0x97e271bfd332cc2e68feb6a4fa21a1dc010c53011d02d9acdbfc4d417d8efadd",
		common.DasContractNameDpCellType:                "0x5aa67507980f2bfdef75bf4f14caf886202f044d84fb4300858c9c823e3bb651",
		common.DasContractNameDidCellType:               "0x62312cd846659e188b05da11dc3f080b083c27371ea701d6026e11e713e0e3de",
	},
	MapSoScript: map[common.SoScriptType]string{
		common.SoScriptTypeEd25519:   "0xac4689bcb177af6ff0fee8c0c3bf0d8803b263381a5d5b22f1ca1b687eb5eb4a",
		common.SoScriptTypeEth:       "0x748005446caeda5b04e30bca42a65bb0b5275d2bd02694f08ecb8d3d07c41433",
		common.SoScriptTypeTron:      "0xb29620c50fd4fde746eae4d27e40b42d0900d1718fbf25ecae704e6a693215db",
		common.SoScriptTypeCkbMulti:  "0xab8a94c35890712763c70130db080fbbb7fe99ec74e7c87e4fdd44f4a6155646",
		common.SoScriptTypeCkbSingle: "0x4eed7961fdeaa659aa9ff59a242a2bc40f63dc043c1eba44c207ee6e05e02e9f",
		common.SoScriptTypeDogeCoin:  "0x2388d84e42714e51cbb1a60c4ab2df3321c632a204b7d963d7f9b89a50ceaa3c",
		common.SoScriptWebauthn:      "0x3dd5981ead36572330cc90b0bd4795062483392bfeea09f6ff8f02c830988998",
		common.SoScriptBitcoin:       "0xbeeb821fccb98acb8713cd0702c2d74729909222608ee78be42c6a15c1eb2014",
	},
}

var EnvTestnet2 = Env{
	THQCodeHash:      "0x96248cdefb09eed910018a847cfb51ad044c2d7db650112931760e3ef34a7e9a",
	ContractArgs:     "0x98801f99896300d249beb551fd08d0ca68497e14",
	ContractCodeHash: "0x00000000000000000000000000000000000000000000000000545950455f4944",
	MapContract: map[common.DasContractName]string{
		common.DasContractNameConfigCellType:            "0x34363fad2018db0b3b6919c26870f302da74c3c4ef4456e5665b82c4118eda51",
		common.DasContractNameAccountCellType:           "0x6f0b8328b703617508d62d1f017b0d91fab2056de320a7b7faed4c777a356b7b",
		common.DasContractNameBalanceCellType:           "0x27560fe2daa6150b771621300d1d4ea127832b7b326f2d70eed63f5333b4a8a9",
		common.DasContractNameDispatchCellType:          "0xeedd10c7d8fee85c119daf2077fea9cf76b9a92ddca546f1f8e0031682e65aee",
		common.DasContractNameIncomeCellType:            "0xd7b9d8213671aec03f3a3ab95171e0e79481db2c084586b9ea99914c00ff3716",
		common.DasContractNameAccountSaleCellType:       "0xed5d7fc00a3f8605bfe3f6717747bb0ed529fa064c2b8ce56e9677a0c46c2c1c",
		common.DasContractNameAlwaysSuccess:             "0x7821c662b7efd50e7f6cf2b036efe53e07eccaf2e3447a2a470ee07ae455ab92",
		common.DasContractNameApplyRegisterCellType:     "0xc78fa9066af1624e600ccfb21df9546f900b2afe5d7940d91aefc115653f90d9",
		common.DasContractNamePreAccountCellType:        "0xd3f7ad59632a2ebdc2fe9d41aa69708ed1069b074cd8b297b205f835335d3a6b",
		common.DasContractNameProposalCellType:          "0x03d0bb128bd10e666975d9a07c148f6abebe811f511e9574048b30600b065b9a",
		common.DasContractNameReverseRecordCellType:     "0x334d7841eb156b8aa5abd7b09277e91e782d840140905496bb5bff0ea6ce9d75",
		common.DASContractNameOfferCellType:             "0x443b2d1b3b00ffab1a2287d84c47b2c31a11aad24b183d732c213a69e3d6d390",
		common.DASContractNameSubAccountCellType:        "0x63ca3e26cc69809f06735c6d9139ec2d84f2a277f13509a54060d6ee19423b5b",
		common.DASContractNameEip712LibCellType:         "0x16549cab7e92afb5f157141bc9da7781ce692a3144e47e2b8879a8d5a57b87c6",
		common.DasContractNameReverseRecordRootCellType: "0xd7dbb55237fca8aa544d160e26a22424b4edc6b9ce136994df977fa057e468ae",
		common.DasKeyListCellType:                       "0xf80a65bd78e54f66614eb5d17c980e66c1944e7808510b6f08bb46ae5407147c",
		common.DasContractNameDpCellType:                "0xf5af4173248dbeaf4a54d1aa93958e094b98c90400be92bb2b1daa61dfc37785",
		common.DasContractNameDidCellType:               "0x80f0d4bf6b3951911aa6b98cc609d477a8a10b903b35cfd528d098e95c36f680",
	},
	MapSoScript: map[common.SoScriptType]string{
		common.SoScriptTypeEd25519:   "0x6cd3e59963ebd6aec1ee99968f2ec418ced3273ad636a561cd5459836c903c2d",
		common.SoScriptTypeEth:       "0x644ead361f2a2fe32431a81489a7a0f8d7332ba17918ab5184461d588f011a82",
		common.SoScriptTypeTron:      "0x2df08d17fe6ff0994b741344fe62e100b7576c31e006f8bc83d26f26ede8d82e",
		common.SoScriptTypeCkbMulti:  "0x393ac0be9d3bf4a282d2cf0b80a4a910206b87451fc6a3fb04d176f3d0c1f73a",
		common.SoScriptTypeCkbSingle: "0x249b06f89e76f98cebcd945bea47926d6ff88908cbbc8be2ccfb0a915e349b21",
		common.SoScriptTypeDogeCoin:  "0x4d801e313bca5426c254fb9ac73fcdb79f9eee216a911ec39526a00e3379f496",
		common.SoScriptWebauthn:      "0xe8e93df527c7dd6f4deb2d5ac3a071649f5794a4ed8d565adca74e7bee634c48",
		common.SoScriptBitcoin:       "0xaaa8f5e85e59392a3c26eb0533851efaf9daf8793349a2d117025d6b474be3fa",
	},
}

var EnvTestnet3 = Env{
	THQCodeHash:      "0x96248cdefb09eed910018a847cfb51ad044c2d7db650112931760e3ef34a7e9a",
	ContractArgs:     "0xbc502a34a430e3e167c82a24db6f9237b15ebf35",
	ContractCodeHash: "0x00000000000000000000000000000000000000000000000000545950455f4944",
	MapContract: map[common.DasContractName]string{
		common.DasContractNameConfigCellType:            "0xe9fa679290f63ba1debc74f04e96f299d1b61a03a24e7a1a51c7ccad416ec16a",
		common.DasContractNameAccountCellType:           "0x1c7d62798d412351ec9a4262aacc2c2837712b780d90e0720051ffa6d6304e32",
		common.DasContractNameBalanceCellType:           "0x301127f501c2620174b60945f305d78a4f5cdf67dd44009001265ca133d0088d",
		common.DasContractNameDispatchCellType:          "0xeedd10c7d8fee85c119daf2077fea9cf76b9a92ddca546f1f8e0031682e65aee",
		common.DasContractNameIncomeCellType:            "0x0af1f2332b1f742136ef0ffec849d415c4b5b8c4426de5d245f2e7ac0a9f5773",
		common.DasContractNameAccountSaleCellType:       "0xaf446d756fca5b51fad915a3f0526221527f1dd5f32898408ee86e73dc9d9814",
		common.DasContractNameAlwaysSuccess:             "0x7821c662b7efd50e7f6cf2b036efe53e07eccaf2e3447a2a470ee07ae455ab92",
		common.DasContractNameApplyRegisterCellType:     "0xb12e454a3896b910030e669e8a81bf99c351cb58be8d1529e1c7023265db6084",
		common.DasContractNamePreAccountCellType:        "0x919e8036ded6670b923958db207527b5d1d3b95955e446e3475bfd7f949898c1",
		common.DasContractNameProposalCellType:          "0x2eecc2c70936609d14872acdedc9ba8b265c26240d49d66b6db027c5cfdb4256",
		common.DasContractNameReverseRecordCellType:     "0x80963278cbdc61cdafd5250555984b71ad016798b8879adc0e6b1ee7e01b7912",
		common.DASContractNameOfferCellType:             "0xc69186c17e41fead0f87eb1f94829778e98a398be202655ac59fdb9567d05bae",
		common.DASContractNameSubAccountCellType:        "0x57498a2df0c0137146ced681fa1854599e404da5804c1a5ff45d954c3cc89bfd",
		common.DASContractNameEip712LibCellType:         "0xd21b1f02b6057d6776bf7f1c7a26f550f1b383a930349620507f328e4cc07f2e",
		common.DasContractNameReverseRecordRootCellType: "0xeb54ad6c4be167c902470c4fe9b04a4d8e2b5ef46f4b2725e814518d7baa15ca",
		common.DasKeyListCellType:                       "0x449bb468a1fefb93e5b034a070d46e6772fcd25712dc56fc8bcc191b4433f1a0",
		common.DasContractNameDpCellType:                "",
		common.DasContractNameDidCellType:               "",
	},
	MapSoScript: map[common.SoScriptType]string{
		common.SoScriptTypeEd25519:   "0x6cd3e59963ebd6aec1ee99968f2ec418ced3273ad636a561cd5459836c903c2d",
		common.SoScriptTypeEth:       "0x644ead361f2a2fe32431a81489a7a0f8d7332ba17918ab5184461d588f011a82",
		common.SoScriptTypeTron:      "0x2df08d17fe6ff0994b741344fe62e100b7576c31e006f8bc83d26f26ede8d82e",
		common.SoScriptTypeCkbMulti:  "0x393ac0be9d3bf4a282d2cf0b80a4a910206b87451fc6a3fb04d176f3d0c1f73a",
		common.SoScriptTypeCkbSingle: "0x249b06f89e76f98cebcd945bea47926d6ff88908cbbc8be2ccfb0a915e349b21",
		common.SoScriptTypeDogeCoin:  "0x4d801e313bca5426c254fb9ac73fcdb79f9eee216a911ec39526a00e3379f496",
		common.SoScriptWebauthn:      "0xe8e93df527c7dd6f4deb2d5ac3a071649f5794a4ed8d565adca74e7bee634c48",
	},
}

func InitEnv(net common.DasNetType) Env {
	switch net {
	case common.DasNetTypeMainNet:
		return EnvMainNet
	case common.DasNetTypeTestnet2:
		return EnvTestnet2
	case common.DasNetTypeTestnet3:
		return EnvTestnet3
	default:
		return EnvMainNet
	}
}

func InitEnvOpt(net common.DasNetType, names ...common.DasContractName) Env {
	switch net {
	case common.DasNetTypeMainNet:
		return initEnvOpt(EnvMainNet, names...)
	case common.DasNetTypeTestnet2:
		return initEnvOpt(EnvTestnet2, names...)
	case common.DasNetTypeTestnet3:
		return initEnvOpt(EnvTestnet3, names...)
	default:
		return initEnvOpt(EnvMainNet, names...)
	}
}

func initEnvOpt(envNet Env, names ...common.DasContractName) Env {
	env := Env{
		THQCodeHash:      envNet.THQCodeHash,
		ContractArgs:     envNet.ContractArgs,
		ContractCodeHash: envNet.ContractCodeHash,
		MapContract:      map[common.DasContractName]string{},
		MapSoScript:      envNet.MapSoScript,
	}
	for _, v := range names {
		if contract, ok := envNet.MapContract[v]; ok {
			env.MapContract[v] = contract
		}
	}
	return env
}
