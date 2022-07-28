package u2ug

var adjectives = []string{
	"Busy",
	"Careless",
	"Clumsy",
	"Nimble",
	"Brave",
	"Mighty",
	"Meek",
	"Clever",
	"Dull",
	"Afraid",
	"Scared",
	"Cowardly",
	"Bashful",
	"Lazy",
	"Proud",
	"Fair",
	"Greedy",
	"Wise",
	"Foolish",
	"Tricky",
	"Truthful",
	"Loyal",
	"Happy",
	"Cheerful",
	"Joyful",
	"Carefree",
	"Friendly",
	"Moody",
	"Crabby",
	"Cranky",
	"Awful",
	"Gloomy",
	"Angry",
	"Worried",
	"Excited",
	"Calm",
	"Bored",
	"Hardworking",
	"Silly",
	"Wild",
	"Crazy",
	"Fussy",
	"Still",
	"Odd",
	"Starving",
	"Stuffed",
	"Alert",
	"Sleepy",
	"Surprised",
	"Tense",
	"Rude",
	"Selfish",
	"Strict",
	"Tough",
	"Polite",
	"Amusing",
	"Kind",
	"Gentle",
	"Quiet",
	"Caring",
	"Hopeful",
	"Rich",
	"Thrifty",
	"Stingy",
	"Spoiled",
	"Generous",
	"Quick",
	"Speedy",
	"Swift",
	"Hasty",
	"Rapid",
	"Good",
	"Fantastic",
	"Splendid",
	"Wonderful",
	"Hard",
	"Difficult",
	"Challenging",
	"Easy",
	"Simple",
	"Chilly",
	"Freezing",
	"Icy",
	"Steaming",
	"Sizzling",
	"Muggy",
	"Cozy",
	"Huge",
	"Great",
	"Vast",
	"Sturdy",
	"Grand",
	"Heavy",
	"Plump",
	"Deep",
	"Puny",
	"Small",
	"Tiny",
	"Petite",
	"Long",
	"Endless",
	"Beautiful",
	"Adorable",
	"Shining",
	"Sparkling",
	"Glowing",
	"Fluttering",
	"Soaring",
	"Crawling",
	"Creeping",
	"Sloppy",
	"Messy",
	"Slimy",
	"Grimy",
	"Crispy",
	"Spiky",
	"Rusty",
	"Smelly",
	"Foul",
	"Stinky",
	"Curly",
	"Fuzzy",
	"Plush",
	"Lumpy",
	"Wrinkly",
	"Smooth",
	"Glassy",
	"Snug",
	"Stiff",
	"Ugly",
	"Hideous",
	"Horrid",
	"Dreadful",
	"Nasty",
	"Cruel",
	"Creepy",
	"Loud",
	"Shrill",
	"Muffled",
	"Creaky",
}

var games = []string{
	"Trivial Pursuit",
	"Zombi",
	"Asphalt",
	"Defender of the Crown",
	"Le Maître des Âmes",
	"Le Nécromancien",
	"Mange Cailloux",
	"ST Krak",
	"Night Hunter",
	"B.A.T.",
	"Defender of the Crown",
	"Final Command",
	"Fred",
	"Ilyad",
	"Intruder",
	"Iron Lord",
	"Jimmy Connors Pro Tennis Tour",
	"Othello Killer",
	"Puffy's Saga",
	"Skateball",
	"TwinWorld: Land of Vision",
	"B.A.T.",
	"Brain Blasters",
	"Fred",
	"Hexsider",
	"Iron Lord",
	"Jupiter's Masterdrive",
	"Night Hunter",
	"Pick 'n Pile",
	"Pool of Radiance",
	"Puffy's Saga",
	"Ranx: The Video Game",
	"Tom and the Ghost",
	"TwinWorld: Land of Vision",
	"Unreal",
	"Zombi",
	"B.A.T.",
	"Bomberman",
	"Celtic Legends",
	"Great Courts 2",
	"Jimmy Connors Pro Tennis Tour",
	"Maupiti Island",
	"Pick 'n Pile",
	"Unreal",
	"Bomberman",
	"First Samurai",
	"Jimmy Connors Pro Tennis Tour",
	"Mega-Lo-Mania",
	"Starush",
	"Star Wars",
	"Star Wars: The Empire Strikes Back",
	"The Koshan Conspiracy",
	"The Perfect General",
	"Vroom",
	"Jimmy Connors Tennis",
	"F1 Pole Position",
	"Indiana Jones and the Last Crusade: The Action Game",
	"Championship Manager",
	"F1 Pole Position",
	"The Koshan Conspiracy",
	"Indiana Jones and the Last Crusade: The Action Game",
	"Street Racer",
	"F1 Pole Position 2",
	"Hyper V-Ball",
	"Soul Blazer",
	"Rayman",
	"Action Euro Soccer 96",
	"Action Soccer",
	"Kiyeko and the Lost Night",
	"Rayman",
	"Street Racer",
	"Payuta and the Ice God",
	"Adventures in Music with the Recorder",
	"Amazing Learning Games with Rayman",
	"Maths and English with Rayman: Volume 2",
	"Maths and English with Rayman: Volume 3",
	"Secrets of the Luxor",
	"The Adventures of Valdo & Marie",
	"F1 Pole Position 64",
	"Sub Culture",
	"Earth 2140",
	"F1 Racing Simulation",
	"POD",
	"POD: Back to Hell",
	"Pro Pinball: Timeshock!",
	"Rayman Designer",
	"Sean Dundee's World Club Football",
	"S.C.A.R.S.",
	"All Star Tennis '99",
	"S.C.A.R.S.",
	"Buck Bumble",
	"Cob to the Rescue!",
	"Football World Manager",
	"Gex: Enter the Gecko",
	"Hell-Copter",
	"Hexcite: The Shapes of Victory",
	"Might and Magic VI: The Mandate of Heaven",
	"Redline Racer",
	"S.C.A.R.S.",
	"Shadow Gunner: The Robot Wars",
	"Speed Busters: American Highways",
	"The Fifth Element",
	"World Football 98",
	"All Star Tennis '99",
	"Star Wars: Episode I - Racer",
	"Monaco Grand Prix Racing Simulation 2",
	"Tonic Trouble",
	"Monaco Grand Prix Racing Simulation 2",
	"Wall Street Tycoon",
	"Rayman 2: The Great Escape",
	"Rocket: Robot on Wheels",
	"Mobil 1 Rally Championship",
	"Tonic Trouble",
	"The Longest Journey",
	"Evolution: The World of Sacred Device",
	"Alex Builds His Farm",
	"Laura's Happy Adventures",
	"Monaco Grand Prix Racing Simulation 2",
	"Requiem: Avenging Angel",
	"Skullcaps",
	"Seven Kingdoms II: The Fryhtan Wars",
	"Shadow Company: Left for Dead",
	"Speed Devils",
	"Suzuki Alstare Extreme Racing",
	"Papyrus",
	"Theocracy",
	"Inspector Gadget: Operation Madkactus",
	"Rayman 2: The Great Escape",
	"Business Tycoon",
	"Rayman",
	"Disney's Dinosaur",
	"Evolution 2: Far off Promise",
	"All Star Tennis 2000",
	"Toonsylvania",
	"Surf Riders",
	"Deep Fighter",
	"Infestation",
	"O'Leary Manager 2000",
	"Rayman 2: The Great Escape",
	"F1 Racing Championship",
	"Deep Fighter",
	"Carl Lewis Athletics 2000",
	"In Cold Blood",
	"Virtual Skipper",
	"Batman Beyond: Return of the Joker",
	"Disney's Donald Duck: Goin' Quackers",
	"Walt Disney's The Jungle Book: Rhythm n' Groove",
	"Gold and Glory: The Road to El Dorado",
	"Disney's Aladdin",
	"F1 Racing Championship",
	"Chessmaster 8000",
	"Walt Disney's The Jungle Book: Mowgli's Wild Adventure",
	"Batman Beyond: Return of the Joker",
	"Disney's Donald Duck: Goin' Quackers",
	"Tom Clancy's Rainbow Six: Rogue Spear",
	"Tom Clancy's Rainbow Six: Rogue Spear Mission Pack - Urban Operations",
	"Disney's Dinosaur",
	"Batman Beyond: Return of the Joker",
	"Grandia II",
	"POD SpeedZone",
	"F1 Racing Championship",
	"Walt Disney's The Jungle Book: Rhythm n' Groove",
	"Disney's Donald Duck: Goin' Quackers",
	"Rayman 2 Forever",
	"Tokyo Xtreme Racer 2",
	"Gold and Glory: The Road to El Dorado",
	"Disney's The Emperor's New Groove",
	"Eternal Ring",
	"Little Nicky",
	"Sno-Cross Championship Racing",
	"Arcatera: The Dark Brotherhood",
	"All Star Tennis 2000",
	"Amazing Learning Games with Rayman",
	"Anne McCaffrey's Freedom: First Resistance",
	"Animorphs",
	"David O'Leary's Total Soccer 2000",
	"Disney's Dinosaur",
	"Disney's Donald Duck: Goin' Quackers",
	"F1 Racing Championship",
	"Grandia",
	"Papyrus: Le Secret de la Cité Perdue",
	"Princesse Sissi et Tempête",
	"Pro Rally 2001",
	"Rayman Revolution",
	"Sea Dogs",
	"Spirou: The Robot Invasion",
	"Suzuki Alstare Extreme Racing",
	"The Mask of Zorro",
	"Tom and Jerry in Fists of Furry",
	"Tonic Trouble",
	"Stupid Invaders",
	"Disney's Dinosaur",
	"F1 Racing Championship",
	"Armored Core 2",
	"The Settlers IV",
	"Flipper & Lopaka",
	"Batman: Chaos in Gotham",
	"Myst III: Exile",
	"Conflict Zone",
	"Eurofighter Typhoon",
	"Hype: The Time Quest",
	"Rayman Advance",
	"Stupid Invaders",
	"Europa Universalis",
	"Roswell Conspiracies: Aliens, Myths & Legends",
	"Moderngroove: Ministry of Sound Edition",
	"V.I.P.",
	"Dragon Riders: Chronicles of Pern",
	"Conquest: Frontier Wars",
	"Planet of the Apes",
	"Final Fight",
	"Pool of Radiance: Ruins of Myth Drannor",
	"Evil Twin: Cyprien's Chronicles",
	"Batman: Vengeance",
	"Kohan: Immortal Sovereigns",
	"Tom Clancy's Rainbow Six: Rogue Spear - Black Thorn",
	"Gadget Tycoon",
	"Sunny Garcia Surfing",
	"Rally Championship 2002",
	"Silent Hunter II",
	"Battle Realms",
	"The Final Cut",
	"Super Street Fighter II: Turbo Revival",
	"Conflict Zone",
	"Tom Clancy's Ghost Recon",
	"Disney's Tarzan Untamed",
	"Batman: Vengeance",
	"IL-2 Sturmovik",
	"Walt Disney's Snow White and the Seven Dwarfs",
	"Super Bust-A-Move",
	"Disney's Tarzan Untamed",
	"Mega Man Battle Network",
	"Rayman Arena",
	"The Legend of Alon D'ar",
	"Evil Twin: Cyprien's Chronicles",
	"Batman: Vengeance",
	"Breath of Fire",
	"Worms World Party",
	"Trevor Chan's Capitalism II",
	"Conflict Zone",
	"Jade Cocoon 2",
	"Donald Duck Advance",
	"Rayman Arena",
	"Alex Ferguson's Player Manager 2001",
	"Batman: Gotham City Racer",
	"Escape from Monkey Island",
	"Gunfighter: The Legend of Jesse James",
	"Inspector Gadget: Gadget's Crazy Maze",
	"Pearl Harbor: Strike at Dawn",
	"Roswell Conspiracies: Aliens, Myths & Legends",
	"Scrabble",
	"Taxi 2",
	"Tokyo Xtreme Racer: Zero",
	"Grandia II",
	"Salt Lake 2002",
	"Echelon",
	"Destroyer Command",
	"Mike Tyson Boxing",
	"Evolution Worlds",
	"Rayman Rush",
	"Bratz",
	"Warlords: Battlecry II",
	"Grandia II",
	"Tom Clancy's Rainbow Six: Rogue Spear",
	"Bratz",
	"E.T. The Extra-Terrestrial: Interplanetary Mission",
	"Dark Planet: Battle for Natrolis",
	"Jim Henson's Bear in the Big Blue House",
	"X-Bladez: Inline Skater",
	"Disney's Donald Duck: Goin' Quackers",
	"Tom Clancy's Ghost Recon: Desert Siege",
	"E.T.: Phone Home Adventure",
	"Pro Rally 2002",
	"Europa Universalis II",
	"Hooters Road Trip",
	"Sabrina, the Teenage Witch: Potion Commotion",
	"Ultimate Fighting Championship: Tapout",
	"Evil Twin: Cyprien's Chronicles",
	"Monster Jam: Maximum Destruction",
	"The Sum of All Fears",
	"Taxi 2",
	"UFC: Throwdown",
	"The Elder Scrolls III: Morrowind",
	"Sven-Göran Eriksson's World Manager",
	"Taxi 2",
	"Breath of Fire II",
	"Muppet Pinball Mayhem",
	"Planet of the Apes",
	"Tom Clancy's Rainbow Six: Lone Wolf",
	"Dokapon: Monster Hunter",
	"V.I.P.",
	"UFC: Throwdown",
	"Tiger Woods PGA Tour Golf",
	"Monster Jam: Maximum Destruction",
	"Worms Blast",
	"Largo Winch: Empire Under Threat",
	"Chessmaster 9000",
	"Colin McRae Rally 2.0",
	"Largo Winch: Empire Under Threat",
	"Myst III: Exile",
	"Super Bust-A-Move 2",
	"The Sum of All Fears",
	"Rayman Arena",
	"Tom Clancy's Ghost Recon: Island Thunder",
	"Worms World Party",
	"Largo Winch: Empire Under Threat",
	"Deathrow",
	"Mega Man Battle Network 2",
	"Wizardry: Tale of the Forsaken Land",
	"Speed Challenge: Jacques Villeneuve's Racing Vision",
	"The Sum of All Fears",
	"Disney's PK: Out of the Shadows",
	"Riding Champion: Legacy of Rosemond Hill",
	"Pro Rally",
	"Rocky",
	"Tomb Raider: The Prophecy",
	"Tom Clancy's Ghost Recon",
	"Rocky",
	"Catz 5",
	"Dogz 5",
	"Dragon's Lair 3D: Return to the Lair",
	"Tom Clancy's Splinter Cell",
	"Monster Jam: Maximum Destruction",
	"Rocky",
	"Speed Challenge: Jacques Villeneuve's Racing Vision",
	"The Elder Scrolls III: Morrowind",
	"The Mummy",
	"Street Fighter Alpha 3",
	"Disney's PK: Out of the Shadows",
	"Globetrotter 2",
	"Monster Jam: Maximum Destruction",
	"Super Bust-A-Move",
	"RS3: Racing Simulation Three",
	"Tom Clancy's Ghost Recon",
	"Lunar Legend",
	"Jim Henson's Bear in the Big Blue House",
	"Bratz",
	"Disney's Treasure Planet",
	"Moto Racer Advance",
	"Batman: Vengeance",
	"E.T. The Extra-Terrestrial: Away From Home",
	"Jet Ion GP",
	"Largo Winch .// Commando SAR",
	"Moorhen 3 ...Chicken Chase",
	"Sven-Göran Eriksson's World Challenge",
	"Trains & Trucks Tycoon",
	"V.I.P.",
	"The Sum of All Fears",
	"Taxi 3",
	"Tom Clancy's Ghost Recon",
	"Sword of the Samurai",
	"Tom Clancy's Splinter Cell",
	"Wild Arms 3",
	"IL-2 Sturmovik: Forgotten Battles",
	"Murakumo: Renegade Mech Pursuit",
	"Rayman 3",
	"Rayman 3: Hoodlum Havoc",
	"Walt Disney's The Jungle Book: Rhythm n' Groove",
	"Gunfighter II: Revenge of Jesse James",
	"Rayman 3: Hoodlum Havoc",
	"Tom Clancy's Rainbow Six 3: Raven Shield",
	"Rayman 3: Hoodlum Havoc",
	"CSI: Crime Scene Investigation",
	"Shadowbane",
	"Rayman 3: Hoodlum Havoc",
	"Tom Clancy's Splinter Cell",
	"Chessmaster",
	"Will Rock",
	"Ape Escape 2",
	"Charlie's Angels",
	"Tom Clancy's Ghost Recon: Island Thunder",
	"Crouching Tiger, Hidden Dragon",
	"RS3: Racing Simulation Three",
	"XIII",
	"Crouching Tiger, Hidden Dragon",
	"Batman: Rise of Sin Tzu",
	"In Memoriam",
	"Warlords IV: Heroes of Etheria",
	"Prince of Persia: The Sands of Time",
	"XIII",
	"Tom Clancy's Rainbow Six 3",
	"XIII",
	"Batman: Rise of Sin Tzu",
	"Beyond Good & Evil",
	"Uru: Ages Beyond Myst",
	"Saddle Up: Time to Ride",
	"Mucha Lucha! Mascaritas of the Lost Code",
	"Prince of Persia: The Sands of Time",
	"Beyond Good & Evil",
	"Lock On: Modern Air Combat",
	"Monster 4x4: Masters of Metal",
	"XIII",
	"Biathlon 2004",
	"Prince of Persia: The Sands of Time",
	"Beyond Good & Evil",
	"Crouching Tiger, Hidden Dragon",
	"Monster 4x4: Masters of Metal",
	"Beyond Good & Evil",
	"Downtown Run",
	"Scrabble: 2003 Edition",
	"TOCA: World Touring Cars",
	"Baldur's Gate: Dark Alliance",
	"Scrabble: 2003 Edition",
	"Tom Clancy's Rainbow Six 3: Athena Sword",
	"Tom Clancy's Rainbow Six 3",
	"Tom Clancy's Ghost Recon: Jungle Storm",
	"CSI: Dark Motives",
	"Far Cry",
	"Tom Clancy's Splinter Cell: Pandora Tomorrow",
	"Harvest Moon: A Wonderful Life",
	"Tom Clancy's Ghost Recon 2",
	"Fatal Frame II: Crimson Butterfly",
	"Tom Clancy's Splinter Cell: Pandora Tomorrow",
	"Champions of Norrath",
	"Tom Clancy's Rainbow Six 3",
	"Tom Clancy's Splinter Cell: Pandora Tomorrow",
	"Tom Clancy's Rainbow Six 3: Black Arrow",
	"The Political Machine",
	"Chessmaster 10th Edition",
	"Advance Guardian Heroes",
	"Star Wars Trilogy: Apprentice of the Force",
	"Rocky Legends",
	"Sherlock Holmes: The Case of the Silver Earring",
	"The Dukes of Hazzard: Return of the General Lee",
	"Myst IV: Revelation",
	"Ape Escape: Pumped & Primed",
	"Pacific Fighters",
	"Chessmaster 10th Edition",
	"CSI: Miami",
	"Tom Clancy's Ghost Recon 2",
	"Asphalt: Urban GT",
	"Alexander",
	"Prince of Persia: Warrior Within",
	"Biathlon 2005",
	"Sprung",
	"CSI: Crime Scene Investigation",
	"Tom Clancy's Ghost Recon 2",
	"Tork: Prehistoric Punk",
	"Disney's Winnie the Pooh's Rumbly Tumbly Adventure",
	"Heritage of Kings: The Settlers",
	"Brothers in Arms: Road to Hill 30",
	"Playboy: The Mansion",
	"Brothers in Arms: Road to Hill 30",
	"Cold Fear",
	"Champions: Return to Arms",
	"Rayman: Hoodlum's Revenge",
	"Silent Hunter III",
	"Rayman DS",
	"The Bard's Tale",
	"Lumines: Puzzle Fusion",
	"Cold Fear",
	"Myst IV: Revelation",
	"Tom Clancy's Splinter Cell: Chaos Theory",
	"Star Wars: Episode III - Revenge of the Sith",
	"Bomberman",
	"Tom Clancy's Splinter Cell: Chaos Theory",
	"Bomberman Hardball",
	"Tom Clancy's Ghost Recon 2: Summit Strike",
	"Darkwatch",
	"187: Ride or Die",
	"Tom Clancy's Rainbow Six: Lockdown",
	"Myst V: End of Ages",
	"Marathon Manager",
	"Myst V: End of Ages",
	"Far Cry Instincts",
	"Lunar: Dragon Song",
	"Tom Clancy's Rainbow Six: Lockdown",
	"Heroes of the Pacific",
	"Brothers in Arms: Earned in Blood",
	"FLOW: Urban Dance Uprising",
	"America's Army: Rise of a Soldier",
	"America's Army: True Soldiers",
	"GripShift",
	"King Kong: The Official Game of the Movie",
	"Kong: The 8th Wonder of the World",
	"Trollz: Hair Affair!",
	"Frantix",
	"Prince of Persia: The Two Thrones",
	"Battles of Prince of Persia",
	"Prince of Persia: Revelations",
	"Prince of Persia: The Two Thrones",
	"King Kong: The Official Game of the Movie",
	"Biathlon 2006: Go for Gold",
	"Pippa Funnell: The Stud Farm Inheritance",
	"Rugby Challenge 2006",
	"Drakengard 2",
	"Exit",
	"Tom Clancy's Rainbow Six: Lockdown",
	"Curling 2006",
	"Tom Clancy's Ghost Recon Advanced Warfighter",
	"CSI: 3 Dimensions of Murder",
	"Tom Clancy's Rainbow Six: Critical Hour",
	"Call of Cthulhu: Dark Corners of the Earth",
	"Blazing Angels: Squadrons of WWII",
	"Monster 4x4: World Circuit",
	"Street Riders",
	"Blazing Angels: Squadrons of WWII",
	"Far Cry: Instincts - Evolution",
	"Far Cry: Instincts - Predator",
	"Pippa Funnell: Take the Reins",
	"Tom Clancy's Splinter Cell: Essentials",
	"Dogz",
	"LostMagic",
	"Paradise",
	"Tom Clancy's Ghost Recon Advanced Warfighter",
	"Heroes of Might and Magic V",
	"Call of Juarez",
	"AND 1 Streetball",
	"Astonishia Story",
	"Over G Fighters",
	"Pirates of the Caribbean: The Legend of Jack Sparrow",
	"Devil May Cry 3: Dante's Awakening - Special Edition",
	"Enchanted Arms",
	"Call of Juarez",
	"Faces of War",
	"Open Season",
	"Import Tuner Challenge",
	"The Settlers II: 10th Anniversary",
	"Tom Clancy's Splinter Cell: Double Agent",
	"Dark Messiah of Might and Magic",
	"Tom Clancy's Splinter Cell: Double Agent",
	"Pippa Funnell: Take the Reins",
	"Mind Quiz",
	"Tom Clancy's Splinter Cell: Double Agent",
	"Asphalt: Urban GT 2",
	"GT Pro Series",
	"Monster 4x4: World Circuit",
	"Rayman Raving Rabbids",
	"Red Steel",
	"Tom Clancy's Rainbow Six Vegas",
	"Dogz",
	"Horsez",
	"Open Season",
	"Tom Clancy's Splinter Cell: Double Agent",
	"Safari Photo Africa: Wild Earth",
	"Brothers in Arms: D-Day",
	"Rayman Raving Rabbids",
	"Star Wars: Lethal Alliance",
	"Rayman Raving Rabbids",
	"Blazing Angels: Squadrons of WWII",
	"Far Cry Vengeance",
	"Tom Clancy's Rainbow Six Vegas",
	"Catz",
	"Dogz",
	"Pippa Funnell: The Golden Stirrup Challenge",
	"Scrabble 2007 Edition",
	"Rocky Balboa",
	"Resident Evil 4",
	"Rayman Raving Rabbids",
	"Tom Clancy's Ghost Recon Advanced Warfighter 2",
	"TMNT",
	"Teenage Mutant Ninja Turtles",
	"Blazing Angels: Squadrons of WWII",
	"TMNT",
	"Go! Sudoku",
	"Silent Hunter: Wolves of the Pacific",
	"Tom Clancy's Splinter Cell: Double Agent",
	"Mind Quiz: Your Brain Coach",
	"Beyond Divinity",
	"Enchanted Arms",
	"Prince of Persia: Rival Swords",
	"Prince of Persia: The Two Thrones",
	"Rayman Raving Rabbids",
	"The Elder Scrolls IV: Oblivion",
	"Driver 76",
	"Platinum Sudoku",
	"WarTech: Senko no Ronde",
	"Surf's Up",
	"Prince of Persia Classic",
	"Brothers in Arms DS",
	"Tom Clancy's Rainbow Six Vegas",
	"Driver: Parallel Lines",
	"Tom Clancy's Rainbow Six Vegas",
	"The Settlers",
	"Tom Clancy's Ghost Recon Advanced Warfighter 2",
	"Top Trumps Adventures!",
	"Tom Clancy's Ghost Recon Advanced Warfighter 2",
	"Cosmic Family",
	"Jam Sessions",
	"Blazing Angels 2: Secret Missions of WWII",
	"CSI: 3 Dimensions of Murder",
	"CSI: Hard Evidence",
	"Brain Spa",
	"The Settlers: Rise of an Empire",
	"Real Football 2008 3D",
	"Top Trumps: Dogs & Dinosaurs",
	"Top Trumps: Horror & Predators",
	"Heroes of Might and Magic V: Tribes of the East",
	"Totally Spies! 3: Secret Agents",
	"Chessmaster: Grandmaster Edition",
	"Imagine: Animal Doctor",
	"Imagine: Babyz",
	"Imagine: Fashion Designer",
	"Imagine: Master Chef",
	"Naruto: Rise of a Ninja",
	"Petz Wild Animals: Dolphinz",
	"Chessmaster: Grandmaster Edition",
	"CSI: Dark Motives",
	"Telly Addicts",
	"Who Wants to Be a Millionaire: 1st Edition",
	"Imagine: Interior Designer",
	"Imagine: Wedding Designer",
	"Blazing Angels 2: Secret Missions of WWII",
	"My French Coach",
	"My Spanish Coach",
	"My Word Coach",
	"Beowulf: The Game",
	"Petz: Catz 2",
	"Petz: Dogz 2",
	"Petz: Horsez 2",
	"Rayman Raving Rabbids 2",
	"Assassin's Creed",
	"Rayman Raving Rabbids 2",
	"Petz: Hamsterz 2",
	"Cranium Kabookii",
	"Horsez",
	"Petz: Horsez 2",
	"Miami Nights: Singles in the City",
	"Nitrobike",
	"No More Heroes",
	"Chessmaster Live",
	"Assassin's Creed: Altaïr's Chronicles",
	"Puppy Palace",
	"Dark Messiah: Might and Magic - Elements",
	"Chessmaster: The Art of Learning",
	"Top Trumps: Dogs & Dinosaurs",
	"Lost: Via Domus",
	"Petz: Wild Animals - Tigerz",
	"Anno 1701: Dawn of Discovery",
	"Imagine: Figure Skater",
	"Petz: Bunnyz",
	"Tom Clancy's Rainbow Six Vegas 2",
	"Nitrobike",
	"Tom Clancy's Rainbow Six Vegas 2",
	"The Dog Island",
	"Haze",
	"Emergency Heroes",
	"Protöthea",
	"Rayman Raving Rabbids 2",
	"Petz: Horsez 2",
	"Stratego: Next Edition",
	"Imagine: Rock Star",
	"My Weight Loss Coach",
	"Gourmet Chef",
	"Animal Genius",
	"Quick Yoga Training",
	"Imagine: Teacher",
	"My Chinese Coach",
	"The Settlers: Rise of Cultures",
	"Hell's Kitchen: The Game",
	"The Price Is Right",
	"Armored Core: For Answer",
	"Hell's Kitchen: The Game",
	"Brothers in Arms: Hell's Highway",
	"My SAT Coach",
	"Brothers in Arms: Double Time",
	"My Secret World by Imagine",
	"Brothers in Arms: Hell's Highway",
	"Imagine: Babysitters",
	"Imagine: Fashion Designer - New York",
	"Imagine: Champion Rider",
	"Battle of Giants: Dinosaurs",
	"Cesar Millan's Dog Whisperer",
	"My Japanese Coach",
	"The Dog Island",
	"CSI: NY - The Game",
	"Circus Games",
	"Ener-G: Dance Squad",
	"Far Cry 2",
	"Prince of Persia Classic",
	"Guitar Rock Tour",
	"Petz: Horse Club",
	"Petz Rescue: Ocean Patrol",
	"Petz Rescue: Wildlife Vet",
	"Rayman Raving Rabbids: TV Party",
	"Family Fest Presents Movie Games",
	"Tom Clancy's EndWar",
	"Imagine: Party Babyz",
	"Shaun White Snowboarding",
	"Shaun White Snowboarding: Road Trip",
	"Imagine: Ballet Star",
	"Imagine: Movie Star",
	"Naruto: The Broken Bond",
	"Petz: Crazy Monkeyz",
	"Petz: Monkeyz House",
	"Rayman Raving Rabbids: TV Party",
	"Petz: Catz Clan",
	"Petz Rescue: Endangered Paradise",
	"Who Wants to Be a Millionaire: 2nd Edition",
	"My Stop Smoking Coach: Allen Carr's EasyWay",
	"Petz Sports",
	"Happy Cooking",
	"My Fun Facts Coach",
	"Prince of Persia",
	"Prince of Persia: The Fallen King",
	"Shaun White Snowboarding",
	"Prince of Persia",
	"Baby Life",
	"Telly Addicts",
	"My Fitness Coach",
	"Imagine: Fashion Party",
	"Imagine: Party Planner",
	"Imagine: Cheerleader",
	"Jojo's Fashion Show: Design in a Dash!",
	"Imagine: Journalist",
	"Petz: Horseshoe Ranch",
	"Tenchu: Shadow Assassins",
	"Hell's Kitchen: The Game",
	"Jake Power: Fireman",
	"Jake Power: Policeman",
	"Imagine: Ice Champions",
	"Tom Clancy's EndWar",
	"Six Flags Fun Park",
	"Tom Clancy's H.A.W.X",
	"Grey's Anatomy: The Video Game",
	"World in Conflict: Soviet Assault",
	"Vacation Sports",
	"Tom Clancy's H.A.W.X",
	"Broken Sword: Shadow of the Templars - The Director's Cut",
	"Wheelman",
	"Imagine: My Restaurant",
	"Fashion Studio Paris Collection",
	"Imagine: Family Doctor",
	"My Fitness Coach",
	"Tenchu: Shadow Assassins",
	"Imagine: Music Fest",
	"Imagine: Makeup Artist",
	"Imagine: Boutique Owner",
	"CellFactor: Psychokinetic Wars",
	"My Spanish Coach",
	"Petz Fashion: Dogz and Catz",
	"Anno 1404",
	"Call of Juarez: Bound in Blood",
	"Teenage Mutant Ninja Turtles: Turtles in Time Re-Shelled",
	"Imagine: Soccer Captain",
	"Imagine: Teacher - Class Trip",
	"Imagine: Detective",
	"Academy of Champions: Soccer",
	"Cloudy with a Chance of Meatballs",
	"Heroes Over Europe",
	"Cloudy with a Chance of Meatballs",
	"NewU: Fitness First Personal Trainer",
	"Teenage Mutant Ninja Turtles: Smash-Up",
	"The Price is Right: 2010 Edition",
	"Where's Waldo? The Fantastic Journey",
	"Battle of Giants: Dragons",
	"Family Feud: 2010 Edition",
	"Imagine: Salon Stylist",
	"Metropolis Crimes",
	"Teenage Mutant Ninja Turtles: Turtles in Time Re-Shelled",
	"Imagine: Zookeeper",
	"Imagine: Sweet 16",
	"Petz: Pony Beauty Pageant",
	"Cover Girl",
	"Imagine: Fashion Designer - World Tour",
	"Petz Dolphinz Encounter",
	"Jam Sessions 2",
	"Panzer General: Allied Assault",
	"CSI: Deadly Intent",
	"Imagine: Artist",
	"Monster 4x4 Stunt Racer",
	"Petz: Dogz Family",
	"Rabbids Go Home: A Comedy Adventure",
	"Battle of Giants: Dragons - Bronze Edition",
	"C.O.P.: The Recruit",
	"Imagine: Babyz Fashion",
	"Rabbids Go Home",
	"Shaun White Snowboarding: World Stage",
	"Knockout Party",
	"Fairyland: Melody Magic",
	"Petz Nursery",
	"Style Lab: Jewelry Design",
	"Style Lab: Makeover",
	"Assassin's Creed II",
	"Assassin's Creed II: Discovery",
	"Just Dance",
	"Petz: Dogz Talent Show",
	"Assassin's Creed: Bloodlines",
	"Cook Wars",
	"Petz: Hamsterz Superstarz",
	"Your Shape",
	"Teenage Mutant Ninja Turtles: Arcade Attack",
	"James Cameron's Avatar: The Game",
	"Might & Magic: Clash of Heroes",
	"Rayman",
	"Petz: Horsez 2",
	"Assassin's Creed II: Discovery",
	"Sleepover Party",
	"No More Heroes 2: Desperate Struggle",
	"Battle of Giants: Mutant Insects",
	"Assassin's Creed II",
	"Silent Hunter 5: Battle of the Atlantic",
	"Boot Camp Academy",
	"Battle of Giants: Dinosaurs - Fight for Survival",
	"Imagine: Gymnast",
	"Red Steel 2",
	"Racquet Sports",
	"The Settlers 7: Paths to a Kingdom",
	"Tom Clancy's Splinter Cell: Conviction",
	"Castle & Co",
	"OK! Puzzle Stars",
	"Tom Clancy's Splinter Cell: Conviction",
	"Prince of Persia: The Forgotten Sands",
	"Prince of Persia: The Forgotten Sands",
	"Voodoo Dice",
	"Prince of Persia",
	"Voodoo Dice",
	"Prince of Persia: The Forgotten Sands",
	"Dance on Broadway",
	"Battle of Giants: Mutant Insects - Revenge",
	"Voodoo Dice",
	"Petz Hamsterz Family",
	"Galaxy Racers",
	"Petz: Dogz Family",
	"Scott Pilgrim vs. the World: The Game",
	"Gold's Gym Dance Workout",
	"Scott Pilgrim vs. the World: The Game",
	"Tom Clancy's H.A.W.X 2",
	"R.U.S.E.",
	"Tom Clancy's H.A.W.X 2",
	"Arthur and the Revenge of Maltazard",
	"Shaun White Skateboarding",
	"The Hollywood Squares",
	"Just Dance 2",
	"Puzzler World 2011",
	"The Settlers Online",
	"Shaun White Skateboarding",
	"CSI: Fatal Conspiracy",
	"Bloody Good Time",
	"Fighters Uncaged",
	"Just Dance Kids",
	"Family Feud: Decades",
	"MotionSports",
	"My Chinese Coach",
	"Tom Clancy's H.A.W.X 2",
	"Assassin's Creed: Brotherhood",
	"Tom Clancy's Ghost Recon",
	"Tom Clancy's Ghost Recon Predator",
	"Raving Rabbids: Travel in Time",
	"CSI: Unsolved!",
	"Michael Jackson: The Experience",
	"Sports Collection: 15 Sports to Master",
	"The Hollywood Squares",
	"Prince of Persia: The Sands of Time HD",
	"Prince of Persia: The Two Thrones HD",
	"Prince of Persia: Warrior Within HD",
	"Petz: Catz Family",
	"Zeit²",
	"Beyond Good & Evil HD",
	"The $1,000,000 Pyramid",
	"We Dare",
	"Fit in Six",
	"Dance on Broadway",
	"Assassin's Creed: Brotherhood",
	"Rayman 3D",
	"Family Feud: Decades",
	"Asphalt 6: Adrenaline",
	"Combat of Giants: Dinosaurs 3D",
	"Tom Clancy's Ghost Recon: Shadow Wars",
	"IL-2 Sturmovik: Cliffs of Dover",
	"Rabbids: Travel in Time 3D",
	"Tom Clancy's Splinter Cell 3D",
	"Might & Magic: Clash of Heroes",
	"Michael Jackson: The Experience",
	"The $1,000,000 Pyramid",
	"Child of Eden",
	"Cubic Ninja",
	"Outland",
	"Petz Fantasy 3D",
	"Beyond Good & Evil HD",
	"Call of Juarez: The Cartel",
	"Just Dance: Summer Party",
	"The Smurfs",
	"The Smurfs Dance Party",
	"From Dust",
	"The Smurfs & Co",
	"Tom Clancy's Splinter Cell Classic Trilogy HD",
	"From Dust",
	"Driver: Renegade",
	"Driver: San Francisco",
	"Call of Juarez: The Cartel",
	"From Dust",
	"Puzzler Mind Gym 3D",
	"TrackMania 2: Canyon",
	"Might & Magic: Clash of Heroes",
	"Child of Eden",
	"Driver: San Francisco",
	"Just Dance 3",
	"Battle of Giants: Dinosaurs Strike",
	"Might & Magic Heroes VI",
	"Family Feud: 2012 Edition",
	"PowerUp Heroes",
	"The Price is Right: Decades",
	"The Adventures of Tintin: The Secret of the Unicorn",
	"Imagine: Babyz",
	"Just Dance Kids 2",
	"Zoo Mania 3D",
	"James Noir's Hollywood Crimes",
	"MotionSports Adrenaline",
	"NCIS",
	"Raving Rabbids: Alive & Kicking",
	"Michael Jackson: The Experience 3D",
	"Puppies 3D",
	"The Black Eyed Peas Experience",
	"Self-Defense Training Camp",
	"Your Shape: Fitness Evolved 2012",
	"ABBA: You Can Dance",
	"Assassin's Creed: Revelations",
	"Imagine: Fashion Designer",
	"Rayman Origins",
	"The Price is Right: Decades",
	"Anno 2070",
	"Assassin's Creed: Revelations",
	"Just Dance 3",
	"Puzzler World 2012",
	"Prince of Persia Classic",
	"Tom Clancy's Splinter Cell HD",
	"Tom Clancy's Splinter Cell: Chaos Theory HD",
	"Tom Clancy's Splinter Cell: Pandora Tomorrow HD",
	"Asphalt: Injection",
	"Dungeon Hunter: Alliance",
	"Lumines: Electronic Symphony",
	"Rayman Origins",
	"Tom Clancy's Splinter Cell: Conviction",
	"I Am Alive",
	"Assassin's Creed: Recollection",
	"Shoot Many Robots",
	"MotoHeroz",
	"Rayman 3 HD",
	"Funky Barn 3D",
	"Horses 3D",
	"Just Dance: Best Of",
	"Rayman Origins",
	"I Am Alive",
	"Shoot Many Robots",
	"Trials Evolution",
	"Tom Clancy's Ghost Recon: Future Soldier",
	"Mad Riders",
	"Babel Rising",
	"Babel Rising 3D",
	"Babel Rising",
	"James Noir's Hollywood Crimes",
	"Just Dance: Greatest Hits",
	"Tom Clancy's Ghost Recon: Future Soldier",
	"The Expendables 2: Videogame",
	"Babel Rising",
	"The Expendables 2: Videogame",
	"I Am Alive",
	"Might & Magic: Duel of Champions",
	"Prince of Persia Classic",
	"NCIS 3D",
	"Rayman Jungle Run",
	"Puzzler World 2012",
	"Rayman Jungle Run",
	"Babel Rising 3D",
	"Just Dance 4",
	"Monster 4x4 3D",
	"Rocksmith",
	"Just Dance: Disney Party",
	"Imagine: Fashion Life",
	"Assassin's Creed III",
	"Assassin's Creed III: Liberation",
	"The Avengers: Battle for Earth",
	"Nutty Fluffies Rollercoaster",
	"Poptropica Adventures",
	"Rayman Origins",
	"Assassin's Creed III",
	"Rabbids Rumble",
	"The Hip Hop Dance Experience",
	"ESPN Sports Connection",
	"Just Dance 4",
	"Rabbids Land",
	"Your Shape: Fitness Evolved 2013",
	"ZombiU",
	"Assassin's Creed III",
	"Far Cry 3",
	"The Avengers: Battle for Earth",
	"Might & Magic: Duel of Champions",
	"Racquet Sports",
	"Nutty Fluffies Rollercoaster",
	"TrackMania 2: Stadium",
	"Trials Evolution: Gold Edition",
	"ShootMania Storm",
	"Far Cry 3: Blood Dragon",
	"Call of Juarez: Gunslinger",
	"Rayman Jungle Run",
	"Splinter Cell: Blacklist - Spider-Bot",
	"The Smurfs 2",
	"Might & Magic: Clash of Heroes",
	"Spartacus Legends",
	"The Smurfs & Co: Spellbound",
	"TrackMania 2: Valley",
	"The Smurfs 2",
	"Prince of Persia: The Shadow and the Flame",
	"Cloudberry Kingdom",
	"Splinter Cell: Blacklist - Spider-Bot",
	"Tom Clancy's Splinter Cell: Blacklist",
	"Rayman Legends",
	"Panzer General Online",
	"Anno Online",
	"Flashback",
	"Just Dance 2014",
	"Rabbids Big Bang",
	"Just Dance Kids 2014",
	"Rocksmith: All-new 2014 Edition",
	"Assassin's Creed IV: Black Flag",
	"Rocksmith: All-new 2014 Edition",
	"Rayman Fiesta Run",
	"Just Dance 2014",
	"Assassin's Creed IV: Black Flag",
	"Fighter Within",
	"Assassin's Creed IV: Black Flag",
	"Assassin's Creed Pirates",
	"Rabbids Big Bang",
	"Assassin's Creed: Liberation HD",
	"Rayman Fiesta Run",
	"Might & Magic X: Legacy",
	"MotoHeroz",
	"Far Cry Classic",
	"Rayman Legends",
	"South Park: The Stick of Truth",
	"Tom Clancy's Ghost Recon Phantoms",
	"Trials Frontier",
	"Trials Fusion",
	"Might & Magic X: Legacy",
	"Child of Light",
	"CSI: Hidden Crimes",
	"Watch Dogs",
	"Valiant Hearts: The Great War",
	"Child of Light",
	"Assassin's Creed: Memories",
	"Little Raiders: Robin's Revenge",
	"Just Dance Now",
	"Petz Beach",
	"Petz Countryside",
	"Just Dance 2015",
	"Valiant Hearts: The Great War",
	"Shape Up: Battle Run",
	"Assassin's Creed Rogue",
	"Assassin's Creed Unity",
	"Shape Up",
	"Tetris Ultimate",
	"Far Cry 4",
	"Far Cry 4: Arcade Poker",
	"Far Cry 4: Arena Master",
	"Rabbids Invasion: The Interactive TV Show",
	"Watch Dogs",
	"Anno: Build an Empire",
	"Valiant Hearts: The Great War",
	"Monopoly Deal",
	"Monopoly Plus",
	"The Crew",
	"Tetris Ultimate",
	"Monopoly Deal",
	"Monopoly Plus",
	"Horse Haven: World Adventures",
	"Risk",
	"Grow Home",
	"The Mighty Quest for Epic Loot",
	"Monkey King Escape",
	"Trivial Pursuit Live!",
	"Assassin's Creed Rogue",
	"Horse Haven: World Adventures",
	"Driver: Speedboat Paradise",
	"Assassin's Creed Chronicles: China",
	"Risk",
	"Tetris Ultimate",
	"Rabbids Appisodes",
	"Scrabble",
	"Boggle",
	"Toy Soldiers: War Chest",
	"Zombi",
	"Grow Home",
	"Care Bears: Belly Match",
	"The Smurfs 3D",
	"Gravity Falls: Legend of the Gnome Gemulets",
	"Just Dance 2016",
	"Just Dance: Disney Party 2",
	"Assassin's Creed Syndicate",
	"Anno 2205: Asteroid Miner",
	"Anno 2205",
	"Grow Home",
	"Assassin's Creed Syndicate",
	"Might & Magic: Heroes Online",
	"Tom Clancy's Rainbow Six Siege",
	"Tom Clancy's EndWar Online",
	"Rayman Adventures",
	"Just Dance Now",
	"Tetris Ultimate",
	"The Smurfs: Epic Run",
	"Assassin's Creed Chronicles: India",
	"Sandstorm: Pirate Wars",
	"Assassin's Creed Chronicles: Russia",
	"Rayman",
	"Far Cry Primal",
	"Assassin's Creed Identity",
	"Far Cry Primal",
	"Tom Clancy's The Division",
	"Rabbids Appisodes",
	"Rayman",
	"TrackMania Turbo",
	"The Smurfs: Epic Run",
	"Hungry Shark: World",
	"Rock Gods: Tap Tour",
	"Assassin's Creed Identity",
	"Trials of the Blood Dragon",
	"NCIS: Hidden Crimes",
	"Battleship",
	"Risk: Urban Assault",
	"Grow Up",
	"Uno",
	"Champions of Anteria",
	"Just Sing",
	"Rabbids Crazy Rush",
	"NCIS: Hidden Crimes",
	"Face Up: The Selfie Game",
	"Rocksmith: All-new 2014 Edition - Remastered",
	"Hungry Shark: World",
	"Eagle Flight",
	"Just Dance 2017",
	"Eagle Flight",
	"Assassin's Creed: The Ezio Collection",
	"Watch Dogs 2",
	"Steep",
	"Werewolves Within",
	"Uno",
	"City of Love: Paris",
	"For Honor",
	"Just Dance 2017",
	"Tom Clancy's Ghost Recon Wildlands",
	"TrackMania 2: Lagoon",
	"Star Trek: Bridge Crew",
	"Mario + Rabbids Kingdom Battle",
	"Monopoly Plus",
	"Rayman Legends: Definitive Edition",
	"Atomega",
	"Hungry Shark: World",
	"South Park: The Fractured But Whole",
	"Just Dance 2018",
	"Assassin's Creed Origins",
	"Monopoly",
	"Jeopardy!",
	"Uno",
	"Wheel of Fortune",
	"South Park: Phone Destroyer",
	"Ode",
	"South Park: The Stick of Truth",
	"Discovery Tour: Assassin's Creed - Ancient Egypt",
	"Hungry Dragon",
	"Assassin's Creed Rogue Remastered",
	"Far Cry 5",
	"South Park: The Fractured But Whole",
	"Far Cry 3: Classic Edition",
	"The Crew 2",
	"Hungry Shark: World",
	"The Settlers: History Edition",
	"Legendary Fishing",
	"Transference",
	"South Park: The Stick of Truth",
	"Assassin's Creed Odyssey",
	"Child of Light: Ultimate Edition",
	"Starlink: Battle for Atlas",
	"Just Dance 2019",
	"Jeopardy!",
	"Sports Party",
	"Trivial Pursuit Live!",
	"Wheel of Fortune",
	"Brawlhalla",
	"Valiant Hearts: The Great War",
	"Assassin's Creed: Rebellion",
	"Far Cry New Dawn",
	"Anno 1800",
	"Trials Rising",
	"Tom Clancy's The Division 2",
	"Space Junkies",
	"Growtopia",
	"Assassin's Creed III Remastered",
	"Starlink: Battle for Atlas",
	"Assassin's Creed III Remastered",
	"The Mighty Quest for Epic Loot",
	"Growtopia",
	"Discovery Tour: Assassin's Creed - Ancient Greece",
	"Rayman Mini",
	"Tom Clancy's Ghost Recon Breakpoint",
	"Rabbids: Coding!",
	"Rayman Mini",
	"Just Dance 2020",
	"Assassin's Creed Odyssey",
	"Just Dance 2020",
	"Might & Magic Heroes: Era of Chaos",
	"Assassin's Creed: The Rebel Collection",
	"Star Trek: Bridge Crew",
	"Tom Clancy's Ghost Recon Breakpoint",
	"Might & Magic: Chess Royale",
	"Tom Clancy's The Division 2",
	"The Crew 2",
	"TrackMania",
	"Brawlhalla",
	"Hyper Scape",
	"Tom Clancy's Elite Squad",
	"Rabbids: Coding!",
	"Rabbids Wild Race",
	"AGOS: A Game of Space",
	"Watch Dogs: Legion",
	"Assassin's Creed Valhalla",
	"Watch Dogs: Legion",
	"Assassin's Creed Valhalla",
	"Family Feud",
	"Just Dance 2021",
	"Watch Dogs: Legion",
	"Tom Clancy's Rainbow Six Siege",
	"Idle Restaurant Tycoon",
	"Immortals Fenyx Rising",
	"Scott Pilgrim vs. The World: The Game - Complete Edition",
	"Far Cry 6",
	"Discovery Tour: Viking Age",
	"Riders Republic",
	"Just Dance 2022",
	"Monopoly Madness",
	"Clash of Beasts",
	"Tom Clancy's Rainbow Six Extraction",
	"Assassin's Creed: The Ezio Collection",
	"Trivial Pursuit Live! 2",
	"Roller Champions",
	"Assassin's Creed Origins",
	"Discovery Tour: Viking Age",
	"Rabbids: Party of Legends",
	"Mario + Rabbids Sparks of Hope",
	"Skull & Bones",
	"Discovery Tour: Viking Age",
	"OddBallers",
	"Prince of Persia: The Sands of Time Remake",
	"Rocksmith+",
	"Roller Champions",
	"Tom Clancy's The Division Heartland",
	"Avatar: Frontiers of Pandora",
	"Assassin's Creed",
	"Assassin's Creed Infinity",
	"Beyond Good and Evil 2",
	"The Settlers",
	"Tom Clancy's Rainbow Six Mobile",
	"Tom Clancy's Splinter Cell",
	"Tom Clancy's The Division Resurgence",
	"XDefiant",
	"Untitled Star Wars game",
}
