package main

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click

// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

/*func sayHello(w http.ResponseWriter, r *http.Request) {
	// 向响应写入器输出固定字符串内容
	_, _ = fmt.Fprintln(w, "<h1>Hello World!</h1>")
}*/

/*
main 程序入口函数
初始化HTTP服务并监听8080端口：
1. 注册路由/hello到sayHello处理器
2. 启动HTTP服务监听
*/

/*func main() {
	// 注册路由与处理函数的映射关系
	http.HandleFunc("/hello", sayHello)
	// 启动HTTP服务器并监听端口
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
		return
	}
}*/

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	"time"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world!",
	})
}

// 定义中间件
// StatCost 是一个统计耗时请求耗时的中间件
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("name", "小王子") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		// 调用该请求的剩余处理程序
		c.Next()
		// 不调用该请求的剩余处理程序
		// c.Abort()
		// 计算耗时
		cost := time.Since(start)
		log.Println(cost)
	}
}

func indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "我是一个中间件",
	})
}

func m1Handler(c *gin.Context) {
	c.Set("name", "lcz")
	fmt.Println("m1 in")
	start := time.Now()
	//c.Set("name", "小王子") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
	// 调用该请求的剩余处理程序
	c.Next()
	// 不调用该请求的剩余处理程序
	// c.Abort()
	// 计算耗时
	cost := time.Since(start)
	log.Println(cost)
	fmt.Println("m1 out")
}

func m2Handler(c *gin.Context) {
	name, ok := c.Get("name")
	if !ok {
		fmt.Println("未取到值")
	}
	fmt.Println(name)

	fmt.Println("m2 in")
	start := time.Now()
	//c.Set("name", "小王子") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
	// 调用该请求的剩余处理程序
	//c.Next()
	// 不调用该请求的剩余处理程序 后续中间件不会调用
	//c.Abort()
	// 如果要使当前中间件后续内容也不执行，则调用return
	// 计算耗时
	cost := time.Since(start)
	log.Println(cost)
	fmt.Println("m2 out")
}

// authMiddleware 鉴权中间件
func authMiddleware(doCheck bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if doCheck {
			fmt.Println("调用authmiddleware")
		} else {
			fmt.Println("调用authmiddleware2")
		}
	}
}

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// 加载静态文件  xxx代指 ./static 加载静态文件需要在解析模板之前进行
	r.Static("/xxx", "./statics")
	// 设置模板函数 safe
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})

	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	/*	r.GET("/hello", sayHello) // 启动HTTP服务，默认在0.0.0.0:8080启动服务

		// RESTful 软件架构风格
		r.GET("/book", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "GET",
			})
		})

		r.POST("/book", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "POST",
			})
		})

		r.PUT("/book", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "PUT",
			})
		})

		r.DELETE("/book", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "DELETE",
			})
		})*/

	// 模板解析
	r.LoadHTMLFiles("./templates/index.html", "./templates/login.html", "./templates/upload.html")
	//r.LoadHTMLFiles("./login.html")
	//r.LoadHTMLGlob("templates/**/*")

	// HTTP请求
	/*	r.GET("/posts/index", func(c *gin.Context) {
			// 状态码 200， 要渲染的模板名称，要传递给模板的数据
			c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
				"title": "lzc",
			})
		})

		r.GET("/users/index", func(c *gin.Context) {
			// 状态码 200， 要渲染的模板名称(如果模板没有取名，则默认为文件名)，要传递给模板的数据
			c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
				"title": "<a href = 'https://www.baidu.com/'>百度一下</a>",
			})
		})*/

	r.GET("/json", func(c *gin.Context) {
		// 方法1：使用map
		/*		data := map[string]interface{}{
				"name": "lzc",
				"age":  18,
				"sex":  "男",
			}*/
		// gin.H = map[string]interface{}
		data := gin.H{
			"name": "lzc",
			"age":  18,
			"sex":  "男",
		}
		c.JSON(http.StatusOK, data)
	})
	// 使用结构体存储存储json 字段名需要大写，否则无法传入前端
	type msg struct {
		Name    string `json:"name"` // 反射到前端为小写
		Message string
		Age     int
	}
	r.GET("/json2", func(c *gin.Context) {
		data := msg{
			Name:    "lzc",
			Message: "Hello World!",
			Age:     18,
		}
		c.JSON(http.StatusOK, data)
	})

	r.GET("/web", func(c *gin.Context) {
		// 获取浏览器向服务器发送请求的 query 参数
		//name := c.Query("query") // 通过Query方法获取请求参数
		name := c.DefaultQuery("query", "defaultValue") // 取不到就用指定的默认值
		/*		name, ok := c.GetQuery("query")
				if !ok { // 取不到，第二个参数则返回 false
					name = "somebody"
				}*/
		age := c.Query("age")
		// 将获取参数渲染到前端页面
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	// 浏览器先向服务端请求登录页面
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	// 浏览器再向服务端提交表单数据
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		/*		username := c.DefaultPostForm("username", "somebody")
				password := c.DefaultPostForm("password", "***")*/
		/*		username, ok := c.GetPostForm("username")
				if !ok {
					username = "somebody"
				}*/
		// 后端返回一个主页
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name":     username,
			"Password": password,
		})
	})

	// 获取URI路径参数
	r.GET("/user/:name/:age", func(c *gin.Context) {
		// 获取路径参数
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.GET("/blog/:year/:month", func(c *gin.Context) {
		year := c.Param("year")
		month := c.Param("month")
		c.JSON(http.StatusOK, gin.H{
			"year":  year,
			"month": month,
		})
	})
	// 结构体
	type UserInfo struct {
		Username string `form:"username" json:"ume"` // 需要与请求参数绑定
		Password string `form:"password" json:"pwd"`
	}

	r.GET("/user", func(c *gin.Context) {
		/*		username := c.Query("username")
				password := c.Query("password")
				u := UserInfo{
					username: username,
					password: password,
				}
				fmt.Printf("%#v\n", u)
				c.JSON(http.StatusOK, gin.H{
					"message": "ok",
				})*/
		var u UserInfo          // 声明一个结构体变量
		err := c.ShouldBind(&u) // 结构体与请求参数进行绑定
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}
	})
	// 利用 form表单与请求参数绑定
	r.POST("/form", func(c *gin.Context) {
		var u UserInfo          // 声明一个结构体变量
		err := c.ShouldBind(&u) // 结构体与请求参数进行绑定
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}
	})
	//  json与请求参数绑定
	r.POST("/json", func(c *gin.Context) {
		var u UserInfo          // 声明一个结构体变量
		err := c.ShouldBind(&u) // 结构体与请求参数进行绑定
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}
	})
	// 上传文件 demo
	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", nil)
	})
	// 单文件上传
	/*	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
		// 拼接文件的存储路径
		dst := fmt.Sprintf("./uploaded/%s", file.Filename)
		// 存储文件
		c.SaveUploadedFile(file, dst)
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%s uploaded!", file.Filename),
		})
	})*/
	// 多文件上传
	r.POST("/upload", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["f1"]

		fmt.Println(files)

		for index, file := range files {
			dst := fmt.Sprintf("./uploaded/%d_%s", index, file.Filename)
			c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d files uploaded!", len(files)),
		})
	})

	// 重定向
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})

	r.GET("a", func(c *gin.Context) {
		c.Request.URL.Path = "/b"
		r.HandleContext(c)
	})
	r.GET("b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "我是b",
		})
	})

	// 请求方法大集合
	r.Any("/all", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet: // case "GET"
			c.JSON(http.StatusOK, gin.H{
				"message": "get",
			})
		case http.MethodPost: // case "POST"
			c.JSON(http.StatusOK, gin.H{
				"message": "post",
			})
		case http.MethodPut: // case "PUT"
			c.JSON(http.StatusOK, gin.H{
				"message": "put",
			})
		case http.MethodDelete: // case "DELETE"
			c.JSON(http.StatusOK, gin.H{
				"message": "delete",
			})
		}
	})

	// 404路由 规定访问不到资源路由
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "wrong!wrong!",
		})
	})

	// 路由组
	viderGroup := r.Group("/a")
	{
		viderGroup.GET("/aa", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "我是a/aa",
			})
		})
		viderGroup.GET("/ab", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "我是a/ab",
			})
		})
		// 嵌套路由组
		xx := viderGroup.Group("xx")
		xx.GET("oo", func(c *gin.Context) {
			// localhost:8080/a/xx/oo
		})
	}

	//r := gin.Default() 默认使用了 Logger 和 Recovery中间件
	// 不想使用默认的中间件，可以使用 r := gin.New() 新建一个没有任何默认中间件的路由
	// 挡在中间件或 handler中启动新的 goroutine时，不能使用原始的上下文 c * gin.Context 必须使用其只读副本 c.Copy()
	// 中间件
	r.Use(m1Handler, m2Handler, authMiddleware(true)) // 全局注册中间件
	// GET(relativePath string, handlers ...HandlerFunc)
	r.GET("/middleware", indexHandler)

	r.GET("/shop")
	r.GET("boat") // 自带中间件 m1Handler

	// 路由组注册中间件 注意 会遵循 r注册的全局中间件
	xxGroup := r.Group("/xx", authMiddleware(true))
	{
		xxGroup.GET("/oo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "111",
			})
			// localhost:8080/xx/oo
		})
	}

	xx2Group := r.Group("/xx2")
	xx2Group.Use(authMiddleware(true))
	{
		xx2Group.GET("/oo2", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "222",
			})
			// localhost:8080/xx2/oo
		})
	}

	r.Run()

}
