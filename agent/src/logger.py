from colorama import Fore, init  # type:ignore

init(autoreset=True)

class Logging:

    def info(self, data: str):
        print(Fore.YELLOW + (str(data)))

    def success(self, data: str):
        print(Fore.GREEN + (str(data)))

    def error(self, data: str):
        print(Fore.RED + (str(data)))


log = Logging()