import flask
import datetime as dt
def hello(event, context):
    print(flask)
    print(dt.datetime.now())
    sum = 0
    for i in range(0, 1000000):
        sum += 1
    print(sum)
    print(dt.datetime.now())
    return "Hello world!"
