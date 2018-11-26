package LZString

import (
	"math"
)

var keyStrBase64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/="

func CompressToBase64(uncompressed string) string {
	if len(uncompressed) == 0 {
		return ""
	}

	res := Compress(uncompressed, 6, func(character int) string {
		return string([]rune(keyStrBase64)[character])
	})

	switch len(res) % 4 {
	default:
	case 0:
		return res
	case 1:
		return res + "==="
	case 2:
		return res + "=="
	case 3:
		return res + "="
	}
	return res
}

func Compress(uncompressed string, bitsPerChar int, getCharFromInt func(character int) string) string {
	if len(uncompressed) == 0 {
		return ""
	}
	var value int
	contextDictionary := make(map[string]int)
	contextDictionaryToCreate := make(map[string]bool)
	var contextC string
	var contextW string
	var contextWc string
	contextEnlargeIn := float64(2)
	contextDictSize := int(3)
	contextNumBits := int(2)
	var contextDataString string
	contextDataVal := int(0)
	contextDataPosition := int(0)

	for ii := 0; ii < len(uncompressed); ii++ {
		contextC = string([]rune(uncompressed)[ii])
		_, in := contextDictionary[contextC]
		if !in {
			contextDictionary[contextC] = contextDictSize
			contextDictSize++
			contextDictionaryToCreate[contextC] = true
		}

		contextWc = contextW + contextC

		_, in = contextDictionary[contextWc]
		if in {
			contextW = contextWc
		} else {
			_, in = contextDictionaryToCreate[contextW]
			if in {
				if []rune(contextW)[0] < 256 {
					for i := 0; i < contextNumBits; i++ {
						contextDataVal = contextDataVal << 1
						if contextDataPosition == bitsPerChar-1 {
							contextDataPosition = 0
							contextDataString += getCharFromInt(contextDataVal)
							contextDataVal = 0
						} else {
							contextDataPosition++
						}
					}
					value = int([]rune(contextW)[0])
					for i := 0; i < 8; i++ {
						contextDataVal = (contextDataVal << 1) | (value & 1)
						if contextDataPosition == bitsPerChar-1 {
							contextDataPosition = 0
							contextDataString += getCharFromInt(contextDataVal)
							contextDataVal = 0
						} else {
							contextDataPosition++
						}
						value = value >> 1
					}
				} else {
					value = 1
					for i := 0; i < contextNumBits; i++ {
						contextDataVal = (contextDataVal << 1) | value
						if contextDataPosition == bitsPerChar-1 {
							contextDataPosition = 0
							contextDataString += getCharFromInt(contextDataVal)
							contextDataVal = 0
						} else {
							contextDataPosition++
						}
						value = 0
					}
					value = int([]rune(contextW)[0])
					for i := 0; i < 16; i++ {
						contextDataVal = (contextDataVal << 1) | (value & 1)
						if contextDataPosition == bitsPerChar-1 {
							contextDataPosition = 0
							contextDataString += getCharFromInt(contextDataVal)
							contextDataVal = 0
						} else {
							contextDataPosition++
						}
						value = value >> 1
					}
				}
				contextEnlargeIn--
				if contextEnlargeIn == 0 {
					contextEnlargeIn = math.Pow(2, float64(contextNumBits))
					contextNumBits++
				}
				delete(contextDictionaryToCreate, contextW)
			} else {
				value = contextDictionary[contextW]
				for i := 0; i < contextNumBits; i++ {
					contextDataVal = (contextDataVal << 1) | (value & 1)
					if contextDataPosition == bitsPerChar-1 {
						contextDataPosition = 0
						contextDataString += getCharFromInt(contextDataVal)
						contextDataVal = 0
					} else {
						contextDataPosition++
					}
					value = value >> 1
				}
			}
			contextEnlargeIn--
			if contextEnlargeIn == 0 {
				contextEnlargeIn = math.Pow(2, float64(contextNumBits))
				contextNumBits++
			}
			contextDictionary[contextWc] = contextDictSize
			contextDictSize++
			contextW = contextC
		}
	}

	if contextW != "" {
		_, in := contextDictionaryToCreate[contextW]
		if in {
			if []rune(contextW)[0] < 256 {
				for i := 0; i < contextNumBits; i++ {
					contextDataVal = contextDataVal << 1
					if contextDataPosition == bitsPerChar-1 {
						contextDataPosition = 0
						contextDataString += getCharFromInt(contextDataVal)
						contextDataVal = 0
					} else {
						contextDataPosition++
					}
				}
				value = int([]rune(contextW)[0])
				for i := 0; i < 8; i++ {
					contextDataVal = (contextDataVal << 1) | (value & 1)
					if contextDataPosition == bitsPerChar-1 {
						contextDataPosition = 0
						contextDataString += getCharFromInt(contextDataVal)
						contextDataVal = 0
					} else {
						contextDataPosition++
					}
					value = value >> 1
				}
			} else {
				value = 1
				for i := 0; i < contextNumBits; i++ {
					contextDataVal = (contextDataVal << 1) | value
					if contextDataPosition == bitsPerChar-1 {
						contextDataPosition = 0
						contextDataString += getCharFromInt(contextDataVal)
						contextDataVal = 0
					} else {
						contextDataPosition++
					}
					value = 0
				}
				value = int([]rune(contextW)[0])
				for i := 0; i < 16; i++ {
					contextDataVal = (contextDataVal << 1) | (value & 1)
					if contextDataPosition == bitsPerChar-1 {
						contextDataPosition = 0
						contextDataString += getCharFromInt(contextDataVal)
						contextDataVal = 0
					} else {
						contextDataPosition++
					}
					value = value >> 1
				}
			}
			contextEnlargeIn--
			if contextEnlargeIn == 0 {
				contextEnlargeIn = math.Pow(2, float64(contextNumBits))
				contextNumBits++
			}
			delete(contextDictionaryToCreate, contextW)
		} else {
			value = contextDictionary[contextW]
			for i := 0; i < contextNumBits; i++ {
				contextDataVal = (contextDataVal << 1) | (value & 1)
				if contextDataPosition == bitsPerChar-1 {
					contextDataPosition = 0
					contextDataString += getCharFromInt(contextDataVal)
					contextDataVal = 0
				} else {
					contextDataPosition++
				}
				value = value >> 1
			}
		}
		contextEnlargeIn--
		if contextEnlargeIn == 0 {
			contextEnlargeIn = math.Pow(2, float64(contextNumBits))
			contextNumBits++
		}
	}

	value = 2
	for i := 0; i < contextNumBits; i++ {
		contextDataVal = (contextDataVal << 1) | (value & 1)
		if contextDataPosition == bitsPerChar-1 {
			contextDataPosition = 0
			contextDataString += getCharFromInt(contextDataVal)
			contextDataVal = 0
		} else {
			contextDataPosition++
		}
		value = value >> 1
	}

	for {
		contextDataVal = contextDataVal << 1
		if contextDataPosition == bitsPerChar-1 {
			contextDataString += getCharFromInt(contextDataVal)
			break
		} else {
			contextDataPosition++
		}
	}
	return contextDataString
}
