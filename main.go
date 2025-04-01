package main

import (
	"chat_upgrade/controller"
	"chat_upgrade/db"
	"chat_upgrade/model"
	"chat_upgrade/repository"
	"chat_upgrade/router"
	"chat_upgrade/usecase"
	"chat_upgrade/validator"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// ✅ .env ファイルの読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// ✅ AWS 設定の読み込み
	config := model.Config{
		AWSAccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
		AWSSecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
		AWSRegion:          os.Getenv("AWS_REGION"),
		S3Bucket:           os.Getenv("S3_BUCKET"),
	}

	// データベースの初期化
	database := db.NewDB() // ここは *gorm.DB を返すように設定されていると仮定

	// バリデーターの初期化
	userValidator := validator.NewUserValidator()

	// ユーザー関連のリポジトリ、ユースケース、コントローラーの初期化
	userRepository := repository.NewUserRepository(database)
	s3Repository := repository.NewS3Repository(config.S3Bucket)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator, s3Repository)
	userController := controller.NewUserController(userUsecase)

	// 経験情報関連のリポジトリ、ユースケース、コントローラーの初期化
	experienceRepository := repository.NewExperienceRepository(database)
	experienceUsecase := usecase.NewExperienceUsecase(experienceRepository, s3Repository)
	experienceController := controller.NewExperienceController(experienceUsecase)

	// Career関連のリポジトリ、ユースケース、コントローラーの初期化
	careerRepository := repository.NewCareerRepository(database)
	careerUsecase := usecase.NewCareerUsecase(careerRepository)
	careerController := controller.NewCareerController(careerUsecase)

	// ルーターの設定
	e := router.NewRouter(userController, experienceController, careerController)

	port := os.Getenv("PORT")
if port == "" {
  port = "8080" 
}
e.Logger.Fatal(e.Start(":" + port))

}
