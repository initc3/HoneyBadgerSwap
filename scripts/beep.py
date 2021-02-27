from time import sleep


def beep(*, t=10):
    while True:
        print("beep")
        sleep(t)


if __name__ == "__main__":
    beep()
