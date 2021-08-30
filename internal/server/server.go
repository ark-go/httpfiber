package server

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func init() {}

func StartServer() {
	configServer := fiber.Config{
		ServerHeader: "My Server", // "", Включает HTTP-заголовок сервера с заданным значением
		//Prefork: ,		//!несколько процессов Go будут прослушивать один и тот же порт, Изучать!!
		//StrictRouting: ,	// false, true - рассматривает /foo и /foo/ как разные маршруты
		//CaseSensitive: ,	// false, true - рассматривает /Foo и /foo как разные маршруты
		//Immutable: ,		 // false,  Если этот параметр включен, все значения, возвращаемые контекстными методами, неизменяемы. По умолчанию они действительны до тех пор, пока вы не вернетесь из обработчика
		//UnescapePath: ,// false, Преобразует все закодированные символы в маршруте перед установкой пути для контекста, чтобы маршрутизация также могла работать со специальными символами, закодированными в URL
		//ETag: ,// false, Включите создание заголовков ETag , поскольку как слабые, так и сильные теги etags создаются с использованием одного и того же метода хеширования ( CRC-32 ).
		//BodyLimit: , //  по умолчанию: 4 * 1024 * 1024. максимально допустимый размер тела запроса
		//Concurrency: , // по умолчанию: 256 * 1024. Максимальное количество одновременных подключений
		//Views: , // это интерфейс , чтобы обеспечить собственный механизм шаблонов
		//ViewsLayout: ,
		//ReadTimeout: , // nil, Время, отведенное на чтение всего запроса, включая тело. Установите nilна неограниченное время ожидания
		//WriteTimeout: , // nil, Максимальная продолжительность до тайм-аута записи ответа. Установите nilна неограниченное время ожидания
		//IdleTimeout: , // nil, Максимальное время ожидания следующего запроса при включении проверки активности
		//ReadBufferSize: , // 4096, Размер буфера для каждого соединения для чтения запросов
		//WriteBufferSize: , // 4096, Размер буфера на соединение для записи ответов
		//CompressedFileSuffix: , // ".fiber.gz", Добавляет суффикс к исходному имени файла и пытается сохранить полученный сжатый файл под новым именем
		//ProxyHeader: , // "", Это позволит ctx.IPвернуть значение заданного ключа заголовка. По умолчанию ctx.IPвернет удаленный IP-адрес из TCP-соединения
		//GETOnly: , // false, Позволяет отклонять все запросы, не относящиеся к GET
		//ErrorHandler: , // DefaultErrorHandler, ErrorHandler выполняется, когда возвращается ошибка fiber.Handler
		//DisableKeepalive: , // false, Отключите поддерживающие соединения. Сервер закроет входящие соединения после отправки первого ответа клиенту
		//DisableDefaultDate: , // false, Если установлено trueзначение, заголовок даты по умолчанию будет исключен из ответа
		//DisableDefaultContentType: , false, Если задано значение true, Content-Typeзаголовок по умолчанию будет исключен из ответа
		//DisableHeaderNormalizing: , false, по умолчанию Например, заголовок cOnteNT-tYPEбудет преобразован в более читаемый Content-Type.
		DisableStartupMessage: true, // false, Если установлено значение true, он не будет распечатывать отладочную информацию и сообщение запуска, например
		//AppName: ,
		//StreamRequestBody: ,
		//DisablePreParseMultipartForm: ,
		//ReduceMemoryUsage: ,
		//JSONEncoder: ,
		//JSONDecoder: ,
		//Network: ,
		//EnableTrustedProxyCheck: ,
		//TrustedProxies: ,

	}
	app := fiber.New(configServer)
	routerMain(app)

	// Print the router stack in JSON format
	data, _ := json.MarshalIndent(app.Stack(), "", "  ")
	fmt.Println(string(data))

	// Установите опцию fasthttp 'MaxConnsPerIP' to '1'  https://pkg.go.dev/github.com/valyala/fasthttp#Server
	//app.Server().MaxConnsPerIP = 1
	// Прочитаем ServerHeader из конфигурации сервера
	fmt.Println("Конфиг:", app.Config().ServerHeader)
	// Создайте новую конфигурацию для статического метода
	config := fiber.Static{
		Compress:      true,
		ByteRange:     true,  // разрешает запросы диапазона байтов
		Browse:        false, // просмотр каталога
		Index:         "index.html",
		CacheDuration: 60 * time.Second, // Срок действия неактивных обработчиков файлов.
		MaxAge:        3600,             // sec, Значение Cache-Controlзаголовка HTTP, установленное в ответе файла.
		Next:          nil,
	}

	// Serve files from './public' directory with config
	app.Static("/", "./public", config)
	app.Listen(":3000")
}

// https://dev.to/koddr/go-fiber-by-examples-delving-into-built-in-functions-1p3k#listener
