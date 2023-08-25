package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
	"unicode"
)

// RegisterRoutes 函数用于将控制器的方法自动注册到路由
func RegisterRoutes(router *gin.RouterGroup, ctrl interface{}) {

	t := reflect.TypeOf(ctrl) // 获取控制器的反射类型
	name := t.Name()
	if t.Kind() == reflect.Ptr { // 检查类型是否为指针
		name = t.Elem().Name() // 如果是指针，则获取指针指向的类型
	}
	v := reflect.New(t.Elem())      // 创建控制器的新实例
	prefix := strings.ToLower(name) // 将控制器名转换为小写，用作 URL 前缀
	// 遍历控制器的所有方法
	for i := 0; i < t.NumMethod(); i++ {
		state := false
		methodName := t.Method(i).Name
		httpMethod := extractHTTPMethod(methodName)                      // 提取 HTTP 方法，例如 "GET"
		routeSuffix := generateRouteSuffix(methodName[len(httpMethod):]) // 生成路由后缀
		// 如果路由后缀以 "by" 结尾，则替换为 ":id" 占位符
		if strings.HasSuffix(routeSuffix, "by") {
			routeSuffix = strings.TrimSuffix(routeSuffix, "by") + ":id"
			state = true
		}
		// 如果成功提取了 HTTP 方法和路由后缀，则注册方法到路由
		if httpMethod != "" {
			route := ""
			if routeSuffix == "" {
				route = fmt.Sprintf("/%s", prefix)
			} else {
				route = fmt.Sprintf("/%s/%s", prefix, routeSuffix)
			}
			router.Handle(httpMethod, route, func(ctx *gin.Context) {
				// 设置 Ctx 字段的值
				v.Elem().FieldByName("Ctx").Set(reflect.ValueOf(ctx))
				method := v.MethodByName(methodName) // 根据方法名获取方法
				var resultValues []reflect.Value
				if state {
					id := ctx.Param("id")
					resultValues = method.Call([]reflect.Value{reflect.ValueOf(id)})
				} else {
					resultValues = method.Call(nil)
				}
				if len(resultValues) > 0 {
					firstReturnValue := resultValues[0].Interface()
					ctx.JSON(200, firstReturnValue)
				}
			})
		}
	}
}

// extractHTTPMethod 函数根据方法名提取 HTTP 方法
func extractHTTPMethod(methodName string) string {
	switch {
	case strings.HasPrefix(methodName, "Get"):
		return "GET"
	case strings.HasPrefix(methodName, "Post"):
		return "POST"
	case strings.HasPrefix(methodName, "Put"):
		return "PUT"
	case strings.HasPrefix(methodName, "Delete"):
		return "DELETE"
	case strings.HasPrefix(methodName, "Patch"):
		return "PATCH"
	default:
		return ""
	}
}

// generateRouteSuffix 函数将驼峰命名格式的方法名转换为 URL 路由格式
func generateRouteSuffix(methodName string) string {
	var result []rune
	for _, char := range methodName {
		if unicode.IsUpper(char) {
			result = append(result, '/')
			result = append(result, unicode.ToLower(char))
		} else {
			result = append(result, char)
		}
	}
	return string(result)
}
