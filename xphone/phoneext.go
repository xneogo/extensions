/*
 *  ┏┓      ┏┓
 *┏━┛┻━━━━━━┛┻┓
 *┃　　　━　　  ┃
 *┃   ┳┛ ┗┳   ┃
 *┃           ┃
 *┃     ┻     ┃
 *┗━━━┓     ┏━┛
 *　　 ┃　　　┃神兽保佑
 *　　 ┃　　　┃代码无BUG！
 *　　 ┃　　　┗━━━┓
 *　　 ┃         ┣┓
 *　　 ┃         ┏┛
 *　　 ┗━┓┓┏━━┳┓┏┛
 *　　   ┃┫┫  ┃┫┫
 *      ┗┻┛　 ┗┻┛
 @Time    : 2024/7/13 -- 14:32
 @Author  : bishop ❤️ MONEY
 @Description: phone_ext.go
*/

package xphone

var PhoneRegExpMap = map[string]*PhoneRegExp{
	ChinaMainland: &PhoneRegExp{
		AreaNumber: "86",
		RegexpAll:  "^((00|\\+){0,1}86-){0,1}((13[0-9])|(14[579])|(15[0-9])|(16[567])|(18[0-9])|(17[0-8])|(19[13589]))\\d{8}$",
		RegionCN:   "中国大陆",
		RegionCode: ChinaMainland,
	},
	"TEST": &PhoneRegExp{
		AreaNumber: "86",
		RegexpAll:  "^((00|\\+)?86-)?678\\d{8}$",
		RegionCN:   "中国大陆",
		RegionCode: "CN",
	},
	ChinaHongKong: &PhoneRegExp{
		AreaNumber: "852",
		RegexpAll:  "^((00|\\+){0,1}852-)(5|6|8|9)\\d{7}$",
		RegionCN:   "中国香港",
		RegionCode: ChinaHongKong,
	},
	ChinaMacao: &PhoneRegExp{
		AreaNumber: "853",
		RegexpAll:  "^((00|\\+){0,1}853-6)\\d{7}$",
		RegionCN:   "中国澳门",
		RegionCode: ChinaMacao,
	},
	ChinaTaiwan: &PhoneRegExp{
		AreaNumber: "886",
		RegexpAll:  "^((00|\\+){0,1}886-09)\\d{8}$",
		RegionCN:   "中国台湾",
		RegionCode: ChinaTaiwan,
	},
	UnitedArabEmirates: &PhoneRegExp{
		AreaNumber: "971",
		RegexpAll:  "^(00|\\+){0,1}971-0?5(0|2|5|6)\\d{7}$",
		RegionCN:   "阿联酋",
		RegionCode: UnitedArabEmirates,
	},
	Australia: &PhoneRegExp{
		AreaNumber: "61",
		RegexpAll:  "^(00|\\+){0,1}61-0?[45]\\d{8}$",
		RegionCN:   "澳大利亚",
		RegionCode: Australia,
	},
	TimorLeste: &PhoneRegExp{
		AreaNumber: "670",
		RegexpAll:  "^(00|\\+){0,1}670-7\\d{7}$",
		RegionCN:   "东帝汶",
		RegionCode: TimorLeste,
	},
	Philippines: &PhoneRegExp{
		AreaNumber: "63",
		RegexpAll:  "^(00|\\+){0,1}63-0?9\\d{9}$",
		RegionCN:   "菲律宾",
		RegionCode: Philippines,
	},
	Korea: &PhoneRegExp{
		AreaNumber: "82",
		RegexpAll:  "^(00|\\+){0,1}82-0?1\\d{9}$",
		RegionCN:   "韩国",
		RegionCode: Korea,
	},
	Canada: &PhoneRegExp{
		AreaNumber: "1",
		RegexpAll:  "^(00|\\+){0,1}1-[1-9]\\d{9}$",
		RegionCN:   "加拿大",
		RegionCode: Canada,
	},
	Kampuchea: &PhoneRegExp{
		AreaNumber: "855",
		RegexpAll:  "^(00|\\+){0,1}855-0?85\\d{6}$",
		RegionCN:   "柬埔寨",
		RegionCode: Kampuchea,
	},
	Laos: &PhoneRegExp{
		AreaNumber: "856",
		RegexpAll:  "^(00|\\+){0,1}856-0?20\\d{7,8}$",
		RegionCN:   "老挝",
		RegionCode: Laos,
	},
	Malaysia: &PhoneRegExp{
		AreaNumber: "60",
		RegexpAll:  "^(00|\\+){0,1}60-0?1(([02346789]\\d{7})|(1\\d{8}))$",
		RegionCN:   "马来西亚",
		RegionCode: Malaysia,
	},
	UnitedStates: &PhoneRegExp{
		AreaNumber: "1",
		RegexpAll:  "^(00|\\+){0,1}1-[1-9]\\d{9}$",
		RegionCN:   "美国",
		RegionCode: UnitedStates,
	},
	Myanmar: &PhoneRegExp{
		AreaNumber: "95",
		RegexpAll:  "^(00|\\+){0,1}95-0?[1-9]\\d{9}$",
		RegionCN:   "缅甸",
		RegionCode: Myanmar,
	},
	Japan: &PhoneRegExp{
		AreaNumber: "81",
		RegexpAll:  "^(00|\\+){0,1}81-0?[1-9]\\d{9}$",
		RegionCN:   "日本",
		RegionCode: Japan,
	},
	Thailand: &PhoneRegExp{
		AreaNumber: "66",
		RegexpAll:  "^(00|\\+){0,1}66-0?[1-9]\\d{8}$",
		RegionCN:   "泰国",
		RegionCode: Thailand,
	},
	Brunei: &PhoneRegExp{
		AreaNumber: "673",
		RegexpAll:  "^(00|\\+){0,1}673-((22[89])|(71\\d)|(72[0-3])|(8[1236789]\\d))\\d{4}$",
		RegionCN:   "文莱",
		RegionCode: Brunei,
	},
	Spain: &PhoneRegExp{
		AreaNumber: "34",
		RegexpAll:  "^(00|\\+){0,1}34-[6-7]\\d{8}$",
		RegionCN:   "西班牙",
		RegionCode: Spain,
	},
	Singapore: &PhoneRegExp{
		AreaNumber: "65",
		RegexpAll:  "^(00|\\+){0,1}65-[89]\\d{7}$",
		RegionCN:   "新加坡",
		RegionCode: Singapore,
	},
	NewZealand: &PhoneRegExp{
		AreaNumber: "64",
		RegexpAll:  "^(00|\\+){0,1}64-0?2[012579]\\d{7,8}$",
		RegionCN:   "新西兰",
		RegionCode: NewZealand,
	},
	Indonesia: &PhoneRegExp{
		AreaNumber: "62",
		RegexpAll:  "^(00|\\+){0,1}62-0?8\\d{8,10}$",
		RegionCN:   "印度尼西亚",
		RegionCode: Indonesia,
	},
	UnitedKingdom: &PhoneRegExp{
		AreaNumber: "44",
		RegexpAll:  "^(00|\\+){0,1}44-0?7\\d{9}$",
		RegionCN:   "英国",
		RegionCode: UnitedKingdom,
	},
	Vietnam: &PhoneRegExp{
		AreaNumber: "84",
		RegexpAll:  "^(00|\\+){0,1}84-0?[1-9]\\d{9}$",
		RegionCN:   "越南",
		RegionCode: Vietnam,
	},
	Germany: &PhoneRegExp{
		AreaNumber: "49",
		RegexpAll:  "^(00|\\+){0,1}49-0?\\d{3}\\s?\\d{8}$",
		RegionCN:   "德国",
		RegionCode: Germany,
	},
	India: &PhoneRegExp{
		AreaNumber: "91",
		RegexpAll:  "^(00|\\+){0,1}91-0?\\d{5}\\s?\\d{5}$",
		RegionCN:   "印度",
		RegionCode: India,
	},
	France: &PhoneRegExp{
		AreaNumber: "33",
		RegexpAll:  "^(00|\\+){0,1}33-0?\\d{2}\\s?\\d{8}$",
		RegionCN:   "法国",
		RegionCode: France,
	},
	Italy: &PhoneRegExp{
		AreaNumber: "39",
		RegexpAll:  "^(00|\\+){0,1}39-0?\\d{3}\\s?\\d{7}$",
		RegionCN:   "意大利",
		RegionCode: Italy,
	},
	Brazil: &PhoneRegExp{
		AreaNumber: "55",
		RegexpAll:  "^(00|\\+){0,1}55-0?\\d{2}\\s?\\d{5}\\s?\\d{4}$",
		RegionCN:   "巴西",
		RegionCode: Brazil,
	},
	Russia: &PhoneRegExp{
		AreaNumber: "7",
		RegexpAll:  "^(00|\\+){0,1}7-0?\\d{3}\\s?\\d{7}$",
		RegionCN:   "俄罗斯",
		RegionCode: Russia,
	},
	Nigeria: &PhoneRegExp{
		AreaNumber: "234",
		RegexpAll:  "^(00|0|\\+){0,1}234-0?\\d{11}$",
		RegionCN:   "尼日利亚",
		RegionCode: Nigeria,
	},
	Mexico: &PhoneRegExp{
		AreaNumber: "52",
		RegexpAll:  "^(00|\\+){0,1}52-0?\\d{2}\\s?\\d{3}\\s?\\d{4}$",
		RegionCN:   "墨西哥",
		RegionCode: Mexico,
	},
	Argentina: &PhoneRegExp{
		AreaNumber: "54",
		RegexpAll:  "^(00|\\+){0,1}54-0?\\d{2}\\s?\\d{4}\\s?\\d{4}$",
		RegionCN:   "阿根廷",
		RegionCode: Argentina,
	},
	Turkey: &PhoneRegExp{
		AreaNumber: "90",
		RegexpAll:  "^(00|\\+){0,1}90-0?\\d{2}\\s?\\d{3}\\s?\\d{4}$",
		RegionCN:   "土耳其",
		RegionCode: Turkey,
	},
	Ireland: &PhoneRegExp{
		AreaNumber: "353",
		RegexpAll:  "^(00|\\+){0,1}353-0?\\d{3}\\s?\\d{4}$",
		RegionCN:   "爱尔兰",
		RegionCode: Ireland,
	},
	Netherlands: &PhoneRegExp{
		AreaNumber: "31",
		RegexpAll:  "^(00|\\+){0,1}31-0?\\d{2}\\s?\\d{3}\\s?\\d{4}$",
		RegionCN:   "荷兰",
		RegionCode: Netherlands,
	},
	SaudiArabia: &PhoneRegExp{
		AreaNumber: "966",
		RegexpAll:  "^(00|\\+){0,1}966-0?\\d{3}\\s?\\d{7}$",
		RegionCN:   "沙特阿拉伯",
		RegionCode: SaudiArabia,
	},
	Switzerland: &PhoneRegExp{
		AreaNumber: "41",
		RegexpAll:  "^(00|\\+){0,1}41-0?\\d{2}\\s?\\d{3}\\s?\\d{2}\\s?\\d{2}$",
		RegionCN:   "瑞士",
		RegionCode: Switzerland,
	},
	Egypt: &PhoneRegExp{
		AreaNumber: "20",
		RegexpAll:  "^(00|\\+){0,1}20-0?\\d{2}\\s?\\d{8}$",
		RegionCN:   "埃及",
		RegionCode: Egypt,
	},
	Vietnam: &PhoneRegExp{
		AreaNumber: "84",
		RegexpAll:  "^(00|\\+){0,1}84-0?\\d{10}$",
		RegionCN:   "越南",
		RegionCode: Vietnam,
	},
	Malaysia: &PhoneRegExp{
		AreaNumber: "60",
		RegexpAll:  "^(00|\\+){0,1}60-0?\\d{2}\\s?\\d{8}$",
		RegionCN:   "马来西亚",
		RegionCode: Malaysia,
	},
	Poland: &PhoneRegExp{
		AreaNumber: "48",
		RegexpAll:  "^(00|\\+){0,1}48-0?\\d{2}\\s?\\d{8}$",
		RegionCN:   "波兰",
		RegionCode: Poland,
	},
	Israel: &PhoneRegExp{
		AreaNumber: "972",
		RegexpAll:  "^(00|\\+){0,1}972-0?\\d{2}\\s?\\d{7}$",
		RegionCN:   "以色列",
		RegionCode: Israel,
	},
	Colombia: &PhoneRegExp{
		AreaNumber: "57",
		RegexpAll:  "^(00|\\+){0,1}57-0?\\d{2}\\s?\\d{7}$",
		RegionCN:   "哥伦比亚",
		RegionCode: Colombia,
	},
	Peru: &PhoneRegExp{
		AreaNumber: "51",
		RegexpAll:  "^(00|\\+){0,1}51-0?\\d{2}\\s?\\d{7}$",
		RegionCN:   "秘鲁",
		RegionCode: Peru,
	},
	Chile: &PhoneRegExp{
		AreaNumber: "56",
		RegexpAll:  "^(00|\\+){0,1}56-0?\\d{2}\\s?\\d{8}$",
		RegionCN:   "智利",
		RegionCode: Chile,
	},
	Iraq: &PhoneRegExp{
		AreaNumber: "964",
		RegexpAll:  "^(00|\\+){0,1}964-0?\\d{2}\\s?\\d{8}$",
		RegionCN:   "伊拉克",
		RegionCode: Iraq,
	},
	Jordan: &PhoneRegExp{
		AreaNumber: "962",
		RegexpAll:  "^(00|\\+){0,1}962-0?\\d{2}\\s?\\d{8}$",
		RegionCN:   "约旦",
		RegionCode: Jordan,
	},
	Kazakhstan: &PhoneRegExp{
		AreaNumber: "7",
		RegexpAll:  "^(00|\\+){0,1}7-0?\\d{10}$",
		RegionCN:   "哈萨克斯坦",
	},
	Thailand: &PhoneRegExp{
		AreaNumber: "66",
		RegexpAll:  "^(00|\\+){0,1}66-0?[1-9]\\d{8}$",
		RegionCN:   "泰国",
		RegionCode: Thailand,
	},
	Egypt: &PhoneRegExp{
		AreaNumber: "20",
		RegexpAll:  "^(00|\\+){0,1}20-0?\\d{2}\\s?\\d{8}$",
		RegionCN:   "埃及",
		RegionCode: Egypt,
	},
	Romania: &PhoneRegExp{
		AreaNumber: "40",
		RegexpAll:  "^(00|\\+){0,1}40-0?\\d{2}\\s?\\d{7}$",
		RegionCN:   "罗马尼亚",
		RegionCode: Romania,
	},
	Serbia: &PhoneRegExp{
		AreaNumber: "381",
		RegexpAll:  "^(00|\\+){0,1}381-0?\\d{2}\\s?\\d{7}$",
		RegionCN:   "塞尔维亚",
		RegionCode: Serbia,
	},
	Finland: &PhoneRegExp{
		AreaNumber: "358",
		RegexpAll:  "^(00|\\+){0,1}358-0?\\d{2}\\s?\\d{7}$",
		RegionCN:   "芬兰",
		RegionCode: Finland,
	},
	Austria: &PhoneRegExp{
		AreaNumber: "43",
		RegexpAll:  "^(00|\\+){0,1}43-0?\\d{2}\\s?\\d{7}$",
		RegionCN:   "奥地利",
		RegionCode: Austria,
	},
	Belgium: &PhoneRegExp{
		AreaNumber: "32",
		RegexpAll:  "^(00|\\+){0,1}32-0?\\d{2}\\s?\\d{7}$",
		RegionCN:   "比利时",
		RegionCode: Belgium,
	},
	Spain: &PhoneRegExp{
		AreaNumber: "34",
		RegexpAll:  "^(00|\\+){0,1}34-0?\\d{2}\\s?\\d{8}$",
		RegionCN:   "西班牙",
		RegionCode: Spain,
	},
	Philippines: &PhoneRegExp{
		AreaNumber: "63",
		RegexpAll:  "^(00|\\+){0,1}63-0?\\d{10}$",
		RegionCN:   "菲律宾",
		RegionCode: Philippines,
	},
	Bulgaria: &PhoneRegExp{
		AreaNumber: "359",
		RegexpAll:  "^(00|\\+){0,1}359-0?\\d{2}\\s?\\d{7}$",
		RegionCN:   "保加利亚",
		RegionCode: Bulgaria,
	},
	Denmark: &PhoneRegExp{
		AreaNumber: "45",
		RegexpAll:  "^(00|\\+){0,1}45-0?\\d{2}\\s?\\d{7}$",
		RegionCN:   "丹麦",
		RegionCode: Denmark,
	},
	Ireland: &PhoneRegExp{
		AreaNumber: "353",
		RegexpAll:  "^(00|\\+){0,1}353-0?\\d{3}\\s?\\d{4}$",
		RegionCN:   "爱尔兰",
		RegionCode: Ireland,
	},
	Portugal: &PhoneRegExp{
		AreaNumber: "351",
		RegexpAll:  "^(00|\\+){0,1}351-0?\\d{2}\\s?\\d{9}$",
		RegionCN:   "葡萄牙",
		RegionCode: Portugal,
	},
	Sweden: &PhoneRegExp{
		AreaNumber: "46",
		RegexpAll:  "^(00|\\+){0,1}46-0?\\d{2}\\s?\\d{7}$",
		RegionCN:   "瑞典",
		RegionCode: Sweden,
	},
	Norway: &PhoneRegExp{
		AreaNumber: "47",
		RegexpAll:  "^(00|\\+){0,1}47-0?\\d{2}\\s?\\d{7}$",
		RegionCN:   "挪威",
		RegionCode: Norway,
	},
	Iceland: &PhoneRegExp{
		AreaNumber: "354",
		RegexpAll:  "^(00|\\+){0,1}354-0?\\d{7}$",
		RegionCN:   "冰岛",
		RegionCode: Iceland,
	},
}
