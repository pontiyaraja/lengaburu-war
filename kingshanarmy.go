package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	horse        = "H"
	elephant     = "E"
	armouredTank = "AT"
	slingGun     = "SG"
)

type Army struct {
	horse        int
	elephant     int
	armouredTank int
	slingGun     int
}
type armyData struct {
	resource        string
	value           int
	required        int
	defect          bool
	reachedMaxPower bool
}

var kingShanArmy Army
var alFalcone Army

var resourceMap map[string]int

func init() {
	kingShanArmy.horse = 100
	kingShanArmy.elephant = 50
	kingShanArmy.armouredTank = 10
	kingShanArmy.slingGun = 5

	resourceMap = make(map[string]int)
	resourceMap[horse] = kingShanArmy.horse
	resourceMap[elephant] = kingShanArmy.elephant
	resourceMap[armouredTank] = kingShanArmy.armouredTank
	resourceMap[slingGun] = kingShanArmy.slingGun

	alFalcone.horse = 300
	alFalcone.elephant = 200
	alFalcone.armouredTank = 40
	alFalcone.slingGun = 20
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	dataString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("invalid input data")
		os.Exit(0)
	}
	//Falicornia attacks with 100 H, 101 E, 20 AT, 5 SG

	strArr := strings.Split(dataString, " ")
	if len(strArr) < 10 {
		fmt.Println("invalid input data")
		os.Exit(0)
	}

	i1I := strArr[3]
	i2I := strArr[5]
	i3I := strArr[7]
	i4I := strArr[9]

	if strings.Compare(horse, strings.TrimSuffix(strArr[4], ",")) != 0 {
		fmt.Println(strArr[4], "invalid input data, please follow `n H, n E, n AT, n SG")
		os.Exit(0)
	}
	if strings.Compare(elephant, strings.TrimSuffix(strArr[6], ",")) != 0 {
		fmt.Println("invalid input data, please follow `n H, n E, n AT, n SG")
		os.Exit(0)
	}
	if strings.Compare(armouredTank, strings.TrimSuffix(strArr[8], ",")) != 0 {
		fmt.Println("invalid input data, please follow `n H, n E, n AT, n SG")
		os.Exit(0)
	}
	if strings.Compare(slingGun, strings.TrimSpace(strArr[10])) != 0 {
		fmt.Println("invalid input data, please follow `n H, n E, n AT, n SG")
		os.Exit(0)
	}

	var fhrs, fel, fat, fsg int
	fhrs, err = strconv.Atoi(i1I)
	if err != nil {
		fmt.Println("invalid input data")
		os.Exit(0)
	}

	fel, err = strconv.Atoi(i2I)
	if err != nil {
		fmt.Println("invalid input data")
		os.Exit(0)
	}
	fat, err = strconv.Atoi(i3I)
	if err != nil {
		fmt.Println("invalid input data")
		os.Exit(0)
	}
	fsg, err = strconv.Atoi(i4I)
	if err != nil {
		fmt.Println("invalid input data")
		os.Exit(0)
	}

	var battalionList []armyData

	battalionList = append(battalionList, getBattalion(fhrs, kingShanArmy.horse, horse))
	battalionList = append(battalionList, getBattalion(fel, kingShanArmy.elephant, elephant))
	battalionList = append(battalionList, getBattalion(fat, kingShanArmy.armouredTank, armouredTank))
	battalionList = append(battalionList, getBattalion(fsg, kingShanArmy.slingGun, slingGun))

	for i, batallion := range battalionList {
		if batallion.defect && batallion.required > 0 {
			if i == 0 && battalionList[i+1].defect {
				break
			} else if i > 0 && i < len(battalionList)-1 && battalionList[i+1].defect && battalionList[i-1].defect {
				break
			} else if i == len(battalionList)-1 && battalionList[i-1].defect {
				break
			}

			if i > 0 && !battalionList[i-1].defect && !battalionList[i-1].reachedMaxPower {
				value := getLowerArmyPower(batallion.required, resourceMap[battalionList[i-1].resource], &battalionList[i-1])
				if value == 0 {
					break
				} else {
					batallion.required = batallion.required - value/2
					if batallion.required <= 0 {
						battalionList[i].defect = false
						batallion.defect = false
					}
				}
			}
			if batallion.defect && i < len(battalionList)-1 && !battalionList[i+1].defect && !battalionList[i+1].reachedMaxPower {
				value := getUpperArmyPower(batallion.required, resourceMap[battalionList[i+1].resource], &battalionList[i+1])
				if value == 0 {
					break
				} else {
					batallion.required = batallion.required - value*2
					if batallion.required <= 0 {
						battalionList[i].defect = false
						batallion.defect = false
					}
				}
			}
		}
	}
	checkWin(battalionList)
}

func checkWin(battalionList []armyData) {
	var lost bool
	var res string
	for i, bat := range battalionList {
		if bat.defect && !lost {
			lost = true
		}
		if i > 0 && i < len(battalionList) {
			res = res + ",  "
		}
		res = res + strconv.Itoa(bat.value) + " " + bat.resource
	}
	if lost {
		res = res + " and loses"
	} else {
		res = res + " and wins"
	}
	fmt.Println("Lengaburu deploys", res)
}

func getBattalion(Flbattalion, kingBattalion int, resource string) armyData {
	var battalionData armyData
	battalionData.resource = resource

	if Flbattalion%2 == 0 {
		battalionData.value = Flbattalion / 2
	} else {
		battalionData.value = (Flbattalion / 2) + 1
	}
	if battalionData.value >= kingBattalion {
		battalionData.required = battalionData.value - kingBattalion
		battalionData.value = kingBattalion
		battalionData.reachedMaxPower = true
		if battalionData.required > 0 {
			battalionData.defect = true
		}
	}
	return battalionData
}

func getUpperArmyPower(requiredPower, maxBattalion int, nextBattalion *armyData) int {
	var power int
	if nextBattalion != nil && !nextBattalion.reachedMaxPower {
		if requiredPower == 1 {
			power = requiredPower
			if power+nextBattalion.value > maxBattalion {
				return 0
			} else if power+nextBattalion.value == maxBattalion {
				nextBattalion.reachedMaxPower = true
				nextBattalion.value += power
				return power
			}
			nextBattalion.value += power
			return power
		}
		power = requiredPower / 2
		if requiredPower%2 > 0 {
			power++
		}

		power = power + nextBattalion.value

		if power >= maxBattalion {
			requiredPower = power - maxBattalion
			power = maxBattalion
			nextBattalion.reachedMaxPower = true
		}
		nextBattalion.value = power
		return requiredPower
	}
	return 0
}

func getLowerArmyPower(requiredPower, maxBattalion int, preBattalion *armyData) int {
	if preBattalion != nil && !preBattalion.reachedMaxPower {
		var power int
		if requiredPower == 1 {
			power = requiredPower * 2
			if power+preBattalion.value > maxBattalion {
				return 0
			} else if power+preBattalion.value == maxBattalion {
				preBattalion.reachedMaxPower = true
				preBattalion.value += power
				return power
			}
			preBattalion.value += power
			return power
		}

		power = requiredPower * 2
		power = power + preBattalion.value
		if power > maxBattalion {
			requiredPower = power - maxBattalion
			power = maxBattalion
			preBattalion.reachedMaxPower = true
		}
		preBattalion.value = power
		return requiredPower
	}
	return 0
}
