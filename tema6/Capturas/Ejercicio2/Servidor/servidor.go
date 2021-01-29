package main

import "github.com/kataras/iris/v12"

func newApp() *iris.Application {
    	app := iris.New()
	
	usuario := app.Party("/usuario")
	{
		//usuario.Post("", nuevoUsuario)
		usuario.Get("", buscarUsuario)
		//usuario.Put("", modificarUsuario)
		//usuario.Delete("", borrarUsuario)
	}
	
   	return app
}

func buscarUsuario(ctx iris.Context){
	ctx.JSON(iris.Map{
		"username": "PabloA",
		"nombre": "Pablo",
		"apellido": "Alfaro",
		"correo": "pabloalfaro@correo.ugr.es",
		"ciudad": "Pamplona",
	})
	return
}

func main() {
	app := newApp()
	app.Listen(":3000")
}
