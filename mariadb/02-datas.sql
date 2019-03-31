use music;

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




insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Undertow"), 1, "Intolerance", 294);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Undertow"), 2, "Prison Sex", 296);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Undertow"), 3, "Sober", 306);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Undertow"), 4, "Bottom", 433);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Undertow"), 5, "Crawl Away", 329);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Undertow"), 6, "Swamp Song", 331);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Undertow"), 7, "Undertow", 321);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Undertow"), 8, "4 Degrees", 363);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Undertow"), 9, "Flood", 465);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Undertow"), 10, "Disgustipated", 947);

insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "AEnima"), 1, "Stinkfist", 311);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "AEnima"), 2, "Eulogy", 509);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "AEnima"), 3, "H.", 363);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "AEnima"), 4, "Useful Idiots", 39);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "AEnima"), 5, "Forty-Six &2 ", 363);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "AEnima"), 6, "Message to Harry Manback", 113);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "AEnima"), 7, "Hooker with a Penis", 274);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "AEnima"), 8, "Intermission", 56);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "AEnima"), 9, "Jimmy", 324);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "AEnima"), 10, "Die Eier von Satan", 137);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "AEnima"), 11, "Pushit", 596);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "AEnima"), 12, "Cesaro Summability", 86);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "AEnima"), 13, "AEnima", 400);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "AEnima"), 14, "(-)Ions", 240);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "AEnima"), 15, "Third Eye", 827);

insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Lateralus"), 1, "The Grudge", 516);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Lateralus"), 2, "Eon Blue Apocalypse", 64);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Lateralus"), 3, "The Patient", 433);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Lateralus"), 4, "Mantra", 72);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Lateralus"), 5, "Schism", 407);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Lateralus"), 6, "Parabol", 184);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Lateralus"), 7, "Parabola", 363);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Lateralus"), 8, "Ticks & Leeches", 490);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Lateralus"), 9, "Lateralus", 564);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Lateralus"), 10, "Disposition", 286);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Lateralus"), 11, "Reflection", 667);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Lateralus"), 12, "Triad", 526);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Lateralus"), 13, "Faaip De Oaid", 159);

insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "10,000 Days"), 1, "Vicarious", 426);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "10,000 Days"), 2, "Jambi", 448);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "10,000 Days"), 3, "Wings for Marie, Pt. 1", 371);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "10,000 Days"), 4, "10,000 Days (Wings, Pt. 2)", 673);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "10,000 Days"), 5, "The Pot", 381);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "10,000 Days"), 6, "Lipan Conjuring", 71);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "10,000 Days"), 7, "Lost Keys (Blame Hofmann)", 226);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "10,000 Days"), 8, "Rosetta Stoned", 671);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "10,000 Days"), 9, "Intension", 441);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "10,000 Days"), 10, "Right in Two", 535);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "10,000 Days"), 11, "Viginti Tres", 302);

insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Facelift"), 1, "We Die Young", 152);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Facelift"), 2, "Man in the Box", 286);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Facelift"), 3, "Sea of Sorrow", 349);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Facelift"), 4, "Bleed the Freak", 241);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Facelift"), 5, "I Can't Remember", 222);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Facelift"), 6, "Love, Hate, Love", 386);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Facelift"), 7, "It Ain't Like That", 277);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Facelift"), 8, "Sunshine", 284);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Facelift"), 9, "Put You Down", 196);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Facelift"), 10, "Confusion", 344);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Facelift"), 11, "I Know Somethin('bout You)", 261);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Facelift"), 12, "Real Thing", 243);

insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Dirt"), 1, "Them Bones", 150);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Dirt"), 2, "Dam That River", 189);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Dirt"), 3, "Rain When I Die", 363);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Dirt"), 4, "Down in a Hole", 339);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Dirt"), 5, "Sickman", 331);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Dirt"), 6, "Rooster", 376);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Dirt"), 7, "Junkhead", 311);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Dirt"), 8, "Dirt", 318);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Dirt"), 9, "God Smack", 232);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Dirt"), 10, "Iron Gland", 43);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Dirt"), 11, "Hate to Feel", 316);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Dirt"), 12, "Angry Chair", 289);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Dirt"), 13, "Would?", 206);

insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Alice in Chains"), 1, "Grind", 284);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Alice in Chains"), 2, "Brush Away", 202);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Alice in Chains"), 3, "Sludge Factory", 432);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Alice in Chains"), 4, "Heaven Beside You", 327);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Alice in Chains"), 5, "Head Creeps", 388);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Alice in Chains"), 6, "Again", 245);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Alice in Chains"), 7, "Shame in You", 335);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Alice in Chains"), 8, "God Am", 248);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Alice in Chains"), 9, "So Close", 165);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Alice in Chains"), 10, "Nothin' Song", 340);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Alice in Chains"), 11, "Frog", 498);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Alice in Chains"), 12, "Over Now", 423);

insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Mars Red Sky"), 1, "Strong Reflection", 331);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Mars Red Sky"), 2, "Curse", 244);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Mars Red Sky"), 3, "Falls", 385);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Mars Red Sky"), 4, "Marble Sky", 361);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Mars Red Sky"), 5, "Way to Rome", 319);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Mars Red Sky"), 6, "Saddle Point", 254);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Mars Red Sky"), 7, "Up the Stairs", 478);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Mars Red Sky"), 8, "The Ravens Are Back", 282);

insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Stranded in Arcadia"), 1, "The Light Beyond", 484);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Stranded in Arcadia"), 2, "Hovering Satellites", 336);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Stranded in Arcadia"), 3, "Holy Mondays", 294);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Stranded in Arcadia"), 4, "Join the Race", 319);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Stranded in Arcadia"), 5, "Arcadia", 357);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Stranded in Arcadia"), 6, "Circles", 300);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Stranded in Arcadia"), 7, "Seen a Ghost", 438);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Stranded in Arcadia"), 8, "Beyond the Light", 154);

insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Apex III - Praise for the Burning Soul"), 1, "(Alien Grounds)", 226);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Apex III - Praise for the Burning Soul"), 2, "Apex III", 429);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Apex III - Praise for the Burning Soul"), 3, "The Whinery", 362);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Apex III - Praise for the Burning Soul"), 4, "Mindreader", 394);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Apex III - Praise for the Burning Soul"), 5, "Under The Hood", 305);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Apex III - Praise for the Burning Soul"), 6, "Friendly Fire", 287);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Apex III - Praise for the Burning Soul"), 7, "Prodigal Sun", 408);
insert into tracks(id_record, number, title, length)
    values((Select id from records where title = "Apex III - Praise for the Burning Soul"), 8, "Shot In Providence", 493);

