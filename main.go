package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func randint(min, max int) int {
	return rand.Intn(max-min) + min
}

func attack(CharName, charClass string) string {
	if charClass == "warrior" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", CharName, 5+randint(3, 5))
	}

	if charClass == "mage" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", CharName, 5+randint(5, 10))
	}

	if charClass == "healer" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", CharName, 5+randint(-3, -1))
	}
	return "неизвестный класс персонажа"
}

// обратите внимание на "if else" и на "else"
func defence(CharName, charClass string) string {
	if charClass == "warrior" {
		return fmt.Sprintf("%s блокировал %d урона.", CharName, 10+randint(5, 10))
	}

	if charClass == "mage" {
		return fmt.Sprintf("%s блокировал %d урона.", CharName, 10+randint(-2, 2))
	}

	if charClass == "healer" {
		return fmt.Sprintf("%s блокировал %d урона.", CharName, 10+randint(2, 5))
	}
	return "неизвестный класс персонажа"
}

// обратите внимание на "if else" и на "else"
func special(CharName, charClass string) string {
	if charClass == "warrior" {
		return fmt.Sprintf("%s применил специальное умение `Выносливость %d`", CharName, 80+25)
	}

	if charClass == "mage" {
		return fmt.Sprintf("%s применил специальное умение `Атака %d`", CharName, 5+40)
	}

	if charClass == "healer" {
		return fmt.Sprintf("%s применил специальное умение `Защита %d`", CharName, 10+30)
	}
	return "неизвестный класс персонажа"
}

// здесь обратите внимание на имена параметров
func start_training(CharName, charClass string) string {
	if charClass == "warrior" {
		fmt.Printf("%s, ты Воитель - отличный боец ближнего боя.\n", CharName)
	}

	if charClass == "mage" {
		fmt.Printf("%s, ты Маг - превосходный укротитель стихий.\n", CharName)
	}

	if charClass == "healer" {
		fmt.Printf("%s, ты Лекарь - чародей, способный исцелять раны.\n", CharName)
	}

	fmt.Println("Потренируйся управлять своими навыками.")
	fmt.Println("Введи одну из команд: attack — чтобы атаковать противника,")
	fmt.Println("defence — чтобы блокировать атаку противника,")
	fmt.Println("special — чтобы использовать свою суперсилу.")
	fmt.Println("Если не хочешь тренироваться, введи команду skip.")

	var cmd string
	for cmd != "skip" {
		fmt.Print("Введи команду: ")
		fmt.Scanf("%s\n", &cmd)
		switch {
		case cmd == "attack":
			fmt.Println(attack(CharName, charClass))
		case cmd == "defence":
			fmt.Println(defence(CharName, charClass))
		case cmd == "special":
			fmt.Println(special(CharName, charClass))
		default:
			fmt.Println("неизвестная команда")
		}
	}
	return "тренировка окончена"
}

// обратите внимание на имя функции и имена переменных
func choiceCharClass() string {
	var approveChoice string
	var charClass string

	for approveChoice != "y" {
		fmt.Print("Введи название персонажа, за которого хочешь играть: Воитель — warrior, Маг — mage, Лекарь — healer: ")
		fmt.Scanf("%s\n", &charClass)
		if charClass == "warrior" {
			fmt.Println("Воитель — дерзкий воин ближнего боя. Сильный, выносливый и отважный.")
		} else if charClass == "mage" {
			fmt.Println("Маг — находчивый воин дальнего боя. Обладает высоким интеллектом.")
		} else if charClass == "healer" {
			fmt.Println("Лекарь — могущественный заклинатель. Черпает силы из природы, веры и духов.")
		}
		fmt.Print("Нажми (Y), чтобы подтвердить выбор, или любую другую кнопку, чтобы выбрать другого персонажа: ")
		fmt.Scanf("%s\n", &approveChoice)
		approveChoice = strings.ToLower(approveChoice)
	}
	return charClass
}

// обратите внимание на имена переменных
func main() {
	fmt.Println("Приветствую тебя, искатель приключений!")
	fmt.Println("Прежде чем начать игру...")

	var name string
	fmt.Print("...назови себя: ")
	fmt.Scanf("%s\n", &name)

	fmt.Printf("Здравствуй, %s\n", name)
	fmt.Println("Сейчас твоя выносливость — 80, атака — 5 и защита — 10.")
	fmt.Println("Ты можешь выбрать один из трёх путей силы:")
	fmt.Println("Воитель, Маг, Лекарь")

	class := choiceCharClass()

	fmt.Println(start_training(name, class))
}
