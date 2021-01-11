package main

import (
	"fmt"

	"errors"
	"net/smtp"
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

func inpAddressTo() int {
	var addrTo [40]byte

	fmt.Scanf(
		"%s\n",
		&addrTo,
	)
	fmt.Println(addrTo[2])
	return 0
}

func inpsubject() {
	fmt.Scanf(
		"%s\n",
		&subjectText,
	)
}

func commandSend() int {
	var komand string

	fmt.Scanf(
		"%s\n",
		&komand,
	)
	if komand == "Y" || komand == "y" {
		return 1
	}
	return 0
}

func sendPost() int {

	auth := smtp.PlainAuth("", addrFrom, passwMail, host)
	if err := smtp.SendMail(host+numport, auth, addrFrom, []string{addrTo}, []byte(message)); err != nil {
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
			err = inpAddressTo()
			if err == 1 {
				fmt.Println("Aдрес получателя некорректен")
			}
		}
		fmt.Print("Введите тему почты:		")
		inpsubject()
		fmt.Println("Введите текст письма")
		inpMessage()
		fmt.Print("Отправляю? (Y)   ")
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
