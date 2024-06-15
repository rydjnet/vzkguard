package antispam

var triggerWords = map[string]int{
	"ищу":              1,
	"ищем":             1,
	"заработаешь":      1,
	"поиске":           1,
	"набираю":          1,
	"команду":          1,
	"поисках":          1,
	"возьму":           1,
	"набор":            1,
	"предоставляем":    1,
	"напарников":       1,
	"нужны":            1,
	"предлагаю":        1,
	"заработка":        1,
	"партнерство":      1,
	"ответственных":    1,
	"umafic":           5,
	"английскому":      1,
	"испанскому":       1,
	"заинтересованных": 1,
	"нyжен":            1,
	"инструмент":       1,
	"поиска":           1,
	"клиентов":         1,
	"требуются":        1,
	"предложение":      1,
	"прибыльное":       1,
	"интересное":       1,
	"бизнесу":          1,
	"товарному":        1,
	"криптoвaлютe":     5,
	"поделюсь":         1,
	"интересoваных":    1,
	"взаимовыгодных":   1,
	"партнера":         1,
	"парнера":          2,
	"партнеров":        1,
	"людей":            1,
	"людях":            1,
	"удалённого":       1,
	"сотрудничества":   1,
	"инвайты":          1,
	"инвайт":           1,
	"люди":             1,
	"ассистента":       1,
	"валяешься":        1,
	"человек":          1,
	"ребят":            1,
	"проекта":          1,
	"возможность":      1,
	"график":           1,
	"проект":           1,
	"подработка":       1,
	"пассивно":         1,
	"получи":           1,
	"прибыль":          1,
	"крипте":           1,
	"гарантированно":   1,
	"сoтрудничествo":   5, // содржит латинские символы
	"дocтoйный":        5, // содржит латинские символы
	"сотрудничество":   1,
	"удалённое":        1,
	"удаленной":        1,
	"удаленному":       1,
	"удалённая":        1,
	"удаленке":         1,
	"удаленная":        1,
	"занятость":        1,
	"удаленном":        1,
	"удаленно":         1,
	"удаленную":        1,
	"удалённый":        1,
	"детали":           1,
	"пишите":           1,
	"лс":               1,
	"доход":            3,
	"финансовое":       1,
	"дохода":           3,
	"оплатой":          1,
	"прибыли":          1,
	"формате":          1,
	"дистанционном":    1,
	"оплата":           1,
	"доходом":          3,
	"доходе":           1,
	"crypto":           2,
	"bitget":           3,
	"mexc":             3,
	"BitGet":           3,
	"kucoin":           3,
	"вакансии":         1,
	"благонадёжного":   1,
	"добросовестного":  1,
	"добросовестных":   1,
	"вакансия":         1,
	"деньги":           1,
	"миллион":          1,
	"банкомат":         1,
	"купюpы":           1,
	"заработку":        1,
	"зарабатывать":     1,
	"oпытa":            5, // содржит латинские символы
	"обучeние":         5, // содржит латинские символы
	"поддeржка":        1, // содржит латинские символы
	"быстрoгo":         5, // содржит латинские символы
	"зapaбoтoк":        5, // содржит латинские символы
	"нужeн":            5, // содржит латинские символы
	"отдам":            5,
	"материал":         1,
	"вeдению":          5,
	"безвозмезднo":     5,
	"заинтересовaн":    5,
	"oбyчaющий":        5,
	"зарабоtок":        5,
	"подробнoсtи":      5,
	"прияtный":         5,
	"идet":             5, // содржит латинские символы
	"haбop":            5, // содржит латинские символы
	"людeй":            5, // содржит латинские символы
	"ha":               2, // содржит латинские символы
	"hobый":            5, // содржит латинские символы
	"пpoekt":           5, // содржит латинские символы
	"гpафик":           5, // содржит латинские символы
	"yдaлённo":         5, // содржит латинские символы
	"любoй":            5, // содржит латинские символы
	"мupa":             5, // содржит латинские символы
	"тoчкu":            5, // содржит латинские символы
	"быстpый":          5, // содржит латинские символы
	"pocт":             5, // содржит латинские символы
	"выcoкue":          5, // содржит латинские символы
	"дoxoды":           5, // содржит латинские символы
	"дeнь":             5, // содржит латинские символы
	"подpобности":      5, // содржит латинские символы
	"пoдробноcти":      5, // содржит латинские символы
	"мeста":            5, // содржит латинские символы
	"огpаничeны":       5, // содржит латинские символы
	"нeдeлю":           5, // содржит латинские символы
	"зa":               5, // содржит латинские символы
	"лc":               5, // содржит латинские символы
	"пoдрoбнoсти":      5, // содржит латинские символы
	"трейдингу":        2,
	"курс":             1,
	"зapa6отка":        5,
	"кoманду":          5,
	"зαραботоκ":        5,
	"сфеρα":            5,
	"чαсα":             5,
	"зα":               5,
	"нa":               5, // содржит латинские символы
	"нαм":              5,
	"плαтите":          5,
	"cтρого":           5,
	"вρемени":          5,
	"заинтересованых":  1,
	"заработок":        3,
	"фaльшивые":        1,
	"Р2Р":              3,
	"продажа":          1,
	"поддельных":       1,
	"купюp":            1,
	"казино":           1,
	"букмекерской":     1,
	"биржах":           1,
	"денег":            1,
	"pынке":            1,
	"чувак":            1,
	"$":                1,
	"18+":              1,
}
var whiteList = map[string]bool{
	"врача":            true,
	"няк":              true,
	"бк":               true,
	"докмед":           true,
	"гастроэнтерологи": true,
	"клиники":          true,
	"врачи":            true,
	"врач":             true,
	"колит":            true,
	"пациенты":         true,
	"стул":             true,
	"кровь":            true,
	"туалет":           true,
	"исследования":     true,
	"причину":          true,
	"болячка":          true,
	"болезнь":          true,
	"болячкой":         true,
	"боли":             true,
	"кишечник":         true,
	"заболеванием":     true,
	"заболевание":      true,
	"лекарства":        true,
	"диарея":           true,
	"срк":              true,
	"диагноз":          true,
	"диагноза":         true,
	"ГИБП":             true,
}

var triggerPhrases = []string{
	"нужны люди на удалённый заработок",
	"ищу партнеров в новый проект если интересно пишите в личные сообщения",
	"дocтoйный зapaбoтoк", //фраза с латинскими символами
	"бeз oпытa",           //фраза с латинскими символами
	"бeрём бeз oпытa",
	"предлагаю сотрудничество удалённо",
	"предлагаю сотрудничество",
	"сотрудничество удаленно",
	"сотрудничество удалённо",
	"всем совершеннолетним удалённое сотрудничество",
	"совершеннолетние люди на удалённое сотрудничество",
	"ответственных ребят с амбициями в новый проект",
	"удаленная деятельность",
	"сотрудничество на удаленке с хорошим доходом",
	"партнеров в команду для получения доп дохода",
	"250$ в день",
	"3000$ в месяц",
	"вариант заработка в удаленном формате",
	"проходят в любой бaнкомaт",
	"пoтенциaл дoхoдa",
	"надежного парнера",
	"надежного партнера",
	"людей на сотрудничество",
	"фaльшивые рубли",
	"фальшивые рубли",
	"поддержка для быстрoгo старта",
	"гарантированно заработаешь",
	"вариант заработка",
	"в новый проект",
	"нового проекта",
	"подробностями  пишите мне в личку",
	"подробности в лс",
	"за детальной информацией",
	"депозит в приложении",
	"на выгодных условиях",
	"программа лояльности",
	"букмекерской конторе",
	"ставок на спорт",
	"интересно пиши + в лс",
	"если интересно пишите в личные сообщения",
	"за деталями в лс",
	"деталями в личные сообщения",
	"за подробностями пишите мне в лс",
	"в лс или на основу",
	"в мире казино",
	"cтавки на спорт",
	"свободный график",
	"взаимовыгодного сотрудничества",
	"интересный проект",
	"в cфepe цифpoвыx вaлют",
	"удаленной основе",
	"sochishark",
	"дропаю туда различную инфу",
	"график плавающий",
	"заработать 1 миллион",
	"отличий от наcтоящих",
	"есть предложение за деталями",
	"по всем вопросам пишите",
	"возможность поднять финансовое положение",
	"за предложением в лс",
	"на удаленку",
	"за регистрацию на сайте",
	"подработка на удаленке",
	"работа удаленная",
	"кому интересно",
	"удалённая занятость",
	"высокой оплатой",
	"за информацией обращайтесь",
	"возможность зарабатывать",
	"зарабатывать пассивно",
	"у тебя совсем нету денег",
	"удалённого сотрудничества",
	"обращаться в личные сообщения",
	"подробностями в лс",
	"у себя в личке",
	"график работы свободный",
	"тысяч в день",
	"подробностями пиши",
	"минут твоего времени",
	"заработаешь свои",
	"подробностями ему в личку",
	"в доп прибыли",
	"возможность удалённого",
	"Тогда пиши \"+\"",
	"курс по ИИ",
	"поделюсь безвозмезднo",
	"нужен курс по",
	"кому-то нужен материал",
	"пишите - поделюсь",
	"товарному бизнесу",
	"инвайт в каналы",
	"крутая платформа",
	"деталями обращайтесь",
	"известных биржах",
	"работаем по готовым алгоритмам",
	"скину просто так",
}

var triggerEarns = []string{
	"день",
	"неделю",
	"мeсяц",
}
