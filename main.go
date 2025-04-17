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
	"net/http"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world!",
	})
}

// 静态文件：

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
	r.LoadHTMLFiles("./templates/index.html", "./templates/login.html")
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
		Username string `form:"username"` // 需要与请求参数绑定
		Password string `form:"password"`
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
	// 利用 post请求
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

	r.Run()
}
