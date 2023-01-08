import os , requests
from re import search as find
from slack_bolt import App
from slack_bolt.adapter.socket_mode import SocketModeHandler

# Initializes your app with your bot token and socket mode handler
app = App(token=os.environ.get("SLACK_BOT_TOKEN"))

# Listens to incoming messages that contain "hello"
# To learn available listener arguments,
# visit https://slack.dev/bolt-python/api-docs/slack_bolt/kwargs_injection/args.html
@app.message(".")
def message_hello(message, say):
    # say() sends a message to the channel where the event was triggered
    say(f"Hey there <@{message['user']}>!" )
    namespace=['praveen','testing']
    if find("namespace|name.*space|pods",message['text']):
        result=find("namespace\s(.*)",message['text'])
        if result.group(1)  in namespace:
            say(f"Requested namespace:   { result.group(1) }")
            url='http://pod-monitor.' + result.group(1) + '/metrics'
            say(url)
            response = requests.get(url)
            say("Pod Status \n################################\n")
            for i in str(response.content).split("\\n"):
                say(f"{i}")
        else:
            say("Namespace not available")
    else:
        say("""To proceed with service please mention namespace in the message ,
        Example 
         namespace testing """)

# Start your app
if __name__ == "__main__":
    SocketModeHandler(app, os.environ["SLACK_APP_TOKEN"]).start()
