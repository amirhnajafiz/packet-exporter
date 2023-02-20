import socket
import sys



# let's open a connection to google with socket
PORT = 80

try:
    # first we need to find google ip
    host = socket.gethostbyname('www.google.com')
except socket.gaierror as err:
    print("cannot find google ip %s" % (err))
    sys.exit()
