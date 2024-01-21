package repo

import (
	"anonymous.net.cn/public-system-backend/config"
	"anonymous.net.cn/public-system-backend/db"
	commonModel "anonymous.net.cn/public-system-backend/module/common/model"
	publicModel "anonymous.net.cn/public-system-backend/module/public/model"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const (
	DocIcpLicence       = "doc_icp_licence"       // IcpLicence 备案信息文档
	DocCopyrightLicence = "doc_copyright_licence" // CopyrightLicence 版权信息文档
	DocTechStack        = "doc_tech_stack"        // TechStack 技术栈信息文档
	DocContactMe        = "doc_contact_me"        // ContactMe 联系方式信息文档
	DocCodePlace        = "doc_code_place"        // CodePlace 代码空间信息文档
	DocSiteEntrance     = "doc_site_entrance"     // SiteEntrance 站点信息文档
)

// GetDisplayIcpLicence 获取 icp 备案展示信息
func GetDisplayIcpLicence() (*publicModel.IcpLicence, error) {
	s, err := db.MongoClient.StartSession()
	if err != nil {
		log.Println("获取数据库会话失败: ", err.Error())
		return nil, err
	}
	defer s.EndSession(context.Background())
	collection := s.Client().Database(config.Config.DB.Mongo.DB).Collection(DocIcpLicence)
	licence := &publicModel.IcpLicence{}
	err = collection.FindOne(context.Background(), bson.M{}).Decode(licence)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		} else {
			return nil, err
		}
	} else {
		return licence, nil
	}
}

// GetDisplayCopyrightLicence 获取版权展示信息
func GetDisplayCopyrightLicence() (*publicModel.CopyrightLicence, error) {
	s, err := db.MongoClient.StartSession()
	if err != nil {
		log.Println("获取数据库会话失败: ", err.Error())
		return nil, err
	}
	defer s.EndSession(context.Background())
	collection := s.Client().Database(config.Config.DB.Mongo.DB).Collection(DocCopyrightLicence)
	licence := &publicModel.CopyrightLicence{}
	err = collection.FindOne(context.Background(), bson.M{}).Decode(licence)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		} else {
			return nil, err
		}
	} else {
		return licence, nil
	}
}

// GetDisplayTechStack 获取技术栈展示信息
func GetDisplayTechStack(limit int64) (*commonModel.Display[publicModel.TechStack], error) {
	s, err := db.MongoClient.StartSession()
	if err != nil {
		log.Println("获取数据库会话失败: ", err.Error())
		return nil, err
	}
	defer s.EndSession(context.Background())
	skip := int64(0)
	res := &commonModel.Display[publicModel.TechStack]{}
	res.Size = uint64(limit)
	records := make([]publicModel.TechStack, 0)
	collection := s.Client().Database(config.Config.DB.Mongo.DB).Collection(DocTechStack)
	filter := bson.M{}
	sort := bson.D{
		bson.E{Key: "weight", Value: -1},
		bson.E{Key: "name", Value: 1},
	}
	opts := options.FindOptions{
		Sort:  sort,
		Limit: &limit,
		Skip:  &skip,
	}
	ctx := context.Background()
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}
	res.Total = uint64(count)
	cursor, err := collection.Find(ctx, filter, &opts)
	if err != nil {
		return nil, err
	}
	if cursor == nil {
		res.Records = records
		return res, nil
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Println(err.Error())
		}
	}(cursor, ctx)
	for cursor.Next(context.Background()) {
		var tmp publicModel.TechStack
		if err = cursor.Decode(&tmp); err != nil {
			return nil, err
		}
		records = append(records, tmp)
	}
	res.Records = records
	return res, nil
}

// GetDisplayContactMe 获取联系方式展示信息
func GetDisplayContactMe(limit int64) (*commonModel.Display[publicModel.Contact], error) {
	s, err := db.MongoClient.StartSession()
	if err != nil {
		log.Println("获取数据库会话失败: ", err.Error())
		return nil, err
	}
	defer s.EndSession(context.Background())
	skip := int64(0)
	res := &commonModel.Display[publicModel.Contact]{}
	res.Size = uint64(limit)
	records := make([]publicModel.Contact, 0)
	collection := s.Client().Database(config.Config.DB.Mongo.DB).Collection(DocContactMe)
	filter := bson.M{}
	sort := bson.D{
		bson.E{Key: "weight", Value: -1},
		bson.E{Key: "name", Value: 1},
	}
	opts := options.FindOptions{
		Sort:  sort,
		Limit: &limit,
		Skip:  &skip,
	}
	ctx := context.Background()
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}
	res.Total = uint64(count)
	cursor, err := collection.Find(ctx, filter, &opts)
	if err != nil {
		return nil, err
	}
	if cursor == nil {
		res.Records = records
		return res, nil
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Println(err.Error())
		}
	}(cursor, ctx)
	for cursor.Next(context.Background()) {
		var tmp publicModel.Contact
		if err = cursor.Decode(&tmp); err != nil {
			return nil, err
		}
		records = append(records, tmp)
	}
	res.Records = records
	return res, nil
}

// GetDisplayCodePlace 获取代码空间展示信息
func GetDisplayCodePlace(limit int64) (*commonModel.Display[publicModel.Contact], error) {
	s, err := db.MongoClient.StartSession()
	if err != nil {
		log.Println("获取数据库会话失败: ", err.Error())
		return nil, err
	}
	defer s.EndSession(context.Background())
	skip := int64(0)
	res := &commonModel.Display[publicModel.Contact]{}
	res.Size = uint64(limit)
	records := make([]publicModel.Contact, 0)
	collection := s.Client().Database(config.Config.DB.Mongo.DB).Collection(DocCodePlace)
	filter := bson.M{}
	sort := bson.D{
		bson.E{Key: "weight", Value: -1},
		bson.E{Key: "name", Value: 1},
	}
	opts := options.FindOptions{
		Sort:  sort,
		Limit: &limit,
		Skip:  &skip,
	}
	ctx := context.Background()
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}
	res.Total = uint64(count)
	cursor, err := collection.Find(ctx, filter, &opts)
	if err != nil {
		return nil, err
	}
	if cursor == nil {
		res.Records = records
		return res, nil
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Println(err.Error())
		}
	}(cursor, ctx)
	for cursor.Next(context.Background()) {
		var tmp publicModel.Contact
		if err = cursor.Decode(&tmp); err != nil {
			return nil, err
		}
		records = append(records, tmp)
	}
	res.Records = records
	return res, nil
}

// GetDisplaySiteEntrance 获取站点入口展示信息
func GetDisplaySiteEntrance(limit int64) (*commonModel.Display[publicModel.SiteEntrance], error) {
	s, err := db.MongoClient.StartSession()
	if err != nil {
		log.Println("获取数据库会话失败: ", err.Error())
		return nil, err
	}
	defer s.EndSession(context.Background())
	skip := int64(0)
	res := &commonModel.Display[publicModel.SiteEntrance]{}
	res.Size = uint64(limit)
	records := make([]publicModel.SiteEntrance, 0)
	collection := s.Client().Database(config.Config.DB.Mongo.DB).Collection(DocSiteEntrance)
	filter := bson.M{}
	sort := bson.D{
		bson.E{Key: "weight", Value: -1},
		bson.E{Key: "name", Value: 1},
	}
	opts := options.FindOptions{
		Sort:  sort,
		Limit: &limit,
		Skip:  &skip,
	}
	ctx := context.Background()
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}
	res.Total = uint64(count)
	cursor, err := collection.Find(ctx, filter, &opts)
	if err != nil {
		return nil, err
	}
	if cursor == nil {
		res.Records = records
		return res, nil
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Println(err.Error())
		}
	}(cursor, ctx)
	for cursor.Next(context.Background()) {
		var tmp publicModel.SiteEntrance
		if err = cursor.Decode(&tmp); err != nil {
			return nil, err
		}
		records = append(records, tmp)
	}
	res.Records = records
	return res, nil
}
