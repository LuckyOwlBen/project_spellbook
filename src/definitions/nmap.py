class Nmap:
    def __init__(self):
        self.nmap_root = "nmap"
        self.version_check = "-sV"
        self.script_scan = "-sC"

    @property
    def nmap_root(self):
        return self._nmap_root
    
    @property
    def version_check(self):
        return self._version_check
    
    @property
    def script_scan(self):
        return self._script_scan