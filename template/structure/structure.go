package structure

import (
	"fmt"
	"gogenerate/generate/generatestructer/gotemplate"
	"gogenerate/template/structure/Init/app"
	"log"
	"os"
)

type GoTemplate struct {
	Path    string
	AppPath string
}

func (gt *GoTemplate) CreateStructure(path string) {
	gt.Path = ""
	gt.AppPath = ""
	gt.Path = path
	gt.CreateInitTemplate()
	gt.CreateAppTemplate()
	gt.CreateDockerComposeTemplate()
	gt.CreateDockerfileTemplate()
	gt.CreateEnvTemplate()
	gt.CreateGitignoreTemplate()
	gt.CreateMainTemplate()
}

func (gt *GoTemplate) CreateInitTemplate() {
	gt.AppPath = fmt.Sprintf("%v/app", gt.Path)
	folderInapp := [...]string{"bootstrap", "controller", "domain", "infrastructure", "interfaces", "middleware", "repository", "routes", "rules", "service", "utils"}
	folderInUtils := [...]string{"encrypt", "jwt", "response", "session", "interfaces", "validation"}

	err := os.RemoveAll(gt.Path)
	if err != nil {
		fmt.Printf(": %v", err)
	}
	errCreateFolderBackend := os.Mkdir(gt.Path, 0755)
	if errCreateFolderBackend != nil {
		log.Println(errCreateFolderBackend)
	}

	errCreateFolderApp := os.Mkdir(gt.AppPath, 0755)
	if errCreateFolderApp != nil {
		log.Println(errCreateFolderApp)
	}

	for _, v := range folderInapp {
		err := os.Mkdir(gt.AppPath+"/"+v, 0755)
		if err != nil {
			log.Println(err)
		}
	}

	for _, v := range folderInUtils {
		err := os.Mkdir(gt.AppPath+"/utils/"+v, 0755)
		if err != nil {
			log.Println(err)
		}
	}
}

func (gt *GoTemplate) CreateAppTemplate() {
	// ------------------------------------------------------------------------------------------------------------
	// bootstrap
	f, ferr := os.Create(gt.AppPath + "/bootstrap/app.go")
	if ferr != nil {
		log.Println(ferr)
	}

	app.AppTemplate.Execute(f, struct{}{})
	// ------------------------------------------------------------------------------------------------------------
	// infrastructure

	f, ferr = os.Create(gt.AppPath + "/infrastructure/casbin.go")
	if ferr != nil {
		log.Println(ferr)
	}

	app.CasbinTemplate.Execute(f, struct{}{})

	f, ferr = os.Create(gt.AppPath + "/infrastructure/env.go")
	if ferr != nil {
		log.Println(ferr)
	}

	app.ENVTemplate.Execute(f, struct{}{})

	f, ferr = os.Create(gt.AppPath + "/infrastructure/logger.go")
	if ferr != nil {
		log.Println(ferr)
	}

	app.LoggerTemplate.Execute(f, struct{}{})

	f, ferr = os.Create(gt.AppPath + "/infrastructure/mongohandler.go")
	if ferr != nil {
		log.Println(ferr)
	}

	app.MongoTemplate.Execute(f, struct{}{})

	f, ferr = os.Create(gt.AppPath + "/infrastructure/rabbitmq.go")
	if ferr != nil {
		log.Println(ferr)
	}

	app.RabbitTemplate.Execute(f, struct{}{})

	f, ferr = os.Create(gt.AppPath + "/infrastructure/redis.go")
	if ferr != nil {
		log.Println(ferr)
	}

	app.RedisTemplate.Execute(f, struct{}{})
	// ------------------------------------------------------------------------------------------------------------
	// interfaces

	f, ferr = os.Create(gt.AppPath + "/interfaces/logger.go")
	if ferr != nil {
		log.Println(ferr)
	}

	app.InterfaceLoggerTemplate.Execute(f, struct{}{})
	// ------------------------------------------------------------------------------------------------------------
	// middleware

	f, ferr = os.Create(gt.AppPath + "/middleware/acl.go")
	if ferr != nil {
		log.Println(ferr)
	}

	app.ACLTemplate.Execute(f, struct{}{})

	f, ferr = os.Create(gt.AppPath + "/middleware/jwt.go")
	if ferr != nil {
		log.Println(ferr)
	}

	app.JWTTemplate.Execute(f, struct{}{})
	// ------------------------------------------------------------------------------------------------------------
	// repository

	f, ferr = os.Create(gt.AppPath + "/middleware/redisrepository.go")
	if ferr != nil {
		log.Println(ferr)
	}

	app.RedisRepoTemplate.Execute(f, struct{}{})
	// ------------------------------------------------------------------------------------------------------------
	// service

	f, ferr = os.Create(gt.AppPath + "/middleware/redisservice.go")
	if ferr != nil {
		log.Println(ferr)
	}

	app.RedisSVTemplate.Execute(f, struct{}{})
	// ------------------------------------------------------------------------------------------------------------
	// utils

	f, ferr = os.Create(gt.AppPath + "/utils/encrypt/argon2.go")
	if ferr != nil {
		log.Println(ferr)
	}

	app.Argon2Template.Execute(f, struct{}{})

	f, ferr = os.Create(gt.AppPath + "/utils/jwt/jwt.go")
	if ferr != nil {
		log.Println(ferr)
	}

	app.JWTUTILTemplate.Execute(f, struct{}{})

	f, ferr = os.Create(gt.AppPath + "/utils/response/response.go")
	if ferr != nil {
		log.Println(ferr)
	}

	app.ResTemplate.Execute(f, struct{}{})

	f, ferr = os.Create(gt.AppPath + "/utils/session/session.go")
	if ferr != nil {
		log.Println(ferr)
	}

	app.SessionTemplate.Execute(f, struct{}{})

	f, ferr = os.Create(gt.AppPath + "/utils/validation/validation.go")
	if ferr != nil {
		log.Println(ferr)
	}

	app.ValidateTemplate.Execute(f, struct{}{})

}

func (gt *GoTemplate) CreateDockerComposeTemplate() {
	f, ferr := os.Create(gt.Path + "/docker-compose.yml")
	if ferr != nil {
		log.Println(ferr)
	}

	app.DockerComposeTemplate.Execute(f, struct{}{})
}

func (gt *GoTemplate) CreateDockerfileTemplate() {
	f, ferr := os.Create(gt.Path + "/Dockerfile")
	if ferr != nil {
		log.Println(ferr)
	}

	app.DockerfileTemplate.Execute(f, struct{}{})
}

func (gt *GoTemplate) CreateEnvTemplate() {
	f, ferr := os.Create(gt.Path + "/.env")
	if ferr != nil {
		log.Println(ferr)
	}

	app.ENVCTemplate.Execute(f, struct{}{})
}

func (gt *GoTemplate) CreateGitignoreTemplate() {
	f, ferr := os.Create(gt.Path + "/.gitignore")
	if ferr != nil {
		log.Println(ferr)
	}

	app.GitignoreTemplate.Execute(f, struct{}{})
}

func (gt *GoTemplate) CreateMainTemplate() {
	log.Println(gt.Path)
	f, ferr := os.Create(gt.Path + "/main.go")
	if ferr != nil {
		log.Println(ferr)
	}

	gotemplate.MainTemplate.Execute(f, struct{}{})
}
