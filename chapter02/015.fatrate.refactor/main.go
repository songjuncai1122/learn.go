package main

import (
	"fmt"
)

func main() {
	for {
		mainFatRateBody()
		if cont := whetherContinue();!cont{
			break
		}
	}
}

func mainFatRateBody(){
	weight,tall,age,sexWeight,sex:=getMaterialsFromInput()
	fatRate := calcFatRate(weight,tall,age,sexWeight)
	getHealthinessSuggestions(sex, age, fatRate)
}

func getHealthinessSuggestions(sex string, age int, fatRate float64) {
	if sex == "男" {
		// 编写男性的体脂率与体脂状态表
		if age >= 18 && age <= 39 {
			if fatRate <= 0.1 {
				fmt.Println("目前是：偏瘦。要多吃多锻炼，增加体质")
			} else if fatRate > 0.1 && fatRate <= 0.16 {
				fmt.Println("目前是：标准。")
			} else if fatRate > 0.16 && fatRate <= 0.21 {
				fmt.Println("目前是：偏胖")
			} else if fatRate > 0.21 && fatRate <= 0.26 {
				fmt.Println("目前是：肥胖。少吃点，多运动")
			} else {
				fmt.Println("目前是:非常肥胖。健身游泳跑步，即刻开始")
			}

		} else if age >= 40 && age <= 59 {
			if fatRate <= 0.11 {
				fmt.Println("目前是：偏瘦。要多吃多锻炼，增加体质")
			} else if fatRate > 0.11 && fatRate <= 0.17 {
				fmt.Println("目前是：标准。")
			} else if fatRate > 0.17 && fatRate <= 0.22 {
				fmt.Println("目前是：偏胖")
			} else if fatRate > 0.22 && fatRate <= 0.27 {
				fmt.Println("目前是：肥胖。少吃点，多运动")
			} else {
				fmt.Println("目前是:非常肥胖。健身游泳跑步，即刻开始")
			}
		} else if age >= 60 {
			if fatRate <= 0.13 {
				fmt.Println("目前是：偏瘦。要多吃多锻炼，增加体质")
			} else if fatRate > 0.13 && fatRate <= 0.19 {
				fmt.Println("目前是：标准。")
			} else if fatRate > 0.19 && fatRate <= 0.24 {
				fmt.Println("目前是：偏胖")
			} else if fatRate > 0.24 && fatRate <= 0.29 {
				fmt.Println("目前是：肥胖。少吃点，多运动")
			} else {
				fmt.Println("目前是:非常肥胖。健身游泳跑步，即刻开始")
			}
		} else {
			fmt.Println("我们不看未成年人的体脂率，变化太大了，无法评判")
		}
	} else {
		//  编写女性的体脂率与体脂状态表
		if age >= 18 && age <= 39 {
			if fatRate <= 0.20 {
				fmt.Println("目前是：偏瘦。要多吃多锻炼，增加体质")
			} else if fatRate > 0.20 && fatRate <= 0.27 {
				fmt.Println("目前是：标准。")
			} else if fatRate > 0.27 && fatRate <= 0.34 {
				fmt.Println("目前是：偏胖")
			} else if fatRate > 0.34 && fatRate <= 0.39 {
				fmt.Println("目前是：肥胖。少吃点，多运动")
			} else {
				fmt.Println("目前是:非常肥胖。健身游泳跑步，即刻开始")
			}
		} else if age >= 40 && age <= 59 {
			if fatRate <= 0.21 {
				fmt.Println("目前是：偏瘦。要多吃多锻炼，增加体质")
			} else if fatRate > 0.21 && fatRate <= 0.28 {
				fmt.Println("目前是：标准。")
			} else if fatRate > 0.28 && fatRate <= 0.35 {
				fmt.Println("目前是：偏胖")
			} else if fatRate > 0.35 && fatRate <= 0.40 {
				fmt.Println("目前是：肥胖。少吃点，多运动")
			} else {
				fmt.Println("目前是:非常肥胖。健身游泳跑步，即刻开始")
			}
		} else if age >= 60 {
			if fatRate <= 0.22 {
				fmt.Println("目前是：偏瘦。要多吃多锻炼，增加体质")
			} else if fatRate > 0.22 && fatRate <= 0.29 {
				fmt.Println("目前是：标准。")
			} else if fatRate > 0.29 && fatRate <= 0.36 {
				fmt.Println("目前是：偏胖")
			} else if fatRate > 0.36 && fatRate <= 0.41 {
				fmt.Println("目前是：肥胖。少吃点，多运动")
			} else {
				fmt.Println("目前是:非常肥胖。健身游泳跑步，即刻开始")
			}
		} else {
			fmt.Println("我们不看未成年人的体脂率，变化太大了，无法评判")
		}
	}
}

func whetherContinue()bool{
	var whetherContinue string
	fmt.Print("是否录入下一个（y/n)?")
	fmt.Scanln(&whetherContinue)
	if whetherContinue != "y" {
		return false
	}
	return true
}

func calcFatRate(weight float64,tall float64,age int,sexWeight int)float64{
	// 计算体脂率
	var bmi float64=weight/(tall*tall)
	var fatRate float64=(1.2*bmi+0.23*float64(age)-5.4-10.8*float64(sexWeight))/100
	fmt.Println("体脂率是：",fatRate)
	return fatRate
}

func getMaterialsFromInput()(float64,float64,int,int,string){
	//录入名项
	var name string
	fmt.Print("姓名：")
	fmt.Scanln(&name)

	var weight float64
	fmt.Print("体重(千克):")
	fmt.Scanln(&weight)

	var tall float64
	fmt.Print("身高（米）：")
	fmt.Scanln(&tall)

	var age int
	fmt.Print("年龄：")
	fmt.Scanln(&age)

	var sexWeight int
	var sex string = "男"
	fmt.Print("性别（男/女):")
	fmt.Scanln(&sex)

	if sex == "男"{
		sexWeight =1
	}else {
		sexWeight =0
	}
	return weight,tall,age,sexWeight,sex

}

