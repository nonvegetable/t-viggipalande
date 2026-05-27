package main

func projectsSection() portfolioSection {
	return portfolioSection{
		key:         "projects",
		title:       "Projects",
		summary:     "A set of product ideas turned into deployed apps and experiments.",
		description: "Projects with a bias toward usefulness, fast feedback, and clear interfaces.",
		items: []portfolioItem{
			{title: "Formula Genie • SaaS Platform", meta: "AI, React.js, backend APIs, Razorpay", description: "Problem: formula workflows were scattered. Built: a full-stack SaaS for generation, auth, payments, and usage tracking.", linkText: "Open project", linkURL: "https://formulagenie.vercel.app/"},
			{title: "PokéGuess • Pokémon Wordle", meta: "React.js, Python, Flask", description: "Problem: make daily guessing feel fair, quick, and a little annoying. Built: a Wordle-style game with progressive hints and fuzzy matching.", linkText: "Open project", linkURL: "https://pokemon-wordle-murex.vercel.app/"},
			{title: "FONO • Phone Memory Game", meta: "JavaScript, HTML, CSS", description: "Problem: make short-term memory training less boring. Built: a browser game for recalling phone numbers digit by digit.", linkText: "Open project", linkURL: "https://phone-number-game.onrender.com/"},
			{title: "Assero Blockchain Platform", meta: "Java, React.js", description: "Problem: manage assets with clearer ownership records. Built: Ethereum smart contracts for asset workflows."},
		},
	}
}
