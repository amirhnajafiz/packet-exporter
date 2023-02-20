import socket
import sys



try:
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    print ("socket successfully created")
except socket.error as err:
    print ("socket creation failed with error %s" %(err))
    sys.exit()

# let's open a connection to google with socket
PORT = 80

try:
    # first we need to find google ip
    host = socket.gethostbyname('www.google.com')
except socket.gaierror as err:
    print("cannot find google ip %s" % (err))
    sys.exit()

s.connect((host, PORT))

print ("the socket has successfully connected to google")