package parser

import (
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"regexp"
	"strconv"
)

//年龄
var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)岁</div>`)
//身高
var heightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)cm</div>`)
//月收入
var incomeRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>月收入:3-5千</div>`)
//体重
var weightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>48kg</div>`)
//性别
var genderRe = regexp.MustCompile(``)
//星座
var xingzuoRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>魔羯座(12.22-01.19)</div>`)
//婚姻状况
var marriageRe = regexp.MustCompile(`<div class="m-btn" data-v-bff6f798>([^<]+)</div>`)
//学历
var educationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>大学本科</div>`)
//职业
var occupationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>其他职业</div>`)
//籍贯
var hokouRe = regexp.MustCompile(`<div class="m-btn pink" data-v-bff6f798>籍贯:河南安阳</div>`)
//住房
var houseRe = regexp.MustCompile(`<div class="m-btn pink" data-v-bff6f798>未买车</div>`)
//是否购车
var carRe = regexp.MustCompile(``)

func ParseProfile(contents []byte) engine.ParseResult {
	profile := model.Profile{}
	
	age,err := strconv.Atoi(extractString(contents,ageRe))
	if err != nil {
		profile.Age = age
	}
	profile.Marriage = extractString(contents,marriageRe)



}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) > 2 {
		return string(match[1])
	}else {
		return ""
	}
}
