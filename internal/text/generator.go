package text

import (
	"math/rand"
	"strings"
	"time"
)

type Generator struct {
	words []string
}

func NewGenerator() *Generator {
	return &Generator{
		words: []string{
			"the", "be", "to", "of", "and", "a", "in", "that", "have", "I",
			"it", "for", "not", "on", "with", "he", "as", "you", "do", "at",
			"this", "but", "his", "by", "from", "they", "we", "say", "her", "she",
			"or", "an", "will", "my", "one", "all", "would", "there", "their", "what",
			"so", "up", "out", "if", "about", "who", "get", "which", "go", "me",
			"when", "make", "can", "like", "time", "no", "just", "him", "know", "take",
			"people", "into", "year", "your", "good", "some", "could", "them", "see", "other",
			"than", "then", "now", "look", "only", "come", "its", "over", "think", "also",
			"back", "after", "use", "two", "how", "our", "work", "first", "well", "way",
			"even", "new", "want", "because", "any", "these", "give", "day", "most", "us",
			"is", "am", "are", "was", "were", "been", "being", "had", "has", "did",
			"doing", "does", "done", "said", "says", "saying", "went", "gone", "made", "making",
			"came", "come", "coming", "saw", "seen", "seeing", "took", "taken", "taking", "got",
			"gotten", "getting", "put", "putting", "found", "finding", "gave", "given", "giving", "told",
			"telling", "felt", "feeling", "let", "letting", "left", "leaving", "kept", "keeping", "began",
			"begun", "beginning", "brought", "bringing", "wrote", "written", "writing", "met", "meeting", "ran",
			"run", "running", "set", "setting", "stood", "standing", "won", "winning", "thought", "thinking",
			"heard", "hearing", "showed", "shown", "showing", "spoke", "spoken", "speaking", "sat", "sitting",
			"lost", "losing", "paid", "paying", "built", "building", "bought", "buying", "sent", "sending",
			"fell", "fallen", "falling", "led", "leading", "spent", "spending", "taught", "teaching", "grew",
			"grown", "growing", "drew", "drawn", "drawing", "drove", "driven", "driving", "held", "holding",
			"meant", "meaning", "kept", "keeping", "let", "letting", "begun", "beginning", "seems", "seemed",
			"seeming", "shown", "showing", "lay", "laid", "lying", "rose", "risen", "rising", "broke",
			"broken", "breaking", "spoke", "spoken", "speaking", "chose", "chosen", "choosing", "wore", "worn",
			"wearing", "ate", "eaten", "eating", "hit", "hitting", "bit", "bitten", "biting", "blew",
			"blown", "blowing", "caught", "catching", "dealt", "dealing", "dug", "digging", "flew", "flown",
			"flying", "forgot", "forgotten", "forgetting", "froze", "frozen", "freezing", "hid", "hidden", "hiding",
			"hurt", "hurting", "rode", "ridden", "riding", "rang", "rung", "ringing", "sang", "sung",
			"singing", "sank", "sunk", "sinking", "slept", "sleeping", "slid", "sliding", "threw", "thrown",
			"throwing", "understood", "understanding", "woke", "woken", "waking", "wore", "worn", "wearing", "tore",
			"torn", "tearing", "beat", "beaten", "beating", "bound", "binding", "cast", "casting", "crept",
			"creeping", "cut", "cutting", "dealt", "dealing", "dug", "digging", "fed", "feeding", "fled",
			"fleeing", "flung", "flinging", "forbade", "forbidden", "forbidding", "forgave", "forgiven", "forgiving", "froze",
			"frozen", "freezing", "got", "gotten", "getting", "gave", "given", "giving", "ground", "grinding",
			"hung", "hanging", "knelt", "kneeling", "knit", "knitting", "laid", "laying", "led", "leading",
			"leapt", "leaping", "learnt", "learning", "left", "leaving", "lent", "lending", "let", "letting",
			"lost", "losing", "made", "making", "meant", "meaning", "met", "meeting", "mown", "mowing",
			"overcome", "overcoming", "overdone", "overdoing", "overtaken", "overtaking", "overthrown", "overthrowing", "paid", "paying",
			"pled", "pleading", "proven", "proving", "put", "putting", "quit", "quitting", "read", "reading",
			"rid", "ridding", "ridden", "riding", "rung", "ringing", "risen", "rising", "run", "running",
			"sawn", "sawing", "said", "saying", "seen", "seeing", "sought", "seeking", "sold", "selling",
			"sent", "sending", "set", "setting", "sewn", "sewing", "shaken", "shaking", "shaven", "shaving",
			"shorn", "shearing", "shed", "shedding", "shone", "shining", "shod", "shoeing", "shot", "shooting",
			"shown", "showing", "shrunk", "shrinking", "shut", "shutting", "sung", "singing", "sunk", "sinking",
			"sat", "sitting", "slain", "slaying", "slept", "sleeping", "slid", "sliding", "slung", "slinging",
			"slit", "slitting", "smitten", "smiting", "sown", "sowing", "spoken", "speaking", "sped", "speeding",
			"spent", "spending", "spilt", "spilling", "spun", "spinning", "spit", "spitting", "split", "splitting",
			"spread", "spreading", "sprung", "springing", "stood", "standing", "stolen", "stealing", "stuck", "sticking",
			"stung", "stinging", "stunk", "stinking", "stridden", "striding", "struck", "striking", "striven", "striving",
			"sworn", "swearing", "swept", "sweeping", "swollen", "swelling", "swum", "swimming", "swung", "swinging",
			"taken", "taking", "taught", "teaching", "torn", "tearing", "told", "telling", "thought", "thinking",
			"thrived", "thriving", "thrown", "throwing", "thrust", "thrusting", "trodden", "treading", "understood", "understanding",
			"upheld", "upholding", "upset", "upsetting", "woken", "waking", "worn", "wearing", "woven", "weaving",
			"wed", "wedding", "wept", "weeping", "wound", "winding", "won", "winning", "withheld", "withholding",
			"withstood", "withstanding", "wrung", "wringing", "written", "writing",
		},
	}
}

func (g *Generator) GenerateText() string {
	rand.NewSource(time.Now().UnixNano())
	var result []string
	wordCount := 0

	for wordCount < 100 {
		word := g.words[rand.Intn(len(g.words))]
		result = append(result, word)
		wordCount += len(word) + 1
	}

	// Trim the last word if it exceeds 100 characters
	fullText := strings.Join(result, " ")
	if len(fullText) > 100 {
		fullText = fullText[:100]
	}

	fullText = strings.ToUpper(fullText[:1]) + fullText[1:]
	if fullText[len(fullText)-1] != '.' {
		fullText += ""
	}

	return fullText
}
