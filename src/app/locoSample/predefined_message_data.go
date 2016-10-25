package main

type PredefinedMessageTranslation struct {
	Text              string
	PictureIdentifier string
	PictureWidth      int
	PictureHeight     int
}

var (
	predefineData = map[string]PredefinedMessageTranslation{
		"swipePromotion.-1000": PredefinedMessageTranslation{
			PictureIdentifier: "Predefine_Message_Swipe_Promotion_-1000",
			PictureWidth:      640,
			PictureHeight:     1136,
		},
		"swipePromotion.-1000.0": PredefinedMessageTranslation{
			Text: "Predefine_Message_Swipe_Promotion_-1000_0",
		},
		"swipePromotion.-1000.1": PredefinedMessageTranslation{
			PictureIdentifier: "Predefine_Message_Swipe_Promotion_-1000_1",
			PictureWidth:      640,
			PictureHeight:     1136,
		},
		"swipePromotion.-1001.0": PredefinedMessageTranslation{
			Text: "Predefine_Message_Swipe_Promotion_-1001_0",
		},
		"swipePromotion.-1002.0": PredefinedMessageTranslation{
			PictureIdentifier: "Predefine_Message_Swipe_Promotion_-1002_0",
			PictureWidth:      640,
			PictureHeight:     1136,
		},
		"teamAccount.Advertisement": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Advertisement",
		},

		"teamAccount.AnswerReply.1928": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_1928",
		},
		"teamAccount.AnswerReply.1929": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_1929",
		},
		"teamAccount.AnswerReply.1930": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_1930",
		},
		"teamAccount.AnswerReply.1932": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_1932",
		},
		"teamAccount.AnswerReply.1922": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_1922",
		},
		"teamAccount.AnswerReply.1934": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_1934",
		},
		"teamAccount.AnswerReply.1935": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_1935",
		},
		"teamAccount.AnswerReply.1936": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_1937",
		},
		"teamAccount.AnswerReply.1938": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_1938",
		},
		"teamAccount.AnswerReply.1939": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_1939",
		},
		"teamAccount.AnswerReply.1940": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_1940",
		},
		"teamAccount.AnswerReply.1941": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_1941",
		},
		"teamAccount.AnswerReply.1942": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_1942",
		},
		"teamAccount.AnswerReply.1943": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_1943",
		},
		"teamAccount.AnswerReply.1944": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_1944",
		},
		"teamAccount.AnswerReply.1945": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_1945",
		},
		"teamAccount.AnswerReply.1946": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_1946",
		},
		"teamAccount.AnswerReply.1947": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_1947",
		},
		"teamAccount.AnswerReply.1948": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_1948",
		},
		"teamAccount.AnswerReply.1949": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_1949",
		},
		"teamAccount.AnswerReply.1950": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_1950",
		},
		"teamAccount.AnswerReply.1951": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_1951",
		},
		"teamAccount.AnswerReply.1953": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_1953",
		},
		"teamAccount.AnswerReply.1954": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_1954",
		},
		"teamAccount.AnswerReply.1955": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_1955",
		},
		"teamAccount.AnswerReply.1957": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_1957",
		},
		"teamAccount.AnswerReply.1958": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_1958",
		},
		"teamAccount.AnswerReply.1960": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_1960",
		},
		"teamAccount.AnswerReply.1961": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_1961",
		},
		"teamAccount.AnswerReply.1962": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_1962",
		},
		"teamAccount.AnswerReply.1963": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_1963",
		},
		"teamAccount.AnswerReply.1964": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_1964",
		},
		"teamAccount.AnswerReply.1965": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_1965",
		},
		"teamAccount.AnswerReply.1966": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_1966",
		},

		"teamAccount.AnswerReply.2180": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_2180",
		},
		"teamAccount.AnswerReply.2181": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_2181",
		},
		"teamAccount.AnswerReply.2182": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_2182",
		},
		"teamAccount.AnswerReply.2184": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_2184",
		},
		"teamAccount.AnswerReply.2185": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_2185",
		},
		"teamAccount.AnswerReply.2186": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_2186",
		},
		"teamAccount.AnswerReply.2187": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_2187",
		},
		"teamAccount.AnswerReply.2188": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_2188",
		},
		"teamAccount.AnswerReply.2191": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_2191",
		},
		"teamAccount.AnswerReply.2192": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_2192",
		},
		"teamAccount.AnswerReply.2193": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_2193",
		},
		"teamAccount.AnswerReply.2195": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_2195",
		},
		"teamAccount.AnswerReply.2196": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_2196",
		},
		"teamAccount.AnswerReply.2197": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_2197",
		},
		"teamAccount.AnswerReply.2198": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_2198",
		},
		"teamAccount.AnswerReply.2199": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_2199",
		},
		"teamAccount.AnswerReply.2200": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_2200",
		},
		"teamAccount.AnswerReply.2201": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_2201",
		},
		"teamAccount.AnswerReply.2202": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_2202",
		},
		"teamAccount.AnswerReply.2203": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_2203",
		},
		"teamAccount.AnswerReply.2205": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_2205",
		},
		"teamAccount.AnswerReply.2206": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_2206",
		},
		"teamAccount.AnswerReply.2207": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_2207",
		},
		"teamAccount.AnswerReply.2209": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_2209",
		},
		"teamAccount.AnswerReply.2210": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_2210",
		},
		"teamAccount.AnswerReply.2212": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_2212",
		},
		"teamAccount.AnswerReply.2213": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_2213",
		},
		"teamAccount.AnswerReply.2214": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_2214",
		},
		"teamAccount.AnswerReply.2215": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_2215",
		},
		"teamAccount.AnswerReply.2216": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_2216",
		},
		"teamAccount.AnswerReply.2217": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_2217",
		},
		"teamAccount.AnswerReply.2218": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Answer_Reply_2218",
		},

		"teamAccount.BadProfile": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Bad_Profile",
		},
		"teamAccount.BannedAdvertisement": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Banned_Advertisement",
		},
		"teamAccount.BannedPornProfilePicture": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Banned_Porn_Profile_Picture",
		},
		"teamAccount.DirtyDating": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Dirty_Dating",
		},
		"teamAccount.FakeProfilePicture": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Fake_Profile_Picture",
		},
		"teamAccount.GoodWarning": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Good_Warning",
		},
		"teamAccount.HiddenUserDefault": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Hidden_User_Default",
		},
		"teamAccount.MomentDeleted": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Moment_Deleted",
		},
		"teamAccount.MomentDeletedByMistake": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Moment_Deleted_By_Mistake",
		},
		"teamAccount.MomentUserAdsBanned": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Moment_User_Ads_Banned",
		},
		"teamAccount.MomentUserSexualContentBanned": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Moment_User_Sexual_Content_Banned",
		},
		"teamAccount.PokeReply": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Poke_Reply",
		},
		"teamAccount.ReportedUserBanned": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Reported_User_Banned",
		},
		"teamAccount.SpamMsg": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Spam_Msg",
		},
		"teamAccount.StolenAccount": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Stolen_Account",
		},
		"teamAccount.UglyAdvertisement": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Ugly_Advertisement",
		},
		"teamAccount.UglyONS": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Ugly_ONS",
		},
		"teamAccount.UglyPornProfilePicture": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Ugly_Porn_Profile_Picture",
		},
		"teamAccount.UglyProfilePicture": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Ugly_Profile_Picture",
		},
		"teamAccount.Welcome": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Welcome",
		},
		"teamAccount.WrongGender": PredefinedMessageTranslation{
			Text: "Predefine_Message_Team_Account_Wrong_Gender",
		},
	}
)
