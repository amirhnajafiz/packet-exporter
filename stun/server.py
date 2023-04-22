import socket


class Server:
    def __int__(self):
        self.peers = []

    def _join(self, ip):
        self.peers.append(ip)

    def _leave(self, ip):
        self.peers.remove(ip)

    def listen(self, port):
        pass
