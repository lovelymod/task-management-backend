package repository

import (
	"context"
	"log"
	"time"

	"github.com/lovelymod/task-management-backend/internal/bootstrap"
	"github.com/lovelymod/task-management-backend/internal/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type authRepository struct {
	mc *bootstrap.MongoCollections
}

func NewAuthHandler(mc *bootstrap.MongoCollections) entity.AuthRepository {
	return &authRepository{mc: mc}
}

func (r *authRepository) GetUserByEmail(ctx context.Context, email string) ([]entity.User, error) {
	var exitingUser []entity.User

	filter := bson.D{{Key: "email", Value: email}}
	opts := options.Find().SetLimit(1)

	cursor, err := r.mc.Users.Find(ctx, filter, opts)
	if err != nil {
		log.Println(err)
		return nil, entity.ErrGlobalServerError
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &exitingUser); err != nil {
		log.Println(err)
		return nil, entity.ErrGlobalServerError
	}

	return exitingUser, nil
}

func (r *authRepository) CreateUser(ctx context.Context, registerUser *entity.User) (*entity.User, error) {
	if _, err := r.mc.Users.InsertOne(ctx, registerUser); err != nil {
		log.Println(err)
		return nil, entity.ErrGlobalServerError
	}

	return registerUser, nil
}

func (r *authRepository) CreateRefreshToken(ctx context.Context, refreshToken *entity.RefreshToken) error {
	if _, err := r.mc.RefreshTokens.InsertOne(ctx, refreshToken); err != nil {
		log.Println(err)
		return entity.ErrGlobalServerError
	}
	return nil
}

func (r *authRepository) RevokeRefreshToken(ctx context.Context, token string) error {
	filter := bson.D{{Key: "token", Value: token}}
	update := bson.M{
		"$set": bson.M{
			"is_revoked": true,
			"updated_at": time.Now(),
		},
	}

	if _, err := r.mc.RefreshTokens.UpdateOne(ctx, filter, update); err != nil {
		log.Println(err)
		return entity.ErrGlobalServerError
	}
	return nil
}

func (r *authRepository) GetRefreshToken(ctx context.Context, token string) (*entity.RefreshToken, error) {
	// สร้าง Pipeline สำหรับการทำ Aggregate (Join)
	pipeline := mongo.Pipeline{
		// 1. $match: กรองข้อมูลหา Token ที่ตรงกับเงื่อนไข (ทำงานคล้ายๆ FindOne)
		bson.D{{Key: "$match", Value: bson.D{{Key: "token", Value: token}}}},

		// 2. $lookup: สั่ง Join ไปที่ Collection "users"
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "users"},        // ชื่อ collection ที่เราจะไปดึงข้อมูลมา (ต้องเป็นชื่อใน DB)
			{Key: "localField", Value: "user_id"}, // ฟิลด์ในตาราง refresh_tokens ปัจจุบัน
			{Key: "foreignField", Value: "_id"},  // ฟิลด์ที่จะไปเทียบในตาราง users
			{Key: "as", Value: "user"},           // ชื่อฟิลด์ผลลัพธ์ที่จะเก็บลงใน Struct (ต้องตรงกับ bson:"user" ด้านบน)
		}}},

		// 3. $unwind: แปลงผลลัพธ์ User จากรูปแบบ Array ให้กลายเป็น Object ตัวเดียว
		bson.D{{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: "$user"},
			{Key: "preserveNullAndEmptyArrays", Value: true}, // ใส่ไว้กันพัง กรณีที่หา User ไม่เจอ
		}}},
	}

	// ใช้คำสั่ง Aggregate แทน FindOne
	cursor, err := r.mc.RefreshTokens.Aggregate(ctx, pipeline)
	if err != nil {
		log.Println(err)
		return nil, entity.ErrGlobalServerError
	}
	defer cursor.Close(ctx) // อย่าลืมปิด cursor เสมอ

	// เนื่องจาก Aggregate คืนค่าเป็น Cursor เราต้องใช้ cursor.Next() เพื่อเช็กว่ามีผลลัพธ์ไหม
	if cursor.Next(ctx) {
		var existingRefreshToken entity.RefreshToken
		if err := cursor.Decode(&existingRefreshToken); err != nil {
			log.Println(err)
			return nil, entity.ErrGlobalServerError
		}
		return &existingRefreshToken, nil
	}

	// ถ้าวิ่งมาถึงตรงนี้ แปลว่าหา Token ไม่เจอเลย
	return nil, mongo.ErrNoDocuments
}
