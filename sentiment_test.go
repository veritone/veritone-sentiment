package sentiment

import "testing"

var model Models

func init() {
	var err error

	//model, err = Train()

	model, err = Restore()
	if err != nil {
		panic(err.Error())
	}
}

func TestPositiveWordSentimentShouldPass1(t *testing.T) {
	t.Parallel()

	w := []string{"happy", "love", "happiness", "humanity", "awesome", "great", "fun", "super", "trust", "fearless", "creative", "dream", "good", "compassion", "joy", "independent", "success"}
	for _, word := range w {
		s := model.SentimentAnalysis(word, English)
		if s.Score < uint8(5) {
			t.Errorf("Sentiment of < %v > (returned %v) should be greater than 0.5!\n", word, s)
		} else {
			t.Logf("Sentiment of < %v > valid\n\tReturned %v\n", word, s)
		}

		if len(s.Words) == 0 {
			t.Fatalf("Words returned should have individual sentiment\n\t%v\n", s)
		}

		for _, score := range s.Words {
			if score.Score < uint8(5) {
				t.Errorf("Probability of < %v > (returned %v) should be greater than 5 always!\n", score.Word, score.Score)
			} else {
				t.Logf("Probability of < %v > valid\n\tReturned %v\n", score.Word, score.Score)
			}
			if score.Probability == 0 {
				t.Errorf("Probability of < %v > (returned %v) should be greater than 0.0 always!\n", score.Word, score.Score)
			} else {
				t.Logf("Probability of < %v > valid\n\tReturned %v\n", score.Word, score.Score)
			}
		}
	}
}

func TestNegativeWordSentimentShouldPass1(t *testing.T) {
	t.Parallel()

	w := []string{"not", "resent", "deplorable", "bad", "terrible", "hate", "scary", "terrible", "concerned", "wrong", "rude!!!", "sad", "horrible", "unimpressed", "useless", "offended", "disrespectful"}
	for _, word := range w {
		s := model.SentimentAnalysis(word, English)
		if s.Score > uint8(5) {
			t.Errorf("Sentiment of < %v > (returned %v) should be less than 0.5!\n", word, s)
		} else {
			t.Logf("Sentiment of < %v > valid\n\tReturned %v\n", word, s)
		}

		if len(s.Words) == 0 {
			t.Fatalf("Words returned should have individual sentiment\n\t%v\n", s)
		}

		for _, score := range s.Words {
			if score.Score > uint8(5) {
				t.Errorf("Probability of < %v > (returned %v) should be less than 0.5 always!\n", score.Word, score.Score)
			} else {
				t.Logf("Probability of < %v > valid\n\tReturned %v\n", score.Word, score.Score)
			}
			if score.Probability == 0 {
				t.Errorf("Probability of < %v > (returned %v) should be greater than 0.0 always!\n", score.Word, score.Score)
			} else {
				t.Logf("Probability of < %v > valid\n\tReturned %v\n", score.Word, score.Score)
			}
		}
	}
}

func TestPositiveSentenceSentimentShouldPass1(t *testing.T) {
	t.Parallel()

	w := []string{
		"I had an awesome time watching this movie",
		"Sometimes I like to say hello to strangers and it's fun",
		"America needs to support the middle class",
		"Harry Potter is a great movie!",
		"The quest for love is a long one, but it ends in happiness",
		"You are a great person",
		"I love the way you can't talk",
		"You are a caring person",
		"I'm quite ambitious, and this job would be a great opportunity for me!",
		"I'm pretty easy-going.",
		"I find it easy to get along with people",
		"I am very hard-working",
		"I'm very methodical and take care over my work",
	}

	for _, sentence := range w {
		s := model.SentimentAnalysis(sentence, English)
		if s.Score < uint8(5) {
			t.Errorf("Sentiment of sentence < %v > (returned %v) should be greater than 0.5!\n", sentence, s)
		} else {
			t.Logf("Sentiment of sentence < %v > is valid.\n\tReturned %v\n", sentence, s)
		}
	}
}

func TestNegativeSentenceSentimentShouldPass1(t *testing.T) {
	t.Parallel()

	w := []string{
		"Jeffery is not a fun guy",
		"I don't enjoy saying hello to strangers",
		"I would compare this person to Donald Trump (ARGH!!!!!) Blind and ignorant!",
		"I'm happy here. I think so, at least.",
		"I hate random people",
		"I don't like your tone right now",
		"I'm not sure you know what you are talking about",
		"The rapture is upon us! Behold!!",
		"I think the growing consensus that China is somehow not a fair player is a bad thing overall",
		"Michelle Bachmann is a total idiot!",
		"How could you say such a thing!",
		"I hate banannas almost as much as I don't love you",
		"Dinner last night sucked",
	}

	for _, sentence := range w {
		s := model.SentimentAnalysis(sentence, English)
		if s.Score > uint8(5) {
			t.Errorf("Sentiment of sentence < %v > (returned %v) should be less than 0.5!\n", sentence, s)
		} else {
			t.Logf("Sentiment of sentence < %v > is valid.\n\tReturned %v\n", sentence, s)
		}
	}
}

func TestSentimentAnalysisShouldPass1(t *testing.T) {
	t.Parallel()
	transcript := `On the cross to put away sin by the sacrifice of himself told ...so infinite are on this is so great that only the sacrifice of jesus christ god son could pay for the enormously of arson thank god. He said his son to die for your ...you could be  and blameless before this is john macarthur praying you're continuing to be corporate Now let's get a check of traffic with charlie simon's ...into chaos by a traffic ...center got a problem ...five northbound if your past ninety nine ...just went past the airport and you're ...the woodland watch out in the left lane we have an accident at old river road traffic is backed up the vietnam veterans memorial bridge so far and getting slower by the second incident cleared fifty eastbound eldorado hills boulevard that's good news ninety nine southbound shoulders blocked by an accident at ...and ...capital city freeway the business eighty portion it's got its usual stop-and-go happening right about ...street until you get past E street that ...driving arbitrarily simon's seven ten K FI a joins us now for basic gospel with bob davis and richard piper recorded earlier for broadcasted this time on seven ten K ...keeping faith in america.</p> Fellow everybody with richard piper I bought a and this is basic gossip of ...dedicated to helping ...loop ...good ... If you level bible question or ...because we ... ...recall right ...number's eight four three two seven four two eight four three two seven four two we're lives off the air and online bright out you can get ...could answer your question or discuss the issues ...important in your life so we both of those fault line in the toll-free number four three two seven four two we'd love to hear from you right now it's basic gospel everybody now here's richard piper thanks bob almost long we've been the studying the< idea of freedom why it's so important how do you find it what's the source of it I don't think we could talk about freedom enough because everything in`

	analysis := model.SentimentAnalysis(transcript, English)

	if analysis.Score < uint8(5) {
		t.Errorf("Analysis of transcript should be greater than 0.5\n\treturned %v\n", analysis.Score)
	} else {
		t.Logf("Analysis of transcript was valid\n\treturned %v\n", analysis.Score)
	}

	if analysis.Words == nil {
		t.Errorf("Analysis of transcript returned nil words array!\n")
	} else {
		t.Logf("Analysis of transcript retuned valid word array\n\treturned %+v\n", analysis.Words)
	}

	if analysis.Sentences == nil {
		t.Errorf("Split sentence analysis should not be nil!\n")
	} else {
		t.Logf("Split sentence analysis returned valid sentence arrat\n\treturned %+v\n", analysis.Sentences)
	}
}

// From Haruki Murakami in Norwegian Wood
func TestSentimentAnalysisShouldPass2(t *testing.T) {
	t.Parallel()
	transcript := `“I like the ferry to Hokkaido. And I have no desire to fly through the air," she said. I accompanied her to Ueno Station. She carried her guitar and I carried her suitcase. We sat on a platform bench waiting for the train to pull in. Reiko wore the same tweed jacket and white trousers she had on when she arrived in Tokyo.`

	analysis := model.SentimentAnalysis(transcript, English)

	if analysis.Score < uint8(5) {
		t.Errorf("Analysis of transcript should be greater than 0.5\n\treturned %v\n", analysis.Score)
	} else {
		t.Logf("Analysis of transcript was valid\n\treturned %v\n", analysis.Score)
	}

	if analysis.Words == nil {
		t.Errorf("Analysis of transcript returned nil words array!\n")
	} else {
		t.Logf("Analysis of transcript retuned valid word array\n\t returned %+v\n", analysis.Words)
	}
}

// Paul Graham essay on immigration laws is _slightly_ negative
func TestSentimentAnalysisShouldPass3(t *testing.T) {
	t.Parallel()
	transcript := `The anti-immigration people have to invent some explanation to account for all the effort technology companies have expended trying to make immigration easier. So they claim it's because they want to drive down salaries. But if you talk to startups, you find practically every one over a certain size has gone through legal contortions to get programmers into the US, where they then paid them the same as they'd have paid an American. Why would they go to extra trouble to get programmers for the same price? The only explanation is that they're telling the truth: there are just not enough great programmers to go around`

	analysis := model.SentimentAnalysis(transcript, English)

	if analysis.Score > uint8(5) {
		t.Errorf("Analysis of transcript should be less than 0.5\n\treturned %v\n", analysis.Score)
	} else {
		t.Logf("Analysis of transcript was valid\n\treturned %v\n", analysis.Score)
	}

	if analysis.Words == nil {
		t.Errorf("Analysis of transcript returned nil words array!\n")
	} else {
		t.Logf("Analysis of transcript retuned valid word array\n\t returned %+v\n", analysis.Words)
	}
}

// From Haruki Murakami - Kafka On The Shore
func TestSentimentAnalysisShouldPass4(t *testing.T) {
	t.Parallel()
	transcript := `I'm inside the cafeteria sipping a free cup of hot tea when this young girl comes over and plunks herself down on the plastic seat next to me. In her right hand she has a paper cup of hot coffee she bought from a vending machine, the steam rising up from it, and in her left hand she's holding a small container with sandwiches inside—another bit of vending-machine gourmet fare, by the looks of it. She's kind of funny looking. Her face is out of balance—broad forehead, button nose, freckled cheeks, and pointy ears. A slammed-together, rough sort of face you can't ignore. Still, the whole package isn't so bad. For all I know maybe she's not so wild about her own looks, but she seems comfortable with who she is, and that's the important thing. There's something childish about her that has a calming effect, at least on me. She isn't very tall, but has good-looking legs and a nice bust for such a slim body.`

	analysis := model.SentimentAnalysis(transcript, English)

	if analysis.Score > uint8(5) {
		t.Errorf("Analysis of transcript should be less than 0.5\n\treturned %v\n", analysis.Score)
	} else {
		t.Logf("Analysis of transcript was valid\n\treturned %v\n", analysis.Score)
	}

	if analysis.Words == nil {
		t.Errorf("Analysis of transcript returned nil words array!\n")
	} else {
		t.Logf("Analysis of transcript retuned valid word array\n\t returned %+v\n", analysis.Words)
	}
}

// From F. Scott Fitzgerald - Great Gatsby
func TestSentimentAnalysisShouldPass5(t *testing.T) {
	t.Parallel()
	transcript := `He smiled understandingly- much more than understandingly. It was one of those rare smiles with a quality of eternal reassurance in it, that you may come across four or five times in life. It faced–or seemed to face–the whole eternal world for an instant, and then concentrated on you with an irresistible prejudice in your favor. It understood you just as far as you wanted to be understood, believed in you as you would like to believe in yourself, and assured you that it had precisely the impression of you that, at your best, you hoped to convey.`

	analysis := model.SentimentAnalysis(transcript, English)

	if analysis.Score < uint8(5) {
		t.Errorf("Analysis of transcript should be greater than 0.5\n\treturned %v\n", analysis.Score)
	} else {
		t.Logf("Analysis of transcript was valid\n\treturned %v\n", analysis.Score)
	}

	if analysis.Words == nil {
		t.Errorf("Analysis of transcript returned nil words array!\n")
	} else {
		t.Logf("Analysis of transcript retuned valid word array\n\t returned %+v\n", analysis.Words)
	}
}

func TestSentimentAnalysisShouldPass6(t *testing.T) {
	t.Parallel()
	transcript := `At the station I pop into the first little diner that catches my eye, and eat my fill of udon. Born and “raised in Tokyo, I haven't had much udon in my life. But now I'm in Udon Central—Shikoku—and confronted with noodles like nothing I've ever seen. They're chewy and fresh, and the soup smells great, really fragrant. And talk about cheap. It all tastes so good I order seconds, and for the first time in who knows how long, I'm happily stuffed. Afterward I plop myself down on a bench in the plaza next to the station and gaze up at the sunny sky. I'm free, I remind myself. Like the clouds floating across the sky, I'm all by myself, totally free.`

	analysis := model.SentimentAnalysis(transcript, English)

	if analysis.Score < uint8(5) {
		t.Errorf("Analysis of transcript should be greater than 0.5\n\treturned %v\n", analysis.Score)
	} else {
		t.Logf("Analysis of transcript was valid\n\treturned %v\n", analysis.Score)
	}

	if analysis.Words == nil {
		t.Errorf("Analysis of transcript returned nil words array!\n")
	} else {
		t.Logf("Analysis of transcript retuned valid word array\n\t returned %+v\n", analysis.Words)
	}
}

func TestSentimentAnalysisShouldPass7(t *testing.T) {
	t.Parallel()
	transcript := `I am a relatively happy sentence!`

	analysis := model.SentimentAnalysis(transcript, English)

	if analysis.Score < uint8(5) {
		t.Errorf("Analysis of transcript should be greater than 0.5\n\treturned %+v\n", analysis)
	} else {
		t.Logf("Analysis of transcript was valid\n\treturned %v\n", analysis.Score)
	}

	if analysis.Words == nil {
		t.Errorf("Analysis of transcript returned nil words array!\n")
	} else {
		t.Logf("Analysis of transcript retuned valid word array\n\t returned %+v\n", analysis.Words)
	}

	if analysis.Sentences != nil {
		t.Errorf("Split sentence analysis should be nil!\n\t%v\n", analysis.Sentences)
	} else {
		t.Logf("Split sentence analysis returned nil sentence array (valid)\n")
	}
}

// Donald Trump snippet from his annoucement presidential speech
func TestAssholeSentimentAnalysisShouldPass1(t *testing.T) {
	t.Parallel()
	transcript := `Thank you. It's true, and these are the best and the finest. When Mexico sends its people, they're not sending their best. They're not sending you. They're not sending you. They're sending people that have lots of problems, and they're bringing those problems with us. They're bringing drugs. They're bringing crime. They're rapists. And some, I assume, are good people. But I speak to border guards and they tell us what we're getting. And it only makes common sense. It only makes common sense. They're sending us not the right people. It's coming from more than Mexico. It's coming from all over South and Latin America, and it's coming probably -- probably -- from the Middle East. But we don't know. Because we have no protection and we have no competence, we don't know what's happening. And it's got to stop and it's got to stop fast.`

	analysis := model.SentimentAnalysis(transcript, English)

	if analysis.Score > uint8(5) {
		t.Errorf("Analysis of transcript should be less than 0.5\n\treturned %v\n", analysis.Score)
	} else {
		t.Logf("Analysis of transcript was valid\n\treturned %v\n", analysis.Score)
	}

	if analysis.Words == nil {
		t.Errorf("Analysis of transcript returned nil words array!\n")
	} else {
		t.Logf("Analysis of transcript retuned valid word array\n\t returned %+v\n", analysis.Words)
	}
}

// Last part of GOP primary debate
func TestAssholeSentimentAnalysisShouldPass2(t *testing.T) {
	t.Parallel()
	transcript := `MACCALLUM: All right.  Lindsey Graham, this conversation will no doubt go to the war on women, and that cutting funding to this group could be a very broad brush against all of you or anybody who will hold this nomination as being against women's health, against these organizations that people will say provide positive things for many women.  GRAHAM: I don't think it's a war on women for all of us as Americans to stand up and stop harvesting organs from little babies. Let's take the money that we would give to Planned Parenthood and put it in women's health care without having to harvest the organs of the unborn. The only way we're going to defund Planned Parenthood is have a pro-life president.  You want to see a war on women? Come with me to Iraq and Afghanistan, folks. I've been there 35 times. I will show you what they do to women. These mythical Arab armies that my friends talk about that are going to protect us don't exist. If I am president of the United States, we're going to send soldiers back to Iraq, back to Syria, to keep us from being attacked here and keep soldiers in Afghanistan because we must.  I cannot tell you how much our nation is threatened and how we need a commander in chief who understands the threats to this nation.  If you're running for president of the United States and you do not understand that we cannot defend this nation without more of our soldiers over there, you are not ready for this job.  HEMMER: Thank you, Senator.  Executive power. It appears that you all have a little bit of an issue with it at the moment. I want to move through this as quickly as I can, from stage left to stage right.  On the second day of his presidency, January 22nd 2009, President Obama signed one of his first executive orders. That was close Guantanamo Bay in Cuba. Still open today. What would be your first executive order?  Governor Gilmore, start.  GILMORE: Well, it's not a matter of what the first executive order would be, Bill. The matter is what orders exist now that shouldn't exist?  The president has done an executive order with respect to illegal immigration that is illegal. Illegal. And it creates a -- a contempt for the law, for the rule of law. If i were the president of the United States, I would go and look at every executive order that exists right now and determine which ones want to be voided, because the president shouldn't be legislating: not through that vehicle or any other. We should be relying upon the leadership of the Congress to pass the laws.  HEMMER: Senator Graham.  GRAHAM: Change the Mexico City policy, not take one dime of taxpayer money to fund abortion organizations overseas, and restore the NSA that's been gutted. We're going dark when it comes to detecting the next attack. We have gutted our ability to detect the next attack. And I would not stand for that as president of the United States. I would take the fight to these guys, whatever it took, as long as it took.  HEMMER: Governor Jindal, your first executive order would be in the White House would be what?  JINDAL: To repeal these unconstitutional illegal orders, whether it's amnesty or whether it's this president going around the Congress, whether it's in Obamacare, to restore the rule of law. I'd also go after these sanctuary cities, do everything we can to make sure that we are not -- that we are actually prosecuting and cutting off funding for cities that are harboring illegal aliens, and then finally making sure the IRS is not going after conservative or religious groups.  I would sign an executive order protecting religious liberty, our first amendment rights, so Christian business owners and individuals don't face discrimination for having a traditional view of marriage.  HEMMER: Governor Perry.  PERRY: It'll be a pretty busy day, but that Iran negotiation is going to be torn up on day one. We're going to start the process of securing that border. I'm also going to take a bottle of White-Out with me to get started on all those executive orders that Mr. Obama has put his name to.  HEMMER: That will be a long day.  PERRY: It will be a long day.  HEMMER: Senator Santorum?  SANTORUM: Just ditto to that.  We're going to suspend -- I've -- I've said this for four years. We're going to suspend and repeal every executive order, every regulation that cost American jobs and is -- is -- is impacting our freedom.  And second, the First Amendment Defense Act, which is protecting religious liberty, if it's not passed by then, which I suspect it won't, because the president will veto it, I will institute an executive order to make sure that people of faith are not being -- not being harassed and persecuted by the federal government for standing up for the religious beliefs.  HEMMER: First order, Carly Fiorina?  FIORINA: I agree with my colleagues. We need to begin by undoing -- I would begin by undoing a whole set of things that President Obama has done, whether it's illegal amnesty or this latest round of EPA regulations. But let me go back to something that's very important. We have been debating right here the core difference between conservatism and progressivism.  Conservatives, I am a conservative because I believe no one of us is any better than any other one of us. Every one of us is gifted by God, whether it is those poor babies being picked over or it's someone whose life is tangled up in a web of dependence.  Progressives don't believe that. They believe some are smarter than others, some are better than others, so some are going to need to take care of others.  That is the fight we have to have, and we have to undo a whole set of things that President Obama has done that get at the heart of his disrespect and disregard for too many Americans.  HEMMER: Governor Pataki?  PATAKI: Bill, I defeated Mario Cuomo. In the first day in office, my first executive order, I revoked every one of the executive orders that he had -- he had enacted over the prior 12 years. I would do that to Barack Obama's executive orders.  But I'd sign a second one, as I did in New York, as well, having a hard hiring freeze on adding new employees except for the military or defense-related positions. I'd sign that executive order.  When I left the workforce, New York State had been reduced by over 15 percent. We can do that in Washington. I will do that in Washington.  HEMMER: Thank you all.  MCCALLUM: Moving on to the next question, President Obama promised hope and change for the country, yet 60 percent of Americans are not satisfied with the shape that the country is in right now. Many think that America has lost its "can do" spirit and that it's not the nation that it once was.  Ronald Reagan was confronted with a similar atmosphere, and he said that it could be morning in America again. JFK said it was a new frontier. FDR said that we had nothing to fear but fear itself.  On this level, Carly Fiorina, can you inspire this nation?  FIORINA: This is a great nation. It is a unique nation in all of human history and on the face of the planet, because here, our founders believed that everyone has a right to fulfill their potential and that that right --they called it life, liberty, the pursuit of happiness -- comes from God and cannot be taken away by government.  We have arrived at a point in our nation's history where the potential of this nation and too many Americans is being crushed by the weight, the power, the cost, the complexity, the ineptitude, the corruption of the federal government, and only someone who will challenge the status quo of Washington, D.C. can lead the resurgence of this great nation.  I will do that.  MCCALLUM: We're talking about tapping into historic levels of leadership and lifting the nation in this kind of way that we're discussing.  So Senator Santorum, how would you do it?  SANTORUM: I came to Washington, D.C. in 1990. That sounds like a long time ago. It was. It was 25 years ago, and I came by defeating the Democratic incumbent. I came as a reformer.  I started the Gang of Seven, and it led to the overtaking of the 40-year Democratic rule of Congress, because I didn't -- I stood up to the old-boy network in Washington, D.C. because I believed that Washington was not the solution, that Washington was the problem, just like Ronald Reagan said. I was a child of Ronald Reagan.  And I went there, and for 16 years, I fought the insiders and was able to get things done. That's the difference. We need to elect someone who will stand with the American people, who understands its greatness, who understands what an open economy and freedom is all about, but at the same time, has a record of being able to get things done in Washington like we've never seen before.  Reforms, everything from moral and cultural issues to economic issues. Those of you health savings accounts. Health savings accounts are something that we introduced. It's a private-sector solution that believes in freedom, not Obamacare that believes in government control.  SANTORUM: Those are the things we brought, and we were able to get things done. If you want someone who's not going to divide Washington, but gets things done, then you should make me your president.  HEMMER: Thank you, senator.  MACCALLUM: (inaudible) Lindsey Graham?  GRAHAM: Thank you.  First thing I'd tell the American people, whatever it takes to defend our nation, I would do.  To the 1 percent who have been fighting this war for over a decade, I'd try my best to be a commander-in-chief worthy of your sacrifice.  We're going to lose Social Security and Medicare if Republicans and Democrats do not come together and find a solution like Ronald Reagan and Tip O'Neill. I will be the Ronald Reagan if I can find a Tip O'Neill.  When I was 21, my mom died. When I was 22, my dad died. We owned a liquor store, restaurant, bar and we lived in the back. Every penny we needed from -- every penny we got from Social Security, because my sister was a minor, we needed. Today, I'm 60, I'm not married, I don't have any kids. I would give up some Social Security to save a system that Americans are going to depend on now and in the future.  Half of American seniors would be in poverty without a Social Security check. If you make your president, I'm going to put the country ahead of the party. I'm going to do what it takes to defend this nation. This nation has been great to me, and that's the only way I know to pay you back.  MACCALLUM: Thank you.  HEMMER: Thank you, Senator. I need a two-word answer to the following query. In 2008, then-Senator Barack Obama described Hillary Clinton as, quote, "likable enough," end quote. What two words would you use to describe the Democratic frontrunner? Governor Pataki to start.  PATAKI: Divisive and with no vision. No vision at all. HEMMER: Wow. Carly Fiorina.  FIORINA: Not trustworthy. No accomplishment.  UNKNOWN: Secretive and untrustworthy.  PERRY: Well, let's go with three. Good at email.  HEMMER: Governor Jindal?  JINDAL: Socialist and government dependent.  GRAHAM: Not the change we need at a time we need it.  HEMMER: Governor?  GILMORE: Professional politician that can't be trusted.  HEMMER: Not a lot of compliments. To be continued.  MACCALLUM: So every candidate will have the opportunity to make a closing statement tonight. Each candidate will have 30 second for that. And we start with Governor Perry.  PERRY: Well, this is going to be a show me, don't tell me election. I think America is just a few good decisions and a leadership change at the top away from the best years we've ever had. And I think that the record of the governor of the last 14 years of the 12th largest economy in the world is just the medicine America is looking for.  1.5 million jobs created during the worst economic time this country has had since the Great Depression while the rest of the country lost 400,000 jobs. We're talking about a state that moved graduation rates forward from 27th in the nation to second-highest. As a matter of fact, if you're Hispanic or African-American in Texas, you have the number one high school graduation rates in America.  Americans are looking for somebody that's going to give them, and there is a place in this country over the last eight years in particular that talked about hope every day, and they didn't just talk about it, they delivered it. And that was the state of Texas. And if we can do that in Texas, that 12th largest economy in the world, we can do it in America.  Our best days are in front of us. We can reform those entitlements, we can change that corporate tax code and lower it. We can put America back on track on a growth level and a growth rate that we've never seen in the history of this country. Manufacturing will flow back into this country. It just needs a corporate executive type at the top that's done it before. And I will suggest to you nobody's done it like Rick Perry has done it over the last eight years. And if you elect me president, we will bring incredible growth back to this country. And as someone who's worn the uniform of the country, that's how we build our military back up.  HEMMER: Thank you Governor. Senator Santorum?  SANTORUM: I'll tell you how optimistic I am about America. Karen and I have seven children. You don't have seven children and bring them into this world if you're not optimistic about the future of this country.  I am, but people are upset, and they're upset for a reason about the future of this country. Donald Trump actually seized on it when he talked about immigration. And I think the reason he did is because immigration is sort of an example of what's broken and what's wrong in Washington, D..C.  You see, you have one side, the Democrats, and with immigration, all they care about is votes. They don't care about American workers, they just care about bringing as many people in so they can get as many votes as they can. ON the other side, you have so many Republicans, and what do they care about? Helping business make profits. There's nobody out there looking out for the American worker.  I'm looking out for the American worker. I'm the only one on this stage who has a plan that's actually reduced -- actually going to reduce immigration. Actually going to do something to help the American worker. And you combine that with a plan to make manufacturing -- this country number one in manufacturing, you've got someone who's going to help revitalize and give hope to America, the place -- the place is that is the most hopeless today.  That's why I ask for your support for president.  HEMMER: All right. Senator thank you.  MACCALLUM: Governor Jindal?  JINDAL: You know, we've got a lot of great talkers running for president. We've already got a great talker in the White House. We cannot afford four more years of on the job training. We need a doer, not a talker. We also need a nominee, a candidate who will endorse our own principles.  Jeb Bush says we've got to be willing to lose the primary in order to win the general. Let me translate that for you. That's the establishment telling us to hide our conservative principles to get the left and the media to like us. That never works. If we do that again, we will lose again, we will deserve to lose again.  One principle, for example, we've got to embrace is on immigration. We must insist on assimilation -- immigration without assimilation is an invasion. We need to tell folks who want to come here, they need to come here legally. They need to learn English, adopt our values, roll up their sleeves and get to work.  I'm tired of the hyphenated Americans and the division. I've got the backbone, I've got the band width, I've got the experience to get us through this. I'm asking folks not just to join my campaign, but join a cause. It is time to believe in America again. MACCALLUM: Thank you, Governor.  HEMMER: Carly Fiorina, closing statement.  FIORINA: Hillary Clinton lies about Benghazi, she lies about e- mails. She is still defending Planned Parenthood, and she is still her party's frontrunner. 2016 is going to be a fight between conservatism, and a Democrat party that is undermining the very character of this nation. We need a nominee who is going to throw every punch, not pull punches, and someone who cannot stumble before he even gets into the ring.  I am not a member of the political class. I am a conservative; I can win this job, I can do this job, I need your help, I need your support. I will, with your help and support, lead the resurgence of this great nation.  Thank you.  HEMMER: Thank you.  MACCALLUM: Senator Graham.  GRAHAM: We need somebody ready to be commander-in-chief on day one, who understands there are no moderates in Iran, they've been killed a long time ago. That the Ayatollah is a radical jihadist who really means it when he chants, "Death to America, death to Israel." And this deal is giving him a pathway to a bomb, a missile to deliver it, and money to pay for it all.  We need a president who can solve our problems, bring us together. We're becoming Greece if we don't work together. At the end of the day, ladies and gentlemen, our best days are ahead of us only if we work together, and I intend to put this country on a path of success by working together and doing the hard things that should have been done a very long time ago.  HEMMER: And to Governor Pataki, closing statement now.  PATAKI: With all the candidates, why me?  My background is different. I look at Washington, and I hear the talk, and I see the promises and it seems nothing ever changes. Washington gets bigger, taxes get higher, and the American people feel more distance from our government. I have the opportunity not just to run, but to win in the deep blue state of New York three times. And not only did I win, but I then worked with a Democratic legislature to put in place the most sweeping conservative reforms of any state in America, taking us from the most dangerous state in America to the fourth safest; reducing our welfare rolls by over 1 million, and replacing over 700,000 private sector jobs.  I can govern by bringing people together. And also, I've been tested in a way no one else has. I was governor on September 11th, and I'm proud of my leadership in bringing New York through that time. And when I left, we were stronger, we were safer, and we were more united than at any time in my lifetime.  We need to bring people together in Washington. The talk has got to stop, the action has got to begin. People can promise you something, I delivered in the blue state of New York. I will deliver for the American people if I have the privilege of leading this country.  HEMMER: Thank you, Governor.  MACCALLUM: Governor Gilmore.  GILMORE: Well, I was a conservative governor of Virginia, I governed that way, and that's my track record. But the key thing that we're seeing now is serious challenges to this country that must change, the direction of this nation must change. And that's why I've offered a specific program to the people of America tonight to address the fundamental problem of getting our country growing again, getting our economy growing, wages up, opportunities for people.  And second, the international crisis we are facing is most dreadful and most dangerous. I have the experience as a prosecutor, attorney general, governor, United States Army intelligence veteran, governor during the 9/11 attack, chairman of the Terrorism Commission for this country. It's time for real substance and real experience.  And that's what I'll offer to the people of the United States in this candidacy for the presidency.  MACCALLUM: Thank you, Governor.  HEMMER: That concludes the first debate of the 2016 Republican primary. We would like to thank all seven of you for being here today.`

	analysis := model.SentimentAnalysis(transcript, English)

	if analysis.Score > uint8(5) {
		t.Errorf("Analysis of transcript should be less than 0.5\n\treturned %v\n", analysis.Score)
	} else {
		t.Logf("Analysis of transcript was valid\n\treturned %v\n", analysis.Score)
	}

	if analysis.Words == nil {
		t.Errorf("Analysis of transcript returned nil words array!\n")
	} else {
		t.Logf("Analysis of transcript retuned valid word array\n\t returned %+v\n", analysis.Words)
	}
}
