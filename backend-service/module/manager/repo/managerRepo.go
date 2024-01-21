package repo

import (
	"anonymous.net.cn/public-system-backend/config"
	"anonymous.net.cn/public-system-backend/db"
	publicModel "anonymous.net.cn/public-system-backend/module/public/model"
	"anonymous.net.cn/public-system-backend/module/public/repo"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// ========== 备案信息存取 ========== //

// SetIcpLicence 设置 icp 备案信息
func SetIcpLicence(icp publicModel.IcpLicence) (primitive.ObjectID, error) {
	licenceExist, err := isIcpLicenceExist()
	if err != nil {
		return primitive.NilObjectID, err
	}
	if licenceExist {
		return modifiedIcpLicence(icp)
	}
	return createIcpLicence(icp)
}

// isIcpLicenceExist icp 备案信息是否存在
func isIcpLicenceExist() (bool, error) {
	s, err := db.MongoClient.StartSession()
	if err != nil {
		log.Println("获取数据库会话失败: ", err.Error())
		return false, err
	}
	defer s.EndSession(context.Background())
	collection := s.Client().Database(config.Config.DB.Mongo.DB).Collection(repo.DocIcpLicence)
	count, err := collection.CountDocuments(context.Background(), bson.M{}, options.Count())
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// createIcpLicence 创建 icp 备案信息
func createIcpLicence(icp publicModel.IcpLicence) (primitive.ObjectID, error) {
	s, err := db.MongoClient.StartSession()
	if err != nil {
		log.Println("获取数据库会话失败: ", err.Error())
		return primitive.NilObjectID, err
	}
	defer s.EndSession(context.Background())
	collection := s.Client().Database(config.Config.DB.Mongo.DB).Collection(repo.DocIcpLicence)
	icp.ID = primitive.NewObjectID()
	icp.GmtCreate = time.Now()
	_, err = collection.InsertOne(context.Background(), icp, options.InsertOne())
	if err != nil {
		return primitive.NilObjectID, err
	}
	return icp.ID, nil
}

// modifiedIcpLicence 修改 icp 备案信息
func modifiedIcpLicence(icp publicModel.IcpLicence) (primitive.ObjectID, error) {
	s, err := db.MongoClient.StartSession()
	if err != nil {
		log.Println("获取数据库会话失败: ", err.Error())
		return primitive.NilObjectID, err
	}
	defer s.EndSession(context.Background())
	collection := s.Client().Database(config.Config.DB.Mongo.DB).Collection(repo.DocIcpLicence)
	licence := &publicModel.IcpLicence{}
	err = collection.FindOne(context.Background(), bson.M{}).Decode(licence)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return primitive.NilObjectID, nil
		} else {
			return primitive.NilObjectID, err
		}
	}
	modified := false
	if icp.Province != "" && icp.Province != licence.Province {
		licence.Province = icp.Province
		modified = true
	}
	if icp.Number != "" && icp.Number != licence.Number {
		licence.Number = icp.Number
		if !modified {
			modified = true
		}
	}
	if icp.No != "" && icp.No != licence.No {
		licence.No = icp.No
		if !modified {
			modified = true
		}
	}
	if modified {
		licence.GmtModified = time.Now()
		filter := bson.D{{"_id", licence.ID}}
		update := bson.D{{"$set", bson.D{{"province", licence.Province}, {"number", licence.Number}, {"no", licence.No}}}}
		_, err = collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return primitive.NilObjectID, err
		}
	}
	return licence.ID, nil
}

// DeleteIcpLicence 删除 icp 备案信息
func DeleteIcpLicence() (int64, error) {
	s, err := db.MongoClient.StartSession()
	if err != nil {
		log.Println("获取数据库会话失败: ", err.Error())
		return 0, err
	}
	defer s.EndSession(context.Background())
	opts := &options.DeleteOptions{}
	collection := s.Client().Database(config.Config.DB.Mongo.DB).Collection(repo.DocIcpLicence)
	res, err := collection.DeleteMany(context.Background(), bson.M{}, opts)
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, nil
}

// ========== 版权信息存取 ========== //

// SetCopyrightLicence 设置版权信息
func SetCopyrightLicence(copyright publicModel.CopyrightLicence) (primitive.ObjectID, error) {
	licenceExist, err := isCopyrightLicenceExist()
	if err != nil {
		return primitive.NilObjectID, nil
	}
	if licenceExist {
		return modifiedCopyrightLicence(copyright)
	}
	return createCopyrightLicence(copyright)
}

// isCopyrightExist 版权信息是否存在
func isCopyrightLicenceExist() (bool, error) {
	s, err := db.MongoClient.StartSession()
	if err != nil {
		log.Println("获取数据库会话失败: ", err.Error())
		return false, err
	}
	defer s.EndSession(context.Background())
	collection := s.Client().Database(config.Config.DB.Mongo.DB).Collection(repo.DocCopyrightLicence)
	count, err := collection.CountDocuments(context.Background(), bson.M{}, options.Count())
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// createCopyrightLicence 创建版权信息
func createCopyrightLicence(copyright publicModel.CopyrightLicence) (primitive.ObjectID, error) {
	s, err := db.MongoClient.StartSession()
	if err != nil {
		log.Println("获取数据库会话失败: ", err.Error())
		return primitive.NilObjectID, err
	}
	defer s.EndSession(context.Background())
	collection := s.Client().Database(config.Config.DB.Mongo.DB).Collection(repo.DocCopyrightLicence)
	copyright.ID = primitive.NewObjectID()
	copyright.GmtCreate = time.Now()
	_, err = collection.InsertOne(context.Background(), copyright, options.InsertOne())
	if err != nil {
		return primitive.NilObjectID, err
	}
	return copyright.ID, nil
}

// modifiedCopyrightLicence 修改版权信息
func modifiedCopyrightLicence(copyright publicModel.CopyrightLicence) (primitive.ObjectID, error) {
	s, err := db.MongoClient.StartSession()
	if err != nil {
		log.Println("获取数据库会话失败: ", err.Error())
		return primitive.NilObjectID, err
	}
	defer s.EndSession(context.Background())
	collection := s.Client().Database(config.Config.DB.Mongo.DB).Collection(repo.DocCopyrightLicence)
	licence := &publicModel.CopyrightLicence{}
	err = collection.FindOne(context.Background(), bson.M{}).Decode(licence)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return primitive.NilObjectID, nil
		} else {
			return primitive.NilObjectID, err
		}
	}
	modified := false
	if copyright.From != "" && copyright.From != licence.From {
		licence.From = copyright.From
		modified = true
	}
	if copyright.To != "" && copyright.To != licence.To {
		licence.To = copyright.To
		if !modified {
			modified = true
		}
	}
	if copyright.Author != "" && copyright.Author != licence.Author {
		licence.Author = copyright.Author
		if !modified {
			modified = true
		}
	}
	if modified {
		licence.GmtModified = time.Now()
		filter := bson.D{{"_id", licence.ID}}
		update := bson.D{{"$set", bson.D{{"from", licence.From}, {"to", licence.To}, {"author", licence.Author}}}}
		_, err = collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return primitive.NilObjectID, err
		}
	}
	return licence.ID, nil
}

// DeleteCopyrightLicence 删除版权信息
func DeleteCopyrightLicence() (int64, error) {
	s, err := db.MongoClient.StartSession()
	if err != nil {
		log.Println("获取数据库会话失败: ", err.Error())
		return 0, err
	}
	defer s.EndSession(context.Background())
	opts := &options.DeleteOptions{}
	collection := s.Client().Database(config.Config.DB.Mongo.DB).Collection(repo.DocCopyrightLicence)
	res, err := collection.DeleteMany(context.Background(), bson.M{}, opts)
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, nil
}

// ========== 技术栈信息存取 ========== //

// AddTechStack 添加技术栈
func AddTechStack(techStack publicModel.TechStack) (primitive.ObjectID, error) {
	s, err := db.MongoClient.StartSession()
	if err != nil {
		log.Println("获取数据库会话失败: ", err.Error())
		return primitive.NilObjectID, err
	}
	defer s.EndSession(context.Background())
	collection := s.Client().Database(config.Config.DB.Mongo.DB).Collection(repo.DocTechStack)
	techStack.ID = primitive.NewObjectID()
	techStack.GmtCreate = time.Now()
	_, err = collection.InsertOne(context.Background(), techStack, options.InsertOne())
	if err != nil {
		return primitive.NilObjectID, err
	}
	return techStack.ID, nil
}

// DeleteTechStack 删除技术栈
func DeleteTechStack() {

}

// ModifiedTechStack 修改技术栈
func ModifiedTechStack() {

}

// ========== 联系方式信息存取 ========== //

// AddContactMe 添加联系方式
func AddContactMe() {

}

// DeleteContactMe 删除联系方式
func DeleteContactMe() {

}

// ModifiedContactMe 修改联系方式
func ModifiedContactMe() {

}

// ========== 代码空间信息存取 ========== //

// AddCodePlace 添加代码空间
func AddCodePlace() {

}

// DeleteCodePlace 删除代码空间
func DeleteCodePlace() {

}

// ModifiedCodePlace 修改代码空间
func ModifiedCodePlace() {

}

// ========== 站点入口信息存取 ========== //

// AddSiteEntrance 添加站点入口
func AddSiteEntrance() {

}

// DeleteSiteEntrance 删除站点入口
func DeleteSiteEntrance() {

}

// ModifiedSiteEntrance 修改站点入口
func ModifiedSiteEntrance() {

}
