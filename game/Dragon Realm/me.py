import random
import time


def displayIntro():
    # 在你面前是一个满地是龙的土地
    print('You are in a land full of dragons. In front of you,')
    # 你看到两个洞穴。在一个山洞里，龙很友善将与你分享宝藏
    print('You see two caves. In one cave,the dragon is friendly')
    print('and will share his treasure with you. The')
    # 另外一只饥饿的龙，会吃掉你
    print('is greedy and hungry, and will eat you on sight.')
    print()


def chooseCave():
    cave = ''
    while cave != '1' and cave != '2':
        # 你会进入这个洞穴吗(1或2)
        print('While cave will you go into? (1 or 2)')
        cave = input()
    return cave


def checkCave(chosenCave):
    # 你靠近山洞
    print('You approach the cave...')
    time.sleep(2)
    # 它是黑暗怪异的
    print('It is dark and spooky...')
    time.sleep(2)
    # 一条巨龙在你面前跳出来！他张开嘴巴
    print('A large dragon jumps out in front of you! He opens his jaws and ...')
    print()
    time.sleep(2)
    friendlyCave = random.randint(1, 2)
    if chosenCave == str(friendlyCave):
        # 给你他的宝藏
        print('Gives you his treasure')
    else:
        # 你被吃掉了
        print('Gobbles you down in one bite!')


playAgain = 'yes'
while playAgain == 'yes' or playAgain == 'y':
    displayIntro()
    caveNumber = chooseCave()
    checkCave(caveNumber)
    print('Do you want to play again? (yes or no)')
    playAgain = input()
