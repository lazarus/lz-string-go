package LZString

import (
	"errors"
	"math"
	"strings"
	"sync"
	"unicode/utf8"
)

// Compress
//goland:noinspection SpellCheckingInspection
const _defaultKeyStrBase64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/="

func Compress(uncompressed string, keyStrBase64 string) string {
	if len(uncompressed) == 0 {
		return ""
	}
	if keyStrBase64 == "" {
		keyStrBase64 = _defaultKeyStrBase64
	}
	charArr := []rune(keyStrBase64)
	res := _compress(uncompressed, 6, charArr)
	switch len(res) % 4 {
	case 3:
		return res + "="
	case 2:
		return res + "=="
	case 1:
		return res + "==="
	}
	return res
}

func _compress(uncompressed string, bitsPerChar int, charArr []rune) string {
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
	contextDictSize := 3
	contextNumBits := 2
	//var contextDataString string
	var contextDataString = strings.Builder{}

	contextDataVal := 0
	contextDataPosition := 0

	for ii := 0; ii < len(uncompressed); ii++ {
		contextC = string(uncompressed[ii])
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
				contextWRune := int(contextW[0])
				if contextWRune < 256 {
					for i := 0; i < contextNumBits; i++ {
						contextDataVal = contextDataVal << 1
						if contextDataPosition == bitsPerChar-1 {
							contextDataPosition = 0
							contextDataString.WriteRune(charArr[contextDataVal])
							contextDataVal = 0
						} else {
							contextDataPosition++
						}
					}
					value = contextWRune
					for i := 0; i < 8; i++ {
						contextDataVal = (contextDataVal << 1) | (value & 1)
						if contextDataPosition == bitsPerChar-1 {
							contextDataPosition = 0
							contextDataString.WriteRune(charArr[contextDataVal])
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
							contextDataString.WriteRune(charArr[contextDataVal])
							contextDataVal = 0
						} else {
							contextDataPosition++
						}
						value = 0
					}
					value = contextWRune
					for i := 0; i < 16; i++ {
						contextDataVal = (contextDataVal << 1) | (value & 1)
						if contextDataPosition == bitsPerChar-1 {
							contextDataPosition = 0
							contextDataString.WriteRune(charArr[contextDataVal])
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
						contextDataString.WriteRune(charArr[contextDataVal])
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
			contextWRune := int(contextW[0])
			if contextWRune < 256 {
				for i := 0; i < contextNumBits; i++ {
					contextDataVal = contextDataVal << 1
					if contextDataPosition == bitsPerChar-1 {
						contextDataPosition = 0
						contextDataString.WriteRune(charArr[contextDataVal])
						contextDataVal = 0
					} else {
						contextDataPosition++
					}
				}
				value = contextWRune
				for i := 0; i < 8; i++ {
					contextDataVal = (contextDataVal << 1) | (value & 1)
					if contextDataPosition == bitsPerChar-1 {
						contextDataPosition = 0
						contextDataString.WriteRune(charArr[contextDataVal])
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
						contextDataString.WriteRune(charArr[contextDataVal])
						contextDataVal = 0
					} else {
						contextDataPosition++
					}
					value = 0
				}
				value = contextWRune
				for i := 0; i < 16; i++ {
					contextDataVal = (contextDataVal << 1) | (value & 1)
					if contextDataPosition == bitsPerChar-1 {
						contextDataPosition = 0
						contextDataString.WriteRune(charArr[contextDataVal])
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
					contextDataString.WriteRune(charArr[contextDataVal])
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
			contextDataString.WriteRune(charArr[contextDataVal])
			contextDataVal = 0
		} else {
			contextDataPosition++
		}
		value = value >> 1
	}

	for {
		contextDataVal = contextDataVal << 1
		if contextDataPosition == bitsPerChar-1 {
			contextDataString.WriteRune(charArr[contextDataVal])
			break
		} else {
			contextDataPosition++
		}
	}
	return contextDataString.String()
}

// Decompress
var baseReverseDic = sync.Map{}

type dataStruct struct {
	input      string
	alphabet   string
	val        int
	position   int
	index      int
	dictionary []string
	enlargeIn  float64
	numBits    int
}

func covertToBaseReverseDic(alphabet string) map[byte]int {
	var val = map[byte]int{}
	charArr := []rune(alphabet)
	for i := 0; i < len(charArr); i++ {
		val[byte(charArr[i])] = i
	}
	return val
}

func getBaseValue(alphabet string, char byte) int {
	vv, ok := baseReverseDic.Load(alphabet)
	var arr map[byte]int
	if ok {
		arr = vv.(map[byte]int)
	} else {
		arr = covertToBaseReverseDic(alphabet)
		baseReverseDic.Store(alphabet, arr)
	}
	return arr[char]
}

// Input is composed of ASCII characters, so accessing it by array has no UTF-8 pb.
func readBits(nb int, data *dataStruct) int {
	result := 0
	power := 1
	for i := 0; i < nb; i++ {
		respB := data.val & data.position
		data.position = data.position / 2
		if data.position == 0 {
			data.position = 32
			data.val = getBaseValue(data.alphabet, data.input[data.index])
			data.index += 1
		}
		if respB > 0 {
			result |= power
		}
		power *= 2
	}
	return result
}

func appendValue(data *dataStruct, str string) {
	data.dictionary = append(data.dictionary, str)
	data.enlargeIn -= 1
	if data.enlargeIn == 0 {
		data.enlargeIn = math.Pow(2, float64(data.numBits))
		data.numBits += 1
	}
}

func getString(last string, data *dataStruct) (string, bool, error) {
	c := readBits(data.numBits, data)
	switch c {
	case 0:
		str := string(rune(readBits(8, data)))
		appendValue(data, str)
		return str, false, nil
	case 1:
		str := string(rune(readBits(16, data)))
		appendValue(data, str)
		return str, false, nil
	case 2:
		return "", true, nil
	}
	if c < len(data.dictionary) {
		return data.dictionary[c], false, nil
	}
	if c == len(data.dictionary) {
		return concatWithFirstRune(last, last), false, nil
	}
	return "", false, errors.New("bad character encoding")
}

// Need to handle UTF-8, so we need to use rune to concatenate
func concatWithFirstRune(str string, getFirstRune string) string {
	r, _ := utf8.DecodeRuneInString(getFirstRune)
	return str + string(r)
}

func Decompress(input string, keyStrBase64 string) (string, error) {
	if keyStrBase64 == "" {
		keyStrBase64 = _defaultKeyStrBase64
	}
	data := dataStruct{input, keyStrBase64, getBaseValue(keyStrBase64, input[0]), 32, 1, []string{"0", "1", "2"}, 5, 2}
	result, isEnd, err := getString("", &data)
	if err != nil || isEnd {
		return result, err
	}
	last := result
	data.numBits += 1
	for {
		str, isEnd, err := getString(last, &data)
		if err != nil || isEnd {
			return result, err
		}
		result = result + str
		appendValue(&data, concatWithFirstRune(last, str))
		last = str
	}
}
