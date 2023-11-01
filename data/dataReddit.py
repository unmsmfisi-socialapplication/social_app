import praw

reddit = praw.Reddit(
    client_id="yi_hZzmBPMkH5QFXz5WfXQ",
    client_secret="pMo2c4ydx_pLA14SZdTTJz1Nt2K0AQ",
    user_agent="app-dataRedi by u/vladnn177",
    username="vladnn177 ",
    password="B/D#Bk?Cf/6$C@U"
)

subreddit = reddit.subreddit("python")

top_posts = subreddit.top(limit=10)
new_posts = subreddit.new(limit=10)

for post in top_posts:
    print("Title - ", post.title)
    print("ID - ", post.id)
    print("Author - ", post.author)
    print("URL - ", post.url)
    print("Score - ", post.score)
    print("Comment count - ", post.num_comments)
    print("Created - ", post.created_utc)
    print("\n")

    post = reddit.submission(id=post.id)

    comments = post.comments

    for comment in comments[:4]:
        print("     Printing comment...")
        print("     Comment body - ", comment.body)
        print("\n")

