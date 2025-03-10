package db

var usernames = [100]string{
	"JohnSmith", "JaneDoe", "AliceJohnson", "BobBrown", "CharlieDavis",
	"DavidClark", "EveLewis", "FrankMartin", "GraceThompson", "HankWhite",
	"IvyHarris", "JackWalker", "KaraYoung", "LeoKing", "MonaHill",
	"NinaScott", "OscarGreen", "PaulAdams", "QuinnBaker", "RoseCarter",
	"SamTurner", "TinaMitchell", "UmaPerez", "VictorRoberts", "WendyCampbell",
	"XanderFlores", "YaraJenkins", "ZaneSimmons", "AaronBrooks", "BellaMorris",
	"CalebRogers", "DianaBell", "EthanReed", "FionaBailey", "GeorgeCooper",
	"HollyBryant", "IanRussell", "JasmineStewart", "KyleSanders", "LilyFisher",
	"MasonDaniels", "NoraHunter", "OliverPalmer", "PiperLawrence", "QuincyGarrett",
	"RubyAlexander", "SeanGraham", "TaraWallace", "UlyssesDaniels", "VioletPeterson",
	"WalterWarren", "XeniaWoods", "YusufBoyd", "ZoeyFoster", "AidanFreeman",
	"BrookeCole", "CarterHughes", "DaisyRichards", "ElijahLane", "FaithBishop",
	"GavinGray", "HarperGriffin", "IsaacPhillips", "JadeOwens", "KieranBurns",
	"LaylaHenderson", "MicahCrawford", "NoahEllis", "OpalOrtiz", "PaigeMyers",
	"QuintonMorris", "ReeseChapman", "SawyerFox", "TessaLane", "UlrichSpencer",
	"VanessaBates", "WesleyHolland", "XimenaBrady", "YvetteMaxwell", "ZachRogers",
	"AriaChambers", "BlakeGregory", "ColtonHarper", "DelilahWalsh", "ElliotBrewer",
	"FelicityVega", "GraysonHolland", "HaileyMonroe", "IslaHarmon", "JamesFletcher",
}

var tags = []string{
	"Technology",
	"Programming",
	"Lifestyle",
	"Travel",
	"Health & Fitness",
	"Finance",
	"Food",
	"Education",
	"Personal Development",
	"Photography",
}

var titles = []string{
	"10 Must-Have Tools for Efficient Programming",
	"How to Stay Healthy While Working a Desk Job",
	"The Ultimate Guide to Budget-Friendly Travel",
	"Mastering React: Tips and Tricks for Beginners",
	"Photography 101: Capturing Stunning Landscapes",
	"Top Financial Habits to Build Wealth Over Time",
	"5 Simple Recipes for Busy Weeknights",
	"The Power of Habit: Transform Your Daily Routine",
	"Exploring the Future of AI in Everyday Life",
	"How to Make Learning Fun: Tips for Educators and Parents",
	"The Beginner’s Guide to Cryptocurrency Investing",
	"How to Declutter Your Life in 7 Simple Steps",
	"10 Best Practices for Writing Clean Code",
	"How to Create a Morning Routine That Works",
	"Top 10 Travel Destinations for 2024",
	"Understanding Cloud Computing for Beginners",
	"Simple Ways to Improve Your Mental Health",
	"Building a Minimalist Wardrobe on a Budget",
	"How to Start a Successful Blog in 2024",
	"5 Quick Workouts You Can Do at Home",
	"The Science Behind Productivity and Focus",
	"Mastering Git for Team Collaboration",
	"How to Cook Restaurant-Quality Meals at Home",
	"The Art of Storytelling in Photography",
	"Financial Independence: What It Means and How to Achieve It",
	"How to Stay Motivated While Working From Home",
	"Top 10 Trends in Web Development",
	"A Beginner's Guide to Digital Marketing",
	"5 Tips for Writing Better Email Newsletters",
	"How to Organize Your Life Using Notion",
	"The Role of UX Design in Product Development",
	"7 Steps to Launching Your First Online Course",
	"How to Travel Solo and Love It",
	"Understanding Machine Learning: A Non-Technical Guide",
	"The Benefits of Journaling for Stress Relief",
	"How to Grow Your Instagram Following Organically",
	"The Basics of Starting Your Own Podcast",
	"How to Build Confidence and Self-Esteem",
	"5 Effective Study Techniques for Students",
	"How to Choose the Best Laptop for Programming",
	"The Ultimate Guide to Meal Prepping for Beginners",
	"How to Build a Personal Brand on LinkedIn",
	"The Importance of Cybersecurity in Today’s World",
	"How to Start Investing in Stocks as a Beginner",
	"The Best Tools for Remote Work Productivity",
	"How to Plan the Perfect Weekend Getaway",
	"Understanding SEO: A Guide for Beginners",
	"How to Stay Creative in a World Full of Distractions",
	"The Health Benefits of Yoga and Meditation",
	"10 Books to Read for Personal Growth in 2024",
	"How to Design a Website That Converts",
	"Tips for Managing Stress During the Holidays",
	"How to Balance Work and Family Life",
	"The History and Future of Artificial Intelligence",
	"How to Choose the Right Camera for Your Needs",
	"The Art of Minimalism: Living with Less",
	"How to Conduct Market Research for Your Business",
	"10 Easy DIY Projects to Try This Weekend",
	"How to Set SMART Goals for the New Year",
	"The Basics of Blockchain Technology",
	"How to Stay Fit Without Going to the Gym",
	"The Evolution of Social Media Over the Years",
	"How to Create a Winning Resume and Cover Letter",
	"The Ultimate Guide to Freelancing in 2024",
	"How to Create Engaging YouTube Videos",
	"The Benefits of Learning a New Language",
	"How to Build a Sustainable Wardrobe",
	"The Importance of Networking for Career Growth",
	"How to Overcome Procrastination and Get Things Done",
	"The Best Budget Apps for Financial Tracking",
	"How to Host a Virtual Event That Stands Out",
	"The Science of Happiness: What Really Makes Us Happy",
	"How to Prepare for a Job Interview Like a Pro",
	"The Top Skills You Need to Succeed in 2024",
	"How to Create a Social Media Content Calendar",
	"The Benefits of Volunteering for Personal Growth",
	"How to Manage Your Time More Effectively",
	"The Pros and Cons of Remote Learning",
	"How to Create a Portfolio That Gets You Hired",
	"The Basics of Data Science for Beginners",
	"How to Write Engaging Blog Posts That Rank",
	"The Benefits of Practicing Gratitude Daily",
	"How to Use Analytics to Grow Your Business",
	"The Best Apps for Learning a New Skill",
	"How to Cultivate a Growth Mindset",
	"The Importance of Diversity in the Workplace",
	"How to Create a Financial Plan for Your Future",
	"The Best Tools for Managing Your To-Do List",
	"How to Stay Safe While Traveling Abroad",
	"The Role of AI in Healthcare Innovation",
	"How to Start a Side Hustle with Minimal Investment",
	"The Benefits of Walking for Physical and Mental Health",
	"How to Make the Most of Your Commute Time",
	"The Future of Renewable Energy Solutions",
	"How to Stay Inspired as a Creative Professional",
	"The Impact of Technology on Education",
	"How to Develop Emotional Intelligence",
	"The Best Online Communities for Learning and Networking",
	"How to Host a Memorable Dinner Party",
	"The Psychology of Consumer Behavior",
	"How to Make Money Online in 2024",
	"The Role of Design Thinking in Problem-Solving",
}

var content = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur."

var commentList = []string{
	"Great article! Really helped me understand the basics of web development.",
	"I disagree with some points, but overall a good read. Looking forward to more content!",
	"This is exactly what I needed! The step-by-step guide made it so easy to follow.",
	"I love how detailed this post is. It's clear you put a lot of effort into it. Thanks!",
	"I have a question about one part. Could you clarify how to implement step 4?",
	"Awesome tips, I’m definitely going to try this out on my next project!",
	"I’ve been struggling with this for a while, and this post made it click. Thanks!",
	"Can you recommend any tools or resources to dive deeper into this topic?",
	"This post was really insightful, but I’d love to see more on this subject in the future.",
	"Great content! But I think adding examples would make it even more useful.",
}
