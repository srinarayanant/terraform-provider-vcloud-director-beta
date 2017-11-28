import logging


class VCDVappCreationError(Exception):
    def __init__(self, msg):
        Exception.__init__(self, msg)


class VDCNotFoundError(Exception):
    def __init__(self, msg):
        Exception.__init__(self, " VDCNotFoundError [" + msg + "]")


class ItemFoundError(Exception):
    def __init__(self, msg):
        Exception.__init__(self, " ItemFoundError [" + msg + "]")


class VCDDiskCreationError(Exception):
    def __init__(self, msg):
        Exception.__init__(self, msg)


class VCDDiskDeletionError(Exception):
    def __init__(self, msg):
        Exception.__init__(self, msg)


if __name__ == '__main__':

    logging.basicConfig(level=logging.DEBUG)

    raise VCDVappCreationError("ee")
