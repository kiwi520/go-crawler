package parser

import (
	"crawler/engine"
	"regexp"
	"crawler/models"
	"strconv"
)
var ageRe  = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var heightRe  = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
var weightRe  = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([\d]+)kg</span></td>`)
var marriageRe  = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var incomeRe  = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var genderRe  = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var educationRe  = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe  = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
var xinzuoRe  = regexp.MustCompile(`<td><span class="label">星座：</span>([^<]+)</td>`)
var hukouRe  = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var houseRe  = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe  = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
func ParseProfile(contents []byte,namer string) engine.ParseResult{

	profile := models.Profile{}
	profile.Name =namer

	age ,err := strconv.Atoi(extracString(contents,ageRe))
	if err == nil{
		profile.Age = age
	}

	height ,err := strconv.Atoi(extracString(contents,heightRe))
	if err == nil{
		profile.Height = height
	}

	weight ,err := strconv.Atoi(extracString(contents,weightRe))
	if err == nil{
		profile.Height = weight
	}

	profile.Marriage = extracString(contents,marriageRe)
	profile.Income = extracString(contents,incomeRe)
	profile.Gender = extracString(contents,genderRe)
	profile.Education = extracString(contents,educationRe)
	profile.Occupation = extracString(contents,occupationRe)
	profile.Xinzuo = extracString(contents,xinzuoRe)
	profile.Hukou = extracString(contents,hukouRe)
	profile.House = extracString(contents,houseRe)
	profile.Car = extracString(contents,carRe)

	//fmt.Println(profile)
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return  result
}

func extracString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >=2{
		return string(match[1])
	}else{
		return ""
	}
}
