package main

var V = []FoodInfo{{FoodName: "bak chor mee", MerchantName: "eatzi gourment bistro", DetailedLocation: "safra yishun", PostalCode: 400345, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},

	{FoodName: "roti prata", MerchantName: "ah sik's curry", DetailedLocation: "tampines ave 7", PostalCode: 432555, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "soya sauce chicken", MerchantName: "sor you chicken", DetailedLocation: "admiralty st 22", PostalCode: 441534, Price: 4.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "rojak", MerchantName: "kim's rojak", DetailedLocation: "tampines ave 7", PostalCode: 451323, Price: 5.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "hai di lao", MerchantName: "hai di lao", DetailedLocation: "tampines ave 7", PostalCode: 462152, Price: 10.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "duck rice", MerchantName: "kay lee duck", DetailedLocation: "tampines ave 7", PostalCode: 475020, Price: 6.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "duck noodles", MerchantName: "little ducklings", DetailedLocation: "tampines ave 7", PostalCode: 482232, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "murtabak", MerchantName: "north india delight", DetailedLocation: "tampines ave 7", PostalCode: 491233, Price: 4.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "salted egg fried rice", MerchantName: "royal palm", DetailedLocation: "orchid country club", PostalCode: 503455, Price: 7.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "satay bee hoon", MerchantName: "woody family cafe", DetailedLocation: "sebawang dr 5", PostalCode: 512455, Price: 20.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "hokkien mee", MerchantName: "la pizzaiola", DetailedLocation: "yio chu kang st 21", PostalCode: 520393, Price: 15.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "fish bee hoon soup", MerchantName: "hualong fishhead steamboat", DetailedLocation: "ang mo kio st 52", PostalCode: 531234, Price: 6.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "ice cream sandwich", MerchantName: "citrus by the pool", DetailedLocation: "woodlands hub", PostalCode: 546287, Price: 5.20, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "prawn noodles", MerchantName: "the famous kitchnen", DetailedLocation: "khati dr 7", PostalCode: 551255, Price: 6.70, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "fried chicken wings", MerchantName: "wingz", DetailedLocation: "ang mo kio st 52", PostalCode: 565789, Price: 17.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "kaya toast", MerchantName: "toastbox", DetailedLocation: "ang mo kio st 52", PostalCode: 575789, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "cendol", MerchantName: "food republic", DetailedLocation: "ang mo kio st 52", PostalCode: 585789, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "durian", MerchantName: "geylang durian limited", DetailedLocation: "ang mo kio st 52", PostalCode: 601242, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "mutton soup", MerchantName: "west india cuisine", DetailedLocation: "ang mo kio st 52", PostalCode: 611242, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "nasi padang", MerchantName: "hualong fishhead steamboat", DetailedLocation: "ang mo kio st 52", PostalCode: 631242, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "claypot rice", MerchantName: "ah ming claypot shop", DetailedLocation: "ang mo kio st 52", PostalCode: 621242, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "nasi briyani", MerchantName: "nasi king", DetailedLocation: "ang mo kio st 52", PostalCode: 641242, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "orh luak", MerchantName: "old airport original", DetailedLocation: "ang mo kio st 52", PostalCode: 651242, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "curry chicken noodles", MerchantName: "kt original", DetailedLocation: "ang mo kio st 52", PostalCode: 421242, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "mee soto", MerchantName: "makcik's favorite", DetailedLocation: "ang mo kio st 52", PostalCode: 441242, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "popiah", MerchantName: "xiao li dishes", DetailedLocation: "ang mo kio st 52", PostalCode: 431242, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "muah chee", MerchantName: "xiao hua dishes", DetailedLocation: "ang mo kio st 52", PostalCode: 451242, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "chicken rice", MerchantName: "huat ah chicken", DetailedLocation: "ang mo kio st 52", PostalCode: 461242, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "tutu kueh", MerchantName: "kueh tutu from klang", DetailedLocation: "ang mo kio st 52", PostalCode: 471242, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "nonya kueh", MerchantName: "pernakan fine dining", DetailedLocation: "ang mo kio st 52", PostalCode: 481242, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "mee rebus", MerchantName: "sedap sedap", DetailedLocation: "ang mo kio st 52", PostalCode: 491242, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "chilli crab", MerchantName: "chinatown seafood limited", DetailedLocation: "ang mo kio st 52", PostalCode: 500393, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "sambal sting ray", MerchantName: "zi char in suntec", DetailedLocation: "sebawang dr 5", PostalCode: 510393, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "kway chap", MerchantName: "hua tiong small dishes", DetailedLocation: "yio chu kang st 21", PostalCode: 520393, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "putu piring", MerchantName: "shiok shiok", DetailedLocation: "ang mo kio st 52", PostalCode: 530393, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "fried carrot cake", MerchantName: "shiok die", DetailedLocation: "ang mo kio st 52", PostalCode: 540393, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "satay", MerchantName: "abang satay club", DetailedLocation: "ang mo kio st 52", PostalCode: 550393, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "bak kut teh", MerchantName: "clarke quay original bkt", DetailedLocation: "ang mo kio st 52", PostalCode: 560393, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "lor mee", MerchantName: "lor lor lor me", DetailedLocation: "ang mo kio st 52", PostalCode: 570393, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "sup tulung", MerchantName: "best soup in town", DetailedLocation: "ang mo kio st 52", PostalCode: 580393, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "fish head curry", MerchantName: "seafood paradise", DetailedLocation: "ang mo kio st 52", PostalCode: 590393, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "curry puff", MerchantName: "old chang kee", DetailedLocation: "ang mo kio st 52", PostalCode: 600393, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "laksa", MerchantName: "kallang laksa", DetailedLocation: "ang mo kio st 52", PostalCode: 610393, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "char kway teow", MerchantName: "tampnes char kway teow", DetailedLocation: "ang mo kio st 52", PostalCode: 620393, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "fish ball noodles", MerchantName: "dodo fishball shop", DetailedLocation: "ang mo kio st 52", PostalCode: 630393, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "pernakan food", MerchantName: "pernakan fine dining", DetailedLocation: "ang mo kio st 52", PostalCode: 640393, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "ban mian", MerchantName: "chiese shiok food", DetailedLocation: "ang mo kio st 52", PostalCode: 650393, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "pig organ soup", MerchantName: "exotic soup shop", DetailedLocation: "ang mo kio st 52", PostalCode: 630393, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "chwee kueh", MerchantName: "kuehhhh best", DetailedLocation: "ang mo kio st 52", PostalCode: 420393, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "chicken porridge", MerchantName: "auntie's porridge", DetailedLocation: "ang mo kio st 52", PostalCode: 450393, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "wanton mee", MerchantName: "potong pasir wanton mee", DetailedLocation: "ang mo kio st 52", PostalCode: 440393, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "korean bbq", MerchantName: "kim's bbq", DetailedLocation: "ang mo kio st 52", PostalCode: 440393, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
	{FoodName: "mala xiang guo", MerchantName: "crazy mlxg", DetailedLocation: "ang mo kio st 52", PostalCode: 595789, Price: 3.50, OpeningPeriods: OpeningPeriods{"monday": []string{"1345", "2300"}, "tuesday": []string{"1345", "2300"}, "wednesday": []string{"1345", "2300"}, "thursday": []string{"1345", "2300"}, "friday": []string{"1345", "2300"}, "saturday": []string{"1100", "2345"}, "sunday": []string{"1345", "2345"}, "public holiday": []string{"1100", "2345"}}},
}