package main

import (
	"errors"
	"fmt"
	"net/smtp"
	"os"
)

const (
	host      = "mail.nic.ru"
	numport   = ":25"
	addrFrom  = "sendertab@catelecom.ru"
	passwMail = "jc39vI71"
	//sslpost	= "true"
	//authent = "plain"

)

var (
	message, addrTo, subjectText string
)

var (
	ErrInvalidAddr = errors.New("invalid mail address")
	ErrSentMail    = errors.New("error sent mail")
)

func inpMessage() {
	fmt.Scanf(
		"%s\n",
		&message,
	)
}

func inpAdreString() int {
	fmt.Scanf(
		"%s\n",
		&addrTo,
	)
	return 0
}

func inpAddrTo() int {
	len := 20
	data := make([]byte, len)
	n, err := os.Stdin.Read(data)
	if err != nil {
		fmt.Print("err_read= ")
	}
	if err == nil && n > 0 {
		fmt.Printf("Получено: %X 	%s\n", data, string(data))
		err := check(data)
		fmt.Println("err_check= ", err)
		if err == 0 {
			addrTo = string(data)
			fmt.Println("addrTo=", addrTo)
			return 0
		}
	} else {
		return 1
	}
	return 1
}

// ****************проверка корректности адреса*************
// *********** флаги адреса при анализе массива байт date[i]*********
// если пробел встретился - ошибка
// multihund = 0  счетчик собак собака должна встретиться только 1 раз не раньше 2-х символов
// multipoint	> "hund" +1 после собаки должен быть хоть 1 символ до встречи с "."
//				если точек после собаки больше 1 ошибка
// если эти флаги установлены то адрес корректен

func check(data []byte) int {

	hundz := 0
	multihund := 0
	multipoint := 0

	// анализ байт адреса
	for i := 1; i < len(data); i++ {
		if data[i] == ' ' {
			return 1 // встретился пробел - ошибка
		}
		if i > 1 && data[i] == '@' {
			hundz = i   // после хотя бы 2-х символов встретилась собака, запомнили место
			multihund++ // считаем собак
			if multihund > 1 {
				return 1 // встретилась вторая собака - ошибка
			}
		}
		if hundz > 1 && i > hundz+1 && data[i] == '.' {
			multipoint++ // сначала собака потом обязательно хотя бы 1 символ затем точка
			// (точки до собаки игнорируем)
			// считаем точки после собаки
			if multipoint > 1 { // проверим после собакии точек больше 2-х ?
				return 1 // после собаки встретилась вторая точка  - ошибка
			}
		}
	}
	if multipoint == 1 {
		return 0 // адрес корректен
	} else {
		return 1 // точек много адрес не корректен
	}

}

func inpsubject() {
	fmt.Scanf(
		"%s\n",
		&subjectText,
	)
}

func commandSend() int {
	var komand string
	println("-------------------------------------------------")
	println("Адрес получателя:    ", addrTo)
	println("Адрес отправителя:   ", addrFrom)
	println("Тема: ", subjectText)
	println("-------------------------------------------------")
	print(message, "\n")
	println("-------------------------------------------------", "\n")
	fmt.Print("Отправляю? (Y)   ", "\n")
	fmt.Scanf(
		"%s\n",
		&komand,
	)
	if komand == "Y" || komand == "y" || komand == "Н" || komand == "н" {
		return 1 // отправляем
	} else {
		return 0
	}
}

func sendPost() int {

	auth := smtp.PlainAuth("", addrFrom, passwMail, host)
	if err := smtp.SendMail(host+numport, auth, addrFrom, []string{addrTo}, []byte(message)); err != nil {
		println("***error_send****")
		return 1
		//os.Exit(1)
	}
	return 0
}

func main() {
	var komand string
	var (
		err, comsend int
	)

	message = ""
	addrTo = ""
	subjectText = ""
	komand = "Y"
	for komand == "Y" || komand == "y" || komand == "Н" || komand == "н" {
		fmt.Println("------------------------------------")
		fmt.Println("|         Отправка эл почты        |")
		fmt.Println("|  Отправлять, не переотправлять!  |")
		fmt.Println("|                                  |")
		fmt.Println("|   (c) jiliaevyp@gmail.com        |")
		fmt.Println("------------------------------------")
		err = 1
		for err == 1 {
			fmt.Print("Введите адрес получателя:	")
			//err = inpAddressTo()
			//err = inpAddrTo()
			err = inpAdreString()
			if err == 1 {
				fmt.Println("Aдрес получателя некорректен")
			}
		}
		fmt.Print("Введите тему почты:		")
		inpsubject()
		fmt.Println("Введите текст письма")
		inpMessage()
		comsend = commandSend()
		if comsend == 1 {
			err = sendPost()
			if err == 1 {
				fmt.Print("*** Ошибка при отправлении почты ***", "\n", "\n")
			} else {
				fmt.Print("---------------------------", "\n")
				fmt.Print("\n", "Письмо отправлено", "\n", "\n")
				fmt.Print("---------------------------", "\n")
			}
		} else {
			fmt.Print("\n", "Отправка письма отменена", "\n", "\n")
		}
		fmt.Print("Продолжить? (Y)   ")
		fmt.Println("Закончить?  (Enter)")
		komand = ""
		fmt.Scanf(
			"%s\n",
			&komand,
		)
	}
	fmt.Println("Рад был с Вами пработать!")
	fmt.Print("Обращайтесь в любое время без колебаний!", "\n", "\n")
}
