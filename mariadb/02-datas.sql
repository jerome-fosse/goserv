insert into artists (name, country) values ("Tool", "USA");
insert into artists (name, country) values ("Alice in Chains", "USA");
insert into artists (name, country) values ("Mars Red Sky", "France");

insert into records(title, id_artist, year, genre)
    values("Undertow", (select id from artists where name = "Tool"), 1993, "Metal progressif / Metal alternatif");
insert into records(title, id_artist, year, genre)
    values("AEnima", (select id from artists where name = "Tool"), 1996, "Metal progressif / Metal alternatif");
insert into records(title, id_artist, year, genre)
    values("Lateralus", (select id from artists where name = "Tool"), 2001, "Metal progressif");
insert into records(title, id_artist, year, genre)
    values("10,000 Days", (select id from artists where name = "Tool"), 2006, "Metal progressif");

insert into records(title, id_artist, year, genre)
    values("Facelift", (select id from artists where name = "Alice in Chains"), 1990, "Heavy Metal");
insert into records(title, id_artist, year, genre)
    values("Dirt", (select id from artists where name = "Alice in Chains"), 1992, "Heavy Metal");
insert into records(title, id_artist, year, genre)
    values("Alice in Chains", (select id from artists where name = "Alice in Chains"), 1995, "Heavy Metal");

insert into records(title, id_artist, year, genre)
    values("Mars Red Sky", (select id from artists where name = "Mars Red Sky"), 2011, "Stoner psychedelic");
insert into records(title, id_artist, year, genre)
    values("Stranded in Arcadia", (select id from artists where name = "Mars Red Sky"), 2014, "Stoner psychedelic");
insert into records(title, id_artist, year, genre)
    values("Apex III - Praise for the Burning Soul", (select id from artists where name = "Mars Red Sky"), 2016, "Stoner psychedelic");




insert into tracks(id_record, number, title)
    values((Select id from records where title = "Undertow"), 1, "Intolerance");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Undertow"), 2, "Prison Sex");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Undertow"), 3, "Sober");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Undertow"), 4, "Bottom");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Undertow"), 5, "Crawl Away");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Undertow"), 6, "Swamp Song");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Undertow"), 7, "Undertow");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Undertow"), 8, "4 Degrees");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Undertow"), 9, "Flood");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Undertow"), 10, "Disgustipated");

insert into tracks(id_record, number, title)
    values((Select id from records where title = "AEnima"), 1, "Stinkfist");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "AEnima"), 2, "Eulogy");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "AEnima"), 3, "H.");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "AEnima"), 4, "Useful Idiots");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "AEnima"), 5, "Forty-Six &2 ");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "AEnima"), 6, "Message to Harry Manback");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "AEnima"), 7, "Hooker with a Penis");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "AEnima"), 8, "Intermission");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "AEnima"), 9, "Jimmy");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "AEnima"), 10, "Die Eier von Satan");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "AEnima"), 11, "Pushit");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "AEnima"), 12, "Cesaro Summability");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "AEnima"), 13, "AEnima");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "AEnima"), 14, "(-)Ions");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "AEnima"), 15, "Third Eye");

insert into tracks(id_record, number, title)
    values((Select id from records where title = "Lateralus"), 1, "The Grudge");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Lateralus"), 2, "Eon Blue Apocalypse");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Lateralus"), 3, "The Patient");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Lateralus"), 4, "Mantra");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Lateralus"), 5, "Schism");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Lateralus"), 6, "Parabol");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Lateralus"), 7, "Parabola");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Lateralus"), 8, "Ticks & Leeches");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Lateralus"), 9, "Lateralus");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Lateralus"), 10, "Disposition");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Lateralus"), 11, "Reflection");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Lateralus"), 12, "Triad");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Lateralus"), 13, "Faaip De Oaid");

insert into tracks(id_record, number, title)
    values((Select id from records where title = "10,000 Days"), 1, "Vicarious");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "10,000 Days"), 2, "Jambi");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "10,000 Days"), 3, "Wings for Marie, Pt. 1");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "10,000 Days"), 4, "10,000 Days (Wings, Pt. 2)");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "10,000 Days"), 5, "The Pot");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "10,000 Days"), 6, "Lipan Conjuring");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "10,000 Days"), 7, "Lost Keys (Blame Hofmann)");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "10,000 Days"), 8, "Rosetta Stoned");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "10,000 Days"), 9, "Intension");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "10,000 Days"), 10, "Right in Two");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "10,000 Days"), 11, "Viginti Tres");

insert into tracks(id_record, number, title)
    values((Select id from records where title = "Facelift"), 1, "We Die Young");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Facelift"), 2, "Man in the Box");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Facelift"), 3, "Sea of Sorrow");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Facelift"), 4, "Bleed the Freak");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Facelift"), 5, "I Can't Remember");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Facelift"), 6, "Love, Hate, Love");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Facelift"), 7, "It Ain't Like That");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Facelift"), 8, "Sunshine");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Facelift"), 9, "Put You Down");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Facelift"), 10, "Confusion");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Facelift"), 11, "I Know Somethin('bout You)");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Facelift"), 12, "Real Thing");

insert into tracks(id_record, number, title)
    values((Select id from records where title = "Dirt"), 1, "Them Bones");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Dirt"), 2, "Dam That River");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Dirt"), 3, "Rain When I Die");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Dirt"), 4, "Down in a Hole");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Dirt"), 5, "Sickman");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Dirt"), 6, "Rooster");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Dirt"), 7, "Junkhead");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Dirt"), 8, "Dirt");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Dirt"), 9, "God Smack");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Dirt"), 10, "Iron Gland");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Dirt"), 11, "Hate to Feel");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Dirt"), 12, "Angry Chair");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Dirt"), 13, "Would?");

insert into tracks(id_record, number, title)
    values((Select id from records where title = "Alice in Chains"), 1, "Grind");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Alice in Chains"), 2, "Brush Away");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Alice in Chains"), 3, "Sludge Factory");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Alice in Chains"), 4, "Heaven Beside You");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Alice in Chains"), 5, "Head Creeps");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Alice in Chains"), 6, "Again");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Alice in Chains"), 7, "Shame in You");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Alice in Chains"), 8, "God Am");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Alice in Chains"), 9, "So Close");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Alice in Chains"), 10, "Nothin' Song");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Alice in Chains"), 11, "Frog");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Alice in Chains"), 12, "Over Now");

insert into tracks(id_record, number, title)
    values((Select id from records where title = "Mars Red Sky"), 1, "Strong Reflection");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Mars Red Sky"), 2, "Curse");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Mars Red Sky"), 3, "Falls");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Mars Red Sky"), 4, "Marble Sky");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Mars Red Sky"), 5, "Way to Rome");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Mars Red Sky"), 6, "Saddle Point");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Mars Red Sky"), 7, "Up the Stairs");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Mars Red Sky"), 8, "The Ravens Are Back");

insert into tracks(id_record, number, title)
    values((Select id from records where title = "Stranded in Arcadia"), 1, "The Light Beyond");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Stranded in Arcadia"), 2, "Hovering Satellites");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Stranded in Arcadia"), 3, "Holy Mondays");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Stranded in Arcadia"), 4, "Join the Race");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Stranded in Arcadia"), 5, "Arcadia");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Stranded in Arcadia"), 6, "Circles");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Stranded in Arcadia"), 7, "Seen a Ghost");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Stranded in Arcadia"), 8, "Beyond the Light");

insert into tracks(id_record, number, title)
    values((Select id from records where title = "Apex III - Praise for the Burning Soul"), 1, "(Alien Grounds)");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Apex III - Praise for the Burning Soul"), 2, "Apex III");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Apex III - Praise for the Burning Soul"), 3, "The Whinery");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Apex III - Praise for the Burning Soul"), 4, "Mindreader");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Apex III - Praise for the Burning Soul"), 5, "Under The Hood");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Apex III - Praise for the Burning Soul"), 6, "Friendly Fire");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Apex III - Praise for the Burning Soul"), 7, "Prodigal Sun");
insert into tracks(id_record, number, title)
    values((Select id from records where title = "Apex III - Praise for the Burning Soul"), 8, "Shot In Providence");

