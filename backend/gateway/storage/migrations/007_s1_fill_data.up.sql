INSERT INTO `users` (`id`, `role_id`, `login`, `password`, `first_name`, `last_name`, `patronymic`) VALUES
    (1, 1, 'Example', '$2a$10$tPeq/dedUCQ4vUsAZJgmp.kN2Bn4vfdbZLNlOKCmDjphhVG0zMX8K', '', '', ''),
    (2, 1, 'Director', '$2a$10$GiAo62sOGxF9O6HWqTKGkOotIxK9i00/nrY/Nb6UXR6iNQ9gnazFi', '', '', ''),
    (3, 1, 'TestCustomer', '$2a$10$f/CmKXLDrB2Mb2Kcr4.3H.4tqoc2dd.kAIVukxllcaq.g1D1.MTuq', '', '', ''),
    (4, 1, 'loginDEsar2018', '$2a$10$fC.UU1N55l6qXbEn/LGgxuZggwtdTRc9WOlMgVbp4HVNBGUeV2AjG', 'Ефим', 'Яковлев', 'Ильяович'),
    (5, 3, 'loginDEvdj2018', '$2a$10$cyrjEw5vnaIwrtkgyz733u6qv0vsesZhNMWxT/LLLnK4873yAZJ06', 'Богдан', 'Никонов', 'Юлианович'),
    (6, 2, 'loginDEyyj2018', '$2a$10$30qepXOeHDsZOmL5wZUAIO31wRb1AJTKzsBPDmHib1CjWAY/V4aEC', 'Валерьян', 'Лебедев', 'Кимович'),
    (7, 4, 'loginDEvva2018', '$2a$10$SE2BRnL55X8Ex5PZMN2WFOSQUTyjt/d836UhlOkiJJMa75Pl3Ozau', 'Александра', 'Ермакова ', 'Куприяновна'),
    (8, 5, 'loginDEepk2018', '$2a$10$Dbbb/LnJh3OQ5Nne0NzpQOn94aQm4BcjQA3I9LdGASFX0q3iqxZtu', 'Гордей', 'Петухов', 'Серапионович');

INSERT INTO `suppliers` 
VALUES 
	(1, "ИП Дмитрий Delivery", "г. Москва ул. Шарикоподшипниковская 42а", 120),
    (2, "Fast Express", "Московская область, г. Красногорск, ул. Маяковская д.3", 300);

insert into `tool_types` (`id`, `name`) values
    (1, "Формы"),
    (2, "Ножи"),
    (3, "Кастрюли")
    ;

INSERT INTO `tools` (`id`, `type_id`, `supplier_id`, `name`, `description`, `degree_of_wear_id`, `purchase_date`) 
VALUES
    (1, 1, 1, 'Форма 1', '', 1, '2024-11-14 12:33:06');

INSERT INTO `component_types` VALUES
    (1, 1, "Премиксы для выпечки"),
    (2, 1, "Сливки"),
    (3, 1, "Фруктовое пюре"),
    (4, 1, "Десертная паста"),
    (5, 1, "топпинг"),
    (6, 1, "Цукаты"),
    (7, 1, "Глазурь"),
    (8, 1, "Шоколад"),
    (9, 1, "Какао"),
    
    (10, 2, "Выпеченные полуфабрикаты"),
    (11, 2, "Шарики желейные"),
    (12, 2, "шоколад"),
    (13, 2, "Посыпка")
    ;

INSERT INTO `components` (
    `article`, 
    `name`, 
    `measure_unit`, 
    `component_type_id`
)
VALUES ("76003","Мука миндальная ПРЕМИУМ","кг",1), ("74007","Комп.пищ.добавка премикс САВОЯРДИ микс","кг",1), ("74009","Комп.пищ.добавка премикс ОВОСИЛ","кг",1), ("74002","Комп.пищ.добавка премикс БИСКВИЗИТ","кг",1), ("74004","Смесь для мучных изделий КЕЙК ЗЛАКИ","кг",1), ("74003","Комп.пищ.добавка премикс МУРБЕЛЛА","кг",1), ("74006","Смесь для мучных изделий ФРОЛЛА ЗЛАКИ","кг",1), ("74081","Смесь сахаристая ГРИЛЬЯЖ","кг",1), ("74046","Комп.пищ.доб.премикс МАФФИН","кг",1), ("74048","Премикс БРАУНИ ШОК","кг",1), ("74050","Комп.пищ.добавка премикс ВАФЛИ/ОЛАДЬИ","кг",1), ("74049","Комп.пищ.добавка премикс АВОЛЕТТА макарунс","кг",1), ("73014","Крем на растит.маслах КОНДИТЕРСКИЙ 26%","л",2), ("73018","Сливки Бризе 30%","л",2), ("73023","Сливки Бризе 36%","л",2), ("73015","Сливки 33%","л",2), ("73016","Сливки 35%","л",2), ("73019","Крем на растит.маслах для взбивания 26%","л",2), ("73067","Комп.пищ.добавка крем холод.приготов.ПЕРФЕКТА","кг",2), ("73069","Комп.пищ.добавка крем горяч.приготовл.СОВРАНА","кг",2), ("73078","Комплексная пищевая добавка премикс ТИРАМИСУ","кг",2), ("73079","Комплексная пищевая добавка премикс БАВАРЕЗЕ","кг",2), ("73080","Комплексная пищ.доб.премикс ПАННА КОТТА","кг",2), ("73082","Крем десертный МУСС КАКАО","кг",2), ("73084","Комп.пищ.доб.крем десертный классич.ЧИЗКЕЙК","кг",2), ("73088","Комп.пищ.доб.крем дес.АМЕРИК.ЧИЗКЕЙК","кг",2), ("76032","Премикс горячий шоколад ШОКОЛАДНАЯ ЧАШКА","кг",2), ("72860","Пюре фруктовое МАНГО","кг",3), ("75144","Паста ПРАЛИНЕ ФИСТАШКА","кг",3), ("71011","Глазурь какаосодерж.кондит.ЗАХЕРКРЕМ","кг",3), ("75046","Глазурь пастообразн.конд.ШОКОСМАРТ","кг",3), ("75045","Глазурь пастообразная жир.ШОКОСМАРТ белая","кг",3), ("71109","Шоколадная начинка ВАНДЕРШОК белая","кг",3), ("72799","Начинка КОБЕРШОК БЬЯНКО","кг",3), ("71123","Глазурь пастообразная ГАБРИЭЛА БЕЛАЯ","кг",3), ("71124","Глазурь пастообразная шок.ГАБРИЭЛА ТЕМНАЯ","кг",3), ("75400","Паста шок.орех.ФАРЧИТОЗА ТОП","кг",3), ("75401","Паста шок.орех.ФАРЧИТОЗА ТОП","кг",3), ("75412","Паста шок.орех.ФАРЧИТОЗА","кг",3), ("75402","Паста шок.орех.НОЧЧОЛИТА СУПРЕМ","кг",3), ("75403","Паста шок.орех.НОЧЧОЛИТА СУПРЕМ","кг",3), ("75413","Паста шок.орех.ФАРЧИТОЗА","кг",3), ("75414","Паста шок.орех.ФАРЧИТОЗА","кг",3), ("75405","Паста шок.орех.НОЧЧОЛИТА ТОП","кг",3), ("75406","Паста шок.орех.НОЧЧОЛИТА ТОП","кг",3), ("75408","Паста ореховая НОЧЧИОБЬЯНКА КРУАССАН","кг",3), ("75407","Паста-наполнитель НОВЕЛЛА АВОРИО","кг",3), ("75000","Паста-наполнитель РОКСИКРЕМ","кг",3), ("75002","Паста шок.орех.ДЖАНДУЯ КРУАССАН термостойкая","кг",3), ("75409","Паста ореховая ФИСТАШКА КРЕМ 10","кг",3), ("75410","Паста ореховая ФИСТАШКА КРЕМ ТОП 15","кг",3), ("75411","Паста ореховая ФИСТАШКА КРУАССАН","кг",3), ("75001","Глазурь пастообразн.конд.ШОКОКРЕМ","кг",3), ("75005","Глазурь пастообразн.конд.ШОКОКРЕМ","кг",3), ("75007","Глазурь пастообр.шок.МОРЕТТА БИТТЕР термост.","кг",3), ("75009","Паста шоколадная МОРЕТТА ЭКСТРИМ","кг",3), ("75011","Паста шоколадная МОРЕТТА ЭКСТРИМ","кг",3), ("72707","Паста шоколадная КРАНФИЛ КАРАМЕЛЬ","кг",3), ("72708","Паста шоколадная КРАНФИЛ ФИСТАШКА","кг",3), ("72709","Паста шоколадная КРАНФИЛ БЕЛЫЙ ШОКОЛАД","кг",3), ("72710","Паста шоколадная КРАНФИЛ ТЕМНЫЙ ШОКОЛАД","кг",3), ("72723","Паста шоколадная ДЕЛИКРИСП соленая карамель","кг",3), ("72724","Паста шоколадная ДЕЛИКРИСП кокос","кг",3), ("72725","Паста шоколадная ДЕЛИКРИСП красные ягоды","кг",3), ("75014","Паста ПРАЛИНЕ ГРЕЦКИЙ ОРЕХ","кг",3), ("75012","Паста ПРАЛИНЕ ДЖАНДУЯ","кг",3), ("75021","Паста ПРАЛИНЕ АРАХИС-ШОКОЛАД","кг",3), ("75119","Паста сахарно-ореховая ПРАЛИНЕ (НУГА) светлая","кг",3), ("75022","Паста ПРАЛИНЕ МИНДАЛЬ-МЯТА","кг",3), ("71237","Начинка ФРУТТИДОР ЯБЛОКО","кг",3), ("72796","Начинка ФРУТТИДОР АБРИКОС","кг",3), ("72795","Начинка ФРУТТИДОР ЧЕРНИКА","кг",3), ("72798","Начинка ФРУТТИДОР АПЕЛЬСИН","кг",3), ("72797","Начинка ФРУТТИДОР ПЕРСИК","кг",3), ("72727","Начинка ФРУТТИДОР АМАРЕНА","кг",3), ("72728","Начинка ФРУТТИДОР ЛЕСНЫЕ ЯГОДЫ","кг",3), ("72729","Начинка ФРУТТИДОР ГРУША","кг",3), ("72730","Начинка ФРУТТИДОР МАНГО","кг",3), ("72731","Начинка ФРУТТИДОР ТРОПИЧЕСКИЕ ФРУКТЫ","кг",3), ("72732","Начинка ФРУТТИДОР КЛУБНИКА","кг",3), ("72733","Начинка ФРУТТИДОР МАЛИНА","кг",3), ("72734","Начинка ФРУТТИДОР КРАСНАЯ ЧЕРЕШНЯ","кг",3), ("75036","Начинка КРЕМ ЛИМОН","кг",3), ("75037","Начинка КРЕМ АПЕЛЬСИН","кг",3), ("75038","Начинка КРЕМ ЛЕСНЫЕ ЯГОДЫ И ЙОГУРТ","кг",3), ("75041","Начинка фруктовая ЯБЛОКО ПИНК кубики 12х12 мм","кг",3), ("75073","Начинка фруктовая ЯБЛОКО ПИНК дольки","кг",3), ("75056","Начинка фруктовая ЯБЛОКО дольки","кг",3), ("75170","Начинка фрукт.КОНФРУТТИ черника","кг",3), ("75173","Начинка фрукт.КОНФРУТТИ абрикос","кг",3), ("75175","Паста-пюре фрукт.КОНФРУТТИ НАТУР клубника","кг",3), ("75176","Паста-пюре фрукт.КОНФРУТТИ НАТУР малина","кг",3), ("75177","Паста-пюре фрукт.КОНФРУТТИ НАТУР садов.ягоды","кг",3), ("75178","Паста-пюре фрукт.КОНФРУТТИ НАТУР манго","кг",3), ("75179","Паста-пюре фрукт.КОНФРУТТИ НАТУР маракуйя","кг",3), ("75180","Паста-пюре фрукт.КОНФРУТТИ НАТУР банан","кг",3), ("75181","Паста-пюре фрукт.КОНФРУТТИ НАТУР ананас","кг",3), ("75190","Паста-пюре фрукт.КОНФРУТТИ НАТУР вишня Гриот","кг",3), ("75191","Паста-пюре фрукт.КОНФРУТТИ НАТУР мандарин","кг",3), ("75192","Паста-пюре фрукт.КОНФРУТТИ НАТУР груша Вильям","кг",3), ("70942","Паста кондитерская ДЕКОР СЕРЕБРО","кг",4), ("70943","Паста кондитерская ДЕКОР ЗОЛОТО","кг",4), ("70944","Паста кондитерская ДЕКОР БРОНЗА","кг",4), ("74084","Паста РОМ КРЕОЛЬСКИЙ АРОМА","кг",4), ("74083","Паста ГОРЬКИЙ МИНДАЛЬ АРОМА","кг",4), ("72241","Паста ВАНИЛЬ АРОМА","кг",4), ("72242","Паста КОФЕ АРОМА","кг",4), ("72245","Паста ЛИМОН АРОМА","кг",4), ("72247","Паста КЛУБНИКА АРОМА","кг",4), ("72248","Паста ФИСТАШКА АРОМА","кг",4), ("72021","Комплексная пищ.добавка премикс ЙОГУРТ-100","кг",4), ("72018","Комплексная пищевая добавка премикс КОКОС-100","кг",4), ("72718","Паста десертная КОКОС","кг",4), ("72234","Паста десертная ЛЕСНОЙ ОРЕХ","кг",4), ("72704","Паста десертная МАСКАРПОНЕ","кг",4), ("72253","Паста десертная ФИСТАШКА","кг",4), ("75204","Паста десертная ФИСТАШКА ИТАЛИЯ НК","кг",4), ("72254","Паста десертная ГРЕЦКИЙ ОРЕХ","кг",4), ("72256","Паста десертная ПЕЧЕНЬЕ","кг",4), ("72258","Паста десертная СЛИВОЧНАЯ КАРАМЕЛЬ","кг",4), ("72263","Паста десертная ВАНИЛЬ КЛАССИК","кг",4), ("72705","Паста десертная ЗАБАЙОНЕ","кг",4), ("72274","Паста десертная МЯТА","кг",4), ("72289","Паста десертная ШОКОЛАД","кг",4), ("72287","Паста десертная КАРАМЕЛЬ","кг",4), ("72702","Паста десертная ТИРАМИСУ","кг",4), ("72719","Паста десертная ЖЕВАТЕЛЬНАЯ РЕЗИНКА","кг",4), ("72521","Паста десертная ЛЕСНАЯ ЯГОДА","кг",4), ("72683","Паста десертная МАЛИНА","кг",4), ("72684","Паста десертная ПЕРСИК с кусочками","кг",4), ("72685","Паста десертная МАРАКУЙЯ","кг",4), ("72695","Паста десертная КЛУБНИКА","кг",4), ("72696","Паста десертная ДЫНЯ с кусочками","кг",4), ("72697","Паста десертная ЗЕЛЕНОЕ ЯБЛОКО с кусочками","кг",4), ("72698","Паста десертная АНАНАС с кусочками","кг",4), ("72699","Паста десертная БАНАН","кг",4), ("72700","Паста десертная МАНГО с кусочками","кг",4), ("72721","Паста десертная ЛИМОННАЯ ЦЕДРА","кг",4), ("72722","Паста десертная АПЕЛЬСИНОВАЯ ЦЕДРА","кг",4), ("72146","Наполнитель МОЛОЧНЫЙ ШОКОЛАД и злаки","кг",5), ("72147","Наполнитель ГРАНАТ","кг",5), ("72713","Топпинг КИВИ","кг",5), ("72706","Наполнитель шоколадный МАКАО","кг",5), ("72712","Наполнитель ВИШНЯ","кг",5), ("72720","Наполнитель ШОКОЛАД С ПЕЧЕНЬЕМ","кг",5), ("72593","Паста десертная АБРИКОС","кг",4), ("72801","Паста десертная ЗЕМЛЯНИКА","кг",4), ("72802","Паста десертная ВИШНЯ","кг",4), ("72522","Паста десертная ГРУША","кг",4), ("72041","Паста десертная ПРЯНОЕ ПЕЧЕНЬЕ","кг",4), ("72120","Паста десертная БИТТЕР ЛИМОН","кг",4), ("72150","Паста десертная МЯТА","кг",4), ("72151","Паста десертная МАНДАРИН","кг",4), ("72506","Паста десертная ВАНИЛЬ","кг",4), ("72507","Паста десертная ВАНИЛЬ БУРБОН","кг",4), ("72508","Паста десертная КАРАМЕЛЬ","кг",4), ("72514","Паста десертная КАПУЧИНО","кг",4), ("72524","Паста десертная КЛУБНИКА","кг",4), ("72525","Паста десертная МАНГО","кг",4), ("72526","Паста десертная АНАНАС","кг",4), ("72527","Паста десертная ДЫНЯ","кг",4), ("72528","Паста десертная КИВИ","кг",4), ("72539","Паста десертная ЧЕРНАЯ СМОРОДИНА","кг",4), ("72562","Паста десертная БАНАН","кг",4), ("72564","Паста десертная АНАНАС","кг",4), ("72565","Паста десертная МАЛИНА","кг",4), ("72577","Паста десертная ЛЕСНЫЕ ЯГОДЫ","кг",4), ("72587","Паста десертная ПЕРСИК","кг",4), ("72591","Паста десертная МАРАКУЙЯ","кг",4), ("72595","Паста десертная РОЗА","кг",4), ("72596","Паста десертная АПЕЛЬСИН","кг",4), ("72599","Паста десертная МАЛАГА","кг",4), ("72555","Паста десертная ПИНА КОЛАДА","кг",4), ("72575","Паста десертная ВИСКИКРЕМ","кг",4), ("72543","Паста десертная БИСКОТТО","кг",4), ("72579","Паста десертная АМАЛЬФИ","кг",4), ("72542","Паста десертная НЬЮ ЙО","кг",4), ("72515","Паста десертная ТИРАМИСУ","кг",4), ("72548","Паста десертная МЕКРАФ","кг",4), ("72553","Паста десертная МИСТЕР НИКО","кг",4), ("72552","Паста десертная НОЧЧИОЛА-ЛА-ЛА","кг",4), ("72554","Паста десертная ФИСТАШКА СИЦИЛИЙСКАЯ","кг",4), ("72556","Паста десертная ФИСТАШКА ЧИСТАЯ","кг",4), ("72572","Паста десертная ЛАЗУРЬ","кг",4), ("72574","Паста десертная БУБЛЬГУМ","кг",4), ("72537","Паста десертная ПЕРСИК","кг",4), ("72576","Паста десертная ЗЕЛЕНЫЙ ЧАЙ","кг",4), ("72516","Паста десертная ГРУША","кг",4), ("72582","Паста десертная МАНДАРИН","кг",4), ("72566","Паста десертная ГУСТОКРЕМ ЛАЙМЕТТА","кг",4), ("72567","Паста десертная ГУСТОКРЕМ АПЕЛЬСИН","кг",4), ("72563","Паста десертная ЧЕРРИ МАНИЯ","кг",4), ("72580","Паста десертная РОЗОВЫЙ ГРЕЙПФРУТ","кг",4), ("72585","Паста десертная ТРОПИК","кг",4), ("72637","Паста десертная ЗЕЛЕНОЕ ЯБЛОКО","кг",4), ("84080","Черешня красная МИНИ 14/16 мм.НАППИ","кг",6), ("84129","Черешня засах.красная 20/22 мм.с веточкой","кг",6), ("84131","Черешня засах.красная 20/22 мм.с веточкой","кг",6), ("84134","Черешня засах.красная 18/20 мм ИТАЛКАНДИТИ","кг",6), ("84081","Черешня засах.красная 20/22 мм.НАППИ","кг",6), ("84017","Черешня засах.красная 20/22 мм.АМБРОЗИО","кг",6), ("84014","Черешня засах.черная 20/22 мм.АМБРОЗИО","кг",6), ("84019","Черешня засах.зеленая 20/22 мм.АМБРОЗИО","кг",6), ("84018","Черешня засах.желтая 20/22 мм.АМБРОЗИО","кг",6), ("84133","Апельсиновые засах. дольки ИТАЛКАНДИТИ","кг",6), ("84023","Апельсин.корочка засах.кубик.6х6 мм.АМБРОЗИО","кг",6), ("84088","Апельсиновая корочка засах.куб.6х6 мм.НАППИ","кг",6), ("84086","Апельсиновая корочка засах.куб.4х4 мм.НАППИ","кг",6), ("84075","Апельсин.корочка засах.кубик.4х4 мм.АМБРОЗИО","кг",6), ("84027","Лимонная.корочка засах.кубик.6х6 мм.АМБРОЗИО","кг",6), ("84085","Лимонная корочка засах.кубик 6х6 мм.НАППИ","кг",6), ("84084","Лимонная корочка засах.кубик 4х4 мм.НАППИ","кг",6), ("84074","Лимонная.корочка засах.кубик.4х4 мм.АМБРОЗИО","кг",6), ("84060","Апельсин.корочка засах.полоска.6х60 мм.","кг",6), ("84061","Лимонная корочка засах.полоска.6х60 мм.","кг",6), ("84125","Фрукты засах.ассорти кубик.6х6 мм","кг",6), ("84057","Фрукты засах.разноцвет.кубик.6х6 мм.АМБРОЗИО","кг",6), ("84073","Фрукты засах.разноцвет.кубик.4х4 мм.АМБРОЗИО","кг",6), ("84041","Мандарин засах.целый АМБРОЗИО","кг",6), ("84043","Груша засах.целая красная АМБРОЗИО","кг",6), ("84047","Груша засах.целая зеленая АМБРОЗИО","кг",6), ("84049","Фрукты засах.целые ассорти АМБРОЗИО","кг",6), ("84076","Черешня АМАРЕНА в сиропе с веточкой 20/22 мм.","кг",6), ("84078","Черешня АМАРЕНА в сиропе 20/22 мм.","кг",6), ("84126","Черешня засах.АМАРЕНА ЭКСТРА14/16 мм","кг",6), ("84062","Лимонная корочка засах.кубик.9х9 мм.ЭКСТРА","кг",6), ("84026","Апельсин.корочка засах.кубик.9х9 мм.ЭКСТРА","кг",6), ("84002","Черешня с вет.коктейльная красн.Мараскино","кг",6), ("84124","Черешня КОКТЕЙЛЬ с веточкой красн.20/22 мм.","кг",6), ("84120","Черешня КОКТЕЙЛЬ с веточкой красн.18/20 мм.","кг",6), ("84001","Черешня с вет.коктейльная красн.Мараскино","кг",6), ("84148","Черешня КОКТЕЙЛЬ с веточкой красн.","кг",6), ("84147","Черешня КОКТЕЙЛЬ с веточкой красн.","кг",6), ("84149","Черешня КОКТЕЙЛЬ с веточкой желтая","кг",6), ("84150","Черешня КОКТЕЙЛЬ с веточкой зеленая","кг",6), ("84151","Черешня КОКТЕЙЛЬ с веточкой красн. НК","кг",6), ("84153","Черешня КОКТЕЙЛЬ с веточкой красн.","кг",6), ("84154","Черешня КОКТЕЙЛЬ с веточкой красн.","кг",6), ("84163","Черешня КОКТЕЙЛЬ с веточкой зелен.","кг",6), ("84173","Черешня КОКТЕЙЛЬ с веточкой желтая","кг",6), ("71152","Глазурь БРИЛЛО КЛУБНИКА","кг",7), ("71119","Глазурь БРИЛЛО КАРАМЕЛЬ","кг",7), ("71117","Глазурь БРИЛЛО ШОКОЛАД","кг",7), ("71118","Глазурь БРИЛЛО БЕЛЫЙ ШОКОЛАД","кг",7), ("70006","Глазурь желейная БРИЛГЕЛЬ","кг",7), ("70001","Глазурь желейная БРИЛГЕЛЬ","кг",7), ("70022","Глазурь желейная ФРИО","кг",7), ("70043","Глазурь желейная КОЛДГЕЛЬ NEW","кг",7), ("70101","Глазурь желейная ДЕКОРГЕЛЬ нейтральная","кг",7), ("70325","Глазурь желейная абрикосовая","кг",7), ("70401","Глазурь желейная НОВОГЕЛЬ нейтральная","кг",7), ("70406","Глазурь желейная НОВОГЕЛЬ нейтральная","кг",7), ("70408","Глазурь желейная НОВОГЕЛЬ нейтральная","кг",7), ("70026","Глазурь желейная РОЯЛГЕЛЬ","кг",7), ("70937","Глазурь гляссажная КОВЕРГЛАС ШОКОЛАД","кг",7), ("70947","Глазурь гляссажная КОВЕРГЛАС ЛАЙМ","кг",7), ("70955","Глазурь гляссажная КОВЕРГЛАС КАРАМЕЛЬ МУ","кг",7), ("70926","Глазурь мягкая МИРУАР АПЕЛЬСИН","кг",7), ("70927","Глазурь мягкая МИРУАР ВАНИЛЬ","кг",7), ("70928","Глазурь мягкая МИРУАР КЛУБНИКА","кг",7), ("70929","Глазурь мягкая МИРУАР ФИСТАШКА","кг",7), ("70930","Глазурь желейная МИРУАР ШОКОЛАД","кг",7), ("71100","Шоколад темный 36/38","кг",8), ("71101","Шоколад молочный 30/32","кг",8), ("71102","Шоколад белый 32/34","кг",8), ("71103","Шоколад темный 36/38","кг",8), ("71104","Шоколад молочный 30/32","кг",8), ("71110","Шоколад белый 32/34","кг",8), ("71112","Шоколад темный ШОКО НЕРО 72%","кг",8), ("71113","Шоколад темный ШОКО НЕРО 52%","кг",8), ("71114","Шоколад молочный ШОКО ЛАТТЕ 30%","кг",8), ("71116","Шоколад белый ШОКО БЬЯНКО 34/36","кг",8), ("71142","Посыпка шоколадная КРОШКА ТЕМНАЯ","кг",8), ("71132","Посыпка шоколадная КРОШКА ТЕМНАЯ","кг",8), ("71143","Посыпка шоколадная КРОШКА МОЛОЧНАЯ","кг",8), ("71144","Посыпка шоколадная КРОШКА БЕЛАЯ","кг",8), ("71136","Посыпка шоколадная ВЕРМИШЕЛЬ БЕЛАЯ","кг",8), ("71145","Посыпка шоколадная ВЕРМИШЕЛЬ ТЕМНАЯ","кг",8), ("71146","Посыпка шоколадная ВЕРМИШЕЛЬ МОЛОЧНАЯ","кг",8), ("71147","Посыпка шоколадная ВЕРМИШЕЛЬ БЕЛАЯ","кг",8), ("71141","Какао-масло","кг",9), ("71190","Какао-масло","кг",9), ("71172","Какао-масло","кг",9), ("71140","Какао-паста","кг",9), ("71138","Какао-порошок 22/24","кг",9), ("71151","Какао-порошок 22/24 10х1 кг.","кг",9), ("71149","Какао-порошок 22/24","кг",9), ("71106","Глазурь сахарная белая","кг",7), ("71115","Глазурь сахарная","кг",7), ("71107","Глазурь сахарная ФОНДАНТ СОФТ","кг",7), ("71210","Глазурь сахарн.ореховая РЕАЛ","кг",7), ("71031","Глазурь кондит.'Шок.термостаб.капли 1200'","кг",7), ("71037","Глазурь кондит.'Шок.термостаб.капли 1600'","кг",7), ("71125","Глазурь белая КИРОН","кг",7), ("71126","Глазурь какаосодержащая КИРОН","кг",7), ("71033","Глазурь кондит.'Шок.термостаб.капли 1200'","кг",7), ("71039","Глазурь кондит.'Шок.термостаб.капли 1600'","кг",7), ("71046","Глазурь белая 'Шок.термостаб.капли 1600'","кг",7), ("71005","Глазурь кондит.'Шок.диски ИТАЛИКА'","кг",7), ("71016","Глазурь какаосодерж.конд.'Шок.диски ТОПКОВЕР'","кг",7), ("71018","Глазурь белая 'Шок.диски ТОПКОВЕР'","кг",7), ("71042","Глазурь белая 'Шок.диски ИТАЛИКА'","кг",7), ("71044","Глазурь белая 'Шок.диски ИТАЛИКА'","кг",7), ("71087","Глазурь 'М' Молочный шоколад","кг",7), ("71079","Глазурь 'М' Шоколад","кг",7), ("71081","Глазурь 'М' Белый шоколад","кг",7), ("71083","Глазурь 'М' Фисташка","кг",7), ("71084","Глазурь 'М' Лесной орех","кг",7), ("71088","Глазурь 'М' Клубника","кг",7), ("71086","Глазурь 'М' Лимон","кг",7);

INSERT INTO `components` (
    `article`, 
    `name`, 
    `measure_unit`, 
    `component_type_id`
)
VALUES ("88240","Ром-баба 5 см.","кг",10), ("88190","Украшение вафельное ДИСК 60 мм","шт",10), ("88191","Украшение вафельное СИГАРА 9 см","шт",10), ("88108","Вафельный рожок 185 ЧЕРНЫЙ естеств.край","шт",10), ("88109","Вафельный рожок 110 ЧЕРНЫЙ ровный край","шт",10), ("88110","Вафельный стаканчик ФАКЕЛ","шт",10), ("88192","Вафельный рожок ФАКЕЛ 126 мм","шт",10), ("88193","Вафельный рожок ФАКЕЛ 141 мм","шт",10), ("88100","Вафельный рожок 185 естеств.край","шт",10), ("88101","Вафельный рожок 110 ровный край","шт",10), ("88102","Вафельный рожок 150 естеств.край","шт",10), ("88103","Вафельный рожок 60 Миниконус ровный край","шт",10), ("88104","Вафельная корзиночка","шт",10), ("88105","Вафельный стаканчик 75","шт",10), ("88172","Вафельный рожок 150 глазированный ест. край","шт",10), ("88171","Вафельный рожок 110 глазированный ровн. край","шт",10), ("88111","Валован пресный 44 мм.","шт",10), ("88112","Валован пресный 38 мм.","шт",10), ("88211","Тарталетка сладкая 44 мм.","шт",10), ("88212","Тарталетка сладкая 74 мм.","шт",10), ("88214","Тарталетка соленая 44 мм.","шт",10), ("88215","Лодочка сладкая 67 мм.","шт",10), ("88216","Лодочка соленая 67 мм.","шт",10), ("88217","Тарталетка соленая 74 мм.","шт",10), ("88218","Тарталетка сладкая Квадрат 45х45 мм.блистер","шт",10), ("88219","Тарталетка соленая Квадрат 45х45 мм.блистер","шт",10), ("88220","Тарталетка соленая Ассорти Квадрат 45х45 мм","шт",10), ("88228","Лодочка шоколадная 67 мм.","шт",10), ("88234","Трубочка сицилийская 6 см.","шт",10), ("88235","Трубочка сицилийская 9 см.","шт",10), ("88229","Тарталетка шоколадная 44 мм.","шт",10), ("88243","Тарталетка Миньон 53 мм.","шт",10), ("88244","Тарталетка Ложка сладкая 80 мм.","шт",10), ("88245","Тарталетка Ложка соленая 80 мм.","шт",10), ("88230","Печенье миндальное МАКАРУНС ассорти 6 цветов","шт",10), ("88252","Бисквитные палочки САВОЙЯРДИ","шт",10), ("88300","Профитроль 37 мм.","шт",10), ("88302","Профитроль 45 мм.","шт",10), ("82001","Шарики желейн.КЛУБНИКА 3/4 мм","кг",11), ("82011","Шарики желейн.ЧЕРЕШНЯ 9/10 мм","кг",11), ("82021","Шарики желейн.ЧЕРЕШНЯ 13/14 мм","кг",11), ("82027","Шарики желейн.ЧЕРЕШНЯ 13/14 мм","кг",11), ("82031","Шарики желейн.ЧЕРЕШНЯ 17/18 мм","кг",11), ("82081","Шарики желейн.ЯБЛОКО 3/4 мм","кг",11), ("82091","Шарики желейн.ЯБЛОКО 9/10 мм","кг",11), ("82101","Шарики желейн.ЯБЛОКО 13/14 мм","кг",11), ("82121","Шарики желейн.ЛЕСНАЯ ЯГОДА 3/4 мм","кг",11), ("82131","Шарики желейн.ЛЕСНАЯ ЯГОДА 9/10 мм","кг",11), ("82141","Шарики желейн.ЛЕСНАЯ ЯГОДА 13/14 мм","кг",11), ("82161","Шарики желейн.ЛИМОН 3/4 мм","кг",11), ("82171","Шарики желейн.ЛИМОН 9/10 мм","кг",11), ("82181","Шарики желейн.ЛИМОН 13/14 мм","кг",11), ("51401","Шок.декор-аэрозоль ДОЛЬЧЕ ВЕЛЮТО красный","шт",12), ("51402","Шок.декор-аэрозоль ДОЛЬЧЕ ВЕЛЮТО оранжевый","шт",12), ("51403","Шок.декор-аэрозоль ДОЛЬЧЕ ВЕЛЮТО жёлтый","шт",12), ("51404","Шок.декор-аэрозоль ДОЛЬЧЕ ВЕЛЮТО розовый","шт",12), ("51405","Шок.декор-аэрозоль ДОЛЬЧЕ ВЕЛЮТО коричневый","шт",12), ("51406","Шок.декор-аэрозоль ДОЛЬЧЕ ВЕЛЮТО светл.корич.","шт",12), ("51407","Шок.декор-аэрозоль ДОЛЬЧЕ ВЕЛЮТО белый","шт",12), ("51408","Шок.декор-аэрозоль ДОЛЬЧЕ ВЕЛЮТО зеленый","шт",12), ("51409","Шок.декор-аэрозоль ДОЛЬЧЕ ВЕЛЮТО голубой","шт",12), ("51412","Шок.декор-аэрозоль ДОЛЬЧЕ ВЕЛЮТО вишнёвый","шт",12), ("51413","Шок.декор-аэрозоль ДОЛЬЧЕ ВЕЛЮТО черный","шт",12), ("71171","Украшение шоколадное ШАРИКИ КРАНЧ ТЕМНЫЕ","кг",12), ("71169","Украшение шоколадное ШАРИКИ КРАНЧ БЕЛЫЕ","кг",12), ("71170","Украшение шоколадное ШАРИКИ КРАНЧ МОЛОЧНЫЕ","кг",12), ("80290","Капли шоколадные ПЕПИТА темные","кг",12), ("80291","Капли шоколадные ПЕПИТА белые","кг",12), ("80001","Посыпка из конд.глазури ШОК.КРОШКА ТЕМНАЯ","кг.",13), ("80003","Посыпка из конд.глазури ШОК.КРОШКА МИНИ ТЕМН.","кг.",13), ("80010","Посыпка из конд.глазури ШОК.КРОШКА ТЕМНАЯ","кг.",13), ("80040","Посыпка из белой глазури ШОК.КРОШКА БЕЛАЯ","кг.",13), ("80043","Посыпка из белой глазури ШОК.КРОШКА БЕЛАЯ","кг.",13), ("80083","Посыпка из белой глазури ШОК.КРОШКА КРАСНАЯ","кг.",13), ("80084","Посыпка из белой глазури ШОК.КРОШКА ЖЕЛТАЯ","кг.",13), ("80085","Посыпка из белой глазури ШОК.КРОШКА РОЗОВАЯ","кг.",13), ("80102","Посыпка из белой глазури ШОК.ВЕРМИШЕЛЬ БЕЛАЯ","кг.",13), ("80110","Посыпка из конд.глазури ШОК.ВЕРМИШЕЛЬ ТЕМНАЯ","кг.",13), ("80113","Посыпка из конд.глазури ШОК.ВЕРМИШЕЛЬ ТЕМНАЯ","кг.",13), ("80183","Посыпка из белой глазури ШОК.ВЕРМИШЕЛЬ КРАСН.","кг.",13), ("80184","Посыпка из белой глазури ШОК.ВЕРМИШЕЛЬ ЖЕЛТАЯ","кг.",13), ("80185","Посыпка из белой глазури ШОК.ВЕРМИШЕЛЬ РОЗОВ.","кг.",13), ("80170","Посыпка сахар.ВЕРМИШЕЛЬ разноцвет.","кг.",13), ("80171","Посыпка сахар.ВЕРМИШЕЛЬ разноцвет.","кг.",13), ("71211","Глазурь сахарно-ореховая РЕАЛ","кг.",13), ("80173","Посыпка сахар.ВЕРМИШЕЛЬ разноцвет. с НК","кг.",13), ("80200","Посыпка сахар.разноцветная МОНПАРЕЛЬ","кг.",13), ("80201","Посыпка сахар.разноцвeт.МОНПАРЕЛЬ","кг.",13), ("80243","Посыпка сахар.АМАРЕТТО","кг.",13), ("80281","Посыпка КОНФЕТТИ","кг.",13), ("80249","Посыпка сахар.АМАРЕЛЛА","кг.",13), ("80250","Посыпка сахар.АМАРЕЛЛА","кг.",13), ("80259","Посыпка сахар.ДУТЫЙ РИС","кг.",13), ("80261","Посыпка ДУТЫЙ РИС САХАРИСТЫЙ","кг.",13), ("80266","Посыпка сахарная декор-пудра НЭВИССИМА","кг.",13), ("80270","Сахарная декор-пудра БУКАНЕВЕ","кг.",13), ("80272","Сахарная декор-пудра БУКАНЕВЕ","кг.",13), ("80273","Сахарная декор-пудра ДОЛОМИТИ","кг.",13), ("80280","Посыпка какао-пудра БУКАО","кг.",13), ("80282","Посыпка какао-пудра БУКАО","кг.",13), ("80284","Посыпка какао-пудра СПОЛВЕРШОК","кг.",13), ("80285","Посыпка какао-пудра СПОЛВЕРШОК","кг.",13), ("80288","Посыпка сахар.МЕРЕНГА крошка","кг.",13), ("80297","Посыпка сахар.ГРАНЕЛЛА мелкая","кг.",13), ("88121","Вафельная сахарная крошка 1-3 мм","кг",13), ("88106","Вафельная сахарная крошка 3-5 мм","кг",13);


-- Доставка
INSERT INTO `supplier_components` (
    `supplier_id`, 
    `component_id`, 
    `price`
) 
VALUES (1,1,12348.72), (1,2,5688.00), (1,3,1124.64), (1,4,3888.00), (1,5,1821.60), (1,6,3384.00), (1,7,1389.60), (1,8,288), (1,9,2196.00), (1,10,1836.00), (1,11,396), (1,12,3960.00), (1,13,1240.00), (1,14,3100.00), (1,15,3600.00), (1,16,3455.00), (1,17,3340.00), (1,18,985), (1,19,3024.00), (1,20,2160.00), (1,21,1008.00), (1,22,936), (1,23,972), (1,24,720), (1,25,588), (1,26,1260.00), (1,27,612), (1,28,1368.00), (1,29,1485.00), (1,30,2275.20), (1,31,1764.00), (1,32,2016.00), (1,33,2736.00), (1,34,2736.00), (1,35,4104.00), (1,36,3456.00), (1,37,1980.00), (1,38,3852.00), (1,39,1598.40), (1,40,2160.00), (1,41,3620.00), (1,42,3240.00), (1,43,5040.00), (1,44,1980.00), (1,45,3852.00), (1,46,2412.00), (1,47,1700.00), (1,48,3150.00), (1,49,3240.00), (1,50,2700.00), (1,51,3204.00), (1,52,3456.00), (1,53,396), (1,54,6120.00), (1,55,540), (1,56,540), (1,57,2160.00), (1,58,2268.00), (1,59,3204.00), (1,60,2592.00), (1,61,2448.00), (1,62,3492.00), (1,63,2844.00), (1,64,4968.00), (1,65,1270.00), (1,66,840), (1,67,680), (1,68,2592.00), (1,69,1485.00), (1,70,1440.00), (1,71,1800.00), (1,72,1872.00), (1,73,1728.00), (1,74,1368.00), (1,75,1893.60), (1,76,1764.00), (1,77,1317.60), (1,78,1548.00), (1,79,1620.00), (1,80,1800.00), (1,81,1908.00), (1,82,1728.00), (1,83,1692.00), (1,84,1692.00), (1,85,1980.00), (1,86,576), (1,87,612), (1,88,468), (1,89,2736.00), (1,90,2232.00), (1,91,720), (1,92,936), (1,93,828), (1,94,756), (1,95,990), (1,96,684), (1,97,864), (1,98,756), (1,99,756), (1,100,684), (1,101,1296.00), (1,102,1368.00), (1,103,1152.00), (1,104,1080.00), (1,105,1008.00), (1,106,936), (1,107,1044.00), (1,108,1008.00), (1,109,900), (1,110,864), (1,111,1152.00), (1,112,1512.00), (1,113,2556.00), (1,114,7200.00), (1,115,1440.00), (1,116,5760.00), (1,117,5904.00), (1,118,6192.00), (1,119,3384.00), (1,120,2736.00), (1,121,5076.00), (1,122,1332.00), (1,123,3096.00), (1,124,2520.00), (1,125,2520.00), (1,126,1512.00), (1,127,2664.00), (1,128,3528.00), (1,129,4320.00), (1,130,2880.00), (1,131,3024.00), (1,132,1512.00), (1,133,2700.00), (1,134,3168.00), (1,135,2808.00), (1,136,2484.00), (1,137,3096.00), (1,138,864), (1,139,936), (1,146,2426.40), (1,147,2167.20), (1,148,324), (1,149,2304.00), (1,150,2304.00), (1,151,2376.00), (1,152,3888.00), (1,153,5040.00), (1,154,3960.00), (1,155,4163.76), (1,156,7200.00), (1,157,4212.00), (1,158,4060.80), (1,159,4464.00), (1,160,5040.00), (1,161,9000.00), (1,162,4896.00), (1,163,4716.00), (1,164,4032.00), (1,165,3996.00), (1,166,3744.00), (1,167,3758.40), (1,168,4032.00), (1,169,5241.60), (1,170,3564.00), (1,171,3715.20), (1,172,3715.20), (1,173,4140.00), (1,174,3715.20), (1,175,4680.00), (1,176,3628.80), (1,177,3852.00), (1,178,6264.00), (1,179,5904.00), (1,180,9360.00), (1,181,6840.00), (1,182,9216.00), (1,183,9000.00), (1,184,5760.00), (1,185,9360.00), (1,186,6480.00), (1,187,19080.00), (1,188,31104.00), (1,189,23832.00), (1,190,6552.00), (1,191,7704.00), (1,192,3096.00), (1,193,4334.40), (1,194,3528.00), (1,195,3600.00), (1,140,2880.00), (1,141,2880.00), (1,142,7632.00), (1,143,3888.00), (1,144,3672.00), (1,145,3801.60), (1,196,2628.00), (1,197,3456.00), (1,198,720), (1,199,1944.00), (1,200,2088.00), (1,201,2160.00), (1,202,2376.00), (1,203,2304.00), (1,204,2232.00), (1,205,4896.00), (1,206,1008.00), (1,207,2016.00), (1,208,1044.00), (1,209,1044.00), (1,210,1224.00), (1,211,1260.00), (1,212,1260.00), (1,213,1260.00), (1,214,2880.00), (1,215,3528.00), (1,216,1224.00), (1,217,1224.00), (1,218,1375.20), (1,219,1800.00), (1,220,2376.00), (1,221,2376.00), (1,222,2376.00), (1,223,2304.00), (1,224,2376.00), (1,225,828), (1,226,1584.00), (1,227,4536.00), (1,228,2016.00), (1,229,2052.00), (1,230,900), (1,231,1224.00), (1,232,1224.00), (1,233,2016.00), (1,234,1080.00), (1,235,1080.00), (1,236,1224.00), (1,237,432), (1,238,432), (1,239,450), (1,240,432), (1,241,1949.04), (1,242,1656.00), (1,243,1728.00), (1,244,2016.00), (1,245,1620.00), (1,246,396), (1,247,1512.00), (1,248,1188.00), (1,249,432), (1,250,4248.00), (1,251,360), (1,252,1332.00), (1,253,2484.00), (1,254,3204.00), (1,255,720), (1,256,792), (1,257,828), (1,258,1512.00), (1,259,1584.00), (1,260,1584.00), (1,261,1512.00), (1,262,1476.00), (1,288,4802.40), (1,289,4464.00), (1,290,4939.20), (1,291,1375.20), (1,292,1296.00), (1,293,1512.00), (1,294,5328.00), (1,295,4248.00), (1,296,4896.00), (1,297,5256.00), (1,298,432), (1,299,2016.00), (1,300,432), (1,301,540), (1,302,2520.00), (1,303,432), (1,304,432), (1,305,504), (1,306,4176.00), (1,307,828), (1,308,4032.00), (1,309,2592.00), (1,310,576), (1,263,5760.00), (1,264,9792.00), (1,265,3636.00), (1,266,2448.00), (1,267,2592.00), (1,268,432), (1,269,432), (1,270,432), (1,271,3600.00), (1,272,3312.00), (1,273,7200.00), (1,274,7344.00), (1,275,3528.00), (1,276,6768.00), (1,277,3960.00), (1,278,3672.00), (1,279,504), (1,280,7920.00), (1,281,4086.72), (1,282,6984.00), (1,283,2736.00), (1,284,6480.00), (1,285,3456.00), (1,286,3600.00), (1,287,3168.00)
    ;


INSERT INTO `supplier_components` (
    `supplier_id`, 
    `component_id`, 
    `price`
) 
VALUES (2,311,2606.40), (2,312,1728.00), (2,313,1339.20), (2,314,1200.00), (2,315,570), (2,316,800), (2,317,1368.00), (2,318,1188.00), (2,319,875), (2,320,400), (2,321,750), (2,322,1725.00), (2,323,975), (2,324,615), (2,325,920), (2,326,1125.00), (2,327,1656.00), (2,328,1980.00), (2,329,1224.00), (2,330,1296.00), (2,331,1224.00), (2,332,1296.00), (2,333,1296.00), (2,334,1296.00), (2,335,1656.00), (2,336,1656.00), (2,337,2376.00), (2,338,1296.00), (2,339,1296.00), (2,340,936), (2,341,1224.00), (2,342,1800.00), (2,343,2880.00), (2,344,2880.00), (2,345,2160.00), (2,346,1411.20), (2,347,1584.00), (2,348,1512.00), (2,349,900), (2,350,918), (2,351,900), (2,352,5112.00), (2,353,842), (2,354,900), (2,355,900), (2,356,900), (2,357,900), (2,358,900), (2,359,882), (2,360,900), (2,361,900), (2,362,900), (2,363,1584.00), (2,364,1584.00), (2,365,1584.00), (2,366,1584.00), (2,367,1584.00), (2,368,1584.00), (2,369,1584.00), (2,370,1584.00), (2,371,1584.00), (2,372,1584.00), (2,373,1584.00), (2,374,1692.00), (2,375,1764.00), (2,376,1764.00), (2,377,1980.00), (2,378,2448.00), (2,379,388), (2,380,2736.00), (2,381,6840.00), (2,382,388), (2,383,7560.00), (2,384,5328.00), (2,385,5328.00), (2,386,5328.00), (2,387,3744.00), (2,388,468), (2,389,8496.00), (2,390,4320.00), (2,391,4320.00), (2,392,4320.00), (2,393,360), (2,394,6120.00), (2,395,10368.00), (2,396,2761.20), (2,397,1610.00), (2,398,360), (2,399,5472.00), (2,400,648), (2,401,648), (2,402,4752.00), (2,403,504), (2,404,2988.00), (2,405,1800.00), (2,406,396), (2,407,1908.00), (2,408,3600.00), (2,409,432), (2,410,3024.00), (2,411,547), (2,412,2664.00), (2,413,8136.00), (2,414,1800.00), (2,415,850), (2,416,1700.00);


INSERT INTO `purchased_components` (`id`, `component_id`, `supplier_id`, `quantity`, `purchase_price`, `shelf_life`) VALUES
(1, 1, 1, 11, 20.00, NULL),
(2, 2, 1, 3, 14.00, NULL),
(3, 3, 1, 5, 15.00, NULL),
(4, 1, 1, 3, 11.02, NULL);

INSERT INTO `products` (`id`, `name`, `sizes`, `is_semiproduct`) VALUES
(1, 'Торт \"Шоколадный\"', '50x30', 0);


INSERT INTO `tool_failure_reasons` (`id`, `name`) VALUES
    (1, 'Custom reason'),
    (2, 'test_reason');


INSERT INTO `tool_failures` (
    `id`, `tool_id`, `master_id`, 
    `failure_reason_id`, `failure_at`, `fixed_at`) 
VALUES
    (1, 7, 11, 1, '2024-11-13 10:39:41', '2024-11-13 00:00:00'),
    (2, 7, 11, 2, '2024-11-13 10:42:32', NULL);