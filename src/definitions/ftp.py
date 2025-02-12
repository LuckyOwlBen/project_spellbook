class fileTransferProtocol:
    def __init__(self):
        self.ftp_root = "ftp"
        self.ftp_anonymous_login = "-a"
    @property
    def ftp_root(self):
        return self._ftp_root

    @property
    def ftp_anonymous_login(self):
        return self._ftp_anonymous_login